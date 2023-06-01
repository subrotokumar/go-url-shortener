package server

import (
	"fmt"
	"strconv"

	"github.com/go-url-shortener/model"
	"github.com/go-url-shortener/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	redirectUrl := c.Params("redirect")
	link, err := model.FindByLinkUrl(redirectUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error link not found" + err.Error(),
		})
	}
	link.Clicked += 1
	err = model.UpdateLink(link)
	if err != nil {
		fmt.Printf("Error updating redirect")
	}
	return c.Redirect(link.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllRedirect(c *fiber.Ctx) error {
	links, err := model.GetAllLinks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all Links" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(links)
}

func getLink(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing id" + err.Error(),
		})
	}
	link, err := model.GetLink(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrieve Link from db" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(link)
}

func createLink(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var link model.Link
	err := c.BodyParser(&link)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing json" + err.Error(),
		})
	}
	if link.Random {
		link.Link = utils.RandomUrl(8)
	}

	err = model.CreateLink(link)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error in creating link in db" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(link)
}

func updateLink(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var link model.Link
	err := c.BodyParser(&link)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error in parsing json" + err.Error(),
		})
	}

	err = model.UpdateLink(link)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error in updating link in DB" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(link)
}

func deleteLink(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error in parsing id from url" + err.Error(),
		})
	}

	err = model.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error in deleting link from db" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Link deleted",
	})
}

func ServeAndListen() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/r/:redirect", redirect)
	app.Get("/link", getAllRedirect)
	app.Get("/link/:id", getLink)
	app.Post("/link", createLink)
	app.Patch("/link", updateLink)
	app.Delete("/link/:id", deleteLink)
	app.Listen(":3000")
}
