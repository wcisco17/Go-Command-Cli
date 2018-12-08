package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Look up Cli"
	app.Usage = "Let's you query IP, CNAMES, MX records and Name Servers!"

	var input string

	fmt.Println("Enter website name: ")

	_, er := fmt.Scan(&input)

	if er != nil {
		fmt.Println(er)
	}

	myFlag := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: input,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the name Servers for a particular Host",
			Flags: myFlag,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for _, value := range ns {
					fmt.Println(value.Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP address for particualar host",
			Flags: myFlag,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for _, value := range ip {
					fmt.Println(value)
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: myFlag,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records of host",
			Flags: myFlag,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))

				if err != nil {
					fmt.Println(err)
					return err
				}

				for _, value := range mx {
					fmt.Println("Host: ", value.Host, "Pref: ", value.Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
