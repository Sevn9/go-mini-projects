package main

import (
	"maps"
)

type RoomKitchen struct {
	mapPlaceItemKitchen map[PlaceName][]ItemsName
}

func NewKitchen() IRoom {
	mapPlace := maps.Clone(mapPlaceItemKitchen)
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
	return ""
}

func (r *RoomKitchen) DeleteItemFromFurniture(itemName string) {
	for _, y := range r.mapPlaceItemKitchen {
		deleteSliceItem(y, ItemsName(itemName))
	}
}
