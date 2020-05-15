// Code generated by go-bindata. DO NOT EDIT.
// sources:
// engine/execution/state/bootstrap/contracts/FeeContract.cdc (1.882kB)
// engine/execution/state/bootstrap/contracts/FlowToken.cdc (7.439kB)
// engine/execution/state/bootstrap/contracts/FungibleToken.cdc (7.083kB)
// engine/execution/state/bootstrap/contracts/ServiceAccount.cdc (1.973kB)

package bootstrap

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _contractsFeecontractCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\xcf\x6f\x9c\x3a\x10\xbe\xf3\x57\xcc\xe5\x49\x10\x65\xe1\x1d\x9e\xde\x21\xdd\x54\x55\x93\xac\xd4\x5b\x55\xa5\xb9\xcf\x9a\x21\xb8\x05\x1b\xd9\x66\xc9\x36\xca\xff\x5e\xf9\x17\xe0\x84\xaa\x6d\x2e\x0e\x9e\xf9\xe6\xd7\xf7\x8d\xb7\xba\xb8\xc8\xb2\x03\x11\xdc\x48\x61\x14\x32\x93\x65\xf7\x2d\xd7\xc0\xc2\x27\x0c\x4a\x9e\x78\x4d\x1a\x10\x4e\x38\x76\x06\xa6\x96\x14\x81\x91\xdf\x49\x68\x40\x45\xc0\x64\xd7\x11\x33\x54\x03\x5a\xaf\x86\xe8\x12\xd0\x80\x69\x09\x48\xd4\x99\x6c\x80\x06\xc9\x5a\x7b\x71\x76\x80\xe3\xa8\x04\xd5\x70\x3c\x3b\x1f\xac\x7b\x2e\x00\x19\x93\xa3\x30\x65\x96\x5d\x54\x59\xc6\xfb\x41\x2a\x03\x87\x51\x3c\xf2\x63\x47\xf7\x36\x19\x34\x4a\xf6\xf0\xef\xd3\x3f\x7a\x36\x77\x72\x7a\x6d\xca\x86\xf1\xb8\x14\x7f\x20\x8a\x7d\xc1\x73\x96\x01\x00\x54\x15\xdc\x9d\x48\xd8\xfa\xd0\x00\xd7\x40\x3d\x37\xb6\xf8\xa9\x25\xe1\x0a\x9a\xd1\x2d\x6a\xe0\x82\x1b\x8e\x1d\xff\x41\xb5\x83\xdb\xf0\xe4\xf0\xab\xd8\x9f\x16\xa7\xbc\xf8\xa3\x34\xcb\xf4\x6a\x1a\xa4\xe6\xd6\x62\x24\x20\x3c\xd8\x19\xbf\xcd\x74\x1b\xbd\x72\xec\xed\x9c\xae\xe0\xeb\x81\x3f\xfd\xff\xdf\xdf\x67\xd3\x46\xc9\xf3\x56\x33\x6e\x92\xfa\xa3\xe3\xe6\x97\x59\x3e\x2b\x7e\x42\x43\x5e\x0a\x97\x30\x71\xd3\xda\x28\x1d\x67\xb1\x11\x68\x46\xc1\x0c\x97\x22\x40\xdc\x81\x8c\x91\xd6\xb9\xa6\xae\x29\xe0\x84\xca\xe3\xaf\xe0\xc3\x4c\x61\xe9\x1b\x8f\x79\x42\xb0\x75\x8c\x70\xd8\x9a\x9b\x51\x44\x8f\xdc\x52\xff\x36\x50\x01\xcf\xce\xdb\xfe\x75\x64\xe0\x88\x1d\x0a\x46\x70\xed\xa4\x52\x86\xcf\xd9\xc5\x16\x56\xba\x9a\xca\x34\xee\x7e\xe7\x00\xc5\xec\x69\x07\xbb\x4d\x48\x88\xe9\x5d\x5f\xe6\x4e\xdc\x40\x55\xda\x08\x7c\x21\x2d\x47\xc5\x08\xe4\xf1\x1b\xb1\x40\x5b\xb2\x08\x1a\x18\x0a\x68\x65\xe7\x74\x61\x37\x26\xd0\x58\xbe\x1e\x86\x8a\xb1\x7c\xa6\xa8\xf3\x90\xc9\x22\x3d\xb3\xab\xdb\xb5\xc3\x2d\x19\x52\x3d\x17\xa4\x9d\xf8\xe3\xa0\x64\xe3\x3e\xed\xcb\xf0\xe0\xb9\xd6\x13\x0e\x3a\xbd\xf4\xfc\xa3\x58\xc7\xa3\x7e\x30\xe7\x88\x09\x72\x7b\x0d\x43\x51\xbb\x41\x6a\x40\x11\x04\xe8\x22\x99\x96\x92\xda\x43\x2d\x6e\x3a\x13\xea\x45\xbd\xe5\x56\x33\x51\x19\x4b\xcb\x79\x11\x15\xbc\x92\x43\x94\x84\xe7\x0d\xae\xd7\x8b\x1c\x34\x10\xe5\x91\x60\xac\x6c\x5d\x73\xbe\x87\xfd\x6e\x79\x7f\x4a\xa6\x08\x0d\xdd\xcd\xd6\xbc\x48\xa0\x6f\x32\xc0\x7e\xf7\x7e\x15\x2c\x71\x0e\x4d\xae\xcd\x89\x3d\x2a\x70\x73\x5f\xfd\x99\x66\x57\x64\xac\x7c\xbc\x69\xb6\x04\x85\x86\xc3\x3e\x73\xf9\x7a\x69\xaa\x0a\x6e\x5c\x53\x8e\xba\x85\x36\x8d\x27\x02\x6e\x80\x0b\xd0\x46\x2a\x7c\xa4\x2d\x2a\x96\x7d\xfa\xed\x9c\xd2\x35\xf5\x12\xde\xef\xc0\x7b\x06\x4d\xe7\xc5\xbb\x34\x74\xfc\xb5\xb0\xd5\xe4\xfb\x9d\x87\x5d\x82\x91\x57\x50\x85\xb2\xaa\x86\xc8\xa3\x8b\x64\x21\xee\xec\xf8\x66\xd9\x39\x69\xe9\x56\x4e\xda\xff\x9b\xbc\xff\xd3\xc6\xfb\xbf\x66\x60\xfb\xf9\xf7\x53\x7d\xc9\x7e\x06\x00\x00\xff\xff\x38\x86\x96\xd9\x5a\x07\x00\x00")

func contractsFeecontractCdcBytes() ([]byte, error) {
	return bindataRead(
		_contractsFeecontractCdc,
		"contracts/FeeContract.cdc",
	)
}

func contractsFeecontractCdc() (*asset, error) {
	bytes, err := contractsFeecontractCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts/FeeContract.cdc", size: 1882, mode: os.FileMode(420), modTime: time.Unix(1589524772, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x98, 0x89, 0x8c, 0x4c, 0xde, 0xef, 0x22, 0xe5, 0x98, 0x58, 0xce, 0xf3, 0xc2, 0x64, 0xd5, 0x79, 0xc7, 0xaf, 0x98, 0x2f, 0x69, 0x50, 0x6d, 0x48, 0xe1, 0xdc, 0x5a, 0x70, 0xed, 0x6, 0x9f, 0x97}}
	return a, nil
}

var _contractsFlowtokenCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\x4d\x6f\x24\xb7\xd1\xbe\xcf\xaf\xa8\xd7\x2f\x12\x8f\xd6\xa3\x91\x0c\x04\x39\x08\x5a\x7b\xb5\xc9\x2e\xb0\x07\x07\x81\xb3\x49\xae\x62\x77\xd7\x68\x68\x75\x93\x03\x92\x3d\xa3\xf1\x42\xff\x3d\xa8\xe2\x47\x93\x9c\x8f\x95\x62\x47\x07\x7b\xd5\x24\xeb\x8b\x55\x4f\x3d\x2c\x5d\xbd\x79\x33\x9b\xfd\x3f\x7c\xec\xf5\xee\xb3\x7e\x44\x05\xf8\x24\x86\x4d\x8f\xd0\x6a\xe5\x8c\x68\xdd\x6c\xf6\x79\x2d\x2d\x48\x0b\x62\x5a\x94\xf4\xdf\x01\x95\x13\x4e\x6a\x05\x7a\x05\x6e\x8d\x2c\x04\x3e\x8e\xea\x41\x36\x3d\x82\x17\x67\x9d\x50\x9d\x30\xdd\x72\xf6\xc9\x82\xd2\x0e\x36\xc2\xb8\x78\x20\x2e\x2e\xa0\x19\x1d\xfc\x32\x5a\x07\x76\xad\x77\x16\xd6\x7a\x07\x83\xb6\x0e\x1c\x09\xb1\x33\xbb\xd6\x63\xdf\x4d\x6a\xab\xd3\x52\xb5\xfd\xd8\x49\xf5\x30\x99\xa1\xd0\xed\xb4\x79\xf4\x02\x40\x3a\x8b\xfd\x6a\x49\xbe\x60\xe6\x6b\xf4\x11\xb4\xea\xf7\xa0\x10\x3b\x0b\x4e\x43\x83\xd0\xe1\xa6\xd7\x7b\xec\x40\x2a\xd0\x0a\x41\xb4\xad\x1e\x95\x5b\xb2\x00\xde\x9d\xfb\x91\xe4\xb8\xb5\x70\xb0\x63\x5b\x1b\x32\x50\x1b\x2f\x02\x45\xbb\x86\xd1\xa2\xf9\xd6\x46\x51\x33\x69\xf9\xec\xbf\xc4\xd8\x3b\xd0\xcd\x2f\xd8\xba\x05\x34\xd8\xeb\x9d\x37\xb3\x8a\x71\xa7\xd1\x07\x90\xcc\x24\x2b\x0d\x76\xb8\x92\x0a\x59\x8a\x54\x0e\xcd\x4a\xb4\x68\xbd\x0d\xc2\xe0\x4c\xf4\x06\x45\xb7\x07\xbf\x8d\xed\xe0\xf0\x94\x17\x94\x4e\xce\x66\x6f\xae\x66\x33\x39\x6c\xb4\x71\x69\x93\xdf\xb3\x32\x7a\x80\xeb\xa7\x3f\xd8\xd9\x6c\x33\x36\x93\xbb\x29\x92\x37\xd5\x81\x2f\xb3\x19\x00\xc0\xd5\x15\x7c\xd6\x4e\xf4\x60\xc7\xcd\xa6\xdf\x53\xb8\x56\x74\x39\xfe\x56\x39\x30\x4f\xd2\x3a\x54\x2d\xf2\x7e\x12\xbe\x15\x06\x1c\x1d\xfa\x07\x9f\xb9\x81\x7f\x7e\x94\x4f\x7f\xfe\x53\x12\xf8\x61\xeb\xef\x5f\x38\xca\x49\x1c\xa4\x73\xd8\xc1\x6e\x8d\xaa\xbc\x0a\x69\xa1\x35\x28\x1c\x76\x49\x34\xf2\xd1\xc2\xd2\x4f\x4a\x3a\x29\x7a\xf9\x2b\x76\x73\xe9\xff\x5d\xea\xbd\x78\x91\x62\xef\x8f\x30\x08\x3b\xe9\xd6\x9d\x11\xbb\x10\x34\xe1\xef\xb7\x32\xe1\xdf\x61\xd3\x5c\x0c\x94\x0a\x51\xd7\x82\xcf\xdc\xc0\x5d\xd7\x19\xb4\xf6\xc7\xd7\xea\xee\x70\xa3\xad\x74\x3e\x3d\x8e\x6b\xfe\xab\xdf\x72\xa0\xd8\xe9\x57\xaa\x55\xb8\xcb\x55\x0f\x94\x46\x75\xa4\x7f\x92\xaa\xd6\xf4\x7a\x97\xac\x33\x54\x88\x95\xe8\xf7\xa3\x51\xff\x8d\x68\x01\x03\xee\xbc\xb5\x06\x0c\x5a\x3d\x9a\x16\x4f\xe7\xca\x4f\xbc\xf1\x2f\x7e\x6d\x2e\xfa\x5e\xef\xb0\xbb\x3b\xa1\x76\x8a\xf7\xd5\x55\x32\x25\x16\xbe\x07\x03\x06\x51\xa9\x08\xb9\x5a\xa4\x72\x60\x28\x99\x60\xc0\x97\xa8\xf4\xbb\xc5\x03\xa6\x2a\x5a\x23\xac\x46\xd5\x12\x12\xd8\x58\xc8\xfe\x88\x50\x1d\x3c\xe8\x2d\x1a\x2a\xf1\xc6\x0b\xdb\x18\xe4\xef\x1b\x82\xd0\x56\xab\x4e\xf2\xc1\x28\x4d\xaa\xaa\x5c\x63\xf5\xec\x39\xe6\xad\xe8\x7b\xec\x96\xb9\xee\x76\x8d\xed\xa3\x85\xb5\xd8\x6c\x28\x88\x0e\xcc\xa8\x9c\x1c\x90\x4f\xe2\x16\x0d\x88\x64\x1f\x47\xb3\x10\x11\x25\xfd\x1c\xe2\x4d\xeb\xca\xbb\xde\x60\x8c\x7c\xf4\x8a\x0a\x18\x9f\x4e\x40\x2b\xd9\x18\xa5\xf9\xf4\x88\xd0\xb6\x00\xab\x69\xd9\xf0\x6d\x2a\x0d\x3b\xb1\x87\x95\x26\xc3\x06\xd1\xcb\x56\xea\xd1\xfa\x8b\x70\x3a\xa8\xf4\x01\x4c\x51\xd1\x63\x50\x2a\x15\x08\x69\x96\x70\x07\x76\x83\xad\x14\x7d\x48\x83\x29\x5f\xca\x4e\xe1\x4d\x70\x9a\xd3\x2a\x4a\x9b\xea\xa3\x88\x02\xa5\x56\x12\xc3\xfa\x2b\xe8\x5c\xfe\xdd\xe8\xad\xec\xd0\x2c\xaa\xef\x3f\x63\x8b\x72\x7b\xf8\xfd\xbd\xe8\x39\x99\x02\xe4\x06\xf5\x6b\xdd\x77\xbe\xbf\x34\x61\x5d\xaf\x40\x70\x00\x6c\x6c\xaa\x71\x7b\x44\xdd\xb0\xb3\x44\xdc\x94\x32\x11\x27\x0b\xa1\x94\x09\xd1\x1b\x0e\x2a\xdd\x3f\x25\x46\x3a\x4b\x07\xe7\x95\xe4\x0b\xf8\x92\xd6\xe9\x87\x7b\x73\x14\xf9\x36\x0a\x4f\x5b\x9e\x0b\x4b\x22\xc0\x66\xdf\xf2\xe5\x8f\x31\x0b\x7d\xbe\x88\xc7\x58\x73\x0e\x1f\x28\x4d\xb9\x76\x41\xf0\x47\x61\x1e\x46\x6a\xb0\xf9\x79\x2a\x9b\xa8\x22\xf6\x52\x7f\x86\xf1\x3c\xd5\xdd\x32\x3f\xf4\xc9\x85\x94\xb2\x20\xfc\xcd\x23\xb5\x51\x61\xf6\xa1\x48\x23\x16\x8d\xd6\x67\x0a\x5d\x4f\x2e\x80\xc4\x0e\x5a\xe1\x3e\xed\x6c\x90\xe9\x8c\x11\xca\xae\xd0\x18\xec\x96\xa4\xc5\xa0\x1b\x8d\xf2\x17\xab\x70\xd7\xef\x73\x21\xb1\x90\x82\x4a\x5d\x94\x13\xcb\xf5\x65\x49\x95\x22\x1d\xd7\x60\x93\x35\x8d\x5c\x14\xf6\x16\x77\x54\x4c\xcb\x63\x61\xa6\x84\x59\x8d\x2a\xc5\xa9\x06\xe2\x1b\x78\x97\x98\xc1\xd2\x5b\x73\xf6\xc2\x8b\x5f\x2f\x43\xbc\x8b\x03\x84\xe1\x87\x5d\xd3\xff\x3f\x76\x4d\x96\xa2\x77\x0a\xcd\x8f\x4b\xe1\x5b\xd9\x45\x21\xc4\x07\x0f\x6e\x2f\xf3\xf2\x9f\x72\xd3\x4b\xbb\x38\x91\x76\x21\x4c\xaf\xc9\xba\x82\xda\x55\x29\xc7\x79\x26\xba\xce\x16\x45\xe6\x6c\xaa\xac\x70\x7d\x59\xf5\xd2\xaf\xec\x9e\x3d\x9e\x81\xc4\xd0\x7d\x83\xa2\xc3\xa1\x71\x7a\xa2\x4c\x0a\xbd\x31\x0d\xb6\x62\xb4\x38\xe5\x71\x51\x5a\x64\x63\x96\xbb\x94\xa5\x68\xa2\xee\x80\x66\xdc\x0f\xf8\xe8\xb7\x93\xb5\x6b\x51\x38\xd2\xa0\x27\xd7\x76\x1c\xb0\x63\x57\x19\x99\x57\x9a\xbb\x4b\x48\xbb\xd0\xd8\x97\x07\x69\x15\x42\x3d\xf7\xd7\xfa\x8e\x75\xd5\x90\xd1\xa3\x83\x2d\x7b\x74\x7b\x19\x98\x96\xfd\xbf\x83\xb4\x7b\x71\xd2\x7d\xe7\x85\x2d\x6b\xe8\x49\xb9\x57\xf3\xa6\x62\xbb\xa7\x4f\x5f\xcd\xbf\xe2\x0c\xbc\x85\xeb\xe5\x75\xb1\x1e\xaf\x6c\x5b\x98\x9e\xa5\x61\xd8\x30\xaf\x83\x31\x79\x9d\x11\x67\x78\x7b\xe2\xfb\x65\xe1\x7a\xa6\x27\xd3\x96\xd0\xe4\xc3\xb0\x71\xfb\x63\xc4\xa6\xcc\xf8\x12\xfd\x7c\xaa\x11\x3a\x80\xc8\x33\xf8\x57\x34\x3a\x75\x6f\xd5\x25\x34\x93\x13\x5a\x89\xbe\x27\xdc\x0b\xa8\x45\x3d\x98\x7b\xf6\x40\x4f\x43\x5a\xa4\xf6\x6c\x13\xd5\xc8\x85\x31\xbf\x62\x21\x5e\x6c\x02\xc2\x9a\x53\xf1\x93\xce\x74\x9e\x09\x70\xc5\xf8\xf5\x24\xcc\xbf\xcf\x42\x7b\x17\xf4\x52\xe2\xd7\x16\xb7\xdf\x98\x9c\x36\x50\x85\xd0\x4c\xc1\xed\x37\x78\xd0\xe8\x29\x99\xeb\x30\xce\xcf\xa3\xe3\x57\x30\xea\x7a\x79\x7d\x51\xdd\x13\x71\x93\x3b\xd5\x11\x21\x3e\x41\xb7\x22\x02\x79\x6c\x62\x73\x45\x37\x10\xcb\xf1\x9e\x5a\x68\xf9\xc5\x9a\x9a\x12\x35\x24\xcf\x66\x38\xb2\x0d\x59\xf4\x12\x3e\x93\x99\x52\xd1\x11\xba\x96\xd0\x46\x29\x6e\x9e\xdf\x47\x46\x17\xe9\x78\x09\x60\x89\x4c\x41\x46\x51\x8e\x32\xf0\x42\x13\x1d\xfb\x5c\xf2\x9b\x73\x50\x4d\xdb\x6d\xe6\xdc\x82\x41\x99\xac\x1a\x62\x4e\x72\xed\x80\x2f\x9e\x45\x4d\x15\xb2\x86\x3c\x9c\x48\xe2\x73\x1d\x74\xb2\xf6\xd5\x3d\x94\x38\x7e\xf9\x85\x7e\x42\x90\x7f\x08\x62\xe6\xd7\x17\x37\xf0\x8d\x0f\x57\x78\xa3\xf9\x5a\x6a\x10\x1e\x38\xc3\x0c\xc5\x41\x71\x69\x7e\x73\x4a\xda\x6d\x40\xcb\x2a\xfa\x27\xe4\xf6\x68\xad\x17\xca\xb7\x1e\x6e\xd4\x8b\x2a\x55\x3c\xff\x06\x14\xfb\xee\x18\x4d\x38\xb4\x12\x8e\x99\x7e\x86\x63\x14\x4f\xd6\x9a\x11\xbc\xa0\x40\xcf\x93\x08\xaa\xa4\x57\x24\x67\xc9\x23\x4a\x02\xb1\x00\x3b\x36\xfc\x32\xb2\x39\x6f\xc8\x85\x25\xc2\xea\xb2\x19\xcc\x82\xbe\xa8\xd8\x4a\xb2\x11\xd4\xa2\xaa\x57\x83\xcd\x1e\x0c\x0e\x7a\x1b\xe7\x6a\xa1\x6c\x59\x6c\x9a\xdb\x1c\x25\x88\x0c\x3e\x13\x53\x0d\x69\xe4\x87\x63\xc2\x72\x18\xdc\x59\x66\x39\x05\xea\xf7\x66\x01\x74\x44\xc4\xd4\x38\xdd\xf4\x8f\xf7\xe2\x94\x26\xc5\xf8\xe1\xfc\x9d\xfb\x3c\xf9\x1b\xee\xfc\x23\xf2\x45\x17\x9f\xda\x69\x86\x30\xbe\xb5\x56\x93\x8b\x73\x31\xac\x14\x9f\x98\x5e\xdc\xc0\xbb\x12\xb6\x0f\x7c\x3d\x3b\x03\x29\x7e\x3d\x5f\x29\x99\x9a\x17\x4a\xc9\x02\xc9\xcf\xc8\xe3\x1e\x1c\x7b\x5b\xd4\x18\x50\xfc\x7e\x92\xf2\xf0\xd4\x35\x3d\x74\x0d\x8f\x0f\xf2\x19\xc4\x12\xee\xfa\x1e\x56\x12\xe9\x6d\x2d\xab\x79\x23\xc1\x5f\x14\xd4\xe4\x72\x3a\x7a\x25\xfb\x79\x32\x15\x2e\x11\x68\x3f\x4f\xe7\x79\x77\x36\x54\xd7\x2b\xd8\xad\x85\x9b\x48\x4d\x3d\x02\xf6\x6d\xba\xd3\x51\x77\x66\x6b\x3d\x6a\x21\x57\xd4\x38\x34\xf4\x5a\x10\x06\x41\x98\x46\x3a\x22\xf5\xc5\x46\x8e\x6a\x1e\x40\x7a\x49\x94\x2f\xfd\x1c\x73\xd9\xf3\xd8\xe5\x82\xf2\x03\xe0\xe1\xf8\x97\x08\xfe\xfd\xf5\x35\x71\xdd\x5c\x8b\x4f\xa7\x6c\x8c\xc5\x6c\xb1\x06\xab\x8c\x2c\x30\xc5\x13\x5b\x24\xb2\x28\x55\x31\x1f\xab\xf2\xbf\x00\x85\xe3\x28\x5d\xdb\x78\x51\x5a\x1f\x87\xfd\xa4\x6f\x7e\x7b\xc9\xc2\x3c\xc7\xbf\x0a\x7a\xaf\x56\x11\x62\x3c\x2e\x1d\xf3\x4d\x50\x25\xf6\xb2\x85\x56\x6c\x44\x23\x7b\xe9\xf6\x31\x76\xe1\x0f\x03\xd9\x6c\x80\x87\x61\xf8\xb4\xd1\x16\x6d\x4d\x9c\xee\x03\xe7\xbc\x87\x01\xdd\x5a\xd3\x63\xca\xe8\xf1\xc1\x87\xeb\x3e\xce\x85\xee\xb3\x49\xfe\x91\xa8\x14\x8e\xf5\x52\x3d\xde\xfe\xf1\xcb\xf1\x01\xd3\xf3\x0f\xf3\xa2\xa0\xae\xbc\x1b\x93\xcb\x69\x12\x55\x6c\x73\xc2\x3c\xa0\x3b\x19\xa2\xb4\xf7\x7f\x1c\xab\x70\xc7\xf7\x31\x59\xf3\x50\xbd\x8f\x6b\xbf\x31\x52\x41\xce\x57\x03\x15\xf6\xfd\xae\x71\xa2\x26\x90\x03\x76\x22\xdf\xd3\x1b\xc8\x57\x48\x7c\xc6\xd4\x95\x42\xe5\x31\x64\x02\xa6\x22\x39\x03\xd0\xdf\x5f\xa7\x97\xc7\xa9\x2a\xc9\x64\x9e\xaa\x95\x4c\x41\xe9\xdd\x07\xea\x33\x84\x84\xd3\xe0\xde\xff\x1d\x30\x3d\x11\x12\xcc\x12\x7f\xc8\xb0\x35\x49\xe1\x56\xf5\xd2\x3f\xed\x1c\xaf\xff\xe7\xd9\xf3\xec\x3f\x01\x00\x00\xff\xff\xd6\x4c\xc7\x42\x0f\x1d\x00\x00")

func contractsFlowtokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_contractsFlowtokenCdc,
		"contracts/FlowToken.cdc",
	)
}

func contractsFlowtokenCdc() (*asset, error) {
	bytes, err := contractsFlowtokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts/FlowToken.cdc", size: 7439, mode: os.FileMode(420), modTime: time.Unix(1589524644, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0x21, 0xbc, 0x97, 0xf0, 0x5, 0x72, 0xda, 0xb1, 0xe8, 0xca, 0x8b, 0x28, 0x76, 0x73, 0x6a, 0x54, 0x71, 0xb2, 0x49, 0xe5, 0xb7, 0xd4, 0x52, 0x99, 0x8e, 0xe9, 0xaf, 0x5, 0xe6, 0xe0, 0x70}}
	return a, nil
}

var _contractsFungibletokenCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\xcf\x73\xdb\xba\x11\xbe\xf3\xaf\xd8\x49\x0e\xb1\x53\xc5\x7e\x87\x4e\x0f\x99\xc9\x6b\x93\x69\x3c\x93\x4b\xdb\x69\xdd\xbe\xab\x20\x72\x29\xa1\x06\x01\x3e\x00\x94\xcc\xbc\xc9\xff\xde\xd9\xc5\x0f\x82\x14\x6d\xcb\xd3\xe7\x8b\x25\x12\x58\xec\x2e\xbe\xfd\xf6\x03\x74\xfb\xfe\x7d\x55\xbd\x85\xfb\x03\xc2\x9d\x32\x27\xb8\x1b\xf4\x5e\xee\x14\xc2\xbd\x79\x40\x0d\xce\x0b\xdd\x08\xdb\x54\xd5\xdb\xb7\xb0\x4d\x2f\xf9\xdd\x16\x6a\xa3\xbd\x15\xb5\x07\xa9\x3d\xda\x56\xd4\x58\x55\x64\x28\x7f\x05\x7f\x10\x1e\x84\x52\xd0\x26\xb3\x9e\xcd\xa6\x99\x0e\x4e\x66\x50\x0d\x1c\xc4\x91\x5e\xd1\xf3\xd6\xd8\x0e\xbc\xb9\xa9\xbe\xb5\x20\x60\x70\x68\x1d\x9c\x84\xf6\x8e\xde\x37\xd8\x2b\x33\x82\x00\x8d\xa7\x85\xa9\x0d\xf8\x03\x4a\x9b\xbf\x57\xc1\xb2\x46\x6c\x68\xa6\xec\x7a\x85\x1d\x6a\x4f\xc3\x60\x16\xc8\xe4\xef\x0d\xfb\x5f\x18\x59\xb8\xd7\x1a\x45\x39\xa2\x80\xc8\x8a\x1d\x14\x3a\x10\xba\x01\x2d\x3a\xa9\xf7\x15\x87\xeb\x67\x19\x70\x3d\xd6\xb2\x95\xe8\x6e\x42\x0a\xff\x23\x06\xe5\xb7\x60\xd1\x99\xc1\x52\xc2\xbe\x8a\xfa\x00\xa2\xae\xcd\xc0\xbe\x09\x0f\xe6\xa4\x5d\x08\x2e\xa5\x27\x05\xc1\x7e\x08\x72\x98\xf6\xa5\xc6\xca\xb4\xbc\x1c\x1b\xcd\x36\xc1\x79\x63\xb1\x01\xa9\x63\x4a\x92\x75\x7a\x2e\xf6\x31\xca\xe5\xa4\x83\x70\xd0\xa1\x3f\x98\xc6\x41\x8e\xc3\x9c\x34\x5a\x8e\xd0\xf8\x03\xda\xb8\x1d\xb5\xd0\x50\x0b\xa5\x62\x48\xff\xb0\xe6\x28\x1b\xb4\xdb\x0d\x6c\xff\x89\x35\xca\x23\x7f\xa6\x59\xdb\x2f\x42\x91\xa3\x53\xc0\x53\x6a\x1c\xbb\xe1\xca\x27\xd0\x60\xad\x84\x45\xe8\x2d\x7e\xa8\x8d\x6e\xa4\x97\x46\x87\x14\xf7\xc6\xf9\xf2\x19\xfb\x68\xd1\x79\x2b\x6b\x5f\x91\xb3\xf8\x88\xf5\x40\x2f\x21\xa6\xa5\x1d\x74\x1d\x06\x87\x54\x84\x90\x43\xf8\x23\xd0\x3a\x0e\x7b\x61\x85\x47\xd8\x61\x2d\x06\xf2\xc5\xc3\x5e\x1e\xd1\xf1\x70\x8a\x96\x3f\x88\x9d\x54\xd2\x8f\xb4\x05\xee\x20\x2c\x56\x02\x2c\xb6\x68\x51\xd7\x8c\x8b\x90\xe6\x90\xd0\xb0\x85\x5a\x8d\x80\x8f\xbd\x71\xd1\x54\x2b\x51\x35\x6e\xf2\xa8\x92\x1a\x8c\x46\x30\x16\x3a\x63\x31\x79\x3c\xa5\xe2\xa6\xaa\xbe\x51\xe9\x38\x13\x1d\x0a\xa9\x5f\x78\xd3\x89\x07\x84\x7a\x70\xde\x74\x39\xc3\x31\x35\x19\xf0\x94\x9b\x79\x96\xa9\x90\x0c\x1c\x85\x95\x66\xa0\xd1\x52\xef\x1d\x9c\xa4\x3f\xb0\xf9\x80\xbc\x9b\xea\xce\x58\xc0\x47\x41\x66\x36\x20\xa0\x15\x43\x8d\x9e\xf7\x7e\x87\x93\x75\x6c\x60\x37\xa6\xba\xe5\x1a\xe0\x74\x40\x02\xc5\xac\xb8\xbe\x8c\x30\x38\xa9\xf7\x85\xaf\xb4\xb5\x93\x6b\x9b\x18\xa6\x69\x17\x25\x9a\x09\xa3\x22\x07\x1c\xea\x86\x67\xda\x00\xb7\x54\x2d\x3d\xa2\xfd\xe0\xcd\x07\xfa\xbf\xe1\x88\xcc\xe0\xa9\x6a\x68\x4d\x22\x01\x5a\x88\xb9\x81\x82\x15\x50\x23\x59\x55\xa0\xb0\xd9\xa3\x05\xd7\x09\xeb\xf3\x52\x37\x70\x6f\xc2\x4a\xd1\xba\x37\x20\xf4\x54\x07\x9b\x2a\xd0\x53\xac\x51\x47\x29\x19\x79\xd1\xc6\x8a\x53\x91\x4a\x68\xad\xe9\x4a\x8c\x30\x55\x85\x12\x62\xe0\x36\xd8\x1b\x27\x7d\x46\x07\x18\x3d\x5b\xe9\x9d\x4b\xd8\x22\x86\xa4\xcc\x7b\x0c\xf6\xad\xd0\xae\x45\x7b\x53\x55\xef\x6f\xab\xea\xf6\x96\x79\xbc\x13\x52\x2f\x79\xbc\xd8\x85\xdb\x5b\xf8\x3b\x9b\x7e\x9a\x93\xa5\x52\x33\xc2\x94\xae\xa0\xf8\xdb\xdb\xaa\x1f\x76\x2b\xe4\xbf\xd8\xb2\xdf\xaa\x0a\x00\x20\x3a\xe5\x8d\x17\x0a\xf4\xd0\xed\xd0\x32\xda\x43\x6a\xa4\x06\x7c\x94\xce\x53\x25\xdd\xa4\xf1\xdf\x3c\x48\x07\x43\x1f\x4b\xab\x00\x9b\xa5\x47\xa8\xdd\x60\x63\x6f\x09\x66\xdd\xd0\xf7\x6a\x4c\xd3\x9d\x17\xa3\x23\xd2\x1b\xb8\xb4\x09\x27\xc1\x56\x23\x3c\xf2\x20\xf2\xff\x28\x6c\x98\xfd\x2f\x9e\xfc\x11\xfe\x7d\x27\x1f\xff\xf4\xc7\xec\xf4\xd7\x23\x26\x42\x96\x0e\xb0\x93\x9e\xb0\x7e\xa2\x8d\x23\x9f\xa6\xf0\x1d\xd4\x16\x85\xc7\x26\x9b\x46\x9e\x3a\xcb\xc6\x37\x2d\xbd\x14\x4a\x7e\xc7\xe6\x4a\x86\xcf\xf3\x75\xaf\x2f\x5a\x38\x24\x8d\x98\x2b\xe1\x4c\x07\x74\x89\x80\x90\x85\x0b\xbf\xc4\x41\x57\xa2\xa3\x06\x90\xd6\xda\xf0\x9c\x8f\xf0\xb9\x69\x2c\x3a\xf7\xe7\xd7\xae\x1d\xf1\x1a\x7a\xd2\xfa\xca\x7f\x0d\x43\xce\x16\xf6\x66\x6d\xd9\x44\x16\xf1\x7b\xc6\xc1\x5c\x47\x20\x71\x4c\x1d\x09\xd5\xe2\xaf\x83\xb4\x8c\x0a\x07\xad\xb1\x39\x21\xc4\x41\x71\xfe\xa2\xfc\x26\x1c\x31\x1d\x8c\x7d\x46\x5c\x9a\xf0\x0b\x42\x63\xf4\xbb\xbc\xd4\x7c\x15\xa3\x61\xbb\x4b\xfd\xec\x80\x16\x37\x69\x5e\xd1\x3d\x14\x0a\x62\x6b\xd3\x47\x9c\xf4\xc6\x39\x19\x09\xdb\xb4\x01\x2a\xb4\x7c\x24\xed\x3e\x46\xee\xb2\xcf\x14\x69\x70\x42\x63\x8d\xce\x09\x2b\xd5\x18\x15\x00\x53\x88\x39\x69\x88\x6e\xcc\xfc\xa7\xe4\x9f\x77\xd9\x89\x88\x63\x3d\xc6\x75\x52\xba\x8a\x67\xe5\xeb\xbb\x44\x45\xec\x8f\x1b\x76\x91\x1a\x96\x29\x65\x81\x90\xf8\xa9\x34\x10\xd8\xd9\x0f\x96\x30\xb3\xd4\x1a\xb9\xd7\x58\xec\xcc\x11\x9b\xdc\x73\xd6\x9d\xb9\x2f\x7a\xf9\x3b\xae\x6c\x74\x0e\x14\x1e\x51\x11\x4c\xfb\x61\xa7\x64\xbd\x81\xdd\x90\xb8\xca\x51\xfa\x04\x25\x77\xa7\xb0\x2b\x4d\xa5\x9d\xe2\x06\x3d\x29\x1c\x6e\x2a\xde\x58\x06\x06\xfb\x95\xf3\x38\xd7\x50\xa5\xad\x9a\x95\x18\xd7\xb2\x1a\x99\xcd\xc3\xf2\xc9\xd5\xe7\xc2\x09\xcb\x76\x62\x84\xbd\x15\xda\x47\x79\x15\x17\xc9\x21\x52\x67\x4d\x80\xa1\x70\xe4\x31\x31\x58\x76\xa1\xcf\x6a\x80\x36\x2a\xb4\x22\xc1\x2a\x35\x2a\xcf\x7a\x26\xdd\xa8\x60\xd9\x76\x69\x85\x71\x9a\x60\x92\x43\xf7\x07\x6b\x86\x3d\xb5\xc9\xac\x75\x2e\x8b\x28\x88\x16\x0e\x8b\x72\xf2\x42\x50\xbc\x79\x97\xc6\x44\xf6\x56\xc3\x99\xc5\x50\x5a\x7b\x75\x38\x54\x46\xed\xa0\x73\x81\x2c\x28\xec\xfa\x23\xfc\x25\xa0\xf9\xb7\x3c\x85\xa7\x19\xb7\x7c\x14\x3d\xd8\x5a\x74\x51\xf3\xb7\xd1\xe7\x00\x31\x2a\x0e\x38\x0a\x35\xe0\xd9\xb4\x30\xe5\x26\x96\x39\x7c\xfa\x04\xd1\x8b\xb3\x91\xf4\xf7\x26\xf1\xbc\x50\x71\x1c\x74\x83\xf3\xa4\xd3\x68\x25\x27\x3a\x04\x11\x52\x94\x2c\x46\xbd\x39\x75\x11\x8e\xe9\xcd\xcc\xfc\x8f\x6a\xfe\xe9\x47\xe6\xeb\xa4\xf2\xff\x1f\xbe\x8e\x4d\xe4\x9c\xae\xa5\x5e\xb6\xfd\x17\xe9\x5a\xea\x5a\x0d\x0d\x92\xa2\x4b\x07\x84\xe0\x42\x7d\xc0\xfa\x61\x1e\x79\x64\x80\x64\xe3\x84\x7c\xba\xa4\x5d\x21\x9d\x7d\x89\xcc\x0e\xb1\x07\x99\x5d\x15\x5c\xd0\x98\x34\x66\x5d\x52\x6f\x40\xc9\x07\x3a\x11\x2a\xc9\xa7\xab\x8e\x64\x89\xd0\x4d\xd6\x2d\xac\x35\xe9\x39\x69\x15\xd9\x32\x4a\x3d\xf4\x2a\x9c\x07\xe0\x45\xaa\x4f\xdb\xb2\xa0\xfa\x98\xe9\x8b\x98\x3e\xca\x7b\x22\xb3\xd0\xde\x93\x36\x0d\x21\x94\x13\xd7\xf7\x69\xaa\xb7\xb1\xc7\xe7\xea\x2b\x1a\xbe\x0a\x52\x24\xd4\xd4\xf5\xb2\xa8\x2c\xae\xd4\x14\xcd\xc8\xa5\xf1\x73\xac\xcb\xab\x9f\xae\x9f\x28\x8e\x28\x45\x32\x00\x52\x69\x04\xf8\x1d\xf1\x52\xd0\xc7\xa3\xec\xf3\x98\x27\x59\x28\xa4\x0e\x08\x9a\xd4\x02\x1f\xff\xa0\x3c\xad\xa7\xf9\xd4\x27\x8b\x42\x21\xbe\x23\xb5\xa5\xf1\x14\xc6\xbd\x0b\x92\x2b\x4a\xcc\x4d\x09\xe5\x64\x82\xd5\x79\x96\x97\x50\x1b\x6b\xb1\xf6\x6a\xbc\x04\x32\x31\xa8\x05\x62\x26\xc1\xbe\xe0\x8b\xc8\xe5\xef\xdc\x12\x0f\x49\x51\xc7\xf1\x73\x35\x4d\x7f\xe4\xe1\xd5\xe2\xed\xd9\x76\xaf\x73\xa8\x43\xd5\x96\x54\x98\xac\xac\x6f\xf7\x97\xc5\x36\x97\xa9\x49\x80\x0d\x8f\x92\xa1\x4b\x01\x50\xee\x5b\x79\xb4\x29\x5a\xcc\x12\x01\xd3\x0d\x84\x37\x4f\x9d\x5a\x97\xac\x76\xcf\x27\xc2\x5a\x09\x2b\xd2\x65\x06\xf3\x5a\x6d\xf9\xd8\x37\xf6\xac\x4a\xc4\xda\x01\xac\x43\xa1\xe7\xbc\x84\x47\xb4\xe3\xf2\x30\x98\x67\xce\x2f\x0a\xdc\xf2\xa4\x17\x6d\x70\x26\x1b\x6c\xa5\xc6\xd2\x93\xd0\x05\xcd\xee\xbf\x18\x2d\x65\x2e\x0c\x97\x01\xb9\xd3\x5d\x76\x41\x54\xdc\x0b\x15\xa5\x11\x99\x9d\xb3\xe9\xf2\x35\x0a\xbd\x99\xae\x52\x9e\x42\x39\x3b\xf8\x31\x0b\xe0\x4d\xe6\xc7\x4d\x01\xfb\x57\xa0\xfe\xf5\xa0\x8f\x46\xa7\xdb\x91\xb0\x7d\x31\xa1\xe1\xaa\x6b\x12\x91\xf2\xfb\x5c\xba\xa4\xcb\x53\x73\x72\xa4\xfd\x28\x8a\xa4\x4f\x17\x08\x4e\x8e\x1e\x17\x2a\xfc\xf9\xea\x5b\x3d\x0f\x14\x5a\x7f\x1b\xb4\xc4\x76\x52\xfb\x6c\xdf\xe5\x55\x53\xef\x82\x24\xb5\x92\xdc\x3f\x96\xd0\xc8\x5d\x30\x59\xc6\x66\xc6\x60\xf0\xfb\xe8\xaf\xd5\x56\x31\xa3\x8e\x9f\x5f\x50\x51\x9f\x83\x74\x9a\x34\x51\xa2\x10\x15\x04\xa6\xd0\x60\x2c\xe0\xaf\x83\x50\xe1\xdb\x8a\xa0\x7a\x56\x46\xc1\xb3\x3a\x91\x8e\x25\x9c\x26\xd2\xed\x42\x4d\xb7\x41\xdb\x1d\xb6\xc6\xe2\x96\x35\x0a\xfa\xb8\x13\x6a\xc8\x8b\x2e\xfa\xcc\x9a\xf1\x78\x9f\xbb\xc3\xbd\xd4\x9a\x50\xb4\xb8\x22\x9d\x2e\x4f\x57\x66\xbf\xcc\xc8\xec\xe0\x55\xf9\xf8\x1a\x3e\x3c\x9f\xed\xbf\xa5\x0e\x77\xd6\x98\xf9\x4a\x2c\xca\x9f\x29\xb3\xbd\xc5\x23\xdf\x57\x16\xe8\x7b\x9d\x86\x5d\xd1\x44\xe0\xc5\x03\x9e\x21\x56\xd0\x93\x5e\x58\xd1\xa1\x8f\xd7\xdf\xa2\x69\xe6\xe2\xa7\x28\x83\x48\x73\x0b\x24\xc4\x4b\xf8\x27\x4b\xf2\x55\x3a\xe8\xc2\xc6\xb8\xb6\x0d\x7f\x48\x8f\x4b\xd9\xf4\x94\x56\x7a\x7e\x53\xdc\xd0\xbd\xb8\x1b\xd3\xdd\xd0\xab\x4e\x14\x41\xdf\x7c\xed\x7a\x3f\xae\xb5\xd9\xcf\x7a\x0c\x77\xad\xe9\xa7\x87\xf9\x49\x9b\x2f\x46\xd9\x42\xfc\x69\xa8\xec\x4d\xb3\x9b\x95\x83\x98\x5c\xfe\xf4\x09\x7e\x5a\x36\x0f\xda\x91\xa5\x2f\x57\x6b\x9c\xb3\xb2\x25\xe7\x87\xb6\x49\x9a\xc2\x1b\x6a\x04\x1a\x4f\x6a\x4c\x5a\x2e\x3a\xc9\x09\xe6\x1f\x76\xbe\xa3\x35\xe7\x9a\x24\x65\xea\x47\xf5\xbf\x00\x00\x00\xff\xff\xce\x85\x62\x0d\xab\x1b\x00\x00")

func contractsFungibletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_contractsFungibletokenCdc,
		"contracts/FungibleToken.cdc",
	)
}

func contractsFungibletokenCdc() (*asset, error) {
	bytes, err := contractsFungibletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts/FungibleToken.cdc", size: 7083, mode: os.FileMode(420), modTime: time.Unix(1589523885, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x53, 0x24, 0x3b, 0x7e, 0x5b, 0x96, 0x55, 0x20, 0x21, 0x25, 0x10, 0x74, 0xe2, 0x3e, 0x38, 0xca, 0x8b, 0xf, 0x41, 0xc6, 0x9c, 0xcc, 0x2f, 0x92, 0x12, 0xf0, 0x2d, 0x4, 0xdc, 0xc0, 0xf8, 0xd8}}
	return a, nil
}

var _contractsServiceaccountCdc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x4d\x6b\xdb\x40\x10\xbd\xeb\x57\x4c\x02\x0d\x12\x24\xb2\x0b\xa5\x07\x91\x0f\x52\xb7\x3e\x97\xd4\xe9\xb5\xac\x56\x23\x6b\xc9\x7a\x57\xac\x46\xfe\xc0\xf8\xbf\x97\x95\x56\xb2\xd6\x1f\x35\x2d\x2d\x34\x27\x25\xf3\xde\xcc\xbc\xb7\x33\x13\xb1\x28\xb5\x21\x98\xd6\x6a\x2e\x52\x89\x33\xfd\x86\x0a\x72\xa3\x17\x30\x5e\xbf\xab\x82\x2e\x2c\xf5\xea\x5c\x08\x71\xa2\x15\x19\xc6\x69\x10\x0c\xca\x3a\x05\xde\xfd\xfd\x1b\x9a\xa5\xe0\xf8\xcc\xb9\xae\x15\xc1\x36\x08\x00\x00\x2c\x44\x22\x01\x19\xa6\x2a\xc6\x49\x68\x35\x45\x4c\xe0\x75\x2a\xd6\x1f\x3f\x78\x10\xd6\x32\x27\x06\xd9\x21\xac\xc7\xe5\xb5\x82\x0c\xb3\x9a\xd3\xcc\x4b\x18\xfe\xb0\x74\x4a\xe0\xb9\xa6\xc2\xb5\x10\xc1\xb6\xe1\xd9\x9f\xa6\x05\xab\xed\x3b\xab\x25\xc1\x03\x54\x28\xf3\x38\xc3\xdc\xfe\x3a\xeb\x03\xa1\x4d\x12\x79\xac\x1c\xb1\xe5\xdc\xdf\x0d\x32\xc4\x2b\x41\x45\x66\xd8\x2a\x64\x0b\x5b\x2b\x69\x13\xfa\x2a\xa3\xa0\xcf\x34\x30\x30\xce\xb0\xd4\x95\xa0\xd0\x1a\x99\xd8\xb4\x5d\x89\xb6\xf0\xee\x94\xda\xe7\x23\x6f\xfe\x13\xc5\xc7\x8f\xf6\x57\x54\x0b\x25\xe8\xf3\xa0\xd7\x4b\x6a\x47\x23\x68\x5a\x40\x60\xa0\x70\x35\x98\xe5\x56\x0a\x53\x19\x94\x35\x81\x20\x10\x0a\x2a\xd2\x86\xcd\xb1\x67\xdb\xd4\x71\xc5\x96\x18\xde\xdf\xf5\xcc\x98\x37\x09\xbf\x2c\x4a\xda\xb4\x56\x45\xb7\x40\x3a\x81\x91\xa3\x8f\xf2\x0e\xea\x74\x9c\xec\xa6\xac\x53\x29\x38\x70\x56\xb2\x54\x48\x41\x1b\x20\x0d\x54\xa0\x6b\x8c\x0a\x46\xa0\x95\xdc\x00\xae\x4b\x5d\x61\x35\x4c\x62\x61\xce\x37\xeb\x4a\x33\x57\x40\x85\xd1\xf5\xbc\x68\x82\x2f\xc8\x51\x2c\xd1\x80\x50\x84\x26\x67\xfc\x40\x93\x14\xea\xed\xfe\x66\xeb\x2d\x7e\xdc\x91\x76\x8f\x61\x8f\x6e\x0a\xb6\xad\xee\x65\x75\xc0\x5b\x0f\x46\xcc\xcc\x91\xce\xda\xd0\x63\xff\x91\x1f\x29\x93\x4c\x71\x84\x5c\xa0\xcc\x3c\x33\x3e\xb9\xc8\xef\x78\xe1\x38\x17\xad\x70\xb8\x3f\x75\xe2\xe4\x5e\xef\xa7\xdb\x65\xef\x87\xfc\x6b\x53\xbd\x1b\xf3\xee\x0a\x1e\x2c\xb7\x33\xe2\x05\x73\x78\x68\x78\x5e\x6f\xf1\x1c\x69\xd2\x7b\x1c\x9e\x13\x14\x5d\xf9\xac\x54\x1b\xa3\x57\xbf\xb0\x29\xba\xda\x3f\xab\x41\xaa\x8d\x1a\x34\x12\xbb\xcf\x4b\x7a\xdb\x75\x3a\xb5\xd2\x09\xdc\xec\x17\xb0\x9d\x88\xed\x61\xbd\xe6\x35\xbb\x46\x0f\xd0\x8f\xee\xbe\x9c\xdb\x51\x78\x7a\x82\x92\x29\xc1\xc3\xeb\x57\xc5\x52\x89\x76\xfa\xda\x5c\x60\x30\x47\x83\x76\x82\xdc\x44\xba\x96\xdb\x23\x08\x4b\xfb\x7d\xed\xbd\xa5\xbd\x52\xe1\xf0\x0c\x9d\xf8\x27\x00\x0f\x30\x8e\xc7\xe3\xf7\x3e\xe6\xf8\x6c\x7a\xb8\x5d\xb0\x0b\x7e\x06\x00\x00\xff\xff\x56\x24\xdf\xe5\xb5\x07\x00\x00")

func contractsServiceaccountCdcBytes() ([]byte, error) {
	return bindataRead(
		_contractsServiceaccountCdc,
		"contracts/ServiceAccount.cdc",
	)
}

func contractsServiceaccountCdc() (*asset, error) {
	bytes, err := contractsServiceaccountCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts/ServiceAccount.cdc", size: 1973, mode: os.FileMode(420), modTime: time.Unix(1589570017, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0x94, 0xa4, 0x32, 0xb, 0xce, 0xcb, 0x4f, 0x94, 0xec, 0x20, 0xa6, 0x43, 0x95, 0xe3, 0xaf, 0x41, 0xf1, 0x4f, 0x13, 0x4c, 0x3e, 0x5d, 0x35, 0xb1, 0xf1, 0xd4, 0x93, 0x19, 0x28, 0x21, 0x49}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"contracts/FeeContract.cdc": contractsFeecontractCdc,

	"contracts/FlowToken.cdc": contractsFlowtokenCdc,

	"contracts/FungibleToken.cdc": contractsFungibletokenCdc,

	"contracts/ServiceAccount.cdc": contractsServiceaccountCdc,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"contracts": &bintree{nil, map[string]*bintree{
		"FeeContract.cdc":    &bintree{contractsFeecontractCdc, map[string]*bintree{}},
		"FlowToken.cdc":      &bintree{contractsFlowtokenCdc, map[string]*bintree{}},
		"FungibleToken.cdc":  &bintree{contractsFungibletokenCdc, map[string]*bintree{}},
		"ServiceAccount.cdc": &bintree{contractsServiceaccountCdc, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
