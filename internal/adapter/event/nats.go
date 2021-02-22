package event

import (
	"context"
	"fmt"
	"time"

	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
	aevent "github.com/burungbangkai/fakesmtp/pkg/adapter/event"
	agrpc "github.com/burungbangkai/fakesmtp/pkg/adapter/grpc"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

const (
	ping_interval       = 10
	ping_max_out        = 5
	connect_wait_second = 5 * time.Second
)

func NewNATSConnection(url, clusterID, clientID string) (stan.Conn, port.StopGracefully, error) {
	natsc, err := nats.Connect(url)
	if err != nil {
		return nil, nil, err
	}
	nc, err := stan.Connect(clusterID, clientID,
		stan.NatsConn(natsc),
		stan.Pings(ping_interval, ping_max_out),
		stan.ConnectWait(connect_wait_second),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Printf("Connection lost, reason: %v", reason)
		}),
	)
	if err != nil {
		return nil, nil, err
	}
	return nc, func() {
		if err := nc.NatsConn().Drain(); err != nil {
			fmt.Println(err.Error())
		}
		if err := nc.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}, nil
}

func NewNATSRawEmailReceivedEventSender(sc stan.Conn, serFn port.Serialize) port.SendRawEmail {
	var ackFn stan.AckHandler = func(s string, e error) {
		if e != nil {
			fmt.Println(e.Error())
		}
	}
	return func(ctx context.Context, e *model.RawEmail) error {
		evt := &agrpc.RawEmailReceivedEvent{
			UserId:    e.UserID,
			InboxName: e.InboxName,
			Raw:       e.RawMsg,
		}
		byts, err := serFn(evt)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		_, err = sc.PublishAsync(aevent.NATS_RAW_EMAIL_RECEIVED_EVT_QUEUE, byts, ackFn)
		if err != nil {
			fmt.Println(err.Error())
		}
		return err
	}
}
