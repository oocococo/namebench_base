package query

import (
	"namebench/common"
	"net"

	"github.com/miekg/dns"
)

// RunWithLimit make dns query with limited goroutine
func RunWithLimit(client *dns.Client, url string, table common.NameserverTable, pool chan struct{}) {
	// block if ch is full
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(url), dns.TypeA)
	pool <- struct{}{}
	// fmt.Printf("%v -> %v\n", url, table.Nameserver)
	_, rtt, err := client.Exchange(m, net.JoinHostPort(table.Nameserver, "53"))
	<-pool
	if err != nil {
		// log.Println(err)
		table.Ch <- 2000
	} else {
		table.Ch <- rtt.Milliseconds()
	}
}
