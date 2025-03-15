package repository

import "github.com/mfaxmodem/tracker/internal/domain/models"

type Repository interface {
    // User operations
    CreateUser(*models.User) error
    GetUserByEmail(string) (*models.User, error)
    GetAllVisitors() ([]models.User, error)
    UpdateUser(*models.User) error
    DeleteUser(int64) error

    // Store operations
    SaveStore(*models.Store) error
    GetStores() ([]models.Store, error)
    UpdateStore(*models.Store) error
    DeleteStore(int64) error

    // Route operations
    SaveRoute(*models.Route) error
    GetAllRoutes() ([]models.Route, error)
    GetVisitorRoutes(int64) ([]models.Route, error)
    UpdateRoute(*models.Route) error
    DeleteRoute(int64) error

    // Location operations
    SaveLocation(*models.Location) error
    GetLocations(int64) ([]models.Location, error)
}