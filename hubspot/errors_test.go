package hubspot

import (
	"context"
	"errors"
	"testing"
)

// TestStatusCodeMatching verifies that ignore/retry matching keys off the
// parsed HTTP status code for raw-HTTP errors (so a correlationId that happens
// to contain "404"/"429" cannot trigger a false match) while still falling back
// to substring matching for generated-client errors.
func TestStatusCodeMatching(t *testing.T) {
	ctx := context.Background()
	ignore404 := shouldIgnoreErrors([]string{"404"})
	retry429 := shouldRetryError([]string{"429"})

	// The reported bug: a 403 whose body carries a correlationId containing
	// "404" and "429" must not be ignored or retried.
	tricky := &hubspotAPIError{
		Path:       "/settings/v3/users",
		StatusCode: 403,
		Body:       `{"message":"missing scope","correlationId":"5a404e1c-429b-4000-9000-000000000000"}`,
	}
	if ignore404(ctx, nil, nil, tricky) {
		t.Error("403 error with '404' in body must NOT be ignored")
	}
	if retry429(ctx, nil, nil, tricky) {
		t.Error("403 error with '429' in body must NOT be retried")
	}

	// Real status codes match.
	if !ignore404(ctx, nil, nil, &hubspotAPIError{StatusCode: 404}) {
		t.Error("real 404 must be ignored")
	}
	if !retry429(ctx, nil, nil, &hubspotAPIError{StatusCode: 429}) {
		t.Error("real 429 must be retried")
	}

	// A 429 is not a 404.
	if ignore404(ctx, nil, nil, &hubspotAPIError{StatusCode: 429}) {
		t.Error("429 must not match the 404 ignore list")
	}

	// Fallback: generated-client / plain errors still match on substring.
	if !ignore404(ctx, nil, nil, errors.New("googleapi: Error 404: not found")) {
		t.Error("plain error containing 404 should match via substring fallback")
	}
}
