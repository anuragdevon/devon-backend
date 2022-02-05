package contact

// Imports
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"fmt"
)

// Global Variables and datatypes
type ContactData struct {
	Email   string `json:"email"`
	Message string `json:"message"`
  
}
  
type JsonResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
  
// Contact Process Functions START--------------------------------------------
func processor(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	response := JsonResponse{}
	if err != nil {
		panic(err)
	}
	var contactdata ContactData
	json.Unmarshal(body, &contactdata)
  
	err = contact(contactdata)
	if err != nil {
	  log.Println("[failed!]:", err)
	  response = JsonResponse{Type: "Error!", Message: "Unable to send email!"}
  
	} else {
	  response = JsonResponse{Type: "Success!", Message: "Mail Sent Successfully!"}
	  log.Println("success!")
	}
	json.NewEncoder(w).Encode(response)
}
  
func contact(contactdata ContactData) error {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
  
	to := os.Getenv("EMAIL_TO")
	subject := "Contact From: " + contactdata.Email
	msg := []byte(createMessage(contactdata.Email, contactdata.Message, to, subject))
  
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
  
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, msg)
	
	return err
}
  
func createMessage(email string, message string, to string, subject string) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("To: %s\r\n", to)
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += fmt.Sprintf("\r\n%s\r\n", message)
  
	return msg
}
  
func HandleRequests() {
	http.HandleFunc("/", processor)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
  // Contact Process Functions END--------------------------------------------
  