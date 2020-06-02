/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package channel

import (
  "fmt"
  "os"
  "strings"
  "log"
  "github.com/streadway/amqp"
)

type ChannelClient struct {
  Conn *amqp.Connection
  Channel *amqp.Channel
  Queue amqp.Queue
  Durable bool
  Exclusive bool   // durable
  Del_when_used bool   // delete when unused
  No_wait bool   // no-wait
}



func init(){
}

// Constructs a new channelClient with the given framing rules
func NewChannel() *ChannelClient {

  conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
  if err != nil {
    log.Fatalf("%s: %s", "Failed to connect RabbitMQ: ", err)
  }

  //defer conn.Close()

	return &ChannelClient{
    Conn: conn,
    Durable: false,
    Exclusive: false,
    Del_when_used: false,
    No_wait: false,
  }
}

func (t *ChannelClient) Name() string {
    return "Channel RabbitMQ"
}

func (t *ChannelClient) failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

// create a channel, which is where most of the API
func (t *ChannelClient) CreateChannel(name string){

  if len(name) == 0 {
        name = ""
  }

  fmt.Printf("ChannelClient: %v\n", t.Conn.Channel)
  ch, err := t.Conn.Channel()
  t.failOnError(err, "Failed to open a channel")
  t.Channel = ch

  //defer ch.Close()
}

/**
* Declaring a queue is idempotent - it will only be created if it doesn't exist already.
* The message content is a byte array, so you can encode whatever you like there.
*/

func (t *ChannelClient) QueueDeclare(name string){

  if len(name) == 0 {
        name = ""
  }

  q, err := t.Channel.QueueDeclare(
    name, // name
    t.Durable,   // durable - survive broker restart
    t.Del_when_used,   // delete when unused - last consumer unsubscribe = delete queue
    t.Exclusive,   // exclusive - for one collection, connection closing = delete queue
    t.No_wait,   // no-wait
    nil,     // arguments, message and queue TTL (time to live), queue length, mirroring settings, max number of prior, consumer prior
  )
  t.failOnError(err, "Failed to declare a queue")
  t.Queue = q
}

func (t *ChannelClient) QueueBind(exchange string, key string){

    err := t.Channel.QueueBind(
      t.Queue.Name, // queue name
      key,     // routing key
      exchange, // exchange
      false,
      nil,
    )

    t.failOnError(err, "Failed to declare a queue")

}

// function thats send simple message to named channel (Directly)
func (t *ChannelClient) SendMessage(exchange string){

  if len(exchange) == 0 {
    exchange = ""
  }

  body := "Hello World!"

  err := t.Channel.Publish(
    exchange,     // exchange
    t.Queue.Name, // routing key
    false,  // mandatory
    false,  // immediate
    amqp.Publishing {
      ContentType: "text/plain",
      Body:        []byte(body),
    })
  t.failOnError(err, "Failed to publish a message")
}

// RECEIVER
// Setting up is the same as the publisher; we open a connection and a channel,
// and declare the queue from which we're going to consume.
// Note this matches up with the queue that send publishes to.

func (t *ChannelClient) ReceiveMessages(){
  msgs, err := t.Channel.Consume(
    t.Queue.Name, // queue
    "",     // consumer
    true,   // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // args
  )
  t.failOnError(err, "Failed to register a consumer")

  // What is chan? :)
  forever := make(chan bool)

      go func() {
        for d := range msgs {
          log.Printf("Received a message: %s", d.Body)
        }
      }()

      log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}

func (t *ChannelClient) SendWorkMessage(msg []string){

  body := t.bodyFrom(msg)
  //body := msg

  err := t.Channel.Publish(
    "",           // exchange
    t.Queue.Name,       // routing key
    false,        // mandatory
    false,
    amqp.Publishing {
      DeliveryMode: amqp.Persistent,
      ContentType:  "text/plain",
      Body:         []byte(body),
    })

    t.failOnError(err, "Failed to publish a message")
    log.Printf(" [x] Sent %s", body)

}

/**
* Available types: direct, topic, headers, fanout
* fanout - simple broadcasts all messages it receives to all the queues (good for logger)
* direct - simple messaing to queues that binding key match routing_key of message
*/
func (t *ChannelClient) CreateExchange(name string, typex string) {

  err := t.Channel.ExchangeDeclare(
      name,   // name
      typex, // type
      true,     // durable
      false,    // auto-deleted
      false,    // internal
      false,    // no-wait
      nil,      // arguments
    )

  t.failOnError(err, "Failed to publish a message")
}

func (t *ChannelClient) bodyFrom(args []string) string {
        var s string
        if (len(args) < 2) || os.Args[1] == "" {
                s = "hello"
        } else {
                s = strings.Join(args[1:], " ")
        }
        return s
}
