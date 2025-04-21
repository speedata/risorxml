package risorxml

import (
	"github.com/risor-io/risor/errz"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
	"github.com/speedata/goxml"
)

// XMLAttribute represents an XML attribute in the Risor XML module.
type XMLAttribute struct {
	Value *goxml.Attribute
}

// Type of the object.
func (xmlattr *XMLAttribute) Type() object.Type {
	return "xml.attribute"
}

// Inspect returns a string representation of the given object.
func (xmlattr *XMLAttribute) Inspect() string {
	return xmlattr.Value.String()
}

// Interface converts the given object to a native Go value.
func (xmlattr *XMLAttribute) Interface() interface{} {
	return xmlattr.Value
}

// Equals returns True if the given object is equal to this object.
func (xmlattr *XMLAttribute) Equals(other object.Object) object.Object {
	return object.NewBool(xmlattr.Value == other.(*XMLAttribute).Value)
}

// GetAttr returns the attribute with the given name from this object.
func (xmlattr *XMLAttribute) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "id":
		return object.NewInt(int64(xmlattr.Value.ID)), true
	case "name":
		return object.NewString(xmlattr.Value.Name), true
	case "value":
		return object.NewString(xmlattr.Value.Value), true
	case "namespace":
		return object.NewString(xmlattr.Value.Namespace), true
	case "prefix":
		return object.NewString(xmlattr.Value.Prefix), true
	default:
		return nil, false
	}
}

// SetAttr sets the attribute with the given name on this object.
func (xmlattr *XMLAttribute) SetAttr(name string, value object.Object) error {
	switch name {
	case "name":
		xmlattr.Value.Name = value.Interface().(string)
	case "space":
		xmlattr.Value.Namespace = value.Interface().(string)
	case "value":
		xmlattr.Value.Value = value.Interface().(string)
	case "prefix":
		xmlattr.Value.Prefix = value.Interface().(string)
	default:
		return errz.ArgsErrorf("unknown attribute: %s", name)
	}
	return nil
}

// IsTruthy returns true if the object is considered "truthy".
func (xmlattr *XMLAttribute) IsTruthy() bool {
	return xmlattr.Value.Value != ""
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (xmlattr *XMLAttribute) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("xml.attribute: operation %s not supported", opType)
}

// Cost returns the incremental processing cost of this object.
func (xmlattr *XMLAttribute) Cost() int {
	return 0
}
