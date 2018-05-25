package main
import "fmt"
import "time"
func one() {
	for i:=1;i<=10;i++ {
	time.Sleep(1*time.Second)
	fmt.Println(i)
	}
}
func ten() {
	for i:=11;i<=20;i++ {
	time.Sleep(3*time.Second)
	fmt.Println(i)
	}
}
func twenty() {
	for i:=21;i<=30;i++ {
	time.Sleep(5*time.Second)
	fmt.Println(i)
	}
}
func tirty() {
	for i:=31;i<=40;i++ {
	time.Sleep(7*time.Second)
	fmt.Println(i)
	}
}
func forty() {
	for i:=41;i<=50;i++ {
	time.Sleep(9*time.Second)
	fmt.Println(i)
	}
}
func main() {
go one()
go ten()
go twenty()
go tirty()
go forty()
time.Sleep(30*time.Second)
fmt.Println("end")
}