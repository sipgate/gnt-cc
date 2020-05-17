package websocket

import (
	"gnt-cc/config"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func checkAndLogMessage(msgType int) {
	var typeStr string
	switch msgType {
	case websocket.BinaryMessage:
		typeStr = "binary"
	case websocket.TextMessage:
		typeStr = "text"
	case websocket.PingMessage:
		typeStr = "ping"
	case websocket.PongMessage:
		typeStr = "pong"
	case websocket.CloseMessage:
		typeStr = "close"
	}
	log.Debugf("Websocket: received message type '%s'", typeStr)
}

func Handler(w http.ResponseWriter, r *http.Request, host string, port int) error {
	log.Infoln("Upgrading connection to websocket")

	var upgrader = websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		Subprotocols:      []string{"binary"},
		EnableCompression: false,
	}
	if bool(config.Get().DevelopmentMode) {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Failed to set websocket upgrade: %s", err)
		return err
	}

	spiceSrv, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	log.Infof("Connecting to spice target %s:%d", host, port)
	if err != nil {
		log.Errorf("Failed to connect to spice target: %s", err)
		conn.Close()
		return err
	}

	go func() {
		defer spiceSrv.Close()
		defer conn.Close()
		counter := 0
		for {
			counter++
			log.Debugf("Spice2Websocket: Loop #%d", counter)
			buf := make([]byte, 1024)
			size, err := spiceSrv.Read(buf)
			if err != nil {
				log.Warningf("Spice2Websocket: failed to read from spice socket: %s", err)
				return
			}
			data := buf[:size]
			log.Debugf("Spice2Websocket: Writing %d Bytes to websocket", len(data))
			err = conn.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				log.Warningf("Spice2Websocket: failed to write to websocket: %s", err)
				return
			}
		}
	}()

	go func() {
		defer spiceSrv.Close()
		defer conn.Close()
		counter := 0
		for {
			counter++
			log.Debugf("Websocket2Spice: Loop #%d", counter)
			msgType, message, err := conn.ReadMessage()
			if _, ok := err.(*websocket.CloseError); ok {
				log.Debug("Websocket2Spice: closing websocket")
				return
			}
			if err != nil {
				log.Warningf("Websocket2Spice: failed to read websocket message: %s", err)
				return
			}
			checkAndLogMessage(msgType)
			log.Debugf("Websocket2Spice: Writing %d Bytes to spice socket", len(message))
			_, err = spiceSrv.Write(message)
			if err != nil {
				log.Warningf("Websocket2Spice: failed to write to spice socket: %s", err)
				return
			}
		}
	}()

	return nil
}
