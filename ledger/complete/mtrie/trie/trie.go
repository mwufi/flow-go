package trie

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common"
	"github.com/onflow/flow-go/ledger/common/utils"
	"github.com/onflow/flow-go/ledger/complete/mtrie/node"
)

// MTrie represents a perfect in-memory full binary Merkle tree with uniform height.
// For a detailed description of the storage model, please consult `mtrie/README.md`
//
// A MTrie is a thin wrapper around a the trie's root Node. An MTrie implements the
// logic for forming MTrie-graphs from the elementary nodes. Specifically:
//   * how Nodes (graph vertices) form a Trie,
//   * how register values are read from the trie,
//   * how Merkle proofs are generated from a trie, and
//   * how a new Trie with updated values is generated.
//
// `MTrie`s are _immutable_ data structures. Updating register values is implemented through
// copy-on-write, which creates a new `MTrie`. For minimal memory consumption, all sub-tries
// that where not affected by the write operation are shared between the original MTrie
// (before the register updates) and the updated MTrie (after the register writes).
//
// DEFINITIONS and CONVENTIONS:
//   * HEIGHT of a node v in a tree is the number of edges on the longest downward path
//     between v and a tree leaf. The height of a tree is the height of its root.
//     The height of a Trie is always the height of the fully-expanded tree.
type MTrie struct {
	root         *node.Node
	pathByteSize int
}

// NewEmptyMTrie returns an empty Mtrie (root is an empty node)
func NewEmptyMTrie(pathByteSize int) (*MTrie, error) {
	if pathByteSize < 1 {
		return nil, errors.New("trie's path size [in bytes] must be positive")
	}
	height := pathByteSize * 8
	return &MTrie{
		root:         node.NewEmptyTreeRoot(height),
		pathByteSize: pathByteSize,
	}, nil
}

// NewMTrie returns a Mtrie given the root
func NewMTrie(root *node.Node) (*MTrie, error) {
	if root.Height()%8 != 0 {
		return nil, errors.New("height of root node must be integer-multiple of 8")
	}
	pathByteSize := root.Height() / 8
	return &MTrie{
		root:         root,
		pathByteSize: pathByteSize,
	}, nil
}

// Height returns the height of the trie, which
// is the height of its root node.
func (mt *MTrie) Height() int { return mt.root.Height() }

// StringRootHash returns the trie's Hex-encoded root hash.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) StringRootHash() string { return hex.EncodeToString(mt.root.Hash()) }

// RootHash returns the trie's root hash (i.e. the hash of the trie's root node).
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) RootHash() []byte { return mt.root.Hash() }

// PathLength return the length [in bytes] the trie operates with.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) PathLength() int { return mt.pathByteSize }

// AllocatedRegCount returns the number of allocated registers in the trie.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) AllocatedRegCount() uint64 { return mt.root.RegCount() }

// MaxDepth returns the length of the longest branch from root to leaf.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) MaxDepth() uint16 { return mt.root.MaxDepth() }

// RootNode returns the Trie's root Node
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) RootNode() *node.Node {
	return mt.root
}

// StringRootHash returns the trie's string representation.
// Concurrency safe (as Tries are immutable structures by convention)
func (mt *MTrie) String() string {
	trieStr := fmt.Sprintf("Trie root hash: %v\n", mt.StringRootHash())
	return trieStr + mt.root.FmtStr("", "")
}

// UnsafeRead read payloads for the given paths.
// UNSAFE: requires _all_ paths to have a length of mt.Height bits.
// CAUTION: while reading the payloads, `paths` is permuted IN-PLACE for optimized processing.
// Return:
//  * `payloads` []*ledger.Payload
//     For each path, the corresponding payload is written into payloads. AFTER
//     the read operation completes, the order of `path` and `payloads` are such that
//     for `path[i]` the corresponding register value is referenced by 0`payloads[i]`.
// TODO move consistency checks from Forest into Trie to obtain a safe, self-contained API
func (mt *MTrie) UnsafeRead(paths []ledger.Path) []*ledger.Payload {
	payloads := make([]*ledger.Payload, len(paths)) // pre-allocate slice for the result
	mt.read(mt.root, paths, payloads)
	return payloads
}

// read reads all the registers in subtree with `head` as root node. For each
// `path[i]`, the corresponding payload is written into `payloads[i]` for the same index `i`.
// CAUTION:
//  * while reading the payloads, `paths` is permuted IN-PLACE for optimized processing.
//  * unchecked requirement: all paths must go through the `head` node
func (mt *MTrie) read(head *node.Node, paths []ledger.Path, payloads []*ledger.Payload) {
	// check for empty paths
	if len(paths) == 0 {
		return
	}

	// path not found
	if head == nil {
		for i := range paths {
			payloads[i] = ledger.EmptyPayload()
		}
		return
	}
	// reached a leaf node
	if head.IsLeaf() {
		for i, p := range paths {
			if bytes.Equal(head.Path(), p) {
				payloads[i] = head.Payload()
			} else {
				payloads[i] = ledger.EmptyPayload()
			}
		}
		return
	}

	// partition step to quick sort the paths:
	// lpaths contains all paths that have `0` at the partitionIndex
	// rpaths contains all paths that have `1` at the partitionIndex
	heightIndex := mt.Height() - head.Height() // distance to the tree root
	partitionIndex := utils.SplitPaths(paths, heightIndex)
	lpaths, rpaths := paths[:partitionIndex], paths[partitionIndex:]
	lpayloads, rpayloads := payloads[:partitionIndex], payloads[partitionIndex:]

	// read values from left and right subtrees in parallel
	parallelRecursionThreshold := 32 // threshold to avoid the parallelization going too deep in the recursion
	if len(lpaths) <= parallelRecursionThreshold {
		mt.read(head.LeftChild(), lpaths, lpayloads)
		mt.read(head.RightChild(), rpaths, rpayloads)
	} else {
		// concurrent read of left and right subtree
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			mt.read(head.LeftChild(), lpaths, lpayloads)
			wg.Done()
		}()
		mt.read(head.RightChild(), rpaths, rpayloads)
		wg.Wait() // wait for all threads
	}
}

// NewTrieWithUpdatedRegisters constructs a new trie containing all registers from the parent trie.
// The key-value pairs specify the registers whose values are supposed to hold updated values
// compared to the parent trie. Constructing the new trie is done in a COPY-ON-WRITE manner:
//   * The original trie remains unchanged.
//   * subtries that remain unchanged are from the parent trie instead of copied.
// UNSAFE: method requires the following conditions to be satisfied:
//   * keys are NOT duplicated
//   * requires _all_ paths to have a length of mt.Height bits.
// CAUTION: `updatedPaths` and `updatedPayloads` are permuted IN-PLACE for optimized processing.
// TODO: move consistency checks from MForest to here, to make API safe and self-contained
func NewTrieWithUpdatedRegisters(parentTrie *MTrie, updatedPaths []ledger.Path, updatedPayloads []ledger.Payload) (*MTrie, error) {
	parentRoot := parentTrie.root
	updatedRoot := parentTrie.update(parentRoot.Height(), parentRoot, updatedPaths, updatedPayloads, nil)
	if parentRoot == updatedRoot {
		return parentTrie, nil
	}
	updatedTrie, err := NewMTrie(updatedRoot)
	if err != nil {
		return nil, fmt.Errorf("constructing updated trie failed: %w", err)
	}
	return updatedTrie, nil
}

// update traverses the subtree and updates the stored registers
// CAUTION: while updating, `paths` and `payloads` are permuted IN-PLACE for optimized processing.
// UNSAFE: method requires the following conditions to be satisfied:
//   * paths all share the same common prefix [0 : mt.maxHeight-1 - nodeHeight)
//     (excluding the bit at index headHeight)
//   * paths are NOT duplicated
func (parentTrie *MTrie) update(
	nodeHeight int, parentNode *node.Node,
	paths []ledger.Path, payloads []ledger.Payload, compactLeaf *node.Node,
) *node.Node {
	// No new paths to write
	if len(paths) == 0 {
		// check is a compactLeaf from a higher height is still left
		if compactLeaf != nil {
			return node.NewLeaf(compactLeaf.Path(), compactLeaf.Payload(), nodeHeight)
		}
		return parentNode
	}

	if len(paths) == 1 && parentNode == nil && compactLeaf == nil {
		return node.NewLeaf(paths[0], &payloads[0], nodeHeight)
	}

	if parentNode != nil && parentNode.IsLeaf() { // if we're here then compactLeaf == nil
		// check if the parent path node is among the updated paths
		parentPath := parentNode.Path()
		found := false
		for i, p := range paths {
			if bytes.Equal(p, parentPath) {
				// the case where the leaf can be reused
				if len(paths) == 1 {
					if !bytes.Equal(parentNode.Payload().Value, payloads[i].Value) {
						return node.NewLeaf(paths[i], &payloads[i], nodeHeight)
					}
					// avoid creating a new node when the same payload is written
					return parentNode
				}
				found = true
				break
			}
		}
		if !found {
			compactLeaf = parentNode
		}
	}

	// in the remaining code: len(paths)>1

	// Split paths and payloads to recurse:
	// lpaths contains all paths that have `0` at the partitionIndex
	// rpaths contains all paths that have `1` at the partitionIndex
	heightIndex := parentTrie.Height() - nodeHeight // distance to the tree root
	partitionIndex := utils.SplitByPath(paths, payloads, heightIndex)
	lpaths, rpaths := paths[:partitionIndex], paths[partitionIndex:]
	lpayloads, rpayloads := payloads[:partitionIndex], payloads[partitionIndex:]

	// check if there is a compact leaf that needs to get deep to height 0
	var lcompactLeaf, rcompactLeaf *node.Node
	if compactLeaf != nil {
		// if yes, check which branch it will go to.
		if utils.Bit(compactLeaf.Path(), parentTrie.Height()-nodeHeight) == 0 {
			lcompactLeaf = compactLeaf
		} else {
			rcompactLeaf = compactLeaf
		}
	}

	// set the parent node children
	var lchildParent, rchildParent *node.Node
	if parentNode != nil {
		lchildParent = parentNode.LeftChild()
		rchildParent = parentNode.RightChild()
	}

	// recurse over each branch
	var lChild, rChild *node.Node
	if len(lpayloads) == 0 || len(rpayloads) == 0 {
		// runtime optimization: if there are _no_ updates for either left or right sub-tree, proceed single-threaded
		lChild = parentTrie.update(nodeHeight-1, lchildParent, lpaths, lpayloads, lcompactLeaf)
		rChild = parentTrie.update(nodeHeight-1, rchildParent, rpaths, rpayloads, rcompactLeaf)
	} else {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			lChild = parentTrie.update(nodeHeight-1, lchildParent, lpaths, lpayloads, lcompactLeaf)
		}()
		rChild = parentTrie.update(nodeHeight-1, rchildParent, rpaths, rpayloads, rcompactLeaf)
		wg.Wait()
	}

	// mitigate storage exhaustion attack: avoids creating a new node when the exact same
	// payload is re-written at a register.
	if lChild == lchildParent && rChild == rchildParent {
		return parentNode
	}
	return node.NewInterimNode(nodeHeight, lChild, rChild)
}

// UnsafeProofs provides proofs for the given paths.
// CAUTION: while updating, `paths` and `proofs` are permuted IN-PLACE for optimized processing.
// UNSAFE: requires _all_ paths to have a length of mt.Height bits.
func (mt *MTrie) UnsafeProofs(paths []ledger.Path, proofs []*ledger.TrieProof) {
	mt.proofs(mt.root, paths, proofs)
}

// proofs traverses the subtree and stores proofs for the given register paths in
// the provided `proofs` slice
// CAUTION: while updating, `paths` and `proofs` are permuted IN-PLACE for optimized processing.
// UNSAFE: method requires the following conditions to be satisfied:
//   * paths all share the same common prefix [0 : mt.maxHeight-1 - nodeHeight)
//     (excluding the bit at index headHeight)
func (mt *MTrie) proofs(head *node.Node, paths []ledger.Path, proofs []*ledger.TrieProof) {
	// check for empty paths
	if len(paths) == 0 {
		return
	}

	// we've reached the end of a trie
	// and path is not found (noninclusion proof)
	if head == nil {
		// by default, proofs are non-inclusion proofs
		return
	}

	// we've reached a leaf
	if head.IsLeaf() {
		for i, path := range paths {
			// value matches (inclusion proof)
			if bytes.Equal(head.Path(), path) {
				proofs[i].Path = head.Path()
				proofs[i].Payload = head.Payload()
				proofs[i].Inclusion = true
			}
		}
		// by default, proofs are non-inclusion proofs
		return
	}

	// increment steps for all the proofs
	for _, p := range proofs {
		p.Steps++
	}

	// partition step to quick sort the paths:
	// lpaths contains all paths that have `0` at the partitionIndex
	// rpaths contains all paths that have `1` at the partitionIndex
	heightIndex := mt.Height() - head.Height() // distance to the tree root
	partitionIndex := utils.SplitTrieProofsByPath(paths, proofs, heightIndex)
	lpaths, rpaths := paths[:partitionIndex], paths[partitionIndex:]
	lproofs, rproofs := proofs[:partitionIndex], proofs[partitionIndex:]

	wg := sync.WaitGroup{}
	parallelRecursionThreshold := 128 // thresold to avoid the parallelization going too deep in the recursion

	if rChild := head.RightChild(); rChild != nil {
		// This code is necessary, because we do not remove nodes from the trie
		// when a register is deleted. Instead, we just set the respective leaf's
		// payload to empty. While this will cause the lead's hash to become the
		// default hash, the node itself remains as part of the trie.
		// TODO: On update, prune subtries which only contain empty registers.
		//       Then, a child is nil if and only if the subtrie is empty.
		nodeHash := rChild.Hash()
		isDef := bytes.Equal(nodeHash, common.GetDefaultHashForHeight(rChild.Height()))
		if !isDef { // in proofs, we only provide non-default value hashes
			for _, p := range lproofs {
				utils.SetBit(p.Flags, heightIndex)
				p.Interims = append(p.Interims, nodeHash)
			}
		}
	}
	if len(lpaths) > parallelRecursionThreshold {
		wg.Add(1)
		go func() {
			mt.proofs(head.LeftChild(), lpaths, lproofs)
			wg.Done()
		}()
	} else {
		mt.proofs(head.LeftChild(), lpaths, lproofs)
	}

	if lChild := head.LeftChild(); lChild != nil {
		nodeHash := lChild.Hash()
		isDef := bytes.Equal(nodeHash, common.GetDefaultHashForHeight(lChild.Height()))
		if !isDef { // in proofs, we only provide non-default value hashes
			for _, p := range rproofs {
				utils.SetBit(p.Flags, heightIndex)
				p.Interims = append(p.Interims, nodeHash)
			}
		}
	}
	mt.proofs(head.RightChild(), rpaths, rproofs)
	wg.Wait()
}

// Equals compares two tries for equality.
// Tries are equal iff they store the same data (i.e. root hash matches)
// and their number and height are identical
func (mt *MTrie) Equals(o *MTrie) bool {
	if o == nil {
		return false
	}
	return o.PathLength() == mt.PathLength() && bytes.Equal(o.RootHash(), mt.RootHash())
}

// DumpAsJSON dumps the trie key value pairs to a file having each key value pair as a json row
func (mt *MTrie) DumpAsJSON(w io.Writer) error {

	// Use encoder to prevent building entire trie in memory
	enc := json.NewEncoder(w)

	err := mt.dumpAsJSON(mt.root, enc)
	if err != nil {
		return err
	}

	return nil
}

func (mt *MTrie) dumpAsJSON(n *node.Node, encoder *json.Encoder) error {
	if n.IsLeaf() {
		err := encoder.Encode(n.Payload())
		if err != nil {
			return err
		}
		return nil
	}

	if lChild := n.LeftChild(); lChild != nil {
		err := mt.dumpAsJSON(lChild, encoder)
		if err != nil {
			return err
		}
	}

	if rChild := n.RightChild(); rChild != nil {
		err := mt.dumpAsJSON(rChild, encoder)
		if err != nil {
			return err
		}
	}
	return nil
}

// EmptyTrieRootHash returns the rootHash of an empty Trie for the specified path size [bytes]
func EmptyTrieRootHash(pathByteSize int) []byte {
	return node.NewEmptyTreeRoot(8 * pathByteSize).Hash()
}

// AllPayloads returns all payloads
func (mt *MTrie) AllPayloads() []ledger.Payload {
	return mt.root.AllPayloads()
}

// IsAValidTrie verifies the content of the trie for potential issues
func (mt *MTrie) IsAValidTrie() bool {
	// TODO add checks on the health of node max height ...
	return mt.root.VerifyCachedHash()
}
