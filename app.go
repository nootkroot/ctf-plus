package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nootkroot/ctf-plus/pkg/log"
	"github.com/nootkroot/ctf-plus/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

func main() {
	cfg := config.LoadConfig()

	viewsPath, err := filepath.Abs(fmt.Sprintf("./themes/"+cfg.Theme+"/views"))
	if err != nil {
		log.Fatal("Error while converting to absolute path", err)
	}
	
	if _, err := os.Stat(viewsPath); err != nil {
		log.Fatal("Error while loading theme", err)
	}

	engine := django.New(viewsPath, ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Listen(":3000")
}