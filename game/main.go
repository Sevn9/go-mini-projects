package main

import (
	"slices"
)

var player *Player
var rooms map[RoomName]IRoom
var conditionForReadyGoingToSchool []ItemsName

var playerAction map[string]ActionDelegate

type ActionDelegate func(*Player, string) string

func main() {
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

	conditionForReadyGoingToSchool = []ItemsName{ItemsNameBackpack, ItemsNameKeys, ItemsNameNotes}
	slices.Sort(conditionForReadyGoingToSchool)
}

// сюды приходит команда с тестов
func handleCommand(command string) string {

	result := player.commandHandler(command)

	return result
}
