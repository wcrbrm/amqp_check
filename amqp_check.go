// Copyright 2017-2020 Tensigma Ltd. All rights reserved.
// Use of this source code is governed by Microsoft Reference Source
// License (MS-RSL) that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var app = cli.App("amqp_check", "amqp connection checker")
var (
	appRabbitConnection = app.String(cli.StringOpt{
		Name:   "R rabbitmq-connection",
		Desc:   "Specify RabbitMQ connection string.",
		EnvVar: "APP_RABBITMQ_CONNECTION",
		Value:  "amqp://localhost:5672?connection_attempts=5&retry_delay=5",
	})
)

func startApplication() {
	fmt.Printf("TESTING %v\n", *appRabbitConnection)
	_, err := amqp.Dial(*appRabbitConnection)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("CONNECTED\n")
	}
}

func main() {
	app.Action = startApplication
	if err := app.Run(os.Args); err != nil {
		log.Fatalln("[ERR]", err)
	}
}
