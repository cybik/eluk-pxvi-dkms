package main

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"html"
	"log"
	"net/http"
)

// TODO: kill using PID via systemd
// TODO: constize sample part of file names
// TODO: access sys files (read, write)
// TODO: write config to /var
// TODO: check /var for config
// TODO: create default configs in /var if none exist
// TODO: read configs from /var if they exist
// TODO: write loaded configs to /sys files on start if they exist
func main() {
	const daemon_name string = "eluk_pxvi_led_daemon"
	cntxt := &daemon.Context{
		PidFileName: fmt.Sprintf("%s.pid", daemon_name),
		PidFilePerm: 0644,
		LogFileName: fmt.Sprintf("%s.log", daemon_name),
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{fmt.Sprintf("[%s]", daemon_name)},
	}
	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	serveHTTP()
}

func serveHTTP() {
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, "go-daemon: %q", html.EscapeString(r.URL.Path))
}
