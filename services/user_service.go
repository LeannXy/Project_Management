package services

import (
	"errors"

	"github.com/LeannXy/Project_Management/models"
	"github.com/LeannXy/Project_Management/repositories"
	"github.com/LeannXy/Project_Management/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
	Login(email,Password string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByPublicID(id string) (*models.User, error)
	GetAllPagination(filter, sort string, limit, offset int)([]models.User, int64, error)
	
}

type userService struct{
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
		return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	//mengecek aakah email sudah terdaftar
	//hasing pw
	//set role
	existingUser , _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("email alredy registered")
	}
	hased, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hased
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.repo.Create(user)
}

func (s *userService) Login(email,Password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credential")
	}
	if !utils.CheckPasswordHash(Password, user.Password){
		return nil, errors.New("invalid credential")
	}
	return user, nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return  s.repo.FindByID(id)
}

func (s *userService) GetByPublicID(id string) (*models.User, error) {
	return  s.repo.FindByPublicID(id)
}
func (s *userService) GetAllPagination(filter, sort string, limit, offset int)([]models.User, int64, error){
	return s.repo.FindAllPagination(filter, sort, limit, offset )
}