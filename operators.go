//comparison operators 
package main
import "fmt"

func main() {
	if 1!=2 {
		fmt.Println("yes \n")
	} else if  2>=3{
		fmt.Println("its elseif")
	} else
	{
		fmt.Println("no")
	}
  fmt.Println("sum of 2 and 3 is: ",add(2,3),"\n")
  fmt.Println("subtract of 2 and 3 is: ",sub(2,3),"\n")
  fmt.Println("multiplication of 2 and 3 is: ",mul(2,3),"\n" )
  fmt.Println("divide of 8 and 5 is: ",divide(5,8),"\n")
  fmt.Println("remainder of 8 and 5 is: ",remainder(8,5),"\n")
  //fmt.Println("fruit is grape:",checkfruit("apple","grape"))
  checkfruit("grape","apple")
  fmt.Println(sassignop(2))
  fmt.Println(massignop(5))
  // assign operators are +=, -=.*=,/=,%=,^= ,etc
  fmt.Println("xor of 2 and 2 is: ", xop(2,2))
  // bit wise operators  and(&) is 1 only if both are 1 .or(|) any one is 1 then 1

}
func add(a int,b int) int{
	return a+ b 
}
func sub(a int,b int) int{
	return a-b 
}
func mul(a int, b int) int{
	return a*b 
}
func divide(a int,b int) int{
	return a/b
}
func remainder(a int, b int)int{
	return a%b 
}

// logical operator 
func checkfruit(a string,  b string)   {
	 // fmt.Println(a,b)
	if a == "apple" && b == "apples" {
		fmt.Println("yes they are apples")
		//return true 
	}
	if a == "grape" || b == "banana" {
		fmt.Println("could be banana or grape")
	}
}

func sassignop(a int) int{
	a+=2
	return a
}

func massignop(a int) int{
	a*=2
	return a
}

func xop(a int, b int) int {
	return a^b
}