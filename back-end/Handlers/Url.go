package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"urlShorter/Database"
)

func CreateUrl(c *fiber.Ctx) error {
	form := new(CreateUrlForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	Token := c.Cookies("Token")
	User, err := Database.FindUserByToken(Token)
	if err != nil {
		return c.Status(404).SendString("kullanıcı bulunamadı")
	}

	err = Database.CreateUrl(User.Token, form.Redirect)
	c.JSON(fiber.Map{
		"message":"url başarıyla kısaltıldı",
	})
	return nil
}

func RedirectPage(c *fiber.Ctx) error {
	Token := c.Params("Token")

	form := new(ClickedForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	Url,err := Database.FindUrl(Token)
	if err != nil {
		return c.Status(404).SendString("Link bulunamadı")
	}

	err = Database.SaveUserClicked(Token,form.IP,form.Country,form.City,form.Latitude,form.Longitude)
	if err != nil {
		return c.Status(500).SendString("sunucuda bir sorun çıktı")
	}

	c.JSON(fiber.Map{
		"redirect":Url.Redirect,
	})
	return nil
}

func DeleteUrl(c *fiber.Ctx) error {
	UserToken := c.Cookies("Token")
	UrlToken := c.Params("Token")
	Url, err := Database.FindUrl(UrlToken);
	if err != nil {
		return c.Status(404).SendString("Link bulunamadı")
	}

	if(Url.OwnerToken != UserToken){
		return c.Status(401).SendString("Yetersiz yetki")
	}

	Database.DeleteUrl(UrlToken)
	c.JSON(fiber.Map{
		"message":"Url başarıyla silindi",
	})
	return nil
}