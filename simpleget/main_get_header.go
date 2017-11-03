package main

import (
  "net/http"
  "log"
)

func main() {
  resp, err := http.Get("http://localhost:18888")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  if err != nil {
    panic(err)
  }
  log.Println("Status:", resp.Status)
  log.Println("Headers:", resp.Header)
}
