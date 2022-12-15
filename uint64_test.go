package nullable_test

import (
	"testing"

	"github.com/tee8z/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanUint64(t *testing.T) {
	nullableInt := nullable.NewUint64(nil)

	// uint8
	nullableInt.Scan(37)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// raw binary of uint8
	nullableInt.Scan("0000000000000000000000000000000000000000000000000000000000010001")
	tests.AssertEqual(t, nullableInt.Get(), 17)

	// uint16
	nullableInt.Scan(1234)
	tests.AssertEqual(t, nullableInt.Get(), 1234)

	// raw binary of uint16
	nullableInt.Scan("0000000000000000000000000000000000000000000000000001000011100001")
	tests.AssertEqual(t, nullableInt.Get(), 4321)

	// uint32
	nullableInt.Scan(654321)
	tests.AssertEqual(t, nullableInt.Get(), 654321)

	// raw binary of uint32
	nullableInt.Scan("0000000000000000000000000000000000000000000001011010101000100010")
	tests.AssertEqual(t, nullableInt.Get(), 371234)

	// uint64
	nullableInt.Scan(50000000000)
	tests.AssertEqual(t, nullableInt.Get(), 50000000000)

	// raw binary of uint64
	nullableInt.Scan("0000000000000000000000000000000000111010110111100110100010110001")
	tests.AssertEqual(t, nullableInt.Get(), 987654321)

	nullableInt.Scan(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestNewUint64(t *testing.T) {
	// uint8
	var basicUint1 uint64 = 37
	nullableUint1 := nullable.NewUint64(&basicUint1)
	tests.AssertEqual(t, nullableUint1.Get(), 37)

	// uint16
	var basicUint2 uint64 = 1234
	nullableUint2 := nullable.NewUint64(&basicUint2)
	tests.AssertEqual(t, nullableUint2.Get(), 1234)

	// uint32
	var basicUint3 uint64 = 654321
	nullableUint3 := nullable.NewUint64(&basicUint3)
	tests.AssertEqual(t, nullableUint3.Get(), 654321)

	// uint64
	var basicUint4 uint64 = 50000000000
	nullableUint4 := nullable.NewUint64(&basicUint4)
	tests.AssertEqual(t, nullableUint4.Get(), 50000000000)

	nullableUint5 := nullable.NewUint64(nil)
	tests.AssertEqual(t, nullableUint5.Get(), nil)
}

func TestSetUint64(t *testing.T) {
	nullableUint := nullable.NewUint64(nil)
	tests.AssertEqual(t, nullableUint.Get(), nil)

	// uint8
	var basicUint1 uint64 = 37
	nullableUint.Set(&basicUint1)
	tests.AssertEqual(t, nullableUint.Get(), 37)
	// uint16
	var basicUint2 uint64 = 1234
	nullableUint.Set(&basicUint2)
	tests.AssertEqual(t, nullableUint.Get(), 1234)

	// uint32
	var basicUint3 uint64 = 654321
	nullableUint.Set(&basicUint3)
	tests.AssertEqual(t, nullableUint.Get(), 654321)

	// uint64
	var basicUint4 uint64 = 50000000000
	nullableUint.Set(&basicUint4)
	tests.AssertEqual(t, nullableUint.Get(), 50000000000)

	nullableUint.Set(nil)
	tests.AssertEqual(t, nullableUint.Get(), nil)
}

func TestJSONUint64(t *testing.T) {
	var basicInt1 uint64 = 37
	marshalUnmarshalJSON(t, nullable.NewUint64(&basicInt1))

	var basicInt2 uint64 = 1234
	marshalUnmarshalJSON(t, nullable.NewUint64(&basicInt2))

	var basicInt3 uint64 = 654321
	marshalUnmarshalJSON(t, nullable.NewUint64(&basicInt3))

	var basicInt4 uint64 = 50000000000
	marshalUnmarshalJSON(t, nullable.NewUint64(&basicInt4))

	marshalUnmarshalJSON(t, nullable.NewUint64(nil))
}

func TestUint64(t *testing.T) {
	type TestNullableUint64 struct {
		ID    uint64
		Name  string
		Value nullable.Uint64
		Unit  string
	}

	DB.Migrator().DropTable(&TestNullableUint64{})
	if err := DB.Migrator().AutoMigrate(&TestNullableUint64{}); err != nil {
		t.Errorf("failed to migrate nullable uint64, got error: %v", err)
	}

	var protonEnergy uint64 = 50000000000
	proton := TestNullableUint64{
		Name:  "proton",
		Value: nullable.NewUint64(&protonEnergy),
		Unit:  "Joule",
	}
	DB.Create(&proton)

	neutron := TestNullableUint64{
		Name:  "neutron",
		Value: nullable.NewUint64(nil),
		Unit:  "Joule",
	}
	DB.Create(&neutron)

	var result1 TestNullableUint64
	if err := DB.First(&result1, "name = ?", "proton").Error; err != nil {
		t.Fatal("Cannot read uint64 test record of \"proton\"")
	}
	tests.AssertEqual(t, result1, proton)

	var result2 TestNullableUint64
	if err := DB.First(&result2, "name = ?", "neutron").Error; err != nil {
		t.Fatal("Cannot read uint64 test record of \"neutron\"")
	}
	tests.AssertEqual(t, result2, neutron)
}
