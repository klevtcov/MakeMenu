package main

import (
	"encoding/json"
	"fmt"
)

type Ingridients struct {
	Proteins
}

// someStruct := SomeStruct{}  // Read errors caught by unmarshal
// fileBytes, _ := os.ReadFile(filename)
// err := json.Unmarshal(fileBytes, spec)

func main() {

	// var m Message
	// err := json.Unmarshal(b, &m)

	fmt.Println("app started")
}
