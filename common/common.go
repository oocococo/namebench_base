package common

type NameserverTable struct {
	Nameserver string
	Ch         chan int64
	Rtt        int64
}
