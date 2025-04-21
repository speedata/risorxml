package risorxml

import (
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
	"github.com/speedata/goxml"
)

// XMLProcInst represents an XML processing instruction in the Risor XML module.
type XMLProcInst struct {
	Value *goxml.ProcInst
}

// Type of the object.
func (pi *XMLProcInst) Type() object.Type {
	return "xml.procinst"
}

// Inspect returns a string representation of the given object.
func (pi *XMLProcInst) Inspect() string {
	return pi.Value.Target + " " + string(pi.Value.Inst)
}

// Interface converts the given object to a native Go value.
func (pi *XMLProcInst) Interface() interface{} {
	return pi.Value
}

// Equals returns True if the given object is equal to this object.
func (pi *XMLProcInst) Equals(other object.Object) object.Object {
	return object.NewBool(pi.Value == other.(*XMLProcInst).Value)
}

// GetAttr returns the attribute with the given name from this object.
func (pi *XMLProcInst) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "is_element":
		return object.False, true
	case "is_text":
		return object.False, true
	case "is_comment":
		return object.False, true
	case "is_procinst":
		return object.True, true
	case "text":
		return object.NewString(string(pi.Value.Inst)), true
	case "target":
		return object.NewString(pi.Value.Target), true
	}

	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (pi *XMLProcInst) SetAttr(name string, value object.Object) error {
	return object.Errorf("xml.procinst: cannot set attribute %s", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (pi *XMLProcInst) IsTruthy() bool {
	return string(pi.Value.Inst) != ""
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (pi *XMLProcInst) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("xml.procinst: operation %s not supported", opType)
}

// Cost returns the incremental processing cost of this object.
func (pi *XMLProcInst) Cost() int {
	return 0
}
