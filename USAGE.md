# Setup the client

## Setup with global authentication

```go
func NetboxClient(host string, token string) *netbox.APIClient {
	configuration := netbox.NewConfiguration()
	configuration.Host = host
	configuration.UserAgent = "myapp/0.1", )
	configuration.Scheme = "https"
	configuration.AddDefaultHeader(
		"Authorization",
		fmt.Sprintf("Token %s", token),
	)
	apiClient := netbox.NewAPIClient(configuration)
	return apiClient
}
```

## Handle authentication per-call

```go
func PrepareNetboxContext(ctx context.Context, token string) context.Context {
	authValues := map[string]netbox.APIKey{
		"tokenAuth": {
			Prefix: "Token",
			Key:    token,
		},
	}
	return context.WithValue(ctx, netbox.ContextAPIKeys, authValues)
}
```

## Generic helper to handle errors

```go
var ErrNetboxUnknownError = errors.New("unknown error raised by Netbox")

func RaiseNetboxError(response *http.Response, err error) error {
	if err == nil {
		err = ErrUnknownNetboxError
	}

	if response == nil {
		return fmt.Errorf("error with empty response: %w", err)
	}

	bodyBytes, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return fmt.Errorf("error reading Netbox body response: %w", readErr)
	}
	bodyString := string(bodyBytes)

	return fmt.Errorf(
		"error during a Netbox request: %w (url: %s, method: %s, body: %s)",
		err,
		response.Request.URL,
		response.Request.Method,
		bodyString,
	)
}
```

# Client usage

## Request a prefix by ID

```go
func GetNetboxPrefixByID(
	ctx context.Context,
	client *netbox.APIClient,
	id int32,
) (*netbox.Prefix, error) {
	prefix, resp, err := client.IpamAPI.
		IpamPrefixesRetrieve(ctx, id).
		Execute()
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		return nil, RaiseNetboxError(resp, err)
	}

	return prefix, nil
}
```
