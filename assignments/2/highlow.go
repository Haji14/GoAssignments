package main
import "fmt"
func main() {
var a[10] int
var highest int
var lowest int
fmt.Println("Enter the marks of 10 students")
for i:=0;i<10;i++ {
fmt.Scanln(&a[i])
}
highest=max(a)
fmt.Println(highest)
lowest=min(a)
fmt.Println(lowest)

}
func max(a[10] int)int {
	max:=a[0]
	for i:= 1; i < 10; i++ {
	  if a[i] > max {
		 max  = a[i];
		}
	}
	return max
}
func min(a[10] int )int {
	for i:=0;i<10;i++ {
        for j:=i+1;j<10;j++ {
			if(a[i]<a[j]) {
				temp:=a[i]
				a[i]=a[j]
				a[j]=temp
			}
		}
	}
	l:=a[9]
    return l
}




