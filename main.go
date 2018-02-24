package main

import (
	"log"
	"os"

	"github.com/covrom/cloud-torrent/server"
	"github.com/jpillora/opts"
)

var VERSION = "0.8.24fsr" //set with ldflags

func main() {
	s := server.Server{
		Title:      "Cloud Torrent",
		Port:       3000,
		ConfigPath: "./cloud-torrent.json",
	}

	wd, _ := os.Getwd()
	log.Println("Working path", wd)

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
