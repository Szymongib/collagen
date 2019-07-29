package main

import (
	"fmt"
	"time"
)

func main() {

	doggos := Doggos{
		{Name: "Fluffy", BirthDate: time.Now().Unix(), Height: 120},
		{Name: "Max", BirthDate: time.Now().Unix(), Height: 20},
	}

	fmt.Println("Doggos:")
	for _, d := range doggos {
		fmt.Println(d)
	}

	fmt.Println("Long doggos:")
	for _, d := range doggos.Filter(func(item Doggo) bool {
		return item.Height > 50
	}) {
		fmt.Println(d)
	}

	fmt.Println("Max exist: ", doggos.Exist(func(element Doggo) bool {
		return element.Name == "Max"
	}))

	slicedDoggos := doggos.ToSlice()

	fmt.Print(slicedDoggos)

	fmt.Println(doggos.Map(func(item Doggo) interface{} {
		return item.Name
	}))
}

//func LargeDoggos(doggos Doggos) Doggos {
//	return doggos.Filter(func(d Doggo) bool {
//		return d.Height >= 50
//	})
//}

func LargeDoggos(doggos Doggos) Doggos {
	return doggos.Filter(func(d Doggo) bool {
		return d.Height >= 50
	})
}

// TODO - several commands like doggos exist? doggos older than etc.

//func loadDoggos() Doggos {
//
//}
