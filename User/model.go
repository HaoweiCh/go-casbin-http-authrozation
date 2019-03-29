package User

import "errors"

// Item is a user
type Item struct {
	ID   int
	Name string
	Role string
}

// Items is a list of users
type Items []Item

// Exists checks if a user with the given id exists in the list
func (u Items) Exists(id int) bool {
	exists := false
	for _, user := range u {
		if user.ID == id {
			return true
		}
	}
	return exists
}

// FindByName returns the user with the given name, or returns an error
func (u Items) FindByName(name string) (Item, error) {
	for _, user := range u {
		if user.Name == name {
			return user, nil
		}
	}
	return Item{}, errors.New("USER_NOT_FOUND")
}
