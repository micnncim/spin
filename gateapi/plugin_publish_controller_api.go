/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/context"
)

// Linger please
var (
	_ context.Context
)

type PluginPublishControllerApiService service

/* PluginPublishControllerApiService Publish a plugin binary and the plugin info metadata.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param plugin plugin
@param pluginId pluginId
@param pluginInfo pluginInfo
@param pluginVersion pluginVersion
@return */
func (a *PluginPublishControllerApiService) PublishPluginUsingPOST(ctx context.Context, plugin *os.File, pluginId string, pluginInfo string, pluginVersion string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/plugins/publish/{pluginId}/{pluginVersion}"
	localVarPath = strings.Replace(localVarPath, "{"+"pluginId"+"}", fmt.Sprintf("%v", pluginId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pluginVersion"+"}", fmt.Sprintf("%v", pluginVersion), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"*/*",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if plugin != nil {
		fbs, _ := ioutil.ReadAll(plugin)
		localVarFileBytes = fbs
		localVarFileName = plugin.Name()
		plugin.Close()
	}
	localVarFormParams.Add("pluginInfo", parameterToString(pluginInfo, ""))
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}
