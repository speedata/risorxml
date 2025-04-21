package risorxml

import (
	"context"
	"os"

	"github.com/risor-io/risor/object"
	"github.com/speedata/goxml"
)

// parse parses an XML file and returns an object representing the XML document.
func parse(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("args error: xml.parse() takes exactly 1 argument (%d given)", len(args))
	}
	firstArg := args[0]

	switch firstArg.Type() {
	case "string":
		// a filename
		r, err := os.Open(firstArg.Interface().(string))
		if err != nil {
			return object.NewError(err)
		}
		d, err := goxml.Parse(r)
		if err != nil {
			return object.NewError(err)
		}
		return &xmlDocument{
			xmlDocument: d,
		}
	case "file":
		var rFile *object.File
		var ok bool
		if rFile, ok = firstArg.(*object.File); !ok {
			return object.TypeErrorf("xml.parse() takes a file as argument, got %s", firstArg.Type())
		}
		d, err := goxml.Parse(rFile)
		if err != nil {
			return object.NewError(err)
		}
		return &xmlDocument{
			xmlDocument: d,
		}

	default:
		return object.TypeErrorf("xml.parse() takes a string or file as argument, got %s", firstArg.Type())
	}
}

func newDocument(ctx context.Context, args ...object.Object) object.Object {
	return &xmlDocument{
		xmlDocument: &goxml.XMLDocument{},
	}
}

func newAttribute(ctx context.Context, args ...object.Object) object.Object {
	if len(args) == 2 {
		firstarg := args[0]
		secondarg := args[1]
		if str, ok := firstarg.(*object.String); ok {
			if str2, ok := secondarg.(*object.String); ok {
				return &XMLAttribute{Value: &goxml.Attribute{Name: str.Value(), Value: str2.Value()}}
			}
		}
	}
	return &XMLAttribute{Value: &goxml.Attribute{}}
}

func newElement(ctx context.Context, args ...object.Object) object.Object {
	if len(args) == 1 {
		if str, ok := args[0].(*object.String); ok {
			return &XMLElement{Value: &goxml.Element{Name: str.Value()}}
		}
	}
	return &XMLElement{Value: &goxml.Element{}}
}

func newChardata(ctx context.Context, args ...object.Object) object.Object {
	if len(args) == 1 {
		if str, ok := args[0].(*object.String); ok {
			return &xmlText{text: &goxml.CharData{Contents: str.Value()}}
		}
	}
	// If no argument is provided, create an empty CharData object.
	return &xmlText{text: &goxml.CharData{}}
}

// Module returns a new module with the XML parser functionality.
func Module() *object.Module {
	return object.NewBuiltinsModule("xml", map[string]object.Object{
		"parse":         object.NewBuiltin("xml.parse", parse),
		"new_attribute": object.NewBuiltin("xml.new_attribute", newAttribute),
		"new_document":  object.NewBuiltin("xml.new_document", newDocument),
		"new_element":   object.NewBuiltin("xml.new_element", newElement),
		"new_chardata":  object.NewBuiltin("xml.new_chardata", newChardata),
	})
}
