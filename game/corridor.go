package main

type RoomCorridor struct {
	isDoorOpen bool
}

func NewCorridor() IRoom {
	return &RoomCorridor{
		isDoorOpen: false,
	}
}

func (r *RoomCorridor) LookAroundInfo() string {
	answer := "ты находишься в коридоре, дверь перекрывает выход на улицу"
	return answer
}

func (r *RoomCorridor) CanExitTo(roomName RoomName) (bool, string) {
	if roomName == RoomNameStreet && !r.isDoorOpen {
		return false, "дверь закрыта"
	}
	return true, ""
}

func (r *RoomCorridor) TransitionInfo() string {

	answer := "ничего интересного. "
	answer += GetExits(RoomNameCorridor)

	return answer
}

func (r *RoomCorridor) DeleteItemFromRoom(itemName string) {
}

func (r *RoomCorridor) IsItemExistThisRoom(itemName string) bool {
	return false
}

func (r *RoomCorridor) ApplyItem(itemName ItemsName, interactionPlace InteractionPlaceName) string {

	if interactionPlace != interactionPlaceNameDoor {
		return "не к чему применить"
	}

	if itemName != ItemsNameKeys {
		return "нельзя применить этот предмет к " + string(interactionPlace)
	}

	r.isDoorOpen = true
	return "дверь открыта"
}
