package main 
import "fmt"
func main() {
	var grosspay float64
	fmt.Println("Enter the grosspay ")
	fmt.Scanln(&grosspay)

	if grosspay<=240 {
		fmt.Println("tax amount is 0")
	} else if(grosspay>240 && grosspay<=480) {
		fmt.Println("tax amount is ",(grosspay*15)/100)
	} else {
		fmt.Println("tax amount is ",(grosspay*28)/100)
	}
}