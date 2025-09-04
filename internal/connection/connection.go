package connection

import (
	"fmt"
	"github.com/KitiHey/tcp-chat-go/internal/db"
	"net"
)

func HandleCon(conn net.Conn, database *db.Database) {
	uuid := database.NewUser(conn)
	usr, _ := database.GetUser(uuid)
	chn := usr.Channel()
	conn.Write([]byte(fmt.Sprintf("Connected to channel '%s'", chn)))
	for {
		var buffer = make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			break;
		}
		database.Send(usr, buffer)
		fmt.Println(string(buffer))
	}
}
