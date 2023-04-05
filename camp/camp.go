package camp

import "sync"

type Client struct {
	Name string `json:"name"`
}

type Leader struct {
	Client
}
type Soldier struct {
	Client
	Ip string
}

var instance *Camp
var once sync.Once

type JsonCamp struct {
	Leader   Leader    `json:"leader"`
	Soldiers []Soldier `json:"soldiers"`
}

type Camp struct {
	Leader   Leader
	Soldiers []Soldier
}

func NewCamp(leaderName string) *Camp {
	once.Do(func() {
		instance = &Camp{
			Leader: Leader{Client{Name: leaderName}},
		}
	})
	return instance
}

func GetCamp() *Camp {
	if instance == nil {
		c := NewCamp("default")
		return c
	}
	return instance
}

func (c *Camp) SetLeader(l Leader) {
	c.Leader = l
}

func (c *Camp) AddSoldier(s Soldier) {
	c.Soldiers = append(c.Soldiers, s)
}

func (c *Camp) RemoveSoldier(ip string) bool {
	for i, sl := range c.Soldiers {
		if sl.Ip == ip {
			c.Soldiers = append(c.Soldiers[:i], c.Soldiers[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Camp) IsSoldierInCamp(ip string) bool {
	for _, sl := range c.Soldiers {
		if sl.Ip == ip {
			return true
		}
	}
	return false
}
