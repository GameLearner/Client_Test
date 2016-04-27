package main
import "Server/Network"
import "Server/PBProto"
import "net"
import "fmt"
import "sync/atomic"

var closeNum int32
var maxConnects int
var exitCh chan struct{}
func init()  {
    maxConnects = 10000;
    exitCh = make(chan struct{})
}

func sendPacket(session *Network.Session) {
        packet := &PBProto.Login{
            Name : "test",
            Passwd : "md5",
        }
        
        proto := new(Network.Protocol)
        proto.ID = Network.LoginID;
        proto.Packet = packet;
        session.SendPacket(proto);
    }
func onSessionClose(session *Network.Session)  {
    atomic.AddInt32(&closeNum, 1)
    if int(closeNum) == maxConnects {
        exitCh <- struct{}{}
    }
}
    
func main()  {
    for i := 0; i < maxConnects; i++ {
        conn, err := net.Dial("tcp", "127.0.0.1:9999")
        if nil != err {
            fmt.Println("connect remote error : " + err.Error())
            return
        }
        
        session, _ := Network.NewSession(conn, onSessionClose)
        
        sendPacket(session)
        
        go session.Run();
    }
    select {
        case <- exitCh:
        {
            return
        }
    }
}