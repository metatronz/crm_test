package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func (env *Database) LoginHandler(ctx *fiber.Ctx) error {
	p := new(models.Person)
	if err := ctx.BodyParser(p); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	var pass sql.NullString
	err := env.BillDB.QueryRow("select pass from accounts where login = ?", p.Login).Scan(&pass)
	if err != nil {
		return ctx.SendStatus(fiber.StatusServiceUnavailable)
	}
	if pass.String == p.Pass {
		authToken, err := GenerateSecureToken()
		if err != nil {
			return ctx.SendStatus(fiber.StatusServiceUnavailable)
		}
		result, err := env.AppDB.Exec("insert into testuser (user, auth_token) values (?, ?)", p.Login, authToken)
		fmt.Println(result, err)
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"token": authToken})
	}
	return ctx.SendStatus(fiber.StatusUnauthorized)
}

func GenerateSecureToken() (string, error) {
	b := make([]byte, 255)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	sEnc := base64.URLEncoding.EncodeToString(b)
	return sEnc[0:255], nil
}
