package main

import (
  contact "devon-backend/main/contact"
	"github.com/joho/godotenv"
  "log"
  "os"
)


// Driver Code
func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("[Error loading .env file]: ", err)
    os.Exit(1)
  }
  log.Println("server started at :8080")
  contact.HandleRequests()
}
