package main

import (
	"fmt"
)

var player *Player
var rooms map[RoomName]IRoom

var playerAction map[string]ActionDelegate

type ActionDelegate func(*Player, string) string

func main() {
	initGame()
	fmt.Println(handleCommand("идти комната"))
}

func initGame() {
	//rooms = make(map[string]*Room)

	//заново инициализируем комнаты
	rooms = map[RoomName]IRoom{
		RoomNameKitchen:  NewKitchen(),
		RoomNameMyRoom:   NewMyRoom(),
		RoomNameCorridor: NewCorridor(),
		RoomNameStreet:   NewStreet(),
		RoomNameHome:     NewHome(),
	}

	//Создать предметы и связи
	//keys := NewItem("ключи")
	//keys.AddUsage("дверь", "дверь открыта")
	//notes := NewItem("конспекты")
	//tea := NewItem("чай")
	//backpack := NewItem("рюкзак")

	//todo: разделить по запятой lookAroundDescription чтобы вставить предметы на столе и стуле?

	//todo: сделать instruction составной ибо
	// result:   ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор
	// expected: ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор

	//создать все комнаты

	/*
		kitchenRoom := NewRoom(
			"кухня",
			map[string]*Item{
				"чай": tea,
			},
			map[string]string{
				"roomDescription": "ты находишься на кухне",
				"instruction":     "надо собрать рюкзак и идти в универ",
			},
			"кухня, ничего интересного.",
			[]*Furniture{
				{Name: "на столе", Items: []string{"чай"}},
			})

		rooms[kitchenRoom.RoomName] = kitchenRoom

		corridor := NewRoom(
			"коридор",
			map[string]*Item{},
			map[string]string{},
			"ничего интересного.",
			[]*Furniture{})

		rooms[corridor.RoomName] = corridor

		myRoom := NewRoom(
			"комната",
			map[string]*Item{
				"ключи":     keys,
				"конспекты": notes,
				"рюкзак":    backpack,
			},
			map[string]string{},
			"ты в своей комнате.",
			[]*Furniture{
				{Name: "на столе", Items: []string{"ключи", "конспекты"}},
				{Name: "на стуле", Items: []string{"рюкзак"}},
			})

		rooms[myRoom.RoomName] = myRoom

		street := NewRoom(
			"улица",
			map[string]*Item{},
			map[string]string{},
			"на улице весна.",
			[]*Furniture{})

		home := NewRoom(
			"домой",
			map[string]*Item{},
			map[string]string{},
			"вы вернулись домой",
			[]*Furniture{})

		rooms[street.RoomName] = street

		//добавить связи между комнатами
		kitchenRoom.AvailableRoomsExit = append(kitchenRoom.AvailableRoomsExit, corridor)
		corridor.AvailableRoomsExit = append(corridor.AvailableRoomsExit, kitchenRoom, myRoom, street)
		myRoom.AvailableRoomsExit = append(myRoom.AvailableRoomsExit, corridor)
		street.AvailableRoomsExit = append(street.AvailableRoomsExit, home)

		//создать игрока, передать текущую локацию
		player = NewPlayer(kitchenRoom)
	*/

	player = NewPlayer(RoomNameKitchen)
	// создать мапу действий игрока: название - функция
	playerAction = map[string]ActionDelegate{
		"осмотреться": (*Player).LookAround,
		"идти":        (*Player).GoToRoom,
		"надеть":      (*Player).PutOn,
		"взять":       (*Player).Take,
		"применить":   (*Player).Apply,
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
