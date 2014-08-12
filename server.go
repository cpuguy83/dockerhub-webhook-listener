package listener

import (
	"encoding/json"
	"log"
	"net/http"
)

var msgHandlers Registry

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

type Config struct {
	ListenAddr string
	Mailgun    mailGunConfig
	Tls        struct {
		Key  string
		Cert string
	}
}

func Serve(config *Config) error {
	msgHandlers = MsgHandlers(config)
	http.HandleFunc("/", reqHandler)
	if config.Tls.Key != "" && config.Tls.Cert != "" {
		log.Print("Starting with SSL")
		return http.ListenAndServeTLS(config.ListenAddr, config.Tls.Cert, config.Tls.Key, Log(http.DefaultServeMux))
	}
	log.Print("Warning: Server is starting without SSL, you should not pass any credentials using this configuration")
	log.Print("To use SSL, you must provide a config file with a [tls] section, and provide locations to a `key` file and a `cert` file")
	return http.ListenAndServe(config.ListenAddr, Log(http.DefaultServeMux))
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
	go handleMsg(imgConfig)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func handleMsg(img HubMessage) {
	msgHandlers.Call(img)
}
