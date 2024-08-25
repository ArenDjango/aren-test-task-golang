package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	//person = &Person{"Alice"} not working
	person.Name = "Changed Name" // working
}

type XYZ struct {
	X, Y, Z int
}

func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)

	var s1 XYZ
	fmt.Println(s1.X, s1.Y, s1.Z)
	p1 := XYZ{23, 12, -2}
	p2 := XYZ{X: 12, Y: 13}
	fmt.Println(p1.X, p1.Y, p1.Z)
	fmt.Println(p2.X, p2.Y, p2.Z)

	pSlice := [4]XYZ{}
	pSlice[2] = p1
	pSlice[0] = p2
	fmt.Println(pSlice)
	p2 = XYZ{1, 2, 3}
	fmt.Println(pSlice)
}

//
