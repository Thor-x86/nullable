package nullable_test

import (
	"testing"

	"github.com/tee8z/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanUint32(t *testing.T) {
	nullableInt := nullable.NewUint32(nil)

	// uint8
	nullableInt.Scan(37)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// uint16
	nullableInt.Scan(1234)
	tests.AssertEqual(t, nullableInt.Get(), 1234)

	// uint32
	nullableInt.Scan(654321)
	tests.AssertEqual(t, nullableInt.Get(), 654321)

	nullableInt.Scan(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestNewUint32(t *testing.T) {
	// uint8
	var basicUint1 uint32 = 37
	nullableUint1 := nullable.NewUint32(&basicUint1)
	tests.AssertEqual(t, nullableUint1.Get(), 37)

	// uint16
	var basicUint2 uint32 = 1234
	nullableUint2 := nullable.NewUint32(&basicUint2)
	tests.AssertEqual(t, nullableUint2.Get(), 1234)

	// uint32
	var basicUint3 uint32 = 654321
	nullableUint3 := nullable.NewUint32(&basicUint3)
	tests.AssertEqual(t, nullableUint3.Get(), 654321)

	nullableUint5 := nullable.NewUint32(nil)
	tests.AssertEqual(t, nullableUint5.Get(), nil)
}

func TestSetUint32(t *testing.T) {
	nullableUint := nullable.NewUint32(nil)
	tests.AssertEqual(t, nullableUint.Get(), nil)

	// uint8
	var basicUint1 uint32 = 37
	nullableUint.Set(&basicUint1)
	tests.AssertEqual(t, nullableUint.Get(), 37)
	// uint16
	var basicUint2 uint32 = 1234
	nullableUint.Set(&basicUint2)
	tests.AssertEqual(t, nullableUint.Get(), 1234)

	// uint32
	var basicUint3 uint32 = 654321
	nullableUint.Set(&basicUint3)
	tests.AssertEqual(t, nullableUint.Get(), 654321)

	nullableUint.Set(nil)
	tests.AssertEqual(t, nullableUint.Get(), nil)
}

func TestJSONUint32(t *testing.T) {
	var basicInt1 uint32 = 37
	marshalUnmarshalJSON(t, nullable.NewUint32(&basicInt1))

	var basicInt2 uint32 = 1234
	marshalUnmarshalJSON(t, nullable.NewUint32(&basicInt2))

	var basicInt3 uint32 = 654321
	marshalUnmarshalJSON(t, nullable.NewUint32(&basicInt3))

	marshalUnmarshalJSON(t, nullable.NewUint32(nil))
}

func TestUint32(t *testing.T) {
	type TestNullableUint32 struct {
		ID    uint32
		Name  string
		Value nullable.Uint32
		Unit  string
	}

	DB.Migrator().DropTable(&TestNullableUint32{})
	if err := DB.Migrator().AutoMigrate(&TestNullableUint32{}); err != nil {
		t.Errorf("failed to migrate nullable uint32, got error: %v", err)
	}

	var protonEnergy uint32 = 654321
	proton := TestNullableUint32{
		Name:  "proton",
		Value: nullable.NewUint32(&protonEnergy),
		Unit:  "Joule",
	}
	DB.Create(&proton)

	neutron := TestNullableUint32{
		Name:  "neutron",
		Value: nullable.NewUint32(nil),
		Unit:  "Joule",
	}
	DB.Create(&neutron)

	var result1 TestNullableUint32
	if err := DB.First(&result1, "name = ?", "proton").Error; err != nil {
		t.Fatal("Cannot read uint32 test record of \"proton\"")
	}
	tests.AssertEqual(t, result1, proton)

	var result2 TestNullableUint32
	if err := DB.First(&result2, "name = ?", "neutron").Error; err != nil {
		t.Fatal("Cannot read uint32 test record of \"neutron\"")
	}
	tests.AssertEqual(t, result2, neutron)
}
