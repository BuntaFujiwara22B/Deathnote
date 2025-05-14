package utils

import (
	"log"
	"time"
	"context"

	"deathback/config"
	"deathback/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ScheduleDeath(person models.Person) {
	if person.ImageURL == "" {
		log.Println("Persona sin foto, no puede morir.")
		return
	}

	go func() {
		time.Sleep(40 * time.Second)

		client := config.ConnectDB()
		collection := client.Database("deathdb").Collection("people")

		var updatedPerson models.Person
		_ = collection.FindOne(context.TODO(), bson.M{"_id": person.ID}).Decode(&updatedPerson)

		if updatedPerson.IsDead {
			return
		}

		if updatedPerson.Cause == "" {
			// Muerte por ataque al corazón
			now := time.Now()
			update := bson.M{
				"$set": bson.M{
					"cause":     "Ataque al corazón",
					"isDead":    true,
					"deathTime": now,
				},
			}
			_, err := collection.UpdateByID(context.TODO(), person.ID, update)
			if err != nil {
				log.Println("Error al programar muerte:", err)
			}
			log.Println("Murió por ataque al corazón:", updatedPerson.FullName)
		} else {
			// Espera 6 minutos y 40 segundos para que se registren los detalles
			time.Sleep(6*time.Minute + 40*time.Second)

			var updatedAgain models.Person
			_ = collection.FindOne(context.TODO(), bson.M{"_id": person.ID}).Decode(&updatedAgain)

			if updatedAgain.Details != "" && !updatedAgain.IsDead {
				time.Sleep(40 * time.Second)
				now := time.Now()
				update := bson.M{
					"$set": bson.M{
						"isDead":    true,
						"deathTime": now,
					},
				}
				_, err := collection.UpdateByID(context.TODO(), person.ID, update)
				if err != nil {
					log.Println("Error al ejecutar muerte tras detalles:", err)
				}
				log.Println("Murió tras registrar detalles:", updatedAgain.FullName)
			}
		}
	}()
}
