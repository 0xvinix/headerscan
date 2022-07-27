package main

import (
    "os"
    "fmt"
    "net/http"
    "net/url"
    "time"
    "runtime"
)

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

func init() {
    if runtime.GOOS == "windows" {
        Reset  = ""
        Red    = ""
        Green  = ""
        Yellow = ""
        Blue   = ""
        Purple = ""
        Cyan   = ""
        Gray   = ""
        White  = ""
    }
}

func main() {
    if len(os.Args) == 1 {
        msgBanner(" [❌] Usage: headerscan http://target.com \n")
        os.Exit(0)
    } else {
        target := os.Args[1]
            if (IsUrl(target) == true){
                checkURL(target)
            }else {
                msgBanner(" [❌] Invalid URL! \n")
            }
    }
}

func checkURL(target string){
    msgBanner("")
    resp,err := http.Get(target)
    if resp != nil {
        simpleRecon(target)
    } else {
    fmt.Println(" [❌] "+Red+"Error: "+Reset+"", err)
    }
}

func getPowered(target string) string{
    resp, _ := http.Get(target)
    xpowered := resp.Header.Get("X-Powered-By")
    if len(xpowered) > 1 {
        return string(xpowered)
    }
    return "hidden"
}

func getServer(target string) string {
    resp, _ := http.Get(target)
    server := resp.Header.Get("Server")
    if len(server) > 1 {
        return string(server)
    }
    return "hidden"
}

func xframeOptions(target string) bool {
    resp, _ := http.Get(target)
    clickjacking := resp.Header.Get("X-Frame-Options")
    if len(clickjacking) == 0 {
        return true
    }
    return false
}

func xssProtection(target string) bool {
    resp, _ := http.Get(target)
    xss := resp.Header.Get("X-Xss-Protection")
    if len(xss) == 0 || xss == "0" {
        return true
    }
    return false
}

func transportSecurity(target string) bool {
    resp, _ := http.Get(target)
    hsts := resp.Header.Get("Strict-Transport-Security")
    if len(hsts) == 0 {
        return true
    }
    return false
}

func contentType(target string) bool {
    resp, _ := http.Get(target)
    mime := resp.Header.Get("X-Content-Type-Options")
    if len(mime) == 0 {
        return true
    }
    return false
}

func referrerPolicy(target string) bool{
    resp, _ := http.Get(target)
    rp := resp.Header.Get("Referrer-Policy")
    if len(rp) == 0 {
        return true
    }
    return false
}

func contentSecurity(target string) bool{
    resp, _ := http.Get(target)
    csp := resp.Header.Get("Content-Security-Policy")
    if len(csp) == 0 {
        return true
    }
    return false
}

func simpleRecon(target string){
    dt := time.Now()
    u, _ := url.Parse(target)
    fmt.Printf(" [🎯] "+White+"Target:"+Reset+" %v\n", u.Host)
    fmt.Printf(" [🕰️] "+White+"Started at:"+Reset+" %v\n\n", dt.String())

    fmt.Println(" [📝] Web Technology: ", getPowered(target))
    fmt.Println(" [🖥️] Web Server: ", getServer(target))
    if xframeOptions(target) == true {
        fmt.Println(" [❌] X-Frame Options Header Not Set")
    } else {
        fmt.Println(" [✔️] X-Frame-Options Header OK")
    }

    if xssProtection(target) == true {
        fmt.Println(" [❌] X-XSS-Protection Header Not Set")
    } else {
        fmt.Println(" [✔️] X-XSS-Protection Header OK")
    }

    if transportSecurity(target) == true {
        fmt.Println(" [❌] Strict-Transport-Security Not Set")
    } else {
        fmt.Println(" [✔️] Strict-Transport-Security OK")
    }

    if contentType(target) == true {
        fmt.Println(" [❌] X-Content-Type-Options Header Not Set")
    } else {
        fmt.Println(" [✔️] X-Content-Type-Options OK")
    }

    if referrerPolicy(target) == true {
        fmt.Println(" [❌] Referrer-Policy Header Not Set")
    } else {
        fmt.Println(" [✔️] Referrer-Policy OK")
    }

    if contentSecurity(target) == true {
        fmt.Println(" [❌] Content-Security-Policy Header Not Set")
    } else {
        fmt.Println(" [✔️] Content-Security-Policy OK")
    }

    fmt.Println("\n [😊] Bye!")
}

func IsUrl(str string) bool {
    u, err := url.Parse(str)
    return err == nil && u.Scheme != "" && u.Host != ""
}

func msgBanner(msg string) {
    ascii := 
    `
                       __         
    |_| _  _  _| _  __(_  _  _ __ 
    | |(/_(_|(_|(/_ | __)(_ (_|| |

    Version 1.0
    by @vinix
    `
    fmt.Printf(ascii)
    fmt.Printf("\n [🕷️] Simple scanner for HTTP headers.\n")
    fmt.Printf("\n%v", msg)
}
