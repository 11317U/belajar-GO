package main

import (
	"apigorm/controller"
	"apigorm/database"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const PORT = ":3000"

type Controllers struct {
	masterDB *gorm.DB
}

// func New(db *gorm.DB) *Controllers {
// 	return &Controllers{
// 		masterDB: db,
// 	}
// }

// func (c *Controllers) CreateBook(ctx *gin.Context) {
// 	var newBook models.Books

// 	err := ctx.ShouldBindJSON(&newBook)

// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	if err := c.masterDB.Create(&newBook).Error; err != nil {
// 		ctx.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}
// }

func main() {
	fmt.Println("Starting server.......")

	// db := database.db()
	// controller := controller.New(db)

	dbs := database.Connect()
	controller := controller.New(dbs)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/book", controller.CreateBook)

	// Users
	// users := r.Group("/v1/users")
	// {
	// 	users.POST("/", controller.CreateBook)
	// 	users.GET("/products", controller.GetUsersWithProducts)
	// }

	// c := Controllers{
	// 	masterDB: db,
	// }

	r.Run(PORT)
}