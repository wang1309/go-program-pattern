package visitor

import "testing"

func TestVisitor(t *testing.T) {
	c := Circle{10}
	r :=  Rectangle{100, 200}
	shapes := []Shape{c, r}
	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}
}
