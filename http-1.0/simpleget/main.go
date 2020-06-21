package main

import (
  "io/ioutil"
  "log"
  "net/http"
)

func main() {
  resp, err := http.Get("http://localhost:8080") // net/httpのGetメソッドでアクセス
  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  log.Println(string(body))
  log.Println("Status: ", resp.Status)
  log.Println("Status Code: ", resp.StatusCode)
  for k, v := range resp.Header {
    log.Println(k, " ", v[0])
  }
}
