package conditions

import "fmt"

// Must function checks if the error given by the caller is nil or not.
// If not, it will display a line in stdout.
func Must(err error) {
	if err != nil {
		fmt.Printf("Error - %v\n", err)
	}
}
