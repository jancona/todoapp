package store

import (
	"context"
	"fmt"
	"log"

	"github.com/jancona/todoapp/vectyui/actions"
	"github.com/jancona/todoapp/vectyui/dispatcher"
	"github.com/jancona/todoapp/vectyui/store/model"
	"github.com/jancona/todoapp/vectyui/store/storeutil"
	"github.com/jancona/todoapp/vectyui/todoclient"
)

var (
	// Items represents all of the TODO items in the store.
	Items []*todoclient.ModelToDo

	// Filter represents the active viewing filter for items.
	Filter = model.All

	// Listeners is the listeners that will be invoked when the store changes.
	Listeners = storeutil.NewListenerRegistry()

	clientConf = todoclient.NewConfiguration()
)

func init() {
	dispatcher.Register(onAction)
	clientConf.Host = ""
	clientConf.Scheme = ""
	clientConf.BasePath = ".."
}

// ActiveItemCount returns the current number of items that are not completed.
func ActiveItemCount() int {
	return count(false)
}

// CompletedItemCount returns the current number of items that are completed.
func CompletedItemCount() int {
	return count(true)
}

func count(completed bool) int {
	count := 0
	for _, item := range Items {
		if item.Completed == completed {
			count++
		}
	}
	return count
}

func onAction(action interface{}) {
	log.Printf("onAction: %#v", action)
	go func() {
		var err error
		switch a := action.(type) {
		case *actions.ReplaceItems:
			Items, err = getAll()
			if err != nil {
				// TODO: display error
				log.Fatalf("Error: %#v", err)
			}

		case *actions.AddItem:
			todo, err := addToDo(todoclient.ModelToDoInput{Title: a.Title})
			if err != nil {
				// TODO: display error
				log.Fatalf("Error: %#v", err)
			}
			Items = append(Items, todo)

		case *actions.DestroyItem:
			err = deleteToDo(*Items[a.Index])
			if err != nil {
				// TODO: display error
				log.Fatalf("Error: %#v", err)
			}
			copy(Items[a.Index:], Items[a.Index+1:])
			Items = Items[:len(Items)-1]

		case *actions.SetTitle:
			Items[a.Index].Title = a.Title
			Items[a.Index], err = updateToDo(*Items[a.Index])
			if err != nil {
				// TODO: display error
				log.Fatalf("Error: %#v", err)
			}

		case *actions.SetCompleted:
			Items[a.Index].Completed = a.Completed
			Items[a.Index], err = updateToDo(*Items[a.Index])
			if err != nil {
				// TODO: display error
				log.Fatalf("Error: %#v", err)
			}

		case *actions.SetAllCompleted:
			for i := range Items {
				Items[i].Completed = a.Completed
				Items[i], err = updateToDo(*Items[i])
				if err != nil {
					// TODO: display error
					log.Fatalf("Error: %#v", err)
				}
			}

		case *actions.ClearCompleted:
			var activeItems []*todoclient.ModelToDo
			for _, item := range Items {
				if item.Completed {
					err = deleteToDo(*item)
					if err != nil {
						// TODO: display error
						log.Fatalf("Error: %#v", err)
					}
				} else {
					activeItems = append(activeItems, item)
				}
			}
			Items = activeItems

		case *actions.SetFilter:
			Filter = a.Filter

		default:
			return // don't fire listeners
		}

		Listeners.Fire(action)
		log.Printf("Listeners.Fire: %#v", action)
	}()
}

func getAll() ([]*todoclient.ModelToDo, error) {
	c := todoclient.NewAPIClient(clientConf)
	todos, r, err := c.TodosApi.Find(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("Unable to get Todos: %v", err)
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status from API call: %d %s", r.StatusCode, r.Status)
	}

	items := make([]*todoclient.ModelToDo, len(todos))
	for i := range todos {
		log.Printf("todo[%d]: %#v", i, todos[i])
		items[i] = &todos[i]
	}
	return items, nil
}

func addToDo(tdi todoclient.ModelToDoInput) (*todoclient.ModelToDo, error) {
	c := todoclient.NewAPIClient(clientConf)
	todo, r, err := c.TodosApi.AddOne(context.TODO(), tdi)
	if err != nil {
		return nil, fmt.Errorf("Unable to add %#v: %v", tdi, err)
	}
	if r.StatusCode != 201 {
		return nil, fmt.Errorf("Bad status from API call: %d %s", r.StatusCode, r.Status)
	}
	return &todo, nil
}

func deleteToDo(todo todoclient.ModelToDo) error {
	c := todoclient.NewAPIClient(clientConf)
	_, r, err := c.TodosApi.Delete(context.TODO(), todo.Id)
	if err != nil {
		return fmt.Errorf("Unable to delete %s: %v", todo.Id, err)
	}
	if r.StatusCode != 204 {
		return fmt.Errorf("Bad status from API call: %d %s", r.StatusCode, r.Status)
	}
	return nil
}

func updateToDo(todo todoclient.ModelToDo) (*todoclient.ModelToDo, error) {
	c := todoclient.NewAPIClient(clientConf)
	todo, r, err := c.TodosApi.Update(
		context.TODO(),
		todo.Id,
		todoclient.ModelToDoInput{
			Title:     todo.Title,
			Completed: todo.Completed,
			Order:     todo.Order,
		})
	if err != nil {
		return nil, fmt.Errorf("Unable to update %s: %v", todo.Id, err)
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status from API call: %d %s", r.StatusCode, r.Status)
	}
	return &todo, nil
}
