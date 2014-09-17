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
	err := remail.Serve(remail.ListenWithPrint)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
