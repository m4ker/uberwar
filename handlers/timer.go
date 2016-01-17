package handlers

import (
	"github.com/hulucat/utils"
	"time"
)

func check() {
	//检查爆炸
	for sid, user := range users {
		for _, tower := range towers {
			if utils.GetEarthDistance(user.Lat, user.Lng, tower.Lat, tower.Lng) < 10 { //Boom!
				user.Hp -= 1
				users[sid] = user
				msgs[sid] = append(msgs[sid], "你碰到地雷，被炸掉1滴血")
				owner, osid := getUser(tower.UserId)
				owner.Score += 10
				owner.TotalScore += 10
				users[osid] = owner

			}

		}
	}
	//检查相遇
	for sid, user := range users {
		for _, user2 := range users {
			if utils.GetEarthDistance(user.Lat, user.Lng, user2.Lat, user2.Lng) < 10 { //Papapa!
				user.Hp += 1
				user.Mana += 1
				users[sid] = user
				msgs[sid] = append(msgs[sid], "你遇到"+user2.Name+",获得1滴血和1魔法")
			}
		}
	}
}

func Start() {
	timer := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-timer.C:
			check()
		}
	}
}
