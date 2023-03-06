package post

import "github.com/labstack/echo/v4"

type Delivery interface {
	Create() echo.HandlerFunc
	GetPosts() echo.HandlerFunc
}
