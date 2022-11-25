package server

import (
	"example/hello/config"
	"fmt"
	"github.com/urfave/negroni"
	"strconv"
)

const addrFormat = ":%s"

func StartAPIServer(conf config.Configurations) {
	server := negroni.New(negroni.NewRecovery())

	server.UseHandler(routes())

	addr := fmt.Sprintf(addrFormat, strconv.Itoa(conf.Port()))

	server.Run(addr)
}
