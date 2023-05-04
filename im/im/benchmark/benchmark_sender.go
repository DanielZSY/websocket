package benchmark

import (
	"fmt"
	"imservice/im"
)
import "net"
import "log"
import "runtime"
import "flag"
import "time"
import "math/rand"
import "github.com/gomodule/redigo/redis"

var first int64
var last int64
var host string
var port int
var redis_pool *redis.Pool
var seededRand *rand.Rand

const redis_address = "127.0.0.1:6379"
const redis_password = ""
const redis_db = 0

func init() {
	flag.Int64Var(&first, "first", 0, "first uid")
	flag.Int64Var(&last, "last", 0, "last uid")

	flag.StringVar(&host, "host", "127.0.0.1", "host")
	flag.IntVar(&port, "port", 23000, "port")
}

func send(uid int64) {
	ip := net.ParseIP(host)
	addr := net.TCPAddr{ip, port, ""}

	token, err := login(uid)
	if err != nil {
		log.Println("login error")
		return
	}

	conn, err := net.DialTCP("tcp4", nil, &addr)
	if err != nil {
		log.Println("connect error")
		return
	}
	seq := 1

	auth := &im.AuthenticationToken{token: token, platform_id: 1, device_id: "00000000"}
	im.SendMessage(conn, &im.Message{cmd: im.MSG_AUTH_TOKEN, seq: seq, version: im.DEFAULT_VERSION, body: auth})
	im.ReceiveMessage(conn)

	for i := 0; i < 18000; i++ {
		r := rand.Int63()
		receiver := r%(last-first) + first
		log.Println("receiver:", receiver)
		content := fmt.Sprintf("test....%d", i)
		seq++
		msg := &im.Message{cmd: im.MSG_IM, seq: seq, version: im.DEFAULT_VERSION, flag: 0, body: &im.IMMessage{uid, receiver, 0, int32(i), content}}
		im.SendMessage(conn, msg)
		for {
			ack := im.ReceiveMessage(conn)
			if ack.cmd == im.MSG_ACK {
				break
			}
		}
	}
	conn.Close()
	log.Printf("%d send complete", uid)
}

func main() {
	runtime.GOMAXPROCS(4)
	flag.Parse()
	fmt.Printf("first:%d, last:%d\n", first, last)
	if last <= first {
		return
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	redis_pool = im.NewRedisPool(redis_address, redis_password, redis_db)
	send(1)
}
