package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Tracing(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
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
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}
		defer span.Finish()

		var traceID, spanID string
		spanContext := span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext := spanContext.(jaeger.SpanContext)
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
		}

		c.Set("trace_id", traceID)
		c.Set("span_id", spanID)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
