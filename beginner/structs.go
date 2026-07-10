// A struct group field into a single type.
// Structs with method are how to build types with state
package beginner

import "fmt"

type User struct {
	Name string
	Age int
}

func (u User) Greet() string {
	return "Hi " + u.Name
}

func Structs() {
	u := User{Name: "Alice", Age: 30}
	fmt.Println(u.Name)		// Alice
	u.Age = 31 						//mutate
}

