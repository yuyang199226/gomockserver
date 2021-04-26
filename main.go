package main
import (
    "flag"
    "fmt"
    "github.com/yuyang199226/gomockserver/config"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)


var configFile *string = flag.String("config", "mock.json", "config file path")
func main() {
    flag.Parse()
    conf,err := config.LoadConfigFromFile(*configFile)
    if err != nil {
        panic(err)
    }
    handlermap := make(map[string]string)
    handler := func(w http.ResponseWriter, req *http.Request) {
        path := req.URL.Path
        method := req.Method
        log.Printf("receive request %s %s %s\n",req.Proto,  req.Method, req.URL.Path)
        b,_ := ioutil.ReadAll(req.Body)
        log.Printf("receive request %s\n", b)
        key := fmt.Sprintf("%s-%s", path, method)
        if resp, ok := handlermap[key]; ok {
            io.WriteString(w, resp+"\n")
        } else {
            http.NotFound(w, req)
            return
        }
    }

    for _, hc := range conf.Handlers {
        key := fmt.Sprintf("%s-%s", hc.Path, strings.ToUpper(hc.Method))
        handlermap[key] = hc.Response
        http.HandleFunc(hc.Path, handler)
    }
    addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
    fmt.Printf("listen: %s\n", addr)
    if conf.UseHttp2 {
        srv := &http.Server{Addr:addr, Handler: nil}
        fmt.Println("start http2 server")
        log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
    }else {
        log.Fatal(http.ListenAndServe(addr, nil))
    }



}
