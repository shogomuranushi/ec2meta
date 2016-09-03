package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
)

func main() {
  metaurl := "http://169.254.169.254/latest/meta-data/"
  url := metaurl + os.Args[1]

  resp, _ := http.Get(url)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray))
}
