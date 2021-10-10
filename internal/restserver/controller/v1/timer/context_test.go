// context_test.go holds auxiliary functions shared within this package.
package timer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

func newTestGinCtxWithReq(method, url string, body map[string]interface{}) *gin.Context {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	var reqBody io.Reader
	switch body {
	case nil:
		reqBody = nil
	default:
		bodyBytes, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(bodyBytes)
	}

	c.Request, _ = http.NewRequest(method, url, reqBody)
	c.Request.Header.Set("Content-Type", "application/json")
	return c
}

func monkeyPatchValidateTimerFunc(ret error) (restore func()) {
	old := validateTimerFunc
	restore = func() { validateTimerFunc = old }
	validateTimerFunc = func(*model.Timer) error { return ret }
	return
}

func monkeyPatchBindJsonFunc(ret error) (restore func()) {
	old := bindJsonFunc
	restore = func() { bindJsonFunc = old }
	bindJsonFunc = func(c *gin.Context, obj interface{}) error { return ret }
	return
}

func monkeypatch_validateTimerCoreFunc(ret error) (restore func()) {
	old := validateTimerCoreFunc
	restore = func() { validateTimerCoreFunc = old }
	validateTimerCoreFunc = func(tc *model.TimerCore) error { return ret }
	return
}

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

func timerExistsInDBByName(db *gorm.DB, name string) bool {
	var fetchedName string
	db.Raw("SELECT name FROM timer WHERE name = ?", name).Scan(&fetchedName)
	return fetchedName != ""
}
