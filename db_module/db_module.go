package db_module

import (
	_ "database/sql"
	"fmt"
	_ "fmt"
	_ "log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	User_id    string
	Start_date string
	End_date   string
	Title      string
	Status     string
}

func InsertTodo(c echo.Context) error {
	dsn := "root:1234@tcp(127.0.0.1:3306)/todo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Todo{})

	db.Create(&Todo{
		User_id:    c.QueryParam("user_id"),
		Start_date: c.QueryParam("start_date"),
		End_date:   c.QueryParam("end_date"),
		Title:      c.QueryParam("title"),
		Status:     c.QueryParam("status"),
	})

	return c.String(http.StatusOK, "insert")
}

func SelectTodo(c echo.Context) error {
	dsn := "root:1234@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	todos := []Todo{}

	db.Where("user_id LIKE ?", "%"+c.QueryParam("user_id")+"%").Find(&todos)

	for _, todo := range todos {
		fmt.Println(todo)
	}

	return c.String(http.StatusOK, "select")
}

func UpdateTodo(c echo.Context) error {
	dsn := "root:1234@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	todo := Todo{}
	db.First(&todo, "user_id = ?", c.QueryParam("id_query"))
	//fmt.Println(c.QueryParam("user_id"))

	if c.QueryParam("user_id") != "" {
		todo.User_id = c.QueryParam("user_id")
	}

	if c.QueryParam("start_date") != "" {
		todo.Start_date = c.QueryParam("start_date")
	}

	if c.QueryParam("end_date") != "" {
		todo.End_date = c.QueryParam("end_date")
	}

	if c.QueryParam("title") != "" {
		todo.Title = c.QueryParam("title")
	}

	if c.QueryParam("status") != "" {
		todo.Status = c.QueryParam("status")
	}

	db.Save(&todo)

	return c.String(http.StatusOK, "update")
}

func DeleteTodo(c echo.Context) error {
	dsn := "root:1234@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//fmt.Println("hello")

	todo := Todo{}
	db.First(&todo, "user_id = ?", c.QueryParam("id_query"))
	db.Delete(&todo)

	return c.String(http.StatusOK, "delete")
}

//func InsertTodo(c echo.Context) error {
//db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/todo")
//if err != nil {
//log.Fatal(err)
//return err
//}
//defer db.Close()

//todo := Todo{
//user_id:    c.QueryParam("user_id"),
//start_date: c.QueryParam("start_date"),
//end_date:   c.QueryParam("end_date"),
//title:      c.QueryParam("title"),
//status:     c.QueryParam("status"),
//}

//_, err = db.Exec(`INSERT INTO
//todo
//VALUES (?, ?, ?, ?, ?)`,
//todo.user_id, todo.start_date, todo.end_date, todo.title, todo.status)
//if err != nil {
//log.Fatal(err)
//return err
//}

//fmt.Println("insert succeeded")

//return c.String(http.StatusOK, "insert succeded")
//}

//func SelectTodo(c echo.Context) error {
//db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/todo")
//if err != nil {
//log.Fatal(err)
//return err
//}
//defer db.Close()

//user_id := c.QueryParam("user_id")

//rows, err := db.Query(`SELECT * FROM todo where user_id like ?`, "%"+user_id+"%")
//if err != nil {
//log.Fatal(err)
//}
//defer rows.Close()

//for rows.Next() {
//todo := Todo{}

//err = rows.Scan(&todo.user_id, &todo.start_date, &todo.end_date, &todo.title, &todo.status)
//if err != nil {
//log.Fatal(err)
//return err
//}

//fmt.Println(todo)

//}

//return c.String(http.StatusOK, "select succeeded")
//}

//func UpdateTodo(c echo.Context) error {

//db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/todo")
//if err != nil {
//log.Fatal(err)
//return err
//}
//defer db.Close()

//query_user := c.QueryParam("query_user")
//todo := Todo{
//user_id:    c.QueryParam("user_id"),
//start_date: c.QueryParam("start_date"),
//end_date:   c.QueryParam("end_date"),
//title:      c.QueryParam("title"),
//status:     c.QueryParam("status"),
//}

//stmt, err := db.Prepare("UPDATE todo SET user_id=?, start_date=?, end_date=?, title=?, status=? WHERE user_id=?;")
//if err != nil {
//log.Fatal(err)
//return err
//}
//defer stmt.Close()

//_, err = stmt.Exec(todo.user_id, todo.start_date, todo.end_date, todo.title, todo.status, query_user)
//if err != nil {
//log.Fatal(err)
//return err
//}

//fmt.Println("update succeeded")
//fmt.Println(todo.user_id, todo.start_date, todo.end_date, todo.title, todo.status)

//return c.String(http.StatusOK, "update succeded")
//}

//func DeleteTodo(c echo.Context) error {
//db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/todo")
//if err != nil {
//log.Fatal(err)
//return err
//}
//defer db.Close()

//user_id := c.QueryParam("user_id")
//stmt, err := db.Prepare("Delete from todo WHERE user_id=?")
//if err != nil {
//log.Fatal(err)
//return err
//}
//defer stmt.Close()

//_, err = stmt.Exec(user_id)
//if err != nil {
//log.Fatal(err)
//return err
//}

//return c.String(http.StatusOK, "deleted okay")
//}
