package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"`
}

func main() {
	acc := Account{Number: 1, Balance: 100}
	bytes, err := json.Marshal(acc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json marshal: %v\n", string(bytes))

	// encoder := json.NewEncoder(os.Stdout)
	// encoder.Encode(acc)

	fmt.Print("json encode: ")
	err = json.
		NewEncoder(os.Stdout).
		Encode(acc)

	if err != nil {
		panic(err)
	}

	jraw := []byte(`{"number":2,"balance":200}`)
	var acc1 Account
	err = json.Unmarshal(jraw, &acc1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json unmarshal: %v\n", acc1)

	reader := strings.NewReader(string(jraw))
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&acc1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("json decode: %v\n", acc1)
}
