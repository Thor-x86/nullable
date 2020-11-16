package nullable_test

import (
	"encoding/json"
	"testing"

	"github.com/Thor-x86/nullable"
	"gorm.io/gorm/utils/tests"
)

func marshalUnmarshalJSON(t *testing.T, target interface{}) {
	serialized, err := json.Marshal(target)
	if err != nil {
		t.Fatalf("Failed to marshal %T", target)
		return
	}

	switch target.(type) {
	case nullable.Bool:
		var unserialized nullable.Bool
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Byte:
		var unserialized nullable.Byte
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Bytes:
		var unserialized nullable.Bytes
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Float32:
		var unserialized nullable.Float32
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Float64:
		var unserialized nullable.Float64
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Int:
		var unserialized nullable.Int
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Int8:
		var unserialized nullable.Int8
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Int16:
		var unserialized nullable.Int16
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Int32:
		var unserialized nullable.Int32
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Int64:
		var unserialized nullable.Int64
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.String:
		var unserialized nullable.String
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Time:
		var unserialized nullable.Time
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Uint:
		var unserialized nullable.Uint
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Uint8:
		var unserialized nullable.Uint8
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Uint16:
		var unserialized nullable.Uint16
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Uint32:
		var unserialized nullable.Uint32
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.Uint64:
		var unserialized nullable.Uint64
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	default:
		t.Fatalf("%T is not registered at json_test.go", target)
	}
}
