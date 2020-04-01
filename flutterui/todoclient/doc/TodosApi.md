# todoclient.api.TodosApi

## Load the API package
```dart
import 'package:todoclient/api.dart';
```

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addOne**](TodosApi.md#addOne) | **POST** /todos | adds a new ToDo
[**delete**](TodosApi.md#delete) | **DELETE** /todos/{id} | deletes a ToDo
[**deleteAll**](TodosApi.md#deleteAll) | **DELETE** /todos | deletes all ToDos
[**find**](TodosApi.md#find) | **GET** /todos | returns all todos
[**getOne**](TodosApi.md#getOne) | **GET** /todos/{id} | gets a ToDo
[**update**](TodosApi.md#update) | **PATCH** /todos/{id} | updates a ToDo


# **addOne**
> ModelToDo addOne(todo)

adds a new ToDo

### Example 
```dart
import 'package:todoclient/api.dart';

var api_instance = TodosApi();
var todo = ModelToDoInput(); // ModelToDoInput | Add ToDo

try { 
    var result = api_instance.addOne(todo);
    print(result);
} catch (e) {
    print("Exception when calling TodosApi->addOne: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todo** | [**ModelToDoInput**](ModelToDoInput.md)| Add ToDo | 

### Return type

[**ModelToDo**](ModelToDo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete**
> ModelToDo delete(id)

deletes a ToDo

### Example 
```dart
import 'package:todoclient/api.dart';

var api_instance = TodosApi();
var id = 38400000-8cf0-11bd-b23e-10b96e4ef00d; // String | ToDo ID

try { 
    var result = api_instance.delete(id);
    print(result);
} catch (e) {
    print("Exception when calling TodosApi->delete: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**String**](.md)| ToDo ID | [default to null]

### Return type

[**ModelToDo**](ModelToDo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteAll**
> List<ModelToDo> deleteAll()

deletes all ToDos

### Example 
```dart
import 'package:todoclient/api.dart';

var api_instance = TodosApi();

try { 
    var result = api_instance.deleteAll();
    print(result);
} catch (e) {
    print("Exception when calling TodosApi->deleteAll: $e\n");
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List<ModelToDo>**](ModelToDo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find**
> List<ModelToDo> find()

returns all todos

### Example 
```dart
import 'package:todoclient/api.dart';

var api_instance = TodosApi();

try { 
    var result = api_instance.find();
    print(result);
} catch (e) {
    print("Exception when calling TodosApi->find: $e\n");
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List<ModelToDo>**](ModelToDo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getOne**
> ModelToDo getOne(id)

gets a ToDo

### Example 
```dart
import 'package:todoclient/api.dart';

var api_instance = TodosApi();
var id = 38400000-8cf0-11bd-b23e-10b96e4ef00d; // String | ToDo ID

try { 
    var result = api_instance.getOne(id);
    print(result);
} catch (e) {
    print("Exception when calling TodosApi->getOne: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**String**](.md)| ToDo ID | [default to null]

### Return type

[**ModelToDo**](ModelToDo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update**
> ModelToDo update(id, todo)

updates a ToDo

### Example 
```dart
import 'package:todoclient/api.dart';

var api_instance = TodosApi();
var id = 38400000-8cf0-11bd-b23e-10b96e4ef00d; // String | ToDo ID
var todo = ModelToDoInput(); // ModelToDoInput | Update ToDo

try { 
    var result = api_instance.update(id, todo);
    print(result);
} catch (e) {
    print("Exception when calling TodosApi->update: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**String**](.md)| ToDo ID | [default to null]
 **todo** | [**ModelToDoInput**](ModelToDoInput.md)| Update ToDo | 

### Return type

[**ModelToDo**](ModelToDo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

