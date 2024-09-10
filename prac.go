package main

import "fmt"

func main(){
	// for i:=0;i<5;i++{
	// 	fmt.Println(i)
	// }

   var temp ,sum int 

   for i:=0; i<5;i++{
	fmt.Scanf("%d\n",&temp)
	sum=sum+temp
   }
   fmt.Printf("%d",sum)


   //while loop like for loop 
   i:=0;
   for i<5{
	fmt.Scanf("%d\n",&temp)
	sum=sum+temp
	i++
   }
   fmt.Printf("%d",sum)
   }

   to exit infinite loop 

   we can use break