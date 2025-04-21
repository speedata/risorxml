package risorxml

import (
	"context"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
	"github.com/speedata/goxml"
)

type xmlDocument struct {
	xmlDocument *goxml.XMLDocument
}

// Type of the object.
func (xp *xmlDocument) Type() object.Type {
	return "xml.document"
}

// Inspect returns a string representation of the given object.
func (xp *xmlDocument) Inspect() string {
	return xp.xmlDocument.String()
}

// Interface converts the given object to a native Go value.
func (xp *xmlDocument) Interface() interface{} {
	return xp.xmlDocument
}

// Returns True if the given object is equal to this object.
func (xp *xmlDocument) Equals(other object.Object) object.Object {
	return object.NewBool(xp.xmlDocument == other.(*xmlDocument).xmlDocument)
}

func (xp *xmlDocument) append(ctx context.Context, args ...object.Object) object.Object {
	for _, arg := range args {
		typ := arg.Type()
		switch typ {
		case "xml.element":
			elt := arg.(*XMLElement)
			xp.xmlDocument.Append(elt.Value)
		case "xml.attribute":
			attr := arg.(*XMLAttribute)
			xp.xmlDocument.Append(attr.Value)
		default:
			// fmt.Printf("~~> t %#v\n", typ)
		}
	}
	return nil

}

// GetAttr returns the attribute with the given name from this object.
func (xp *xmlDocument) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "root":
		r, err := xp.xmlDocument.Root()
		if err != nil {
			return object.NewError(err), true
		}
		elt := &XMLElement{
			Value: r,
		}
		return elt, true
	case "children":
		children := xp.xmlDocument.Children()
		n := &xmlNodes{n: children}
		return n, true
	case "append":
		return object.NewBuiltin("append", xp.append), true
	case "xml":
		return object.NewString(xp.xmlDocument.ToXML()), true
	}

	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (xp *xmlDocument) SetAttr(name string, value object.Object) error {
	return object.Errorf("xml.document: cannot set attribute %s", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (xp *xmlDocument) IsTruthy() bool {
	return xp.xmlDocument != nil
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (xp *xmlDocument) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("xml.document: operation %s not supported", opType)
}

// Cost returns the incremental processing cost of this object.
func (xp *xmlDocument) Cost() int {
	return 0
}
