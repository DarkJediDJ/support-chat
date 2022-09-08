package sender

import (
	"encoding/json"
	"net/http"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Handler struct {
	conn *kafka.Conn
}

type Message struct {
	Id   int    `json:"ID"`
	Text string `json:"message"`
}

func Init(c *kafka.Conn) *Handler {

	return &Handler{
		conn: c,
	}
}

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var mes Message

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&mes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rBody, err := json.Marshal(mes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	_, err = h.conn.Write(rBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
