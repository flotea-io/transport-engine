package main

import (
	//_ "github.com/lib/pq"
	"fmt"
	//"os"
	//"reflect"
	//"flt/utils"
	"flt/channel"
	//"time"
	//"strconv"
	"github.com/tidwall/gjson"
)


func init() {}


  // 1. route per carrier
  // 2. send messages to carriers
  // 3. carrier decision to move/remove/accept
  // 4. ....

func main() {

  carriersJson := `{
  "carriers": [
      {
        "name": "Przewoznik I",
        "uuid": "fb3c0be3-3941-409e-8cfe-496eff10d651"
      },
      {
        "name": "Przewoznik II",
        "uuid": "8b3cff90-838c-490d-8dae-af4e653d5208"
      },
      {
        "name": "Przewoznik III",
        "uuid": "7875c845-0e13-4266-b50f-2f5da6bfaf1d"
      }
    ]
  }`

  carriers := gjson.Get(carriersJson, "carriers")
  data := carriers.Array()
  fmt.Printf("Len of Carriers: %v\n", len(data))

  for key, value := range data {
      _ = value
      carrier := data[key].Map();

      fmt.Printf("selector: %s\n", carrier["uuid"])
  }

  channel_s := channel.NewChannel()
  fmt.Printf("ChannelClient: %v\n", channel_s.Name())

  /*
  channel_s := channel.NewChannel()
  fmt.Printf("ChannelClient: %v\n", channel_s.Name())
  channel_s.CreateChannel("")

  //channel_s.Exclusive = true

  channel_s.CreateExchange("test_exchange", "direct")
  channel_s.SendMessage("test_exchange");

  //fmt.Printf("ChannelClient: %v\n", channel_s.Channel)


  // Exmple: go run main.go {pasazer: Testowy, count: 1, tel: 791 139 005, Cieszyn, {lat: 123, lng: 512}, Berlin {10,20}, {1,0,0,0}, 2020-01-01T15:00:00}
  // pasazer, ilosc, telefon, miasto, szerokosci, miasto, szerokosci, {pasazer, paczka, zwierze, niepelnosprawny}, data wyjazdu
  //channel_s.SendWorkMessage(os.Args)

  //channel_s.ReceiveMessages()
  */
  _ = channel_s

}
