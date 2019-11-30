package main

import (
	"fmt"
	"time"
)

//go:generate collagen --name Doggo --plural DoggoPointers --pointer
//go:generate collagen --name Doggo
type Doggo struct {
	Name      string
	BirthDate int64
	Height    float32
}

func (d Doggo) String() string {
	return fmt.Sprintf("Name: %s, BirthDate: %s, Height: %v",
		d.Name, time.Unix(d.BirthDate, 0).String(), d.Height)
}
