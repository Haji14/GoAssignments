package main
import "fmt"
func main() {
var a int
var b int
var c int
fmt.Print("Enter a")
fmt.Scanln(&a)
fmt.Print("Enter b")
fmt.Scanln(&b)
fmt.Print("Enter c")
fmt.Scanln(&c)
if (a>b && a>c) {
fmt.Println("a is greater");
} else if (b>a && b>c) {
fmt.Println("b is greater");
}else {
fmt.Println("c is greater");
}
}