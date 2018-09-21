package main

import(
	"fmt"
	// "net"
	"net/http"
	// "strings"
	"log"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

var cip string

func sayHello(w http.ResponseWriter,r *http.Request){
	// ip := RemoteIp(r)
	ip := r.RemoteAddr
	/*
	if strings.Contains(ip,"[::1]") {
		strings.Replace(ip,"[::1]","127.0.0.1",-1)
	}
	*/
	// fmt.Fprintf(w,ip)
	if cip == ""{
		cip = ip
		fmt.Fprint(w,"minigan buzai,gogogo")
	}else if ip != cip{
		fmt.Fprint(w,"can not login")
	}
	
	//fmt.Println(cip)
}

func sayBye(w http.ResponseWriter,r *http.Request){
	cip = ""
	fmt.Fprint(w,"bye")
}

func main(){
	http.HandleFunc("/hi",sayHello)
	http.HandleFunc("/bye",sayBye)
	// fmt.Println(cip)
	err := http.ListenAndServe(":9090",nil)
	
	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}

/*
func RemoteIp(req *http.Request) string {
    remoteAddr := req.RemoteAddr
    if ip := req.Header.Get(XRealIP); ip != "" {
        remoteAddr = ip
    } else if ip = req.Header.Get(XForwardedFor); ip != "" {
        remoteAddr = ip
    } else {
        remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
    }

    if remoteAddr == "::1" {
        remoteAddr = "127.0.0.1"
    }

    return remoteAddr
}
*/