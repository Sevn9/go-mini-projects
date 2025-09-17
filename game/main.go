package main

import (
	"fmt"
)

var player *Player
var rooms map[RoomName]IRoom

var playerAction map[string]ActionDelegate

type ActionDelegate func(*Player, string) string

func main() {
	//initGame()
	//fmt.Println(handleCommand("идти комната"))
	//initGame()
}

func initGame() {
	//заново инициализируем комнаты
	rooms = map[RoomName]IRoom{
		RoomNameKitchen:  NewKitchen(),
		RoomNameMyRoom:   NewMyRoom(),
		RoomNameCorridor: NewCorridor(),
		RoomNameStreet:   NewStreet(),
		RoomNameHome:     NewHome(),
	}

	player = NewPlayer(RoomNameKitchen)
	// создать мапу действий игрока: название - функция
	playerAction = map[string]ActionDelegate{
		"осмотреться": (*Player).LookAround,
		"идти":        (*Player).GoToRoom,
		"надеть":      (*Player).PutOn,
		"взять":       (*Player).Take,
		"применить":   (*Player).Apply,
	}

	room, isRoomExist := rooms[RoomNameMyRoom]
	if isRoomExist {
		fmt.Println("Y test", room.LookAroundInfo())
	}
}

// сюды приходит команда с тестов
func handleCommand(command string) string {

	result := player.commandHandler(command)

	return result
}

// удаляет одно вхождение в slice
func deleteSliceItem[T comparable](items []T, value T) []T {
	for i, v := range items {
		if v == value {
			return append(items[:i], items[i+1:]...)
		}
	}
	return items
}
