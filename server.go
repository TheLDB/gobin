package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/:code", func(c echo.Context) error {
		status := c.Param("code")

		s, err := strconv.Atoi(status)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, &Response{
				Success: false,
				Message: "Invalid Status Code",
			})
		}

		return c.JSON(s, &Response{
			Success: s >= 200 && s <= 399,
			Message: http.StatusText(s),
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
