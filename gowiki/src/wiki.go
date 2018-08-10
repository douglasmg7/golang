package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "html/template"
  )

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
      return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        fmt.Println("Title:", title)
        p = &Page{Title: title}
    } else {
      fmt.Println("No title.")
      fmt.Println(err)
    }
    t, _ := template.ParseFiles("../template/edit.html")
    t.Execute(w, p)
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
//   title := r.URL.Path[len("/view/"):]
//   p, _ := loadPage(title)
//   fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

func main() {
  // http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))

  // body, err := ioutil.ReadFile("../template/edit.html")
  // if err == nil {
  //   // fmt.Fprintf("A%s", body);
  //   fmt.Println(body);
  // }
  // if err != nil {
  //   fmt.Println(err)
  // }
}