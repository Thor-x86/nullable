package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Int SQL type that can retrieve NULL value
type Int struct {
	realValue int
	isValid   bool
}

// NewInt creates a new nullable integer
func NewInt(value *int) Int {
	if value == nil {
		return Int{
			realValue: 0,
			isValid:   false,
		}
	}
	return Int{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or integer
func (n Int) Get() *int {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or integer
func (n *Int) Set(value *int) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Int) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed int
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Int) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}
	n.isValid = true
	return convertAssign(&n.realValue, value)
}

// Value implements the driver Valuer interface.
func (n Int) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return int64(n.realValue), nil
}

// GormDataType gorm common data type
func (Int) GormDataType() string {
	return "int_null"
}

// GormDBDataType gorm db data type
func (Int) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "BIGINT"
	case "postgres":
		return "bigint"
	}
	return ""
}
