package main
import "fmt"
type Product struct {
    id int
    name string
    quantity int
    price int
    total_price int
}

func main() {
    var entry Product
    var ide int
    var qty int
    var up int
    var tp int
    var nam string
    fmt.Println("Enter the id")
    fmt.Scanln(&ide)
    fmt.Println("Enter the Name")
    fmt.Scanln(&nam)
    fmt.Println("Enter the quantity")
    fmt.Scanln(&qty)
    fmt.Println("Enter the price") 
    fmt.Scanln(&up)
   
    entry.id=ide
    entry.name=nam
    entry.quantity=qty
    entry.price=up
    tp=qty*up
    entry.total_price=tp
    print(entry)

}
    func print(entry Product) {
        fmt.Println("Id= ",entry.id)
        fmt.Println("Name= ",entry.name)
        fmt.Println("Quantity= ",entry.quantity)
        fmt.Println("Price= ",entry.price)
        fmt.Println("Total= ",entry.total_price)

    }
