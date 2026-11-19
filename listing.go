package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"google.golang.org/api/appengine/v1"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//go:embed assets
var assets embed.FS

func main() {
	http.Handle("/assets/", http.FileServerFS(assets))
	http.HandleFunc("/api/services", getServiceListingAPI)
	http.HandleFunc("/", getServiceListingHTML)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func getServiceListing(w http.ResponseWriter, r *http.Request) (appEngineServicesStruct, error) {
	appName := os.Getenv("GOOGLE_CLOUD_PROJECT")
	log.Println("appName:" + appName)
	// add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx := context.Background()
	var appEngineServices appEngineServicesStruct
	appengineService, err := appengine.NewService(ctx)
	if err != nil {
		log.Print("err creating appengine service ", err)
		http.Error(w, "service unavailable", http.StatusInternalServerError)
		return appEngineServices, err
	}
	response, err := appengine.NewAppsServicesService(appengineService).List(appName).Do()
	if err != nil {
		log.Print("err calling NewAppsServicesService.List", err)
		return appEngineServices, err
	}
	log.Println("Loading JSON")
	appEngineServices = transform(response.Services)
	return appEngineServices, nil
}

func getServiceListingAPI(w http.ResponseWriter, r *http.Request) {
	appEngineServices, err := getServiceListing(w, r)
	if err != nil {
		return
	}
	data, err := json.Marshal(appEngineServices)
	if err != nil {
		log.Print("err mashal", err)
	}
	fmt.Fprintf(w, "%s", data)
}

func getServiceListingHTML(w http.ResponseWriter, r *http.Request) {
	appEngineServices, err := getServiceListing(w, r)
	if err != nil {
		return
	}
	transformAndDisplay(appEngineServices, w)
}

type appEngineServicesStruct struct {
	Services []appEngineServiceStruct `json:"services"`
}

type appEngineServiceStruct struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Title      string `json:"title"`
	Track      string `json:"track,omitempty"` // topic key: cloud, web, data, agile; empty if unknown
	TrackLabel string `json:"trackLabel,omitempty"`
}

// tracks maps keywords found in a service ID to a topic (first match wins).
var tracks = []struct {
	key, label string
	keywords   []string
}{
	{"data", "Data & IA", []string{"data", "ai", "llm", "claude", "ml", "elastic", "spark", "kafka"}},
	{"agile", "Agilité", []string{"agile", "scrum", "kanban", "product"}},
	{"cloud", "Cloud & DevOps", []string{"docker", "podman", "kube", "k8s", "terraform", "ansible", "sre", "cloud", "gcp", "aws", "azure", "devops", "linux", "ci"}},
	{"web", "Web & Craft", []string{"angular", "react", "vue", "js", "javascript", "typescript", "java", "spring", "go", "git", "test", "clean", "front", "web"}},
}

func trackOf(id string) (string, string) {
	words := strings.Split(strings.ToLower(id), "-")
	for _, t := range tracks {
		for _, w := range words {
			for _, k := range t.keywords {
				if w == k || (len(k) > 3 && strings.HasPrefix(w, k)) {
					return t.key, t.label
				}
			}
		}
	}
	return "", ""
}

func formatName(name string) string {
	// Function replacing words (assuming lower case input)
	replace := func(word string) string {
		switch word {
		case "zen-" /*, "formation-"*/ :
			return ""
		}

		if word[len(word)-1] == '-' {
			word = word[:len(word)-1] + " "
		}
		return strings.Title(word)
	}

	r := regexp.MustCompile(`(\w+-|\w+)`)
	formattedName := r.ReplaceAllStringFunc(name, replace)

	log.Println(formattedName)
	return formattedName
}

func transform(services []*appengine.Service) appEngineServicesStruct {
	appEngineServices := appEngineServicesStruct{}
	for index := range services {
		appEngineService := appEngineServiceStruct{
			ID:    services[index].Id,
			Name:  services[index].Name,
			URL:   fmt.Sprintf("https://%s-dot-zen-formations.appspot.com/", services[index].Id),
			Title: formatName(services[index].Id),
		}
		appEngineService.Track, appEngineService.TrackLabel = trackOf(appEngineService.ID)
		appEngineServices.Services = append(appEngineServices.Services, appEngineService)
	}
	return appEngineServices
}

func transformAndDisplay(appEngineServices appEngineServicesStruct, w http.ResponseWriter) {
	tmpl := template.Must(template.New("listing.gohtml").Funcs(template.FuncMap{
		// icon inlines a file of assets/icons/ so that currentColor follows the theme.
		"icon": func(name string) (template.HTML, error) {
			b, err := assets.ReadFile("assets/icons/" + name)
			return template.HTML(b), err
		},
	}).ParseFiles("template/listing.gohtml"))
	tc := make(map[string]interface{})
	tc["Services"] = appEngineServices.Services
	if err := tmpl.Execute(w, tc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
