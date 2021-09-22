package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"

	//"github.com/uber/jaeger-client-go/zipkin"

	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	c := &config.Configuration{
		ServiceName: "test",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "127.0.0.1:6831",
		},
	}
	tracer, closer, err := c.NewTracer()
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	fn := func(c *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				tracer,
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
			)
		}
		defer span.Finish()
		c.Request = c.Request.WithContext(newCtx)
		c.Next()

	}

	r := gin.New()
	r.Use(fn)
	group := r.Group("")
	group.GET("", func(c *gin.Context) {
	})
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Printf("Error non-nil %v", err)
	}
	r.ServeHTTP(w, req)
}
