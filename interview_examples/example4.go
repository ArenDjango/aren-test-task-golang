package main

//
//func main() {
//	a := []int32{1}           // cap 1
//	a = append(a, 1, 2, 3, 4) // 8
//	fmt.Println(cap(a))
//
//	// 6
//	// выравнивание
//	// b := []byte{1} // cap 1
//	// b = append(b, 1,2,3,4) // 8
//	// fmt.Println(cap(b))
//	// // 8
//
//	fmt.Println(a[:1]) // a[:1] exclusive => [x:y] x  - inclusive, y - exclusive => will be only first element
//	f(a)
//	//for i := 0; i < 10; i++ {
//	//f(a[:1])
//	//}
//
//	fmt.Println(a)
//	// 1, 1, 2, 3, 4
//	// 1, 10, 2, 3, 4
//}
//
//func f(a []int32) { // get copy of slice, and here will be copied pointer also to the original arrray
//	a = append(a, 10)
//}
//
//// 1 pointer, len, cap
//// 2 pointer, len, cap
