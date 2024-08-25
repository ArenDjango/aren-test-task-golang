package main

//Task: Distributed Task Processing System
//
//Requirements:
//
//1.	Producer Goroutine: Continuously generates tasks and sends them to a channel. Each task is represented as a string (e.g., “Task 1”, “Task 2”, etc.).
//2.	Worker Goroutines: Three worker goroutines process tasks from the task channel.
//3.	Result Collection: A result channel collects the results of processed tasks.
//4.	Graceful Shutdown: Implement a mechanism to gracefully shut down all goroutines upon receiving a signal (e.g., SIGINT).
//5.	Timeout Handling: Implement context-based timeout handling for task processing. If a task takes more than 2 seconds to process, it should be aborted, and the next task should be processed.
//6.	Logging: Log the status of task production, task processing, and shutdown.

//type Task struct {
//	ID int
//}
//
//type Result struct {
//	TaskID int
//	Body   string
//	Error  error
//}

//func main() {
//	taskCh := make(chan Task)
//	resultCh := make(chan Result)
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	numWorkers := runtime.NumCPU()
//
//	go func() {
//		signCh := make(chan os.Signal, 1)
//		signal.Notify(signCh, syscall.SIGINT, syscall.SIGTERM)
//		<-signCh
//		log.Println("Received shutdown signal")
//		cancel()
//		close(taskCh)
//	}()
//
//	wg := sync.WaitGroup{}
//
//	// Producer goroutine
//	go func() {
//		defer close(taskCh)
//		for i := 0; i < 10; i++ {
//			select {
//			case <-ctx.Done():
//				return
//			case taskCh <- Task{ID: i}:
//			}
//			time.Sleep(100 * time.Millisecond)
//		}
//	}()
//
//	// Worker Goroutines
//	for i := 0; i < numWorkers; i++ {
//		wg.Add(1)
//		go worker(ctx, &wg, i, taskCh, resultCh)
//	}
//
//	// Result collector
//
//	go func() {
//		for result := range resultCh {
//			if result.Error != nil {
//				log.Printf("Error processing task %d: %v", result.TaskID, result.Error)
//			} else {
//				log.Printf("Processed task %d: %s", result.TaskID, result.Body)
//			}
//		}
//	}()
//
//	wg.Wait()
//	close(resultCh)
//
//	time.Sleep(4 * time.Second)
//}
//
//func worker(ctx context.Context, wg *sync.WaitGroup, workerID int, taskCh <-chan Task, resultCh chan<- Result) {
//	defer wg.Done()
//	for {
//		select {
//		case <-ctx.Done():
//			return
//		case task, ok := <-taskCh:
//			if !ok {
//				return
//			}
//			// Process task with timeout
//			result := processTask(ctx, task)
//			resultCh <- result
//		}
//	}
//}
//
//func processTask(ctx context.Context, task Task) Result {
//	// Simulate processing with a timeout
//	taskCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
//	defer cancel()
//
//	select {
//	case <-taskCtx.Done():
//		return Result{TaskID: task.ID, Error: taskCtx.Err()}
//	case <-time.After(1 * time.Second): // Simulate work
//		return Result{TaskID: task.ID, Body: fmt.Sprintf("Processed by worker")}
//	}
//}
