package models

import "time"

type Location struct {
    ID        int64     `json:"id"`
    VisitorID int64     `json:"visitor_id"`
    Latitude  float64   `json:"latitude"`
    Longitude float64   `json:"longitude"`
    Timestamp time.Time `json:"timestamp"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}