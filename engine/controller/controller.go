package controller

import (
	"github.com/zserge/webview"
	"log"
	"net"
	"net/http"
	"perv/engine/providers"
)

type EngineController struct {
	model *providers.WindowModel
}

func (c *EngineController) Init() *EngineController {
	c.model = providers.NewWindowModel()
	return nil
}

func (c *EngineController) StartServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(c.model.IndexHTML()))
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}

func (c *EngineController) HandleRPC(w webview.WebView, data string){
	c.model.HandleRPC(&w, &data)
	return
}

