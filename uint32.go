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

// Uint32 SQL type that can retrieve NULL value
type Uint32 struct {
	realValue uint32
	isValid   bool
}

// NewUint32 creates a new nullable 32-bit unsigned integer
func NewUint32(value *uint32) Uint32 {
	if value == nil {
		return Uint32{
			realValue: 0,
			isValid:   false,
		}
	}
	return Uint32{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or 32-bit unsigned integer
func (n Uint32) Get() *uint32 {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or 32-bit unsigned integer
func (n *Uint32) Set(value *uint32) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Uint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Uint32) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed uint32
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Uint32) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}

	var scanned string
	if err := convertAssign(&scanned, value); err != nil {
		return err
	}

	radix := 10
	if len(scanned) == 32 {
		radix = 2
	}

	parsed, err := strconv.ParseUint(scanned, radix, 32)
	if err != nil {
		return err
	}
	n.realValue = uint32(parsed)

	n.isValid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n Uint32) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return strconv.FormatUint(uint64(n.realValue), 10), nil
}

// GormValue implements the driver Valuer interface via GORM.
func (n Uint32) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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
		return clause.Expr{SQL: "?", Vars: []interface{}{n.realValue}}
	}
	return clause.Expr{}
}

// GormDataType gorm common data type
func (Uint32) GormDataType() string {
	return "uint32_null"
}

// GormDBDataType gorm db data type
func (Uint32) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "INT UNSIGNED"
	case "postgres":
		return "numeric"
	}
	return ""
}
