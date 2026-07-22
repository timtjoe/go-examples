// Wrapping, Inspecting, and Comparing errors across boundaries
package advance

import (
	"errors"
	"fmt"
	"os"
)

type DB struct{}

func (d DB) Query(query string) ([]string, error) {
	if query == ""{
		return nil, fmt.Errorf("empty query")
	}
	return []string{
		fmt.Sprintf("executed: %s", query),
		"row 1",
		"row 3",
	}, nil
}

// defining sentinel errors
	var (
		ErrNotFound 		= errors.New("not found")
		ErrUnauthorized = errors.New("unauthorized")
	)

	func loopup(id string)(*User, error){
		fmt.Print(id)
			return nil, ErrNotFound
	}

	type ValidationError struct {
		Field string
		Rule	string 
	}

	func (v *ValidationError) Error() string {
		return fmt.Sprintf("%s failed %s", v.Field, v.Rule)
	}

	type QueryError struct {
		Query	string 
		Err 	error
	}

	func (q *QueryError) Error() string {
		return fmt.Sprintf("query: %q: %v", q.Query, q.Err)
	}

	func (q *QueryError) Unwrap() error {return q.Err}

	// main function
func ErrorExample(){
	db := DB{}

	results, err := db.Query("SELECT * FROM people");
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	for _, r := range results {
		fmt.Println(r)
	}

	// Error.Is - sentinel comparison
	f, err := os.Open("config.toml")
	if errors.Is(err, os.ErrNotExist) {
		// file is missing - handle differently
	}
	fmt.Println(f)

	// error.As - type assertion through the chain
	// for extracting a specific error TYPE from anywhere in the chain:
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Println("path: ", pathErr.Path);
		fmt.Println("op:", pathErr.Op);
	}

	if _, err := loopup("x");
	errors.Is(err, ErrNotFound) {
		//handle missing
	}

	var ve *ValidationError 
	if errors.As(err, &ve) {
		fmt.Println("bad field:", ve.Field);
	}

}