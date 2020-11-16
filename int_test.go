package nullable_test

import (
	"testing"

	"github.com/Thor-x86/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanInt(t *testing.T) {
	nullableInt := nullable.NewInt(nil)

	// uint8
	nullableInt.Scan(37)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// int8
	nullableInt.Scan(-37)
	tests.AssertEqual(t, nullableInt.Get(), -37)

	// uint16
	nullableInt.Scan(1234)
	tests.AssertEqual(t, nullableInt.Get(), 1234)

	// int16
	nullableInt.Scan(-1234)
	tests.AssertEqual(t, nullableInt.Get(), -1234)

	// uint32
	nullableInt.Scan(654321)
	tests.AssertEqual(t, nullableInt.Get(), 654321)

	// int32
	nullableInt.Scan(-654321)
	tests.AssertEqual(t, nullableInt.Get(), -654321)

	// uint64
	nullableInt.Scan(50000000000)
	tests.AssertEqual(t, nullableInt.Get(), 50000000000)

	// int64
	nullableInt.Scan(-50000000000)
	tests.AssertEqual(t, nullableInt.Get(), -50000000000)

	nullableInt.Scan(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestNewInt(t *testing.T) {
	// uint8
	var basicInt1 int = 37
	nullableInt1 := nullable.NewInt(&basicInt1)
	tests.AssertEqual(t, nullableInt1.Get(), 37)

	// int8
	var basicInt2 int = -37
	nullableInt2 := nullable.NewInt(&basicInt2)
	tests.AssertEqual(t, nullableInt2.Get(), -37)

	// uint16
	var basicInt3 int = 1234
	nullableInt3 := nullable.NewInt(&basicInt3)
	tests.AssertEqual(t, nullableInt3.Get(), 1234)

	// int16
	var basicInt4 int = -1234
	nullableInt4 := nullable.NewInt(&basicInt4)
	tests.AssertEqual(t, nullableInt4.Get(), -1234)

	// uint32
	var basicInt5 int = 654321
	nullableInt5 := nullable.NewInt(&basicInt5)
	tests.AssertEqual(t, nullableInt5.Get(), 654321)

	// int32
	var basicInt6 int = -654321
	nullableInt6 := nullable.NewInt(&basicInt6)
	tests.AssertEqual(t, nullableInt6.Get(), -654321)

	// uint64
	var basicInt7 int = 50000000000
	nullableInt7 := nullable.NewInt(&basicInt7)
	tests.AssertEqual(t, nullableInt7.Get(), 50000000000)

	// int64
	var basicInt8 int = -50000000000
	nullableInt8 := nullable.NewInt(&basicInt8)
	tests.AssertEqual(t, nullableInt8.Get(), -50000000000)
}

func TestSetInt(t *testing.T) {
	nullableInt := nullable.NewInt(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)

	// uint8
	var basicInt1 int = 37
	nullableInt.Set(&basicInt1)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// int8
	var basicInt2 int = -37
	nullableInt.Set(&basicInt2)
	tests.AssertEqual(t, nullableInt.Get(), -37)

	// uint16
	var basicInt3 int = 1234
	nullableInt.Set(&basicInt3)
	tests.AssertEqual(t, nullableInt.Get(), 1234)

	// int16
	var basicInt4 int = -1234
	nullableInt.Set(&basicInt4)
	tests.AssertEqual(t, nullableInt.Get(), -1234)

	// uint32
	var basicInt5 int = 654321
	nullableInt.Set(&basicInt5)
	tests.AssertEqual(t, nullableInt.Get(), 654321)

	// int32
	var basicInt6 int = -654321
	nullableInt.Set(&basicInt6)
	tests.AssertEqual(t, nullableInt.Get(), -654321)

	// uint64
	var basicInt7 int = 50000000000
	nullableInt.Set(&basicInt7)
	tests.AssertEqual(t, nullableInt.Get(), 50000000000)

	// int64
	var basicInt8 int = -50000000000
	nullableInt.Set(&basicInt8)
	tests.AssertEqual(t, nullableInt.Get(), -50000000000)

	nullableInt.Set(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestJSONInt(t *testing.T) {
	var basicInt1 int = 37
	marshalUnmarshalJSON(t, nullable.NewInt(&basicInt1))

	var basicInt2 int = -37
	marshalUnmarshalJSON(t, nullable.NewInt(&basicInt2))

	var basicInt3 int = 1234
	marshalUnmarshalJSON(t, nullable.NewInt(&basicInt3))

	var basicInt4 int = -1234
	marshalUnmarshalJSON(t, nullable.NewInt(&basicInt4))

	var basicInt5 int = 654321
	marshalUnmarshalJSON(t, nullable.NewInt(&basicInt5))

	var basicInt6 int = -654321
	marshalUnmarshalJSON(t, nullable.NewInt(&basicInt6))

	var basicInt7 int = 50000000000
	marshalUnmarshalJSON(t, nullable.NewInt(&basicInt7))

	var basicInt8 int = -50000000000
	marshalUnmarshalJSON(t, nullable.NewInt(&basicInt8))

	marshalUnmarshalJSON(t, nullable.NewInt(nil))
}

func TestInt(t *testing.T) {
	type TestNullableInt struct {
		ID    uint
		Name  string
		Value nullable.Int
		Unit  string
	}

	DB.Migrator().DropTable(&TestNullableInt{})
	if err := DB.Migrator().AutoMigrate(&TestNullableInt{}); err != nil {
		t.Errorf("failed to migrate nullable int, got error: %v", err)
	}

	matterEnergy := 50000000000
	matter := TestNullableInt{
		Name:  "matter",
		Value: nullable.NewInt(&matterEnergy),
		Unit:  "Joule",
	}
	DB.Create(&matter)

	antimatterEnergy := -50000000000
	antimatter := TestNullableInt{
		Name:  "antimatter",
		Value: nullable.NewInt(&antimatterEnergy),
		Unit:  "Joule",
	}
	DB.Create(&antimatter)

	neutron := TestNullableInt{
		Name:  "neutron",
		Value: nullable.NewInt(nil),
		Unit:  "Joule",
	}
	DB.Create(&neutron)

	var result1 TestNullableInt
	if err := DB.First(&result1, "name = ?", "matter").Error; err != nil {
		t.Fatal("Cannot read int test record of \"matter\"")
	}
	tests.AssertEqual(t, result1, matter)

	var result2 TestNullableInt
	if err := DB.First(&result2, "name = ?", "antimatter").Error; err != nil {
		t.Fatal("Cannot read int test record of \"antimatter\"")
	}
	tests.AssertEqual(t, result2, antimatter)

	var result3 TestNullableInt
	if err := DB.First(&result3, "name = ?", "neutron").Error; err != nil {
		t.Fatal("Cannot read int test record of \"neutron\"")
	}
	tests.AssertEqual(t, result3, neutron)
}
