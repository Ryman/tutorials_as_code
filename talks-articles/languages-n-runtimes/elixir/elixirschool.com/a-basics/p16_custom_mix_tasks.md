
## Custom mix tasks

### Intro

* mix provide some pre-existing tasks like

```
mix help  ## list pre-existing mix tasks

mix new my_project  ## creates new elixir project

mix local  ## lists local tasks

mix phx.new my_phoenix  ## creating new phoenix app

mix release  ## assembles a self-contained release

mix run  ## starts and run current project

mix test  ## run project's test
```


### Setup

* create a new mix application say by `mix new hello_elixir ; cd hello_elixir ; mix test`

* in `lib/hello_elixir.ex` file generated by Mix, add following

```
defmodule HelloElixir do
  @doc """
  Outputs `Hello!` every time.
  """
  def say do
    IO.puts("Hello!")
  end
end
```


### Custom mix tasks

* create a new dir and file `lib/mix/tasks/hello_elixir.ex` for custom mix task code as below

```
defmodule Mix.Tasks.HelloElixir do
  use Mix.Task

  @shortdoc "Simply calls the HelloElixir.say/0 function."
  def run(_) do
    HelloElixir.say()
  end
end
```


### mix tasks in action

* to run custom command snake-case of name in `Mix.Tasks.<name>` is used as `mix hello_elixir`

* to list this new custom task in `mix help`, compile is needed which can be triggered by `mix compile` or running the task itself

---
