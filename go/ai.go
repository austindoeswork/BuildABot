package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	serverURL := "ws://npcompete.io/wsplay"
	devkey := "YOURDEVKEYHERE"

	//open websocket
	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial(serverURL, nil)
	checkErr(err)
	defer conn.Close()

	//write our devkey
	err = conn.WriteMessage(1, []byte(devkey))
	checkErr(err)

	//receive acknowledgement
	ginfo := &GameInfo{}
	_, msg, err := conn.ReadMessage()
	checkErr(err)

	json.Unmarshal(msg, ginfo)
	fmt.Println("I am player:", ginfo.Player)

	//send game inputs
	// frame := &Frame{}
	for {
		_, msg, err = conn.ReadMessage()
		_ = msg

		checkErr(err)
		// json.Unmarshal(msg, frame)
		// fmt.Println(frame.P1.Income)
		conn.WriteMessage(1, []byte("b00 02"))
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println("FATAL:", err)
		os.Exit(1)
	}
}

type GameInfo struct {
	Player   int    `json:"Player"`
	UserName string `json:"UserName"`
	GameName string `json:"GameName"`
}

type Frame struct {
	W  int `json:"w"`
	H  int `json:"h"`
	P1 struct {
		Owner  int      `json:"owner"`
		Income int      `json:"income"`
		Bits   int      `json:"bits"`
		Towers []string `json:"towers"`
		Troops []struct {
			Owner int `json:"owner"`
			X     int `json:"x"`
			Y     int `json:"y"`
			Maxhp int `json:"maxhp"`
			Hp    int `json:"hp"`
			Enum  int `json:"enum"`
		} `json:"troops"`
		MainTower struct {
			Owner int `json:"owner"`
			X     int `json:"x"`
			Y     int `json:"y"`
			Maxhp int `json:"maxhp"`
			Hp    int `json:"hp"`
			Enum  int `json:"enum"`
		} `json:"mainTower"`
	} `json:"p1"`
	P2 struct {
		Owner  int      `json:"owner"`
		Income int      `json:"income"`
		Bits   int      `json:"bits"`
		Towers []string `json:"towers"`
		Troops []struct {
			Owner int `json:"owner"`
			X     int `json:"x"`
			Y     int `json:"y"`
			Maxhp int `json:"maxhp"`
			Hp    int `json:"hp"`
			Enum  int `json:"enum"`
		} `json:"troops"`
		MainTower struct {
			Owner int `json:"owner"`
			X     int `json:"x"`
			Y     int `json:"y"`
			Maxhp int `json:"maxhp"`
			Hp    int `json:"hp"`
			Enum  int `json:"enum"`
		} `json:"mainTower"`
	} `json:"p2"`
}
