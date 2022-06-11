package handler

import (
	"fmt"
	"goly/model"
	util "goly/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Redirect(c *fiber.Ctx) error {
	golyUrl := c.Params("redirect")
	goly, err := model.FindByGolyUrl(golyUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find goly in DB " + err.Error(),
		})
	}
	// update click count
	goly.Clicked += 1

	err = model.UpdateGoly(goly)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func CreateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}
    goly.Goly = util.GenerateShortUrl(goly.Redirect)
	err = model.CreateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create goly in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)

}

func UpdateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly

	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json " + err.Error(),
		})
	}

	golyID, err := strconv.Atoi(c.Params("id"))
	existing_goly, err := model.GetGoly(golyID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not goly with given ID" + err.Error(),
		})
	}

	existing_goly.Redirect = goly.Redirect
	

	err = model.UpdateGoly(existing_goly)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update goly link in DB " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)

}

func GetGoly(c *fiber.Ctx) error{
	golyID, _ := strconv.Atoi(c.Params("id"))
	goly , err := model.GetGoly(golyID)
	if err !=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find goly with given ID " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)

}

func DeleteGoly(c *fiber.Ctx ) error {
	golyID , _ := strconv.Atoi(c.Params("id"))
	goly, err := model.GetGoly(golyID)
	if err !=nil {
		if err !=nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "could not find goly with given ID" + err.Error(),
			})
		}
	}
	deleteErr := model.DeleteGoly(goly)
	if deleteErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not  delete goly with given ID" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}
