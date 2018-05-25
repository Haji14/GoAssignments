package main
import "fmt"
type details struct {
	name,address string
}

func main() {
telephone:=make(map[int]details)
telephone[1203456988]=details{"haji","madurai"}
telephone[7845123698]=details{"abbu","chennai"}
fmt.Println("Enter the number")
var search int
fmt.Scanln(&search)
value,yes:=telephone[search]
if yes==true {
	fmt.Println("address of",search,"is",value)
} else {
	fmt.Println("address of",search,"is not found")
}
}

