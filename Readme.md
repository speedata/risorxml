
## Reading XML

```xml
<?xml version="1.0" encoding="UTF-8"?>
<?procinst foobar?>

<!-- a
 multiline
comment -->
<data xmlns:ns="mynamespace" ns:foo="hello" a="b">
    <ns:p>hello world!</ns:p>
    <p>hello XML!</p>
</data>
```

`main.risor`:

```
import xml

// xmldoc := xml.parse("data.xml")
// or
r := os.open("data.xml")
xmldoc := xml.parse(r)

func dump(elt) {
    for _, child := range elt.children {
        if child.is_element {
            printf("Element: %q, namespace: %q\n", child.name, child.namespace)
            dump(child)
        } else if child.is_text {
            printf("Text: %q\n", child.text)
        } else if child.is_comment {
            printf("Comment: %q\n", child.text)
        } else if child.is_procinst {
            printf("Processing instruction: target %q, text %q\n",child.target, child.text)
        }
    }
}


dump(xmldoc)
```

prints

```
Processing instruction: target "xml", text "version=\"1.0\" encoding=\"UTF-8\""
Text: "\n"
Processing instruction: target "procinst", text "foobar"
Text: "\n\n"
Comment: " a\n multiline\ncomment "
Text: "\n"
Element: "data", namespace: ""
Text: "\n    "
Element: "p", namespace: "mynamespace"
Text: "hello world!"
Text: "\n    "
Element: "p", namespace: ""
Text: "hello XML!"
Text: "\n"
Text: "\n"
```


## Creating XML

```
import xml

xmldoc := xml.new_document()
root := xml.new_element("root")


chardata := xml.new_chardata("\n   ")
root.append(chardata)

elt1 := xml.new_element("elt1")
elt1.append(xml.new_attribute("name","value"))

root.append(elt1)
root.append(chardata)


elt2 := xml.new_element("element")
elt2.append(xml.new_attribute("name",`va<l"u&e`))
root.append(elt2)
root.append(xml.new_chardata("\n"))

xmldoc.append(root)

print(xmldoc.xml)
```

creates this XML text:

```xml
<root>
   <elt1 name="value" />
   <element name="va&lt;l&quot;u&amp;e" />
</root>
```



