package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Call statement
type Call struct {
	block    value.Value
	name     string
	function value.Value
	operands []value.Value
}

func NewCall(block value.Value, name string, function value.Value, operands ...value.Value) *Call {
	fType := assertFunction(function.Type())
	types := make([]types.Type, len(operands))
	for i, o := range operands {
		types[i] = o.Type()
	}

	// Check we have the correct ammount of operands
	fArgs := fType.Arguments()
	if len(types) != len(fArgs) {
		panic(fmt.Sprintf("Function takes %d operands, you have %d",
			len(fArgs),
			len(types)))
	}

	// Check they are of the correct type
	for i := 0; i < len(types); i++ {
		if !fArgs[i].Equal(types[i]) {
			panic("Operand(s) dont equal function type")
		}
	}

	return &Call{block, name, function, operands}
}

func (i *Call) Block() value.Value {
	return i.block
}

func (i *Call) IsTerminator() bool {
	return false
}

func (i *Call) Type() types.Type {
	return i.function.Type().(types.Function).ReturnType()
}

func (i *Call) Ident() string {
	return "%" + i.name
}

func (i *Call) Llvm() string {
	arguments := ""
	for opIndex, op := range i.operands {
		arguments += fmt.Sprintf("%s %s",
			op.Type().String(),
			op.Ident())

		if opIndex < len(i.operands)-1 {
			arguments += ", "
		}
	}

	return fmt.Sprintf("%%%s = call %s %s(%s)",
		i.name,
		i.function.Type().(types.Function).ReturnType().String(),
		i.function.Ident(),
		arguments)
}
