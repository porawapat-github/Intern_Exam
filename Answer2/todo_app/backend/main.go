package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	IsCompleted bool      `json:"is_completed" db:"is_completed"`
	Priority    int       `json:"priority" db:"priority"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Connect("postgres", "host=localhost port=5432 user=dev_todo_app password=1234 dbname=todo_app sslmode=disable search_path=public")
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", createTodo)
	router.PUT("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run(":8080")
}

func getTodoByID(c *gin.Context) {
	// รับค่า ID จาก Path
	id := c.Param("id")
	var todo Todo
	// ค้นหา Todo ตาม ID
	err := db.Get(&todo, "SELECT * FROM todos WHERE id=$1", id)
	// ตรวจสอบว่าไม่พบข้อมูล
	if err != nil {
		log.Printf("❌ getTodoByID error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	// ส่งข้อมูล Todo กลับ
	c.JSON(http.StatusOK, todo)
}

func getTodos(c *gin.Context) {
	var todos []Todo
	err := db.Select(&todos, "SELECT * FROM todos ORDER BY priority ASC")
	if err != nil {
		log.Printf("❌ getTodos error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var t Todo
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	// 🔍 Log ข้อมูลก่อน insert
	log.Printf("📥 Inserting: %+v", t)

	err := db.Get(&t, `
		INSERT INTO todos (title, description, is_completed, priority, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, description, is_completed, priority, created_at, updated_at
	`, t.Title, t.Description, t.IsCompleted, t.Priority, t.CreatedAt, t.UpdatedAt)

	if err != nil {
		log.Printf("❌ createTodo error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ✅ Log หลัง insert สำเร็จ
	log.Printf("✅ Inserted ID: %d", t.ID)

	c.JSON(http.StatusOK, t)
}

func updateTodo(c *gin.Context) {
	var t Todo
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	t.UpdatedAt = time.Now()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = db.Exec(`
		UPDATE todos
		SET title=$1, description=$2, is_completed=$3, priority=$4, updated_at=$5
		WHERE id=$6
	`, t.Title, t.Description, t.IsCompleted, t.Priority, t.UpdatedAt, idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	confirm := c.Query("confirm")

	if confirm != "true" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please confirm delete with ?confirm=true"})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = db.Exec("DELETE FROM todos WHERE id=$1", idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
