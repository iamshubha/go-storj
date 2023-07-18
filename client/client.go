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

func StorjClient(ctx context.Context) *ClientImpl {
	access, err := uplink.ParseAccess(os.Getenv("ACCESS_TOKEN"))

	if err != nil {
		panic(err)
	}
	storjClient, err := uplink.OpenProject(ctx, access)
	if err != nil {
		panic(err)
	}
	defer storjClient.Close()
	return &ClientImpl{
		StorjClient: storjClient,
	}
}
