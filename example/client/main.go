package main

import (
	"bytes"
	"flag"
	"io"
	"net/http"
	"sync"

	quic "github.com/gpfnw/quic-go"
	"github.com/gpfnw/quic-go/h2quic"
	"github.com/gpfnw/quic-go/internal/protocol"
	"github.com/gpfnw/quic-go/internal/utils"
)

func main() {
	verbose := flag.Bool("v", false, "verbose")
	tls := flag.Bool("tls", false, "activate support for IETF QUIC (work in progress)")
	flag.Parse()
	urls := flag.Args()

	if *verbose {
		utils.SetLogLevel(utils.LogLevelDebug)
	} else {
		utils.SetLogLevel(utils.LogLevelInfo)
	}
	utils.SetLogTimeFormat("")

	versions := protocol.SupportedVersions
	if *tls {
		versions = append([]protocol.VersionNumber{protocol.VersionTLS}, versions...)
	}

	roundTripper := &h2quic.RoundTripper{
		QuicConfig: &quic.Config{Versions: versions},
	}
	defer roundTripper.Close()
	hclient := &http.Client{
		Transport: roundTripper,
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, addr := range urls {
		utils.Infof("GET %s", addr)
		go func(addr string) {
			rsp, err := hclient.Get(addr)
			if err != nil {
				panic(err)
			}
			utils.Infof("Got response for %s: %#v", addr, rsp)

			body := &bytes.Buffer{}
			_, err = io.Copy(body, rsp.Body)
			if err != nil {
				panic(err)
			}
			utils.Infof("Request Body:")
			utils.Infof("%s", body.Bytes())
			wg.Done()
		}(addr)
	}
	wg.Wait()
}
