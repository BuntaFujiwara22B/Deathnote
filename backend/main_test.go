package main

import (
    "database/sql"
    "encoding/json"
    "io"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "deathback/config"
    "deathback/controllers"
    "github.com/gorilla/mux"
)

// Configurar la base de datos para pruebas
func setupTestDB() *sql.DB {
    return config.ConnectDB()
}

// Test para crear una víctima
func TestCreateVictima(t *testing.T) {
    db := setupTestDB()
    defer db.Close()

    body := `{"full_name":"John Doe","image_url":"http://example.com/image.jpg"}`
    req := httptest.NewRequest("POST", "/victimas", io.NopCloser(strings.NewReader(body)))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    controllers.CreateVictima(db, w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("Código de estado incorrecto: got %d, expected %d", res.StatusCode, http.StatusOK)
    }
}

// Test para obtener una víctima por ID
func TestGetVictima(t *testing.T) {
    db := setupTestDB()
    defer db.Close()

    req := httptest.NewRequest("GET", "/victimas/1", nil)
    w := httptest.NewRecorder()

    router := mux.NewRouter()
    router.HandleFunc("/victimas/{id}", func(w http.ResponseWriter, r *http.Request) {
        controllers.GetVictima(db, w, r)
    }).Methods("GET")

    router.ServeHTTP(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("Código de estado incorrecto: got %d, expected %d", res.StatusCode, http.StatusOK)
    }
}

// Test para actualizar una víctima
func TestUpdateVictima(t *testing.T) {
    db := setupTestDB()
    defer db.Close()

    body := `{"full_name":"Jane Doe","cause":"Accidente","image_url":"http://example.com/image2.jpg"}`
    req := httptest.NewRequest("PUT", "/victimas/1", io.NopCloser(strings.NewReader(body)))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    controllers.UpdateVictima(db, w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("Código de estado incorrecto: got %d, expected %d", res.StatusCode, http.StatusOK)
    }
}

// Test para eliminar una víctima
func TestDeleteVictima(t *testing.T) {
    db := setupTestDB()
    defer db.Close()

    req := httptest.NewRequest("DELETE", "/victimas/1", nil)
    w := httptest.NewRecorder()

    controllers.DeleteVictima(db, w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusNoContent {
        t.Errorf("Código de estado incorrecto: got %d, expected %d", res.StatusCode, http.StatusNoContent)
    }
}
