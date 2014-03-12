package main

import (
	"flag"
	"fmt"
	"github.com/oxtoacart/oauther/oauth"
	"os"
)

var (
	clientId     = flag.String("id", "", "Client ID")
	clientSecret = flag.String("secret", "", "Client Secret")
	scope        = flag.String("scope", "", "OAuth scope")
	port         = flag.String("port", "9000", "Port for callback web server")
	authURL      = flag.String("auth_url", "https://accounts.google.com/o/oauth2/auth", "Authentication URL")
	tokenURL     = flag.String("token_url", "https://accounts.google.com/o/oauth2/token", "Token URL")
)

const usageMsg = `
oauther obtains a request token from an OAuth2 provider (Google by default)
and prints out a JSON structure that can later be used to initialize an
oauth.OAuther.

To obtain a request token you must specify, -id, -secret and -scope.

To obtain Client ID and Secret, see the "OAuth 2 Credentials" section under
the "API Access" tab on this page: https://code.google.com/apis/console/
`

func main() {
	flag.Parse()
	if *clientId == "" || *clientSecret == "" || *scope == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, usageMsg)
		os.Exit(1)
	}
	oauther := &oauth.OAuther{
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
		TokenURL:     *tokenURL,
		Scope:        *scope,
		Port:         *port,
		AuthURL:      *authURL,
	}
	if err := oauther.ObtainToken(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	} else {
		if jsonOauther, err := oauther.ToJSON(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(3)
		} else {
			fmt.Println(string(jsonOauther))
		}
	}
}
