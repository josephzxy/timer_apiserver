package timer

import (
	"gorm.io/gorm"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

var dbDeleteByNameFunc = func(db *gorm.DB, name string) error {
	return db.Where("name = ?", name).Delete(&model.Timer{}).Error
}

func (s *TimerStore) DeleteByName(name string) error {
	_, err := s.GetByName(name)
	if err != nil {
		return err
	}

	delErr := dbDeleteByNameFunc(s.DB, name)
	if err != nil {
		return delErr
	}
	return nil
}
