package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pongor "github.com/marceloagmelo/pongor-echo"
	r "gitlab.produbanbr.corp/paas-brasil/go-restore-openshift/router"
	"gitlab.produbanbr.corp/paas-brasil/go-restore-openshift/variaveis"
)

func init() {

	traceHandle := ioutil.Discard
	infoHandle := os.Stdout
	warningHandle := os.Stdout
	errorHandle := os.Stderr

	variaveis.Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	variaveis.Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	variaveis.Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	variaveis.Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	e := r.App
	e.Static("/static", "./static")

	p := pongor.GetRenderer()
	p.Directory = "views"

	e.Renderer = p
	e.Start(":8080")
}
