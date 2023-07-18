package client

import (
	"context"
	"fmt"
	"time"
)

func (c *ClientImpl) Upload(ctx context.Context, msg string) error {

	upload, err := c.StorjClient.UploadObject(ctx, "blob1", fmt.Sprintf("%v.txt", time.Now().UnixNano()), nil)
	if err != nil {
		return err
	}
	_, err = upload.Write([]byte(msg))
	if err != nil {
		upload.Abort()
		return err
	}

	err = upload.Commit()
	if err != nil {
		upload.Abort()
		return err
	}
	return nil
}
