package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

const (
	user   string = "root"
	passwd        = "root"
	dbname        = "gorm_test"
)

type Group struct {
	gorm.Model
	Name sql.NullString
	Type int `gorm:"type:enum('SVOD','TVOD')"`
}

type User struct {
	gorm.Model
	Name    sql.NullString
	User    uint
	Orders  []Order
	Group   Group
	GroupID uint
	Follows []*User `gorm:"many2many:user_follows"`
}

type Order struct {
	gorm.Model
	Currency sql.NullString
	Cents    sql.NullInt32
	Name     string
	UserID   uint
	Items    []Item
}

type Item struct {
	gorm.Model
	Amount   int
	OrderID  uint
	Products []Product `gorm:"many2many:item_products;"`
}

type Product struct {
	gorm.Model
	Cents    int
	Currency string
	Name     string
}

func main() {
	db := SetupDB()
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Group{}, &User{}, &Order{}, &Item{}, &Product{})
	if err != nil {
		panic(err)
	}

	//CreateUser(db)
	//Find(db)

	WhereName(db, "daniel")
	WhereName(db, "notfound")
}

func gormPlugin(db *gorm.DB) {
	db.Logger.Info(context.Background(), "gorm info schema=%v primaryfield=%v relationship=%v",
		db.Statement.Schema,
		db.Statement.Schema.PrioritizedPrimaryField,
		db.Statement.Schema.Relationships,
	)
}

func SetupDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "", log.Lshortfile), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, dbname)), &gorm.Config{
		Logger: newLogger,
	})

	db.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas: []gorm.Dialector{mysql.Open(fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, dbname))},
			Policy:   dbresolver.RandomPolicy{},
		},
	))
	db.Callback().Query().Register("gorm_schema_info", gormPlugin)

	if err != nil {
		panic(err)
	}
	return db
}

func Find(db *gorm.DB) {
	user := User{Model: gorm.Model{ID: 14}}
	err := db.Clauses(dbresolver.Read).Preload("Group").Preload("Orders").Preload("Orders.Items").Preload("Orders.Items.Products").Find(&user).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("user: %+v\n", user)
}

func WhereName(db *gorm.DB, name string) {
	users := []User{}
	err := db.Where("name = ?", name).Select("id").Find(&users).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("users:", users)
}

func GroupBy(db *gorm.Model) {
	//r
	//err := db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
	//if err != nil {
	//panic(err)
	//}

}

func CreateUser(db *gorm.DB) {
	user := User{
		Name: sql.NullString{String: "daniel", Valid: true},
		Group: Group{
			Name: sql.NullString{String: "kktix", Valid: true},
		},
		Orders: []Order{
			{Items: []Item{
				{
					Products: []Product{
						{
							Name:     "macbook",
							Cents:    10000,
							Currency: "TWD",
						},
					},
				},
			}},
		},
	}
	db.Omit("Orders.*").Save(&user) //item 還是會被建立
	// db.Omit("Orders.Items").Save(&user) items 不會被建立
	// db.Omit("Orders.Items.*").Save(&user) // 還是有 Items & product = =
	//db.Omit("Orders.Items.*").Save(&user) // result:
	fmt.Printf("%+v\n", user)
}
