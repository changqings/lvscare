module github.com/labring/lvscare

go 1.22

require (
	github.com/lithammer/dedent v1.1.0
	github.com/moby/ipvs v1.0.1
	github.com/spf13/cobra v1.2.1
	github.com/vishvananda/netlink v1.1.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vishvananda/netns v0.0.0-20191106174202-0a2b9b5464df // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
)

replace (
	github.com/moby/ipvs => github.com/moby/ipvs v1.0.1
	github.com/vishvananda/netlink => github.com/vishvananda/netlink v1.1.0
	github.com/vishvananda/netns => github.com/vishvananda/netns v0.0.0-20200728191858-db3c7e526aae
)
