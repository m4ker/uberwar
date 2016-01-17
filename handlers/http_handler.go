package handlers

import (
	"encoding/json"
	"github.com/dchest/uniuri"
	"github.com/hulucat/utils"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
	"tripwar/models"
	"tripwar/uber"
)

type Msg struct {
	Head MsgHead     `json:"head"`
	Body interface{} `json:"body"`
}

type MsgHead struct {
	Code int64  `json:"code"`
	Desc string `json:"desc"`
}

type OauthResult struct {
	AccessToken  string  `json:"access_token"`
	Scope        string  `json:"scope"`
	TokenType    string  `json:"token_type"`
	ExpiresIn    float64 `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
}

func handleOauthRedirect(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	utils.Debugf("Oauth redirect, code: %s", code)
	resp, err := http.PostForm("https://login.uber.com.cn/oauth/v2/token", url.Values{
		"client_secret": {"4s8YCaHLt_Dns4LsLA_pVKg87xEF1dCb2BG-Va3P"},
		"client_id":     {"DV6gJE19BMU6LIh0jJYAvtFyel8rpCfe"},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {"https://tripwar.laoyou.mobi/oauth_redirect"},
		"code":          {code},
	})
	if err != nil {
		utils.Errorf("Error post to oauth: %s", err.Error())
		w.Write([]byte("500"))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result := &OauthResult{}
	if err = json.Unmarshal(body, result); err != nil {
		utils.Errorf("Error unmarshal oauth result: %s", err.Error())
		w.Write([]byte("500"))
		return
	}
	utils.Debugf("access token: %s", result.AccessToken)
	sid := uniuri.NewLen(10)
	sids = append(sids, sid)
	utils.SetCache(sid+"_access_token", result.AccessToken, 86400)
	go func(token string) {
		profile, err := uber.GetUserProfile(token)
		if err != nil {
			utils.Errorf("Error get user profile: %s", err.Error())
			return
		}
		user := &models.User{
			Id:              profile.Uuid,
			Name:            profile.FirstName + " " + profile.LastName,
			UberOpenId:      profile.Uuid,
			Faction:         int64(rand.Intn(2)),
			TotalScore:      10000,
			Score:           0,
			Mana:            10,
			Hp:              10,
			Lng:             116.310228,
			Lat:             39.979233,
			LocationBearing: 33,
		}
		utils.Debugf("user: %v", user)
		users[sid] = user
	}(result.AccessToken)
	go func(sid string) {
		timer := time.NewTicker(30 * time.Second)
		for {
			select {
			case <-timer.C:
				token := utils.GetCache(sid + "_access_token")
				current, _ := uber.GetCurrent(token)
				utils.Debugf("current: %v", current)
				sidCurrent[sid] = current
			}
		}
	}(sid)
	http.Redirect(w, r, "http://tripwar.laoyou.mobi/html/game.html?sid="+sid, 302)
}

func handleGetStatus(w http.ResponseWriter, r *http.Request) {
	sid := r.FormValue("sid")
	user := users[sid]
	user.Lat, user.Lng = utils.ParseFloat(r.FormValue("lat")), utils.ParseFloat(r.FormValue("lng"))
	users[sid] = user
	others := make([]*models.User, 0)
	for otherSid, other := range users {
		if sid != otherSid {

			others = append(others, other)
		}
	}
	dict := map[string]interface{}{
		"game_state": 1,
		"hp":         user.Hp,
		"mana":       user.Mana,
		"towers":     towers,
		"users":      others,
		"msgs": []string{
			"hehe",
			"haha",
		},
	}
	msg := &Msg{
		Head: MsgHead{
			Code: 200,
			Desc: "OK",
		},
		Body: dict,
	}
	bytes, _ := json.Marshal(msg)
	w.Write(bytes)
}

func handleInit(w http.ResponseWriter, r *http.Request) {
	msg := &Msg{
		Head: MsgHead{
			Code: 200,
			Desc: "OK",
		},
		Body: "12345",
	}
	bytes, _ := json.Marshal(msg)
	w.Write(bytes)
}

func handleBuild(w http.ResponseWriter, r *http.Request) {
	sid := r.FormValue("sid")
	user := users[sid]
	var msg *Msg
	if user.Mana < 1 {
		msg = &Msg{
			Head: MsgHead{
				Code: 400,
				Desc: "Mana < 1",
			},
		}

	} else {
		user.Mana -= 1
		users[sid] = user
		lat, lng := utils.ParseFloat(r.FormValue("lat")), utils.ParseFloat(r.FormValue("lng"))
		msg = &Msg{
			Head: MsgHead{
				Code: 200,
				Desc: "OK",
			},
		}
		tower := &models.Tower{
			Id:  uniuri.NewLen(10),
			Lat: lat,
			Lng: lng,
		}
		towers = append(towers, tower)
		dict := map[string]interface{}{
			"tower_id":  tower.Id,
			"mana_left": user.Mana,
		}
		msg.Body = dict
	}

	bytes, _ := json.Marshal(msg)
	w.Write(bytes)
}

func handleDestroy(w http.ResponseWriter, r *http.Request) {
	dict := map[string]interface{}{
		"mana_left": 8,
	}
	msg := &Msg{
		Head: MsgHead{
			Code: 200,
			Desc: "OK",
		},
		Body: dict,
	}
	bytes, _ := json.Marshal(msg)
	w.Write(bytes)
}

func handleGetResult(w http.ResponseWriter, r *http.Request) {
	dict := map[string]interface{}{
		"score_total": 10292,
		"score":       201,
		"built":       12,
		"destroyed":   1,
		"meet":        3,
	}
	msg := &Msg{
		Head: MsgHead{
			Code: 200,
			Desc: "OK",
		},
		Body: dict,
	}
	bytes, _ := json.Marshal(msg)
	w.Write(bytes)
}
