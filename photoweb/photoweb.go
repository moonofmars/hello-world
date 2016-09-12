// photoweb
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
)

var i int

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, `<html><head><title></title></head><body>
			<form method="POST" action="/upload"
				enctype="multipart/form-data">
			      Choose an image to upload: <input name="image" type="file" />
			      <input type="submit" value="Upload" />
			   </form>
			</body></html> `)
		return
	}

	if r.Method == "POST" {
		fmt.Println("r.Method == \"POST\"")
		f, _, err := r.FormFile("image")
		if err != nil {
			fmt.Println("err1")
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}

		filename := "jj.jpg" //h.Filename
		defer f.Close()
		//t, err := os.Create("/uploads")
		t, err := os.Create(filename)
		if err != nil {
			fmt.Println("err2")
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			fmt.Println("err3")
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		fmt.Println("Redirect: ", "/view?id="+filename)
		http.Redirect(w, r, "/views", http.StatusFound)
		//http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
		//fmt.Println("filename: ", filename, "\t i:", i)
		//i++
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := imageId
	/*
		fmt.Println("imagePath: ", imagePath)
		exists, _ := isExists(imagePath)
		fmt.Println("exists: ", exists)
		if !exists {
			http.NotFound(w, r)
			return
		}
	*/
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)

}
func isExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//	if e, ok := err.(*os.PathError); ok && e.Error == os.ENOENT {
	//	return false, nil
	//}
	return false, err
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listen and serve: ", err.Error())
	}
}
