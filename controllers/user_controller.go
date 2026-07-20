package controllers

import (

	"github.com/LeannXy/Project_Management/models"
	"github.com/LeannXy/Project_Management/services"
	"github.com/LeannXy/Project_Management/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.Bind().Body(user); err != nil{
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return  utils.BadRequest(ctx, "Registrasi gagal", err.Error())
	}
	var userResp models.UserRespons
	_ = copier.Copy(&userResp, &user)
	return  utils.Success(ctx, "Register success", userResp)
}

func (c *UserController) Login(ctx fiber.Ctx) error {
	var body struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.Bind().Body(&body);err != nil {
		return utils.BadRequest(ctx, "invalid request", err.Error())
	}

	user, err := c.service.Login(body.Email, body.Password)
	if err != nil {
		return utils.Unauthorized(ctx, "Login Failed", err.Error())

		
	}
	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)
		var userResp models.UserRespons
	_ = copier.Copy(&userResp, &user)
	return utils.Success(ctx, "Login Succesful", fiber.Map{
		"access_token": token,
		"refresh_token": refreshToken,
		"user": userResp,
	})
}