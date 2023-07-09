package search

import (
	"github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/search"
)

type Service struct {
	storage search.StorageInterface
}

func NewService(storage search.StorageInterface) *Service {
	return &Service{storage: storage}
}
