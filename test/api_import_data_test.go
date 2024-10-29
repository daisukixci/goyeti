/*
FastAPI

Testing ImportDataAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package goyeti

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/yeti-platform/goyeti"
	"testing"
)

func Test_goyeti_ImportDataAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImportDataAPIService ImportMispJsonApiV2ImportDataImportMispJsonPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ImportDataAPI.ImportMispJsonApiV2ImportDataImportMispJsonPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}