package visitor

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Visitor func(shape Shape)

type Shape interface {
	accept(visitor Visitor)
}

type Circle struct {
	Radius int
}

func (c Circle) accept(visitor Visitor)  {
	visitor(c)
}

type Rectangle struct {
	Width, High int
}

func (r Rectangle) accept(visitor Visitor)  {
	visitor(r)
}

func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}