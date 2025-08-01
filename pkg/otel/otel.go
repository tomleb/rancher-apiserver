package otel

import (
	"go.opentelemetry.io/otel"
)

const name = "github.com/rancher/apiserver/pkg/otel"

var (
	Tracer = otel.Tracer(name)
)
