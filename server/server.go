package main

import (

	// "fmt"
	"kns_server/graph"
	"kns_server/graph/generated"

	// "kns_server/graph/model"

	// "kns_server/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var defaultPort = "8080"
var db *gorm.DB

// func initDB() {
//     var err error

// 	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
//     db, err = gorm.Open("mysql", dataSourceName)

//     if err != nil {
//         fmt.Println(err)
//         panic("failed to connect database")
//     }

//     db.LogMode(true)

//     // Create the database. This is a one-time step.
//     // Comment out if running multiple times - You may see an error otherwise
//     db.Exec("CREATE DATABASE test_db")
//     db.Exec("USE test_db")

//     // Migration to create tables for Order and Item schema
//     db.AutoMigrate(&model.Order{}, &model.Item{})
// }

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", addCORSHeaders(srv))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// model.Facebook()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func addCORSHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow any origin to make requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow the Content-Type header in requests
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Allow GET, POST, PUT, DELETE, and OPTIONS methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// If this is a preflight request, return immediately with a 200 status
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		// Call the actual handler
		h.ServeHTTP(w, r)
	})
}
