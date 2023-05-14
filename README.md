# software-engineer-test

## **Problem Statement**

You are asked to develop a Command Line Interface (CLI) to simulate an interaction of an ATM with a retail bank. 

## Commands

- `register [name]` - Register as create subscriber if not exist
- `login [name]` - Login as this customer and creates the customer if not exist
- `deposit [amount]` - Deposits this amount to the logged in customer
- `withdraw [amount]` - Withdraws this amount from the logged in customer
- `transfer [target] [amount]` - Transfers this amount from the logged in customer to the target customer
- `logout` - Logout of the current customer


## In the project directory, you can run
```
$ go run main.go
```

## Example session

Your console output should contain at least the following output depending on the scenario and commands. But feel free to add extra output as you see fit.

```
$ register vini
Registered user: vini

$ login vini
Hello, vini !
Your balance is $ 0

$ deposit 100
Your balance is $ 100

$ logout
Goodbye, vini !

$ login bob
Hello, bob !
Your balance is $ 0

$ deposit 80
Your balance is $ 80

$ transfer vini 50
Transferred $ 50  to vini
Your balance is $ 30

$ transfer vini 100
Insufficient balance

$ deposit 30
Your balance is $ 60

$ transfer Alice 10
Target user not found

$ logout
Goodbye, bob !

$ login vini
Hello, vini !
Your balance is $ 150

$ logout
Goodbye, vini !

```