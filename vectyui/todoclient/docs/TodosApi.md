# \TodosApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOne**](TodosApi.md#AddOne) | **Post** /todos | adds a new ToDo
[**Delete**](TodosApi.md#Delete) | **Delete** /todos/{id} | deletes a ToDo
[**DeleteAll**](TodosApi.md#DeleteAll) | **Delete** /todos | deletes all ToDos
[**Find**](TodosApi.md#Find) | **Get** /todos | returns all todos
[**GetOne**](TodosApi.md#GetOne) | **Get** /todos/{id} | gets a ToDo
[**Update**](TodosApi.md#Update) | **Patch** /todos/{id} | updates a ToDo



## AddOne

> ModelToDo AddOne(ctx, todo)

adds a new ToDo

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todo** | [**ModelToDoInput**](ModelToDoInput.md)| Add ToDo | 

### Return type

[**ModelToDo**](model.ToDo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> ModelToDo Delete(ctx, id)

deletes a ToDo

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| ToDo ID | 

### Return type

[**ModelToDo**](model.ToDo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAll

> []ModelToDo DeleteAll(ctx, )

deletes all ToDos

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ModelToDo**](model.ToDo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Find

> []ModelToDo Find(ctx, )

returns all todos

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ModelToDo**](model.ToDo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOne

> ModelToDo GetOne(ctx, id)

gets a ToDo

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| ToDo ID | 

### Return type

[**ModelToDo**](model.ToDo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> ModelToDo Update(ctx, id, todo)

updates a ToDo

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| ToDo ID | 
**todo** | [**ModelToDoInput**](ModelToDoInput.md)| Update ToDo | 

### Return type

[**ModelToDo**](model.ToDo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

