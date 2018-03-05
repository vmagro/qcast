package lib

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"

	"qcast"
)

type Server interface {
	ListenAndServe() error
}

type server struct {
}

func NewServer() Server {
	return &server{}
}

func (s *server) ListenAndServe() error {
	// TODO: get port from arg
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	glog.Infof("Listening on :8080")

	grpcServer := grpc.NewServer()
	qcast.RegisterSpotifyServer(grpcServer, s)
	return grpcServer.Serve(lis)
}

// GetSpotifyClient gets an authenticated Spotify client from the request context metadata
func GetSpotifyClient(ctx context.Context) (*spotify.Client, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Error getting request metadata")
	}

	tokenString := md["spotify_token"][0]

	auth := spotify.NewAuthenticator("")

	token := &oauth2.Token{
		AccessToken: tokenString,
	}

	client := auth.NewClient(token)
	return &client, nil
}

func ConvertTrack(spt *spotify.FullTrack) *qcast.Track {
	track := qcast.Track{
		Uri:  string(spt.URI),
		Name: spt.Name,
	}
	return &track
}

func ConvertImage(spt *spotify.Image) *qcast.Image {
	image := qcast.Image{
		Url:    spt.URL,
		Width:  int32(spt.Width),
		Height: int32(spt.Height),
	}
	return &image
}

func ConvertImages(sptImages []spotify.Image) []*qcast.Image {
	images := make([]*qcast.Image, len(sptImages))
	for idx, spt := range sptImages {
		images[idx] = ConvertImage(&spt)
	}
	return images
}

func ConvertArtist(spt spotify.SimpleArtist) *qcast.Artist {
	// only the FullArtist has images included, so otherwise use an empty array
	images := []*qcast.Image{}
	// if full, ok := spt.(spotify.FullArtist); ok {
	// 	images = ConvertImages(full.Images)
	// }
	artist := qcast.Artist{
		Uri:    string(spt.URI),
		Name:   spt.Name,
		Images: images,
	}
	return &artist
}

func ConvertAlbum(spt spotify.SimpleAlbum) *qcast.Album {
	artists := make([]*qcast.Artist, len(spt.Artists))
	for idx, spt := range spt.Artists {
		artists[idx] = ConvertArtist(spt)
	}
	album := qcast.Album{
		Uri:     string(spt.URI),
		Name:    spt.Name,
		Artists: artists,
		Art:     ConvertImages(spt.Images),
	}
	return &album
}

func (s *server) Search(ctx context.Context, request *qcast.SearchRequest) (*qcast.SearchResponse, error) {
	client, err := GetSpotifyClient(ctx)
	if err != nil {
		return nil, err
	}

	spotifyResults, err := client.Search(request.Query, spotify.SearchTypeAlbum|spotify.SearchTypeArtist|spotify.SearchTypeTrack)
	if err != nil {
		return nil, err
	}

	// TODO: support pagination (and deduplication?)

	tracks := make([]*qcast.Track, len(spotifyResults.Tracks.Tracks))
	for idx, spt := range spotifyResults.Tracks.Tracks {
		tracks[idx] = ConvertTrack(&spt)
	}
	artists := make([]*qcast.Artist, len(spotifyResults.Artists.Artists))
	for idx, spt := range spotifyResults.Artists.Artists {
		artists[idx] = ConvertArtist(spt.SimpleArtist)
	}
	albums := make([]*qcast.Album, len(spotifyResults.Albums.Albums))
	for idx, spt := range spotifyResults.Albums.Albums {
		albums[idx] = ConvertAlbum(spt)
	}

	results := &qcast.SearchResponse{
		Tracks:  tracks,
		Artists: artists,
		Albums:  albums,
	}

	return results, nil
}
