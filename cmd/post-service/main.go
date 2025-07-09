package postservice

import (
	"log"
	"net/http"
	"os"

	"github.com/dev-akshaygupta/thread-it/internal/domain/post/delivery"
	"github.com/dev-akshaygupta/thread-it/internal/domain/post/handler"
	"github.com/dev-akshaygupta/thread-it/internal/domain/post/repository"
	"github.com/dev-akshaygupta/thread-it/internal/domain/post/service"
	"github.com/dev-akshaygupta/thread-it/pkg/db"
)

func main() {
	dbConn := db.NewPostgresDb()
	repo := repository.NewPostgresRepo(dbConn)
	svc := service.NewPostService(repo)
	handler := &handler.PostHandler{Service: svc}
	r := delivery.NewRouter(handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Post service listening on: %s", port)
	http.ListenAndServe(":"+port, r)
}
