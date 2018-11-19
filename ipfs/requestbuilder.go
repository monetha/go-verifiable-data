package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// requestBuilder is an IPFS commands request builder.
type requestBuilder struct {
	command string
	args    []string
	opts    map[string]string
	headers map[string]string
	body    io.Reader

	ipfs *IPFS
}

// Arguments adds the arguments to the args.
func (r *requestBuilder) Arguments(args ...string) *requestBuilder {
	r.args = append(r.args, args...)
	return r
}

// BodyString sets the request body to the given string.
func (r *requestBuilder) BodyString(body string) *requestBuilder {
	return r.Body(strings.NewReader(body))
}

// BodyBytes sets the request body to the given buffer.
func (r *requestBuilder) BodyBytes(body []byte) *requestBuilder {
	return r.Body(bytes.NewReader(body))
}

// Body sets the request body to the given reader.
func (r *requestBuilder) Body(body io.Reader) *requestBuilder {
	r.body = body
	return r
}

// Option sets the given option.
func (r *requestBuilder) Option(key string, value interface{}) *requestBuilder {
	var s string
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		// slow case.
		s = fmt.Sprint(value)
	}
	if r.opts == nil {
		r.opts = make(map[string]string, 1)
	}
	r.opts[key] = s
	return r
}

// Header sets the given header.
func (r *requestBuilder) Header(name, value string) *requestBuilder {
	if r.headers == nil {
		r.headers = make(map[string]string, 1)
	}
	r.headers[name] = value
	return r
}

// Send sends the request and return the response.
func (r *requestBuilder) Send(ctx context.Context) (*response, error) {
	req := newRequest(r.ipfs.url, r.command, r.args...)
	req.Opts = r.opts
	req.Headers = r.headers
	req.Body = r.body
	return req.Send(ctx, r.ipfs.httpcli)
}

// Exec sends the request a request and decodes the response.
func (r *requestBuilder) Exec(ctx context.Context, res interface{}) error {
	httpRes, err := r.Send(ctx)
	if err != nil {
		return err
	}

	if res == nil {
		httpRes.Close()
		if httpRes.Error != nil {
			return httpRes.Error
		}
		return nil
	}

	return httpRes.Decode(res)
}
