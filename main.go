// Private channel subscription example.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/centrifugal/centrifuge-go"
	"github.com/galt-tr/tx-filter/pkg/common"
	"github.com/galt-tr/tx-filter/pkg/socket"
	//	"github.com/golang-jwt/jwt"
)

/*func connToken(user string, exp int64) string {
	// NOTE that JWT must be generated on backend side of your application!
	// Here we are generating it on client side only for example simplicity.
	claims := jwt.MapClaims{"sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return t
}

func subscribeToken(channel string, client string, exp int64) string {
	// NOTE that JWT must be generated on backend side of your application!
	// Here we are generating it on client side only for example simplicity.
	claims := jwt.MapClaims{"channel": channel, "client": client}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return t
}*/

type eventHandler struct {
	Processor *socket.Processor
}

/*func (h *eventHandler) OnPrivateSub(_ *centrifuge.Client, e centrifuge.PrivateSubEvent) (string, error) {
	token := subscribeToken(e.Channel, e.ClientID, time.Now().Unix()+10)
	return token, nil
}*/

func (h *eventHandler) OnConnect(_ *centrifuge.Client, _ centrifuge.ConnectEvent) {
	log.Println("Connected")
}

func (h *eventHandler) OnError(_ *centrifuge.Client, e centrifuge.ErrorEvent) {
	log.Println("Error", e.Message)
}

func (h *eventHandler) OnDisconnect(_ *centrifuge.Client, e centrifuge.DisconnectEvent) {
	log.Println("Disconnected", e.Reason)
}

func (h *eventHandler) OnMessage(_ *centrifuge.Client, e centrifuge.MessageEvent) {
	log.Println(fmt.Sprintf("New message received from channel: %s", string(e.Data)))
}
func (h *eventHandler) OnServerPublish(_ *centrifuge.Client, e centrifuge.ServerPublishEvent) {
	tx, err := h.Processor.FilterMempoolPublishEvent(e)
	if err != nil {
		log.Printf("error processing event: %v", err)
	}
	if tx == nil {
		log.Printf("filtered transaction")
		return
	}
	log.Printf("filter accepted transaction: %#v", tx.GetTxID())
}

/*type subEventHandler struct{}

func (h *subEventHandler) OnSubscribeSuccess(sub *centrifuge.Subscription, _ centrifuge.SubscribeSuccessEvent) {
	log.Println(fmt.Sprintf("Successfully subscribed to private channel %s", sub.Channel()))
}

func (h *subEventHandler) OnSubscribeError(sub *centrifuge.Subscription, e centrifuge.SubscribeErrorEvent) {
	log.Println(fmt.Sprintf("Error subscribing to private channel %s: %v", sub.Channel(), e.Error))
}

func (h *subEventHandler) OnUnsubscribe(sub *centrifuge.Subscription, _ centrifuge.UnsubscribeEvent) {
	log.Println(fmt.Sprintf("Unsubscribed from private channel %s", sub.Channel()))
}

func (h *subEventHandler) OnPublish(sub *centrifuge.Subscription, e centrifuge.PublishEvent) {
	log.Println(fmt.Sprintf("New message received from channel %s: %s", sub.Channel(), string(e.Data)))
}*/

func newClient(txType common.TransactionType) *centrifuge.Client {
	wsURL := "wss://socket.whatsonchain.com/mempool"
	c := centrifuge.NewJsonClient(wsURL, centrifuge.DefaultConfig())

	//c.SetToken(connToken("112", 0))

	handler := &eventHandler{
		Processor: socket.NewProcessor(txType),
	}
	//c.OnPrivateSub(handler)
	c.OnDisconnect(handler)
	c.OnConnect(handler)
	//c.OnMessage(handler)
	c.OnServerPublish(handler)
	c.OnError(handler)
	return c
}

func main() {
	filterFlag := flag.String("filter", "metanet", "transaction filter type")
	flag.Parse()

	txType := common.TransactionType(*filterFlag)
	c := newClient(txType)
	defer func() { _ = c.Close() }()

	log.Printf("Starting TX filter service for transaction type [%v]", txType)
	err := c.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	/*sub, err := c.NewSubscription("mempool")
	if err != nil {
		log.Fatalln(err)
	}

	subEventHandler := &subEventHandler{}
	sub.OnSubscribeSuccess(subEventHandler)
	sub.OnSubscribeError(subEventHandler)
	sub.OnUnsubscribe(subEventHandler)
	sub.OnPublish(subEventHandler)

	// Subscribe on private channel.
	_ = sub.Subscribe()*/

	select {}
}
