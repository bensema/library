// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Based on https://github.com/DataDog/dd-trace-go/blob/8fb554ff7cf694267f9077ae35e27ce4689ed8b6/contrib/gin-gonic/gin/option.go

package blademaster

import (
	otelpropagation "go.opentelemetry.io/otel/api/propagation"
	oteltrace "go.opentelemetry.io/otel/api/trace"
)

type config struct {
	TracerProvider oteltrace.Provider
	Propagators    otelpropagation.Propagators
}

// Option specifies instrumentation configuration options.
type Option func(*config)

// WithPropagators specifies propagators to use for extracting
// information from the HTTP requests. If none are specified, global
// ones will be used.
func WithPropagators(propagators otelpropagation.Propagators) Option {
	return func(cfg *config) {
		cfg.Propagators = propagators
	}
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider oteltrace.Provider) Option {
	return func(cfg *config) {
		cfg.TracerProvider = provider
	}
}