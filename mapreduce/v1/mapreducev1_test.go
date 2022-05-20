package v1

import (
	"fmt"
	"testing"
)

func TestMapReduce(t *testing.T)  {
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people: %d\n", old)
	//old people: 2

	high_pay := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Salary >= 6000
	})
	fmt.Printf("High Salary people: %d\n", high_pay)
	//High Salary people: 4

	no_vacation := EmployeeFilterIn(list, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("People no vacation: %v\n", no_vacation)
	//People no vacation: [{Hao 44 0 8000} {Jack 26 0 4000} {Marry 29 0 6000}]

	total_pay := EmployeeSumIf(list, func(e *Employee) int {
		return e.Salary
	})

	fmt.Printf("Total Salary: %d\n", total_pay)
	//Total Salary: 43500

	/*younger_pay := EmployeeSumIf(list, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})*/
}
