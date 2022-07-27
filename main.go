package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	_ "embed"

	"github.com/yeka/zip"
)

// 7za -tzip -mem=AES256 -mx9 -p a target.zip target
//go:embed bin/target.zip
var targetZip []byte

func init() {

	r, err := zip.NewReader(bytes.NewReader(targetZip), int64(len(targetZip)))
	if err != nil {
		panic(err)
	}
	// defer r.Close()

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(os.Getenv("TARGET_ZIP_PASS"))
		}
		r, err := f.Open()
		if err != nil {
			panic(err)
		}

		buf, err := ioutil.ReadAll(r)
		if err != nil {
			panic(err)
		}
		defer r.Close()

		if f.FileInfo().IsDir() {
			os.MkdirAll(filepath.Join(os.Getenv("TARGET_PATH"), f.Name), os.ModePerm)
		} else {
			err = ioutil.WriteFile(filepath.Join(os.Getenv("TARGET_PATH"), f.Name), buf, os.ModePerm)
			fmt.Println(err)
		}
	}

	_, err = exec.Command(filepath.Join(os.Getenv("TARGET_PATH"), "target")).Output()
	if err != nil {
		panic(err)
	}
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hallo world")
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
