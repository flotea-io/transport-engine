/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package controllers

import (
    "github.com/astaxie/beego"
    "github.com/gorilla/websocket"
    "flt/models"
    "fmt"
    "net/http"
    //"github.com/astaxie/beego/toolbox"
)

type MyWebSocketController struct {
    beego.Controller
}

var upgrader = websocket.Upgrader{}

func (c *MyWebSocketController) Get() {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
    if err != nil {
        fmt.Println(err);
    }
    //  defer ws.Close()
    var data = models.WsData()
    data.Clients[ws] = true
    c.TplName = "empty.tpl"
}
