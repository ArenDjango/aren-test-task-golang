package main

// check preempting

// #1 we will se finished

//func main() {
//	runtime.GOMAXPROCS(1)
//
//	done := false
//
//	go func() {
//		done = true
//	}()
//
//	for !done {
//	}
//	fmt.Println("finished")
//}

// #2
//func main() {
//	data := make([]int, 4)
//	loopData := func(handleData chan<- int) {
//		defer close(handleData)
//		for i := range data {
//			handleData <- data[i]
//		}
//	}
//	handleData := make(chan int)
//	go loopData(handleData)
//	for num := range handleData {
//		fmt.Println(num)
//	}
//}

// confinement pattern
//
//type Counter struct {
//	value int
//}
//
//func main() {
//	increment := make(chan struct{})
//	getValue := make(chan int)
//
//	// Counter goroutine
//	go func() {
//		counter := &Counter{}
//		for {
//			select {
//			case <-increment:
//				counter.value++
//			case getValue <- counter.value:
//			}
//		}
//	}()
//
//	// Increment counter
//	go func() {
//		for {
//			increment <- struct{}{}
//			time.Sleep(1 * time.Second)
//		}
//	}()
//
//	// Retrieve counter value
//	go func() {
//		for {
//			fmt.Println(<-getValue)
//			time.Sleep(1 * time.Second)
//		}
//	}()
//
//	// Prevent main from exiting immediately
//	select {}
//}

//#3 error handling

//func main() {
//
//	type Result struct {
//		Error    error
//		Response *http.Response
//	}
//
//	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
//		results := make(chan Result)
//		go func() {
//			defer close(results)
//
//			for _, url := range urls {
//				var result Result
//				resp, err := http.Get(url)
//				result = Result{Error: err, Response: resp}
//				select {
//				case <-done:
//					return
//				case results <- result:
//				}
//			}
//		}()
//		return results
//	}
//
//	done := make(chan interface{})
//	defer close(done)
//
//	urls := []string{"https://www.google.com", "https://badhost"}
//	for result := range checkStatus(done, urls...) {
//		if result.Error != nil {
//			fmt.Printf("error: %v", result.Error)
//			continue
//		}
//		fmt.Printf("Response: %v - %v\n", result.Response.Request.URL, result.Response.Status)
//	}
//
//}

//
