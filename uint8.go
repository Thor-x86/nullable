package nullable

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// Uint8 SQL type that can retrieve NULL value
type Uint8 struct {
	realValue uint8
	isValid   bool
}

// NewUint8 creates a new nullable 8-bit unsigned integer
func NewUint8(value *uint8) Uint8 {
	if value == nil {
		return Uint8{
			realValue: 0,
			isValid:   false,
		}
	}
	return Uint8{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or 8-bit unsigned integer
func (n Uint8) Get() *uint8 {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or 8-bit unsigned integer
func (n *Uint8) Set(value *uint8) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Uint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Uint8) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed uint8
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Uint8) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}

	var scanned string
	if err := convertAssign(&scanned, value); err != nil {
		return err
	}

	radix := 10
	if len(scanned) == 8 {
		radix = 2
	}

	parsed, err := strconv.ParseUint(scanned, radix, 8)
	if err != nil {
		return err
	}
	n.realValue = uint8(parsed)

	n.isValid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n Uint8) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return strconv.FormatUint(uint64(n.realValue), 10), nil
}

// GormValue implements the driver Valuer interface via GORM.
func (n Uint8) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		// MySQL and SQLite are using Value() instead of GormValue()
		value, err := n.Value()
		if err != nil {
			db.AddError(err)
			return clause.Expr{}
		}
		return clause.Expr{SQL: "?", Vars: []interface{}{value}}
	case "postgres":
		if !n.isValid {
			return clause.Expr{SQL: "?", Vars: []interface{}{nil}}
		}
		value := strconv.FormatUint(uint64(n.realValue), 2)
		leadingZero := 8 - len(value)
		for leadingZero > 0 {
			value = "0" + value
			leadingZero--
		}
		return clause.Expr{SQL: "?", Vars: []interface{}{value}}
	}
	return clause.Expr{}
}

// GormDataType gorm common data type
func (Uint8) GormDataType() string {
	return "uint8_null"
}

// GormDBDataType gorm db data type
func (Uint8) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "TINYINT UNSIGNED"
	case "postgres":
		return "bit(8)"
	}
	return ""
}
