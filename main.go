package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-git/go-git/v5"
	//"github.com/go-git/go-git"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	os.RemoveAll(path + "/gunstein_vatnar_no")
	_, err = git.PlainClone(path+"/gunstein_vatnar_no", false, &git.CloneOptions{
		URL:      "https://github.com/gunstein/gunstein_vatnar_no.git",
		Progress: os.Stdout,
		Depth:    1,
	})

	//CheckIfError(err)

	fs := http.FileServer(http.Dir("./gunstein_vatnar_no"))
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
