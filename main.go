package main

import (
    "os"
    "log"
    "net"
    "fmt"
    "github.com/urfave/cli"
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
            ns, err := net.LookupNS(c.String("host"))
            if err != nil {
                return err 
            }
            for _, v := range ns{
                fmt.Println(v.Host)
                }
                return nil 
            },
        },
        {
            Name: "ip",
            Usage: "Looks Up IP for a host",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                ip, err := net.LookupIP(c.String("host"))
                if err != nil {
                    fmt.Println(err)
                }
                for _, v := range ip {
                    fmt.Println(v)
                }
                return nil
            },
        },
        {
            Name: "cname",
            Usage: "Looks up the CNAME",
            Flags: myFlags,
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
            Name: "mx",
            Usage: "Looks up mx records ",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                mx, err := net.LookupMX(c.String("host"))
                if err != nil {
                    fmt.Println(err)
                    return err
                }
                for _, v := range mx{
                    fmt.Println(v.Host, v.Pref)
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
