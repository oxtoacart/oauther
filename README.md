oauther
=======

oauther provides a utility for obtaining Google OAuth2 tokens and a library for
using these from within a Go application.

Compiled releases of oauther are available in the [releases](releases) folder.

### Command-line Usage (Obtain a Token)

Macintosh% oauther
Usage of oauther:
  -auth_url="https://accounts.google.com/o/oauth2/auth": Authentication URL
  -id="": Client ID
  -port="9000": Port for callback web server
  -scope="": OAuth scope
  -secret="": Client Secret
  -token_url="https://accounts.google.com/o/oauth2/token": Token URL

oauther obtains a request token from an OAuth2 provider (Google by default)
and prints it to stdout.

To obtain a request token you must specify, -id, -secret and -scope.

To obtain Client ID and Secret, see the "OAuth 2 Credentials" section under
the "API Access" tab on this page: https://code.google.com/apis/console/

### Library Usage (Using a Token)

```go
import (
    "github.com/oxtoacart/oauther/oauth"
    "net/http"
)

func myFunc() {
    var jsonToken []byte
    // read jsonToken from somewhere
    transport, err := oauth.transportWithToken(jsonToken)
    // transport is a *code.google.com/p/goauth2/oauth.Transport
    var client *http.Client
    client := transport.Client()
    // now use your client to make http requests
}
```


