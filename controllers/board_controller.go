package controllers

import (
	"github.com/LeannXy/Project_Management/models"
	"github.com/LeannXy/Project_Management/services"
	"github.com/LeannXy/Project_Management/utils"
	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{service: s}
}
func (c *BoardController) CreateBoard(ctx fiber.Ctx) error {
	var userID uuid.UUID
	board := new(models.Board)
	user := jwtware.FromContext(ctx)
	claims := user.Claims.(jwt.MapClaims)

	if err := ctx.Bind().Body(board); err != nil {
		return utils.BadRequest(ctx, "Gagal membaca request", err.Error())
	}

	userID, err := uuid.Parse(claims["pub_id"].(string))
	if err != nil {
		return utils.BadRequest(ctx, "Gagal membaca request", err.Error())
	}
	board.OwnerPublicID = userID

	if err := c.service.Create(board); err != nil {
		return utils.BadRequest(ctx, "Gagal Menyimpan Data", err.Error())
	}
	return utils.Success(ctx, "Board berhasil dibuat", board)

}
