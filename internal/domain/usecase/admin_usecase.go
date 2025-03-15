package usecase

import (
    "fmt" // Add this import
    "github.com/mfaxmodem/tracker/internal/domain/models"
    "github.com/mfaxmodem/tracker/internal/repository/postgres"
)

type adminUsecase struct {
    repo *postgres.Repository
}

func NewAdminUsecase(repo *postgres.Repository) AdminUsecase {
    return &adminUsecase{
        repo: repo,
    }
}

type AdminUsecase interface {
    GetAllVisitors() ([]models.User, error)
    CreateVisitor(*models.User) error
    GetAllStores() ([]models.Store, error)
    CreateStore(*models.Store) error
    GetAllRoutes() ([]models.Route, error)
    CreateRoute(*models.Route) error
    CreateUser(*models.User) error
    GetUserByEmail(email string) (*models.User, error) // Add this method
}

func (au *adminUsecase) GetAllVisitors() ([]models.User, error) {
    return au.repo.GetAllVisitors()
}

func (au *adminUsecase) CreateVisitor(user *models.User) error {
    return au.repo.CreateUser(user)
}

func (au *adminUsecase) GetAllStores() ([]models.Store, error) {
    return au.repo.GetStores()
}

func (au *adminUsecase) CreateStore(store *models.Store) error {
    // Check if a store with the same name and address already exists
    existingStore, err := au.repo.GetStoreByNameAndAddress(store.Name, store.Address)
    if err != nil {
        return err
    }

    if existingStore != nil {
        return fmt.Errorf("store with the same name and address already exists")
    }

    // Create the new store
    return au.repo.SaveStore(store)
}

func (au *adminUsecase) GetAllRoutes() ([]models.Route, error) {
    return au.repo.GetAllRoutes()
}

func (au *adminUsecase) CreateRoute(route *models.Route) error {
    return au.repo.SaveRoute(route)
}

func (au *adminUsecase) CreateUser(user *models.User) error {
    return au.repo.CreateUser(user)
}

func (au *adminUsecase) GetUserByEmail(email string) (*models.User, error) {
    return au.repo.GetUserByEmail(email)
}