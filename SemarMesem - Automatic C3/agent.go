package main

import (
        "net"
        "strconv"
        "fmt"
        "time"
        cf "github.com/redcode-labs/Coldfire"
        ps "github.com/kotakanbe/go-pingscanner"
)

func PrintWarnMessage(msg string){
        dt := time.Now()
        formatMessage := "["+dt.Format(time.UnixDate)+"] [SemarMesemC3] [WARNING] "
        fmt.Printf(formatMessage+msg)
} 

func PrintErrorMessage(msg string){
	dt := time.Now()
	formatMessage := "["+dt.Format(time.UnixDate)+"] [SemarMesemC3] [ERROR] "
	fmt.Printf(formatMessage+msg)
}

func PrintMessage(msg string) {
        dt := time.Now()
        formatMessage := "["+dt.Format(time.UnixDate)+"] [SemarMesemC3] [INFO] "
        fmt.Printf(formatMessage+msg)
}

func MessageOnError(e error) {
        if e != nil {
                PrintMessage(e.Error())
        }
}

func tcpConnect(host string,port string) (net.Conn, error) {
    timeoutSecond := 10 * time.Second // Default Timeout Change if Need 
    conn , err := net.DialTimeout("tcp",net.JoinHostPort(host,port),timeoutSecond)
        if err != nil {
                return nil,err
        }
    return conn,nil
}

func getPortNumbers(start_number int, last_number int) []int{
        var list_allport []int
        for port := start_number; port <= last_number; port++{
        list_allport = append(list_allport,port)
        }
        return list_allport
}

func getAllOpenedPorts(host string) []string {
	
 var allopenedports []string	
 all_port := getPortNumbers(0, 65535)

 for _ , x:= range all_port {
      
    conn, _:= tcpConnect(host,strconv.Itoa(x))

        if conn != nil {

          defer conn.Close()
          PrintMessage("Local Opened Ports "+conn.RemoteAddr().String())
   	  allopenedports = append(allopenedports,conn.RemoteAddr().String())

	}
   }
	return allopenedports
}


func getAllDiscoveredIp() ([]string, error)  {  
  var allDiscoverIps []string
  scanner := ps.PingScanner{CIDR:cf.GetLocalIp()+"/24",PingOptions: []string{"-c1"},NumOfConcurrency:50}

  discoverIps, err := scanner.Scan()

  if err != nil {
    PrintMessage("Error When Disconver Host")
    return nil,err
  }

  if len(discoverIps) > 0{
      for _, ip := range discoverIps {
	allDiscoverIps = append(allDiscoverIps,ip)
	}
  }
   return allDiscoverIps,nil 
}

func runReconnaisance() {

	PrintMessage("Checking Administration Privilleges...")

	if cf.IsRoot() {
		PrintMessage("Your Privillege is Administration Privilleges")
	} else {
		PrintMessage("Your Privillege is not Administration Privilleges")
	}
	
	PrintMessage("Getting Local Open Ports...")

	allopenports := getAllOpenedPorts(cf.GetLocalIp())
	
	if allopenports != nil {
		PrintMessage("There's no Opened Ports")
	}

	if len(allopenports) > 0 {

		for _, x := range allopenports {
			PrintMessage("Opened Ports: "+x)
		}

	} else {
		PrintMessage("There's no opened Ports")
	}

	PrintMessage("Checking Discover IPs...")

	discoveredIps, err := getAllDiscoveredIp()

	if err != nil {
		PrintErrorMessage("Error when Discover Ips")
	}
	
	if discoveredIps != nil {
		PrintMessage("Not Discoverd IPs Host")
	}	

	if len(discoveredIps) > 0 {

		for _, x := range discoveredIps {
			PrintMessage("Discover IPs "+x)
		}

	}	 				
}

func main(){
  runReconnaisance()
}
