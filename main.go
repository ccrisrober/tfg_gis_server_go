// Copyright (c) 2015, maldicion069 (Cristian Rodríguez) <ccrisrober@gmail.con>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.package com.example

package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/jsonq"
	"log"
	"math/rand"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8089"
	CONN_TYPE = "tcp"
)

func RandomValue(min, max int) int {
	return rand.Intn(max-min) + min
}

var maps []TMap
var keys map[string]TKeyObject
var positions map[int]TObjectUser

//
// Eventos servidor al cliente
//
type TServerSendMap struct {
	Action string //  `json:"Action"`
	Id     int
	Map    TMap
	X      float64
	Y      float64
	//Users map[int]TObjectUser
}

// PROBLEMS: GO have problems with maps marshalling
func SendMap(num_map int, id int) string {
	size := len(positions)
	log.Printf("Hay %d usuarios", size)

	ret := &TNewConnection{Action: "sendMap", Map: maps[num_map], X: 5 * 64, Y: 564, Id: id}

	b, _ := json.Marshal(ret)
	msg := string(b)
	log.Println("FINISH", msg)
	return msg
}

type TMsgWithId struct {
	message string
	id      int
}

func main() {

	runtime.GOMAXPROCS(2)
	rand.Seed(time.Now().Unix()) // Set the rand seed

	keys = make(map[string]TKeyObject)
	positions = make(map[int]TObjectUser)

	keys["Red"] = CreateKeyObject(1, 5*64, 5*64, "Red")
	keys["Blue"] = CreateKeyObject(2, 6*64, 5*64, "Blue")
	keys["Yellow"] = CreateKeyObject(3, 7*64, 5*64, "Yellow")
	keys["Green"] = CreateKeyObject(4, 8*64, 5*64, "Green")

	maps = make([]TMap, 1)

	maps[0] = CreateTMap(1, // ID
		"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 4, 0, 0, 0, 1,"+
			"1, 1, 0, 0, 0, 0, 0, 0, 0, 5, 5, 7, 5, 5, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 5, 5, 5, 5, 5, 5, 5, 1,"+
			"1, 1, 0, 0, 4, 6, 5, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 4, 0, 5, 5, 5, 0, 8, 8, 8, 0, 0, 1,"+
			"1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 4, 0, 0, 0, 1, 1, 0, 5, 5, 5, 0, 0, 0, 8, 8, 8, 4, 0, 1,"+
			"1, 0, 1, 0, 0, 0, 0, 4, 0, 0, 1, 1, 1, 1, 4, 0, 0, 0, 1, 0, 5, 0, 4, 4, 0, 0, 8, 8, 8, 1, 4, 1,"+
			"4, 0, 1, 0, 0, 0, 0, 4, 4, 0, 1, 1, 1, 1, 1, 1, 4, 0, 1, 0, 5, 0, 4, 4, 4, 0, 8, 8, 8, 1, 1, 1,"+
			"1, 0, 1, 0, 0, 4, 4, 4, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 5, 4, 4, 4, 0, 0, 0, 0, 1, 1, 1, 1,"+
			"1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 4, 0, 0, 0, 1,"+
			"1, 1, 0, 0, 0, 0, 0, 0, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 5, 5, 5, 5, 5, 5, 5, 5,"+
			"1, 1, 0, 0, 4, 0, 5, 5, 5, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 4, 0, 5, 5, 5, 0, 1, 1, 1, 0, 0, 1,"+
			"1, 1, 1, 0, 5, 5, 5, 0, 0, 0, 1, 1, 1, 4, 0, 0, 0, 1, 1, 0, 5, 5, 5, 0, 0, 0, 1, 1, 1, 4, 0, 1,"+
			"1, 0, 1, 0, 5, 0, 4, 4, 0, 0, 1, 1, 1, 1, 4, 0, 0, 0, 1, 0, 5, 0, 4, 4, 0, 0, 1, 1, 1, 1, 4, 1,"+
			"4, 0, 1, 0, 5, 0, 4, 4, 4, 0, 1, 1, 1, 1, 1, 1, 4, 0, 1, 0, 5, 0, 4, 4, 4, 0, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 4, 4, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 4, 4, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"+
			"1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,",
		32, // WIDTH
		25, // HEIGHT
		[]TKeyObject{keys["Red"], keys["Blue"], keys["Green"], keys["Yellow"]})

	// Number of people whom ever connected
	//
	clientCount := 0

	// All people who are connected; a map wherein
	// the keys are net.Conn objects and the values
	// are client "ids", an integer.
	//
	allClients := make(map[net.Conn]int)

	// Channel into which the TCP server will push
	// new connections.
	//
	newConnections := make(chan net.Conn)

	// Channel into which we'll push dead connections
	// for removal from allClients.
	//
	deadConnections := make(chan net.Conn)

	// Channel into which we'll push messages from
	// connected clients so that we can broadcast them
	// to every connection in allClients.
	//
	messages := make(chan TMsgWithId)
	battles := make(chan string)

	// Start the TCP server
	//
	server, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Iniciamos servidor xD %s", server.Addr().String())

	// Tell the server to accept connections forever
	// and push new connections into the newConnections channel.
	//
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			newConnections <- conn
		}
	}()

	for {

		select {

		// Accept new clients
		//
		case conn := <-newConnections:

			log.Printf("Accepted new client, #%d", clientCount)

			// Add this connection to the `allClients` map
			//
			allClients[conn] = clientCount

			player := CreateObjectUser(clientCount, 320, 320)

			positions[clientCount] = player

			msg := SendMap(0, clientCount)

			//auxCount := clientCount

			clientCount += 1

			_, err := conn.Write([]byte(msg))

			log.Printf("Conectados %d usuarios", len(positions))
			// Si hay un error en la comunicación, enviamos un mensaje para desconectar el socket
			if err != nil {
				deadConnections <- conn
			}

			// Constantly read incoming messages from this
			// client in a goroutine and push those onto
			// the messages channel for broadcast to others.
			//
			go func(conn net.Conn, clientId int) {
				buf := make([]byte, 512)
				init := false
				for {
					msg := TMsgWithId{id: clientId, message: ""}
					if !init {
						init = true
						newMsg := &TServerNewConnection{Action: "new", Id: player.GetId(), PosX: player.GetPosX(), PosY: player.GetPosY(), Map: player.GetMap()}
						b, _ := json.Marshal(newMsg)
						msg.message = string(b)
						log.Println("Iniciado")
					} else {
						_, err := conn.Read(buf)
						if err != nil {
							log.Println("Error reading:", err.Error())
							break
						}

						message := fmt.Sprintf("%s", buf)

						//log.Printf("Mensaje %s para cliente %d", message, clientId)

						data := map[string]interface{}{}
						dec := json.NewDecoder(strings.NewReader(message))
						_ = dec.Decode(&data)
						jq := jsonq.NewQuery(data)
						action, _ := jq.String("Action")

						action = string(action)

						log.Print("Accion => ", action)

						msg.message = message

						exitSocket := false
						//onlyToOneSocket := false
						//callBackSocket := true

						/*if action == "exit" {
							//msg = fmt.Sprintf("%d||%s", clientId, buf)
							msg = TMsgWithId{id: clientId, message: message}
							exitSocket = true
						}*/
						//log.Printf("Broadcast: (%d, %s)", msg.id, msg.message)

						messages <- msg

						if exitSocket {
							break
						}
						/*log.Println("  Mensaje  => ", msg.message)


						if action == "new" {
							//msg = fmt.Sprintf("%d||%s", clientId, buf)
							msg = TMsgWithId{id: clientId, message: message}
						} else if action == "move" {
							pos_x, _ := jq.Float("PosX")
							pos_y, _ := jq.Float("PosY")
							positions[clientId].SetPosX(pos_x)
							positions[clientId].SetPosY(pos_y)
						} else if action == "position" {
							position := positions[clientId]
							ret := &TServerSendPosition{Action: "position", Id: position.GetId(), RollDie: position.GetRollDie(), PosX: position.GetPosX(), PosY: position.GetPosY(), Map: position.GetMap()}
							b, _ := json.Marshal(ret)
							msg := string(b)
							_, err := conn.Write([]byte(msg))
							// Si hay un error en la comunicación, enviamos un mensaje para desconectar el socket
							if err != nil {
								deadConnections <- conn
							}
							callBackSocket = false
						} else if action == "fight" {
							//msg = fmt.Sprintf("%d||%s", clientId, buf)
							msg = TMsgWithId{id: clientId, message: message}
						} else if action == "finishBattle" {

						} else if action == "exit" {
							//msg = fmt.Sprintf("%d||%s", clientId, buf)
							msg = TMsgWithId{id: clientId, message: message}
							exitSocket = true
						}

						log.Printf("Broadcast: (%d, %s)", msg.id, msg.message)

						if callBackSocket {
							if onlyToOneSocket {
								//battles <- msg
							} else {
								messages <- msg
								if exitSocket {
									break
								}
							}
						}*/
					}

					messages <- msg
				}

				// When we encouter `err` reading, send this
				// connection to `deadConnections` for removal.
				//
				deadConnections <- conn

			}(conn, allClients[conn])

		// Accept battle messages from connected clients
		//
		case fight_msg := <-battles:
			log.Println(fight_msg)

		// Accept messages from connected clients
		//
		case message := <-messages:
			//log.Printf("----" + message.message)
			identf := message.id
			msg := message.message

			// Loop over all connected clients
			//
			for conn, idf := range allClients {
				if identf != idf {
					// Send them a message in a go-routine
					// so that the network operation doesn't block
					//
					go func(conn net.Conn, message string) {
						_, err := conn.Write([]byte(message))

						// If there was an error communicating
						// with them, the connection is dead.
						if err != nil {
							deadConnections <- conn
						}
					}(conn, msg)
				}
			}
			//log.Printf("New message: %s", msg)
			//log.Printf("Broadcast to %d clients", len(allClients))

		// Remove dead clients
		//
		case conn := <-deadConnections:
			log.Printf("Client %d disconnected", allClients[conn])
			conn.Close()
			delete(allClients, conn)
			// TODO : Falta eliminar el objeto de positions
		}
	}
}
