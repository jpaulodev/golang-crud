package server

import mgo "gopkg.in/mgo.v2"

func GetSession() *mgo.Session {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	return session
}
