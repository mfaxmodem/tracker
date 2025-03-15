package usecase

import (
    "github.com/mfaxmodem/tracker/internal/domain/models"
    "github.com/mfaxmodem/tracker/internal/repository/postgres"
)

type locationUsecase struct {
    repo *postgres.Repository
}

type LocationUsecase interface {
    SaveLocation(location *models.Location) error
    GetVisitorLocations(visitorID int64) ([]models.Location, error)
    TrackLocation(location *models.Location) error
    GetVisitorRoutes(visitorID int64) ([]models.Route, error)
}

func NewLocationUsecase(repo *postgres.Repository) LocationUsecase {
    return &locationUsecase{
        repo: repo,
    }
}

func (lu *locationUsecase) SaveLocation(location *models.Location) error {
    return lu.repo.SaveLocation(location)
}

func (lu *locationUsecase) GetVisitorLocations(visitorID int64) ([]models.Location, error) {
    return lu.repo.GetVisitorLocations(visitorID)
}

func (lu *locationUsecase) TrackLocation(location *models.Location) error {
    return lu.repo.SaveLocation(location)
}

func (lu *locationUsecase) GetVisitorRoutes(visitorID int64) ([]models.Route, error) {
    return lu.repo.GetVisitorRoutes(visitorID)
}