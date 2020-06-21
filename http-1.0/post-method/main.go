package main

import (
  "log"
  "net/http"
  "net/url"
)

func main() {
  values := url.Values{
    "test" : {"value"},
  }
  resp, _ := http.PostForm("http://localhost:8080", values)
  log.Println("Status: ", resp.Status)
}
