package main

import (
	"bufio"
	"fmt"
	"os"
)

// House represents the 
type House struct {
	numberOfRooms int
	city          string
	address       string
	price         int
}

//  allow users to enter information for one to many houses
func getHouseData() []House {
	var house House
	var address, city string
	var numberOfRooms, price int
	
	// shoutout to Jerome for suggesting check implementation
	var check int
	check = 0
	reader := bufio.NewReader(os.Stdin)

	// list of dictionaries to store houses info
	listOfHouses := []House{}

	for check != 1 {

		fmt.Println("Enter address of house: ")
		address, _ = reader.ReadString('\n')

		house.address = address

		fmt.Print("Enter city of house: ")
		city, _ = reader.ReadString('\n')
		house.city = city

		fmt.Print("Enter number of rooms in the house: ")
		fmt.Scanln(&numberOfRooms)
		house.numberOfRooms = numberOfRooms

		fmt.Print("Enter the price of house: ")
		fmt.Scanln(&price)
		house.price = price

		listOfHouses = append(listOfHouses, house)

		fmt.Println("Press 1 to exit or 0 to continue: ")
		fmt.Scanln(&check)
	}

	return listOfHouses

}

func showAllHouses(houses []House) {
	for i := 0; i < len(houses); i++ {
		fmt.Printf("%-40s %-20s %2d Rooms %d \n", houses[i].address, houses[i].city, houses[i].numberOfRooms, houses[i].price)
	}
}

func main() {
	// TODO: review formatting -- idk why it's a mess rn
	showAllHouses(getHouseData())
}
