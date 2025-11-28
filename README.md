# Flight Aggregator (Go)

- **http.Server** + **http.ServeMux** (2 routes)
- `GET /health` → **200 OK**
- `GET /flight` → client HTTP vers 2 APIs, **json.NewDecoder** → structs, tri
- Tri par `price`, `time_travel`, `departure_date` via query params
- Architecture **controller / service / repo**
- Repositories avec **interface commune**
- **Viper** : `viper.AutomaticEnv()` + `viper.Get("MY_VAR")`
- **Tests** : algorithmes de tri + service (repositories mockés)
- **Docker Compose** + **air** (reload auto)

## Démarrer
```powershell
docker compose up --build
```

Endpoints :
- `http://localhost:8080/health`
- `http://localhost:8080/flight`
- `http://localhost:8080/flight?sort_by=time_travel&order=asc`
- `http://localhost:8080/flight?sort_by=departure_date&order=desc`

Providers (host) :
- `http://localhost:4001/flights` (j-server1)
- `http://localhost:4002/flights` (j-server2)
(Intra-Docker : `http://j-server1:4001/flights`, `http://j-server2:4001/flights`)

## Tests
```powershell
docker compose exec aggregator make test
```

Structure :
```
internal/
  controller/      # routes HTTP (ServeMux)
  service/         # logique d'agrégation/tri
  repository/      # interface + 2 implémentations
  model/           # Flight & DTO
  config/          # Viper (env)
  httpserver/      # http.Server + mux
tests/             # tests unitaires (sort + service)
```
