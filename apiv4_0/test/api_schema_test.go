/*
NetBox REST API

Testing SchemaAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package apiv4_0

import (
	"context"
	"testing"

	openapiclient "github.com/jqueuniet/netbox-go/apiv4_0"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_apiv4_0_SchemaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SchemaAPIService SchemaRetrieve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SchemaAPI.SchemaRetrieve(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}