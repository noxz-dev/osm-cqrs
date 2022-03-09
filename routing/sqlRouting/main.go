package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"noxz.dev/routing/db"
	"noxz.dev/routing/updater"
)

func main() {
	go updater.RunUpdate()

	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Get("/count", func(c *fiber.Ctx) error {

		rowCount, err := db.CountRowsInNetwork() 

		if err != nil {
			return c.SendStatus(500)
		}

		return c.SendString(fmt.Sprint(rowCount))
	})

	type DijkstraQueryParams struct{
		FromLng     float64     `query:"fromLng"`
    	FromLat     float64     `query:"fromLat"`
    	ToLng 		float64   	`query:"toLng"`
		ToLat		float64		`query:"toLat"`
	}


	app.Get("/routeDijkstra", func(c *fiber.Ctx) error {

		dqp := new(DijkstraQueryParams)

		if err := c.QueryParser(dqp); err != nil {
			log.Println(err.Error())
			return c.SendStatus(500)
		}

		route, err := db.GetRouteDijkstra(dqp.FromLat, dqp.FromLng, dqp.ToLat, dqp.ToLng) 

		if err != nil {
			log.Println(err.Error())
			return c.SendStatus(500)
		}

		return c.JSON(route)
	})

    log.Fatal(app.Listen(":3000"))
}

