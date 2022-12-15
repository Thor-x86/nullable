package nullable_test

import (
	"testing"

	"github.com/tee8z/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanFloat32(t *testing.T) {
	nullableFloat := nullable.NewFloat32(nil)

	nullableFloat.Scan(24.78)
	tests.AssertEqual(t, nullableFloat.Get(), 24.780000686645508)

	nullableFloat.Scan(-24.78)
	tests.AssertEqual(t, nullableFloat.Get(), -24.780000686645508)

	nullableFloat.Scan(782.873129836256643728346128238420)
	tests.AssertEqual(t, nullableFloat.Get(), 782.8731079101562)

	nullableFloat.Scan(-782.873129836256643728346128238420)
	tests.AssertEqual(t, nullableFloat.Get(), -782.8731079101562)

	nullableFloat.Scan(nil)
	tests.AssertEqual(t, nullableFloat.Get(), nil)
}

func TestNewFloat32(t *testing.T) {
	// Check if follow IEEE754 rules
	var basicFloat1 float32 = 24.78
	nullableFloat1 := nullable.NewFloat32(&basicFloat1)
	tests.AssertEqual(t, nullableFloat1.Get(), 24.780000686645508)

	// Check if follow IEEE754 rules for negative value
	var basicFloat2 float32 = -24.78
	nullableFloat2 := nullable.NewFloat32(&basicFloat2)
	tests.AssertEqual(t, nullableFloat2.Get(), -24.780000686645508)

	var basicFloat3 float32 = 782.873129836256643728346128238420
	nullableFloat3 := nullable.NewFloat32(&basicFloat3)
	tests.AssertEqual(t, nullableFloat3.Get(), 782.8731079101562)

	var basicFloat4 float32 = -782.873129836256643728346128238420
	nullableFloat4 := nullable.NewFloat32(&basicFloat4)
	tests.AssertEqual(t, nullableFloat4.Get(), -782.8731079101562)

	nullableFloat5 := nullable.NewFloat32(nil)
	tests.AssertEqual(t, nullableFloat5.Get(), nil)
}

func TestSetFloat32(t *testing.T) {
	nullableFloat := nullable.NewFloat32(nil)
	tests.AssertEqual(t, nullableFloat.Get(), nil)

	var basicFloat1 float32 = 24.78
	nullableFloat.Set(&basicFloat1)
	tests.AssertEqual(t, nullableFloat.Get(), 24.780000686645508)

	var basicFloat2 float32 = -24.78
	nullableFloat.Set(&basicFloat2)
	tests.AssertEqual(t, nullableFloat.Get(), -24.780000686645508)

	var basicFloat3 float32 = 782.873129836256643728346128238420
	nullableFloat.Set(&basicFloat3)
	tests.AssertEqual(t, nullableFloat.Get(), 782.8731079101562)

	var basicFloat4 float32 = -782.873129836256643728346128238420
	nullableFloat.Set(&basicFloat4)
	tests.AssertEqual(t, nullableFloat.Get(), -782.8731079101562)

	nullableFloat.Set(nil)
	tests.AssertEqual(t, nullableFloat.Get(), nil)
}

func TestJSONFloat32(t *testing.T) {
	var basicFloat1 float32 = 24.78
	marshalUnmarshalJSON(t, nullable.NewFloat32(&basicFloat1))

	var basicFloat2 float32 = -24.78
	marshalUnmarshalJSON(t, nullable.NewFloat32(&basicFloat2))

	var basicFloat3 float32 = 782.873129836256643728346128238420
	marshalUnmarshalJSON(t, nullable.NewFloat32(&basicFloat3))

	var basicFloat4 float32 = -782.873129836256643728346128238420
	marshalUnmarshalJSON(t, nullable.NewFloat32(&basicFloat4))

	marshalUnmarshalJSON(t, nullable.NewFloat32(nil))
}

func TestFloat32(t *testing.T) {
	type TestNullableFloat32 struct {
		ID        uint
		Name      string
		Latitude  nullable.Float32
		Longitude nullable.Float32
	}

	DB.Migrator().DropTable(&TestNullableFloat32{})
	if err := DB.Migrator().AutoMigrate(&TestNullableFloat32{}); err != nil {
		t.Errorf("failed to migrate nullable float32, got error: %v", err)
	}

	var lat1 float32 = 24.78
	var long1 float32 = -282.873129836256643728346128238420
	user1 := TestNullableFloat32{
		Name:      "Athaariq Ardhiansyah",
		Latitude:  nullable.NewFloat32(&lat1),
		Longitude: nullable.NewFloat32(&long1),
	}
	DB.Create(&user1)

	var lat2 float32 = 282.873129836256643728346128238420
	var long2 float32 = -24.78
	user2 := TestNullableFloat32{
		Name:      "Bad Cat",
		Latitude:  nullable.NewFloat32(&lat2),
		Longitude: nullable.NewFloat32(&long2),
	}
	DB.Create(&user2)

	user3 := TestNullableFloat32{
		Name:      "Charles Robinson",
		Latitude:  nullable.NewFloat32(nil),
		Longitude: nullable.NewFloat32(nil),
	}
	DB.Create(&user3)

	var result1 TestNullableFloat32
	if err := DB.First(&result1, "name = ?", "Athaariq Ardhiansyah").Error; err != nil {
		t.Fatal("Cannot read float32 test record of \"Athaariq Ardhiansyah\"")
	}
	tests.AssertEqual(t, result1, user1)

	var result2 TestNullableFloat32
	if err := DB.First(&result2, "name = ?", "Bad Cat").Error; err != nil {
		t.Fatal("Cannot read float32 test record of \"Bad Cat\"")
	}
	tests.AssertEqual(t, result2, user2)

	var result3 TestNullableFloat32
	if err := DB.First(&result3, "name = ?", "Charles Robinson").Error; err != nil {
		t.Fatal("Cannot read float32 test record of \"Charles Robinson\"")
	}
	tests.AssertEqual(t, result3, user3)
}
