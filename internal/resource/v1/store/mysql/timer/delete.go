package timer

import (
	"gorm.io/gorm"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

var dbDeleteByNameFunc = func(db *gorm.DB, name string) error {
	return db.Where("name = ?", name).Delete(&model.Timer{}).Error
}

var tsGetByNameFunc = func(ts *TimerStore, name string) (*model.Timer, error) {
	return ts.GetByName(name)
}

func (s *TimerStore) DeleteByName(name string) error {
	_, err := tsGetByNameFunc(s, name)
	if err != nil {
		return err
	}

	delErr := dbDeleteByNameFunc(s.DB, name)
	if delErr != nil {
		return delErr
	}
	return nil
}
