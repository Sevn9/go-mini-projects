package main

type (
	RoomName             string
	ActionName           string
	ItemsName            string
	PlaceName            string
	InteractionPlaceName string
)

const (
	RoomNameKitchen  RoomName = "кухня"
	RoomNameMyRoom   RoomName = "комната"
	RoomNameCorridor RoomName = "коридор"
	RoomNameStreet   RoomName = "улица"
	RoomNameHome     RoomName = "домой"
)

const (
	ActionNameLookAround ActionName = "осмотреться"
	ActionNameGoToRoom   ActionName = "идти"
	ActionNamePutOn      ActionName = "надеть"
	ActionNameTake       ActionName = "взять"
	ActionNameApply      ActionName = "применить"
)

const (
	ItemsNameTea      ItemsName = "чай"
	ItemsNameBackpack ItemsName = "рюкзак"
	ItemsNameNotes    ItemsName = "конспекты"
	ItemsNameKeys     ItemsName = "ключи"
	ItemsNamePhone    ItemsName = "телефон"
)

const (
	PlaceNameTable PlaceName = "на столе"
	PlaceNameChair PlaceName = "на стуле"
)
const (
	interactionPlaceNameDoor   InteractionPlaceName = "дверь"
	interactionPlaceNameСloset InteractionPlaceName = "шкаф"
)

var (
	//переход между комнатами
	mapTransitionRoom = map[RoomName][]RoomName{
		RoomNameKitchen:  []RoomName{RoomNameCorridor},
		RoomNameCorridor: []RoomName{RoomNameKitchen, RoomNameMyRoom, RoomNameStreet},
		RoomNameMyRoom:   []RoomName{RoomNameCorridor},
		RoomNameStreet:   []RoomName{RoomNameHome},
	}

	//какие вещи на какой мебели лежат в комнате
	mapRoomPlace = map[RoomName]map[PlaceName][]ItemsName{
		RoomNameKitchen: mapPlaceItemKitchen,
		RoomNameMyRoom:  mapPlaceItemMyRoom,
	}

	//какая вещь в комнате Kitchen
	mapPlaceItemKitchen = map[PlaceName][]ItemsName{
		PlaceNameTable: []ItemsName{ItemsNameTea},
	}

	//какая вещь в комнате с
	mapPlaceItemMyRoom = map[PlaceName][]ItemsName{
		PlaceNameTable: []ItemsName{ItemsNameKeys, ItemsNameNotes},
		PlaceNameChair: []ItemsName{ItemsNameBackpack},
	}

	//порядок в котором должны располагаться предметы в комнате MyRoom
	placesOrderMyRoom = []PlaceName{PlaceNameTable, PlaceNameChair}

	//взаимодействие предметов
	itemsRules = map[ItemsName]map[InteractionPlaceName]string{
		ItemsNameKeys: {interactionPlaceNameDoor: "дверь открыта"},
	}
)

type IRoom interface {
	LookAroundInfo() string
	TransitionInfo() string
	DeleteItemFromFurniture(string)
	//DeleteItemFromRoom(string)
}

/*
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
*/
