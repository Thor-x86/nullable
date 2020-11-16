package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Float32 SQL type that can retrieve NULL value
type Float32 struct {
	realValue float32
	isValid   bool
}

// NewFloat32 creates a new nullable float
func NewFloat32(value *float32) Float32 {
	if value == nil {
		return Float32{
			realValue: 0,
			isValid:   false,
		}
	}
	return Float32{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or float
func (n Float32) Get() *float32 {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or float
func (n *Float32) Set(value *float32) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Float32) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed float32
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Float32) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}

	var f64 float64
	if err := convertAssign(&f64, value); err != nil {
		return err
	}
	n.realValue = float32(f64)

	n.isValid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n Float32) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return float64(n.realValue), nil
}

// GormDataType gorm common data type
func (Float32) GormDataType() string {
	return "float32_null"
}

// GormDBDataType gorm db data type
func (Float32) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "FLOAT"
	case "postgres":
		return "real"
	}
	return ""
}
