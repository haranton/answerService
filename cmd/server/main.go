package models

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет! Это простой HTTP сервер на Go 🚀")
}

func main() {

	// logger
	// config
	// db connect
	// migrate
	//

	http.HandleFunc("/", handler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Errorf("failed run server err: %s", err)
	}

}
