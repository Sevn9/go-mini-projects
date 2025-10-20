package main

import (
	"slices"
	"strings"
)

type RoomKitchen struct {
	mapPlaceItemKitchen map[PlaceName][]ItemsName
}

func NewKitchen() IRoom {
	mapPlace := make(map[PlaceName][]ItemsName)

	for place, items := range mapPlaceItemKitchen {
		copied := make([]ItemsName, len(items))
		copy(copied, items)
		mapPlace[place] = copied
	}
	return &RoomKitchen{
		mapPlaceItemKitchen: mapPlace,
	}
}

func (r *RoomKitchen) LookAroundInfo() string {
	answer := "ты находишься на кухне, "

	var parts []string

	for _, placeName := range placesOrderKitchen {

		itemsName, isPlaceItemExist := r.mapPlaceItemKitchen[placeName]

		if !isPlaceItemExist {
			return "не существует такого места" + string(placeName)
		}

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

	answer += strings.Join(parts, ", ") + ", "

	return answer
}

func (r *RoomKitchen) CanExitTo(roomName RoomName) (bool, string) {
	return true, ""
}

func (r *RoomKitchen) TransitionInfo() string {
	answer := "кухня, ничего интересного. "

	answer += GetExits(RoomNameKitchen)

	return answer
}

func (r *RoomKitchen) DeleteItemFromRoom(itemName string) {
	for place, y := range r.mapPlaceItemKitchen {
		r.mapPlaceItemKitchen[place] = deleteSliceItem(y, ItemsName(itemName))
	}
}

func (r *RoomKitchen) IsItemExistThisRoom(itemName string) bool {
	for _, placeName := range r.mapPlaceItemKitchen {
		if slices.Contains(placeName, ItemsName(itemName)) {
			return true
		}
	}
	return false
}

func (r *RoomKitchen) ApplyItem(itemName ItemsName, interactionPlace InteractionPlaceName) string {
	return ""
}
