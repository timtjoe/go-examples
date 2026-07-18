package intermediate

import "fmt"

type Counter struct {
	count int
}

// Value receiver - operates on a COPY
func (c Counter) GetValue() int {
	return c.count
}

// Pointer reciever = operates on the original
func (c *Counter) Increment() {
	c.count++
}



func UseCounter() {
	counter := Counter{count: 0}
	counter.Increment() //works = Go auto-addresses
	fmt.Println(counter.count) // 1

	// The Map element trap
m := map[string]Counter{"a": {count: 0}}
// m["a"].count++ // compile error - can't take address of mapped value
// Workaround: replace the whole entry
c := m["a"]
c.count++ 
m["a"] = c
}

// Interface satisfaction
type Greeter interface {
	Greet()
}

type IDog struct{
	Name string
}

func (d IDog) Greet() {} //method on Dog
func (d *IDog) Bark() {} //method on *Dog

var g Greeter = IDog{} // Dog has greet (value receiver)
var _g Greeter = &IDog{} // *Dog also has Greet