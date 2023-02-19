# fp-Kanbanboard
 
## Pembagian Tugas

| Bagian     | Detail             | Dikerjakan oleh          |
|------------|--------------------|--------------------------|
| Helper     | Generate Token     | Dion Fauzi |
| Helper     | Validate Token     | Dion Fauzi |
| Helper     | Generate Token (Test)     | M Rifqi Al Furqon |
| Helper     | Validate Token (Test)     | M Rifqi Al Furqon |
| Middleware | Authorization      | Dion Fauzi dan M Rifqi Al Furqon |
| Middleware | Authentication      | Dion Fauzi dan M Rifqi Al Furqon |
| Endpoint   | POST /users/register         | Dion Fauzi |
| Endpoint   | POST /users/login         | Dion Fauzi |
| Endpoint   | POST /users/register         | Dion Fauzi |
| Endpoint   | PUT /users/update-account         | Dion Fauzi |
| Endpoint   | DELETE /users/delete-account         | Dion Fauzi |
| Endpoint   | POST /categories         | M Rifqi Al Furqon |
| Endpoint   | GET /categories         | M Rifqi Al Furqon |
| Endpoint   | DELETE /categories/:id         | M Rifqi Al Furqon |
| Endpoint   | PATCH /categories/:id         | M Rifqi Al Furqon |
| Endpoint   | GET /tasks         | M Rifqi Al Furqon |
| Endpoint   | POST /tasks         | M Rifqi Al Furqon |
| Endpoint   | PUT /tasks/:id         | Juan Simon Damanik |
| Endpoint   | PATCH /tasks/update-status/:id         | Juan Simon Damanik |
| Endpoint   | PATCH /tasks/update-category/:id         | Juan Simon Damanik |
| Endpoint   | DELETE /tasks/:id         | Juan Simon Damanik |
| Swagger   | Add Swagger Docs    | Juan Simon Damanik |

## Deployment
Projek dideploy di Railway dengan link berikut [https://fp-kanbanboard-production.up.railway.app/](https://fp-kanbanboard-production.up.railway.app/)

## How to Run
### Locally
- Clone this repo
```
git clone https://github.com/DionFauzi/fp-Kanbanboard
```
- Run PostgreSQL Docker script
```
chmod +x ./scripts/run-postgres.sh && ./scripts/run-postgres.sh
```
- Copy .env.example to .env
```
cp .env.example .env
```
- Run go webserver
```
go run ./main.go
```
- Enjoy!
