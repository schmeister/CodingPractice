package tutorialEdge

import "fmt"


type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	fmt.Println("Implement Me")

	dev := Developer{
		Name: name.(string),
		Age:  age.(int),
	}
	return dev
}
