package handlers

import (
	"encoding/json"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const indexFilePath = "index.html"

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	data, err := os.ReadFile(indexFilePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	w.Write(data)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	convertedData := service.Convert(string(data))

	results, err := os.Create(generateFileName())
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	_, err = results.WriteString(convertedData)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"original": string(data),
		"result":   convertedData,
		"file":     results.Name(),
	}

	resp, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

// generateFileName генерирует название для файла, возвращает строку
func generateFileName() string {
	s := "result_" + time.Now().Format("2006-01-02_15-04-05") + ".txt"
	return s
}
