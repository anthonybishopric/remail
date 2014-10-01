package main

import (
	"github.com/anthonybishopric/remail/pkg"

	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetOutput(ioutil.Discard)

	listen := os.Getenv("SMTP_PORT")
	if listen == "" {
		listen = ":25"
	}

	err := remail.Serve(listen, remail.ListenWithPrint)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
