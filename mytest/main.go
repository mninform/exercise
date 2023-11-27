package main

import "fmt"

const maß float64 = 1000 //ml

var durst float64 = 0.3 //ml pro sekunde

func main() {
	zeit := 0.0
	zeit = maß / durst
	fmt.Println("Ich trinke die Maß in ", zeit, "sekunden")
}
