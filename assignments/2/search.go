package main
import "fmt"
func linearsearch(arr [5] int,search int) int {
	for i,item:=range arr {
		if item==search {
			return i
		} 
		
	}
return 1
}
func main() {
	var arr[5] int 
	var search int
	fmt.Println("Enter the numbers")
	for i:=0;i<5;i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println("Enter the element to be searched")
	fmt.Scanln(&search)
	value:=linearsearch(arr,search)
	if(value==1) {
		fmt.Println("Search Failed")
	} else {
		fmt.Println("Successfully searched")
	}
}