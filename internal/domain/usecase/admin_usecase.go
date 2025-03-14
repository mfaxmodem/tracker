package usecase

import (
    "tracker/internal/domain/models"
    "tracker/internal/domain/repository"
)

type adminUsecase struct {
    repo repository.Repository
}

func NewAdminUsecase(repo repository.Repository) AdminUsecase {
    return &adminUsecase{
        repo: repo,
    }
}

func (u *adminUsecase) GetAllVisitors() ([]models.User, error) {
    return u.repo.GetAllVisitors()
}

func (u *adminUsecase) CreateVisitor(user *models.User) error {
    user.Role = "visitor"
    return u.repo.CreateUser(user)
}

func (u *adminUsecase) UpdateVisitor(user *models.User) error {
    return u.repo.UpdateUser(user)
}

func (u *adminUsecase) DeleteVisitor(id int64) error {
    return u.repo.DeleteUser(id)
}

func (u *adminUsecase) GetAllStores() ([]models.Store, error) {
    return u.repo.GetStores()
}

func (u *adminUsecase) CreateStore(store *models.Store) error {
    return u.repo.SaveStore(store)
}

func (u *adminUsecase) UpdateStore(store *models.Store) error {
    return u.repo.UpdateStore(store)
}

func (u *adminUsecase) DeleteStore(id int64) error {
    return u.repo.DeleteStore(id)
}

func (u *adminUsecase) GetAllRoutes() ([]models.Route, error) {
    return u.repo.GetAllRoutes()
}

func (u *adminUsecase) CreateRoute(route *models.Route) error {
    return u.repo.SaveRoute(route)
}

func (u *adminUsecase) UpdateRoute(route *models.Route) error {
    return u.repo.UpdateRoute(route)
}

func (u *adminUsecase) DeleteRoute(id int64) error {
    return u.repo.DeleteRoute(id)
}

func (u *adminUsecase) GetUserByEmail(email string) (*models.User, error) {
    return u.repo.GetUserByEmail(email)
}