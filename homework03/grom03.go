package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" // 以 MySQL 为例，可替换为 sqlite、postgres 等
	"gorm.io/gorm"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"size:100;unique;not null"`
	Password   string `gorm:"size:255;not null"` // 实际项目中应加密存储
	PostsCount uint   `gorm:"default:0"`         // 文章数量统计
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Posts []Post `gorm:"foreignKey:UserID"` // 一对多：用户有多个文章
}

type Post struct {
	ID           uint   `gorm:"primaryKey"`
	Title        string `gorm:"size:200;not null"`
	Content      string `gorm:"type:text"`
	CommentCount uint   `gorm:"default:0"` // 文章数量统计
	UserID       uint   // 外键
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Comments []Comment `gorm:"foreignKey:PostID"` // 一对多：文章有多个评论
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   // 外键
	UserID    uint   // 评论者（可选，这里也关联用户）
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	// 连接数据库（替换为你的 DSN）
	dsn := "root:123456@tcp(localhost:3306)/community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//使用 AutoMigrate 创建表
	// err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// if err != nil {
	// 	panic("failed to migrate")
	// }
	fmt.Println("数据库表创建/迁移完成")
	// 测试钩子函数1
	testFuncHook(db)

	// // 示例数据插入（可选，用于测试查询）
	// insertSampleData(db)

	// // 3. 查询某个用户的所有文章及其评论
	//queryUserPostsWithComments(db, 1) // 假设查询 ID 为 1 的用户

	// // 4. 查询评论数量最多的文章
	// queryPostWithMostComments(db)
}

func (c *Comment) BeforeCreate(db *gorm.DB) error {
	if c.PostID > 0 {
		var count int64
		db.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)

		// 更新文章的 CommentsCount 和 CommentStatus
		status := "有评论"
		if count == 0 {
			status = "无评论"
		}
		fmt.Printf("status: %v\n", status)
		fmt.Printf("count: %v\n", count)

	}
	return nil
}

func queryUserPostsWithComments(db *gorm.DB, userID uint) {
	var user User
	// Preload("Posts") 预加载文章，嵌套 Preload("Posts.Comments") 预加载每篇文章的评论
	err := db.Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("用户 ID %d 不存在\n", userID)
			return
		}
		fmt.Println("查询失败:", err)
		return
	}

	fmt.Printf("用户 %s 发布的文章（共 %d 篇）:\n", user.Username, len(user.Posts))
	for _, post := range user.Posts {
		fmt.Printf("  文章标题: %s (ID: %d, 评论数: %d)\n", post.Title, post.ID, len(post.Comments))
		for _, comment := range post.Comments {
			fmt.Printf("    评论: %s\n", comment.Content)
		}
	}
}

func queryPostWithMostComments(db *gorm.DB) {
	type Result struct {
		Post
		CommentCount int64
	}

	var result Result
	// 子查询统计评论数，按评论数降序，取第一条
	subQuery := db.Model(&Comment{}).Select("COUNT(*)").Where("post_id = posts.id")
	err := db.Model(&Post{}).
		Select("posts.*, (?) as comment_count", subQuery).
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}

	if result.ID == 0 {
		fmt.Println("暂无文章或评论")
		return
	}

	fmt.Printf("评论最多的文章:\n")
	fmt.Printf("  ID: %d\n", result.ID)
	fmt.Printf("  标题: %s\n", result.Title)
	fmt.Printf("  作者ID: %d\n", result.UserID)
	fmt.Printf("  评论数量: %d\n", result.CommentCount)
}

func testFuncHook(db *gorm.DB) {
	// 2. 按照 name 精确查找（最常用）
	var user User
	result := db.Where("username = ?", "bob1").First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("没有找到名为 张三 的用户")
		} else {
			log.Fatal("查询失败:", result.Error)
		}
	} else {
		fmt.Printf("找到用户: ID=%d, Name=%s", user.ID, user.Username)
	}
	println("\n")
	var post Post
	result = db.Where("user_id = ?", 1).First(&post)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("没有找到user_id为 1 的用户")
		} else {
			log.Fatal("查询失败:", result.Error)
		}
	} else {
		fmt.Printf("找到用户: ID=%d, Name=%s", post.UserID, post.Title)
	}
	db.Create(&Comment{Content: "好评11111！", PostID: post.ID, UserID: user.ID})
	var updatedPost Post
	db.First(&updatedPost, post.ID)
	fmt.Printf("updatedPost.CommentCount: %v\n", updatedPost.CommentCount)
	// result := db.Where("username = ?", "bob").Delete(&User{})
	// if result.Error != nil {
	// 	log.Fatal(result.Error)
	// }
	// fmt.Printf("删除了 %d 条记录\n", result.RowsAffected)
	// 1. 创建用户
	// user := User{Username: "bob1", Password: "secret"}
	// db.Create(&user)
	// fmt.Printf("创建用户 %s,初始文章数: %d\n", user.Username, user.PostsCount)

	// // 2. 创建文章 → 触发 Post.AfterCreate → 用户 PostsCount +1
	// post := Post{Title: "GORM 钩子函数", Content: "···内容...", UserID: user.ID}
	// db.Create(&post)
	// db.Model(&user).First(&user) // 刷新查看最新计数
	// fmt.Printf("创建文章后，用户文章数: %d\n", user.PostsCount)

	// // 3. 创建两条评论
	// db.Create(&Comment{Content: "好评！", PostID: post.ID, UserID: user.ID})
	// db.Create(&Comment{Content: "学习了", PostID: post.ID, UserID: user.ID})

	// // 刷新文章查看状态
	// var updatedPost Post
	// db.First(&updatedPost, post.ID)
	// fmt.Printf("updatedPost.CommentCount: %v\n", updatedPost.CommentCount)
}

// Hook: 创建文章后，自动更新用户的 PostsCount
func (p *Post) AfterCreate(db *gorm.DB) (err error) {
	if p.UserID > 0 {
		db.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("posts_count", gorm.Expr("posts_count + ?", 1))
	}
	return nil
}

func insertSampleData(db *gorm.DB) {
	user := User{Username: "alice", Password: "123456"}
	db.Create(&user)

	post1 := Post{Title: "第一篇文章", Content: "内容1", UserID: user.ID}
	post2 := Post{Title: "第二篇文章", Content: "内容2", UserID: user.ID}
	db.Create(&post1)
	db.Create(&post2)

	db.Create(&Comment{Content: "评论1", PostID: post1.ID, UserID: user.ID})
	db.Create(&Comment{Content: "评论2", PostID: post1.ID, UserID: user.ID})
	db.Create(&Comment{Content: "评论3", PostID: post1.ID, UserID: user.ID})
	db.Create(&Comment{Content: "评论4", PostID: post2.ID, UserID: user.ID})
}
