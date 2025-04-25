package risorxml

import (
	"context"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
	"github.com/speedata/goxml"
)

type XMLDocument struct {
	Value *goxml.XMLDocument
}

// Type of the object.
func (xp *XMLDocument) Type() object.Type {
	return "xml.document"
}

// Inspect returns a string representation of the given object.
func (xp *XMLDocument) Inspect() string {
	return xp.Value.String()
}

// Interface converts the given object to a native Go value.
func (xp *XMLDocument) Interface() interface{} {
	return xp.Value
}

// Returns True if the given object is equal to this object.
func (xp *XMLDocument) Equals(other object.Object) object.Object {
	return object.NewBool(xp.Value == other.(*XMLDocument).Value)
}

func (xp *XMLDocument) append(ctx context.Context, args ...object.Object) object.Object {
	for _, arg := range args {
		typ := arg.Type()
		switch typ {
		case "xml.element":
			elt := arg.(*XMLElement)
			xp.Value.Append(elt.Value)
		case "xml.attribute":
			attr := arg.(*XMLAttribute)
			xp.Value.Append(attr.Value)
		default:
			// fmt.Printf("~~> t %#v\n", typ)
		}
	}
	return nil

}

// GetAttr returns the attribute with the given name from this object.
func (xp *XMLDocument) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "root":
		r, err := xp.Value.Root()
		if err != nil {
			return object.NewError(err), true
		}
		elt := &XMLElement{
			Value: r,
		}
		return elt, true
	case "children":
		children := xp.Value.Children()
		n := &xmlNodes{n: children}
		return n, true
	case "append":
		return object.NewBuiltin("append", xp.append), true
	case "xml":
		return object.NewString(xp.Value.ToXML()), true
	}

	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (xp *XMLDocument) SetAttr(name string, value object.Object) error {
	return object.Errorf("xml.document: cannot set attribute %s", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (xp *XMLDocument) IsTruthy() bool {
	return xp.Value != nil
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (xp *XMLDocument) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("xml.document: operation %s not supported", opType)
}

// Cost returns the incremental processing cost of this object.
func (xp *XMLDocument) Cost() int {
	return 0
}
