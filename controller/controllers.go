package controller

import (
	"echoServise/model"
	"echoServise/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//get
func GetUser(store repository.UserStore) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, e := strconv.Atoi(id)
		if e != nil {
			return e
		}
		user := store.GetUser(idInt)
		if user == nil {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return c.JSON(200, user)

	}

}

//get
func GetUsersList(store repository.UserStore) func(c echo.Context) error {
	return func(c echo.Context) error {
		users := store.GetUsers()
		return c.JSON(200, users)
	}

}

//post
func AddUser(store repository.UserStore) func(c echo.Context) error {
	return func(c echo.Context) error {
		u := model.User{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		err := store.AddUser(u.Name)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	}

}

//put
func UpdateUser(store repository.UserStore) func(c echo.Context) error {
	return func(c echo.Context) error {
		u := model.User{}
		if err := c.Bind(&u); err != nil {
			return err
		}
		id := c.Param("id")
		idInt, e := strconv.Atoi(id)
		if e != nil {
			return e
		}
		u.ID = idInt

		err := store.UpdateUser(u)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, nil)
	}

}

//delete
func RemoveUser(store repository.UserStore) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, e := strconv.Atoi(id)
		if e != nil {
			return e
		}
		err := store.RemoveUser(idInt)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, nil)
	}

}
