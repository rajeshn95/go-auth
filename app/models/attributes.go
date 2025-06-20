package models

import (
	"time"

	"github.com/google/uuid"
)

// Book struct to describe book object.
type Attribute struct {
	Uuid       uuid.UUID `db:"uuid" json:"uuid" validate:"required,uuid"`
	UserUuid   uuid.UUID `db:"user_uuid" json:"user_uuid" validate:"required,uuid"`
	Type       string    `db:"type" json:"type" validate:"required,lte=255"`
	Value      string    `db:"value" json:"value" validate:"required,lte=255"`
	Purpose    string    `db:"purpose" json:"purpose" validate:"required,lte=255"`
	VerifiedAt time.Time `db:"verified_at" json:"verified_at"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at" json:"deleted_at"`
}

// Profile struct to describe profile attributes.
// type Profile struct {
// }

// Value make the Profile struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
// func (b *Profile) Value() (driver.Value, error) {
// 	return json.Marshal(b)
// }

// Scan make the Profile struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
// func (b *Profile) Scan(value interface{}) error {
// 	j, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("type assertion to []byte failed")
// 	}

// 	return json.Unmarshal(j, &b)
// }
