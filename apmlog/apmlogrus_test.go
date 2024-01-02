package apmlog_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/iqbalnew/pkg-helper/apmlog"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func Test_TraceContext(t *testing.T) {
	md := metadata.New(make(map[string]string))
	fmt.Println("log jjjj")

	md.Set("traceparent", "00-fe4e957011188e96aaf32a0ad4890177-fd95b5f5dae15b98-01")
	md.Set("elastic-apm-traceparent", "00-fe4e957011188e96aaf32a0ad4890177-fd95b5f5dae15b98-01")

	logrus.WithField("md", md).Println("md")
	ctx := metadata.NewOutgoingContext(context.TODO(), md)
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		logrus.WithField("md", md).Println("data nya md")
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}

	TraceContext := apmlog.TraceContext(ctx)
	if TraceContext == nil {
		fmt.Println(TraceContext)
		t.Fail()
	}

}