package main

import (
	"github.com/urfave/cli"
	"os"
	"uk/purplebooth/flattrackstats-bout-rss/server"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Usage = "Convert Flattrackstats bout status page to an RSS feed to use with IFTTT or other things."

	app.Commands = []cli.Command{
		{
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "port",
					EnvVar: "PORT",
					Value:  "8081",
					Usage:  "Port to listen on",
				},
				cli.StringFlag{
					Name:   "feed-url",
					EnvVar: "FEEDURL",
					Value:  "https://rollerderbyrss.purplebooth.co.uk",
					Usage:  "URL the feed is listening on",
				},
				cli.StringFlag{
					Name:   "flattrackstats-url",
					EnvVar: "FLATTRACKSTATSURL",
					Value:  "http://flattrackstats.com",
					Usage:  "URL for flattrackstats",
				},
			},
			Name:    "listen",
			Aliases: []string{"l"},
			Usage:   "Listen for connections",
			Action:  listen,
		},
	}

	app.Run(os.Args)
}

func listen(c *cli.Context) error {
	server := server.FlattrackstatRssServer{
		Port:              c.Int("port"),
		FeedUrl:           c.String("feed-url"),
		FlattrackstatsUrl: c.String("flattrackstats-url"),
	}
	server.Listen()
	return nil
}
