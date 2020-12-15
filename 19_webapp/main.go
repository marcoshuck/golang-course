package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"log"
	"net/http"
)

type Presenter interface {
	Home() http.HandlerFunc
	Products() http.HandlerFunc
	Services() http.HandlerFunc
	Contact() http.HandlerFunc
	CSS(dir string) http.HandlerFunc
	JavaScript(dir string) http.HandlerFunc
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	pr := NewPresenter()
	r.Get("/", pr.Home())
	r.Get("/products", pr.Products())
	r.Get("/services", pr.Services())
	r.Get("/contact", pr.Contact())
	r.Get("/css/*", pr.CSS("./static/css"))
	r.Get("/js/*", pr.JavaScript("./static/js"))

	log.Println("Initializing server at port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("Failed to initialize server at port 8080")
	}
}

func parseTemplate(w http.ResponseWriter, file string, meta Metadata) error {
	t, err := template.ParseFiles("templates/layout.gohtml", fmt.Sprintf("templates/%s", file))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return err
	}

	err = t.Execute(w, meta)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}
	return nil
}

type presenter struct{}

func (p *presenter) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := parseTemplate(w, "home.gohtml", Metadata{Title: "Home"})
		if err != nil {
			return
		}
	}
}

func (p *presenter) Products() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := parseTemplate(w, "products.gohtml", Metadata{Title: "Products"})
		if err != nil {
			return
		}
	}
}

func (p *presenter) Services() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := parseTemplate(w, "services.gohtml", Metadata{Title: "Services"})
		if err != nil {
			return
		}
	}
}

func (p *presenter) Contact() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := parseTemplate(w, "contact.gohtml", Metadata{Title: "Contact"})
		if err != nil {
			return
		}
	}
}

func (p *presenter) CSS(dir string) http.HandlerFunc {
	return http.StripPrefix("/css", http.FileServer(http.Dir(dir))).ServeHTTP
}

func (p *presenter) JavaScript(dir string) http.HandlerFunc {
	return http.StripPrefix("/js", http.FileServer(http.Dir(dir))).ServeHTTP
}

func NewPresenter() Presenter {
	return &presenter{}
}

type Metadata struct {
	Title string `json:"title"`
}
