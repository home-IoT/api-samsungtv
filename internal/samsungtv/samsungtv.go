package samsungtv

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/home-IoT/api-samsungtv/internal/log"

	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

const keyPrefix = "KEY_"
const statusURL = "http://%s:8001/api/v2/"

// CheckConnection checks if a connection to the TV is possible
func CheckConnection() (bool, string) {
	var text string
	conn, err := connect()
	if err == nil {
		defer closeConnection(conn)
		text, err = getStatus()
	}

	return (err == nil), text
}

// SendKey sends a key to the TV
func SendKey(key string) error {
	conn, err := connect()
	if err != nil {
		log.Debugf("Error establishing a connection.")
		return err
	}
	defer closeConnection(conn)

	keyCommand := getKeyCommand(key)
	wErr := conn.WriteJSON(keyCommand)
	if wErr != nil {
		log.Debugf("Error sending key: %v", wErr)
		return wErr
	}
	jsonStr, jErr := json.Marshal(keyCommand)
	if jErr == nil {
		log.Infof("Sent '%s'.", jsonStr)
	}

	return nil
}

// connect opens a Websocket connection to the TV
func connect() (*websocket.Conn, error) {
	host := fmt.Sprintf("%s:%s", configuration.TV.Host, *configuration.TV.Port)
	path := "/api/v2/channels/samsung.remote.control"
	query := fmt.Sprintf("name=%s", base64.StdEncoding.EncodeToString([]byte(configuration.Controller.Name)))
	u := url.URL{Scheme: *configuration.TV.Protocol, Host: host, Path: path, RawQuery: query}

	log.Infof("Opening connection to %s ...", u.String())

	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	connection, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Debugf("%v", err)
		return nil, err
	}

	log.Infof("Connection is established.")

	return connection, nil
}

// closeConnection closes the Websocket if it is open
func closeConnection(connection *websocket.Conn) {
	defer connection.Close()

	log.Infof("Closing the connection...")

	err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Infof("Error closing connection: %v", err)
		return
	}
	time.Sleep(time.Second / 2)

	connection = nil
	log.Infof("Connection closed.")
}

func getStatus() (string, error) {
	resp, err := http.Get(fmt.Sprintf(statusURL, configuration.TV.Host))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var text []byte
	text, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(text), nil
}

type keyCommandParams struct {
	Cmd          string `json:"Cmd"`
	DataOfCmd    string `json:"DataOfCmd"`
	Option       string `json:"Option"`
	TypeOfRemote string `json:"TypeOfRemote"`
}

type keyCommand struct {
	Method string           `json:"method"`
	Params keyCommandParams `json:"params"`
}

var keyTemplate *keyCommand

func getKeyCommand(key string) keyCommand {
	if !strings.HasPrefix(key, keyPrefix) {
		key = keyPrefix + key
	}

	key = strings.ToUpper(key)

	if keyTemplate == nil {
		keyTemplate = &keyCommand{
			Method: "ms.remote.control",
			Params: keyCommandParams{
				Cmd:          "Click",
				DataOfCmd:    key,
				Option:       "false",
				TypeOfRemote: "SendRemoteKey",
			},
		}
	}

	keyTemplate.Params.DataOfCmd = key

	return *keyTemplate
}
