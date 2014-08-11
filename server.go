package listener

import (
	"encoding/json"
	"log"
	"net/http"
)

var msgHandlers = MsgHandlers()

type HubMessage struct {
	Repository struct {
		Status    string
		RepoUrl   string `json:"repo_url"`
		Owner     string
		IsPrivate bool `json:"is_private"`
		Name      string
		StarCount int    `json:"star_count"`
		RepoName  string `json:"repo_name"`
	}

	Push_data struct {
		PushedAt int `json:"pushed_at"`
		Images   []string
		Pusher   string
	}
}

func Serve(addr string) {
	http.HandleFunc("/", reqHandler)
	http.ListenAndServe(addr, Log(http.DefaultServeMux))
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var imgConfig HubMessage

	err := decoder.Decode(&imgConfig)
	if err != nil {
		http.Error(w, "Could not decode json", 500)
		log.Print(err)
		return
	}
	go notify(imgConfig)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func notify(img HubMessage) {
	msgHandlers.Call(img)
}
