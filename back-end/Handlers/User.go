package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"urlShorter/Utils"
	"urlShorter/Database"
)

func Register(c *fiber.Ctx) error {
	form := new(RegisterForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	if !Utils.ValidateEmail(form.Email){
		return c.Status(400).SendString("Bu email geçersiz")
	}

	if !Utils.ValidatePassword(form.Pass) {
		return c.Status(400).SendString("Şifre yeterince güvenli değil")
	}

	if !Utils.ValidateUsername(form.Username) {
		return c.Status(400).SendString("Kullanıcı Adı Boşluk içermemeli")
	}

	if !Database.Register(form.Username, form.Email, form.Pass){
		return c.Status(400).SendString("Bu Email kullanılıyor.")
	}

	login, _ := Database.Login(form.Email,form.Pass)

	c.JSON(fiber.Map{
		"Token": login,
	})
	return nil
}

func Login(c *fiber.Ctx) error {
	form := new(LoginForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	user := Database.FindUserByEmail(form.Email)
	if !user {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}

	login, err := Database.Login(form.Email,form.Pass)
	if err != nil{
		return c.Status(400).SendString("E-mail yada şifre yanlış")
	}

	c.JSON(fiber.Map{
		"Token": login,
	})
	return nil
}
