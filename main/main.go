package main

import (
	"cyoa"
	"html/template"
	"net/http"
	"os"
)

func main() {
	// open json file
	file, err := os.ReadFile("gopher.json")
	if err != nil {
		panic(err)
	}
	adventure := cyoa.ParseJSON(file)

	// open template file
	storyTemplate := openTemplate("story.html")

	// start server
	http.HandleFunc("/", AdventureHandler(adventure, storyTemplate))
	http.ListenAndServe(":8080", nil)

}

func openTemplate(filename string) *template.Template {
	t, err := template.ParseFiles("story.html")
	if err != nil {
		panic(err)
	}
	return t
}

func AdventureHandler(adventure cyoa.Adventure, storyTemplate *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get path from url
		path := r.URL.Path[1:]
		path = path[:len(path)-5]

		// get scene from path
		scene := adventure[path]

		// render scene
		err := storyTemplate.Execute(w, scene)
		if err != nil {
			panic(err)
		}
	}
}
