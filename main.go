package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func uploadFile(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "uploading file")
	// 1. Parse input, type multipart/form-data
	r.ParseMultipartForm(10 << 20)

	//2. retrieve file from posted form data
	file, handler, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println("error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("uploaded file %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME header %+v\n", handler.Header)

	// 3. write temporary file on our server
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	//write file
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	tempFile.Write(fileBytes)


	// 4. return if successfull
	fmt.Fprintf(w, "Success")
}
func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}


func main() {
	fmt.Println("working")
	setupRoutes()
}
