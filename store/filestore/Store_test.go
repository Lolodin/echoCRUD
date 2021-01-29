package filestore

import (
	"echoServise/model"
	"fmt"
	"testing"
)

func TestNewStore(t *testing.T) {
	s:= NewStore()
	fmt.Println(s)
}
func TestUserStore_AddUser(t *testing.T) {
	s:= NewStore()
	e := s.AddUser("TEST CASE")
	if e!= nil {
		fmt.Println(s)
		t.Fatal(e)
	}
}
func TestUserStore_GetUser(t *testing.T) {
	s:= NewStore()
	u:=s.GetUser(1)
	if u.ID != 1 {
		t.Error(u)
	}
}
func TestUserStore_RemoveUser(t *testing.T) {
	s:= NewStore()
	e:=s.RemoveUser(1)
	if e != nil {
		t.Error(e)
	}
	e=s.RemoveUser(6)
	if e != nil {
		t.Log("OK")
	}

}
func TestUserStore_UpdateUser(t *testing.T) {
	s := NewStore()
	m := model.User{1, "TESTUPDATE"}
	e := s.UpdateUser(m)
	if e != nil {
		t.Error(e)
	}
	m = model.User{33, "TESTUPDATE"}
	err := s.UpdateUser(m)
	if err != nil {
		t.Log("OK")
	}
}
func TestUserFileWrite(t *testing.T) {

}