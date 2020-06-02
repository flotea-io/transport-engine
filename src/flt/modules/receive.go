/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package main

import (
	//"fmt"
	//"flt/channel"
	"os"
	"log"
)

func init() {
	/*
    orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", "postgres", "postgres://flotea:Fckgw-Rhqq2@127.0.0.1:5432/flotea_platform?sslmode=disable")
  */
}


func main() {

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [carier uuid]...", os.Args[0])
    os.Exit(0)
	}

	/*
  channel_s := channel.NewChannel()
	//channel_s2 := channel.NewChannel()

  //fmt.Printf("ChannelClient: %v\n", channel_s.Name())

  channel_s.CreateChannel("")
	//channel_s2.CreateChannel("")
  //fmt.Printf("ChannelClient: %v\n", channel_s.Channel)

	//channel_s.CreateExchange("test_exchange", "")
	//channel_s2.CreateExchange("test_exchange")

	//channel_s.SendMessage()
	channel_s.QueueDeclare("")
	//channel_s2.QueueDeclare("")
	channel_s.QueueBind("test_exchange", "black")
	channel_s2.QueueBind("test_exchange", "white")

  channel_s.ReceiveMessages()
	channel_s2.ReceiveMessages()
	*/
  //_ = channel_s

}
