package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Bytes SQL type that can retrieve NULL value
type Bytes struct {
	realValue []byte
	isValid   bool
}

// NewBytes creates a new nullable array of bytes
func NewBytes(value *[]byte) Bytes {
	if value == nil {
		return Bytes{
			realValue: []byte{},
			isValid:   false,
		}
	}
	return Bytes{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or array of bytes
func (n Bytes) Get() *[]byte {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or array of bytes
func (n *Bytes) Set(value *[]byte) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = []byte{}
	}
}

// MarshalJSON converts current value to JSON
func (n Bytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Bytes) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = []byte{}
		return nil
	}

	var parsed []byte
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Bytes) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = []byte{}, false
		return nil
	}
	n.isValid = true
	return convertAssign(&n.realValue, value)
}

// Value implements the driver Valuer interface.
func (n Bytes) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return n.realValue, nil
}

// GormDataType gorm common data type
func (Bytes) GormDataType() string {
	return "bytes_null"
}

// GormDBDataType gorm db data type
func (Bytes) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "BLOB"
	case "postgres":
		return "bytea"
	}
	return ""
}
