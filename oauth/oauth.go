// package oauther provides the ability to obtain OAuth tokens and to use
// these inside a Go application.
//
// OAuther is based heavily on this example code:
// https://code.google.com/p/goauth2/source/browse/oauth/example/oauthreq.go
package oauth

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"fmt"
	"github.com/oxtoacart/webbrowser"
	"net/http"
	"net/url"
)

var (
	codeChannel  = make(chan string)
	errorChannel = make(chan error)
)

// ObtainToken uses an interactive browser session to obtain an oauth.Token
// in json form.
func ObtainToken(
	clientId string,
	clientSecret string,
	port string,
	scope string,
	authURL string,
	tokenURL string) (jsonToken []byte, err error) {

	go runServer(port)

	// Set up a configuration.
	config := &oauth.Config{
		AccessType:   "offline",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  fmt.Sprintf("http://localhost:%s/", port),
		Scope:        scope,
		AuthURL:      authURL,
		TokenURL:     tokenURL,
	}

	// Set up a Transport using the config.
	transport := &oauth.Transport{Config: config}

	// Get an authorization code from the data provider.
	// ("Please ask the user if I can access this resource.")
	authCodeUrl, _ := url.QueryUnescape(config.AuthCodeURL(""))
	webbrowser.Open(authCodeUrl)

	select {
	case authCode := <-codeChannel:
		// Exchange the authorization code for an access token.
		// ("Here's the code you gave the user, now give me a token!")
		var token *oauth.Token
		if token, err = transport.Exchange(authCode); err != nil {
			return
		}
		jsonToken, err = json.Marshal(token)
		return
	case err = <-errorChannel:
		return
	}
}

func TransportWithToken(jsonToken []byte) (transport *oauth.Transport, err error) {
	token := &oauth.Token{}
	if err = json.Unmarshal(jsonToken, token); err != nil {
		return
	}
	transport = &oauth.Transport{}
	transport.Token = token
	return
}

func runServer(port string) {
	handler := http.NewServeMux()
	handler.HandleFunc("/", handleCallback)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler,
	}
	server.ListenAndServe()
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	errStrings := r.URL.Query()["error"]
	if len(errStrings) == 1 {
		err := fmt.Errorf("Unable to obtain authorization: %s", errStrings[0])
		errorChannel <- err
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("%s", err)))
	} else {
		codes := r.URL.Query()["code"]
		if codes != nil {
			authCode := codes[0]
			codeChannel <- authCode
			w.WriteHeader(200)
			w.Write([]byte("Authorization received, Thank You!"))
		}
	}
}
