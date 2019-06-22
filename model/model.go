package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENTlcolumn:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:CreatedAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type Token struct {
	Token string `json:"token"`
}

//老代码
// type Database struct {
// 	Self   *gorm.DB
// 	Docker *gorm.DB
// }

// var DB *Database

// func openDB(username, password, addr, name string) *gorm.DB {
// 	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
// 		username,
// 		password,
// 		addr,
// 		name,
// 		true,
// 		"Local")

// 	db, err := gorm.Open("mysql", config)
// 	if err != nil {
// 		log.Errorf(err, "Database connection failed. Database name: %s", name)
// 	}

// 	setupDB(db)
// 	return db
// }

// func setupDB(db *gorm.DB) {
// 	db.LogMode(viper.GetBool("gormlog"))
// 	db.DB().SetMaxIdleConns(0)
// }

// func InitSelfDB() *gorm.DB {
// 	return openDB(viper.GetString("db.username"),
// 		viper.GetString("db.password"),
// 		viper.GetString("db.addr"),
// 		viper.GetString("db.name"))
// }

// func GetSelfDB() *gorm.DB {
// 	return InitSelfDB()
// }

// func InitDockerDB() *gorm.DB {
// 	return openDB(viper.GetString("docker_db.username"),
// 		viper.GetString("docker_db.password"),
// 		viper.GetString("docker_db.addr"),
// 		viper.GetString("docker_db.name"))
// }

// func GetDockerDB() *gorm.DB {
// 	return InitDockerDB()
// }

// func (db *Database) Init() {
// 	DB = &Database{
// 		Self:   GetSelfDB(),
// 		Docker: GetDockerDB(),
// 	}
// }

// func (db *Database) Close() {
// 	DB.Self.Close()
// 	DB.Docker.Close()
// }
