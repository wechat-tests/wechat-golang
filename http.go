package wxxx

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpMethod string

const (
	POST HttpMethod = "POST"
	GET  HttpMethod = "GET"
)

type HttpCli interface {
	DO(ctx *RequestCtx) (*ResponseCtx, error)
}

type RequestCtx struct {
	HttpMethod  HttpMethod
	Url         string
	ContentType string
	UserAgent   string
	//KeepAlive     string
	//Accept        string
	CustomHeaders map[string]string
	BodyData      []byte
}
type ResponseCtx struct {
	StatusCode int
	Status     string
	Headers    http.Header
	BodyData   []byte
}

type defaultHttpCli struct {
	cli http.Client
	ctx context.Context
}

//默认 http 访问的超时时间
const kDefaultHttpTimeout = 60 * time.Second

func NewDefaultHttpCli(ctx context.Context) HttpCli {
	if ctx == nil {
		ctx = context.Background()
	}
	return &defaultHttpCli{
		cli: http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: kDefaultHttpTimeout,
		},
		ctx: ctx,
	}
}

func (this *defaultHttpCli) DO(ctx *RequestCtx) (*ResponseCtx, error) {
	AssertNotNil(ctx, "request ctx must not be nil")
	AssertNotEmpty(ctx.UserAgent, "request UserAgent must not empty")
	rspBytes := make([]byte, 0)
	reqBytes := ctx.BodyData
	start := time.Now()
	statusCode := 0
	method := ctx.HttpMethod
	fullUrl := ctx.Url
	defer func() {
		stop := time.Now()
		latency := stop.Sub(start)
		buf := bytes.Buffer{}
		buf.WriteString(fmt.Sprintf("[%s](%s) status:%d latency : %d ms\n", method, fullUrl, statusCode, latency.Milliseconds()))
		buf.WriteString(fmt.Sprintf("request : len= %d >>> %s \n", len(reqBytes), ArrToHexStrWithSp(reqBytes, " ")))
		buf.WriteString(fmt.Sprintf("response : len= %d >>> %s \n", len(rspBytes), ArrToHexStrWithSp(rspBytes, " ")))
		Log.PrintLn(buf.String())
	}()
	withTimeout, cancelFunc := context.WithTimeout(this.ctx, kDefaultHttpTimeout)
	defer cancelFunc()
	var reqBodyReader io.Reader = nil
	if len(ctx.BodyData) > 0 {
		reqBodyReader = bytes.NewReader(ctx.BodyData)
	}
	req, err := http.NewRequestWithContext(withTimeout, string(ctx.HttpMethod), ctx.Url, reqBodyReader)
	if err != nil {
		return nil, errors.WithMessage(err, "build http request failed")
	}
	req.Header.Set("User-Agent", ctx.UserAgent)
	if len(ctx.CustomHeaders) > 0 {
		for k, v := range ctx.CustomHeaders {
			if IsNotEmptyStr(k) && IsNotEmptyStr(v) {
				req.Header.Set(k, v)
			}
		}
	}
	if ctx.HttpMethod == POST {
		AssertNotEmpty(ctx.ContentType, "request ContentType must not empty")
		req.Header.Set("Content-Type", ctx.ContentType)
		if len(ctx.BodyData) > 0 {
			req.Header.Set("Content-Length", fmt.Sprintf("%d", len(ctx.BodyData)))
		}
	}
	response, err := this.cli.Do(req)
	if err != nil {
		return nil, errors.WithMessagef(err, " http '%s' '%s' failed !", ctx.HttpMethod, ctx.Url)
	}
	rspBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithMessage(err, " http read body failed !")
	}
	statusCode = response.StatusCode
	resCtx := &ResponseCtx{
		StatusCode: response.StatusCode,
		Status:     response.Status,
		Headers:    response.Header,
		BodyData:   rspBytes,
	}
	return resCtx, nil
}
