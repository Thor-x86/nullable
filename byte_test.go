package nullable_test

import (
	"testing"

	"github.com/tee8z/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanByte(t *testing.T) {
	nullableByte := nullable.NewByte(nil)

	nullableByte.Scan(0)
	tests.AssertEqual(t, nullableByte.Get(), 0)

	nullableByte.Scan(0x7f)
	tests.AssertEqual(t, nullableByte.Get(), 0x7f)

	nullableByte.Scan(0xff)
	tests.AssertEqual(t, nullableByte.Get(), 0xff)

	nullableByte.Scan(nil)
	tests.AssertEqual(t, nullableByte.Get(), nil)
}

func TestNewByte(t *testing.T) {
	basicByte1 := byte(0)
	nullableByte1 := nullable.NewByte(&basicByte1)
	tests.AssertEqual(t, nullableByte1.Get(), 0)

	basicByte2 := byte(0x7f)
	nullableByte2 := nullable.NewByte(&basicByte2)
	tests.AssertEqual(t, nullableByte2.Get(), 0x7f)

	basicByte3 := byte(0xff)
	nullableByte3 := nullable.NewByte(&basicByte3)
	tests.AssertEqual(t, nullableByte3.Get(), 0xff)

	nullableByte4 := nullable.NewByte(nil)
	tests.AssertEqual(t, nullableByte4.Get(), nil)
}

func TestSetByte(t *testing.T) {
	nullableByte := nullable.NewByte(nil)
	tests.AssertEqual(t, nullableByte.Get(), nil)

	basicByte1 := byte(0)
	nullableByte.Set(&basicByte1)
	tests.AssertEqual(t, nullableByte.Get(), 0)

	basicByte2 := byte(0x7f)
	nullableByte.Set(&basicByte2)
	tests.AssertEqual(t, nullableByte.Get(), 0x7f)

	basicByte3 := byte(0xff)
	nullableByte.Set(&basicByte3)
	tests.AssertEqual(t, nullableByte.Get(), 0xff)

	nullableByte.Set(nil)
	tests.AssertEqual(t, nullableByte.Get(), nil)
}

func TestJSONByte(t *testing.T) {
	basicByte1 := byte(0)
	marshalUnmarshalJSON(t, nullable.NewByte(&basicByte1))

	basicByte2 := byte(0x7f)
	marshalUnmarshalJSON(t, nullable.NewByte(&basicByte2))

	basicByte3 := byte(0xff)
	marshalUnmarshalJSON(t, nullable.NewByte(&basicByte3))

	marshalUnmarshalJSON(t, nullable.NewByte(nil))
}

func TestByte(t *testing.T) {
	type TestNullableByte struct {
		ID   uint
		Name string
		Flag nullable.Byte
	}

	DB.Migrator().DropTable(&TestNullableByte{})
	if err := DB.Migrator().AutoMigrate(&TestNullableByte{}); err != nil {
		t.Errorf("failed to migrate nullable byte, got error: %v", err)
	}

	flagged := byte(0x3c)
	flaggedUser := TestNullableByte{
		Name: "Athaariq Ardhiansyah",
		Flag: nullable.NewByte(&flagged),
	}
	DB.Create(&flaggedUser)

	unflagged := byte(0x0)
	unflaggedUser := TestNullableByte{
		Name: "Bad Cat",
		Flag: nullable.NewByte(&unflagged),
	}
	DB.Create(&unflaggedUser)

	unknownUser := TestNullableByte{
		Name: "Charles Robinson",
		Flag: nullable.NewByte(nil),
	}
	DB.Create(&unknownUser)

	var result1 TestNullableByte
	if err := DB.First(&result1, "name = ?", "Athaariq Ardhiansyah").Error; err != nil {
		t.Fatal("Cannot read byte test record of \"Athaariq Ardhiansyah\"")
	}
	tests.AssertEqual(t, result1, flaggedUser)

	var result2 TestNullableByte
	if err := DB.First(&result2, "name = ?", "Bad Cat").Error; err != nil {
		t.Fatal("Cannot read byte test record of \"Bad Cat\"")
	}
	tests.AssertEqual(t, result2, unflaggedUser)

	var result3 TestNullableByte
	if err := DB.First(&result3, "name = ?", "Charles Robinson").Error; err != nil {
		t.Fatal("Cannot read byte test record of \"Charles Robinson\"")
	}
	tests.AssertEqual(t, result3, unknownUser)
}
