package main
import "fmt"
type persons struct{
	name string
	age int 
	salary float64
}
func main(){
	// type struct students {
	// 	name string
	// 	age int 
	
	
	
	person1 := persons{"raj", 22,10000}
	fmt.Println(person1)
	// try accessing element in struct
	fmt.Println(person1.age)
	// can we modify value in person1
	person1.age = 23
	fmt.Println(person1.age)
	// can we delete age va
}

// func test(){
// 	person2 := persons{"red",23,344444}
// 	//person3 := persons{}
// }