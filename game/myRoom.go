package main

import (
	"fmt"
	"slices"
	"strings"
)

type MyRoom struct {
	mapPlaceItemMyRoom map[PlaceName][]ItemsName
}

func NewMyRoom() IRoom {
	mapPlace := make(map[PlaceName][]ItemsName)

	for place, items := range mapPlaceItemMyRoom {
		copied := make([]ItemsName, len(items))
		copy(copied, items)
		mapPlace[place] = copied
	}
	return &MyRoom{mapPlaceItemMyRoom: mapPlace}
}

func (r *MyRoom) LookAroundInfo() string {
	fmt.Println("LOG LookAroundInfo: MyRoom")
	var answer string
	var parts []string

	for _, placeName := range placesOrderMyRoom {

		fmt.Println("LOG placeName: ", placeName)

		itemsName, isPlaceItemExist := r.mapPlaceItemMyRoom[placeName]

		if !isPlaceItemExist {
			return "не существует такого места" + string(placeName)
		}

		fmt.Println("LOG itemsName: ", itemsName)

		if len(itemsName) == 0 {
			continue
		}

		var itemStrs []string
		for _, itemName := range itemsName {
			itemStrs = append(itemStrs, string(itemName))
		}

		part := string(placeName) + ": " + strings.Join(itemStrs, ", ")
		parts = append(parts, part)
	}

	if len(parts) == 0 {
		answer += "пустая комната. "
	} else {
		answer += strings.Join(parts, ", ") + ". "
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
	fmt.Println("LOG DeleteItemFromFurniture: ", itemName)
	for place, y := range r.mapPlaceItemMyRoom {
		if slices.Contains(y, ItemsName(itemName)) {
			r.mapPlaceItemMyRoom[place] = deleteSliceItem(y, ItemsName(itemName))
			fmt.Println("LOG DeleteItemFromFurniture complete: ", r.mapPlaceItemMyRoom[place])
		}
	}
}

func (r *MyRoom) IsItemExistThisRoom(itemName string) bool {
	for _, placeName := range r.mapPlaceItemMyRoom {
		if slices.Contains(placeName, ItemsName(itemName)) {
			return true
		}
	}
	return false
}
