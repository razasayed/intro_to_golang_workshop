package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uuid)
	}
}
