package main

import (
	r "remidial/internal/adapters/http"
)

func main() {
	e := r.InitRoutes()
	e.Debug = true
	e.Logger.Fatal(e.Start(":3000"))
}
