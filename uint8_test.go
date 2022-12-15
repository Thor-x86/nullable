package nullable_test

import (
	"testing"

	"github.com/tee8z/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanUint8(t *testing.T) {
	nullableInt := nullable.NewUint8(nil)

	// uint8
	nullableInt.Scan(37)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// raw binary of uint8
	nullableInt.Scan("00010001")
	tests.AssertEqual(t, nullableInt.Get(), 17)

	nullableInt.Scan(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestNewUint8(t *testing.T) {
	var basicUint1 uint8 = 37
	nullableUint1 := nullable.NewUint8(&basicUint1)
	tests.AssertEqual(t, nullableUint1.Get(), 37)

	nullableUint2 := nullable.NewUint8(nil)
	tests.AssertEqual(t, nullableUint2.Get(), nil)
}

func TestSetUint8(t *testing.T) {
	nullableUint := nullable.NewUint8(nil)
	tests.AssertEqual(t, nullableUint.Get(), nil)

	var basicUint1 uint8 = 37
	nullableUint.Set(&basicUint1)
	tests.AssertEqual(t, nullableUint.Get(), 37)

	nullableUint.Set(nil)
	tests.AssertEqual(t, nullableUint.Get(), nil)
}

func TestJSONUint8(t *testing.T) {
	var basicInt1 uint8 = 37
	marshalUnmarshalJSON(t, nullable.NewUint8(&basicInt1))

	marshalUnmarshalJSON(t, nullable.NewUint8(nil))
}

func TestUint8(t *testing.T) {
	type TestNullableUint8 struct {
		ID    uint
		Name  string
		Value nullable.Uint8
		Unit  string
	}

	DB.Migrator().DropTable(&TestNullableUint8{})
	if err := DB.Migrator().AutoMigrate(&TestNullableUint8{}); err != nil {
		t.Errorf("failed to migrate nullable uint8, got error: %v", err)
	}

	var protonEnergy uint8 = 115
	proton := TestNullableUint8{
		Name:  "proton",
		Value: nullable.NewUint8(&protonEnergy),
		Unit:  "Joule",
	}
	DB.Create(&proton)

	neutron := TestNullableUint8{
		Name:  "neutron",
		Value: nullable.NewUint8(nil),
		Unit:  "Joule",
	}
	DB.Create(&neutron)

	var result1 TestNullableUint8
	if err := DB.First(&result1, "name = ?", "proton").Error; err != nil {
		t.Fatal("Cannot read uint8 test record of \"proton\"")
	}
	tests.AssertEqual(t, result1, proton)

	var result2 TestNullableUint8
	if err := DB.First(&result2, "name = ?", "neutron").Error; err != nil {
		t.Fatal("Cannot read uint8 test record of \"neutron\"")
	}
	tests.AssertEqual(t, result2, neutron)
}
