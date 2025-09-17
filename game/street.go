package main

import "fmt"

type Street struct{}

func NewStreet() IRoom {
	return &Street{}
}

func (s *Street) LookAroundInfo() string {
	answer := "ты находишься на улице"
	return answer
}

func (s *Street) TransitionInfo() string {
	fmt.Println("LOG TransitionInfo: street")
	answer := "на улице весна. "

	answer += "можно пройти - "
	for placeNameRoom, exitsRoom := range mapTransitionRoom {

		exitsNameCounter := len(exitsRoom)

		if placeNameRoom == RoomNameStreet {
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

func (s *Street) DeleteItemFromFurniture(itemName string) {
}

func (r *Street) IsItemExistThisRoom(itemName string) bool {
	return false
}
