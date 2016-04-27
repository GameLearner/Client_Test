package main
import "Server/Network"
import "Server/PBProto"
import "net"
import "fmt"

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
    
func main()  {
    maxConnects := 1000;
    for i := 0; i < maxConnects; i++ {
        conn, err := net.Dial("tcp", "127.0.0.1:9999")
        if nil != err {
            fmt.Println("connect remote error : " + err.Error())
            return
        }
        
        session, _ := Network.NewSession(conn)
        
        sendPacket(session)
        
        go session.Run();
    }
    for {
        
    }
}