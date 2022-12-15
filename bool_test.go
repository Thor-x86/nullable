package nullable_test

import (
	"testing"

	"gorm.io/gorm/utils/tests"

	"github.com/tee8z/nullable"
)

func TestScanBool(t *testing.T) {
	nullableBool := nullable.NewBool(nil)

	nullableBool.Scan(true)
	tests.AssertEqual(t, nullableBool.Get(), true)

	nullableBool.Scan(false)
	tests.AssertEqual(t, nullableBool.Get(), false)

	nullableBool.Scan(nil)
	tests.AssertEqual(t, nullableBool.Get(), nil)
}

func TestNewBool(t *testing.T) {
	basicBool1 := true
	nullableBool1 := nullable.NewBool(&basicBool1)
	tests.AssertEqual(t, nullableBool1.Get(), true)

	basicBool2 := false
	nullableBool2 := nullable.NewBool(&basicBool2)
	tests.AssertEqual(t, nullableBool2.Get(), false)

	nullableBool3 := nullable.NewBool(nil)
	tests.AssertEqual(t, nullableBool3.Get(), nil)
}

func TestSetBool(t *testing.T) {
	nullableBool := nullable.NewBool(nil)
	tests.AssertEqual(t, nullableBool.Get(), nil)

	trueBool := true
	nullableBool.Set(&trueBool)
	tests.AssertEqual(t, nullableBool.Get(), true)

	falseBool := false
	nullableBool.Set(&falseBool)
	tests.AssertEqual(t, nullableBool.Get(), false)

	nullableBool.Set(nil)
	tests.AssertEqual(t, nullableBool.Get(), nil)
}

func TestJSONBool(t *testing.T) {
	trueBool := true
	marshalUnmarshalJSON(t, nullable.NewBool(&trueBool))

	falseBool := false
	marshalUnmarshalJSON(t, nullable.NewBool(&falseBool))

	marshalUnmarshalJSON(t, nullable.NewBool(nil))
}

func TestBool(t *testing.T) {
	type TestNullableBool struct {
		ID      uint
		Name    string
		Checked nullable.Bool
	}

	DB.Migrator().DropTable(&TestNullableBool{})
	if err := DB.Migrator().AutoMigrate(&TestNullableBool{}); err != nil {
		t.Errorf("failed to migrate nullable boolean, got error: %v", err)
	}

	checked := true
	checkedUser := TestNullableBool{
		Name:    "Athaariq Ardhiansyah",
		Checked: nullable.NewBool(&checked),
	}
	DB.Create(&checkedUser)

	unchecked := false
	uncheckedUser := TestNullableBool{
		Name:    "Bad Cat",
		Checked: nullable.NewBool(&unchecked),
	}
	DB.Create(&uncheckedUser)

	unknownUser := TestNullableBool{
		Name:    "Charles Robinson",
		Checked: nullable.NewBool(nil),
	}
	DB.Create(&unknownUser)

	var result1 TestNullableBool
	if err := DB.First(&result1, "name = ?", "Athaariq Ardhiansyah").Error; err != nil {
		t.Fatal("Cannot read boolean test record of \"Athaariq Ardhiansyah\"")
	}
	tests.AssertEqual(t, result1, checkedUser)

	var result2 TestNullableBool
	if err := DB.First(&result2, "name = ?", "Bad Cat").Error; err != nil {
		t.Fatal("Cannot read boolean test record of \"Bad Cat\"")
	}
	tests.AssertEqual(t, result2, uncheckedUser)

	var result3 TestNullableBool
	if err := DB.First(&result3, "name = ?", "Charles Robinson").Error; err != nil {
		t.Fatal("Cannot read boolean test record of \"Charles Robinson\"")
	}
	tests.AssertEqual(t, result3, unknownUser)
}
