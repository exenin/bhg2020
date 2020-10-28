package main
import (
    "net/http"
    "fmt"
    "io/ioutil"
)
func main(){

    resp, _:= http.Get("https://sharpspring.com")
    fmt.Println(resp.Status)
    body, _:= ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    resp.Body.Close()
}
