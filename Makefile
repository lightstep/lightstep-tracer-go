# tools
GO=go

default: build

all: thrift proto build

.PHONY: thrift proto

# LightStep-specific: rebuilds the LightStep thrift protocol files.  Assumes
# the command is run within the LightStep development environment (i.e. the
# LIGHTSTEP_HOME environment variable is set).
thrift:
	thrift --gen go:package_prefix='github.com/lightstep/lightstep-tracer-go/',thrift_import='github.com/lightstep/lightstep-tracer-go/thrift_0_9_2/lib/go/thrift' -out . $(LIGHTSTEP_HOME)/go/src/crouton/crouton.thrift
	rm -rf lightstep_thrift/reporting_service-remote

proto:
	@if [ ! -r lightstep-tracer-common/collector.proto ]; then \
	  echo "Must run 'git submodule update --init' before generating protobuf stubs"; \
	  false; \
	fi
	docker run --rm -v $(shell pwd)/lightstep-tracer-common:/input:ro -v $(shell pwd)/collectorpb:/output \
	  lightstep/protoc:latest \
	  protoc --go_out=plugins=grpc:/output --proto_path=/input /input/collector.proto

build: thrift proto
	${GO} build github.com/lightstep/lightstep-tracer-go/...

test:
	$(GO) test
	docker run --rm -v $(GOPATH):/input:ro lightstep/noglog:latest noglog github.com/lightstep/lightstep-tracer-go
