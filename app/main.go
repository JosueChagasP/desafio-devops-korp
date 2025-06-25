package main
import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//CMD p resposta Json
type Response struct {
	Nome string `json:"nome"`
	Horario string `json:"horario"`
}

//Hdl for all requisicao
func handler(w http.ResponseWriter, r *http.Request) {
// manipula headers p json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

//pegar hr atual em utc
	now :=time.Now().UTC()
	horarioFormatado := now.Format("2006-01-02 15:04:05 UTC")

//gera resposta
	response := Response{
	Nome: "Projeto Korp",
	Horario: horarioFormatado,
}

//transforma em json
	jsonData, err := json.Marshal(response)
	if err != nil {
	http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
	return
}

//recebe resposta
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

// log da requisicao 
	log.Printf("Requisicao atendida: %s %s - Horario: %s", r.Method, r.URL.Path, horarioFormatado)
}
	func main() {
//conf rota 
	http.HandleFunc("/", handler)
//conf server
	port := ":8080"
	log.Printf("Servidor http-server-projeto-korp iniciado na porta 8080")
	log.Printf("Endpoint: http://localhost:8080")
	log.Printf("Retornando horario UTC dinamicamente")

// Iniciar servidor
	if err := http.ListenAndServe(port, nil); err != nil {
	log.Fatalf ("X Erro ao iniciar servidor: %v", err)
	}
}
