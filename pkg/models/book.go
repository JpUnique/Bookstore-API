//nolint:gochecknoinits,gochecknoglobals //lint issue ignored
package models

import (
	"errors"

	"github.com/JpUnique/Verbatim/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// InitDB initializes the database connection and runs auto-migration.
// It should be called once from the main function of the application.
func InitDB() error {
	err := config.Connect()
	if err != nil {
		err = errors.New("failed to connect to database")
		return err
	}
	db = config.GetDB()
	return db.AutoMigrate(&Book{}).Error
}

func (b *Book) CreateBook() (createBook *Book) {
	db.NewRecord(createBook)
	db.Create(&createBook)
	return createBook
}

func GetAllBooks() (Books []Book) {
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (book *Book, dbConn *gorm.DB) {
	var getBook Book
	dbConn = db.Where("ID=?", ID).Find(&getBook)
	return &getBook, dbConn
}

func DeleteBook(ID int64) (book *Book, dbConn *gorm.DB) {
	var deleteBook Book
	dbConn = db.Where("ID=?", ID).Delete(&deleteBook)
	return &deleteBook, dbConn
}

func UpdateBook(ID int64, updateBook *Book) (updatedBook *Book, dbConn *gorm.DB) {
	var book Book
	dbConn = db.Where("ID=?", ID).First(&book)
	if dbConn.Error != nil {
		return nil, dbConn
	}
	dbConn = db.Model(&book).Updates(updateBook)
	return &book, dbConn
}
