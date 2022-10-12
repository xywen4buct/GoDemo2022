package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	fmt.Println("jobs", jobs)
	fmt.Println("resutls", results)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
}

/**

func test(){
	messages := make(chan string)
	signals := make(chan string)

	go func() {messages <- "hi"} ()
	go func() {signals <- "how"} ()


	select {
		case msg := <-messages:
			fmt.Println("received message 1", msg)
		//default:
			//	fmt.Println("no message received")
	}

	//select {
	//	case messages <- msg:
	//		fmt.Println("sent message", msg)
		//default:
		//	fmt.Println("no message sent")
	//}

	select {
		case msg:= <-messages:
			fmt.Println("received message 2", msg)
		case msg := <-signals:
			fmt.Println("received signal 3", msg)
		//default:
		//fmt.Println("no activity")
	}

}

func main() {
	test()
}

**/
