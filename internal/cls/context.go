package cls

import (
	"context"
	"strconv"

	cls "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls/v20201016"
)

// GetContextInput holds parameters for DescribeLogContext.
type GetContextInput struct {
	TopicId  string
	PkgId    string
	PkgLogId int64
	BTime    string
	PrevLogs int64
	NextLogs int64
	Query    string
}

// GetContext runs DescribeLogContext and returns the context logs.
func (c *Client) GetContext(ctx context.Context, in GetContextInput) ([]*cls.LogContextInfo, error) {
	req := cls.NewDescribeLogContextRequest()
	req.TopicId = &in.TopicId
	req.PkgId = &in.PkgId
	req.PkgLogId = &in.PkgLogId
	if in.BTime != "" {
		req.BTime = &in.BTime
	}
	if in.PrevLogs > 0 {
		req.PrevLogs = &in.PrevLogs
	}
	if in.NextLogs > 0 {
		req.NextLogs = &in.NextLogs
	}
	if in.Query != "" {
		req.Query = &in.Query
	}

	resp, err := c.api.DescribeLogContextWithContext(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Response == nil || resp.Response.LogContextInfos == nil {
		return nil, nil
	}
	return resp.Response.LogContextInfos, nil
}

// PkgLogIdFromString parses PkgLogId from string (e.g. "65536").
func PkgLogIdFromString(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
