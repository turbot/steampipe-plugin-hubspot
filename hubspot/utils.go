package hubspot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	hubspot "github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/properties"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const hubspotAPIBaseURL = "https://api.hubapi.com"

var hubspotHTTPClient = &http.Client{Timeout: 60 * time.Second}

// hubspotPaging models the shared `paging` envelope returned by HubSpot v3
// cursor-paginated endpoints that this plugin calls over raw HTTP.
type hubspotPaging struct {
	Next *hubspotNextPage `json:"next"`
}

type hubspotNextPage struct {
	After string `json:"after"`
	Link  string `json:"link"`
}

// hubspotGet performs an authenticated GET against the HubSpot API and
// unmarshals a successful JSON response into out.
func hubspotGet(ctx context.Context, d *plugin.QueryData, path string, out any) error {
	return hubspotDo(ctx, d, http.MethodGet, path, nil, out)
}

// hubspotPost performs an authenticated POST against the HubSpot API, sending
// body as JSON and unmarshaling a successful JSON response into out.
func hubspotPost(ctx context.Context, d *plugin.QueryData, path string, body any, out any) error {
	return hubspotDo(ctx, d, http.MethodPost, path, body, out)
}

// hubspotDo is the single Bearer-token HTTP path shared by every table that
// the generated library does not cover. It exists to check the response status:
// a non-2xx must surface as an error whose text contains the status code, so the
// plugin's DefaultRetryConfig ("429") and DefaultIgnoreConfig ("404") can act on
// it and a missing-scope 403 fails loudly instead of yielding empty rows.
func hubspotDo(ctx context.Context, d *plugin.QueryData, method, path string, body any, out any) error {
	authorizer, err := connect(ctx, d)
	if err != nil {
		return err
	}

	var reqBody io.Reader
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reqBody = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, hubspotAPIBaseURL+path, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+authorizer.Token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := hubspotHTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("hubspot API %s returned %d: %s", path, resp.StatusCode, string(responseBody))
	}

	if out != nil {
		if err := json.Unmarshal(responseBody, out); err != nil {
			return err
		}
	}

	return nil
}

func connect(ctx context.Context, d *plugin.QueryData) (*hubspot.TokenAuthorizer, error) {
	conn, err := connectAppTokenCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}

	return conn.(*hubspot.TokenAuthorizer), nil
}

var connectAppTokenCached = plugin.HydrateFunc(connectAppTokenUncached).Memoize()

func connectAppTokenUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	// Default to the env var settings
	appToken := os.Getenv("HUBSPOT_PRIVATE_APP_TOKEN")

	// Prefer config settings
	hubSpotConfig := GetConfig(d.Connection)
	if hubSpotConfig.PrivateAppToken != nil {
		appToken = *hubSpotConfig.PrivateAppToken
	}

	if appToken == "" {
		return nil, errors.New("'private_app_token' must be configured")
	}

	authorizer := hubspot.NewTokenAuthorizer(appToken)

	return authorizer, nil
}

func listAllPropertiesByObjectType(ctx context.Context, d *plugin.QueryData, objectType string) ([]properties.Property, error) {
	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAllPropertiesByObjectType", "connection_error", err)
		return []properties.Property{}, nil
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := properties.NewAPIClient(properties.NewConfiguration())
	resp, _, err := client.CoreApi.GetAll(context, objectType).Execute()
	if err != nil {
		plugin.Logger(ctx).Error("listAllPropertiesByObjectType", "api_error", err)
		return []properties.Property{}, nil
	}

	return resp.Results, nil
}

func setDynamicColumnTypes(property properties.Property, column *plugin.Column) {
	switch property.Type {
	case "string":
		column.Type = proto.ColumnType_STRING
	case "number":
		column.Type = proto.ColumnType_DOUBLE
	case "bool":
		column.Type = proto.ColumnType_BOOL
	case "datetime":
		column.Type = proto.ColumnType_TIMESTAMP
	case "enumeration":
		column.Type = proto.ColumnType_STRING
	default:
		column.Type = proto.ColumnType_STRING
	}
}
