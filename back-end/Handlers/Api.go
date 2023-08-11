package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"urlShorter/Database"
)

func HelloWorld(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"hello":"world",
	})
	return nil
}

func FindUser(c *fiber.Ctx) error {
	Token := c.Params("Token")
	User, err := Database.FindUserByToken(Token);
	if err != nil {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}

	c.JSON(fiber.Map{
		"data":User,
	})
	return nil
}

func ListClicks(c *fiber.Ctx) error {
	urlToken := c.Params("UrlToken")
	Url,err  := Database.FindUrl(urlToken)
	if err != nil {
		return c.Status(404).SendString("Url bulunamadı")
	}

	List, _ := Database.ListClicks(Url.Token)
	c.JSON(fiber.Map{
		"data":List,
	})
	return nil
}

func ListUrl(c *fiber.Ctx) error {
	Token := c.Cookies("Token")
	User,err := Database.FindUserByToken(Token)
	if err != nil {
		return c.Status(404).SendString("kullanıcı bulunamadı")
	}

	Urls, _ := Database.ListUrls(User.Token)

	c.JSON(fiber.Map{
		"data":Urls,
	})
	return nil
}