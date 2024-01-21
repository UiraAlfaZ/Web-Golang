package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Abydos struct {
	Name   string
	Age    int
	Hobby  string
	Height float64
	Year   string
}

type Gehenna struct {
	Name   string
	Age    int
	Hobby  string
	Height float64
	Year   string
}
type Millenium struct {
	Name   string
	Age    int
	Hobby  string
	Height float64
	Year   string
}

type Trinity struct {
	Name   string
	Age    int
	Hobby  string
	Height float64
	Year   string
}

type Hyakkiyako struct {
	Name   string
	Age    int
	Hobby  string
	Height float64
	Year   string
}

// Abydos
func HandlerAbydos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { //validari method dengan menggunakan "GET" untuk eksekusi data didalam
		taskForce := []Abydos{
			{"Takanashi Hoshino", 17, "Nap", 1.45, "Third year"},
			{"Izayoi Nonomi", 16, "Shopping", 1.60, "Second year"},
			{"Sunaookami Shiroko", 16, "Jogging, Cycling", 1.56, "Second year"},
			{"Kuromi Serika", 15, "Saving money, Partime-job", 1.53, "First year"},
			{"Okusora Ayane", 15, "Accounting, Antique Collecting", 1.53, "First year"},
		}
		res, err := json.Marshal(taskForce) //mengubah struct mejadi json, contoh struct TaskForce yang akan dirubah
		if err != nil {
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError) //kondisi saat gagal konversi ke json
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed) //saat metod lain gagal konversi ke json
	}
}

// Gehenna
func HandlerGehenna(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		PrefectTeam := []Gehenna{
			{"Sorasaki Hina", 17, "Sleeping, Resting", 1.42, "Third year"},
			{"Amou Ako", 17, "Hina", 1.65, "Third year"},
			{"Shiromi Iori", 16, "Patrolling, Shouting", 1.57, "Second year"},
			{"Hinomiya Chinatsu", 15, "Reading philosphy book", 1.59, "First year"},
		}
		res, err := json.Marshal(PrefectTeam)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

// Millenium
func HandlerMillenium(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Seminar := []Millenium{
			{"Ushio Noa", 16, "Reading, Recitation", 1.61, "Second year"},
			{"Hayase Yuuka", 16, "Calculation", 1.56, "Second year"},
			{"Kurosaki Koyuki", 15, "Test one's luck", 1.49, "First year"},
		}
		res, err := json.Marshal(Seminar)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

// Trinity
func HandlerTrinity(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		SLD := []Trinity{
			{"Urawa Hanako", 16, "Wandering around", 1.61, "Second year"},
			{"Ajitani Hifumi", 16, "Collecting cute things", 1.58, "Second year"},
			{"Shirasu Azusa", 16, "None", 1.49, "Second year"},
			{"Shimoe Koharu", 15, "Fantasies, Delusions", 1.48, "First year"},
		}
		res, err := json.Marshal(SLD)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
func HandlerHyakkiyako(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ETS := []Trinity{
			{"Kasuga Tsubaki", 16, "Afternoon nap", 1.62, "Second year"},
			{"Mizuha Mimori", 16, "Shoujo manga, General housework", 1.56, "Second year"},
			{"Isami Kaede", 15, "Catching beetles", 1.42, "First year"},
		}
		res, err := json.Marshal(ETS)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke JSON", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
func main() {
	http.HandleFunc("/Abydos", HandlerAbydos) //tanpa nama function
	http.HandleFunc("/Gehenna", HandlerGehenna)
	http.HandleFunc("/Millenium", HandlerMillenium)
	http.HandleFunc("/Trinity", HandlerTrinity)
	http.HandleFunc("/Hyakkiyako", HandlerHyakkiyako)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
