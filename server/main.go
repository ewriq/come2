package main

import (
	"fmt"

	"github.com/ewriq/come2/database"
)

func main() {
	if (database.InitDB()) != nil {
		fmt.Println("Database init failed")
	}
}
