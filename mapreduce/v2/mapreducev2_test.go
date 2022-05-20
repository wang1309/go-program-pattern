package v2

import (
	"fmt"
	"testing"
)

func TestV2Mapper(t *testing.T) {
	list := []string{"1", "2", "3", "4", "5", "6"}
	result := Transform(list, func(a string) string {
		return a + a + a
	})
	fmt.Println(result)
	//{"111","222","333","444","555","666"}

	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	TransformInPlace(list1, func(a int) int {
		return a * 3
	})
	fmt.Println(list1)
	//{3, 6, 9, 12, 15, 18, 21, 24, 27}

	type Employee struct {
		Name     string
		Age      int
		Vacation int
		Salary   int
	}

	var list2 = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
	}

	result = TransformInPlace(list2, func(e Employee) Employee {
		e.Salary += 1000
		e.Age += 1
		return e
	})

	fmt.Println(result)
}

func TestV2Reduce(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := Reduce(list1, func(a, b int) int {
		return a + b
	}, nil)
	fmt.Println(result)
}

func TestFilter(t *testing.T)  {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := Filter(list1, func(a int) bool {
		return a > 5
	})
	fmt.Println(result)
}