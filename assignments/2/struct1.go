package main
import "fmt"
type Account struct {
	accnumber int
	balance float64
}
func main() {
	var customer Account
	var accno int
	var bal float64
	var choice int
	fmt.Println("Enter the account number")
    fmt.Scanln(&accno)
    fmt.Println("Enter the deposit")
	fmt.Scanln(&bal)
	customer.accnumber=accno
	customer.balance=bal
	fmt.Println("select\n 1.Deposit\n 2.Withdrawal")
	fmt.Scanln(&choice)
	if(choice==1) {
		deposit(customer)
	} else {
		Withdrawal(customer)
	}

}
func deposit(customer Account) {
	var amount float64
	fmt.Println("Enter account number: ",customer.accnumber)
	fmt.Println("Enter the amount")
	fmt.Scanln(&amount)
	customer.balance=customer.balance+amount
	fmt.Println("New amount: ",customer.balance)
} 
func Withdrawal(customer Account) {
	var amount float64
	fmt.Println("Enter account number: ",customer.accnumber)
	fmt.Println("Enter the amount")
	fmt.Scanln(&amount)
	if (customer.balance>amount) {
	customer.balance=customer.balance-amount
	fmt.Println("New amount: ",customer.balance)
	} else {
		fmt.Println("Insufficient money")
	}
} 