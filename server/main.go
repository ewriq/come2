package main

import (
	"fmt"

	"github.com/ewriq/come2/database"
)

func main() {
	if err := database.InitDB(); err != nil {
		fmt.Println("Database init failed:", err)
	}
	db := database.GetDB()
	fmt.Println("DB initialized:", db != nil)
}
