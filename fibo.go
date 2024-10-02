package main
import "fmt"
// 3
// 3*2*1
func fibo(a int) int{
	if a == 1{
		return 1
	}
	return a* fibo(a-1) 

}

func main(){
	fmt.Println(fibo(5))
}