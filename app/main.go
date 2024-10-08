package main

import "github.com/gofiber/fiber/v2"

func main() {
	router := fiber.New()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("golang-argocd-kustomize is start")
	})

	err := router.Listen(":8080")
	if err != nil {
		panic("start server error")
	}
}
