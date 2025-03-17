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

  app.Get("/premier-league/fixtures", func(c *fiber.Ctx) error {
    fixtures, err := scraper.ScrapePremierLeagueFixtures()
    if err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "error": "Error scraping fixtures: " + err.Error(),
      })
    }

    return c.JSON(fixtures)
  })

  app.Get("/premier-league/match-results", func(c *fiber.Ctx) error {
    results, err := scraper.ScrapePremierLeagueMatchResults()
    if err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "error": "Error scraping match results: " + err.Error(),
      })
    }

    return c.JSON(results)
  })

  app.Get("/premier-league/rankings", func(c *fiber.Ctx) error {
    results, err := scraper.ScrapePremierLeagueRankings()
    if err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "error": "Error scraping rankings: " + err.Error(),
      })
    }

    return c.JSON(results)
  })

  // 404 Not Found Handling
  app.Use(func(c *fiber.Ctx) error {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "error": "Not Found",
    })
  })

  if err := app.Listen(":" + strconv.Itoa(port)); err != nil {
    panic(err)
  }
}