package main

import "fmt"

type DistanceUnit int

const (
	Kilometer DistanceUnit = 0
	Mile      DistanceUnit = 1
)

func (sc DistanceUnit) String() string {
	units := []string{"km", "mi"}
	return units[sc]
}

func main() {

	kmUnit := Kilometer
	fmt.Println(kmUnit.String())

	miUnit := Mile
	fmt.Println(miUnit.String())

}
