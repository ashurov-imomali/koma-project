package service

import "main/internal/repository"

type Service struct {
	Repos *repository.Repository
}

func NewService(s *repository.Repository) *Service {
	return &Service{Repos: s}
}
