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

// Uint16 SQL type that can retrieve NULL value
type Uint16 struct {
	realValue uint16
	isValid   bool
}

// NewUint16 creates a new nullable 16-bit unsigned integer
func NewUint16(value *uint16) Uint16 {
	if value == nil {
		return Uint16{
			realValue: 0,
			isValid:   false,
		}
	}
	return Uint16{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or 16-bit unsigned integer
func (n Uint16) Get() *uint16 {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or 16-bit unsigned integer
func (n *Uint16) Set(value *uint16) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Uint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Uint16) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed uint16
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Uint16) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}

	var scanned string
	if err := convertAssign(&scanned, value); err != nil {
		return err
	}

	radix := 10
	if len(scanned) == 16 {
		radix = 2
	}

	parsed, err := strconv.ParseUint(scanned, radix, 16)
	if err != nil {
		return err
	}
	n.realValue = uint16(parsed)

	n.isValid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n Uint16) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return strconv.FormatUint(uint64(n.realValue), 10), nil
}

// GormValue implements the driver Valuer interface via GORM.
func (n Uint16) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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
		leadingZero := 16 - len(value)
		for leadingZero > 0 {
			value = "0" + value
			leadingZero--
		}
		return clause.Expr{SQL: "?", Vars: []interface{}{value}}
	}
	return clause.Expr{}
}

// GormDataType gorm common data type
func (Uint16) GormDataType() string {
	return "uint16_null"
}

// GormDBDataType gorm db data type
func (Uint16) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "SMALLINT UNSIGNED"
	case "postgres":
		return "bit(16)"
	}
	return ""
}
