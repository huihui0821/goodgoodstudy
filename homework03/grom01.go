package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User 用户
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true"`
	Username  string    `gorm:"size:64;uniqueIndex;not null"`
	Email     string    `gorm:"size:128;uniqueIndex;not null"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// 一对多
	Posts []Post `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Post 文章
type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true"`
	Title     string    `gorm:"size:200;not null"`
	Slug      string    `gorm:"size:200;uniqueIndex;not null"` // URL 友好唯一标识
	Content   string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null;index"` // 外键（GORM 会自动关联）
	Published bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// 反向关联（可选，用于查询时自动加载作者）
	User User `gorm:"foreignKey:UserID"`

	// 一对多：文章拥有多个评论
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Comment 评论（属于一篇文章）
type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true"`
	Content   string    `gorm:"type:text;not null"`
	Author    string    `gorm:"size:64;not null"` // 可匿名，也可后续关联 User
	Email     string    `gorm:"size:128"`
	PostID    uint      `gorm:"not null;index"` // 外键
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// 反向关联（可选）
	Post Post `gorm:"foreignKey:PostID"`
}

func main() {
	dsn := "root:123456@tcp(localhost:3306)/community?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	println("连接数据库成功")
}
