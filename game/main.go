package main

import (
	"slices"
)

var player IPlayer
var rooms map[RoomName]IRoom
var conditionForReadyGoingToSchool []ItemsName

var playerAction map[string]ActionDelegate

type ActionDelegate func(IPlayer, string) string

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
		"осмотреться": IPlayer.LookAround,
		"идти":        IPlayer.GoToRoom,
		"надеть":      IPlayer.PutOn,
		"взять":       IPlayer.Take,
		"применить":   IPlayer.Apply,
	}

	conditionForReadyGoingToSchool = []ItemsName{ItemsNameBackpack, ItemsNameKeys, ItemsNameNotes}
	slices.Sort(conditionForReadyGoingToSchool)
}

// сюды приходит команда с тестов
func handleCommand(command string) string {

	result := player.CommandHandler(command)

	return result
}
