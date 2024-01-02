package apmlog

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/v2"
)

const (
	// FieldKeyTraceID is the field key for the trace ID.
	FieldKeyTraceID = "traceid"

	// FieldKeyTransactionID is the field key for the transaction ID.
	FieldKeyTransactionID = "transactionid"

	// FieldKeySpanID is the field key for the span ID.
	FieldKeySpanID	 = "spanid"
)

// this function create insert field logrus
//	TODO test failed data context
func TraceContext(ctx context.Context) logrus.Fields {
	tx := apm.TransactionFromContext(ctx)
	if tx == nil {
		return nil
	}

	traceContext := tx.TraceContext()
	fields := logrus.Fields{
		FieldKeyTraceID:       traceContext.Trace.String(),
		FieldKeyTransactionID: traceContext.Span.String(),
	}
	if span := apm.SpanFromContext(ctx); span != nil {
		fields[FieldKeySpanID] = span.TraceContext().Span.String()
	}

	return fields
}
