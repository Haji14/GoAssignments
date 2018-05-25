
package main
import "fmt"
func main() {
    var children[10] string
    var count int
    fmt.Println("Enter the names of 10 children")
    for i:=0;i<10;i++ {
        fmt.Scanln(&children[i])
    }
    for i:=0;i<10;i++ {
        for _, value := range children[i] {
            switch value {
			case 'a':
				count+=1
			case 'A':
				count+=1
			case 'e':
				count+=1
			case 'i':
				count+=2
			case 'I':
				count+=2
			case 'o':
				count+=3
			case 'O':
				count+=3
			case 'u':
				count+=4
			case 'U':
				count+=4
            }
        }

        
    }
    rgift:=50-count
    fmt.Println("Remaining gifts =",rgift)

}