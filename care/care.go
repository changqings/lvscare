package care

import (
	"time"

	"github.com/fanux/lvscare/create"
	"github.com/fanux/lvscare/service"
)

//VsAndRsCare is
func VsAndRsCare(vs string, rs []string, beat int64, path string, schem string) {
	lvs := service.BuildLvscare()
	t := time.NewTicker(time.Duration(beat) * time.Second)
	for {
		select {
		case <-t.C:
			//check virturl server
			service, _ := lvs.GetVirtualServer()
			if service == nil {
				create.VsAndRsCreate(vs, rs, lvs)
			}

			//check real server
			lvs.CheckRealServers(path, schem)
		}
	}

}
