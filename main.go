package main

import (
	"time"

	"./utils"
	"gopkg.in/mgo.v2"
)

var isDropMe, insertMore = true, true

func main() {
	//mqtt.DEBUG = log.New(os.Stderr, "DEBUG ", log.Ltime)

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"127.0.0.1:27017"},
		Direct:    false,
		FailFast:  true,
		Timeout:   10 * time.Second,
		PoolLimit: 0,
		Username:  "",
		Password:  "",
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//utils.Execute(session)
	//utils.QueryDocuments(session)
	utils.Aggregate(session)
	//utils.UpdateValueArray(session)

}

