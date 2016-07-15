package main
import (
	"strings"
	"fmt"
)
func main() {
	line := "[1790]:d10:mp10165:7263.868126398500:LINK(1070):CAPLEN( 256): IPv4: UDP: 192.168.186.251:37717 --> 192.168.186.221:1500 len=1036 check=63034"
	in := strings.Split(line, "=")
	fmt.Println(in[2])
	fmt.Println(in)
}
