package testing

import (
	"fmt"

	"github.com/nancyhamdan/golang-study/udemy-courses/golang-basics/testing/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
