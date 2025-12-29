// In the consumer package
package consumer

import "fmt"

// Define the interface here, with only the methods needed
type NameGetter interface {
	GetName() string
}

func ProcessPerson(p NameGetter) {
	// Only uses the GetName method
	fmt.Println(p.GetName())
}
