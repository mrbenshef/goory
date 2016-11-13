package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"
	"github.com/bongo227/goory/value"
)

type Extractvalue struct {
	name     string
	location value.Value
	position int
}

func NewExtractvalue(name string, location value.Value, position int) *Extractvalue {
	return &Extractvalue{name, location, position}
}

func (i *Extractvalue) String() string {
	return "extractvalue"
}

func (i *Extractvalue) IsTerminator() bool {
	return false
}

func (i *Extractvalue) Type() types.Type {
	return i.location.Type().(types.Aggregate).Position(i.position)
}

func (i *Extractvalue) Ident() string {
	return "%" + i.name
}

func (i *Extractvalue) Llvm() string {
	return fmt.Sprintf("%%%s = extractvalue %s %s, %d",
		i.name,
		i.location.Type().String(),
		i.location.Ident(),
		i.position)
}