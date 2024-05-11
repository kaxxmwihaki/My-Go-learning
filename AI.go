package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt the user to enter their name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Store the name in a people.txt file on the desktop
	filePath := os.Getenv("HOME") + "/Desktop/people.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(name + "\n"); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name stored successfully!")

	// Ask the user if they want to search for any name in the file
	fmt.Print("Do you want to search for a name in the file? (yes/no): ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if strings.ToLower(answer) == "yes" {
		searchNameInFile(filePath)
	} else {
		fmt.Println("Thank you!")
	}
}

func searchNameInFile(filePath string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name you want to search for: ")
	searchName, _ := reader.ReadString('\n')
	searchName = strings.TrimSpace(searchName)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	found := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == searchName {
			fmt.Println("Name found!")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Name not found.")
	}
}
