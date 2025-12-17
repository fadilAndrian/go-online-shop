package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"github.com/fadilAndrian/go-online-shop/internal/dto"
	"github.com/fadilAndrian/go-online-shop/internal/usecase/mock"
	"github.com/gin-gonic/gin"
)

func SetupProductRoute(h *ProductHandler) *gin.Engine {
	r := gin.New()

	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)

	return r
}

func ProductMockData() *mock.ProductUsecaseMock {
	return &mock.ProductUsecaseMock{
		ListFn: func() ([]domain.Product, error) {
			return []domain.Product{
				{
					Name:  "Dimsum",
					Price: 10000,
				},
				{
					Name:  "Sosis Solo",
					Price: 15000,
				},
			}, nil
		},

		GetFn: func(id int64) (*domain.Product, error) {
			if id != 1 {
				return nil, errors.New("Product not found")
			}

			return &domain.Product{
				Name:  "Dimsum",
				Price: 10000,
			}, nil
		},

		CreateFn: func(req dto.CreateProductRequest) (*domain.Product, error) {
			return &domain.Product{
				Name:  req.Name,
				Price: req.Price,
			}, nil
		},

		UpdateFn: func(id int64, req dto.UpdateProductRequest) (*domain.Product, error) {
			if id != 1 {
				return nil, errors.New("Product not found")
			}

			return &domain.Product{
				Name:  req.Name,
				Price: req.Price,
			}, nil
		},

		DeleteFn: func(id int64) error {
			if id != 1 {
				return errors.New("Product not found")
			}

			return nil
		},
	}
}

func SetupProductHandlerTest() (*gin.Engine, *httptest.ResponseRecorder) {
	mockData := ProductMockData()
	handler := NewProductHandler(mockData)
	router := SetupProductRoute(handler)

	w := httptest.NewRecorder()

	return router, w
}

func TestSuccessGetProductList(t *testing.T) {
	router, w := SetupProductHandlerTest()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status code 202, got %d", w.Code)
	}
}

func TestSuccessGetProductById(t *testing.T) {
	router, w := SetupProductHandlerTest()

	req := httptest.NewRequest(http.MethodGet, "/1", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status code 202, got %d", w.Code)
	}
}

func TestSuccessCreateProduct(t *testing.T) {
	router, w := SetupProductHandlerTest()

	body, _ := json.Marshal(gin.H{
		"name":  "Gohyong",
		"price": 20000,
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status code 202, got %d", w.Code)
	}
}

func TestSuccessUpdateProduct(t *testing.T) {
	router, w := SetupProductHandlerTest()

	body, _ := json.Marshal(gin.H{
		"name":  "Gohyong",
		"price": 21000,
	})

	req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status code 202, got %d", w.Code)
	}
}

func TestSuccessDeleteProduct(t *testing.T) {
	router, w := SetupProductHandlerTest()

	req := httptest.NewRequest(http.MethodDelete, "/1", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status code 202, got %d", w.Code)
	}
}
