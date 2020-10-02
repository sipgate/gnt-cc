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

func PassThrough(w http.ResponseWriter, r *http.Request, host string, port int) error {
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

	remoteSrv, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	log.Infof("Connecting to remote target %s:%d", host, port)
	if err != nil {
		log.Errorf("Failed to connect to remote target: %s", err)
		conn.Close()
		return err
	}

	go func() {
		defer remoteSrv.Close()
		defer conn.Close()
		counter := 0
		for {
			counter++
			log.Debugf("node->gnt-cc: Loop #%d", counter)
			buf := make([]byte, 1024)
			size, err := remoteSrv.Read(buf)
			if err != nil {
				log.Warningf("node->gnt-cc: failed to read from remote socket: %s", err)
				return
			}
			data := buf[:size]
			log.Debugf("node->gnt-cc: Writing %d Bytes to websocket", len(data))
			err = conn.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				log.Warningf("node->gnt-cc: failed to write to websocket: %s", err)
				return
			}
		}
	}()

	go func() {
		defer remoteSrv.Close()
		defer conn.Close()
		counter := 0
		for {
			counter++
			log.Debugf("gnt-cc->node: Loop #%d", counter)
			msgType, message, err := conn.ReadMessage()
			if _, ok := err.(*websocket.CloseError); ok {
				log.Debug("gnt-cc->node: closing websocket")
				return
			}
			if err != nil {
				log.Warningf("gnt-cc->node: failed to read websocket message: %s", err)
				return
			}
			checkAndLogMessage(msgType)
			log.Debugf("gnt-cc->node: Writing %d Bytes to remote socket", len(message))
			_, err = remoteSrv.Write(message)
			if err != nil {
				log.Warningf("gnt-cc->node: failed to write to remote socket: %s", err)
				return
			}
		}
	}()

	return nil
}
