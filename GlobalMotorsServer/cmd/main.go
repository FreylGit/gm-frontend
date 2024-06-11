package main

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/auto", GetCars)
	r.Get("/auto/{id}", GetCar)
	http.ListenAndServe(":3000", r)
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	cars := generateFakeAutoData(20)
	json, _ := json.Marshal(cars)
	w.Write(json)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	car := Auto{
		ID:          id,
		Title:       gofakeit.CarModel(),
		Price:       gofakeit.Price(10000, 50000),
		Address:     gofakeit.Address().Address,
		Description: gofakeit.Sentence(5),
		ImagePath:   "https://car-images.bauersecure.com/wp-images/2697/bmwi4_029.jpg",
	}
	json, _ := json.Marshal(car)
	w.Write(json)
}

func generateFakeAutoData(count int) []Auto {
	var autos []Auto
	for i := 0; i < count; i++ {
		auto := Auto{
			ID:          int64(i + 1),
			Title:       gofakeit.CarModel(),
			Price:       gofakeit.Price(10000, 50000),
			Address:     gofakeit.Address().Address,
			Description: gofakeit.Sentence(5),
			ImagePath:   "https://car-images.bauersecure.com/wp-images/2697/bmwi4_029.jpg",
		}
		autos = append(autos, auto)
	}

	return autos
}

type Auto struct {
	ID          int64
	Title       string
	Price       float64
	Address     string
	Description string
	ImagePath   string
}
