/*
FastAPI

Testing ObservablesAPIService

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

func Test_goyeti_ObservablesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObservablesAPIService AddContextApiV2ObservablesObservableIdContextPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var observableId interface{}

		resp, httpRes, err := apiClient.ObservablesAPI.AddContextApiV2ObservablesObservableIdContextPost(context.Background(), observableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService AddTextApiV2ObservablesAddTextPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ObservablesAPI.AddTextApiV2ObservablesAddTextPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService BulkAddApiV2ObservablesBulkPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ObservablesAPI.BulkAddApiV2ObservablesBulkPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService DeleteApiV2ObservablesObservableIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var observableId string

		resp, httpRes, err := apiClient.ObservablesAPI.DeleteApiV2ObservablesObservableIdDelete(context.Background(), observableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService DeleteContextApiV2ObservablesObservableIdContextDeletePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var observableId interface{}

		resp, httpRes, err := apiClient.ObservablesAPI.DeleteContextApiV2ObservablesObservableIdContextDeletePost(context.Background(), observableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService DetailsApiV2ObservablesObservableIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var observableId interface{}

		resp, httpRes, err := apiClient.ObservablesAPI.DetailsApiV2ObservablesObservableIdGet(context.Background(), observableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService NewApiV2ObservablesPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ObservablesAPI.NewApiV2ObservablesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService NewExtendedApiV2ObservablesExtendedPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ObservablesAPI.NewExtendedApiV2ObservablesExtendedPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService ObservablesRootApiV2ObservablesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ObservablesAPI.ObservablesRootApiV2ObservablesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService PatchApiV2ObservablesObservableIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var observableId interface{}

		resp, httpRes, err := apiClient.ObservablesAPI.PatchApiV2ObservablesObservableIdPatch(context.Background(), observableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService SearchApiV2ObservablesSearchPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ObservablesAPI.SearchApiV2ObservablesSearchPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObservablesAPIService TagObservableApiV2ObservablesTagPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ObservablesAPI.TagObservableApiV2ObservablesTagPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
