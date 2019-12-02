package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"t-mk-opentrace/ext/log-driver/log"

	"github.com/opentracing/opentracing-go"
)

// Request get
func Request(
	ctx context.Context,
	client *http.Client,
	url string,
	in io.Reader,
	out interface{},
	method string,
	fnOpts ...func(r *http.Request)) error {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		span = opentracing.StartSpan(method + "-" + url)
	}
	httpReq, err := http.NewRequest(method, url, in)
	if err != nil {
		return err
	}
	for _, fn := range fnOpts {
		fn(httpReq)
	}
	// Transmit the span's TraceContext as HTTP headers on our
	// outbound request.
	c := span.Context()
	if err := opentracing.GlobalTracer().Inject(
		c,
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(httpReq.Header)); err != nil {
		log.Warn(err)
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	if err := ReadJSON(resp, out); err != nil {
		return err
	}
	return fmt.Errorf("必须有span")
}

// ReadJSON reads JSON from http.Response and parses it into `out`
func ReadJSON(resp *http.Response, out interface{}) error {
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Warn(err)
		}
	}()

	if resp.StatusCode >= 400 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}

	if out == nil {
		if _, err := io.Copy(ioutil.Discard, resp.Body); err != nil {
			return err
		}
		return nil
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
