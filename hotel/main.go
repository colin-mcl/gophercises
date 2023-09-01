package main

import (
	"fmt"
	"math/rand"
)

const roomsTotal int = 134

func main() {
	hotelName := "Colin's Inn"
	fmt.Printf("Hotel: %s\n", hotelName)
	var occupiedRooms = rand.Intn(134)
	var occupancyRate = (float32(occupiedRooms) / float32(roomsTotal)) * 100

	fmt.Printf("Occupancy rate: %0.2f%\n", occupancyRate)
	if occupancyRate <= 30 {
		fmt.Println("Occupancy level: low")
	} else if occupancyRate <= 60 {
		fmt.Println("Occupancy level: medium")
	} else {
		fmt.Println("Occupancy level: high")
	}

	fmt.Println("Rooms: ")
	for i := 0; i < roomsTotal-occupiedRooms; i++ {
		var people = rand.Intn(10) + 1
		var nights = rand.Intn(10) + 1
		fmt.Printf("\t - %d : %d people / %d nights\n", 110+i, people, nights)
	}
}
