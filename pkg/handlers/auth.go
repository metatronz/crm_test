package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HeaderAuth(ctx *fiber.Ctx) error {
	authString := ctx.Get("authorization")
	var authArr = []string{}
	if len(authString) > 0 {
		authArr = strings.SplitN(authString, " ", 2)
		if len(authArr) == 2 && authArr[0] == "Bearer" {
			ctx.Locals("token", authArr[1])

			return ctx.Next()

		}
	}
	// return ctx.SendStatus(fiber.StatusUnauthorized)
	return ctx.Next()

}

func (env Database) AuthMiddleware(ctx *fiber.Ctx) error {
	bearer := ctx.Locals("token")
	fmt.Println(bearer)
	var user string
	err := env.AppDB.QueryRow("select user from testuser where auth_token = ? ", bearer).Scan(&user)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	} else {
		return ctx.Next()
	}

}
