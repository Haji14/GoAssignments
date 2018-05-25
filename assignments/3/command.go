package main
import "fmt"
func main() {
	var name string 
	var marks[3] int
	fmt.Println("Enter the name")
	fmt.Scanln(&name)
	fmt.Println("Enter the marks")
	for i:=0;i<3;i++ {
		fmt.Scanln(&marks[i])
	}
	fmt.Println("Name: ",name)
	fmt.Println("Sum is ",marks[0]+marks[1]+marks[2])
	fmt.Println("Average is ",marks[0]+marks[1]+marks[2]/3)
}