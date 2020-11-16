package nullable_test

import (
	"testing"

	"github.com/Thor-x86/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanBytes(t *testing.T) {
	nullableBytes := nullable.NewBytes(nil)

	nullableBytes.Scan([]byte{})
	tests.AssertEqual(t, nullableBytes.Get(), []byte{})

	nullableBytes.Scan([]byte{0x0, 0x7f, 0xff})
	tests.AssertEqual(t, nullableBytes.Get(), []byte{0x0, 0x7f, 0xff})
	tests.AssertEqual(t, (*nullableBytes.Get())[0], 0x0)
	tests.AssertEqual(t, (*nullableBytes.Get())[1], 0x7f)
	tests.AssertEqual(t, (*nullableBytes.Get())[2], 0xff)

	nullableBytes.Scan(nil)
	tests.AssertEqual(t, nullableBytes.Get(), nil)
}

func TestNewBytes(t *testing.T) {
	basicBytes1 := []byte{}
	nullableBytes1 := nullable.NewBytes(&basicBytes1)
	tests.AssertEqual(t, nullableBytes1.Get(), basicBytes1)

	basicBytes2 := []byte{0x0, 0x7f, 0xff}
	nullableBytes2 := nullable.NewBytes(&basicBytes2)
	tests.AssertEqual(t, nullableBytes2.Get(), basicBytes2)
	tests.AssertEqual(t, (*nullableBytes2.Get())[0], 0x0)
	tests.AssertEqual(t, (*nullableBytes2.Get())[1], 0x7f)
	tests.AssertEqual(t, (*nullableBytes2.Get())[2], 0xff)

	nullableBytes3 := nullable.NewBytes(nil)
	tests.AssertEqual(t, nullableBytes3.Get(), nil)
}

func TestSetBytes(t *testing.T) {
	nullableBytes := nullable.NewBytes(nil)
	tests.AssertEqual(t, nullableBytes.Get(), nil)

	basicBytes1 := []byte{0x0, 0x7f, 0xff}
	nullableBytes.Set(&basicBytes1)
	tests.AssertEqual(t, nullableBytes.Get(), []byte{0x0, 0x7f, 0xff})

	basicBytes2 := []byte{}
	nullableBytes.Set(&basicBytes2)
	tests.AssertEqual(t, nullableBytes.Get(), []byte{})

	basicBytes3 := []byte{0xff, 0x7f, 0x20}
	nullableBytes.Set(&basicBytes3)
	tests.AssertEqual(t, nullableBytes.Get(), []byte{0xff, 0x7f, 0x20})
	tests.AssertEqual(t, (*nullableBytes.Get())[0], 0xff)
	tests.AssertEqual(t, (*nullableBytes.Get())[1], 0x7f)
	tests.AssertEqual(t, (*nullableBytes.Get())[2], 0x20)

	nullableBytes.Set(nil)
	tests.AssertEqual(t, nullableBytes.Get(), nil)
}

func TestMutateBytes(t *testing.T) {
	basicBytes := []byte{0x33, 0xcc, 0xfd}
	nullableBytes := nullable.NewBytes(&basicBytes)
	tests.AssertEqual(t, nullableBytes.Get(), basicBytes)

	// Element modification must affect nullable bytes
	basicBytes[0] = 0x21
	basicBytes[1] = 0x77
	basicBytes[2] = 0x8f

	tests.AssertEqual(t, (*nullableBytes.Get())[0], 0x21)
	tests.AssertEqual(t, (*nullableBytes.Get())[1], 0x77)
	tests.AssertEqual(t, (*nullableBytes.Get())[2], 0x8f)

	// Append must not affect nullable bytes
	basicBytes = append(basicBytes, 0x31)
	tests.AssertEqual(t, len(*nullableBytes.Get()), 3)
}

func TestJSONBytes(t *testing.T) {
	basicBytes1 := []byte{0x0, 0x7f, 0xff}
	marshalUnmarshalJSON(t, nullable.NewBytes(&basicBytes1))

	basicBytes2 := []byte{}
	marshalUnmarshalJSON(t, nullable.NewBytes(&basicBytes2))

	marshalUnmarshalJSON(t, nullable.NewBytes(nil))
}

func TestBytes(t *testing.T) {
	type TestNullableByteArray struct {
		ID       uint
		Name     string
		HexColor nullable.Bytes
	}

	DB.Migrator().DropTable(&TestNullableByteArray{})
	if err := DB.Migrator().AutoMigrate(&TestNullableByteArray{}); err != nil {
		t.Errorf("failed to migrate nullable bytes, got error: %v", err)
	}

	blue := []byte{0x00, 0x00, 0xff}
	blueUser := TestNullableByteArray{
		Name:     "Athaariq Ardhiansyah",
		HexColor: nullable.NewBytes(&blue),
	}
	DB.Create(&blueUser)

	darkYellow := []byte{0x7f, 0x7f, 0x00}
	darkYellowUser := TestNullableByteArray{
		Name:     "Bad Cat",
		HexColor: nullable.NewBytes(&darkYellow),
	}
	DB.Create(&darkYellowUser)

	unknownUser := TestNullableByteArray{
		Name:     "Charles Robinson",
		HexColor: nullable.NewBytes(nil),
	}
	DB.Create(&unknownUser)

	var result1 TestNullableByteArray
	if err := DB.First(&result1, "name = ?", "Athaariq Ardhiansyah").Error; err != nil {
		t.Fatal("Cannot read byte array test record of \"Athaariq Ardhiansyah\"")
	}
	tests.AssertEqual(t, result1, &blueUser)

	var result2 TestNullableByteArray
	if err := DB.First(&result2, "name = ?", "Bad Cat").Error; err != nil {
		t.Fatal("Cannot read byte array test record of \"Bad Cat\"")
	}
	tests.AssertEqual(t, result2, &darkYellowUser)

	var result3 TestNullableByteArray
	if err := DB.First(&result3, "name = ?", "Charles Robinson").Error; err != nil {
		t.Fatal("Cannot read byte array test record of \"Charles Robinson\"")
	}
	tests.AssertEqual(t, result3, &unknownUser)
}
