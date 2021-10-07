package timer

// This file holds context functions that will be shared by all integration tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

var testDB = getGormDBOrDie()

// getGormDBOrDie returns a value of gorm.DB or panic error occurs
func getGormDBOrDie() *gorm.DB {
	dsn := fmt.Sprintf(
		`%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s`,
		"root",
		"root",
		"localhost",
		3306,
		"test",
		"utf8mb4",
		true,
		"Local",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to get db session for test")
	}
	return db
}

func queryRaw(db *gorm.DB, sql string, result interface{}, args ...interface{}) *gorm.DB {
	return db.Raw(sql, args...).Scan(result)
}

func execRaw(db *gorm.DB, sql string, args ...interface{}) *gorm.DB {
	return db.Exec(sql, args...)
}

func plantTimerOrDie(db *gorm.DB, tc *model.TimerCore) {
	if err := execRaw(db, "INSERT INTO timer (name, trigger_at) VALUES (?, ?)", tc.Name, tc.TriggerAt).Error; err != nil {
		panic(fmt.Sprintf("failed to plant timer with name %s, trigger_at %s", tc.Name, tc.TriggerAt.Format(time.RFC3339)))
	}
}

func assertTimerNotExistByName(t *testing.T, db *gorm.DB, name string) {
	sql := `SELECT * FROM timer WHERE name = ? AND deleted_at IS NULL`
	var timer model.Timer
	result := queryRaw(db, sql, &timer, name)
	assert.Equal(t, result.RowsAffected, int64(0))
	assert.Equal(t, timer.ID, uint(0))
}

func assertTimerExists(t *testing.T, db *gorm.DB, tc *model.TimerCore) {
	sql := `SELECT * FROM timer WHERE name = ? AND deleted_at IS NULL`
	var timer model.Timer
	result := queryRaw(db, sql, &timer, tc.Name)
	assert.Equal(t, result.RowsAffected, int64(1))
	assert.NotEqual(t, timer.ID, uint(0))
	assert.NotEmpty(t, timer.CreatedAt)
	assert.NotEmpty(t, timer.UpdatedAt)
	assert.Equal(t, timer.Name, tc.Name)
	assert.Equal(t, timer.TriggerAt, tc.TriggerAt)
	assert.True(t, timer.Alive)
	assert.False(t, timer.DeletedAt.Valid)
}
