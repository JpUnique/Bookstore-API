//nolint:gochecknoglobals //lint issue ignored
package config

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() (err error) {
	d, err := gorm.Open("mysql", "jpunique:zyxwvu@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		err = errors.New("failed to connect to database")
		return err
	}
	db = d
	return err
}
func GetDB() *gorm.DB {
	return db
}
