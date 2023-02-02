package nullable_test

import (
	"testing"
	"unicode/utf8"

	"github.com/tee8z/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestNonUTF8SafeString(t *testing.T) {
	// use a regular utf8 valid string
	utf8String := "meow i am a cat"
	isValid := utf8.Valid([]byte(utf8String))
	tests.AssertEqual(t, isValid, true)
	nullableString := nullable.NewString(&utf8String)
	isValid = utf8.Valid([]byte(*nullableString.Get()))
	tests.AssertEqual(t, isValid, true)

	// pass in a non utf8 valid string, ensure we get out a valid utf8 string
	notUtf8string := "\x17\x87\x1c\x97\xbbs\x159X\xad\xb0[\xca\xe0O\x0b2\xad\xe5\xa3lp\x9b.\xfe\x8eݘ\x15\xfe\xb5Q"
	isValid = utf8.Valid([]byte(notUtf8string))
	tests.AssertEqual(t, isValid, false)
	nullableString = nullable.NewString(&notUtf8string)
	isValid = utf8.Valid([]byte(*nullableString.Get()))
	tests.AssertEqual(t, isValid, true)
}

func TestScanString(t *testing.T) {
	nullableString := nullable.NewString(nil)

	nullableString.Scan("")
	tests.AssertEqual(t, nullableString.Get(), "")

	nullableString.Scan("This is a test string")
	tests.AssertEqual(t, nullableString.Get(), "This is a test string")

	nullableString.Scan("and This is also a test string that really really long, just in case something fails after somebody enter a really long string like this. You know what? Coding unit test is a lot stressful and spend longer time than making the real code itself. So please show me a little respect of writting this really long string. Thank you!")
	tests.AssertEqual(t, nullableString.Get(), "and This is also a test string that really really long, just in case something fails after somebody enter a really long string like this. You know what? Coding unit test is a lot stressful and spend longer time than making the real code itself. So please show me a little respect of writting this really long string. Thank you!")

	nullableString.Scan("~!@#$%^&*()_+`-=:;\"'/\\")
	tests.AssertEqual(t, nullableString.Get(), "~!@#$%^&*()_+`-=:;\"'/\\")

	nullableString.Scan(nil)
	tests.AssertEqual(t, nullableString.Get(), nil)
}

func TestNewString(t *testing.T) {
	basicString1 := ""
	nullableString1 := nullable.NewString(&basicString1)
	tests.AssertEqual(t, nullableString1.Get(), "")

	basicString2 := "This is a test string"
	nullableString2 := nullable.NewString(&basicString2)
	tests.AssertEqual(t, nullableString2.Get(), "This is a test string")

	basicString3 := "and This is also a test string that really really long, just in case something fails after somebody enter a really long string like this. You know what? Coding unit test is a lot stressful and spend longer time than making the real code itself. So please show me a little respect of writting this really long string. Thank you!"
	nullableString3 := nullable.NewString(&basicString3)
	tests.AssertEqual(t, nullableString3.Get(), "and This is also a test string that really really long, just in case something fails after somebody enter a really long string like this. You know what? Coding unit test is a lot stressful and spend longer time than making the real code itself. So please show me a little respect of writting this really long string. Thank you!")

	basicString4 := "~!@#$%^&*()_+`-=:;\"'/\\"
	nullableString4 := nullable.NewString(&basicString4)
	tests.AssertEqual(t, nullableString4.Get(), "~!@#$%^&*()_+`-=:;\"'/\\")

	nullableString5 := nullable.NewString(nil)
	tests.AssertEqual(t, nullableString5.Get(), nil)
}

func TestSetString(t *testing.T) {
	nullableString := nullable.NewString(nil)
	tests.AssertEqual(t, nullableString.Get(), nil)

	basicString1 := ""
	nullableString.Set(&basicString1)
	tests.AssertEqual(t, nullableString.Get(), "")

	basicString2 := "This is a test string"
	nullableString.Set(&basicString2)
	tests.AssertEqual(t, nullableString.Get(), "This is a test string")

	basicString3 := "and This is also a test string that really really long, just in case something fails after somebody enter a really long string like this. You know what? Coding unit test is a lot stressful and spend longer time than making the real code itself. So please show me a little respect of writting this really long string. Thank you!"
	nullableString.Set(&basicString3)
	tests.AssertEqual(t, nullableString.Get(), "and This is also a test string that really really long, just in case something fails after somebody enter a really long string like this. You know what? Coding unit test is a lot stressful and spend longer time than making the real code itself. So please show me a little respect of writting this really long string. Thank you!")

	basicString4 := "~!@#$%^&*()_+`-=:;\"'/\\"
	nullableString.Set(&basicString4)
	tests.AssertEqual(t, nullableString.Get(), "~!@#$%^&*()_+`-=:;\"'/\\")

	basicString5 := ""
	nullableString.Set(&basicString5)
	tests.AssertEqual(t, nullableString.Get(), "")

	nullableString.Set(nil)
	tests.AssertEqual(t, nullableString.Get(), nil)
}

func TestJSONString(t *testing.T) {
	basicString1 := ""
	marshalUnmarshalJSON(t, nullable.NewString(&basicString1))

	basicString2 := "This is a test string"
	marshalUnmarshalJSON(t, nullable.NewString(&basicString2))

	basicString3 := "and This is also a test string that really really long, just in case something fails after somebody enter a really long string like this. You know what? Coding unit test is a lot stressful and spend longer time than making the real code itself. So please show me a little respect of writting this really long string. Thank you!"
	marshalUnmarshalJSON(t, nullable.NewString(&basicString3))

	basicString4 := "~!@#$%^&*()_+`-=:;\"'/\\"
	marshalUnmarshalJSON(t, nullable.NewString(&basicString4))

	basicString5 := ""
	marshalUnmarshalJSON(t, nullable.NewString(&basicString5))

	marshalUnmarshalJSON(t, nullable.NewString(nil))
}

func TestString(t *testing.T) {
	type TestNullableString struct {
		ID          uint
		Title       string
		Description nullable.String
	}

	DB.Migrator().DropTable(&TestNullableString{})
	if err := DB.Migrator().AutoMigrate(&TestNullableString{}); err != nil {
		t.Errorf("failed to migrate nullable string, got error: %v", err)
	}

	productDesc1 := "The real java coffee, directly came from Java Island, Indonesia"
	product1 := TestNullableString{
		Title:       "Java Coffee",
		Description: nullable.NewString(&productDesc1),
	}
	DB.Create(&product1)

	productDesc2 := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	product2 := TestNullableString{
		Title:       "Lorem Ipsum",
		Description: nullable.NewString(&productDesc2),
	}
	DB.Create(&product2)

	productDesc3 := "~!@#$%^&*()_+`-=:;\"'/\\"
	product3 := TestNullableString{
		Title:       "Abstract Thing",
		Description: nullable.NewString(&productDesc3),
	}
	DB.Create(&product3)

	productDesc4 := ""
	product4 := TestNullableString{
		Title:       "No Title",
		Description: nullable.NewString(&productDesc4),
	}
	DB.Create(&product4)

	product5 := TestNullableString{
		Title:       "Unknown Book",
		Description: nullable.NewString(nil),
	}
	DB.Create(&product5)

	productDesc6 := "This text\"--supposed to not causing any harm"
	product6 := TestNullableString{
		Title:       "XSS Injection",
		Description: nullable.NewString(&productDesc6),
	}
	DB.Create(&product6)

	var result1 TestNullableString
	if err := DB.First(&result1, "title = ?", "Java Coffee").Error; err != nil {
		t.Fatal("Cannot read string test record of \"Java Coffee\"")
	}
	tests.AssertEqual(t, result1, product1)

	var result2 TestNullableString
	if err := DB.First(&result2, "title = ?", "Lorem Ipsum").Error; err != nil {
		t.Fatal("Cannot read string test record of \"Lorem Ipsum\"")
	}
	tests.AssertEqual(t, result2, product2)

	var result3 TestNullableString
	if err := DB.First(&result3, "title = ?", "Abstract Thing").Error; err != nil {
		t.Fatal("Cannot read string test record of \"Abstract Thing\"")
	}
	tests.AssertEqual(t, result3, product3)

	var result4 TestNullableString
	if err := DB.First(&result4, "title = ?", "No Title").Error; err != nil {
		t.Fatal("Cannot read string test record of \"No Title\"")
	}
	tests.AssertEqual(t, result4, product4)

	var result5 TestNullableString
	if err := DB.First(&result5, "title = ?", "Unknown Book").Error; err != nil {
		t.Fatal("Cannot read string test record of \"Unknown Book\"")
	}
	tests.AssertEqual(t, result5, product5)

	var result6 TestNullableString
	if err := DB.First(&result6, "title = ?", "XSS Injection").Error; err != nil {
		t.Fatal("Cannot read string test record of \"XSS Injection\"")
	}
	tests.AssertEqual(t, result6, product6)
}
