package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mfritschdotgo/techchallengefase2/configs"

	_ "github.com/mfritschdotgo/techchallengefase2/docs"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/presenters"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/repositories"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/usecases"

	httpSwagger "github.com/swaggo/http-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title			Skina Lanches Management API
// @version		1.0
// @description	APIs for using the management system and sales orders
// @BasePath					/
func main() {
	config := configs.GetConfig()

	client, err := connectDatabase(config.MONGO_USER, config.MONGO_PASSWORD, config.MONGO_HOST, config.MONGO_PORT, config.MONGO_DATABASE)
	if err != nil {
		panic(err)
	}
	db := client.Database(config.MONGO_DATABASE)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryUseCases := usecases.NewCategory(categoryRepo)
	categoryPresenters := presenters.NewCategoryHandler(categoryUseCases)

	err = categoryUseCases.InitializeCategories(context.Background())
	if err != nil {
		panic(err)
	}

	productRepo := repositories.NewProductRepository(db)
	productUseCases := usecases.NewProduct(productRepo, categoryUseCases)
	productPresenters := presenters.NewProductHandler(productUseCases)

	clientRepo := repositories.NewClientRepository(db)
	clientUseCases := usecases.NewClient(clientRepo)
	clientPresenters := presenters.NewClientHandler(clientUseCases)

	orderRepo := repositories.NewOrderRepository(db)
	orderUseCases := usecases.NewOrder(orderRepo, clientUseCases, productUseCases)
	orderPresenters := presenters.NewOrderHandler(orderUseCases)

	paymentRepo := repositories.NewPaymentRepository(db)
	paymentUseCases := usecases.NewPaymentStatusUsecase(paymentRepo, orderUseCases)
	paymentPresenters := presenters.NewPaymentHandler(paymentUseCases)

	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/categories", func(r chi.Router) {
		r.Post("/", categoryPresenters.CreateCategory)
		r.Patch("/{id}", categoryPresenters.UpdateCategory)
		r.Put("/{id}", categoryPresenters.ReplaceCategory)
		r.Get("/{id}", categoryPresenters.GetCategoryByID)
		r.Get("/", categoryPresenters.GetCategories)
		r.Delete("/{id}", categoryPresenters.DeleteCategory)
	})

	r.Route("/products", func(r chi.Router) {
		r.Post("/", productPresenters.CreateProduct)
		r.Put("/{id}", productPresenters.ReplaceProduct)
		r.Patch("/{id}", productPresenters.UpdateProduct)
		r.Get("/{id}", productPresenters.GetProductByID)
		r.Get("/", productPresenters.GetProducts)
		r.Delete("/{id}", productPresenters.DeleteProduct)
	})

	r.Route("/clients", func(r chi.Router) {
		r.Post("/", clientPresenters.CreateClient)
		r.Get("/{cpf}", clientPresenters.GetClientByCPF)
	})

	r.Route("/orders", func(r chi.Router) {
		r.Get("/", orderPresenters.GetOrders)
		r.Get("/{id}", orderPresenters.GetOrderByID)
		r.Post("/", orderPresenters.CreateOrder)
		r.Patch("/{id}/{status}", orderPresenters.SetOrderStatus)
	})

	r.Route("/payment", func(r chi.Router) {
		r.Post("/", paymentPresenters.UpdatePaymentStatus)
		r.Get("/{id}", paymentPresenters.GeneratePayment)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("/docs/doc.json")))

	http.ListenAndServe(":9090", r)
}

func connectDatabase(user string, password string, host string, port string, dbname string) (*mongo.Client, error) {

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s", user, password, host, port, dbname, user)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
