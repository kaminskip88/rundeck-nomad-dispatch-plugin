package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kaminskip88/nomad-dispatcher/dispatcher"
)

const (
	defaultNomadAddr   = "http://localhost:4646"
	defaultNomadRegion = "global"
)

func main() {
	var c dispatcher.Config
	// rundeck settings > nomad native env > default values
	var addr string
	addr, ok := os.LookupEnv("RD_CONFIG_NOMAD_ADDR")
	if !ok {
		addr, ok = os.LookupEnv("NOMAD_ADDR")
		if !ok {
			addr = defaultNomadAddr
		}
	}
	c.Addr = addr

	var reg string
	reg, ok = os.LookupEnv("RD_CONFIG_NOMAD_REGION")
	if !ok {
		reg, ok = os.LookupEnv("NOMAD_REGION")
		if !ok {
			reg = defaultNomadAddr
		}
	}
	c.Region = reg

	// got default values, so always set
	c.Job, _ = os.LookupEnv("RD_CONFIG_JOB_ID")
	c.Interval, _ = os.LookupEnv("RD_CONFIG_CLIENT_INTERVAL")

	// parse int
	intE, _ := os.LookupEnv("RD_CONFIG_CLIENT_RETRY")
	int, err := strconv.Atoi(intE)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	c.AllocRetry = int
	c.JobRetry = c.AllocRetry

	// parse map
	mE, ok := os.LookupEnv("RD_CONFIG_JOB_META")
	if ok {
		m, err := godotenv.Unmarshal(mE)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		c.Meta = m
	}

	// set payload
	pl, ok := os.LookupEnv("RD_CONFIG_JOB_PAYLOAD")
	if ok {
		c.Payload = pl
	}

	// set debug
	dbg, ok := os.LookupEnv("RD_JOB_LOGLEVEL")
	if ok {
		if dbg == "DEBUG" {
			c.LogDebug = true
		}
	}

	// dispatch
	d, err := dispatcher.NewDispatcher(&c)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = d.Dispatch()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
