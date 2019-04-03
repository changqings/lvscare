package service

import (
	"fmt"
	"net"
	"strconv"
	"syscall"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
)

//EndPoint  is
type EndPoint struct {
	IP   string
	Port string
}

//Lvser is
type Lvser interface {
	CreateInterface(name string, CIRD string) error
	CreateVirtualServer() error
	AddRealServer(ip, port string) error
	GetVirtualServer() (vs *EndPoint, rs *[]EndPoint)
	RemoveVirtualServer() error
	RemoveRealServer(ip, port string) error
}

type lvscare struct {
	vs      EndPoint
	rs      []EndPoint
	service *Service
}

func (l *lvscare) CreateInterface(name string, CIRD string) error {
	interfa := &netlink.Dummy{LinkAttrs: netlink.LinkAttrs{Name: name}}
	err := netlink.LinkAdd(interfa)
	if err != nil {
		return fmt.Errorf("create net interface failed: %s", err)
	}

	kubeIpvs, err := netlink.LinkByName(name)
	if err != nil {
		return fmt.Errorf("get interface failed: %s", err)
	}

	fmt.Println("cird is: ", CIRD)
	addr, err := netlink.ParseAddr(CIRD)
	if err != nil {
		return fmt.Errorf("CIRD failed: %s", err)
	}
	err = netlink.AddrAdd(kubeIpvs, addr)
	if err != nil {
		return fmt.Errorf("bind IP addr failed: %s", err)
	}

	return err
}

func (l *lvscare) CreateVirtualServer() error {
	handle, err := New("")
	if err != nil {
		return fmt.Errorf("New ipvs handle failed: %s", err)
	}
	p, err := strconv.Atoi(l.vs.Port)
	if err != nil {
		return fmt.Errorf("port is %s : %s", l.vs.Port, err)
	}

	s := Service{
		AddressFamily: nl.FAMILY_V4,
		SchedName:     RoundRobin,
		Protocol:      syscall.IPPROTO_TCP,
		Port:          uint16(p),
		Address:       net.ParseIP(l.vs.IP),
		Netmask:       0xFFFFFFFF,
	}

	err = handle.NewService(&s)
	if err != nil {
		return fmt.Errorf("New ipvs failed: %s", err)
	}

	l.service = &s
	return nil
}

func (l *lvscare) AddRealServer(ip, port string) error {
	handle, err := New("")
	if err != nil {
		return fmt.Errorf("New ipvs handle failed: %s", err)
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("port is %s : %s", port, err)
	}

	if l.service == nil {
		return fmt.Errorf("service is nil: %s", err)
	}

	d := Destination{
		AddressFamily:   nl.FAMILY_V4,
		Address:         net.ParseIP(ip),
		Port:            uint16(p),
		Weight:          1,
		ConnectionFlags: ConnectionFlagMasq,
	}

	err = handle.NewDestination(l.service, &d)
	if err != nil {
		return fmt.Errorf("new destination failed: %s", err)
	}

	return nil
}

func (l *lvscare) GetVirtualServer() (vs *EndPoint, rs *[]EndPoint) {
	return nil, nil
}

func (l *lvscare) RemoveVirtualServer() error {
	return nil
}

func (l *lvscare) RemoveRealServer(ip, port string) error {
	return nil
}

//NewLvscare is
func NewLvscare(ip, port string) Lvser {
	return &lvscare{
		vs: EndPoint{IP: ip, Port: port},
	}
}
