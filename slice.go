// //declare slice 
package  main 
import "fmt"


var beers = []string {}

func main(){
var fruits = []string{"apple","banana","goa"}
fmt.Println(fruits)
_ = beers
//add element in index 1 in slice
// how to print index 1 
// print from 1 to 3 elements in slice
fruits = append(fruits, "kiwi","pineapple")
fmt.Println(fruits)
fmt.Println(fruits[1:3])
fmt.Println(fruits[:len(fruits)])
fmt.Println(len(fruits))
// apple banana goa kiwi pineapple 
//delete banana and print it
fruits = append(fruits[:1],fruits[2:len(fruits)]...)
//tfruits := fruits[:len(fruits)]
fmt.Println(fruits)
}

func deleteelement(index int, slice[])



