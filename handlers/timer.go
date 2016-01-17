package handlers

import (
	"fmt"
	"github.com/hulucat/utils"
	"time"
)

func check() {
	//检查爆炸
	for sid, user := range users {
		for _, tower := range towers {
			if tower.Faction == user.Faction {
				continue
			}
			if utils.GetEarthDistance(user.Lat, user.Lng, tower.Lat, tower.Lng) < 10 { //Boom!
				cacheKey := fmt.Sprintf("tripwar_boom_%s_%s", user.Id, tower.Id)
				if utils.GetCache(cacheKey) != "" {
					continue
				}
				utils.Debugf("Boom! user: %s, tower: %s", user.Name, tower.Id)
				user.Hp -= 1
				users[sid] = user
				msgs[sid] = append(msgs[sid], "你碰到地雷，被炸掉1滴血")
				owner, osid := getUser(tower.UserId)
				owner.Score += 10
				owner.TotalScore += 10
				users[osid] = owner
				utils.SetCache(cacheKey, "1", 300)
			}

		}
	}
	//检查相遇
	for sid, user := range users {
		for _, user2 := range users {
			if user.Id == user2.Id {
				continue
			}
			if utils.GetEarthDistance(user.Lat, user.Lng, user2.Lat, user2.Lng) < 10 { //Papapa!
				cacheKey := fmt.Sprintf("tripwar_meets_%s_%s", user.Id, user2.Id)
				if utils.GetCache(cacheKey) != "" {
					continue
				}
				utils.Debugf("Papapa! user: %s, meets: %s", user.Name, user2.Name)
				user.Hp += 1
				user.Mana += 1
				users[sid] = user
				msgs[sid] = append(msgs[sid], "你遇到"+user2.Name+",获得1滴血和1魔法")
				utils.SetCache(cacheKey, "1", 300)
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
