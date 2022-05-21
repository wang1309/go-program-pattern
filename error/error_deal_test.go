package error

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T)  {
	p := Person{}
	p.ReadName().ReadAge().ReadWeight().Print()
	fmt.Println(p.err)  // EOF 错误
}
