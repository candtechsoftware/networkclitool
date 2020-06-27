package finder

import (
    "net"
    "fmt"
)

// FindMX will find return an error or find the mx records as a slice 
func FindMX(str string) error {
    mx, err := net.LookupMX(str)
    if err != nil {
        fmt.Println(err)
        return err
    }
    for _, v := range mx{
        fmt.Println("HOST: ",v.Host,"Pref: " ,v.Pref)
    }
    return nil
}
// FindCname will return an error or find the cname
func FindCname(str string) error {
    cname, err := net.LookupCNAME(str)
    if err != nil {
        fmt.Println(err)
        return err 
    }
    fmt.Println("CNAME: ",cname)
    return nil 
}

// FindIP will return an error or find the ip address
func FindIP(str string) error {
    ip, err := net.LookupIP(str) 
    if err != nil {
        fmt.Println(err)
    }
    for _, v := range ip {
        fmt.Println("IP: ", v)
    }
    return nil
} 


// FindNS will return error or find a slice of ns records
func FindNS(str string) error {
    ns, err := net.LookupNS(str)
    if err != nil {
        return err 
    }
    for _, v := range ns{
        fmt.Println("NS: ", v.Host)
        }
        return nil 
    
}


// FindAll will return all info
func FindAll(str string) ([]error) {
    return []error {
        FindIP(str),
        FindCname(str),
        FindNS(str), 
        FindMX(str),
    }
}