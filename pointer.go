package main
import "fmt"

func main(){
	a := 10
	p := &a
	//initialize string pointe 
	//var s *string ="rajesh"
	fmt.Println("value of p:",*p)
	//a = &p 
	fmt.Println("value of a is:",a)
	//update p value to 20 
	*p = 20
	fmt.Println("value of a is:",a)
	s := "rajesh"
	st := &s
	
	modifys(st)
	fmt.Println("value of s is: ",s)
	fmt.Println("address of a :", p)
	fmt.Println("address of p:",&a)

}

func modifys( s *string) {
       //s := *string
	   *s = "reddy"



}