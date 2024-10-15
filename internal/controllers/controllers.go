package controllers

import (
	"ProductService/internal/db"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func CreateProduct(c *gin.Context) {
	body := Product{}
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(400, "Bad Input")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		c.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	_, err = db.DB.Exec(`INSERT INTO products (title, price) VALUES (?, ?)`, body.Title, body.Price)
	if err != nil {
		log.Println("err", err)
		c.AbortWithStatusJSON(400, "Couldn't create the new product.")
	} else {
		c.JSON(http.StatusOK, "User is successfully created.")
	}
}

func GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := idToInt(idParam)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Input")
		return
	}

	var product Product
	err = db.DB.Exec(`SELECT * FROM products WHERE id = ?`, id).Scan(&product)
	if err != nil {
		c.AbortWithStatusJSON(400, "Couldn't find the user.")
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetAllProducts(c *gin.Context) {
	products, err := db.DB.Exec("select * from products")
	if err != nil {
		c.AbortWithStatusJSON(400, "Couldn't find the products.")
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := idToInt(idParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Input")
		return
	}
	log.Println("id", id)
	body := Product{}

	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Input")
		return
	}

	err = json.Unmarshal(data, &body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Input")
		return
	}

	_, err = db.DB.Exec(`UPDATE products SET title = ?, price = ? WHERE id = ?`, body.Title, body.Price, id)
	if err != nil {
		log.Println("Error updating product:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Couldn't update the product.")
		return
	}

	c.JSON(http.StatusOK, "Product updated successfully.")
}

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := idToInt(idParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Input")
		return
	}
	_, err = db.DB.Exec(`delete from products where id = ?`, id)
	if err != nil {
		c.AbortWithStatusJSON(400, "Couldn't delete the product.")
	} else {
		c.JSON(http.StatusOK, "Product is successfully deleted.")
	}
}

func idToInt(idParam string) (int, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}
	return id, nil
}
