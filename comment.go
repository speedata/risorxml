package risorxml

import (
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
	"github.com/speedata/goxml"
)

type XMLComment struct {
	Value goxml.Comment
}

// Type of the object.
func (c *XMLComment) Type() object.Type {
	return "xml.comment"
}

// Inspect returns a string representation of the given object.
func (c *XMLComment) Inspect() string {
	return c.Value.Contents
}

// Interface converts the given object to a native Go value.
func (c *XMLComment) Interface() interface{} {
	return c.Value
}

// Returns True if the given object is equal to this object.
func (c *XMLComment) Equals(other object.Object) object.Object {
	return object.NewBool(c.Value.Contents == other.(*XMLComment).Value.Contents)
}

// GetAttr returns the attribute with the given name from this object.
func (c *XMLComment) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "is_element":
		return object.False, true
	case "is_text":
		return object.False, true
	case "is_comment":
		return object.True, true
	case "is_procinst":
		return object.False, true
	case "text":
		return object.NewString(c.Value.Contents), true
	}

	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (c *XMLComment) SetAttr(name string, value object.Object) error {
	return object.Errorf("xml.comment: cannot set attribute %s", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (c *XMLComment) IsTruthy() bool {
	return c.Value.Contents != ""
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (c *XMLComment) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("xml.comment: operation %s not supported", opType)
}

// Cost returns the incremental processing cost of this object.
func (c *XMLComment) Cost() int {
	return 0
}
