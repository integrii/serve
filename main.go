// Serves the current working directory over HTTP (static file server).  Has a directory listing and all that stuff.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var path string
var listenAddress string
var spin *spinner.Spinner

func init() {

	// Use the working directory as the default location to serve
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not determine current working directory.", err)
		os.Exit(1)
	}

	// input flags
	flag.StringVar(&listenAddress, "l", ":8000", "The address for the server to listen on. Examples: :80, 127.0.0.1:8000")
	flag.StringVar(&path, "p", wd, "The path for the server to serve.")
	flag.Parse()
}

func main() {

	// watch for kill signals and exit nicely
	go watchForKill()

	// setup a spinner
	spin = spinner.New(spinner.CharSets[14], time.Millisecond*50)
	spin.Color("green")
	spin.FinalMSG = "" // causes it to erase the current message when stopped

	// formulate the proper clickable listen address for output
	var listenAddressClickable string
	if len(strings.Split(listenAddress, ".")) < 4 {
		listenAddressClickable = "http://0.0.0.0" + listenAddress
	} else {
		listenAddressClickable = "http://" + listenAddress
	}

	// configure the spinner output and start it up
	var spinnerMessage string
	spinnerMessage = color.WhiteString(" %s", "Server running at ")
	spinnerMessage = spinnerMessage + color.YellowString("%s ", path)
	spinnerMessage = spinnerMessage + color.WhiteString("%s ", "on")
	spinnerMessage = spinnerMessage + color.GreenString("%s", listenAddressClickable)
	spin.Suffix = spinnerMessage
	spin.Start()

	// initialze a file server handler
	http.Handle("/", http.FileServer(http.Dir(path)))
	err := http.ListenAndServe(listenAddress, nil)
	spin.Stop()
	if err != nil {
		fmt.Println("Server exited with error: ", err)
		os.Exit(254)
	}
	os.Exit(0)
}

// watchForKill watches for kill and interrupt signals
func watchForKill() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	<-c
	spin.Stop()
	os.Exit(0)
}
