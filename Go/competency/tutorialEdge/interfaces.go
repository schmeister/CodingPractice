package tutorialEdge

import "fmt"

type Employee interface {
    Language() string
    Age() int
}

type Engineer struct {
    Name string
}

func (e Engineer) Language() string {
    return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
    return 23
}

func (e Engineer) Something() int {
    return 23
}


func Interface() {
    // This will throw an error
    var programmers []Employee
    elliot := Engineer{Name: "Elliot"}
    // Engineer does not implement the Employee interface
    // you'll need to implement Age() and Random()
	programmers = make([]Employee,0)
    programmers = append(programmers, elliot)
	fmt.Println(programmers)
	fmt.Println(elliot.Something())
}
