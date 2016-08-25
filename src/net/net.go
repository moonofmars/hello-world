// net
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

//处理以/view/起头的HTTP请求
func viewHandler(w http.ResponseWriter, r *http.Request) {
	//截取HTTP路径参数
	title := r.URL.Path[len("/view/"):]
	//加载页面
	p := loadPage(title)
	// 打印内容到 HTTP Response
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
	//fmt.Fprintf(w, "<h1>%s is</h1><div>\nbody %s\n</div>", p.Title, p.Body)
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[:])
	//io.WriteString(w,"Hello, world!\n")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p := loadPage(title)
	//fmt.Fprintf(w, "title is:", p.Title)
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	body := r.FormValue("body2")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler) //注册一个Handler，Handler是逻辑处理接口，它包含一个ServeHTTP方法，由业务实现该方法，并注册给http server
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil) //http.FileServer(http.Dir("."))
}

/*


func main() {
	fmt.Println("Hello World!")
	rd, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		fmt.Printf("%s \n", err.Error())
		return
	}
	b, err := ioutil.ReadAll(rd.Body)
	rd.Body.Close()
	if err == nil {
		fmt.Printf("%s\n\n", string(b))
		fmt.Printf("%s", rd.Header)
	}

*/
