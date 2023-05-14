package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Customer struct {
	Name    string
	Balance int
}

var customers []Customer
var loggedInCustomer *Customer

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "register":
			register(args[1:])
		case "login":
			login(args[1])
		case "deposit":
			amount, _ := strconv.Atoi(args[1])
			deposit(amount)
		case "withdraw":
			amount, _ := strconv.Atoi(args[1])
			withdraw(amount)
		case "transfer":
			target := args[1]
			amount, _ := strconv.Atoi(args[2])
			transfer(target, amount)
		case "logout":
			logout()
		default:
			fmt.Println("Invalid command")
		}
	}
}

func register(args []string) {
	name := args[0]
	for _, customer := range customers {
		if customer.Name == name {
			fmt.Println("The user has already been registered")
			return
		}
	}
	customers = append(customers, Customer{Name: name})
	fmt.Println("Registered user:", name)
}

func login(name string) {
	for i, customer := range customers {
		if customer.Name == name {
			loggedInCustomer = &customers[i]
			fmt.Println("Hello,", name, "!")
			fmt.Println("Your balance is $", loggedInCustomer.Balance)
			return
		}
	}
	fmt.Println("The user has not been registered")
}

func deposit(amount int) {
	if loggedInCustomer == nil {
		fmt.Println("Please log in first")
		return
	}
	loggedInCustomer.Balance += amount
	fmt.Println("Your balance is $", loggedInCustomer.Balance)
}

func withdraw(amount int) {
	if loggedInCustomer == nil {
		fmt.Println("Please log in first")
		return
	}
	if loggedInCustomer.Balance < amount {
		fmt.Println("Insufficient balance")
		return
	}
	loggedInCustomer.Balance -= amount
	fmt.Println("Your balance is $", loggedInCustomer.Balance)
}

func transfer(target string, amount int) {
	if loggedInCustomer == nil {
		fmt.Println("Please log in first")
		return
	}
	var targetCustomer *Customer
	for i, customer := range customers {
		if customer.Name == target {
			targetCustomer = &customers[i]
			break
		}
	}
	if targetCustomer == nil {
		fmt.Println("Target user not found")
		return
	}
	if loggedInCustomer.Balance < amount {
		fmt.Println("Insufficient balance")
		return
	}
	loggedInCustomer.Balance -= amount
	targetCustomer.Balance += amount
	fmt.Println("Transferred $", amount, " to", target)
	fmt.Println("Your balance is $", loggedInCustomer.Balance)
	if loggedInCustomer.Balance == 0 {
		fmt.Println("Owed $", targetCustomer.Balance, " to", target)
	}
}

func logout() {
	if loggedInCustomer == nil {
		fmt.Println("Please log in first")
		return
	}
	fmt.Println("Goodbye,", loggedInCustomer.Name, "!")
	loggedInCustomer = nil
}
