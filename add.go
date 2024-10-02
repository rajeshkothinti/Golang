package main 
import "fmt" 

func main() {
	fmt.Println("this is main function")

	var sum int = add(2,3)
fmt.Println(sum)
}

func add(a int, b int) int{
    return a + b 
}

