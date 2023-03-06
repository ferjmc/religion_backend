package user

import "github.com/labstack/echo/v4"

type Delivery interface {
	CreateUser() echo.HandlerFunc
	GetSingleUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
}
