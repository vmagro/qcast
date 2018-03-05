package lib

import (
	"context"
	"errors"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"io"
	"net"

	"qcast"
)

type Server interface {
	ListenAndServe() error
}

type server struct {
	amqpAddr string
	updates  chan *qcast.QueueUpdate
	conn     *amqp.Connection
}

func NewServer(amqpAddr string) Server {
	return &server{
		amqpAddr: amqpAddr,
		updates:  make(chan *qcast.QueueUpdate),
	}
}

func (s *server) ListenAndServe() error {
	glog.Infof("Connecting to amqp %s", s.amqpAddr)

	conn, err := amqp.Dial(s.amqpAddr)
	if err != nil {
		glog.Errorf("Error connecting: %+v", err)
		return err
	}
	defer conn.Close()
	s.conn = conn

	// start the grpc server
	// TODO: get port from arg
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	glog.Infof("Listening on :8080")

	grpcServer := grpc.NewServer()
	qcast.RegisterDjServer(grpcServer, s)
	return grpcServer.Serve(lis)
}

func (s *server) QueueUpdates(empty *empty.Empty, stream qcast.Dj_QueueUpdatesServer) error {
	ctx := stream.Context()

	code, err := PartyCodeFromContext(ctx)
	if err != nil {
		glog.Error(err)
		return err
	}
	updates, err := QueueUpdates(code, s.conn)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			glog.Infof("Client disconnected from queue updates")
			return errors.New("Client disconnected")
		case update := <-updates:
			glog.Infof("Got update, forwarding to client")
			err := stream.Send(update)
			if err == io.EOF {
				return err
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *server) AddTrack(ctx context.Context, track *qcast.Track) (*qcast.QueuedTrack, error) {
	t := &qcast.QueuedTrack{
		Track: track,
	}
	update := &qcast.QueueUpdate{
		Track: t,
	}

	code, err := PartyCodeFromContext(ctx)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	err = SendToRabbitmq(code, update, s.conn)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return t, nil
}
