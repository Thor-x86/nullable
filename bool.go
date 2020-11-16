package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Bool SQL type that can retrieve NULL value
type Bool struct {
	realValue bool
	isValid   bool
}

// NewBool creates a new nullable boolean
func NewBool(value *bool) Bool {
	if value == nil {
		return Bool{
			realValue: false,
			isValid:   false,
		}
	}
	return Bool{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or boolean
func (n Bool) Get() *bool {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or boolean
func (n *Bool) Set(value *bool) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = false
	}
}

// MarshalJSON converts current value to JSON
func (n Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Bool) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = false
		return nil
	}

	var parsed bool
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Bool) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = false, false
		return nil
	}
	n.isValid = true
	return convertAssign(&n.realValue, value)
}

// Value implements the driver Valuer interface.
func (n Bool) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return n.realValue, nil
}

// GormDataType gorm common data type
func (Bool) GormDataType() string {
	return "bool_null"
}

// GormDBDataType gorm db data type
func (Bool) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "BOOLEAN"
	case "postgres":
		return "boolean"
	}
	return ""
}
