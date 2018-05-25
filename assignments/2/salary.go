package main
import  
	"fmt"
func main() {
	var salary[5] int 
	var highest,lowest int
	var high,low int
	fmt.Println("Enter the salries of 5 employees")
	for i:=0;i<5;i++ {
		fmt.Scanln(&salary[i])
	}
	highest,lowest=sal(salary)
	fmt.Println(high,low)
	}
func sal(salary[5] int) int  {
	for i:=0;i<5;i++ {
        for j:=i+1;j<5;j++ {
            if(salary[i]<salary[j]) {
                temp:=salary[i]
                salary[i]=salary[j]
                salary[j]=temp
            }
        }
	}
	high=salary[0]
	low=salary[4]
	return high,low
}

	

