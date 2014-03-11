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
	codeChannel = make(chan string)
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
	authCode := <-codeChannel

	// Exchange the authorization code for an access token.
	// ("Here's the code you gave the user, now give me a token!")
	var token *oauth.Token
	if token, err = transport.Exchange(authCode); err != nil {
		return
	}
	jsonToken, err = json.Marshal(token)
	return
}

func runServer(port string) {
	http.HandleFunc("/", handleCallback)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	codes := r.URL.Query()["code"]
	if codes != nil {
		authCode := codes[0]
		codeChannel <- authCode
		w.WriteHeader(200)
		w.Write([]byte("Authorization Code Received, Thank You"))
	}
}
