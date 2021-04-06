package main

//import (
//  "fmt"
//  "github.com/jinzhu/gorm"
//  _"github.com/lib/pq"
//  "time"
//)
//
//type Post struct {
//  Id int
//  Content string
//  Author string
//  Comments []Comment
//  CreatedAt time.Time
//}
//
//type Comment struct {
//  Id int
//  Content string
//  Author string
//  PostId int
//  CreatedAt time.Time
//}
//
//var Db *gorm.DB
//
//func init() {
//  var err error
//  Db, err = gorm.Open("postgres", "user=postgres dbname=go-web-app sslmode=disable")
//  if err != nil {
//    panic(err)
//  }
//  Db.AutoMigrate(&Post{}, &Comment{})
//}
//
//func main() {
//  post := Post{Content: "Hello World!", Author: "Sau Sheong"}
//  Db.Create(&post)
//  fmt.Println(post)
//
//  comment := Comment{Content: "Nice!", Author: "Sau Sheong"}
//  Db.Model(&post).Association("Comments").Append(comment)
//  Db.Create(&comment)
//  fmt.Println(comment)
//
//  fmt.Println(post)
//  var comments []Comment
//  Db.Model(&post).Related(&comments)
//  fmt.Println(comments)
//}
