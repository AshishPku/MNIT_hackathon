package main

import (
	"context"
	"easyjobBackend/db"
	"easyjobBackend/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":10000"
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		panic(err)
	}
	jobStore := db.NewMongoJobStore(client)
	jobHandler := handlers.NewJobHandler(jobStore)
	mentorStore := db.NewMongoMentorStore(client)
	mentorHandler := handlers.NewMentorHandler(mentorStore)

	app := fiber.New(config)
	app.Use(cors.New())
	// apiv1 := app.Group("/api/v1", middleware.JwtAuth)

	//PROTECTED ROUTES
	// apiv1.Get("/accountposts", postHandler.HandleGetPostsByUser)
	// apiv1.Post("/post", postHandler.HandleInsertPost)
	// apiv1.Post("/postImages", imageHandler.HandlePostImage)
	// apiv1.Post("/search", postHandler.HandleSearchUser)

	//UNPROTECTED ROUTES

	// app.Post("/signup", userHandler.HandleCreateUser)
	// app.Post("/login", userHandler.HandleLoginUser)
	// app.Static("/", "./public/build")
	// app.Static("/static", "./files")
	app.Get("/home", jobHandler.HandleGetAllJobs)
	app.Get("/search", jobHandler.HandleGetJobsByFilter)
	app.Get("/mentors", mentorHandler.HandleGetAllMentors)

	app.Listen(port)

}
