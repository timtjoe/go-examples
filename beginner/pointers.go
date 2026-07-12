package beginner

import "fmt"

// Pointers as function parameter
// Go passes everything by VALUE.
func increment(n int) {
	n++ // modified the LOCAL copy; caller's n is unchanged
	fmt.Println(n)
}

func incrementPtr(n *int) {
	*n++ // modified through pointer; caller's is chagned
}


// Pointers to structs
// Go automatically dereferences for field access
type pUser struct {
	Name string 
	Age int
}

var p *int //decleared but unintialized - nil

func PointerMain() {
	p := new(int)
	*p = 42

	type User struct{Name string; Age int}
	u := new(User) // *User with empty Name and 0 Age
	u.Name = "Ada"
}


