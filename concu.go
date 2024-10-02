package main 
import ("fmt"
        "time"
	     "sync")
var wg sync.WaitGroup
func calculate(){

	for i:=0;i<500;i++{
		fmt.Println("square is",i*i)
	}
	wg.Done()
}

func main(){
	
start := time.Now()
     wg.Add(1)
	 go calculate()
 //time.Sleep(1 * time.Second)

elapsed := time.Since(start)
wg.Wait()
fmt.Println("total time it took is:",elapsed)

}