package object

import "fmt"

const INTEGER_OBJ = "INTEGER"
const BOOLEAN_OBJ = "BOOLEAN"
const NULL_OBJ = "NULL"
const RETURN_OBJ = "RETURN"

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type Null struct {}
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }

type Return struct {
	Value Object
}
func (r *Return) Type() ObjectType { return RETURN_OBJ }
func (r *Return) Inspect() string { return r.Value.Inspect() }
