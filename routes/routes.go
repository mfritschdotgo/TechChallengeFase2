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

	clientRepo := gateways.NewClientRepository(db)
	clientUseCases := usecases.NewClient(clientRepo)
	clientPresenter := presenters.NewClientPresenter()
	clientController := controllers.NewClientController(clientUseCases, clientPresenter)

	productRepo := gateways.NewProductRepository(db)
	productUseCases := usecases.NewProduct(productRepo, categoryUsecase)
	productPresenter := presenters.NewProductPresenter()
	productController := controllers.NewProductController(productUseCases, productPresenter)

	orderRepo := gateways.NewOrderRepository(db)
	orderUseCases := usecases.NewOrder(orderRepo, clientUseCases, productUseCases)
	orderPresenter := presenters.NewOrderPresenter()
	orderController := controllers.NewOrderController(orderUseCases, orderPresenter)

	paymentRepo := gateways.NewPaymentRepository(db)
	paymentUseCases := usecases.NewPaymentStatusUsecase(paymentRepo, orderUseCases)
	paymentPresenter := presenters.NewPaymentPresenter()
	paymentController := controllers.NewPaymentController(paymentUseCases, paymentPresenter)

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
		r.Post("/", productController.CreateProduct)
		r.Put("/{id}", productController.ReplaceProduct)
		r.Patch("/{id}", productController.UpdateProduct)
		r.Get("/{id}", productController.GetProductByID)
		r.Get("/", productController.GetProducts)
		r.Delete("/{id}", productController.DeleteProduct)
	})

	r.Route("/clients", func(r chi.Router) {
		r.Post("/", clientController.CreateClient)
		r.Get("/{cpf}", clientController.GetClientByCPF)
	})

	r.Route("/orders", func(r chi.Router) {
		r.Get("/", orderController.GetOrders)
		r.Get("/{id}", orderController.GetOrderByID)
		r.Post("/", orderController.CreateOrder)
		r.Patch("/{id}/{status}", orderController.SetOrderStatus)
	})

	r.Route("/payment", func(r chi.Router) {
		r.Post("/", paymentController.UpdatePaymentStatus)
		r.Get("/{id}", paymentController.GeneratePayment)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("/docs/doc.json")))

	return r
}
