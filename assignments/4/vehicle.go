package main
import "fmt"
type employees struct {
	name string
	desg string
	sal int
	v vehicle
}
type vehicle[] string
func main() {
	var person=map[int]employees { 1:employees{"a","PA",25000,vehicle{"duke","honda"},},
	  2:employees{"b","PA",25000,vehicle{"r15","skoda"},},
	  3:employees{"c","PA",25000,vehicle{"ninja"},},}
		 for _,value:=range person {
			 fmt.Println(value.name,value.desg,value.sal,value.v,"\n")
		 }
}