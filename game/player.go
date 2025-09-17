package main

import (
	"fmt"
	"strings"
)

type IPlayer interface {
	GoToRoom(roomName string) string
	LookAround(instruction string) string
	PutOn(itemName string) string
	Take(itemName string) string
	Apply(instruction string) string
}

type Player struct {
	currentRoom    RoomName
	isBackpackOn   bool
	inventoryItems map[ItemsName]struct{}
}

func NewPlayer(room RoomName) *Player {
	newPlayer := Player{
		currentRoom:    room,
		isBackpackOn:   false,
		inventoryItems: make(map[ItemsName]struct{}),
	}

	return &newPlayer
}

func (p *Player) commandHandler(command string) string {

	var answer string
	commandItems := strings.Split(command, " ")

	targetArray := commandItems[1:]

	//если 3 и более параметра то суммируем последние 2 и более в одну строку с разделителем
	// и передаем в нужную функцию чтобы соответствовать сигнатуре
	target := strings.Join(targetArray, ";")

	action, exist := playerAction[commandItems[0]]

	if exist {
		answer = action(p, target)
	} else {
		return "неизвестная команда"
	}

	return answer
}

func (p *Player) GoToRoom(roomName string) string {

	fmt.Println("LOG GoToRoom player:" + roomName)

	var answer string

	//проверить можно ли попасть в эту комнату из текущей
	answerChangeRoom, isExist := mapTransitionRoom[p.currentRoom]

	if !isExist {
		return "нет пути в " + roomName
	}

	for _, roomNameChange := range answerChangeRoom {
		if roomNameChange == RoomName(roomName) {

			fmt.Println("LOG roomName:" + roomName)

			p.currentRoom = RoomName(roomName)
			fmt.Println("LOG p.currentRoom:" + p.currentRoom)

			room, isRoomExist := rooms[p.currentRoom]

			if !isRoomExist {
				fmt.Println("LOG GoToRoom room not exist:" + answer)
				return "такой комнаты не существует"
			}

			answer += room.TransitionInfo()

			fmt.Println("LOG GoToRoom answer:" + answer)
			return answer
		}
	}

	//изменить состояние игрока
	return "нет пути в " + roomName
}

func (p *Player) LookAround(instruction string) string {
	room, isExist := rooms[p.currentRoom]

	if !isExist {
		return "нет такой комнаты"
	}

	answer := room.LookAroundInfo()

	return answer
}

func (p *Player) PutOn(itemName string) string {

	if itemName != string(ItemsNameBackpack) {
		return "предмет нельзя надеть"
	}

	//существует ли рюкзак в этой комнате? надо достать mapPlaceItemMyRoom

	itemPlacesCurrentRoom, isPlaceExist := mapRoomPlace[p.currentRoom]

	if !isPlaceExist {
		return "в этой комнате нет предметов"
	}

	//изменить состояние игрока
	p.inventoryItems[ItemsNameBackpack] = struct{}{}

	p.isBackpackOn = true

	//удалить из комнаты предмет т.к. его больше нет
	for place, itemsName := range itemPlacesCurrentRoom {
		for i, item := range itemsName {
			if item == ItemsName(itemName) {
				itemPlacesCurrentRoom[place] = append(itemsName[:i], itemsName[i+1:]...)
				break
			}
		}
	}

	answer := "вы надели: " + itemName
	return answer
}

func (p *Player) Take(itemName string) string {

	fmt.Println("LOG: Take() Был взят: " + itemName)
	if !p.isBackpackOn {
		return "некуда класть"
	}

	//проверить наличие любых предметов в этой комнате
	itemPlacesCurrentRoom, isPlaceExist := mapRoomPlace[p.currentRoom]

	if !isPlaceExist {
		return "в этой комнате нет предметов"
	}

	for placeName, items := range itemPlacesCurrentRoom {
		fmt.Println("LOG: Take() NowInventory: ")
		for i, roomItem := range items {
			if roomItem == ItemsName(itemName) {
				//положить в инвентарь
				p.inventoryItems[ItemsName(itemName)] = struct{}{}
				//удалить из комнаты
				itemPlacesCurrentRoom[placeName] = append(items[:i], items[i+1:]...)
				break
			}
		}
	}

	_, isItemBackpackExist := p.inventoryItems[ItemsName(itemName)]

	var answer string

	if isItemBackpackExist {
		answer = "предмет добавлен в инвентарь: " + itemName
	} else {
		answer = "такого предмета нет в комнате"
	}

	return answer
}

func (p *Player) Apply(instruction string) string {

	instructionsSlice := strings.Split(instruction, ";")

	if len(instructionsSlice) != 2 {
		return "неправильные входные данные"
	}

	itemName := instructionsSlice[0]
	interactionsPlaceName := instructionsSlice[1]

	//проверить предмет в инвентаре
	_, exist := p.inventoryItems[ItemsName(itemName)]

	fmt.Println("LOG: Apply() ")

	if !exist {
		return "нет предмета в инвентаре - " + itemName
	}

	//проверить можно ли применить к предмету
	interactionMap, isItemRuleExist := itemsRules[ItemsName(itemName)]

	if !isItemRuleExist {
		return "не к чему применить"
	}

	resultAction := interactionMap[InteractionPlaceName(interactionsPlaceName)]

	return resultAction
}

// удаляет одно вхождение в slice
func deleteSliceItemOld(items []string, value string) []string {
	for i, v := range items {
		if v == value {
			return append(items[:i], items[i+1:]...)
		}
	}

	return items
}

func ContainsSliceItem(items []string, value string) bool {
	for _, elem := range items {
		if elem == value {
			return true
		}
	}
	return false
}
