// Package oapigen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package oapigen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Triggers a dryrun for a provided task
	// (POST /v1/dryrun/task)
	ExecuteTaskDryrun(w http.ResponseWriter, r *http.Request)
	// Creates a new task
	// (POST /v1/tasks)
	CreateTask(w http.ResponseWriter, r *http.Request, params CreateTaskParams)
	// Deletes a task by name
	// (DELETE /v1/tasks/{name})
	DeleteTaskByName(w http.ResponseWriter, r *http.Request, name string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// ExecuteTaskDryrun operation middleware
func (siw *ServerInterfaceWrapper) ExecuteTaskDryrun(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ExecuteTaskDryrun(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateTaskParams

	// ------------- Optional query parameter "run" -------------
	if paramValue := r.URL.Query().Get("run"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "run", r.URL.Query(), &params.Run)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter run: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTask(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteTaskByName operation middleware
func (siw *ServerInterfaceWrapper) DeleteTaskByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", chi.URLParam(r, "name"), &name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter name: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteTaskByName(w, r, name)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/dryrun/task", wrapper.ExecuteTaskDryrun)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/tasks", wrapper.CreateTask)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/v1/tasks/{name}", wrapper.DeleteTaskByName)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
