package server

import (
	"example/hello/app"
	"fmt"
	"github.com/urfave/negroni"
	"strconv"
)

const addrFormat = ":%s"

func StartAPIServer(dependency *app.Dependency) {
	serverConfig := dependency.Config.ServerConfiguration()

	server := negroni.New(negroni.NewRecovery())

	server.UseHandler(routes(dependency))

	addr := fmt.Sprintf(addrFormat, strconv.Itoa(serverConfig.Port))

	server.Run(addr)
}
