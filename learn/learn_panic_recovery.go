package main

//
//func a() {
//	fmt.Println("Inside a()")
//	defer func() {
//		if c := recover(); c != nil {
//			fmt.Println("Recovered in a()")
//		}
//	}()
//	fmt.Println("About to call b()")
//	b()
//	fmt.Println("b() exited")  // not will continue
//	fmt.Println("Exiting a()") // not wll continue
//}
//
//func b() {
//	fmt.Println("Inside b()")
//	defer fmt.Println("defer b()")
//	panic("Panic in b()")
//	fmt.Println("Exiting b()")
//}
//
//func main() {
//	a()
//	fmt.Println("main() ended")
//
//	fmt.Print("You are using ", runtime.Compiler, " ")
//	fmt.Println("on a", runtime.GOARCH, "machine")
//	fmt.Println("Using Go version", runtime.Version())
//	fmt.Println("Number of CPUs:", runtime.NumCPU())
//	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
//}
//
////
