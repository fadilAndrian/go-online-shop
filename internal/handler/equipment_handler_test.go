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

func SetupEquipmentRoute(h *EquipmentHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)

	return r
}

func EuqipmentMockData() *mock.EquipmentUsecaseMock {
	return &mock.EquipmentUsecaseMock{
		ListFn: func() ([]domain.Equipment, error) {
			return []domain.Equipment{
				{
					Name:         "Baskom",
					CurrentStock: 1,
				},
				{
					Name:         "Nampan",
					CurrentStock: 1,
				},
			}, nil
		},

		GetFn: func(id int64) (*domain.Equipment, error) {
			if id != 1 {
				return nil, errors.New("equipment not found")
			}

			return &domain.Equipment{
				Name:         "Baskom",
				CurrentStock: 1,
			}, nil
		},

		CreateFn: func(req dto.CreateEquipmentRequest) (*domain.Equipment, error) {
			return &domain.Equipment{
				Name:         req.Name,
				CurrentStock: req.CurrentStock,
			}, nil
		},

		UpdateFn: func(id int64, req dto.UpdateEquipmentRequest) (*domain.Equipment, error) {
			if id != 1 {
				return nil, errors.New("equipment not found")
			}

			return &domain.Equipment{
				Name:         req.Name,
				CurrentStock: req.CurrentStock,
			}, nil
		},

		DeleteFn: func(id int64) error {
			if id != 1 {
				return errors.New("equipment not found")
			}

			return nil
		},
	}
}

func SetupEquipmentHandlerTest() (*gin.Engine, *httptest.ResponseRecorder) {
	mockData := EuqipmentMockData()

	handler := NewEquipmentHandler(mockData)
	router := SetupEquipmentRoute(handler)

	w := httptest.NewRecorder()

	return router, w
}

func TestGetAllEquipment(t *testing.T) {
	router, w := SetupEquipmentHandlerTest()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected 202, got %d", w.Code)
	}
}

func TestGetEquipmentById(t *testing.T) {
	router, w := SetupEquipmentHandlerTest()

	req := httptest.NewRequest(http.MethodGet, "/1", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected is 202, got %d", w.Code)
	}
}

func TestCreateEquipment_Success(t *testing.T) {
	router, w := SetupEquipmentHandlerTest()

	body, _ := json.Marshal(gin.H{
		"name":          "Baskom",
		"current_stock": 1,
	})

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status 202, got %d", w.Code)
	}
}

func TestUpdateEquipment(t *testing.T) {
	router, w := SetupEquipmentHandlerTest()

	body, _ := json.Marshal(gin.H{
		"name":          "Sepatula",
		"current_stock": 1,
	})

	req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status 202, got %d", w.Code)
	}
}

func TestDeleteEquipment(t *testing.T) {
	router, w := SetupEquipmentHandlerTest()

	req := httptest.NewRequest(http.MethodDelete, "/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status 202, got %d", w.Code)
	}
}
