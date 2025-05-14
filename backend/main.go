package main

import (
	
	"fmt"
	"log"
	"net/http"
	"github.com/BuntaFujiwara22B/deathback/config"
	"github.com/BuntaFujiwara22B/deathback/routes"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

func main() {
	// Conexión a PostgreSQL
	db := config.ConnectDB()
	defer db.Close()

	// Inicializar DB
	if err := config.InitDB(db); err != nil {
		log.Fatalf("Error inicializando la base de datos: %v", err)
	}
	fmt.Println("✅ Base de datos inicializada")

	// Configurar rutas con CORS
	router := routes.SetupRoutes(db)
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:5173"}) // Ajusta si usas otro puerto
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})
	// Iniciar servidor
	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, origins, methods)(router)))
}

