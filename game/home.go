package main

type Home struct {
}

func NewHome() IRoom {
	return &Home{}
}

func (r *Home) LookAroundInfo() string {
	answer := "ты находишься в коридоре,"

	//все места где можно выйти
	//беру из mapChangeRoom и добавляю в answer
	return answer
}

func (r *Home) TransitionInfo() string {
	return ""
}

func (r *Home) DeleteItemFromFurniture(itemName string) {
}

func (r *Home) IsItemExistThisRoom(itemName string) bool {
	return false
}
