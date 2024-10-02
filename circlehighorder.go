package main
import "fmt"

//area, premimeter,circu
func area(r float64) float64{
	return 3.14 *r * r
}
func perimeter(r float64) float64{
	return 2 * 3.14 * r 
}
func circumference(r float64) float64{
	return 3.14*r 
}
func main(){
	var i int
	var r float64
	fmt.Println("enter radius of circle")
	fmt.Scanf("%f", &r)
	fmt.Println("enter 1 for area\n 2 for perimeter \n 3 for circumference")
	fmt.Scanf("%d",&i)
	if i==1{
		fmt.Println("area is :",area(r))
	}
	if i ==2{
		fmt.Println("perimeter is: ",perimeter(r))
	}
	if i ==3{
		fmt.Println("circumference is: ",circumference(r))
	}
}