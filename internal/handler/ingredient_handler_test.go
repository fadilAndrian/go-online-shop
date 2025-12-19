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

func SetupIngredientRoute(h *IngredientHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.GET("/", h.GetAllEquipment)
	r.GET("/:id", h.GetEquipmentById)
	r.POST("/", h.CreateEquipment)
	r.PUT("/:id", h.UpdateEquipment)
	r.DELETE("/:id", h.DeleteEquipment)

	return r
}

func IngredientMockData() *mock.IngredientUsecaseMock {
	return &mock.IngredientUsecaseMock{
		ListFn: func() ([]domain.Ingredient, error) {
			return []domain.Ingredient{
				{
					Name:         "Bawang",
					Unit:         "gr",
					InitialStock: 0,
					CurrentStock: 10,
				},
				{
					Name:         "Tomat",
					Unit:         "gr",
					InitialStock: 0,
					CurrentStock: 50,
				},
			}, nil
		},

		GetFn: func(id int64) (*domain.Ingredient, error) {
			if id != 1 {
				return nil, errors.New("ingredient not found")
			}

			return &domain.Ingredient{
				Name:         "Bawang",
				Unit:         "gr",
				InitialStock: 0,
				CurrentStock: 10,
			}, nil
		},

		CreateFn: func(req dto.IngredientCreateRequest) (*domain.Ingredient, error) {
			return &domain.Ingredient{
				Name:         req.Name,
				Unit:         req.Unit,
				InitialStock: req.InitialStock,
				CurrentStock: req.CurrentStock,
			}, nil
		},

		UpdateFn: func(id int64, req dto.IngredientUpdateRequest) (*domain.Ingredient, error) {
			if id != 1 {
				return nil, errors.New("ingredient not found")
			}

			return &domain.Ingredient{
				Name:         req.Name,
				Unit:         req.Unit,
				CurrentStock: req.CurrentStock,
			}, nil
		},

		DeleteFn: func(id int64) error {
			if id != 1 {
				return errors.New("ingredient not found")
			}

			return nil
		},
	}
}

func SetupIngredientHandlerTest() (*gin.Engine, *httptest.ResponseRecorder) {
	mockData := IngredientMockData()

	handler := NewIngredientHandler(mockData)
	router := SetupIngredientRoute(handler)

	w := httptest.NewRecorder()

	return router, w
}

func TestGetAllIngredient(t *testing.T) {
	router, w := SetupIngredientHandlerTest()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestGetIngredientById(t *testing.T) {
	router, w := SetupIngredientHandlerTest()

	req := httptest.NewRequest(http.MethodGet, "/1", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected is 200, got %d", w.Code)
	}
}

func TestCreateIngredient_Success(t *testing.T) {
	router, w := SetupIngredientHandlerTest()

	body, _ := json.Marshal(gin.H{
		"name":          "Bawang",
		"unit":          "gr",
		"initial_stock": 0,
		"current_stock": 100,
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestUpdateIngredient(t *testing.T) {
	router, w := SetupIngredientHandlerTest()

	body, _ := json.Marshal(gin.H{
		"name":          "Bawang",
		"unit":          "gr",
		"current_stock": 100,
	})

	req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d, err: %s", w.Code, w.Body.String())
	}
}

func TestDeleteIngredient(t *testing.T) {
	router, w := SetupIngredientHandlerTest()

	req := httptest.NewRequest(http.MethodDelete, "/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d, err: %s", w.Code, w.Body.String())
	}
}
