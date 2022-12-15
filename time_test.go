package nullable_test

import (
	"testing"
	"time"

	"github.com/tee8z/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanTime(t *testing.T) {
	nullableTime := nullable.NewTime(nil)

	basicTime := time.Now()

	nullableTime.Scan(basicTime.Local())
	tests.AssertEqual(t, nullableTime.Get().Unix(), basicTime.Unix())

	nullableTime.Scan(basicTime.UTC())
	tests.AssertEqual(t, nullableTime.Get().Unix(), basicTime.Unix())

	nullableTime.Scan(nil)
	tests.AssertEqual(t, nullableTime.Get(), nil)
}

func TestNewTime(t *testing.T) {
	basicTime1 := time.Now()
	nullableTime1 := nullable.NewTime(&basicTime1)
	tests.AssertEqual(t, nullableTime1.Get().Unix(), basicTime1.Unix())

	nullableTime2 := nullable.NewTime(nil)
	tests.AssertEqual(t, nullableTime2.Get(), nil)
}

func TestSetTime(t *testing.T) {
	nullableTime := nullable.NewTime(nil)
	tests.AssertEqual(t, nullableTime.Get(), nil)

	basicTime1 := time.Now()
	nullableTime.Set(&basicTime1)
	tests.AssertEqual(t, nullableTime.Get().Unix(), basicTime1.Unix())

	nullableTime.Set(nil)
	tests.AssertEqual(t, nullableTime.Get(), nil)
}

func TestJSONTime(t *testing.T) {
	basicTime := time.Now()
	marshalUnmarshalJSON(t, nullable.NewTime(&basicTime))

	marshalUnmarshalJSON(t, nullable.NewTime(nil))
}

func TestTime(t *testing.T) {
	type TestNullableTime struct {
		UserID     uint `gorm:"primaryKey"`
		BookName   string
		BorrowedAt time.Time
		ReturnedAt nullable.Time
	}

	DB.Migrator().DropTable(&TestNullableTime{})
	if err := DB.Migrator().AutoMigrate(&TestNullableTime{}); err != nil {
		t.Errorf("failed to migrate nullable time, got error: %v", err)
	}

	userReturn1 := time.Now().Add(-7 * 24 * time.Hour)
	user1 := TestNullableTime{
		UserID:     777,
		BookName:   "How To Mating like Cats",
		BorrowedAt: time.Now().Add(-14 * 24 * time.Hour),
		ReturnedAt: nullable.NewTime(&userReturn1),
	}
	DB.Create(&user1)

	user2 := TestNullableTime{
		UserID:     234,
		BookName:   "Return this or DIE",
		BorrowedAt: time.Now().Add(-7 * 24 * time.Hour),
		ReturnedAt: nullable.NewTime(nil),
	}
	DB.Create(&user2)

	var result1 TestNullableTime
	if err := DB.First(&result1, "user_id = ?", 777).Error; err != nil {
		t.Errorf("failed to migrate nullable time, got error: %v", err)
	}
	tests.AssertEqual(t, result1, user1)

	var result2 TestNullableTime
	if err := DB.First(&result2, "user_id = ?", 234).Error; err != nil {
		t.Errorf("failed to migrate nullable time, got error: %v", err)
	}
	tests.AssertEqual(t, result2, user2)
}
