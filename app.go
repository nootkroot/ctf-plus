package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nootkroot/ctf-plus/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

func main() {
	conf := config.LoadConfig()

	viewsPath, err := filepath.Abs(fmt.Sprintf("./themes/"+conf.Theme+"/views"))
	if err != nil {
		log.Fatalf("Error while converting to absolute path: %s", err)
	}
	
	if _, err := os.Stat(viewsPath); err != nil {
		log.Fatalf("Error while loading theme: %s", err)
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
