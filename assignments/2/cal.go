package main
import "fmt"
func main() {
	var first,second,result float64
	var op string
	fmt.Println("Enter the 1 no")
	fmt.Scanln(&first)
	fmt.Println("Enter the 2 no")
	fmt.Scanln(&second)
	fmt.Println("Enter the operation")
	fmt.Scanln(&op)
	switch(op) {
	case "+":
		result=first+second
		fmt.Println("Value is ",result)
	case "-":
		result=first-second
		fmt.Println("Value is ",result)
	case "*":
		result=first*second
		fmt.Println("Value is ",result)
	case "/":
		result=first/second
		fmt.Println("Value is ",result)
	default:
		fmt.Println("No operation can be done")
	}
}