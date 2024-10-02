package main 
import ("fmt"
        //"time"
	    "sync")

// buy and sell 

var wg sync.WaitGroup

func sell( c chan string){
	c <- "sofa"
	fmt.Println("sent value")
	close(c)
	wg.Done()
}
func buy(c chan string){
    fmt.Println("waiting for value")
	val := <-c 
	fmt.Println("received value",val)
    wg.Done()

}
func main(){

	c := make(chan string)
	wg.Add(2)
	go sell(c)
	go buy(c)
	wg.Wait()
}