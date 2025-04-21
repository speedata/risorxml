package risorxml

import (
	"context"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
	"github.com/speedata/goxml"
)

type xmlNodes struct {
	pos     int
	n       []goxml.XMLNode
	current object.Object
}

// Type of the object.
func (xn *xmlNodes) Type() object.Type {
	return "xml.nodes"
}

// Inspect returns a string representation of the given object.
func (xn *xmlNodes) Inspect() string {
	return "xml.nodes"
}

// Interface converts the given object to a native Go value.
func (xn *xmlNodes) Interface() interface{} {
	return xn.n
}

// Returns True if the given object is equal to this object.
func (xn *xmlNodes) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (xn *xmlNodes) GetAttr(name string) (object.Object, bool) {
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (xn *xmlNodes) SetAttr(name string, value object.Object) error {
	return object.Errorf("xml.nodes: cannot set attribute %s", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (xn *xmlNodes) IsTruthy() bool {
	return len(xn.n) > 0
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (xn *xmlNodes) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("xml.nodes: operation %s not supported", opType)
}

// Cost returns the incremental processing cost of this object.
func (xn *xmlNodes) Cost() int {
	return 0
}

// Next advances the iterator and then returns the current object and a
// bool indicating whether the returned item is valid. Once Next() has been
// called, the Entry() method can be used to get an IteratorEntry.
func (xn *xmlNodes) Next(_ context.Context) (object.Object, bool) {
	if xn.pos >= len(xn.n) {
		return nil, false
	}

	switch t := xn.n[xn.pos].(type) {
	case *goxml.Element:
		xn.current = &XMLElement{Value: t}
	case goxml.CharData:
		xn.current = &xmlText{text: &t}
	case goxml.ProcInst:
		xn.current = &XMLProcInst{Value: &t}
	case goxml.Comment:
		xn.current = &XMLComment{Value: t}
	default:
		return nil, false
	}
	xn.pos++
	return xn.current, true
}

// Entry returns the current entry in the iterator and a bool indicating
// whether the returned item is valid.
func (xn *xmlNodes) Entry() (object.IteratorEntry, bool) {
	if xn.current == nil {
		return nil, false
	}
	return object.NewEntry(object.NewInt(int64(xn.pos)), xn.current), true
}
