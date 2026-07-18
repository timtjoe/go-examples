package intermediate

import (
	"fmt"
)

type Animal struct {
    Name string
}

func (a Animal) Greet() string {
    return "Hi, i'am " + a.Name
}

type Dog struct {
    Animal
    Breed string
}

type Cat struct {
    Animal
}

func (c Cat) Greet() string {
    return "Meow. I'm a " + c.Name
}

type Reader interface {
    Read(p []byte) (int, error)
}

type Writer interface {
    Write(p []byte) (int, error)
}

type ReadWrite interface {
    Reader 
    Writer
}

// Multiple embedding
type Server struct {
	
}


func Embed() {
    dog := Dog{Animal: Animal{Name: "Rex"}, Breed: "Lab"}
    cat := Cat{Animal: Animal{Name: "Cathrine"}}

    fmt.Println(dog.Name)
    fmt.Println(dog.Greet())
    fmt.Println(dog.Animal.Name)

    fmt.Println(cat.Name)
    fmt.Println(cat.Greet())
    fmt.Println(cat.Animal.Name)
    fmt.Println(cat.Animal.Greet())
}
