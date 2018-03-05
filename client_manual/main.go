package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"time"

	"qcast"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	glog.Infof("Hello world!")

	// addr := "127.0.0.1:8080"
	addr := "192.168.64.2:31442"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	check(err)
	defer conn.Close()

	// client := qcast.NewSpotifyClient(conn)

	// token := "BQCiJxxaFkgbK638NoxKnWDSVu7bPw0l-JPNMqT-ANkEk2gOaNgWI5kE0bTtmmkT2RuljqnK1Gm-gumAISplm3_4aC9tlS8GJa_qSNoIY1miNErlKGhJdP25KN1Fii7NL8QzOTuwPHYlRo9H7A"

	// ctx := metadata.AppendToOutgoingContext(context.Background(), "spotify_token", token)

	// result, err := client.Search(ctx, &qcast.SearchRequest{"Kodaline"})
	// check(err)

	// for _, track := range result.Tracks {
	// 	fmt.Printf("%s\n\n", track)
	// }
	// for _, artist := range result.Artists {
	// 	fmt.Printf("%s\n\n", artist)
	// }
	// for _, album := range result.Albums {
	// 	fmt.Printf("%s\n\n", album)
	// }
	// fmt.Print("\n\n\n")
	// fmt.Printf("Got %d tracks, %d artists and %d albums\n", len(result.Tracks), len(result.Artists), len(result.Albums))

	ctx := metadata.AppendToOutgoingContext(context.Background(), "party_code", "ABCDE")
	client := qcast.NewDjClient(conn)
	wait := make(chan bool)
	go func() {
		stream, err := client.QueueUpdates(ctx, &empty.Empty{})
		glog.Infof("waiting for updates")
		check(err)
		for {
			update, err := stream.Recv()
			if err == io.EOF {
				break
			}
			check(err)
			fmt.Printf("%+v\n", update)
		}
		wait <- true
	}()

	// wait 2 seconds before sending
	<-time.After(2 * time.Second)

	track := &qcast.Track{
		Name: "Test track name",
	}
	client.AddTrack(ctx, track)
	glog.Infof("added track")

	<-wait
}
