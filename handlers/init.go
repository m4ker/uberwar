package handlers

import (
	"math/rand"
	"tripwar/models"
	"tripwar/uber"
)

var sids = make([]string, 0)
var users = make(map[string]*models.User)           //sid => user
var sidCurrent = make(map[string]*uber.CurrentInfo) //sid => currentInfo
var towers = make([]*models.Tower, 0)
var msgs = make(map[string][]string) //sid => []msg string
func getUser(userId string) (user *models.User, sid string) {
	for sid, user = range users {
		if user.Id == userId {
			break
		}
	}
	return
}

func init() {
	users["test1"] = &models.User{
		Id:              "test1",
		Name:            "Ji weicai",
		UberOpenId:      "test1",
		Faction:         int64(rand.Intn(2)),
		TotalScore:      10000,
		Score:           0,
		Mana:            6,
		Hp:              8,
		Lng:             116.308372,
		Lat:             39.979661,
		LocationBearing: 33,
	}
	users["test2"] = &models.User{
		Id:              "test2",
		Name:            "Gan lanjie",
		UberOpenId:      "test2",
		Faction:         int64(rand.Intn(2)),
		TotalScore:      10000,
		Score:           0,
		Mana:            6,
		Hp:              8,
		Lng:             116.309144,
		Lat:             39.979052,
		LocationBearing: 33,
	}
	users["test3"] = &models.User{
		Id:              "test3",
		Name:            "Niu pengch",
		UberOpenId:      "test3",
		Faction:         int64(rand.Intn(2)),
		TotalScore:      10000,
		Score:           0,
		Mana:            6,
		Hp:              8,
		Lng:             116.311076,
		Lat:             39.978674,
		LocationBearing: 33,
	}
}
