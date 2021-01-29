package model

import "fmt"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Users struct {
	List []User `json:"users"`
}

func (u Users) GetUser(id int) *User {
	for _, v := range u.List {
		if v.ID == id {
			return &v
		}
	}
	return nil
}
func (u *Users) RemoveUser(id int) error {
	var pivot int
	if len(u.List) == 0 {
		return fmt.Errorf("user does not found")
	}
	for i, v := range u.List {
		if v.ID == id {
			pivot = i
			break
		}
		return fmt.Errorf("user does not found")

	}
	if id == 0 {
		pivot++
		u.List = u.List[pivot:]
		return nil
	}
	list1 := u.List[:pivot]
	pivot++
	list2 := u.List[pivot:]
	u.List = append(list1, list2...)
	return nil
}
func (u *Users) UpdateUser(user User) error {
	for i, v := range u.List {
		if v.ID == user.ID {
			u.List[i] = user
			return nil
		}

	}
	return fmt.Errorf("user does not found")
}
func (u *Users) AddUser(name string) {
	id := len(u.List)
	id++
	user := User{ID: id, Name: name}
	u.List = append(u.List, user)
}
