package main
import ("fmt"
	"net"
	
)

func main() {
	fmt.Printf("Welcome to a PortScanner written in Go\n")
	fmt.Printf("This is a Basic SOC Analysis Project\n")
	fmt.Printf("We are going to be seeing if Google has port 80 open\n")
	
	host := "www.cia.gov" //Target Host
	port := "80" //Port to Scan
	
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Port %s is closed \n", port)
		return
	}
	defer conn.Close()
	fmt.Printf("Port %s is open\n", port) //Returns if port is open/closed
}
