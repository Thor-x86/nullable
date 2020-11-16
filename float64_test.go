package nullable_test

import (
	"testing"

	"github.com/Thor-x86/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanFloat64(t *testing.T) {
	nullableFloat := nullable.NewFloat64(nil)

	nullableFloat.Scan(24.78)
	tests.AssertEqual(t, nullableFloat.Get(), 24.78)

	nullableFloat.Scan(-24.78)
	tests.AssertEqual(t, nullableFloat.Get(), -24.78)

	nullableFloat.Scan(782.873129836256643728346128238420)
	tests.AssertEqual(t, nullableFloat.Get(), 782.8731298362567)

	nullableFloat.Scan(-782.873129836256643728346128238420)
	tests.AssertEqual(t, nullableFloat.Get(), -782.8731298362567)

	nullableFloat.Scan(nil)
	tests.AssertEqual(t, nullableFloat.Get(), nil)
}

func TestNewFloat64(t *testing.T) {
	// Check if follow IEEE754 rules
	var basicFloat1 float64 = 24.78
	nullableFloat1 := nullable.NewFloat64(&basicFloat1)
	tests.AssertEqual(t, nullableFloat1.Get(), 24.78)

	// Check if follow IEEE754 rules for negative value
	var basicFloat2 float64 = -24.78
	nullableFloat2 := nullable.NewFloat64(&basicFloat2)
	tests.AssertEqual(t, nullableFloat2.Get(), -24.78)

	var basicFloat3 float64 = 782.873129836256643728346128238420
	nullableFloat3 := nullable.NewFloat64(&basicFloat3)
	tests.AssertEqual(t, nullableFloat3.Get(), 782.8731298362567)

	var basicFloat4 float64 = -782.873129836256643728346128238420
	nullableFloat4 := nullable.NewFloat64(&basicFloat4)
	tests.AssertEqual(t, nullableFloat4.Get(), -782.8731298362567)

	nullableFloat5 := nullable.NewFloat64(nil)
	tests.AssertEqual(t, nullableFloat5.Get(), nil)
}

func TestSetFloat64(t *testing.T) {
	nullableFloat := nullable.NewFloat64(nil)
	tests.AssertEqual(t, nullableFloat.Get(), nil)

	var basicFloat1 float64 = 24.78
	nullableFloat.Set(&basicFloat1)
	tests.AssertEqual(t, nullableFloat.Get(), 24.78)

	var basicFloat2 float64 = -24.78
	nullableFloat.Set(&basicFloat2)
	tests.AssertEqual(t, nullableFloat.Get(), -24.78)

	var basicFloat3 float64 = 782.873129836256643728346128238420
	nullableFloat.Set(&basicFloat3)
	tests.AssertEqual(t, nullableFloat.Get(), 782.8731298362567)

	var basicFloat4 float64 = -782.873129836256643728346128238420
	nullableFloat.Set(&basicFloat4)
	tests.AssertEqual(t, nullableFloat.Get(), -782.8731298362567)

	nullableFloat.Set(nil)
	tests.AssertEqual(t, nullableFloat.Get(), nil)
}

func TestJSONFloat64(t *testing.T) {
	var basicFloat1 float64 = 24.78
	marshalUnmarshalJSON(t, nullable.NewFloat64(&basicFloat1))

	var basicFloat2 float64 = -24.78
	marshalUnmarshalJSON(t, nullable.NewFloat64(&basicFloat2))

	var basicFloat3 float64 = 782.873129836256643728346128238420
	marshalUnmarshalJSON(t, nullable.NewFloat64(&basicFloat3))

	var basicFloat4 float64 = -782.873129836256643728346128238420
	marshalUnmarshalJSON(t, nullable.NewFloat64(&basicFloat4))

	marshalUnmarshalJSON(t, nullable.NewFloat64(nil))
}

func TestFloat64(t *testing.T) {
	type TestNullableFloat64 struct {
		ID        uint
		Name      string
		Latitude  nullable.Float64
		Longitude nullable.Float64
	}

	DB.Migrator().DropTable(&TestNullableFloat64{})
	if err := DB.Migrator().AutoMigrate(&TestNullableFloat64{}); err != nil {
		t.Errorf("failed to migrate nullable float64, got error: %v", err)
	}

	var lat1 float64 = 24.78
	var long1 float64 = -282.873129836256643728346128238420
	user1 := TestNullableFloat64{
		Name:      "Athaariq Ardhiansyah",
		Latitude:  nullable.NewFloat64(&lat1),
		Longitude: nullable.NewFloat64(&long1),
	}
	DB.Create(&user1)

	var lat2 float64 = 282.873129836256643728346128238420
	var long2 float64 = -24.78
	user2 := TestNullableFloat64{
		Name:      "Bad Cat",
		Latitude:  nullable.NewFloat64(&lat2),
		Longitude: nullable.NewFloat64(&long2),
	}
	DB.Create(&user2)

	user3 := TestNullableFloat64{
		Name:      "Charles Robinson",
		Latitude:  nullable.NewFloat64(nil),
		Longitude: nullable.NewFloat64(nil),
	}
	DB.Create(&user3)

	var result1 TestNullableFloat64
	if err := DB.First(&result1, "name = ?", "Athaariq Ardhiansyah").Error; err != nil {
		t.Fatal("Cannot read float64 test record of \"Athaariq Ardhiansyah\"")
	}
	tests.AssertEqual(t, result1, user1)

	var result2 TestNullableFloat64
	if err := DB.First(&result2, "name = ?", "Bad Cat").Error; err != nil {
		t.Fatal("Cannot read float64 test record of \"Bad Cat\"")
	}
	tests.AssertEqual(t, result2, user2)

	var result3 TestNullableFloat64
	if err := DB.First(&result3, "name = ?", "Charles Robinson").Error; err != nil {
		t.Fatal("Cannot read float64 test record of \"Charles Robinson\"")
	}
	tests.AssertEqual(t, result3, user3)
}
