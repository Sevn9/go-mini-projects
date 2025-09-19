package main

import (
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
	var answer string
	var parts []string

	for _, placeName := range placesOrderMyRoom {

		itemsName, _ := r.mapPlaceItemMyRoom[placeName]

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

	return answer
}

func (r *MyRoom) CanExitTo(roomName RoomName) (bool, string) {
	return true, ""
}

func (r *MyRoom) TransitionInfo() string {

	answer := "ты в своей комнате. "
	answer += GetExits(RoomNameMyRoom)

	return answer
}

func (r *MyRoom) DeleteItemFromRoom(itemName string) {
	for place, y := range r.mapPlaceItemMyRoom {
		if slices.Contains(y, ItemsName(itemName)) {
			r.mapPlaceItemMyRoom[place] = deleteSliceItem(y, ItemsName(itemName))
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

func (r *MyRoom) ApplyItem(itemName ItemsName, interactionPlace InteractionPlaceName) string {
	return ""
}
