//+build ignore

package main 

import(
"log"
"fmt"
"io/ioutil"
// "."

"net/http"
"net/http/cookiejar"
"net/url"
"strings"
"time"
"encoding/json"
)

const (
	loginURL     string = "https://sso.pokemon.com/sso/login?service=https%3A%2F%2Fsso.pokemon.com%2Fsso%2Foauth2.0%2FcallbackAuthorize";
	loginOAuth   string = "https://sso.pokemon.com/sso/oauth2.0/accessToken";
);

type LogOnDetails struct{
	Username,
	Password,
	AuthType string;
};

func main() {
	details := LogOnDetails{
		"hbonly",
		"hay111",
		"ptc",
	};
	authWithPTC(&details);
}

func authWithGoogle(details *LogOnDetails) (string, error) {
	_, masterToken, err := gpsoauth.Login(details.Username, details.Password, androidID)
	if err != nil {
		return "", fmt.Errorf("[!] Failed to Login with Google\nUsername: %s\nPassword: %s\nAndroidID: %s", details.Username, details.Password, androidID)
	}
	body, err := gpsoauth.OAuth(details.Username, masterToken, androidID, oAuthService, app, clientSIG)
	if err != nil {
		return "", fmt.Errorf("[!] Failed to Login with Google\nUsername: %s\nPassword: %s\nAndroidID: %s", details.Username, details.Password, androidID)
	}

	if _, ok := body["Auth"]; !ok {
		return "", fmt.Errorf("[!] Missing AUTH. Could be an incorrect Email or Password, or 2 step authentication failure. (This package does not support 2 step auth)")
	}

	return body["Auth"], nil
}

func authWithPTC(details *LogOnDetails) (string, error) {
	// Initiate HTTP Client / Cookie JAR

	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", fmt.Errorf("Failed to create new cookiejar for client")
	}
	newClient := &http.Client{Jar: jar, Timeout: 15 * time.Second}

	// First Request

	req, err := http.NewRequest("GET", loginURL, nil)
	if err != nil {
		return "", fmt.Errorf("认证 Pokemon Trainers Club\n Details: \n\n Username: %s\n Password: %s\n AuthType: %s\n", details.Username, details.Password, details.AuthType)
	}
	req.Header.Set("User-Agent", "niantic")
	firstResp, err := newClient.Do(req); //发送请求
	if err != nil {
		return "", fmt.Errorf("Failed to send intial handshake: Possible wrong Username or Password", err)
	}
	respJSON := make(map[string]string)
	err = json.NewDecoder(firstResp.Body).Decode(&respJSON)
	if err != nil {
		return "", fmt.Errorf("Failed to decode JSON Body: %v", err)
	}

	defer firstResp.Body.Close();
	log.Println(fmt.Sprintf("收到握手应答: %v", respJSON));

	// Second Request

	form := url.Values{}
	form.Add("lt", respJSON["lt"])
	form.Add("execution", respJSON["execution"])
	form.Add("_eventId", "submit")
	form.Add("username", details.Username)
	form.Add("password", details.Password)
	req, err = http.NewRequest("POST", loginURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("Failed to send second request authing with PTC: %v", err)
	}
	req.Header.Set("User-Agent", "niantic")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	secResp, err := newClient.Do(req); //发送请求
	if err != nil {
		return "", fmt.Errorf("Failed to send second request authing with PTC: %v", err)
	}

	ticket := secResp.Request.URL.String()

	if strings.Contains(ticket, "ticket") {
		ticket = strings.Split(ticket, "ticket=")[1]
	} else {
		return "", fmt.Errorf("Failed could not get the Ticket from the second request\n.. Possible wrong Username or Password")
	}
	defer secResp.Body.Close();
	log.Println(fmt.Sprintf("收到票根应答: %v", ticket));

	// Third Request

	form = url.Values{}
	form.Add("client_id", "mobile-app_pokemon-go")
	form.Add("redirect_uri", "https://www.nianticlabs.com/pokemongo/error")
	form.Add("client_secret", "w8ScCUXJQc6kXKw8FiOhd8Fixzht18Dq3PEVkUCP5ZPxtgyWsbTvWHFLm2wNY0JR")
	form.Add("grant_type", "refresh_token")
	form.Add("code", ticket)
	req, err = http.NewRequest("POST", loginOAuth, strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("Failed to send the third request authing with PTC: %v", err)
	}
	req.Header.Add("User-Agent", "niantic")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	thirdResp, err := newClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send the third request authing with PTC: %v", err)
	}
	defer thirdResp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(thirdResp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to decode the body of the third request")
	}

	body := string(bodyBytes);
	log.Println(fmt.Sprintf("收到认证应答: %v", body));

	if strings.Contains(body, "token=") {
		token := strings.Split(body, "token=")[1]
		token = strings.Split(token, "&")[0]
		return token, nil
	} else {
		return "", fmt.Errorf("Failed to get the token on the third request \nBody:\n\n%v", body)
	}
}