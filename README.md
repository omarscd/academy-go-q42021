# Omar Saucedo's `Golang Bootcamp` capstone project

Go version: `1.17`

## API

| ENDPOINT             | HTTP Method    | Parameters    |
|----------------------|----------------|---------------|
| /pokemons            | GET            |               |
| /pokemons/:id        | GET            |               |
| /pokemons/ext/:name  | GET            |               |
| /pokemons/type       | GET            |  type: "odd" \| "even", items: int, items_per_worker: int  |

### TODO
- Handle external API error
- Remove config folder or implement config
- Do actual formatting with Pokemon presenters
- Abstract WorkerPool
