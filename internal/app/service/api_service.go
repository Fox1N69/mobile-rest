package service

import (
	"mobile/internal/app/repository"
)

type ApiService struct {
	repository repository.ApiRepository
}

func NewApiService(apiRepository repository.ApiRepository) *ApiService {
	return &ApiService{repository: apiRepository}
}

