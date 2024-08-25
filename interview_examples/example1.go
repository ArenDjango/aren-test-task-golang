package main

//func main() {
//	wg := sync.WaitGroup{}
//	ch := make(chan string) // нельзя использовать буферизированный канал
//
//	wg.Add(5)
//	for i := 0; i < 5; i++ {
//		go func(group *sync.WaitGroup, i int) {
//			defer group.Done()
//
//			ch <- fmt.Sprintf("Goroutine %d", i)
//		}(&wg, i)
//	}
//
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	for {
//		select {
//		case s, ok := <-ch:
//			if !ok {
//				return // Exit the infinite loop when the channel is closed
//			}
//			fmt.Println(s)
//	}
//}

//
//Тут правильно закрыть канал и после wg.wait слушать событие закрытия канала
//
//
//Event driven, event sourcing
//
//Зачем Кафка если везде grpc пихать - чтобы одно событие слушали кучу консюмеров и легко масштабировать

// improved variant
//func main() {
//	wg := sync.WaitGroup{}
//	ch := make(chan string) // нельзя использовать буферизированный канал
//	var once sync.Once
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	wg.Add(5)
//	for i := 0; i < 5; i++ {
//		go func(group *sync.WaitGroup, i int) {
//			defer group.Done()
//			select {
//			case ch <- fmt.Sprintf("Goroutine %d", i):
//			case <-ctx.Done():
//				fmt.Println("Context cancelled, goroutine exiting")
//				return
//			}
//		}(&wg, i)
//	}
//
//	go func() {
//		wg.Wait()
//		once.Do(func() {
//			close(ch)
//		})
//	}()
//
//	for {
//		select {
//		case s, ok := <-ch:
//			if !ok {
//				return
//			}
//			fmt.Println(s)
//		case <-ctx.Done():
//			log.Println("Context timeout, main loop exiting")
//			return
//		}
//	}
//}

// improved
//func main() {
//	type Result struct {
//		Body  string
//		Error error
//	}
//
//	urlsCh := make(chan string)
//	resultsCh := make(chan Result)
//	ctx, cancell := context.WithTimeout(context.Background(), 6*time.Second)
//	defer cancell()
//
//	go func() {
//		for i := 0; i < 10; i++ {
//			urlsCh <- "http://www.sdsdfsdf.com?query"
//		}
//		close(urlsCh)
//	}()
//
//	numWorkers := 5
//	wg := sync.WaitGroup{}
//	wg.Add(numWorkers)
//	for i := 0; i < numWorkers; i++ {
//		go func() {
//			defer wg.Done()
//			for url := range urlsCh {
//				_, err := http.NewRequestWithContext(ctx, "GET", url, nil)
//				if err != nil {
//					resultsCh <- Result{Error: err}
//					continue
//				}
//				// Simulate HTTP request (replace with actual HTTP call)i
//				time.Sleep(500 * time.Millisecond) // Simulate delay
//				resultsCh <- Result{Body: "somebody", Error: nil}
//			}
//		}()
//	}
//
//	// Goroutine to close resultsCh when all workers are done
//	go func() {
//		wg.Wait()
//		close(resultsCh)
//	}()
//
//	// WaitGroup to wait for the results printing goroutine
//	var printWg sync.WaitGroup
//	printWg.Add(1)
//
//	// Goroutine to read and print results
//	go func() {
//		defer printWg.Done()
//		for rs := range resultsCh {
//			if rs.Error != nil {
//				log.Println("Error:", rs.Error)
//			} else {
//				fmt.Println("Body:", rs.Body)
//			}
//		}
//	}()
//	// Wait for the results printing goroutine to finish
//	printWg.Wait()
//}
