package main

// Imports
import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  "net/smtp"
  "os"
  "github.com/joho/godotenv"
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

// Process Functions START--------------------------------------------
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

func handleRequests() {
  http.HandleFunc("/", processor)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
// Process Functions END--------------------------------------------

// Driver Code
func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("[Error loading .env file]: ", err)
    os.Exit(1)
  }
  log.Println("server started at :8000")
  handleRequests()
}
