package main

import (
	"fmt"
	"time"
)

func main() {

	doggos := Doggos{
		{
			Name:      "Fluffy",
			BirthDate: time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
			Height:    25,
		},
		{
			Name:      "Occy",
			BirthDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
			Height:    102,
		},
		{
			Name:      "Boo",
			BirthDate: time.Date(2013, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
			Height:    49,
		},
		{
			Name:      "Poo",
			BirthDate: time.Date(2009, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
			Height:    17,
		},
		{
			Name:      "Rex",
			BirthDate: time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
			Height:    86,
		},
	}

	fmt.Println("Doggos:")
	for _, d := range doggos {
		fmt.Println(d)
	}

	fmt.Println("Doggos higher than 50:")
	for _, d := range doggos.Filter(func(item Doggo) bool {
		return item.Height > 50
	}) {
		fmt.Println(d)
	}

	fmt.Println("Is there Doggo named Rex?")
	fmt.Println(doggos.Exist(func(element Doggo) bool {
		return element.Name == "Rex"
	}))

	fmt.Println("Doggos names: ")
	fmt.Println(doggos.Map(func(item Doggo) interface{} {
		return item.Name
	}))
}
