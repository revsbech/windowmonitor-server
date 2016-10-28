package main

import (
	"flag"
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	"math/rand"
	"net/http"
	"time"


	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{} // use default options
var channels = make(map[string]*Channel)

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	rand.Seed(time.Now().UnixNano())

	r := mux.NewRouter()

	fmt.Printf("Starting server on %s\n", *addr)

	channels["demochannel"] = NewChannel("demochannel", "MyTestChannel")

	r.HandleFunc("/echo", echo)
	r.HandleFunc("/", home)
	r.HandleFunc("/channel/{channelId}/listen", getListenHandler())
	r.HandleFunc("/stats", statsHandler)
	http.Handle("/", r)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Unable to start server %v", err)
	}
	log.Fatal(http.ListenAndServe(*addr, nil))
}


func statsHandler(w http.ResponseWriter, r *http.Request) {

	clients := channels["demochannel"].GetClients()
	cvs := make([]ClientView, len(clients))

	for i, cl := range clients {
		cvs[i] = ClientView{IpAddress: cl.Ip}
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cvs); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

}


func getListenHandler() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		channelId := vars["channelId"]
		var channel *Channel
		var ok bool

		if channel, ok = channels[channelId]; !ok {
			log.Printf("Unable to find channel with id %s\n", channelId)
			http.Error(w, "Unable to find channel", http.StatusNotFound)
			return
		}
		log.Printf("Attaching client to channel %s", channel.name)

		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		channel.RegisterClient(&Client{
			conn: c,
			Ip:   r.RemoteAddr,
		})


		for {
			//Loop forever, or untill connection clsoses...

		}
		fmt.Println("Closing connection")
	}
}


// helper function
func generateEvents(c chan DataPoint) {
	for {
		value := rand.Float64()
		p := DataPoint{
			Timestamp: time.Now(),
			Value: value,
		}
		//fmt.Printf("Event. Random: %f\n", value);
		c <- p
		time.Sleep(1 * time.Second)
	}
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server,
"Send" to send a message to the server and "Close" to close the connection.
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
