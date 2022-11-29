package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/repentance/rest_vid_api/localdata"
)

func ReadData(ctx *fiber.Ctx) error {
	data := localdata.InitData()
	return ctx.JSON(data.GetData())
}
