package main

import (
	"net/smtp"
	//	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
	Address string
	Name    string
}
type Sender struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Message string `json:"message"`
}
type Response struct {
	Message string `json:"message"`
}

func (e *Email) BuildMessage() []byte {
	m := "Subject :" + e.Subject + "\n" + "from: " + e.Address + "\n" + e.Name + "\n" + e.Body + "\n" + e.Address
	fmt.Println(e)
	return []byte(m)
}

type SmtpServer struct {
	Host string
	Port string
}

func (s *SmtpServer) HostName() string {
	host := s.Host + ":" + s.Port
	return host
}

func SendEmail(w http.ResponseWriter, pass string, sender Sender) {
	email := Email{
		From:    "nader.special.api@gmail.com",
		To:      []string{"nader_atef80@outlook.com"},
		Subject: "Email from portfolio site",
		Body:    sender.Message,
		Address: sender.Address,
		Name:    sender.Name,
	}
	server := SmtpServer{
		Host: "smtp.gmail.com",
		Port: "587",
	}
	auth := smtp.PlainAuth("", email.From, pass, server.Host)
	err := smtp.SendMail(server.HostName(), auth, email.From, email.To, email.BuildMessage())
	if err != nil {
		log.Printf("smtp error: %s", err)
		ServeErr(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	resp := Response{
		"sent successfully",
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		ServeErr(w, http.StatusInternalServerError, err)
	}
}

func main() {
	port := os.Getenv("PORT")
	pass := os.Getenv("APIPASS")
	mux := http.NewServeMux()
	mux.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}
		decoder := json.NewDecoder(r.Body)
		var sender Sender
		err := decoder.Decode(&sender)
		if err != nil {
			ServeErr(w, http.StatusInternalServerError, err)
			return
		}
		SendEmail(w, pass, sender)
	})
	handler := cors.Default().Handler(mux)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		panic(err)
	}
}

func ServeErr(w http.ResponseWriter, status int, err error) {
	log.Println(err.Error())
	w.WriteHeader(status)
	w.Write([]byte("that should be here" + err.Error()))
}
