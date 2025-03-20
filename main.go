package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var puanTablosu = make(map[int]int)
var mu sync.Mutex

func baslangicPuaniAyarla(oyuncuSayisi int) {
	mu.Lock()
	defer mu.Unlock()

	for i := 1; i <= oyuncuSayisi; i++ {
		puanTablosu[i] = 500
	}
}

func puanTablosuHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(puanTablosu)
}

func puanGuncelleHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	oyuncuID, err1 := strconv.Atoi(r.URL.Query().Get("oyuncu"))
	puanDegisimi, err2 := strconv.Atoi(r.URL.Query().Get("puan"))

	if err1 != nil || err2 != nil {
		http.Error(w, "Geçersiz oyuncu ID veya puan değişimi.", http.StatusBadRequest)
		return
	}

	if _, exists := puanTablosu[oyuncuID]; !exists {
		http.Error(w, "Oyuncu bulunamadı.", http.StatusNotFound)
		return
	}

	puanTablosu[oyuncuID] += puanDegisimi
	fmt.Fprintf(w, "Oyuncu %d'in puanı %d olarak güncellendi.\n", oyuncuID, puanTablosu[oyuncuID])
}

func main() {
	var oyuncuSayisi int
	fmt.Print("Oyuncu sayısını girin: ")
	fmt.Scanln(&oyuncuSayisi)

	baslangicPuaniAyarla(oyuncuSayisi)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/puanlar", puanTablosuHandler)
	http.HandleFunc("/puanGuncelle", puanGuncelleHandler)

	fmt.Println("Sunucu 8080 portunda çalışıyor...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
