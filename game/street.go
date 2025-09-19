package main

type Street struct {
}

func NewStreet() IRoom {
	return &Street{}
}

func (s *Street) LookAroundInfo() string {
	answer := "ты находишься на улице"
	return answer
}

func (s *Street) CanExitTo(roomName RoomName) (bool, string) {
	return true, ""
}

func (s *Street) TransitionInfo() string {
	answer := "на улице весна. "
	answer += GetExits(RoomNameStreet)

	return answer
}

func (s *Street) DeleteItemFromRoom(itemName string) {
}

func (r *Street) IsItemExistThisRoom(itemName string) bool {
	return false
}

func (r *Street) ApplyItem(itemName ItemsName, interactionPlace InteractionPlaceName) string {
	return ""
}
