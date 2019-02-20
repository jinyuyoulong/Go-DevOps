package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch() {
	cmd := exec.Command("sh", "./deploay.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}
func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1> hello , this is My deploy page</h1>")
	reLaunch()
}
func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
}
