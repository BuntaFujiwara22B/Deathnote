package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
	
	"github.com/gorilla/mux"
	"github.com/BuntaFujiwara22B/deathback/models"
)

// Crear una víctima en la base de datos
func CreateVictima(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var victima models.Victima
	if err := json.NewDecoder(r.Body).Decode(&victima); err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validaciones
	if victima.FullName == "" {
		http.Error(w, "Se requiere nombre completo", http.StatusBadRequest)
		return
	}
	if victima.ImageURL == "" {
		http.Error(w, "Se requiere URL de imagen", http.StatusBadRequest)
		return
	}

	// Establecer tiempos
	victima.CreatedAt = time.Now()
	victima.DeathTime = victima.CreatedAt.Add(40 * time.Second)
	victima.CauseAdded = false
	victima.DetailsAdded = false
	victima.IsDead = false

	// Insertar en base de datos
	_, err := db.Exec(`
		INSERT INTO victimas 
		(full_name, cause, details, created_at, death_time, image_url, cause_added, details_added, is_dead) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, 
		victima.FullName, victima.Cause, victima.Details, victima.CreatedAt, 
		victima.DeathTime, victima.ImageURL, victima.CauseAdded, victima.DetailsAdded, victima.IsDead)

	if err != nil {
		http.Error(w, "Error al guardar víctima: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(victima)
}

// Obtener víctima por ID
func GetVictima(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var victima models.Victima

	err := db.QueryRow(`
		SELECT id, full_name, cause, details, created_at, death_time, 
		       image_url, cause_added, details_added, is_dead 
		FROM victimas WHERE id = $1`, id).
		Scan(&victima.ID, &victima.FullName, &victima.Cause, &victima.Details, 
		     &victima.CreatedAt, &victima.DeathTime, &victima.ImageURL, 
		     &victima.CauseAdded, &victima.DetailsAdded, &victima.IsDead)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Víctima no encontrada", http.StatusNotFound)
		} else {
			http.Error(w, "Error al obtener víctima: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(victima)
}

// Listar todas las víctimas
func ListVictimas(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query(`
		SELECT id, full_name, cause, details, created_at, death_time, 
		       image_url, cause_added, details_added, is_dead 
		FROM victimas ORDER BY created_at DESC`)
	if err != nil {
		http.Error(w, "Error al obtener víctimas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var victimas []models.Victima
	for rows.Next() {
		var v models.Victima
		if err := rows.Scan(&v.ID, &v.FullName, &v.Cause, &v.Details, 
			&v.CreatedAt, &v.DeathTime, &v.ImageURL, 
			&v.CauseAdded, &v.DetailsAdded, &v.IsDead); err != nil {
			http.Error(w, "Error al leer datos: "+err.Error(), http.StatusInternalServerError)
			return
		}
		victimas = append(victimas, v)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Error al procesar resultados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(victimas)
}

// Actualizar causa de muerte
func UpdateCause(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var data struct {
		Cause string `json:"cause"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar tiempo límite (40 segundos)
	var createdAt time.Time
	err := db.QueryRow("SELECT created_at FROM victimas WHERE id = $1", id).Scan(&createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Víctima no encontrada", http.StatusNotFound)
		} else {
			http.Error(w, "Error al verificar víctima: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if time.Since(createdAt) > 40*time.Second {
		http.Error(w, "Tiempo para especificar causa ha expirado (máximo 40 segundos)", http.StatusBadRequest)
		return
	}

	// Actualizar en base de datos
	_, err = db.Exec(`
		UPDATE victimas 
		SET cause = $1, 
			death_time = CASE 
				WHEN $1 != '' THEN created_at + INTERVAL '6 minutes 40 seconds'
				ELSE created_at + INTERVAL '40 seconds'
			END,
			cause_added = TRUE
		WHERE id = $2`, data.Cause, id)

	if err != nil {
		http.Error(w, "Error al actualizar causa: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Causa actualizada correctamente"})
}

// Actualizar detalles de muerte
func UpdateDetails(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var data struct {
		Details string `json:"details"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar que existe causa y no ha expirado el tiempo (6m40s)
	var createdAt time.Time
	var hasCause bool
	err := db.QueryRow(`
		SELECT created_at, cause_added 
		FROM victimas WHERE id = $1`, id).
		Scan(&createdAt, &hasCause)
	
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Víctima no encontrada", http.StatusNotFound)
		} else {
			http.Error(w, "Error al verificar víctima: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if !hasCause {
		http.Error(w, "Primero debe especificar una causa de muerte", http.StatusBadRequest)
		return
	}

	if time.Since(createdAt) > (6*time.Minute + 40*time.Second) {
		http.Error(w, "Tiempo para especificar detalles ha expirado (máximo 6 minutos 40 segundos)", http.StatusBadRequest)
		return
	}

	// Actualizar en base de datos
	_, err = db.Exec(`
		UPDATE victimas 
		SET details = $1, 
			death_time = created_at + INTERVAL '7 minutes 20 seconds',
			details_added = TRUE
		WHERE id = $2`, data.Details, id)

	if err != nil {
		http.Error(w, "Error al actualizar detalles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Detalles actualizados correctamente"})
}