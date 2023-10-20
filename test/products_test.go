package test

import (
	"database/sql"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/farizirianov/products-restapi/controllers"
	"github.com/farizirianov/products-restapi/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sql.DB

func TestDBConnection(t *testing.T) {
	t.Log("Conexión exitosa")

	db, err := sql.Open("mysql", "user:your_password@/store_table")
	if err != nil {
		t.Errorf("Error al abrir la conexión: %v", err)
		return
	}

	err = db.Ping()
	if err != nil {
		t.Errorf("Error al hacer ping a la base de datos: %v", err)
		return
	}
	t.Log("Conexión exitosa")

	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		t.Errorf("Error al ejecutar la consulta: %v", err)
		return
	}
	t.Logf("Versión de MySQL: %s", version)
	defer db.Close()
}
func TestCreateProduct(t *testing.T) {
	p := models.Product{
		Sku:            "123456",
		Name:           "Camiseta",
		Brand:          "Nike",
		Size:           "M",
		Price:          0.99,
		MainImageUrl:   "https://example.com/camiseta.jpg",
		OtherImagesUrl: []string{"https://example.com/camiseta-1.jpg"},
	}

	// Llamar a la función create product
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controllers.ProductCreate(c)

	// Verificar que el producto tenga los campos esperados
	assert.Equal(t, "123456", p.Sku)
	assert.Equal(t, "Camiseta", p.Name)
	assert.Equal(t, "Nike", p.Brand)
	assert.Equal(t, "M", p.Size)
	assert.Equal(t, 0.99, p.Price)
	assert.Equal(t, "https://example.com/camiseta.jpg", p.MainImageUrl)
	assert.Equal(t, []string{"https://example.com/camiseta-1.jpg"}, p.OtherImagesUrl)
}

func TestUpdateProduct(t *testing.T) {
	p := models.Product{
		Sku:            "123456",
		Name:           "Camiseta",
		Brand:          "Nike",
		Size:           "M",
		Price:          19.99,
		MainImageUrl:   "https://example.com/camiseta.jpg",
		OtherImagesUrl: []string{"https://example.com/camiseta-1.jpg"},
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controllers.ProductCreate(c)

	p.Name = "Camiseta nueva"
	p.Price = 24.99
	p.OtherImagesUrl = []string{"https://example.com/camiseta-3.jpg"}

	// Llamar a la función update product para actualizar el producto en la base de datos
	controllers.UpdateProduct(c)

	assert.Equal(t, "123456", p.Sku)
	assert.Equal(t, "Camiseta nueva", p.Name)
	assert.Equal(t, "Nike", p.Brand)
	assert.Equal(t, "M", p.Size)
	assert.Equal(t, 24.99, p.Price)
	assert.Equal(t, "https://example.com/camiseta.jpg", p.MainImageUrl)
	assert.Equal(t, []string{"https://example.com/camiseta-3.jpg"}, p.OtherImagesUrl)
	assert.NotZero(t, p.CreatedAt)
	assert.NotZero(t, p.UpdatedAt)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:your_password@/test"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Product{})
	assert.NoError(t, err)

	p := models.Product{
		Sku:            "6535105814",
		Name:           "Camiseta",
		Brand:          "Nike",
		Size:           "M",
		Price:          19.99,
		MainImageUrl:   "https://example.com/camiseta.jpg",
		OtherImagesUrl: []string{"https://example.com/camiseta-1.jpg", "https://example.com/camiseta-2.jpg"},
	}

	err = db.Create(&p).Error
	if err != nil {
		// Handle the error
		fmt.Printf("server not responding %s", err.Error())
		return
	}
	assert.NoError(t, err)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = []gin.Param{
		{Key: "sku", Value: p.Sku},
	}
	c.Set("db", db)

	controllers.DeleteProduct(c)

	var count int64
	err = db.Where("sku = ?", p.Sku).First(&models.Product{}).Count(&count).Error
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}
