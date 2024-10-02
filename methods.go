package main 
import "fmt"
func main(){

	type person struct{
		 name string
		 age int
		 salary float64
		 area string
	}
	// methods are like function but with receiver before function name
	// where is it used; practical example ;does it have return statement
	func (s string) somemethod(){

	}
}

// method, will have receiver usually struct and modify its value 
// example srtuct name is circle ; with radius and area so next time you see
//play with that to check how much you know on methods an
// create instance of circle with pointer so changing at one can change at other so here receiver is pointer . recievr can be value also 