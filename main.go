package main 
import (
	"fmt"
	"time"
)

func calcu(i int){
	time.Sleep(1 * time.Second)
	fmt.Println(i * 1)
}
func main(){
	start := time.Now()
	for i := 1; i < 11; i++ {
		go calcu(i)
	}
	time.Sleep( 2 * time.Second)
	during := start.Sub(time.Now())
	fmt.Println(during)

}