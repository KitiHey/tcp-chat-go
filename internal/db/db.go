package db

import (
	"fmt"
	"math/rand"
	"net"
)

type UUID [16]byte
type Users map[UUID]*User

type User struct {
	uuid UUID
	nick string
	channel string
	conn net.Conn
}

type Database struct {
	users Users
	channels map[string]Users
}

func NewDatabase() *Database {
	db := &Database{channels: make(map[string]Users), users: make(Users)}
	db.NewChannel("general")
	return db
}

// Channels
func (db *Database) NewChannel(newchannel string) {
	db.channels[newchannel] = Users{}
}

func (db *Database) GetChannel(channel string) (Users, bool) {
	_, ok := db.channels[channel]
	if !ok {
		return nil, false
	}
	return db.channels[channel], true
}

func (db *Database) AddUsrToChnl(usr *User, channel string) error {
	chn, ok := db.GetChannel(channel)
	if !ok {
		return fmt.Errorf("Channel %s doesnt exist", channel)
	}

	chnprev, ok := db.GetChannel(usr.channel)
	if ok {
		delete(chnprev, usr.uuid)
	}
	chn[usr.uuid] = usr
	usr.channel = channel
	return nil
}

func (db *Database) Send(usr *User, send []byte) error {
	chn, ok := db.GetChannel(usr.Channel())
	if !ok {
		return fmt.Errorf("Channel %s doesnt exist", usr.Channel())
	}
	for _, chnUsr := range chn {
		if chnUsr.uuid == usr.uuid {
			continue
		}
		msg := fmt.Sprintf("%s (%s): %s", usr.nick, usr.channel, string(send))
		chnUsr.conn.Write([]byte(msg))
	}
	return nil
}

// Users
func (db *Database) NewUser(conn net.Conn) (uuid UUID) {
	for i := range uuid {
		uuid[i] = byte(rand.Intn(256))
	}
	usr := &User{uuid: uuid, conn: conn}
	db.AddUsrToChnl(usr, "general")
	usr.SetNick("unknown")
	db.users[uuid] = usr
	return uuid
}

func (usr *User) SetNick(nick string) {
	usr.nick = nick
}

func (db *Database) GetUser(uuid UUID) (*User, bool) {
	usr, ok := db.users[uuid]
	return usr, ok
}

func (usr *User) Channel() string {
	return usr.channel
}
