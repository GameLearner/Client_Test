package main
import "Server/Network"
import "Server/PBProto"
import "net"
import "fmt"
func main()  {
    conn, err := net.Dial("tcp", "127.0.0.1:9999")
    if nil != err {
        fmt.Println("connect remote error : " + err.Error())
        return
    }
    
    session, _ := Network.NewSession(conn)
    
    go func() {
        packet := &PBProto.Login{
            Name : "test",
            Passwd : "md5",
        }
        
        proto := new(Network.Protocol)
        proto.ID = Network.LoginID;
        proto.Packet = packet;
        session.SendPacket(proto);
    }()
    session.Run();

}