package main 
import "fmt"
func main(){
	var num int = 10 
	num1 := 20
	var fname string = "rajesh"
	lname := "reddy"
	var ismale = true
// declared int ,string, bool 
// to declare float , array, slice, map,
// to use operators - to do
    var f float64 = 3.22	
	ff := 3.33
	var cars = [3]string{"tata","ford","mg"}
	numbers := [4]int{1,2,3,4}
	fmt.Println(ff ,f,ismale,fname,lname,num1,num)
	fmt.Println(numbers)
	fmt.Println(cars)


    var mapss =map[string]string {"city":"hyd",}
	fmt.Println(mapss["city"])
	fmt.Println("adding elements to map")
	mapss["address"]="gach"
	fmt.Println(len(mapss))
	fmt.Printf("%T",mapss)
	fmt.Printf("%T \n %T\n",cars,fname)
	mapss["thirdkey"]="thirdvalue"
	fmt.Println("map is: \n",mapss)
	delete(mapss,"address")
	fmt.Println("map after deleting value: ", mapss)
}
