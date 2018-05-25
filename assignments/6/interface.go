package main
import "fmt"
type loan interface {
	getemi() float64
}
type HousingLoan struct {
	pamount float64
}
type PersonalLoan struct {
	pamount float64
}
func (h HousingLoan) getemi() float64 {
	var amount float64
	amount=(9*h.pamount)/100
	return amount

}
func (p PersonalLoan) getemi() float64 {
	var amount float64
	amount=(12*p.pamount)/100
return amount
}
func main() {

	hl:=HousingLoan{pamount: 10000}
var l loan
l=hl
fmt.Println("housingloan = ",l.getemi())
		pl:=PersonalLoan{pamount: 10000}
l=pl

fmt.Println("Personalloan = ",l.getemi())
	


	
}