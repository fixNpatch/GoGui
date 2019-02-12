package providers

import (
	"github.com/zserge/webview"
	"log"
	"strconv"
	"strings"
)

type WindowModel struct {}

func NewWindowModel() *WindowModel {
	return &WindowModel{}
}

func (m *WindowModel) IndexHTML() string  {
	return 	`
<!doctype html>
<html>
	<head>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
	</head>
	<body>
		<button onclick="external.invoke('close')">Close</button>
		<button onclick="external.invoke('fullscreen')">Fullscreen</button>
		<button onclick="external.invoke('unfullscreen')">Unfullscreen</button>
		<button onclick="external.invoke('open')">Open</button>
		<button onclick="external.invoke('opendir')">Open directory</button>
		<button onclick="external.invoke('save')">Save</button>
		<button onclick="external.invoke('message')">Message</button>
		<button onclick="external.invoke('info')">Info</button>
		<button onclick="external.invoke('warning')">Warning</button>
		<button onclick="external.invoke('error')">Error</button>
		<button onclick="external.invoke('changeTitle:'+document.getElementById('new-title').value)">
			Change title
		</button>
		<input id="new-title" type="text" />
		<button onclick="external.invoke('changeColor:'+document.getElementById('new-color').value)">
			Change color
		</button>
		<input id="new-color" value="#e91e63" type="color" />
	</body>
</html>
`
}


func (m *WindowModel) HandleRPC(w *webview.WebView, data *string){
	wb := *w
	dt := *data
	switch {
	case dt == "close":
		wb.Terminate()
	case dt == "fullscreen":
		wb.SetFullscreen(true)
	case dt == "unfullscreen":
		wb.SetFullscreen(false)
	case dt == "open":
		log.Println("open", wb.Dialog(webview.DialogTypeOpen, 0, "Open file", ""))
	case dt == "opendir":
		log.Println("open", wb.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", ""))
	case dt == "save":
		log.Println("save", wb.Dialog(webview.DialogTypeSave, 0, "Save file", ""))
	case dt == "message":
		wb.Dialog(webview.DialogTypeAlert, 0, "Hello", "Hello, world!")
	case dt == "info":
		wb.Dialog(webview.DialogTypeAlert, webview.DialogFlagInfo, "Hello", "Hello, info!")
	case dt == "warning":
		wb.Dialog(webview.DialogTypeAlert, webview.DialogFlagWarning, "Hello", "Hello, warning!")
	case dt == "error":
		wb.Dialog(webview.DialogTypeAlert, webview.DialogFlagError, "Hello", "Hello, error!")
	case strings.HasPrefix(dt, "changeTitle:"):
		wb.SetTitle(strings.TrimPrefix(dt, "changeTitle:"))
	case strings.HasPrefix(dt, "changeColor:"):
		hex := strings.TrimPrefix(strings.TrimPrefix(dt, "changeColor:"), "#")
		num := len(hex) / 2
		if !(num == 3 || num == 4) {
			log.Println("Color must be RRGGBB or RRGGBBAA")
			return
		}
		i, err := strconv.ParseUint(hex, 16, 64)
		if err != nil {
			log.Println(err)
			return
		}
		if num == 3 {
			r := uint8((i >> 16) & 0xFF)
			g := uint8((i >> 8) & 0xFF)
			b := uint8(i & 0xFF)
			wb.SetColor(r, g, b, 255)
			return
		}
		if num == 4 {
			r := uint8((i >> 24) & 0xFF)
			g := uint8((i >> 16) & 0xFF)
			b := uint8((i >> 8) & 0xFF)
			a := uint8(i & 0xFF)
			wb.SetColor(r, g, b, a)
			return
		}
	}
}
