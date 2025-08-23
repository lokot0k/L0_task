package handler

import (
	"L0_task/internal/models"
	"L0_task/pkg/cache"
	"L0_task/tools/order"
	"encoding/json"
	"fmt"
	"github.com/go-faker/faker/v4"
	"gorm.io/driver/postgres"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testDSN = "host=0.0.0.0 port=5432 user=user password=password dbname=appdb sslmode=disable" // креды к тестовой базе

func setupTestEnvironment(t *testing.T) (*gin.Engine, *gorm.DB, *cache.Cache) {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(postgres.Open(testDSN), &gorm.Config{})
	if err != nil {
		t.Fatalf("DB isnt' created: %v", err)
	}
	err = db.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Payment{}, &models.Item{})
	if err != nil {
		t.Fatalf("Failed creating schemas: %v", err)
	}
	cache := cache.NewCache(100)
	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("database", db)
		c.Set("cache", cache)
		c.Next()
	})

	return router, db, cache
}

func TestGetOrder_200(t *testing.T) {
	router, db, _ := setupTestEnvironment(t)

	testOrder := order.GenerateMockOrders(1)[0]
	jsonOrder, err := json.Marshal(testOrder) // сериализуем и десереализуем модель для дефолтного представления в бд
	if err != nil {
		t.Fatalf("Failed marshalling order: %v", err)
	}
	var generatedOrder models.Order
	err = json.Unmarshal(jsonOrder, &generatedOrder)
	if err != nil {
		t.Fatalf("Failed unmarshalling order: %v", err)
	}
	err = db.Create(&generatedOrder).Error
	if err != nil {
		t.Fatalf("Failed creating order: %v", err)
	}
	assert.NoError(t, err)
	router.GET("/orders/:orderID", GetOrder)
	req, err := http.NewRequest("GET", fmt.Sprintf("/orders/%s", generatedOrder.ID), nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Order
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, generatedOrder.ID, response.ID)
}

func TestGetOrder_Cache(t *testing.T) {
	router, _, cache := setupTestEnvironment(t)
	testOrder := order.GenerateMockOrders(1)[0]
	cacheableOrder := &models.CachableOrder{
		Order:    testOrder,
		LastUsed: time.Now(),
	}
	cache.Put(testOrder.ID, cacheableOrder)
	router.GET("/orders/:orderID", GetOrder)
	// для проверки кэша даже не надо класть ентри в бд
	req, err := http.NewRequest("GET", fmt.Sprintf("/orders/%s", testOrder.ID), nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Order
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, testOrder.ID, response.ID)
	assert.Equal(t, testOrder.TrackNumber, response.TrackNumber)
}

func TestGetOrder_404(t *testing.T) {
	router, _, _ := setupTestEnvironment(t)
	router.GET("/orders/:orderID", GetOrder)
	req, err := http.NewRequest("GET", fmt.Sprintf("/orders/%s", faker.UUIDHyphenated()), nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "order not found", response["message"])
}
