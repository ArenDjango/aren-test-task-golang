package main

//func main() {
//	//anArray := [4]int{1, 2, 4, -4}
//	twoD := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
//	//threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}
//	//
//	//fmt.Println("The length of", anArray, "is", len(anArray))
//	//fmt.Println("The first element of", twoD, "is", twoD[0][0])
//	//fmt.Println("The length of", threeD, "is", len(threeD))
//	//
//	//for i := 0; i < len(threeD); i++ {
//	//	v := threeD[i]
//	//	for j := 0; j < len(v); j++ {
//	//		m := v[j]
//	//		for k := 0; k < len(m); k++ {
//	//			fmt.Print(m[k], " ")
//	//		}
//	//	}
//	//}
//	//
//	//for _, v := range threeD {
//	//	for _, m := range v {
//	//		for _, s := range m {
//	//			fmt.Print(s, " ")
//	//		}
//	//	}
//	//	fmt.Println()
//	//}
//	//
//	////s := make([]byte, 5)
//	//
//	//a6 := []int{-10, 1, 2, 3, 4, 5}
//	//a4 := []int{-1, -2, -3, -4}
//	//fmt.Println("a6:", a6)
//	//fmt.Println("a4:", a4)
//	//
//	//copy(a6, a4)
//	//fmt.Println("a6:", a6)
//	//fmt.Println("a4:", a4)
//	//fmt.Println()
//	//
//	//b6 := []int{-10, 1, 2, 3, 4, 5}
//	//b4 := []int{-1, -2, -3, -4}
//	//fmt.Println("b6:", b6)
//	//fmt.Println("b4:", b4)
//	//copy(b4, b6)
//	//fmt.Println("b6:", b6)
//	//fmt.Println("b4:", b4)
//	//
//	//s1 := make([][]int, 4)
//
//	anArray := [5]int{-1, -2, -3, -4, -5}
//	refAnArray := anArray[:] // slice with pointer to existing array
//	fmt.Println(anArray)
//	fmt.Println(refAnArray)
//	anArray[4] = -100
//	fmt.Println(refAnArray)
//
//	for i := 0; i < len(twoD); i++ {
//		for j := 0; j < len(twoD[i]); j++ {
//			twoD[i] = append(twoD[i], i*j)
//		}
//	}
//
//}

//type aStructure struct {
//	person string
//	height int
//	weight int
//}
//
//func main() {
//	mySlice := make([]aStructure, 0)
//	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
//	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
//	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
//	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
//	mySlice = append(mySlice, aStructure{"Athina", 134, 40})
//
//	fmt.Println("0:", mySlice)
//
//	sort.Slice(mySlice, func(i, j int) bool {
//		return mySlice[i].height < mySlice[j].height
//	})
//	fmt.Println("increase", mySlice)
//	sort.Slice(mySlice, func(i, j int) bool {
//		return mySlice[i].height > mySlice[j].height
//	})
//	fmt.Println("decrease", mySlice)
//
//	sl := []int{1, 2, 3, 4, 5}
//	ar := [2]int{1, 2}
//	sl = append(sl, ar[:]...)
//
//	iMap := map[string]int{
//		"mihalis":       18,
//		"bill":          13,
//		"jkaooploamspl": 15,
//	}
//
//	delete(iMap, "jkaooploamspl")
//
//	for key, value := range iMap {
//		fmt.Println("key:", key, "value:", value) // every time will be random in order
//	}
//
//	_, ok := iMap["doesItExist"]
//	if ok {
//		fmt.Println("exists")
//	} else {
//		fmt.Println("not exists")
//	}
//
//	aMap := map[string]int{}
//	aMap["test"] = 1
//	aMap = nil
//	fmt.Println(aMap)
//	aMap["test"] = 1
//}
