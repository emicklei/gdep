package main

import (
	"bytes"
	"fmt"
	"log"
	"reflect"

	"github.com/urfave/cli/v2"
)

func logBegin(c *cli.Context) func() {
	buf := new(bytes.Buffer)
	fmt.Fprint(buf, "[gdep]")

	appendFlag := func(each cli.Flag) {
		fv := reflect.ValueOf(each)
		hide := reflect.Indirect(fv).FieldByName("Hidden").Bool()
		name := each.Names()[0]
		value := c.Generic(name)
		if hide {
			value = "**hidden**"
		}
		fmt.Fprintf(buf, " %s=%v", name, value)
	}
	fmt.Fprintf(buf, " %s (", c.Command.Name)
	for _, each := range c.App.Flags {
		appendFlag(each)
	}
	for _, each := range c.Command.Flags {
		appendFlag(each)
	}
	fmt.Fprint(buf, " )")
	log.Println(buf.String())
	return func() {
		if err := recover(); err != nil {
			// no way to communicate error to cli so exit here.
			log.Fatalln(c.Command.Name, "FAILED with error:", err)
		}
	}
}

func logEnd(c *cli.Context, err error) error {
	if err != nil {
		log.Printf("[gdep] %s failed:%v\n", c.Command.Name, err)
	} else {
		log.Printf("[gdep] %s done\n", c.Command.Name)
	}
	return err
}
