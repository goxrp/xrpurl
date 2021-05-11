package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/grokify/simplego/encoding/jsonutil"
	"github.com/grokify/simplego/net/http/httpsimple"
	"github.com/jessevdk/go-flags"

	gorippled "github.com/goxrp/go-rippled"
	"github.com/goxrp/go-rippled/data"
)

type Options struct {
	Method   string `short:"m" long:"method" description:"Method" required:"true"`
	Data     string `short:"d" long:"data" description:"Data"`
	TestData []bool `short:"t" long:"testdata" description:"Use Test Data"`
	Exec     []bool `short:"x" long:"exec" description:"Execute API Call"`
	Verbose  []bool `short:"v" long:"verbose" description:"Verbose"`
	Pretty   []bool `short:"p" long:"pretty" description:"Pretty Print Result"`
}

const (
	RippledJsonRpcUrl = "https://s1.ripple.com:51234/"
)

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	var reqBody interface{}

	if len(opts.Data) > 0 {
		reqBodyStruct, err := gorippled.BuildJsonRpcRequestBody(opts.Method, []byte(opts.Data))
		if err != nil {
			log.Fatal(err)
		}
		reqBody = reqBodyStruct
	} else if len(opts.TestData) > 0 {
		reqBodyBytes, err := data.ExampleJsonRequest(opts.Method, "")
		if err != nil {
			log.Fatal(err)
		}
		reqBody = reqBodyBytes
		if len(opts.Verbose) > 0 {
			fmt.Println(string(reqBodyBytes))
		}
	}

	if len(opts.Exec) > 0 {
		req := httpsimple.SimpleRequest{
			Method: http.MethodPost,
			URL:    RippledJsonRpcUrl,
			Body:   reqBody,
			IsJSON: true}

		resp, err := httpsimple.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		resBodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if len(opts.Pretty) > 0 {
			fmt.Println(string(jsonutil.PrettyPrint(resBodyBytes, "", "  ")))
		} else {
			fmt.Println(string(resBodyBytes))
		}
	}
	fmt.Println("DONE")
}
