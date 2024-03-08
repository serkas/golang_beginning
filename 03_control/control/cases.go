package main

import (
	"fmt"
	"time"
)

func cases() {
	timestamp := time.Now().Unix() // Unix timestamp of a current moment
	fmt.Printf("Current timestamp: %d\n\n", timestamp)

	// switch cases with variable argument
	role := "any"
	switch role {
	case "user", "any":
		fmt.Println("give user permissions")
	case "admin":
		fmt.Println("give all admin permissions")
	case "superadmin":
		fmt.Println("give all admin permissions and allow to create other admins")

	default:
		fmt.Println("access forbidden. no permissions")
	}

	// switch cases without argument - bool expressions in cases
	switch {
	case timestamp%10 == 0 && role == "any":
		fmt.Printf("10s second %d\n", timestamp)
	case timestamp%5 == 0:
		fmt.Printf("5s second %d\n", timestamp)
	default:
		fmt.Printf("default case %d\n", timestamp)
	}

	fmt.Println()
}
