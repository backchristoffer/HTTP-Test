package main

import (
	"fmt"
	"net/http"
	"github.com/shirou/gopsutil/v3/mem"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "This is a Hello message\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func virtmem(w http.ResponseWriter, req *http.Request) {
	v, _ := mem.VirtualMemory()
    fmt.Fprintf(w,"Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
    fmt.Fprintf(w,"%v\n", v)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/mem", virtmem)
	http.ListenAndServe(":8080", nil)
}
