package middlewaries

import (
	"context"
	"net/http"

	"githib.com/igomonov88/game-service/internal/random"
)

type RequestIDContextKey string

const requestID = RequestIDContextKey("request_id")

func RequestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestID).(string)
}

func RequestIDFromRequest(r *http.Request) string {
	return RequestIDFromContext(r.Context())
}

func WithRequestIDContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestID, id)
}

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(WithRequestIDContext(r.Context(), random.String(8)))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
