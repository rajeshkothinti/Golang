package main
import ("fmt")

func main(){
	fmt.Println(addmul(2,3))
	fmt.Println(addm(3,4))
	fmt.Println(sumofnumber(2,3,4,5))

}

//multiple return
func addmul(a int , b int) (int, int){
	return a+b, a*b
}
// named return
func addm(a int,b int)(s int, m int){
	s = a+b
	m = a*b 
	return
}

// variadic function 

func sumofnumber(c ...int) int {
	su:=0
	y := len(c)
	x:=0
	for x<y{
	su = su + c[x]
	x++
	}
	return su
}

