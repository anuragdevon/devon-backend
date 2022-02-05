package main

// Imports
import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

// Driver Code
func main() {

  url := "http://127.0.0.1:8080/"
  method := "POST"

  payload := strings.NewReader(`{
    "email": "newemail@gmail.com", 
    "message": "mashiro"
  }`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}