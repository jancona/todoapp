/*
 * To Do List API
 *
 * This is a simple todo list API.
 *
 * API version: 0.1.0
 * Contact: jim@anconafamily.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package todoclient
// ModelToDo struct for ModelToDo
type ModelToDo struct {
	Completed bool `json:"completed,omitempty"`
	Id string `json:"id,omitempty"`
	Order int32 `json:"order,omitempty"`
	Title string `json:"title"`
	Url string `json:"url,omitempty"`
}
