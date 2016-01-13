package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "RRD CLI tool"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "filename, f",
			Value: "",
			Usage: "RDD file name",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "init RRD file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "columns, c",
					Value: "",
					Usage: "columns definition in form: function[:col name],function[:col name],.... Functions: average/avg/sum/min/minimum/max/maximum/count/last",
				},
				cli.StringFlag{
					Name:  "archives, a",
					Value: "",
					Usage: "archives definitions in form: rows:step[:archive name],rows:step[:name]...",
				},
			},
			Action: initDB,
		},
		{
			Name:  "put",
			Usage: "put many values into db (as args)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ts",
					Value: "",
					Usage: "time stamp (in sec, date, N/now/NOW)",
				},
				cli.StringFlag{
					Name:  "columns, c",
					Value: "",
					Usage: "optional destination columns number separated by comma",
				},
			},
			Action: putValues,
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get value from RRD file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ts",
					Value: "",
					Usage: "time stamp (in sec, date, N/now/NOW)",
				},
				cli.StringFlag{
					Name:  "columns, c",
					Value: "",
					Usage: "optional columns to get",
				},
			},
			Action: getValue,
		},
		{
			Name:    "get-range",
			Aliases: []string{"gr"},
			Usage:   "get values from RRD file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "begin, b",
					Value: "",
					Usage: "time stamp (in sec, date, N/now/NOW)",
				},
				cli.StringFlag{
					Name:  "end, e",
					Value: "now",
					Usage: "time stamp (in sec, date, N/now/NOW)",
				},
			},
			Action: getRangeValues,
		},
		{
			Name:   "info",
			Usage:  "show informations about rddfile",
			Action: showInfo,
		},
		{
			Name:   "last",
			Usage:  "get last time stamp from database",
			Action: showLast,
		},
	}
	app.Run(os.Args)
}

func getFilenameParam(c *cli.Context) (string, bool) {
	filename := c.GlobalString("filename")
	if !c.GlobalIsSet("filename") || filename == "" {
		fmt.Println("Missing database file name")
		return "", false
	}
	return filename, true
}