package main

import (
	"fmt"
	"strings"
)

type IRoom interface {
	LookAroundInfo() string
	TransitionInfo() string
	DeleteItemFromFurniture(string)
	DeleteItemFromRoom(string)
	AccessExitsAnswer() string
}

type Furniture struct {
	Name  string
	Items []string
}

type Room struct {
	RoomName              string
	AvailableItems        map[string]*Item
	AvailableRoomsExit    []*Room
	LookAroundDescription map[string]string
	TransitionDescription string
	FurnitureItem         []*Furniture
}

func NewRoom(roomName string,
	availableItems map[string]*Item,
	lookAroundDescription map[string]string,
	transitionDescription string,
	furnitureItem []*Furniture) *Room {
	newRoom := Room{
		RoomName:              roomName,
		AvailableItems:        availableItems,
		LookAroundDescription: lookAroundDescription,
		TransitionDescription: transitionDescription,
		FurnitureItem:         furnitureItem,
	}
	return &newRoom
}

func (r *Room) LookAroundInfo() string {
	var resultAnswer string
	var instruction string
	//instruction := "пустая комната"

	//проверить есть ли начальное описание и инструкция для игрока
	if len(r.LookAroundDescription) != 0 {
		resultAnswer = r.LookAroundDescription["roomDescription"] + ", "
		instruction = r.LookAroundDescription["instruction"]
	} else {
		instruction = "."
	}

	furnitureLength := len(r.FurnitureItem)
	//есть ли мебель и если есть то перечислить предметы
	if furnitureLength != 0 {
		for _, item := range r.FurnitureItem {
			resultAnswer += item.Name + ": " + strings.Join(item.Items, ", ")
			furnitureLength -= 1
			if furnitureLength != 0 {
				resultAnswer += ", "
			}
		}
	} else {
		resultAnswer += "пустая комната"
	}

	//вынести логику в отдельный метод
	exitAnswer := r.AccessExitsAnswer()

	if instruction != "." {
		resultAnswer += ", " + instruction
		fmt.Println("LOG:", instruction)
	}

	resultAnswer += "." + exitAnswer

	return resultAnswer
}

func (r *Room) TransitionInfo() string {
	fmt.Println("LOG:" + r.TransitionDescription)
	desc := r.TransitionDescription

	exitAnswer := r.AccessExitsAnswer()

	fmt.Println("LOG:" + desc + exitAnswer)
	return desc + exitAnswer
}

func (r *Room) DeleteItemFromRoom(itemName string) {

	for _, v := range r.FurnitureItem {

		fmt.Println("LOG DeleteItemFromRoom before: " + v.Name + v.Items[0])

	}
	//удаляем из доступных
	delete(r.AvailableItems, itemName)

	//удаляем из комнаты
	for i, v := range r.FurnitureItem {
		for y, item := range v.Items {
			if item == itemName {
				fmt.Println("LOG DeleteItemFromRoom: " + itemName)
				v.Items = append(v.Items[:y], v.Items[y+1:]...)

				if len(v.Items) == 0 {
					r.FurnitureItem = append(r.FurnitureItem[:i], r.FurnitureItem[i+1:]...)
				}
				break
			}
		}
	}

	for _, v := range r.FurnitureItem {

		fmt.Println("LOG DeleteItemFromRoom foreach: " + v.Name)

	}
}

func (r *Room) AccessExitsAnswer() string {
	//логика перечисления выходов
	var exitsAnswerSlice []string

	for _, item := range r.AvailableRoomsExit {
		exitsAnswerSlice = append(exitsAnswerSlice, item.RoomName)
	}

	exitAnswer := fmt.Sprintf(" можно пройти - %s", strings.Join(exitsAnswerSlice, ", "))
	return exitAnswer
}
