package main
import "fmt"
func main() {
var str string
fmt.Println("Enter a char")
fmt.Scanln(&str)
if(str=="a" || str=="e" || str=="i" || str=="o" || str=="u" || str=="A" || str=="e" || str=="I" || str=="O" || str == "U") {
    fmt.Println("It is vowel");
} else {
    fmt.Println("It is npot a vowel")
}
}