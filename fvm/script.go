package fvm

import (
	"github.com/onflow/cadence"
	"github.com/onflow/cadence/runtime"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/model/hash"
)

func Script(script []byte) *InvokableScript {
	return &InvokableScript{
		Script: script,
	}
}

type InvokableScript struct {
	Script    []byte
	Arguments [][]byte
	ID        flow.Identifier
	Value     cadence.Value
	Logs      []string
	Events    []cadence.Event
	Err       Error
}

type ScriptProcessor interface {
	Process(*VirtualMachine, Context, *InvokableScript, Ledger) error
}

func (inv *InvokableScript) WithArguments(args [][]byte) *InvokableScript {
	return &InvokableScript{
		Script:    inv.Script,
		Arguments: args,
	}
}

func (inv *InvokableScript) Invoke(vm *VirtualMachine, ctx Context, ledger Ledger) error {
	scriptHash := hash.DefaultHasher.ComputeHash(inv.Script)
	inv.ID = flow.HashToID(scriptHash)

	for _, p := range ctx.ScriptProcessors {
		err := p.Process(vm, ctx, inv, ledger)
		vmErr, fatalErr := handleError(err)
		if fatalErr != nil {
			return fatalErr
		}

		if vmErr != nil {
			inv.Err = vmErr
			return nil
		}
	}

	return nil
}

type ScriptInvocator struct{}

func NewScriptInvocator() ScriptInvocator {
	return ScriptInvocator{}
}

func (i ScriptInvocator) Process(
	vm *VirtualMachine,
	ctx Context,
	inv *InvokableScript,
	ledger Ledger,
) error {
	env := newEnvironment(ctx, ledger)

	location := runtime.ScriptLocation(inv.ID[:])

	value, err := vm.Runtime.ExecuteScript(inv.Script, inv.Arguments, env, location)
	if err != nil {
		return err
	}

	inv.Value = value
	inv.Logs = env.getLogs()
	inv.Events = env.getEvents()

	return nil
}
