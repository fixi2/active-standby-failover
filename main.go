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
	zkConn *zk.Conn // zookeeper connection
}

func (app SimpleApplication) Run() {
	lock := zk.NewLock(app.zkConn, "/lock", zk.WorldACL(zk.PermAll))

	for {
		lock.Lock()
		exist, _, w, err := app.zkConn.ExistsW("/master")
		if err != nil {
			log.Fatalln(err)
		}
		lock.Unlock()

		wg := sync.WaitGroup{}
		wg.Add(1)

		if !exist {
			// application is Active
			app.zkConn.Create("/master", []byte("master_test"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
			fmt.Println("This Process's Role is Active")

			// Destroy Active app
			time.Sleep(time.Second * 5)
			os.Exit(-1)
			return
		}

		// application is Standby
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
