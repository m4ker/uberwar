package uber

import (
	"encoding/json"
	"github.com/hulucat/utils"
	"io/ioutil"
	"net/http"
)

type UserProfile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
	PromoCode string `json:"promo_code"`
	Uuid      string `json:"uuid"`
}
type Location struct {
	Lat     float64 `json:"latitude"`
	Lng     float64 `json:"longitude"`
	Bearing float64 `json:"bearing"`
}

type CurrentInfo struct {
	RequestId string   `json:"request_id"`
	Status    string   `json:"status"`
	Location  Location `json:"location"`
}

func do(accessToken string, apiUrl string, method string) (bytes []byte, err error) {
	client := &http.Client{}
	r, err := http.NewRequest(method, apiUrl, nil)
	if err != nil {
		utils.Errorf("Error new http request: %s", err.Error())
		return
	}
	r.Header.Add("Authorization", " Bearer "+accessToken)
	response, err := client.Do(r)
	defer response.Body.Close()
	bytes, err = ioutil.ReadAll(response.Body)
	return
}

func GetUserProfile(accessToken string) (profile *UserProfile, err error) {
	bytes, err := do(accessToken, "https://sandbox-api.uber.com.cn/v1/me", "GET")
	profile = &UserProfile{}
	utils.Debugf("User profile returns: %s", string(bytes))
	json.Unmarshal(bytes, profile)
	return
}

func GetCurrent(accessToken string) (current *CurrentInfo, err error) {
	bytes, err := do(accessToken, "https://sandbox-api.uber.com.cn/v1/requests/current", "GET")
	current = &CurrentInfo{}
	utils.Debugf("Current returns: %s", string(bytes))
	json.Unmarshal(bytes, current)
	return
}
