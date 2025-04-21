package risorxml

import (
	"context"
	"encoding/xml"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
	"github.com/speedata/goxml"
)

// XMLElement represents an XML element in the Risor XML library.
type XMLElement struct {
	Value *goxml.Element
	attr  *object.Map
}

// Type of the object.
func (elt *XMLElement) Type() object.Type {
	return "xml.element"
}

// Inspect returns a string representation of the given object.
func (elt *XMLElement) Inspect() string {
	return elt.Value.String()
}

// Interface converts the given object to a native Go value.
func (elt *XMLElement) Interface() interface{} {
	return elt.Value
}

// Equals returns True if the given object is equal to this object.
func (elt *XMLElement) Equals(other object.Object) object.Object {
	if other.Type() != elt.Type() {
		return object.False
	}
	otherElt := other.(*XMLElement)
	if elt.Value == nil && otherElt.Value == nil {
		return object.True
	}
	if elt.Value == nil || otherElt.Value == nil {
		return object.False
	}
	return object.NewBool(elt.Value == otherElt.Value)
}

func (elt *XMLElement) append(ctx context.Context, args ...object.Object) object.Object {
	for _, arg := range args {
		typ := arg.Type()
		switch typ {
		case "xml.element":
			thisElt := arg.(*XMLElement)
			elt.Value.Append(thisElt.Value)
		case "xml.attribute":
			attr := arg.(*XMLAttribute)
			elt.Value.Append(attr.Value)
		case "xml.text":
			cd := arg.(*xmlText)
			elt.Value.Append(cd.text)
		default:
			// fmt.Println(`~~> typ`, typ)
			return nil
		}
	}
	return nil
}

// GetAttr returns the attribute with the given name from this object.
func (elt *XMLElement) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "append":
		return object.NewBuiltin("xml.elt_append", elt.append), true
	case "attributes":
		attrs := elt.Value.Attributes()
		if elt.attr == nil {
			elt.attr = object.NewMap(nil)
		} else {
			elt.attr = object.NewMap(elt.attr.Value())
		}
		for _, attr := range attrs {
			elt.attr.Set(attr.Name, &XMLAttribute{Value: attr})
		}
		return elt.attr, true
	case "children":
		children := elt.Value.Children()
		n := &xmlNodes{n: children}
		return n, true
	case "is_element":
		return object.True, true
	case "is_text":
		return object.False, true
	case "is_comment":
		return object.False, true
	case "is_procinst":
		return object.False, true
	case "name":
		return object.NewString(elt.Value.Name), true
	case "namespace":
		return object.NewString(elt.Value.Namespaces[elt.Value.Prefix]), true
	case "stringvalue":
		return object.NewString(elt.Value.Stringvalue()), true
	case "xml":
		if elt.attr != nil {
			for _, value := range elt.attr.Value() {
				attr := value.(*XMLAttribute)
				attrName := xml.Name{Local: attr.Value.Name}
				if attr.Value.Prefix != "" {
					attrName.Space = attr.Value.Namespace
				}
				elt.Value.SetAttribute(xml.Attr{Name: attrName, Value: attr.Value.Value})
			}
		}
		return object.NewString(elt.Value.ToXML()), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (elt *XMLElement) SetAttr(name string, value object.Object) error {
	switch name {
	case "name":
		elt.Value.Name = value.Interface().(string)
	}
	return nil
}

// IsTruthy returns true if the object is considered "truthy".
func (elt *XMLElement) IsTruthy() bool {
	return elt.Value.Name != ""
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (elt *XMLElement) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("xml.element: operation %s not supported", opType)
}

// Cost returns the incremental processing cost of this object.
func (elt *XMLElement) Cost() int {
	return 0
}
