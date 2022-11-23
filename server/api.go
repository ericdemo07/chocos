package server

import "github.com/urfave/negroni"

func StartAPIServer() {
	println("hello111")
	server := negroni.New(negroni.NewRecovery())

	server.UseHandler(routes())
	server.Run("localhost:3030")
}
