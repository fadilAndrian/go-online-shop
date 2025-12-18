package main

import (
	"log"

	"github.com/fadilAndrian/go-online-shop/internal/config"
	"github.com/fadilAndrian/go-online-shop/internal/handler"
	"github.com/fadilAndrian/go-online-shop/internal/repository"
	"github.com/fadilAndrian/go-online-shop/internal/router"
	"github.com/fadilAndrian/go-online-shop/internal/usecase"
	"github.com/fadilAndrian/go-online-shop/pkg/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("failed to connect database, err : ", err)
	}

	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	equipmentRepository := repository.NewEquipmentRepository(db)
	equipmentUsecase := usecase.NewEquipmentUsecase(equipmentRepository)
	equipmentHandler := handler.NewEquipmentHandler(equipmentUsecase)

	ingredientRepository := repository.NewIngredientRepository(db)
	ingredientUsecase := usecase.NewIngredientUsecase(ingredientRepository)
	ingredientHandler := handler.NewIngredientHandler(ingredientUsecase)

	r := router.Routes(
		productHandler,
		equipmentHandler,
		ingredientHandler,
	)
	r.Run()
}
