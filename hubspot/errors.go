package hubspot

import (
	"context"
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

// shouldIgnoreErrors:: function which returns an ErrorPredicate for HubSpot API calls
func shouldIgnoreErrors(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		return matchStatusCodes(err, notFoundErrors)
	}
}

func shouldRetryError(retryErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		return matchStatusCodes(err, retryErrors)
	}
}

// matchStatusCodes reports whether err matches any of the given status codes.
// For errors raised by the raw HTTP helpers it compares the parsed status code
// exactly; for errors from the generated client library it falls back to a
// substring match on the error text.
func matchStatusCodes(err error, codes []string) bool {
	var apiErr *hubspotAPIError
	if errors.As(err, &apiErr) {
		return slices.Contains(codes, strconv.Itoa(apiErr.StatusCode))
	}
	for _, code := range codes {
		if strings.Contains(err.Error(), code) {
			return true
		}
	}
	return false
}
