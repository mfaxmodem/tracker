package validator

import (
    "github.com/go-playground/validator/v10"
)

type CustomValidator struct {
    validator *validator.Validate
}

func NewValidator() *CustomValidator {
    v := validator.New()
    
    // Custom validation for coordinates
    v.RegisterValidation("latitude", validateLatitude)
    v.RegisterValidation("longitude", validateLongitude)
    
    return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}

func validateLatitude(fl validator.FieldLevel) bool {
    lat := fl.Field().Float()
    return lat >= -90 && lat <= 90
}

func validateLongitude(fl validator.FieldLevel) bool {
    lon := fl.Field().Float()
    return lon >= -180 && lon <= 180
}