package main 
import "fmt"
func main() {
	var deposit float64
	fmt.Println("Enter the deposit amount")
	fmt.Scanln(&deposit)

	if deposit<=1000 {
		fmt.Println("interest amount is ",(deposit*4)/100)
	} else if(deposit>1000 && deposit<=5000) {
		fmt.Println("interest amount is ",(deposit*4.5)/100)
	} else {
		fmt.Println("interest amount is ",(deposit*5)/100)
	}
}