package handler

import (
	"strconv"

	"github.com/asadbek21coder/eleanshop/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	// router.Use(cors.New(config))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")

	{
		auth.POST("/signup", h.signup)
		auth.POST("/signin", h.signin)

	}

	product := router.Group("/product")
	{

		product.POST("/", h.isAdmin, h.createProduct)
		product.GET("/", h.getAllProducts)
		product.GET("/:id", h.getProductById)
		product.PUT("/:id", h.isAdmin, h.updateProduct)
		product.DELETE("/:id", h.isAdmin, h.deleteProduct)
	}

	admin := router.Group("/admin", h.isAdmin)
	{
		admin.PUT("/set-admin")

		category := admin.Group("/category")
		{
			category.GET("/", h.getAllCategories)
			category.POST("/", h.setAdmin, h.createCategory)
			category.GET("/:id", h.getCategoryById)
			category.PUT("/:id", h.setAdmin, h.updateCategory)
			category.DELETE("/:id", h.setAdmin, h.deleteCategory)
		}

		sizes := admin.Group("/sizes")
		{
			sizes.POST("/", h.setAdmin, h.createSize)
			sizes.GET("/", h.getAllSizes)
			sizes.GET("/:id", h.getSizesById)
			sizes.DELETE("/:id", h.setAdmin, h.deleteSize)
		}
	}

	feedback := router.Group("/feedback")
	{
		feedback.POST("/", h.userIdentity, h.createFeedback)
		feedback.GET("/", h.getAllFeedbacks)
		feedback.GET("/:id", h.getFeedbackById)
		feedback.PUT("/:id", h.userIdentity, h.updateFeedback)
		feedback.DELETE("/:id", h.userIdentity, h.deleteFeedback)
	}

	question := router.Group("/question")
	{
		question.POST("/", h.userIdentity, h.createQuestion)
		question.GET("/", h.getAllQuestions)
		question.GET("/:id", h.getQuestionById)
		question.PUT("/:id", h.userIdentity, h.updateQuestion)
		question.DELETE("/:id", h.userIdentity, h.deleteQuestion)
	}
	router.Static("assets", "./assets")
	return router
}

func (h *Handler) parseOffsetQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("offset", "0"))
}

func (h *Handler) parseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", "10"))
}
