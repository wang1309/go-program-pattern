package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T)  {
	s1 := timedSumFunc(sum1)
	s2 := timedSumFunc(sum2)

	fmt.Println(s1(-10000, 10000000))
	fmt.Println(s2(-10000, 10000000))

}
