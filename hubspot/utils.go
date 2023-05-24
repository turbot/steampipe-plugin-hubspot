package hubspot

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/clarkmcc/go-hubspot/generated/v1/oauth"

	hubspot "github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/properties"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type tokenInfo struct {
	Token          string
	ExpirationTime time.Time
}

func connect(ctx context.Context, d *plugin.QueryData) (*hubspot.TokenAuthorizer, error) {
	conn, err := connectAppTokenCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	if conn != nil {
		plugin.Logger(ctx).Error("app")
		return conn.(*hubspot.TokenAuthorizer), nil
	} else {
		plugin.Logger(ctx).Error("oauhth")
		token, err := fetchOAuthToken(ctx, d)
		if err != nil {
			return nil, err
		}
		if token != nil {
			return hubspot.NewTokenAuthorizer(*token), nil
		}
	}

	// Credentials not set
	return nil, errors.New("private_app_token or oauth credentials must be configured")
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
		return nil, nil
	}

	authorizer := hubspot.NewTokenAuthorizer(appToken)

	return authorizer, nil
}

func fetchOAuthToken(ctx context.Context, d *plugin.QueryData) (*string, error) {
	// Default to the env var settings
	clientID := os.Getenv("HUBSPOT_CLIENT_ID")
	clientSecret := os.Getenv("HUBSPOT_CLIENT_SECRET")
	refreshToken := os.Getenv("HUBSPOT_REFRESH_TOKEN")

	// Prefer config settings
	hubSpotConfig := GetConfig(d.Connection)
	if hubSpotConfig.ClientID != nil {
		clientID = *hubSpotConfig.ClientID
	}
	if hubSpotConfig.ClientSecret != nil {
		clientSecret = *hubSpotConfig.ClientSecret
	}
	if hubSpotConfig.RefreshToken != nil {
		refreshToken = *hubSpotConfig.RefreshToken
	}

	if refreshToken == "" || clientID == "" || clientSecret == "" {
		return nil, nil
	}

	cacheKey := refreshToken + clientID + clientSecret
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		tokenDetails := cachedData.(tokenInfo)
		if !tokenDetails.ExpirationTime.Before(time.Now()) {
			return &tokenDetails.Token, nil
		}
	}
	plugin.Logger(ctx).Error("not cached")
	apiClient := oauth.NewAPIClient(oauth.NewConfiguration())
	resp, _, err := apiClient.TokensApi.CreateToken(context.Background()).GrantType("refresh_token").ClientId(clientID).ClientSecret(clientSecret).RefreshToken(refreshToken).Execute()
	if err != nil {
		return nil, err
	}

	// set the expiration time for the token to be 5 minutes less than the actual expiry time
	expirationTime := time.Now().Add(time.Second * time.Duration(resp.ExpiresIn-300))
	tokenData := tokenInfo{
		Token:          resp.AccessToken,
		ExpirationTime: expirationTime,
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, tokenData)

	return &resp.AccessToken, nil
}

func listAllPropertiesByObjectType(ctx context.Context, d *plugin.QueryData, objectType string) ([]string, error) {
	cacheKey := objectType
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.([]string), nil
	}

	authorizer, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAllPropertiesByObjectType", "connection_error", err)
		return nil, err
	}
	context := hubspot.WithAuthorizer(context.Background(), authorizer)
	client := properties.NewAPIClient(properties.NewConfiguration())
	resp, _, err := client.CoreApi.GetAll(context, objectType).Execute()
	if err != nil {
		plugin.Logger(ctx).Error("listAllPropertiesByObjectType", "api_error", err)
		return nil, err
	}

	propertyNames := []string{}
	for _, property := range resp.Results {
		propertyNames = append(propertyNames, property.Name)
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, propertyNames)

	return propertyNames, nil
}
