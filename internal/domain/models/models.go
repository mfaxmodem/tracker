package models

import "time"

type User struct {
    ID           int64     `json:"id"`
    Name         string    `json:"name"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"`
    Role         string    `json:"role"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type Store struct {
    ID          int64     `json:"id"`
    Name        string    `json:"name"`
    Address     string    `json:"address"`
    Latitude    float64   `json:"latitude"`
    Longitude   float64   `json:"longitude"`
    ManagerName string    `json:"manager_name"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Route struct {
    ID        int64     `json:"id"`
    VisitorID int64     `json:"visitor_id"`
    Status    string    `json:"status"`
    StartDate time.Time `json:"start_date"`
    EndDate   time.Time `json:"end_date"`
    StoreIDs  []int64   `json:"store_ids"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Location struct {
    ID        int64     `json:"id"`
    VisitorID int64     `json:"visitor_id"`
    Latitude  float64   `json:"latitude"`
    Longitude float64   `json:"longitude"`
    Timestamp time.Time `json:"timestamp"`
}