package advance

import (
	"errors"
	"fmt"
)

func readFile() error {
	return errors.New("file not found")
}

func Wrapped(){
	err := readFile()

	if err != nil {
		// wrap the original error with context
		wrapped := fmt.Errorf("readFile fail: %w", err)
		fmt.Println(wrapped)

		// Unwrap to see the original
		fmt.Println(errors.Unwrap(wrapped))
	}
}