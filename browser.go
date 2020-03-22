package mplayer

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os/exec"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// BrowserPlayer : Plays a movie in a HTML web page
type BrowserPlayer struct {
	Props
}

const tpl = `
<!DOCTYPE HTML>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{.Title}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="Live Gophie server" />
    <meta name="author" content="Bisoncorps" />
    <style type="text/css">
        #content {
					margin: 0;
					padding: 0;
					height: 100vh;
					left: 0;
					right: 0;
					bottom: 0;
					top: 0px;
        }
    </style>
</head>

<body>
		<video id="content" controls autoplay>
      <source src="{{ .StreamURL}}">
      Your browser does not support the video tag.
    </video> 
</body>
</html>
`

func (p *BrowserPlayer) Play() {
	tmpl, err := template.New("streamingPage").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, &p)
	})

	// Create a listener. This allows for dynamic port allocation
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	url := fmt.Sprintf("http://localhost:%d", port)

	log.Info("Opening Browser at ", url)
	// Launch Browser
	go open(url)
	// Start Server
	log.Error(http.Serve(listener, nil))
}

// Opens URL in browser
func open(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
