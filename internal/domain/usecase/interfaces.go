package usecase

import "tracker/internal/domain/models"

type AdminUsecase interface {
    GetUserByEmail(email string) (*models.User, error)
    GetAllVisitors() ([]models.User, error)
    CreateVisitor(*models.User) error
    UpdateVisitor(*models.User) error
    DeleteVisitor(int64) error

    GetAllStores() ([]models.Store, error)
    CreateStore(*models.Store) error
    UpdateStore(*models.Store) error
    DeleteStore(int64) error

    GetAllRoutes() ([]models.Route, error)
    CreateRoute(*models.Route) error
    UpdateRoute(*models.Route) error
    DeleteRoute(int64) error
}

type LocationUsecase interface {
    TrackLocation(*models.Location) error
    GetVisitorRoutes(visitorID int64) ([]models.Route, error)
}