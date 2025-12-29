// In the implementation package
package person

type Person struct {
	Name string
	Age  int
}

func (p Person) GetName() string { return p.Name }
func (p Person) GetAge() int     { return p.Age }
