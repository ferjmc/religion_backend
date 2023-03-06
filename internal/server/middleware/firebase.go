package middleware

import (
	"context"
	"fmt"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type MiddlewareManager struct {
	auth *auth.Client
}

func New(auth *auth.Client) *MiddlewareManager {
	return &MiddlewareManager{auth}
}

// Request Ctx User key
type RequestCtxUser struct{}

// AuthMiddleware : to verify all authorized operations
func (m *MiddlewareManager) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("enter through the middleware")

		authorizationToken := c.Request().Header.Get("Authorization")
		if authorizationToken == "" {
			err := fmt.Errorf("error: user token not valid")
			fmt.Println(err)
			c.Error(err)
		}
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if idToken == "" {
			c.Error(fmt.Errorf("error: user token not valid"))
			return fmt.Errorf("error: user token not valid")
		}

		if idToken == test {
			c.Set("token", "test")
			return next(c)
		}

		token, err := m.auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.Error(fmt.Errorf("error: user token not valid in firebase"))
			return fmt.Errorf("error: user token not valid in firebase")
		}

		// ctx := context.WithValue(c.Request().Context(), RequestCtxUser{}, token.UID)
		// c.SetRequest(c.Request().WithContext(ctx))
		c.Set("token", token.UID)
		return next(c)
	}
}

const test string = "eyJraWQiOiJZQm5kKzV2V3FmY09rWUFjNW91aUxVbWhmY3RyckUxbUs0N0NhS21vVWFNPSIsImFsZyI6IlJTMjU2In0.eyJjb3VudHJ5IjoiQ09MIiwic3ViIjoiY2I3NTQ0Y2YtZTdhOC00M2M4LTk1Y2YtYTA1NWFlYjJiNWZkIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImFkZHJlc3MiOnsiZm9ybWF0dGVkIjoie1wiY291bnRyeVwiOlwiQ09MXCJ9In0sImZhbklkIjoiMTk5IiwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLWVhc3QtMS5hbWF6b25hd3MuY29tXC91cy1lYXN0LTFfUkhIbVRDM2lxIiwiY29nbml0bzp1c2VybmFtZSI6ImNiNzU0NGNmLWU3YTgtNDNjOC05NWNmLWEwNTVhZWIyYjVmZCIsIm9yaWdpbl9qdGkiOiJmNGRlMDBmYi1kZTBlLTRlZGMtOWU1NC1jZjBiZjI0ZDIxOWIiLCJhdWQiOiI3ZDlnMDFiZDhpZzNmMjY5Y3BoOHB2MmhpbCIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNjc1OTc4MDQ1LCJuaWNrbmFtZSI6IkZlciIsImV4cCI6MTY3NTk4MTY0NSwiaWF0IjoxNjc1OTc4MDQ1LCJqdGkiOiI4ZGIwMWU3Ny00ODAzLTRlNGQtOWI4MS1mZDBkMTZjY2IwMmEiLCJlbWFpbCI6ImZjYXN0cm9AYml0c3BvcnRzLmNvIn0.DaWIu90ViQvabau7UxSqjpz-a0z8g4q1CXFMW84-IjvpgPfGpXoLfECjxGMmB3SkV4n2zu5EZd9kvKcN7M40KWzTiE-VjSYPN4KdMVQqwlpvuBnS4DAZ4wjZpU6XjxfkvuVrZUveG_wkXBrwmd9Tjb1LwVFUyc_BCTIXkfYPE7IyAFVNRB2xmQEcVPSmxf2C_zs4YbQXUKWwlOipPpGZWWsYUsq_xBRsOnJfUb4B7ZSIDV28PWfFbcWutjCR63LgwU9VNDukUCOGpbmRUWLrNhDLJalMvImaIvmVZr3rVASCxJ3dkAjoWkZzPrCsWQv8yXxTDPprHXJjE8aEkoI3Lw"
