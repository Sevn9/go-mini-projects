package main

type IItem interface {
	AddUsage(target string, result string)
	UseItem(target string) string
}

type Item struct {
	ItemName        string
	ItemInteraction map[string]string
}

func NewItem(itemName string) *Item {

	item := Item{
		ItemName:        itemName,
		ItemInteraction: make(map[string]string),
	}

	return &item
}

func (i *Item) AddUsage(target string, result string) {
	i.ItemInteraction[target] = result
}

func (i *Item) UseItem(target string) string {
	answer, exist := i.ItemInteraction[target]

	if !exist {
		return "не к чему применить"
	}
	return answer
}
