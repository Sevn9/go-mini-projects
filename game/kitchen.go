package main

import (
	"slices"
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

	//todo: исправить под новую структуру т.к. map не выдает объекты в фиксированном порядке
	for placeName, itemsName := range r.mapPlaceItemKitchen {

		if len(itemsName) == 0 {
			continue
		}

		answer += string(placeName) + ": "

		for _, itemName := range itemsName {
			answer += string(itemName) + ", "
		}
	}

	//если ничего не собрано:
	answer += "надо собрать рюкзак и идти в универ. "

	//все места где можно выйти
	answer += "можно пройти - "
	for placeNameRoom, exitsRoom := range mapTransitionRoom {
		if placeNameRoom == RoomNameKitchen {
			for _, exit := range exitsRoom {
				answer += string(exit)
			}
		}
	}

	return answer
}

func (r *RoomKitchen) TransitionInfo() string {
	answer := "кухня, ничего интересного. "

	answer += "можно пройти - "
	for placeNameRoom, exitsRoom := range mapTransitionRoom {
		if placeNameRoom == RoomNameKitchen {
			for _, exit := range exitsRoom {
				answer += string(exit)
			}
		}
	}

	return answer
}

func (r *RoomKitchen) DeleteItemFromFurniture(itemName string) {
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
