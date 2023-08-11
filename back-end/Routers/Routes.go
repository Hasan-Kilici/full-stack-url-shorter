package Routers

import(
	"urlShorter/Middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"urlShorter/Handlers"
)

func Initalize(app *fiber.App){
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/", Handlers.HelloWorld)
	app.Post("/r/:Token", Handlers.RedirectPage)

	app.Post("/login", Handlers.Login)
	app.Post("/register",Handlers.Register)

	app.Get("/find/user/:Token", Handlers.FindUser)
	app.Get("/list/user/urls", Handlers.ListUrl)
	app.Get("/list/clicked/:UrlToken",Handlers.ListClicks)

	app.Post("/create/url", Handlers.CreateUrl)

	app.Delete("/delete/url/:Token", Handlers.DeleteUrl)

	app.Use(Middleware.Security)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}