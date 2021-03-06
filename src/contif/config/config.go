package config

import (
	"flag"

	"gopkg.in/fzerorubigd/onion.v2"
)

var (
	Verbose      = flag.Bool("v", false, "Verbose status of connection")
	DialTimeout  = flag.Int("t", 10, "Timeout of sytem on dial host in seconds")
	Domain       = flag.String("d", "google.com", "Critiria doamin for checking connectivity")
	Port         = flag.String("p", "80", "Port of host which you wanna dial with it")
	CycleTimeout = flag.Int("c", 1, "Cycle time to check connectivity in seconds")
)

const (
	organization = "Mimtim"
	appName      = "Netchecker"
)

var Config AppConfig
var o = onion.New()

type AppConfig struct {
	DevelMode    bool   `onion:"devel_mode"`
	Verbose      bool   `onion:"verbose"`
	DialTimeout  int    `onion:"dial_timeout"`
	Domain       string `onion:"domain"`
	Port         string `onion:"port"`
	CycleTimeout int    `onion:"cycle_timeout"`
}

func defaultLayer() onion.Layer {
	d := onion.NewDefaultLayer()
	err := d.SetDefault("devel_mode", true)
	panicOnErr(err)
	err = d.SetDefault("verbose", false)
	panicOnErr(err)
	err = d.SetDefault("dial_timeout", 10)
	panicOnErr(err)
	err = d.SetDefault("domain", "google.com")
	panicOnErr(err)
	err = d.SetDefault("cycle_timeout", 1)
	panicOnErr(err)
	return d

}
func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
