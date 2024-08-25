package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

type user struct {
	name  string
	email string
}

func stayOnStack() user {
	u := user{
		name:  "Bill",
		email: "bill@email.com",
	}
	return u
}

func escapeToHeap() *user {
	u2 := user{
		name:  "Bill",
		email: "bill@email.com",
	}
	return &u2
}

func main() {
	f, err := os.Create("mem.pprof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()

	u1 := stayOnStack()
	u2 := escapeToHeap()

	runtime.GC() // Вызовите сборку мусора для получения более точных данных.
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}

	fmt.Println(u1, u2)
}
