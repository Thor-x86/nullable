package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Byte SQL type that can retrieve NULL value
type Byte struct {
	realValue byte
	isValid   bool
}

// NewByte creates a new nullable single byte
func NewByte(value *byte) Byte {
	if value == nil {
		return Byte{
			realValue: 0,
			isValid:   false,
		}
	}
	return Byte{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or single byte
func (n Byte) Get() *byte {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or single byte
func (n *Byte) Set(value *byte) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Byte) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Byte) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed byte
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Byte) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}

	switch value.(type) {
	case int:
		n.realValue = byte(value.(int))
	default:
		var buffer []byte
		if err := convertAssign(&buffer, value); err != nil {
			return err
		}
		n.realValue = buffer[0]
	}

	n.isValid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n Byte) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return []byte{
		byte(n.realValue),
	}, nil
}

// GormDataType gorm common data type
func (Byte) GormDataType() string {
	return "byte_null"
}

// GormDBDataType gorm db data type
func (Byte) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "TINYINT UNSIGNED"
	case "mysql":
		return "BINARY"
	case "postgres":
		return "bytea"
	}
	return ""
}
