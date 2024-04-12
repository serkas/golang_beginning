package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Stored accounts: ")
		// TODO: implement accounts list

		return
	}

	if args[0] == "add" && len(args) == 2 {
		account := args[1]
		var masterPassword, password string
		fmt.Println("Enter master password")
		_, err := fmt.Scan(&masterPassword)
		if err != nil {
			fmt.Println("Err:", err)
			os.Exit(1)
		}

		fmt.Println("New password for", account)
		_, err = fmt.Scan(&password)
		if err != nil {
			fmt.Println("Err:", err)
			os.Exit(1)
		}

		fmt.Printf("will store password for %s\n", account)

		return
	}
}
