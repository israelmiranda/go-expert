package main

import "fmt"

type NotFoundError struct {
	Resource string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found\n", e.Resource)
}

type PermissionError struct {
	Resource string
	User     string
}

func (e PermissionError) Error() string {
	return fmt.Sprintf("user %s does not have permission to access %s\n", e.User, e.Resource)
}

func getResource(name, user string) (string, error) {
	// Check if resource exists
	if name == "nonexistent" {
		return "", NotFoundError{Resource: name}
	}

	// Check permissions
	if user != "admin" && name == "config" {
		return "", PermissionError{Resource: name, User: user}
	}

	return "Resource data", nil
}

func main() {
	if res, err := getResource("config", "user"); err != nil {
		switch e := err.(type) {
		case NotFoundError:
			fmt.Println("Not found error:", e.Error())
		case PermissionError:
			fmt.Println("Permission denied:", e.Error())
		default:
			fmt.Println("Unknown error:", err)
		}
	} else {
		fmt.Println("Got resource:", res)
	}
}
