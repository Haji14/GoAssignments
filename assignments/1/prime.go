package main
import "fmt"
func main() {
	var number int
	var count int=0
	fmt.Println("Enter number")
	fmt.Scanln(&number)
	for i:=2;i<number;i++ {
		if number%i==0 {
			count++
		}
	}
	if count==0 {
		fmt.Println("It is a prime number")
	} else {
		fmt.Println("It is not a prime number")
	}
}