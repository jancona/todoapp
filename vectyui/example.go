package main

import (
	"github.com/gopherjs/vecty"
	"github.com/jancona/todoapp/vectyui/actions"
	"github.com/jancona/todoapp/vectyui/components"
	"github.com/jancona/todoapp/vectyui/dispatcher"
	"github.com/jancona/todoapp/vectyui/store"
)

func main() {
	dispatcher.Dispatch(&actions.ReplaceItems{})

	vecty.SetTitle("Vecty • TodoMVC")
	vecty.AddStylesheet("https://rawgit.com/tastejs/todomvc-common/master/base.css")
	vecty.AddStylesheet("https://rawgit.com/tastejs/todomvc-app-css/master/index.css")
	p := &components.PageView{}
	store.Listeners.Add(p, func(action interface{}) {
		p.Items = store.Items
		vecty.Rerender(p)
	})
	vecty.RenderBody(p)
}
