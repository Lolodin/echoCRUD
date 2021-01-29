package filestore

import (
	"echoServise/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	 PATHFILE = "bd/BD.json"
	 DIR = "bd"
)


type UserStore struct {
	os.File
	
}

func NewStore() UserStore  {
	us:= UserStore{}
	os.Mkdir(DIR, os.ModeDir)
	f,e:=os.Create(PATHFILE)
	defer f.Close()
	if e != nil {
		log.Fatal(e)
	}
	us.File = *f
	us.File.Write([]byte(`{ "users" : [ {"id": 1, "name" : "test"},{"id": 2, "name" : "test3"}]}`))

	return us
}

func (u UserStore) GetUser(id int) *model.User {
	um := u.getModel()
	return um.GetUser(id)
}

func (u UserStore) GetUsers() []model.User {
	um := u.getModel()
	return um.List
}

func (u UserStore) RemoveUser(id int) error {
	um := u.getModel()
	e:=um.RemoveUser(id)
	if e != nil {
		return e
	}
	b, e := json.Marshal(um)
	if e != nil {
		return e
	}
	return ioutil.WriteFile(PATHFILE,b, os.ModeAppend)


}

func (u UserStore) AddUser(name string) error {
	um := u.getModel()
	um.AddUser(name)
	b, e := json.Marshal(um)
	if e != nil {
		return e
	}
	return ioutil.WriteFile(PATHFILE,b, os.ModeAppend)


}

func (u UserStore) UpdateUser(user model.User) error {
	um := u.getModel()
	e :=um.UpdateUser(user)
	if e != nil {
		return e
	}

	b, e := json.Marshal(um)
	if e != nil {
		return e
	}
	return ioutil.WriteFile(PATHFILE,b, os.ModeAppend)

}

func (u UserStore) getModel() *model.Users {
	b, e := ioutil.ReadFile(PATHFILE)
	if e!= nil {
		return nil
	}
	um := model.Users{}
	json.Unmarshal(b, &um)
	return &um
}