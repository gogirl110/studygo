package main

// 学习 chan 无缓冲
// 这里用到 go  for chan  select基本技术

import "time"
import "fmt"

func main() {
	c1 := make(chan int)
	go func() {
		i := 0 
		for {
			time.Sleep(1 * time.Second) 
			c1 <- i
			if i == 10 {
				close(c1)
				break
			}
			i ++
		}
	}()

	for {
		select {
		case n1, ok:= <-c1:
			if !ok{
				return 
			}
			fmt.Println("recvie", n1)
		 default:
			break
		}
	}
}