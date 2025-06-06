// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получение тестового токена
	// (POST /dummyLogin)
	PostDummyLogin(w http.ResponseWriter, r *http.Request)
	// Авторизация пользователя
	// (POST /login)
	PostLogin(w http.ResponseWriter, r *http.Request)
	// Добавление товара в текущую приемку (только для сотрудников ПВЗ)
	// (POST /products)
	PostProducts(w http.ResponseWriter, r *http.Request)
	// Получение списка ПВЗ с фильтрацией по дате приемки и пагинацией
	// (GET /pvz)
	GetPvz(w http.ResponseWriter, r *http.Request, params GetPvzParams)
	// Создание ПВЗ (только для модераторов)
	// (POST /pvz)
	PostPvz(w http.ResponseWriter, r *http.Request)
	// Закрытие последней открытой приемки товаров в рамках ПВЗ
	// (POST /pvz/{pvzId}/close_last_reception)
	PostPvzPvzIdCloseLastReception(w http.ResponseWriter, r *http.Request, pvzId openapi_types.UUID)
	// Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)
	// (POST /pvz/{pvzId}/delete_last_product)
	PostPvzPvzIdDeleteLastProduct(w http.ResponseWriter, r *http.Request, pvzId openapi_types.UUID)
	// Создание новой приемки товаров (только для сотрудников ПВЗ)
	// (POST /receptions)
	PostReceptions(w http.ResponseWriter, r *http.Request)
	// Регистрация пользователя
	// (POST /register)
	PostRegister(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Получение тестового токена
// (POST /dummyLogin)
func (_ Unimplemented) PostDummyLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Авторизация пользователя
// (POST /login)
func (_ Unimplemented) PostLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Добавление товара в текущую приемку (только для сотрудников ПВЗ)
// (POST /products)
func (_ Unimplemented) PostProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получение списка ПВЗ с фильтрацией по дате приемки и пагинацией
// (GET /pvz)
func (_ Unimplemented) GetPvz(w http.ResponseWriter, r *http.Request, params GetPvzParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Создание ПВЗ (только для модераторов)
// (POST /pvz)
func (_ Unimplemented) PostPvz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Закрытие последней открытой приемки товаров в рамках ПВЗ
// (POST /pvz/{pvzId}/close_last_reception)
func (_ Unimplemented) PostPvzPvzIdCloseLastReception(w http.ResponseWriter, r *http.Request, pvzId openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)
// (POST /pvz/{pvzId}/delete_last_product)
func (_ Unimplemented) PostPvzPvzIdDeleteLastProduct(w http.ResponseWriter, r *http.Request, pvzId openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Создание новой приемки товаров (только для сотрудников ПВЗ)
// (POST /receptions)
func (_ Unimplemented) PostReceptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Регистрация пользователя
// (POST /register)
func (_ Unimplemented) PostRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostDummyLogin operation middleware
func (siw *ServerInterfaceWrapper) PostDummyLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostDummyLogin(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostLogin operation middleware
func (siw *ServerInterfaceWrapper) PostLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostLogin(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostProducts operation middleware
func (siw *ServerInterfaceWrapper) PostProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostProducts(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPvz operation middleware
func (siw *ServerInterfaceWrapper) GetPvz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPvzParams

	// ------------- Optional query parameter "startDate" -------------

	err = runtime.BindQueryParameter("form", true, false, "startDate", r.URL.Query(), &params.StartDate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "startDate", Err: err})
		return
	}

	// ------------- Optional query parameter "endDate" -------------

	err = runtime.BindQueryParameter("form", true, false, "endDate", r.URL.Query(), &params.EndDate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "endDate", Err: err})
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPvz(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostPvz operation middleware
func (siw *ServerInterfaceWrapper) PostPvz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostPvz(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostPvzPvzIdCloseLastReception operation middleware
func (siw *ServerInterfaceWrapper) PostPvzPvzIdCloseLastReception(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "pvzId" -------------
	var pvzId openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "pvzId", runtime.ParamLocationPath, chi.URLParam(r, "pvzId"), &pvzId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "pvzId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostPvzPvzIdCloseLastReception(w, r, pvzId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostPvzPvzIdDeleteLastProduct operation middleware
func (siw *ServerInterfaceWrapper) PostPvzPvzIdDeleteLastProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "pvzId" -------------
	var pvzId openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "pvzId", runtime.ParamLocationPath, chi.URLParam(r, "pvzId"), &pvzId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "pvzId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostPvzPvzIdDeleteLastProduct(w, r, pvzId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostReceptions operation middleware
func (siw *ServerInterfaceWrapper) PostReceptions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostReceptions(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostRegister operation middleware
func (siw *ServerInterfaceWrapper) PostRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostRegister(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
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
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/dummyLogin", wrapper.PostDummyLogin)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/login", wrapper.PostLogin)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/products", wrapper.PostProducts)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/pvz", wrapper.GetPvz)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/pvz", wrapper.PostPvz)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/pvz/{pvzId}/close_last_reception", wrapper.PostPvzPvzIdCloseLastReception)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/pvz/{pvzId}/delete_last_product", wrapper.PostPvzPvzIdDeleteLastProduct)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/receptions", wrapper.PostReceptions)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/register", wrapper.PostRegister)
	})

	return r
}

type PostDummyLoginRequestObject struct {
	Body *PostDummyLoginJSONRequestBody
}

type PostDummyLoginResponseObject interface {
	VisitPostDummyLoginResponse(w http.ResponseWriter) error
}

type PostDummyLogin200JSONResponse Token

func (response PostDummyLogin200JSONResponse) VisitPostDummyLoginResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostDummyLogin400JSONResponse Error

func (response PostDummyLogin400JSONResponse) VisitPostDummyLoginResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostLoginRequestObject struct {
	Body *PostLoginJSONRequestBody
}

type PostLoginResponseObject interface {
	VisitPostLoginResponse(w http.ResponseWriter) error
}

type PostLogin200JSONResponse Token

func (response PostLogin200JSONResponse) VisitPostLoginResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostLogin401JSONResponse Error

func (response PostLogin401JSONResponse) VisitPostLoginResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type PostProductsRequestObject struct {
	Body *PostProductsJSONRequestBody
}

type PostProductsResponseObject interface {
	VisitPostProductsResponse(w http.ResponseWriter) error
}

type PostProducts201JSONResponse Product

func (response PostProducts201JSONResponse) VisitPostProductsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostProducts400JSONResponse Error

func (response PostProducts400JSONResponse) VisitPostProductsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostProducts403JSONResponse Error

func (response PostProducts403JSONResponse) VisitPostProductsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type GetPvzRequestObject struct {
	Params GetPvzParams
}

type GetPvzResponseObject interface {
	VisitGetPvzResponse(w http.ResponseWriter) error
}

type GetPvz200JSONResponse []struct {
	Pvz        *PVZ `json:"pvz,omitempty"`
	Receptions *[]struct {
		Products  *[]Product `json:"products,omitempty"`
		Reception *Reception `json:"reception,omitempty"`
	} `json:"receptions,omitempty"`
}

func (response GetPvz200JSONResponse) VisitGetPvzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostPvzRequestObject struct {
	Body *PostPvzJSONRequestBody
}

type PostPvzResponseObject interface {
	VisitPostPvzResponse(w http.ResponseWriter) error
}

type PostPvz201JSONResponse PVZ

func (response PostPvz201JSONResponse) VisitPostPvzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostPvz400JSONResponse Error

func (response PostPvz400JSONResponse) VisitPostPvzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostPvz403JSONResponse Error

func (response PostPvz403JSONResponse) VisitPostPvzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PostPvzPvzIdCloseLastReceptionRequestObject struct {
	PvzId openapi_types.UUID `json:"pvzId"`
}

type PostPvzPvzIdCloseLastReceptionResponseObject interface {
	VisitPostPvzPvzIdCloseLastReceptionResponse(w http.ResponseWriter) error
}

type PostPvzPvzIdCloseLastReception200JSONResponse Reception

func (response PostPvzPvzIdCloseLastReception200JSONResponse) VisitPostPvzPvzIdCloseLastReceptionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostPvzPvzIdCloseLastReception400JSONResponse Error

func (response PostPvzPvzIdCloseLastReception400JSONResponse) VisitPostPvzPvzIdCloseLastReceptionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostPvzPvzIdCloseLastReception403JSONResponse Error

func (response PostPvzPvzIdCloseLastReception403JSONResponse) VisitPostPvzPvzIdCloseLastReceptionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PostPvzPvzIdDeleteLastProductRequestObject struct {
	PvzId openapi_types.UUID `json:"pvzId"`
}

type PostPvzPvzIdDeleteLastProductResponseObject interface {
	VisitPostPvzPvzIdDeleteLastProductResponse(w http.ResponseWriter) error
}

type PostPvzPvzIdDeleteLastProduct200Response struct {
}

func (response PostPvzPvzIdDeleteLastProduct200Response) VisitPostPvzPvzIdDeleteLastProductResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostPvzPvzIdDeleteLastProduct400JSONResponse Error

func (response PostPvzPvzIdDeleteLastProduct400JSONResponse) VisitPostPvzPvzIdDeleteLastProductResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostPvzPvzIdDeleteLastProduct403JSONResponse Error

func (response PostPvzPvzIdDeleteLastProduct403JSONResponse) VisitPostPvzPvzIdDeleteLastProductResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PostReceptionsRequestObject struct {
	Body *PostReceptionsJSONRequestBody
}

type PostReceptionsResponseObject interface {
	VisitPostReceptionsResponse(w http.ResponseWriter) error
}

type PostReceptions201JSONResponse Reception

func (response PostReceptions201JSONResponse) VisitPostReceptionsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostReceptions400JSONResponse Error

func (response PostReceptions400JSONResponse) VisitPostReceptionsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostReceptions403JSONResponse Error

func (response PostReceptions403JSONResponse) VisitPostReceptionsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PostRegisterRequestObject struct {
	Body *PostRegisterJSONRequestBody
}

type PostRegisterResponseObject interface {
	VisitPostRegisterResponse(w http.ResponseWriter) error
}

type PostRegister201JSONResponse User

func (response PostRegister201JSONResponse) VisitPostRegisterResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostRegister400JSONResponse Error

func (response PostRegister400JSONResponse) VisitPostRegisterResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Получение тестового токена
	// (POST /dummyLogin)
	PostDummyLogin(ctx context.Context, request PostDummyLoginRequestObject) (PostDummyLoginResponseObject, error)
	// Авторизация пользователя
	// (POST /login)
	PostLogin(ctx context.Context, request PostLoginRequestObject) (PostLoginResponseObject, error)
	// Добавление товара в текущую приемку (только для сотрудников ПВЗ)
	// (POST /products)
	PostProducts(ctx context.Context, request PostProductsRequestObject) (PostProductsResponseObject, error)
	// Получение списка ПВЗ с фильтрацией по дате приемки и пагинацией
	// (GET /pvz)
	GetPvz(ctx context.Context, request GetPvzRequestObject) (GetPvzResponseObject, error)
	// Создание ПВЗ (только для модераторов)
	// (POST /pvz)
	PostPvz(ctx context.Context, request PostPvzRequestObject) (PostPvzResponseObject, error)
	// Закрытие последней открытой приемки товаров в рамках ПВЗ
	// (POST /pvz/{pvzId}/close_last_reception)
	PostPvzPvzIdCloseLastReception(ctx context.Context, request PostPvzPvzIdCloseLastReceptionRequestObject) (PostPvzPvzIdCloseLastReceptionResponseObject, error)
	// Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)
	// (POST /pvz/{pvzId}/delete_last_product)
	PostPvzPvzIdDeleteLastProduct(ctx context.Context, request PostPvzPvzIdDeleteLastProductRequestObject) (PostPvzPvzIdDeleteLastProductResponseObject, error)
	// Создание новой приемки товаров (только для сотрудников ПВЗ)
	// (POST /receptions)
	PostReceptions(ctx context.Context, request PostReceptionsRequestObject) (PostReceptionsResponseObject, error)
	// Регистрация пользователя
	// (POST /register)
	PostRegister(ctx context.Context, request PostRegisterRequestObject) (PostRegisterResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHttpHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHttpMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// PostDummyLogin operation middleware
func (sh *strictHandler) PostDummyLogin(w http.ResponseWriter, r *http.Request) {
	var request PostDummyLoginRequestObject

	var body PostDummyLoginJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostDummyLogin(ctx, request.(PostDummyLoginRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostDummyLogin")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostDummyLoginResponseObject); ok {
		if err := validResponse.VisitPostDummyLoginResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostLogin operation middleware
func (sh *strictHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	var request PostLoginRequestObject

	var body PostLoginJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostLogin(ctx, request.(PostLoginRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostLogin")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostLoginResponseObject); ok {
		if err := validResponse.VisitPostLoginResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostProducts operation middleware
func (sh *strictHandler) PostProducts(w http.ResponseWriter, r *http.Request) {
	var request PostProductsRequestObject

	var body PostProductsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostProducts(ctx, request.(PostProductsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProducts")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostProductsResponseObject); ok {
		if err := validResponse.VisitPostProductsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetPvz operation middleware
func (sh *strictHandler) GetPvz(w http.ResponseWriter, r *http.Request, params GetPvzParams) {
	var request GetPvzRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetPvz(ctx, request.(GetPvzRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPvz")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetPvzResponseObject); ok {
		if err := validResponse.VisitGetPvzResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostPvz operation middleware
func (sh *strictHandler) PostPvz(w http.ResponseWriter, r *http.Request) {
	var request PostPvzRequestObject

	var body PostPvzJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostPvz(ctx, request.(PostPvzRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostPvz")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostPvzResponseObject); ok {
		if err := validResponse.VisitPostPvzResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostPvzPvzIdCloseLastReception operation middleware
func (sh *strictHandler) PostPvzPvzIdCloseLastReception(w http.ResponseWriter, r *http.Request, pvzId openapi_types.UUID) {
	var request PostPvzPvzIdCloseLastReceptionRequestObject

	request.PvzId = pvzId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostPvzPvzIdCloseLastReception(ctx, request.(PostPvzPvzIdCloseLastReceptionRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostPvzPvzIdCloseLastReception")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostPvzPvzIdCloseLastReceptionResponseObject); ok {
		if err := validResponse.VisitPostPvzPvzIdCloseLastReceptionResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostPvzPvzIdDeleteLastProduct operation middleware
func (sh *strictHandler) PostPvzPvzIdDeleteLastProduct(w http.ResponseWriter, r *http.Request, pvzId openapi_types.UUID) {
	var request PostPvzPvzIdDeleteLastProductRequestObject

	request.PvzId = pvzId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostPvzPvzIdDeleteLastProduct(ctx, request.(PostPvzPvzIdDeleteLastProductRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostPvzPvzIdDeleteLastProduct")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostPvzPvzIdDeleteLastProductResponseObject); ok {
		if err := validResponse.VisitPostPvzPvzIdDeleteLastProductResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostReceptions operation middleware
func (sh *strictHandler) PostReceptions(w http.ResponseWriter, r *http.Request) {
	var request PostReceptionsRequestObject

	var body PostReceptionsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostReceptions(ctx, request.(PostReceptionsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostReceptions")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostReceptionsResponseObject); ok {
		if err := validResponse.VisitPostReceptionsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostRegister operation middleware
func (sh *strictHandler) PostRegister(w http.ResponseWriter, r *http.Request) {
	var request PostRegisterRequestObject

	var body PostRegisterJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostRegister(ctx, request.(PostRegisterRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostRegister")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostRegisterResponseObject); ok {
		if err := validResponse.VisitPostRegisterResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
