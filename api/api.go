package api

import (
	"server_deploy/repository"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	res := struct {
		Status  int    `json:"Status"`
		Message string `jsons:"message"`
	}{
		Status:  fiber.StatusOK,
		Message: "Connection Success",
	}
	return c.JSON(res)
}

func Follower(c *fiber.Ctx) error {
	username := c.Params("username")
	fol, err := repository.FindFollowers(username)
	if err != nil {
		return c.JSON(err.Error())
	}

	res := struct {
		Followers int `json:"followers"`
	}{
		Followers: fol,
	}
	return c.JSON(res)
}

func Userid(c *fiber.Ctx) error {
	userid := c.Params("userid")
	user, fol, err := repository.FindUser(userid)
	if err != nil {
		return c.JSON(err.Error())
	}

	res := struct {
		Username  string `json:"username"`
		Followers int    `json:"followers"`
	}{
		Username:  user,
		Followers: fol,
	}
	return c.JSON(res)
}
