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
	placesOrderMyRoom  = []PlaceName{PlaceNameTable, PlaceNameChair}
	placesOrderKitchen = []PlaceName{PlaceNameTable}
)

type IRoom interface {
	LookAroundInfo() string
	TransitionInfo() string
	DeleteItemFromRoom(string)
	IsItemExistThisRoom(string) bool
	ApplyItem(item ItemsName, interactionPlace InteractionPlaceName) string
	CanExitTo(roomName RoomName) (bool, string)
}

func GetExits(roomName RoomName) string {
	_, roomNameExist := rooms[roomName]

	if !roomNameExist {
		return "такой комнаты не существует"
	}

	answer := "можно пройти - "
	for placeNameRoom, exitsRoom := range mapTransitionRoom {

		exitsNameCounter := len(exitsRoom)

		if placeNameRoom == roomName {
			for i, exit := range exitsRoom {
				answer += string(exit)
				if i < exitsNameCounter-1 {
					answer += ", "
				}
			}
		}
	}
	return answer
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
