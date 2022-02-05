package main

import (
  contact "./contact"
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
  log.Println("server started at :8000")
  contact.HandleRequests()
}
