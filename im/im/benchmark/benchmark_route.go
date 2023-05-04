package benchmark

import (
	"imservice/im"
	"imservice/im/route"
	"time"
)
import "flag"
import "log"
import "bytes"

var route_addr string = "127.0.0.1:4444"
var appid int64 = 8

func Dispatch(amsg *route.RouteMessage) {
	log.Printf("amsg appid:%d receiver:%d", amsg.appid, amsg.receiver)
}
func main() {
	flag.Parse()

	channel1 := im.NewChannel(route_addr, Dispatch, nil, nil)
	channel1.Start()

	channel1.Subscribe(appid, 1000, true)

	time.Sleep(1 * time.Second)

	channel2 := im.NewChannel(route_addr, Dispatch, nil, nil)
	channel2.Start()

	im := &im.IMMessage{}
	im.sender = 1
	im.receiver = 1000
	im.content = "test"
	msg := &im.Message{cmd: im.MSG_IM, body: im}

	mbuffer := new(bytes.Buffer)
	im.WriteMessage(mbuffer, msg)
	msg_buf := mbuffer.Bytes()

	amsg := &im.RouteMessage{}
	amsg.appid = appid
	amsg.receiver = 1000
	amsg.msg = msg_buf
	channel2.Publish(amsg)

	time.Sleep(3 * time.Second)

	channel1.Unsubscribe(appid, 1000, true)

	time.Sleep(1 * time.Second)
}
