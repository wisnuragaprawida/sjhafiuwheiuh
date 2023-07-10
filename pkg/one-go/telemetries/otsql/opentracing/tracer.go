package opentracing

import (
	"context"
	"database/sql/driver"

	"github.com/luna-duclos/instrumentedsql"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

type tracer struct {
	tracerOrphans bool
}

type span struct {
	tracer
	parent opentracing.Span
}

func NewTracer(tracerOrphans bool) instrumentedsql.Tracer {
	return tracer{tracerOrphans: tracerOrphans}
}

func (t tracer) GetSpan(ctx context.Context) instrumentedsql.Span {
	if ctx == nil {
		return span{
			parent: nil,
			tracer: t,
		}
	}
	return span{parent: opentracing.SpanFromContext(ctx), tracer: t}
}

func (s span) NewChild(name string) instrumentedsql.Span {
	if s.parent == nil {
		if s.tracerOrphans {
			return span{parent: opentracing.StartSpan(name), tracer: s.tracer}
		}
		return s
	}
	return span{parent: s.parent.Tracer().StartSpan(name, opentracing.ChildOf(s.parent.Context())), tracer: s.tracer}
}
func (s span) SetLabel(k, v string) {
	if s.parent == nil {
		return
	}
	s.parent.SetTag(k, v)

}
func (s span) SetError(err error) {
	if err == nil || err == driver.ErrSkip {
		return
	}
	if s.parent == nil {
		return
	}

	ext.Error.Set(s.parent, true)
	s.parent.LogFields(
		log.String("Event", "Error"),
		log.String("message", err.Error()),
	)
}

func (s span) Finish() {
	if s.parent == nil {
		return
	}
	s.parent.Finish()
}
