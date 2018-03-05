package lib

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

func PartyCodeFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("Error getting request metadata")
	}

	partyCode, ok := md["party_code"]
	if !ok {
		return "", errors.New("Error: party_code not in metadata")
	}
	return partyCode[0], nil
}
