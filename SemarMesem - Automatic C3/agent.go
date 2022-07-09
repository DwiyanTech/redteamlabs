package main

import (
        "net"
        "strconv"
        "fmt"
        "time"
        cf "github.com/redcode-labs/Coldfire"
        ps "github.com/kotakanbe/go-pingscanner"
)

func sendMessageToSlack(msg string){


}

func PrintMessage(msg string) {
        dt := time.Now()
        t := dt.Format("15:30")
        fmt.Printf("[%s] [SemarMesemC3] [INFO] %s \n", t, msg)
}

func MessageOnError(e error) {
        if e != nil {
                PrintMessage(e.Error())
        }
}

func tcp_connect(host string,port string) (net.Conn, error) {
    timeoutSecond := 10 * time.Second
    conn , err := net.DialTimeout("tcp",net.JoinHostPort(host,port),timeoutSecond)
        if err != nil {
                return nil,err
        }
    return conn,nil
}

func get_port_numbers(start_number int, last_number int) []int{
        var list_allport []int

        for port := start_number; port <= last_number; port++{
      list_allport = append(list_allport,port)
        }
        return list_allport
}

func getall_openedports(host string) {

        all_port := get_port_numbers(0,65535)
        for _ , x:= range all_port {
                conn, _:= tcp_connect(host,strconv.Itoa(x))
        if conn != nil {
                  defer conn.Close()
          PrintMessage("Local Opened Ports "+conn.RemoteAddr().String())
                }
        }
}


func getAllDiscoveredIp() {

  scanner := ps.PingScanner{CIDR:cf.GetLocalIp()+"/24",PingOptions: []string{"-c1"},NumOfConcurrency:50}

  discoverIps, err := scanner.Scan()

  if err != nil {
    PrintMessage("Error When Disconver Host")
  }

  if len(discoverIps) == 0{
    PrintMessage("No One IP alive")
  }

  if len(discoverIps) > 0{
      for _, ip := range discoverIps {
        PrintMessage("Found Active IP : "+ip)
    }
  }

}

func runNetworkDiscoverScan(){



}

func main(){
   // Fill This
}
