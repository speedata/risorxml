package risorxml

import (
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
	"github.com/speedata/goxml"
)

// xmlText represents a text node in an XML document.
type xmlText struct {
	text *goxml.CharData
}

// Type of the object.
func (cd *xmlText) Type() object.Type {
	return "xml.text"
}

// Inspect returns a string representation of the given object.
func (cd *xmlText) Inspect() string {
	return cd.text.Contents
}

// Interface converts the given object to a native Go value.
func (cd *xmlText) Interface() interface{} {
	return cd.text.Contents
}

// Returns True if the given object is equal to this object.
func (cd *xmlText) Equals(other object.Object) object.Object {
	return object.NewBool(cd.text.Contents == other.(*xmlText).text.Contents)
}

// GetAttr returns the attribute with the given name from this object.
func (cd *xmlText) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "is_element":
		return object.False, true
	case "is_text":
		return object.True, true
	case "is_comment":
		return object.False, true
	case "is_procinst":
		return object.False, true
	case "text":
		return object.NewString(cd.text.Contents), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (cd *xmlText) SetAttr(name string, value object.Object) error {
	return object.Errorf("xml.text: cannot set attribute %s", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (cd *xmlText) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (cd *xmlText) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("xml.chardata: operation %s not supported", opType)
}

// Cost returns the incremental processing cost of this object.
func (cd *xmlText) Cost() int {
	return 0
}
