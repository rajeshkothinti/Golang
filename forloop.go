package main 
import (
	"fmt"
	"os"
)

func main(){
// 	for i:=0; i<5; i ++{
// 		fmt.Println(i)
// }
//i :=0
// for i < 5{
// 	fmt.Println(i)
// 	i++
// }
// fruits := []string{"apple", "grapes", "kiwi"}

// for i,v:= range fruits{
// 	fmt.Println("at index ",i,"fruit is",v)
// }
i := os.Args
fmt.Printf("type of i is %T: ", i)
x :=1
fmt.Println("entered arguments are: \n")
for x<len(i){
	
	fmt.Println(i[x])
	x++
}

}