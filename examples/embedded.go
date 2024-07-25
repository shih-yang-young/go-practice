package examples

import "fmt"

type Employee struct {
	Name string
	Id   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.Id)
}

type Manager struct {
	Employee
	Reports []Employee
}

func GoEmbedded() {
	m := Manager{Employee{"John Doe", "12345"}, []Employee{{"Jane Smith", "67890"}}}
	fmt.Println(m.Name)          // John Doe
	fmt.Println(m.Id)            // 12345
	fmt.Println(m.Description()) // John Doe (12345)
	m.Employee.Name = "Louis Chen"
	fmt.Println(m.Employee.Description()) // Louis Chen (12345)
	m.Reports[0].Name = "Sky Chen"
	fmt.Println(m.Reports[0].Description()) // Sky Chen (67890)
}
