package main

type RoomCorridor struct {
}

func NewCorridor() IRoom {
	return &RoomCorridor{}
}

func (r *RoomCorridor) LookAroundInfo() string {
	answer := "ты находишься в коридоре,"

	//все места где можно выйти
	//беру из mapChangeRoom и добавляю в answer
	return answer
}

func (r *RoomCorridor) TransitionInfo() string {
	answer := "ничего интересного. "

	answer += "можно пройти - "
	for placeNameRoom, exitsRoom := range mapTransitionRoom {

		exitsNameCounter := len(exitsRoom)

		if placeNameRoom == RoomNameCorridor {
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

func (r *RoomCorridor) DeleteItemFromFurniture(itemName string) {
}

func (r *RoomCorridor) IsItemExistThisRoom(itemName string) bool {
	return false
}
