package main  
import (  
   "sync"  
    "fmt"  
)  
var wait sync.WaitGroup  
 var openbalance int=10000  
var mutex sync.Mutex  
func  bank(s string,amount int)  {  
	var balance int
                 mutex.Lock()  
                  if(s=="deposit") {
                balance=openbalance+amount  
                   } else if (s=="withdrawal") {
                                balance=openbalance-amount  
                                  }
                      fmt.Println("Balance: ",balance)  
                  mutex.Unlock() 
                  wait.Done()  
      }  
func main(){  
               
                wait.Add(2)  
                for{
     fmt.Println("1.deposit\n2.Withdrawal\n3.Exit\nSelect any option")
                var choice int
                fmt.Scanln(&choice)
                if(choice==1) {
                fmt.Println("Enter the AMOUNT")
                 var x int
                 fmt.Scanln(&x)
                 go bank("deposit",x)  
                 } else if(choice==2) {
                 fmt.Println("Enter the AMOUNT")
                     var y int
                  fmt.Scanln(&y)
               go bank("withdrawal",y)
                  }else if(choice==3) {
					   break;
                                
  
                }
  
}


   wait.Wait()  

}  
