// Go interfaces are structurally typed - a type satisfied an interface just by having its methods. No implements keyword.

// Keep interfaces small. Many of the most useful Go interfaces have just one method: Stringer, Reader, Writer, error.

package intermediate

import "fmt"

type Stringer interface {
	String() string
}

type Point struct {
	X, Y int
}

// Just by having a String() method, Point satisfies Stringer
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func describe(v any) {
	fmt.Printf("%T: %v\n", v, v)
}

// Type assertion extracts the concrete type:
var i any = "Hello"
func Assertions() {
	s := i.(string)  //panic if i is not string
	s, ok := i.(string) //ok-form: doesn't panic, ok=false on failure
	// Type switch dispatches on the concrete type:
	switch v := i.(type) {
		case int: fmt.Printf("int: %d\n", v)
		case string: fmt.Printf("string: %q\n", v)
		case default: fmt.Printf("other: %v\n", v)
		
	}
}
