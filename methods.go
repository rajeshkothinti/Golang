package main 
import "fmt"
func main(){

	type person struct{
		 name string
		 age int
		 salary float64
		 area string
	}
	func (s string) somemethod(){

	}
}

// method, will have receiver usually struct and modify its value 
// example srtuct name is circle ; with radius and area so next time you see
//play with that to check 
// create instance of circle with pointer so changing at one can change at other so here receiver is pointer . recievr can be value also 
