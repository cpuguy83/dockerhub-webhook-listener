package listener

import (
	"encoding/json"
	"log"
	"net/http"
	"bytes"
)

var msgHandlers Registry

type HubMessage struct {
	Callback_url string
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

type CallbackMessage struct {
	State            string `json:"state"`
	Description      string `json:"description"`
}

type Config struct {
	ListenAddr string
	Mailgun    mailGunConfig
	Tls        struct {
		Key  string
		Cert string
	}
	Apikeys struct {
		Key []string
	}
}

var ServerConfig *Config
var client = &http.Client{}

func Serve(config *Config) error {
	ServerConfig = config
	if len(ServerConfig.Apikeys.Key) == 0 {
		log.Print("Warning: The server is about to start without any authentication.  Anyone can trigger handlers to fire off")
		log.Print("To enable authentication, you must add an `apikeys` section with at least 1 `key`")
	}
	msgHandlers = MsgHandlers()
	http.HandleFunc("/", reqHandler)
	if config.Tls.Key != "" && config.Tls.Cert != "" {
		log.Print("Starting with SSL")
		return http.ListenAndServeTLS(config.ListenAddr, config.Tls.Cert, config.Tls.Key, Log(http.DefaultServeMux))
	}
	log.Print("Warning: Server is starting without SSL, you should not pass any credentials using this configuration")
	log.Print("To use SSL, you must provide a config file with a [tls] section, and provide locations to a `key` file and a `cert` file")
	return http.ListenAndServe(config.ListenAddr, Log(http.DefaultServeMux))
}

// Send callback request
func sendCallback(callbackUrl string, msg *CallbackMessage) {
	log.Printf("Send callback to %s", callbackUrl)

	jsonStr, err := json.Marshal(msg)
	if err != nil {
		log.Print("Failed to marshal callback message")
		log.Print(err)
		return
	}

	req, err := http.NewRequest("POST", callbackUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Print("Failed to make callback request")
		log.Print(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)

	if err != nil {
		log.Print("Failed to request callback")
		log.Print(err)
		return
	}

	log.Print("Succeeded to request callback")
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

	if authenticateRequest(r) {
		go handleMsg(imgConfig)
		return
	}

	http.Error(w, "Not Authorized", 401)
	sendCallback(imgConfig.Callback_url, &CallbackMessage{
		State: "failure",
		Description: "Not authorized",
	})
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.RemoteAddr, r.Method)
		handler.ServeHTTP(w, r)
	})
}

func handleMsg(img HubMessage) {
	msgHandlers.Call(img)
	sendCallback(img.Callback_url, &CallbackMessage{
		State: "success",
		Description: "Hook successfully received",
	})
}

func authenticateRequest(r *http.Request) bool {
	key := r.URL.Query().Get("apikey")
	for _, cfgKey := range ServerConfig.Apikeys.Key {
		if key == cfgKey {
			return true
		}
		continue
	}
	return (len(ServerConfig.Apikeys.Key) == 0) || false
}
