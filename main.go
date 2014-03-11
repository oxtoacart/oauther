package main

import (
	"./oauth"
	"flag"
	"fmt"
	"os"
)

var (
	clientId     = flag.String("id", "", "Client ID")
	clientSecret = flag.String("secret", "", "Client Secret")
	port         = flag.String("port", "8080", "Port for callback web server")
	scope        = flag.String("scope", "https://www.googleapis.com/auth/userinfo.profile", "OAuth scope")
	authURL      = flag.String("auth_url", "https://accounts.google.com/o/oauth2/auth", "Authentication URL")
	tokenURL     = flag.String("token_url", "https://accounts.google.com/o/oauth2/token", "Token URL")
)

const usageMsg = `
To obtain a request token you must specify both -id and -secret.

To obtain Client ID and Secret, see the "OAuth 2 Credentials" section under
the "API Access" tab on this page: https://code.google.com/apis/console/

Once you have completed the OAuth flow, the credentials should be stored inside
the file specified by -cache and you may run without the -id and -secret flags.
`

func main() {
	flag.Parse()
	if *clientId == "" || *clientSecret == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, usageMsg)
		os.Exit(1)
	}
	jsonToken, err := oauth.ObtainToken(
		*clientId,
		*clientSecret,
		*port,
		*scope,
		*authURL,
		*tokenURL)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	} else {
		fmt.Println(string(jsonToken))
	}
}
