package main
import "fmt"
func main() {
	var number int
	fmt.Println("Enter a range value")
	fmt.Scanln(&number)
	fmt.Println("The prime numbers are")
	for i:=2;i<number;i++ {
		var count int
		for j:=1;j<=i;j++ {
			if (i%j ==0) {
				count++
			}
		}
		if (count==2) {
			fmt.Println(i)
		}
	}
}
