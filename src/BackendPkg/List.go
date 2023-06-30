package BackendPkg

import (
	"time"
)

type List struct {
	ListOwner    User
	TimeUpdated  time.Time
	ShoppingList []FoodItem
}

func NewList(listOwner User) List {
	return List{
		ListOwner:    listOwner,
		TimeUpdated:  time.Now(),
		ShoppingList: make([]FoodItem, 0),
	}
}

func (l *List) AddItem(item FoodItem) {
	l.ShoppingList = append(l.ShoppingList, item)
	l.TimeUpdated = time.Now()
}

func (l *List) UpdateItemQuantity(item FoodItem, quantity int) {
	for i, listItem := range l.ShoppingList {
		if listItem.Name == item.Name {
			if quantity > 0 {
				l.ShoppingList[i].Quantity = quantity
			} else {
				// Remove item from the list if quantity is zero
				l.ShoppingList = append(l.ShoppingList[:i], l.ShoppingList[i+1:]...)
			}
			l.TimeUpdated = time.Now()
			return
		}
	}

	// Item not found in the list, add it with the specified quantity
	if quantity > 0 {
		item.Quantity = quantity
		l.AddItem(item)
	}
}

func (l *List) RemoveItem(item FoodItem) {
	for i, listItem := range l.ShoppingList {
		if listItem.Name == item.Name {
			l.ShoppingList = append(l.ShoppingList[:i], l.ShoppingList[i+1:]...)
			l.TimeUpdated = time.Now()
			return
		}
	}
}

// Implement any other methods or functionalities specific to the List type
