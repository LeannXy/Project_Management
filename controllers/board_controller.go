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

func (c *BoardController) UpdateBoard(ctx fiber.Ctx) error {
	publicID := ctx.Params("id")
	board := new (models.Board)

	if err := ctx.Bind().Body(board); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing  data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "ID tidak valid", err.Error())
	}
	existingBoard, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return  utils.NotFound(ctx, "Board tidak ditemukan", err.Error())
	}

	board.InternalID = existingBoard.InternalID
	board.PublicID = existingBoard.PublicID
	board.OwnerPublicID = existingBoard.OwnerPublicID
	board.CreatedAt = existingBoard.CreatedAt


	if err := c.service.Update(board); err != nil {
		return utils.BadRequest(ctx, "Gagal update board", err.Error())
	}
	return utils.Success(ctx, "Board berhasil diperbaharui", board)
}

func (c *BoardController) AddBoardMembers(ctx fiber.Ctx) error{
	publicID := ctx.Params("id")

	var userIDs []string
	if err := ctx.Bind().Body(&userIDs); err != nil{
		return utils.BadRequest(ctx, "Gagal Parshing Data", err.Error())
	}
	if err := c.service.AddMember(publicID, userIDs); err != nil{
		return utils.BadRequest(ctx, "Gagal Menambahkan Members", err.Error())
	}
	return utils.Success(ctx, "Members berhasil ditambahkan", nil)
}
func (c *BoardController) RemoveBoardMembers(ctx fiber.Ctx) error{
	publicID := ctx.Params("id")

	var userIDs []string
	if err := ctx.Bind().Body(&userIDs); err != nil{
		return utils.BadRequest(ctx, "Gagal Parshing Data", err.Error())
	}
	if err := c.service.AddMember(publicID, userIDs); err != nil{
		return utils.BadRequest(ctx, "Gagal Menghapus Member", err.Error())
	}
	return utils.Success(ctx, "Members berhasil dihapus", nil)
}
