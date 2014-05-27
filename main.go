package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
m.Get("/:one/:action/:two", func(params martini.Params) string {
  return params["one"] + params["action"]+ params["two"]
})
  m.Run()
}
