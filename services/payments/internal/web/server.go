package web

import (
	"eda-payments/internal"
	"log"
	"net/http"
)

type Server struct {
	Kakfa *internal.KafkaClient
}

func NewServer(kf *internal.KafkaClient) *Server {
	return &Server{Kakfa: kf}
}

func (s Server) Payment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get Token -> verify claims

	err := s.Kakfa.SendInventory()

	if err != nil {
		http.Error(w, "Failed to send inventory message", http.StatusInternalServerError)
		return
	}

	err = s.Kakfa.SendNotification()

	if err != nil {
		http.Error(w, "Failed to send notification message", http.StatusInternalServerError)
		return
	}

	output := "Payment processed successfully. Inventory and Notification messages sent to Kafka."

	log.Println(output)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Done"))
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mux := http.NewServeMux()

	mux.HandleFunc("/pay", s.Payment)

	mux.ServeHTTP(w, r)

}
