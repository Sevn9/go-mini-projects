package main

import (
	"fmt"
	"maps"
)

type MyRoom struct {
	mapPlaceItemMyRoom map[PlaceName][]ItemsName
}

func NewMyRoom() IRoom {
	mapPlace := maps.Clone(mapPlaceItemMyRoom)
	return &MyRoom{
		mapPlaceItemMyRoom: mapPlace,
	}
}

func (r *MyRoom) LookAroundInfo() string {
	fmt.Println("LOG LookAroundInfo: MyRoom")
	var answer string

	furnitureCounter := len(placesOrderMyRoom)
	fmt.Println("LOG furnitureCounter : ", furnitureCounter)

	for _, placeName := range placesOrderMyRoom {

		itemsName := mapPlaceItemMyRoom[placeName]

		if len(itemsName) == 0 {
			continue
		}
		answer += string(placeName) + ": "

		itemsNameCounter := len(itemsName)
		//пройдем по itemsName
		for i, itemName := range itemsName {
			answer += string(itemName)
			if i < itemsNameCounter-1 {
				answer += ", "
			}
		}

		furnitureCounter--

		if furnitureCounter > 0 {
			answer += ", "
		} else {
			answer += ". "
		}
	}

	//перечислить выходы
	answer += "можно пройти - "
	for placeNameRoom, exitsRoom := range mapTransitionRoom {

		exitsNameCounter := len(exitsRoom)

		if placeNameRoom == RoomNameMyRoom {
			for i, exit := range exitsRoom {
				answer += string(exit)
				if i < exitsNameCounter-1 {
					answer += ", "
				}
			}
		}
	}
	return answer
}

func (r *MyRoom) TransitionInfo() string {
	fmt.Println("LOG TransitionInfo: MyRoom")
	answer := "ты в своей комнате. "

	answer += "можно пройти - "
	for placeNameRoom, exitsRoom := range mapTransitionRoom {

		exitsNameCounter := len(exitsRoom)

		if placeNameRoom == RoomNameMyRoom {
			for i, exit := range exitsRoom {
				answer += string(exit)
				if i < exitsNameCounter-1 {
					answer += ", "
				}
			}
		}
	}

	return answer
}

func (r *MyRoom) DeleteItemFromFurniture(itemName string) {
	for _, y := range r.mapPlaceItemMyRoom {
		deleteSliceItem(y, ItemsName(itemName))
	}
}
