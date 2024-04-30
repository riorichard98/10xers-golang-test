package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	// entity of phone
	Phone struct {
		ID uuid.UUID `json:"id" db:"id" gorm:"type:uuid;primaryKey"`
		gorm.Model
		Name          string  `json:"name" db:"name"`
		Brand         string  `json:"brand" db:"brand"`
		Price         float64 `json:"price" db:"price"`
		StockQuantity int     `json:"stock" db:"stock"`
	}

	// phone domain (there is only db that will be used for now)
	PhoneDomain struct {
		db *gorm.DB
	}

	// methods to be used
	PhoneInterface interface {
		CreatePhone(phone *Phone) error
		GetPhoneByID(id uuid.UUID) (*Phone, error)
		UpdatePhone(phone *Phone) error
		DeletePhoneByID(id uuid.UUID) error
		SearchPhonesByName(name string) ([]Phone, error)
	}
)

// function for creating New Phone domain for use in usecase
// this function provide flexibility(example: diffrent database) and validation in early setup
func NewPhoneDomain(db *gorm.DB) *PhoneDomain {
	// first check if db is there or not before initiate
	if db == nil {
		panic("db is nil")
	}

	return &PhoneDomain{
		db: db,
	}
}

// Create a new phone record
func (pd *PhoneDomain) CreatePhone(phone *Phone) error {
	result := pd.db.Create(phone)
	return result.Error
}

// Get a phone record by ID
func (pd *PhoneDomain) GetPhoneByID(id uuid.UUID) (*Phone, error) {
	var phone Phone
	result := pd.db.First(&phone, id)
	return &phone, result.Error
}

// Update a phone record
func (pd *PhoneDomain) UpdatePhone(phone *Phone) error {
	result := pd.db.Save(phone)
	return result.Error
}

// Delete a phone record by ID
func (pd *PhoneDomain) DeletePhoneByID(id uuid.UUID) error {
	result := pd.db.Delete(&Phone{}, id)
	return result.Error
}

// Search phones by name
func (pd *PhoneDomain) SearchPhonesByName(name string) ([]Phone, error) {
	var phones []Phone
	result := pd.db.Where("name LIKE ?", "%"+name+"%").Find(&phones)
	return phones, result.Error
}
