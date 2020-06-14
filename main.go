package main

import (
	"context"
	"github.com/ednailson/api-base-project-go/app"
	"github.com/micro/go-micro/config"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
)

func main() {
	var flagConfig string
	var flagFileName string
	cliApp := cli.NewApp()
	cliApp.Name = ApplicationName
	cliApp.Version = Version + "(" + GitCommit + ")"
	cliApp.EnableBashCompletion = true
	cliApp.Commands = []cli.Command{
		{
			Name:    "config sample generator",
			Aliases: []string{"csg"},
			Action: func(cli *cli.Context) error {
				return configSampleGenerator(flagFileName)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "file-name, f",
					Value:       "./config-sample.json",
					Usage:       "Config sample file name",
					Destination: &flagFileName,
				},
			},
		},
		{
			Name:    "run application",
			Aliases: []string{"run"},
			Action: func(cli *cli.Context) error {
				return runApplication(flagConfig)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config, c",
					Value:       "./config.json",
					Usage:       "config cliApp file",
					Destination: &flagConfig,
				},
			},
		},
	}
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runApplication(flagConfig string) error {
	var cfg app.Config
	if err := errors.Wrap(config.LoadFile(flagConfig), "failed to get config file"); err != nil {
		return err
	}
	if err := errors.Wrap(config.Scan(&cfg), "failed to read from config file"); err != nil {
		return err
	}
	application := app.LoadApp(cfg)
	ctx := gracefullyShutdown()
	chErr := application.Run()
	select {
	case err := <-chErr:
		if err != nil {
			application.Close()
			return err
		}
	case <-ctx.Done():
		application.Close()
	}
	return nil
}

func gracefullyShutdown() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Println("gracefully shutdown")
		cancel()
	}()
	return ctx
}

func configSampleGenerator(flagFileName string) error {
	return errors.Wrap(app.NewConfigFile(flagFileName), "could not create a new config file")
}
