package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/mfritschdotgo/techchallengefase2/docs"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/controllers"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/gateways"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/handlers"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/presenters"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/repositories"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/usecases"
	"go.mongodb.org/mongo-driver/mongo"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			Skina Lanches Management API
// @version		1.0
// @description	APIs for using the management system and sales orders
// @BasePath	/
func SetupRoutes(db *mongo.Database) *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Repositories and Use Cases
	categoryRepo := repositories.NewMongoDBCategoryRepository(db)
	categoryGateway := gateways.NewCategoryGateway(categoryRepo)
	categoryUsecase := usecases.NewCategoryUseCase(categoryGateway)
	categoryController := controllers.NewCategoryController(categoryUsecase)
	categoryPresenter := presenters.NewCategoryPresenter()
	categoryHandler := handlers.NewCategoryHandler(categoryController, categoryPresenter)

	clientRepo := repositories.NewMongoDBClientRepository(db)
	clientGateway := gateways.NewClientGateway(clientRepo)
	clientUseCases := usecases.NewClient(clientGateway)
	clientPresenter := presenters.NewClientPresenter()
	clientController := controllers.NewClientController(clientUseCases)
	clientHandler := handlers.NewClientHandler(clientController, clientPresenter)

	productRepo := repositories.NewMongoDBProductRepository(db)
	productGateway := gateways.NewProductGateway(productRepo)
	productUseCases := usecases.NewProduct(productGateway, categoryUsecase)
	productPresenter := presenters.NewProductPresenter()
	productController := controllers.NewProductController(productUseCases, productPresenter)
	productHandler := handlers.NewProductHandler(productController, productPresenter)

	orderRepo := repositories.NewMongoDBOrderRepository(db)
	orderGateway := gateways.NewOrderGateway(orderRepo)
	orderUseCases := usecases.NewOrder(orderGateway, clientUseCases, productUseCases)
	orderPresenter := presenters.NewOrderPresenter()
	orderController := controllers.NewOrderController(orderUseCases, orderPresenter)
	orderHandler := handlers.NewOrderHandler(orderController, orderPresenter)

	paymentRepo := repositories.NewMongoDBPaymentRepository(db)
	paymentGateway := gateways.NewPaymentGateway(paymentRepo)
	paymentUseCases := usecases.NewPaymentStatusUsecase(paymentGateway, orderUseCases)
	paymentPresenter := presenters.NewPaymentPresenter()
	paymentController := controllers.NewPaymentController(paymentUseCases, paymentPresenter)
	paymentHandler := handlers.NewPaymentHandler(paymentController, paymentPresenter)

	// Configuração das rotas
	r.Route("/categories", func(r chi.Router) {
		r.Post("/", categoryHandler.CreateCategory)
		r.Patch("/{id}", categoryHandler.UpdateCategory)
		r.Put("/{id}", categoryHandler.ReplaceCategory)
		r.Get("/{id}", categoryHandler.GetCategoryByID)
		r.Get("/", categoryHandler.GetCategories)
		r.Delete("/{id}", categoryHandler.DeleteCategory)
	})

	r.Route("/products", func(r chi.Router) {
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.ReplaceProduct)
		r.Patch("/{id}", productHandler.UpdateProduct)
		r.Get("/{id}", productHandler.GetProductByID)
		r.Get("/", productHandler.GetProducts)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/clients", func(r chi.Router) {
		r.Post("/", clientHandler.CreateClient)
		r.Get("/{cpf}", clientHandler.GetClientByCPF)
	})

	r.Route("/orders", func(r chi.Router) {
		r.Get("/", orderHandler.GetOrders)
		r.Get("/{id}", orderHandler.GetOrderByID)
		r.Post("/", orderHandler.CreateOrder)
		r.Patch("/{id}/{status}", orderHandler.SetOrderStatus)
	})

	r.Route("/payment", func(r chi.Router) {
		r.Post("/", paymentHandler.UpdatePaymentStatus)
		r.Get("/{id}", paymentHandler.GeneratePayment)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("/docs/doc.json")))

	return r
}
