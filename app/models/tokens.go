package models

import (
	"time"

	"github.com/google/uuid"
)

// Book struct to describe book object.
type Token struct {
	Uuid        uuid.UUID `db:"uuid" json:"uuid" validate:"required,uuid"`
	Sub         uuid.UUID `db:"sub" json:"sub" validate:"uuid"`
	Issuer      string    `db:"issuer" json:"issuer" validate:"lte=255"`
	Title       string    `db:"title" json:"title" validate:"lte=255"`
	Permissions any       `db:"permissions" json:"permissions"`
	OtherInfo   any       `db:"other_info" json:"other_info"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at" json:"deleted_at"`
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
