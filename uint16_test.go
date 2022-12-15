package nullable_test

import (
	"testing"

	"github.com/tee8z/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanUint16(t *testing.T) {
	nullableInt := nullable.NewUint16(nil)

	// uint8
	nullableInt.Scan(37)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// raw binary of uint8
	nullableInt.Scan("0000000000010001")
	tests.AssertEqual(t, nullableInt.Get(), 17)

	// uint16
	nullableInt.Scan(1234)
	tests.AssertEqual(t, nullableInt.Get(), 1234)

	// raw binary of uint16
	nullableInt.Scan("0001000011100001")
	tests.AssertEqual(t, nullableInt.Get(), 4321)

	nullableInt.Scan(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestNewUint16(t *testing.T) {
	// uint8
	var basicUint1 uint16 = 37
	nullableUint1 := nullable.NewUint16(&basicUint1)
	tests.AssertEqual(t, nullableUint1.Get(), 37)

	// uint16
	var basicUint2 uint16 = 1234
	nullableUint2 := nullable.NewUint16(&basicUint2)
	tests.AssertEqual(t, nullableUint2.Get(), 1234)

	nullableUint5 := nullable.NewUint16(nil)
	tests.AssertEqual(t, nullableUint5.Get(), nil)
}

func TestSetUint16(t *testing.T) {
	nullableUint := nullable.NewUint16(nil)
	tests.AssertEqual(t, nullableUint.Get(), nil)

	// uint8
	var basicUint1 uint16 = 37
	nullableUint.Set(&basicUint1)
	tests.AssertEqual(t, nullableUint.Get(), 37)
	// uint16
	var basicUint2 uint16 = 1234
	nullableUint.Set(&basicUint2)
	tests.AssertEqual(t, nullableUint.Get(), 1234)

	nullableUint.Set(nil)
	tests.AssertEqual(t, nullableUint.Get(), nil)
}

func TestJSONUint16(t *testing.T) {
	var basicInt1 uint16 = 37
	marshalUnmarshalJSON(t, nullable.NewUint16(&basicInt1))

	var basicInt2 uint16 = 1234
	marshalUnmarshalJSON(t, nullable.NewUint16(&basicInt2))

	marshalUnmarshalJSON(t, nullable.NewUint16(nil))
}

func TestUint16(t *testing.T) {
	type TestNullableUint16 struct {
		ID    uint16
		Name  string
		Value nullable.Uint16
		Unit  string
	}

	DB.Migrator().DropTable(&TestNullableUint16{})
	if err := DB.Migrator().AutoMigrate(&TestNullableUint16{}); err != nil {
		t.Errorf("failed to migrate nullable uint16, got error: %v", err)
	}

	var protonEnergy uint16 = 16321
	proton := TestNullableUint16{
		Name:  "proton",
		Value: nullable.NewUint16(&protonEnergy),
		Unit:  "Joule",
	}
	DB.Create(&proton)

	neutron := TestNullableUint16{
		Name:  "neutron",
		Value: nullable.NewUint16(nil),
		Unit:  "Joule",
	}
	DB.Create(&neutron)

	var result1 TestNullableUint16
	if err := DB.First(&result1, "name = ?", "proton").Error; err != nil {
		t.Fatal("Cannot read uint16 test record of \"proton\"")
	}
	tests.AssertEqual(t, result1, proton)

	var result2 TestNullableUint16
	if err := DB.First(&result2, "name = ?", "neutron").Error; err != nil {
		t.Fatal("Cannot read uint16 test record of \"neutron\"")
	}
	tests.AssertEqual(t, result2, neutron)
}
