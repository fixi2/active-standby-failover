package main

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"log"
	"os"
	"sync"
	"time"
)

type SimpleApplication struct {
	Role   string   // Active, StandBy or Master, Slave
	zkConn *zk.Conn // zookeeper connection
}

func (app SimpleApplication) Run() {

	for {
		exist, _, w, err := app.zkConn.ExistsW("/master")
		if err != nil {
			log.Fatalln(err)
		}

		wg := sync.WaitGroup{}
		wg.Add(1)

		if !exist {
			app.zkConn.Create("/master", []byte("master_test"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
			app.Role = "Active"
			fmt.Println("This Process's Role is Active")

			// master destory
			time.Sleep(time.Second * 5)
			os.Exit(-1)
			return
		}

		app.Role = "StandBy"
		fmt.Println("This Process's Role is Standby")

		// waiting for watch notify
		<-w
	}

	return
}

func main() {

	app := SimpleApplication{}
	app.zkConn = ConnectZookeeper()

	app.Run()
}

var servers []string

func ConnectZookeeper() *zk.Conn {
	conn, _, err := zk.Connect(servers, time.Second * 3)
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func init() {
	servers = []string{"10.113.78.147:2181", "10.113.79.117:2181", "10.113.97.243:2181"}
}
