package main

import (
    "os"
    "log"
    "github.com/urfave/cli"
    finder "network-cli/finder"
)



func main() {
    app := cli.NewApp()
    app.Name = "Website Look up"
    app.Usage = "Lets you look Ips, Cname, Mx reccords and Name servers"

    myFlags := []cli.Flag{
        cli.StringFlag {
            Name: "host",
            Value: "candtechsoftware.com",
        },
    }

    app.Commands = []cli.Command{
        {
        Name: "ns",
        Usage: "Looks up names servers",
        Flags: myFlags,
        Action: func(c *cli.Context) error {
            return finder.FindNS(c.String("host"))
            },
        },
        {
            Name: "ip",
            Usage: "Looks Up IP for a host",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                return finder.FindIP(c.String("host"))
            },
        },
        {
            Name: "cname",
            Usage: "Looks up the CNAME",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                return finder.FindCname(c.String("host"))
            },
        },
        {
            Name: "mx",
            Usage: "Looks up mx records ",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                return finder.FindMX(c.String("host"))
             },

        },
        {
            Name: "all",
            Usage: "Finds MX, CNAME, IP and NS Records",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                 finder.FindAll(c.String("host"))
                
                 return nil 
            },
        },
            
    }
    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err) 
}   
}
