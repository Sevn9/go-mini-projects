package main

import (
	"slices"
	"strings"
)

type IPlayer interface {
	GoToRoom(roomName string) string
	LookAround(instruction string) string
	PutOn(itemName string) string
	Take(itemName string) string
	Apply(instruction string) string
	CommandHandler(command string) string
}

type Player struct {
	currentRoom    RoomName
	isBackpackOn   bool
	inventoryItems map[ItemsName]struct{}
}

func NewPlayer(room RoomName) IPlayer {
	newPlayer := Player{
		currentRoom:    room,
		isBackpackOn:   false,
		inventoryItems: make(map[ItemsName]struct{}),
	}

	return &newPlayer
}

func (p *Player) CommandHandler(command string) string {

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

	var answer string
	exitRoom := RoomName(roomName)

	//существует ли эта комната
	room, isRoomExist := rooms[exitRoom]

	if !isRoomExist {
		return "такой комнаты не существует: " + roomName
	}

	//можно ли туда пройти
	roomNames, _ := mapTransitionRoom[p.currentRoom]

	if !slices.Contains(roomNames, exitRoom) {
		return "нет пути в " + roomName
	}

	//проверка на условия выхода из комнаты
	playerCurrentRoom, _ := rooms[p.currentRoom]
	for _, elem := range roomNames {
		if elem == exitRoom {
			isCan, answerExit := playerCurrentRoom.CanExitTo(exitRoom)

			if !isCan {
				answer += answerExit
				//если вернется неудовлетворительный результат то нельзя менять комнату
				break
			}

			answer += room.TransitionInfo()
			p.currentRoom = exitRoom

			break
		}
	}
	return answer
}

func (p *Player) LookAround(instruction string) string {
	room, isExist := rooms[p.currentRoom]

	if !isExist {
		return "нет такой комнаты"
	}

	answer := room.LookAroundInfo()

	//проверить что собран рюкзак
	if p.currentRoom == RoomNameKitchen {
		var keys []ItemsName
		for item := range p.inventoryItems {
			keys = append(keys, item)
		}

		slices.Sort(keys)

		if slices.Equal(keys, conditionForReadyGoingToSchool) {
			answer += "надо идти в универ. "
		} else {
			answer += "надо собрать рюкзак и идти в универ. "
		}
	}

	answer += GetExits(p.currentRoom)

	return answer
}

func (p *Player) PutOn(itemName string) string {

	if itemName != string(ItemsNameBackpack) {
		return "предмет нельзя надеть"
	}

	room, isRoomExist := rooms[p.currentRoom]

	if !isRoomExist {
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
	room.DeleteItemFromRoom(itemName)

	answer := "вы надели: " + itemName
	return answer
}

func (p *Player) Take(itemName string) string {

	if !p.isBackpackOn {
		return "некуда класть"
	}

	room, isRoomExist := rooms[p.currentRoom]

	if !isRoomExist {
		return "такой комнаты не существует"
	}

	//существует ли этот предмет в этой комнате?
	isItemExist := room.IsItemExistThisRoom(itemName)

	if !isItemExist {
		return "нет такого"
	}

	p.inventoryItems[ItemsName(itemName)] = struct{}{}

	//удалить из комнаты предмет т.к. его больше нет
	room.DeleteItemFromRoom(itemName)

	answer := "предмет добавлен в инвентарь: " + itemName

	return answer
}

func (p *Player) Apply(instruction string) string {

	instructionsSlice := strings.Split(instruction, ";")

	if len(instructionsSlice) != 2 {
		return "неправильные входные данные"
	}

	itemName := ItemsName(instructionsSlice[0])
	interactionsPlaceName := InteractionPlaceName(instructionsSlice[1])

	//проверить предмет в инвентаре
	_, exist := p.inventoryItems[itemName]

	if !exist {
		return "нет предмета в инвентаре - " + string(itemName)
	}

	room, _ := rooms[p.currentRoom]

	resultAction := room.ApplyItem(itemName, interactionsPlaceName)

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
