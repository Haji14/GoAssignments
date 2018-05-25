package main
import "fmt"
func main() {
	var name,capital,cm string
	fmt.Println("Enter the city name")
	fmt.Scanln(&name)
	capital,cm =function(name)
	fmt.Printf("%s : capital - %q, cm - %q",name,capital,cm)

}
func function(name string) (string,string) {
	var capital,cm string
	switch name {
	case "AP ,AndhraPradesh":
		capital="Amaravathi"
		cm="Chandrababu Nayudu"
	case "TM, Tamilnadu":
		capital="Chennai"
		cm="Palani Swamy"
	case "TN, Telangana":
		capital="Hyderabad"
		cm="K.Chandra Sekhar Rao"
	}
	return capital,cm
}
