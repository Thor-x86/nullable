package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// String SQL type that can retrieve NULL value
type String struct {
	realValue string
	isValid   bool
}

// NewString creates a new nullable string
func NewString(value *string) String {
	if value == nil {
		return String{
			realValue: "",
			isValid:   false,
		}
	}
	return String{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or string
func (n String) Get() *string {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or string
func (n *String) Set(value *string) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = ""
	}
}

// MarshalJSON converts current value to JSON
func (n String) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *String) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = ""
		return nil
	}

	var parsed string
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *String) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = "", false
		return nil
	}
	n.isValid = true
	return convertAssign(&n.realValue, value)
}

// Value implements the driver Valuer interface.
func (n String) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return n.realValue, nil
}

// GormDataType gorm common data type
func (String) GormDataType() string {
	return "string_null"
}

// GormDBDataType gorm db data type
func (String) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "TEXT"
	case "postgres":
		return "text"
	}
	return ""
}
