package v1

// https://coolshell.cn/articles/21164.html
type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}


var list = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

// EmployeeCountIf
func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i:=0; i<len(list);i++ {

		if fn(&list[i]) {
			count++
		}
	}

	return count
}

// EmployeeFilterIn
func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for i, _ := range list {
		if fn(&list[i]) {
			newList = append(newList, list[i])
		}
	}

	return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i, _ := range list {
		sum += fn(&list[i])
	}
	return sum
}


