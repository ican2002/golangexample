package main
 
import (
    "fmt"
    "net"
    "os"
)
 
/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
 
func main() {
    /* Lets prepare a address at any address at port 40000*/   
    ServerAddr,err := net.ResolveUDPAddr("udp",":40000")
    CheckError(err)
 
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 4096)
 
    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)
 
        if err != nil {
            fmt.Println("Error: ",err)
        }
		n, err = ServerConn.WriteTo(buf[0:n], addr); 
		if err != nil {
			fmt.Println("reply error:", err)	
		}
	}
}
