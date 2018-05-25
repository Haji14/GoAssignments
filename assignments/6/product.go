package main
import "fmt"
import "io/ioutil"
import "strings"
type productdetails struct {
	id string
	name string
	brand string
	price string
}
func main() {
c:=[]productdetails{}
    file,err :=ioutil.ReadFile("file2.txt")
    if err !=nil {
        fmt.Println("error is :",err)
    }
    str :=string (file)

	a:=strings.Split(str, ",")
	for i:=3;i<len(a);i+=4 {
        x:=productdetails{id:a[i-3],name:a[i-2],brand:a[i-1],price:a[i]}
        c=append(c,x)
	}
	fmt.Println("Enter the id")
    var find string
    fmt.Scanln(&find)
for i:=0;i<len(c);i++ {
    p:=c[i]
    if(p.id==find){
fmt.Println(p.price)    
}
}
}
