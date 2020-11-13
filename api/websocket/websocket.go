package websocket

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gnt-cc/config"
	"io"
	"net"
	"net/http"
	"nhooyr.io/websocket"
)

/*func checkAndLogMessage(msgType int) {
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
}*/

func copyData(dst io.Writer, src io.Reader, doneChan chan<- bool) {
	_, err := io.Copy(dst, src)

	if err != nil {
		panic(err)
	}

	doneChan <- true
}

func PassThrough(w http.ResponseWriter, r *http.Request, host string, port int) error {
	log.Infoln("Upgrading connection to websocket")

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:         []string{"binary"},
		OriginPatterns:       []string{"*"},

	})

	if err != nil {
		log.Errorf("Failed to set websocket upgrade: %s", err)
		return err
	}

	if bool(config.Get().DevelopmentMode) {

	}

	log.Infof("Connecting to remote target %s:%d", host, port)
	ip, err := net.LookupIP(host)

	if err != nil {
		log.Errorf("Cannot look up host")
		return nil
	}

	remoteSrv, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   ip[0],
		Port: port,
	})

	if err != nil {
		log.Errorf("Failed to connect to remote target: %s", err)
		conn.Close(websocket.StatusBadGateway, "cannot reach target")
		return nil
	}

	doneChan := make(chan bool)
	//ctx := context.

	_, reader, err := conn.Reader(context.TODO())

	if err != nil {
		log.Errorf("Cannot create websocket reader", err)
		return nil
	}

	writer, err := conn.Writer(context.TODO(), websocket.MessageBinary)

	if err != nil {
		log.Errorf("Cannot create websocket writer", err)
		return nil
	}

	go copyData(remoteSrv, reader, doneChan)
	go copyData(writer, remoteSrv, doneChan)

	<-doneChan
	conn.CloseRead(context.TODO())
	remoteSrv.CloseRead()
	<-doneChan


	return nil



	//remoteSrv, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), 3 * time.Second)
	if err != nil {
		log.Errorf("Failed to connect to remote target: %s", err)
		conn.Close(websocket.StatusBadGateway, "cannot reach target")
		return nil
	}

	go func() {
		defer remoteSrv.CloseRead()
		defer conn.CloseRead(context.TODO())
		counter := 0
		for {
			counter++
			log.Debugf("node->gnt-cc: Loop #%d", counter)
			buf := make([]byte, 1024)
			size, err := remoteSrv.Read(buf)

			if err != nil {
				if err.Error() == "EOF" {
					return
				}

				log.Warningf("node->gnt-cc: failed to read from remote socket: %s", err)
				return
			}
			data := buf[:size]
			log.Debugf("node->gnt-cc: Writing %d Bytes to websocket", len(data))
			err = conn.Write(context.TODO(), websocket.MessageBinary, data)
			if err != nil {
				log.Warningf("node->gnt-cc: failed to write to websocket: %s", err)
				return
			}
		}
	}()

	go func() {
		defer remoteSrv.CloseRead()
		defer conn.CloseRead(context.TODO())
		counter := 0
		for {
			counter++
			log.Debugf("gnt-cc->node: Loop #%d", counter)
			_, message, err := conn.Read(context.TODO())
			if _, ok := err.(*websocket.CloseError); ok {
				log.Debug("gnt-cc->node: closing websocket")
				return
			}
			if err != nil {
				log.Warningf("gnt-cc->node: failed to read websocket message: %s", err)
				return
			}
			//checkAndLogMessage(msgType)
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
