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
	currentRoom    *Room
	isBackpackOn   bool
	inventoryItems map[string]*Item
}

func NewPlayer(room *Room) *Player {
	newPlayer := Player{
		currentRoom:    room,
		isBackpackOn:   false,
		inventoryItems: map[string]*Item{},
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

	//проверить можно ли попасть в эту комнату из текущей
	isCanToGo := false

	for _, room := range p.currentRoom.AvailableRoomsExit {
		if room.RoomName == roomName {
			isCanToGo = true
		}
	}

	if !isCanToGo {
		return "нет пути в " + roomName
	}

	//изменить состояние игрока
	p.currentRoom = rooms[roomName]
	fmt.Println("LOG:" + p.currentRoom.RoomName)

	answer := p.currentRoom.TransitionInfo()
	return answer
}

func (p *Player) LookAround(instruction string) string {
	answer := p.currentRoom.LookAroundInfo()
	return answer
}

func (p *Player) PutOn(itemName string) string {

	if itemName == "рюкзак" {
		//изменить состояние игрока
		p.isBackpackOn = true

		//удалить из комнаты предмет т.к. его больше нет
		p.currentRoom.DeleteItemFromRoom(itemName)

		answer := "вы надели: " + itemName
		return answer
	}

	return "нельзя надеть этот предмет"
}

func (p *Player) Take(itemName string) string {

	fmt.Println("LOG: Take() Был взят: " + itemName)
	if !p.isBackpackOn {
		return "некуда класть"
	}

	//проверить наличие предмета в этой комнате
	item, exist := p.currentRoom.AvailableItems[itemName]

	if !exist {
		return "нет такого"
	}

	p.inventoryItems[itemName] = item

	for _, item := range p.inventoryItems {
		fmt.Println("LOG: Take() NowInventory: " + item.ItemName)
	}
	answer := "предмет добавлен в инвентарь: " + item.ItemName

	//удалить из комнаты предмет т.к. его больше нет
	p.currentRoom.DeleteItemFromRoom(itemName)

	return answer
}

func (p *Player) Apply(instruction string) string {

	instructionsSlice := strings.Split(instruction, ";")

	if len(instructionsSlice) != 2 {
		return "неправильные входные данные"
	}

	itemName := instructionsSlice[0]
	interactions := instructionsSlice[1]

	item, exist := p.inventoryItems[itemName]

	fmt.Println("LOG: Apply() ")

	if !exist {
		return "нет предмета в инвентаре - " + itemName
	}

	useAnswer := item.UseItem(interactions)

	return useAnswer
}

// удаляет одно вхождение в slice
func deleteSliceItem(items []string, value string) []string {
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
