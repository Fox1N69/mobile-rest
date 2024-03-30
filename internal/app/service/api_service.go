package service

import (
	"mobile/internal/app/repository"

	"github.com/gofiber/fiber/v3"
)

type ApiService struct {
	repository repository.ApiRepository
}

func NewApiService(apiRepository repository.ApiRepository) *ApiService {
	return &ApiService{repository: apiRepository}
}

func (s *ApiService) GetAllData(c fiber.Ctx) error {
	return nil
}
