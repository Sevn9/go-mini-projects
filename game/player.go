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

	room, isRoomExist := rooms[p.currentRoom]

	if !isRoomExist {
		fmt.Println("LOG PutOn room not exist:" + p.currentRoom)
		return "такой комнаты не существует"
	}

	//существует ли рюкзак в этой комнате?
	isItemExist := room.IsItemExistThisRoom(itemName)

	if !isItemExist {
		return "нет такого"
	}

	//изменить состояние игрока
	p.inventoryItems[ItemsNameBackpack] = struct{}{}

	p.isBackpackOn = true

	//удалить из комнаты предмет т.к. его больше нет
	room.DeleteItemFromFurniture(itemName)

	answer := "вы надели: " + itemName
	return answer
}

func (p *Player) Take(itemName string) string {

	fmt.Println("LOG: Take() Был взят: " + itemName)
	if !p.isBackpackOn {
		return "некуда класть"
	}

	room, isRoomExist := rooms[p.currentRoom]

	if !isRoomExist {
		fmt.Println("LOG PutOn room not exist:" + p.currentRoom)
		return "такой комнаты не существует"
	}

	//существует ли этот предмет в этой комнате?
	isItemExist := room.IsItemExistThisRoom(itemName)

	if !isItemExist {
		return "нет такого"
	}

	p.inventoryItems[ItemsName(itemName)] = struct{}{}

	//удалить из комнаты предмет т.к. его больше нет
	room.DeleteItemFromFurniture(itemName)

	answer := "предмет добавлен в инвентарь: " + itemName

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

func ContainsSliceItem(items []string, value string) bool {
	for _, elem := range items {
		if elem == value {
			return true
		}
	}
	return false
}
