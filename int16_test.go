package nullable_test

import (
	"testing"

	"github.com/tee8z/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestJSONInt16(t *testing.T) {
	var basicInt1 int16 = 37
	marshalUnmarshalJSON(t, nullable.NewInt16(&basicInt1))

	var basicInt2 int16 = -37
	marshalUnmarshalJSON(t, nullable.NewInt16(&basicInt2))

	var basicInt3 int16 = 1234
	marshalUnmarshalJSON(t, nullable.NewInt16(&basicInt3))

	var basicInt4 int16 = -1234
	marshalUnmarshalJSON(t, nullable.NewInt16(&basicInt4))

	marshalUnmarshalJSON(t, nullable.NewInt16(nil))
}

func TestNewInt16(t *testing.T) {
	// uint8
	var basicInt1 int16 = 37
	nullableInt1 := nullable.NewInt16(&basicInt1)
	tests.AssertEqual(t, nullableInt1.Get(), 37)

	// int8
	var basicInt2 int16 = -37
	nullableInt2 := nullable.NewInt16(&basicInt2)
	tests.AssertEqual(t, nullableInt2.Get(), -37)

	// uint16
	var basicInt3 int16 = 1234
	nullableInt3 := nullable.NewInt16(&basicInt3)
	tests.AssertEqual(t, nullableInt3.Get(), 1234)

	// int16
	var basicInt4 int16 = -1234
	nullableInt4 := nullable.NewInt16(&basicInt4)
	tests.AssertEqual(t, nullableInt4.Get(), -1234)
}

func TestSetInt16(t *testing.T) {
	nullableInt := nullable.NewInt16(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)

	// uint8
	var basicInt1 int16 = 37
	nullableInt.Set(&basicInt1)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// int8
	var basicInt2 int16 = -37
	nullableInt.Set(&basicInt2)
	tests.AssertEqual(t, nullableInt.Get(), -37)

	// uint16
	var basicInt3 int16 = 1234
	nullableInt.Set(&basicInt3)
	tests.AssertEqual(t, nullableInt.Get(), 1234)

	// int16
	var basicInt4 int16 = -1234
	nullableInt.Set(&basicInt4)
	tests.AssertEqual(t, nullableInt.Get(), -1234)

	nullableInt.Set(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestInt16(t *testing.T) {
	type TestNullableInt16 struct {
		ID    uint
		Name  string
		Value nullable.Int16
		Unit  string
	}

	DB.Migrator().DropTable(&TestNullableInt16{})
	if err := DB.Migrator().AutoMigrate(&TestNullableInt16{}); err != nil {
		t.Errorf("failed to migrate nullable int16, got error: %v", err)
	}

	var matterEnergy int16 = 1234
	matter := TestNullableInt16{
		Name:  "matter",
		Value: nullable.NewInt16(&matterEnergy),
		Unit:  "Joule",
	}
	DB.Create(&matter)

	var antimatterEnergy int16 = -1234
	antimatter := TestNullableInt16{
		Name:  "antimatter",
		Value: nullable.NewInt16(&antimatterEnergy),
		Unit:  "Joule",
	}
	DB.Create(&antimatter)

	neutron := TestNullableInt16{
		Name:  "neutron",
		Value: nullable.NewInt16(nil),
		Unit:  "Joule",
	}
	DB.Create(&neutron)

	var result1 TestNullableInt16
	if err := DB.First(&result1, "name = ?", "matter").Error; err != nil {
		t.Fatal("Cannot read int16 test record of \"matter\"")
	}
	tests.AssertEqual(t, result1, matter)

	var result2 TestNullableInt16
	if err := DB.First(&result2, "name = ?", "antimatter").Error; err != nil {
		t.Fatal("Cannot read int16 test record of \"antimatter\"")
	}
	tests.AssertEqual(t, result2, antimatter)

	var result3 TestNullableInt16
	if err := DB.First(&result3, "name = ?", "neutron").Error; err != nil {
		t.Fatal("Cannot read int16 test record of \"neutron\"")
	}
	tests.AssertEqual(t, result3, neutron)
}
