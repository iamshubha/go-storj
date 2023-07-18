package client

import (
	"context"
	"os"

	"storj.io/uplink"
)

type Client interface {
	Upload(ctx context.Context, msg string) error
}
type ClientImpl struct {
	StorjClient *uplink.Project
}

func StorjClient() *ClientImpl {
	accessToken := os.Getenv("accessToken")
	access, err := uplink.ParseAccess(accessToken)

	if err != nil {
		panic(err)
	}
	storjClient, err := uplink.OpenProject(context.Background(), access)
	if err != nil {
		panic(err)
	}
	defer storjClient.Close()
	return &ClientImpl{
		StorjClient: storjClient,
	}
}
