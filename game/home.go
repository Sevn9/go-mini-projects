package main

type Home struct {
}

func NewHome() IRoom {
	return &Home{}
}

func (r *Home) LookAroundInfo() string {
	answer := "ты находишься дома"
	return answer
}

func (r *Home) CanExitTo(roomName RoomName) (bool, string) {
	return true, ""
}

func (r *Home) TransitionInfo() string {
	return ""
}

func (r *Home) DeleteItemFromRoom(itemName string) {
}

func (r *Home) IsItemExistThisRoom(itemName string) bool {
	return false
}

func (r *Home) ApplyItem(itemName ItemsName, interactionPlace InteractionPlaceName) string {
	return ""
}
