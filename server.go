package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func updateUser(c echo.Context) error {
	// Get name and email
	id := c.FormValue("id")
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.String(http.StatusOK, "name:"+name+", email:"+email+", id:"+id)
}

func deleteUser(c echo.Context) error {
	// Get name and email
	id := c.Param("id")

	return c.String(http.StatusOK, "id:"+id)
}
