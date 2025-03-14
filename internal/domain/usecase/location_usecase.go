package usecase

import (
    "tracker/internal/domain/models"
    "tracker/internal/domain/repository"
)

type locationUsecase struct {
    repo repository.Repository
}

func NewLocationUsecase(repo repository.Repository) LocationUsecase {
    return &locationUsecase{
        repo: repo,
    }
}

func (u *locationUsecase) TrackLocation(location *models.Location) error {
    return u.repo.SaveLocation(location)
}

func (u *locationUsecase) GetVisitorRoutes(visitorID int64) ([]models.Route, error) {
    return u.repo.GetVisitorRoutes(visitorID)
}