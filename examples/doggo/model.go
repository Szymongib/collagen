package main

//go:generate collagen --name Doggo
type Doggo struct {
	Name      string
	BirthDate int64
	Height    float32
}
