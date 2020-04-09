package main

import (
	"fmt"

	uuid "github.com/google/uuid"
)

func main() {
	u, err := uuid.NewUUID()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("UUID generated: %s\n", u)
}
