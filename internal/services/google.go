package services

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:9090/callback-gl",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	oauthStateStringGl = ""
)

func InitializeOauthGoogle() {
	oauthConfGl.ClientID = viper.GetString("google.clientID")
	oauthConfGl.ClientSecret = viper.GetString("google.clientSecret")
	oauthStateStringGl = viper.GetString("oauthStateString")
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	HandleLogin(w, r, oauthConfGl, oauthStateStringGl)
}

/*
CallBackFromGoogle Function
*/
func CallBackFromGoogle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Callback-gl..")

	state := r.FormValue("state")
	fmt.Println(state)
	if state != oauthStateStringGl {
		//logger.Log.Info("invalid oauth state, expected " + oauthStateStringGl + ", got " + state + "\n")
		fmt.Println("invalid oauth state, expected " + oauthStateStringGl + ", got " + state + "\n")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	fmt.Println(code)

	if code == "" {
		fmt.Printf("Code not found..")
		w.Write([]byte("Code Not Found to provide AccessToken..\n"))
		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("User has denied Permission.."))
		}
		// User has denied access..
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Printf("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			return
		}
		fmt.Printf("TOKEN>> AccessToken>> " + token.AccessToken)
		fmt.Printf("TOKEN>> Expiration Time>> " + token.Expiry.String())
		fmt.Printf("TOKEN>> RefreshToken>> " + token.RefreshToken)

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			fmt.Println("Get: " + err.Error() + "\n")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ReadAll: " + err.Error() + "\n")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		fmt.Println("parseResponseBody: " + string(response) + "\n")

		w.Write([]byte("Hello, I'm protected\n"))
		w.Write([]byte(string(response)))
		return
	}
}
