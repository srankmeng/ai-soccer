package main

import (
  "fmt"
  "strconv"
  "github.com/gofiber/fiber/v2"
  "soccer-api/pkg/scraper"
)

func main() {
  app := fiber.New()
  port := 3000
  fmt.Println("Soccer API server listening on http://localhost:" + strconv.Itoa(port))

  app.Get("/health", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
      "status": "ok",
    })
  })

  app.Get("/matches", func(c *fiber.Ctx) error {
    matches, err := scraper.ScrapeMatches()
    if err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "error": "Error scraping matches: " + err.Error(),
      })
    }

    return c.JSON(matches)
  })

  if err := app.Listen(":" + strconv.Itoa(port)); err != nil {
    panic(err)
  }
}