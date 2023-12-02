# Go Monstérdex
The Go Monstérdex App is a simple yet comprehensive application that serves as a digital encyclopedia for Monstér enthusiasts. It provides information about various Monstér species, their abilities, types, moves, and other relevant details.

## App Features
- User Authentication: Create an account or log in securely.
- Token-Based Authentication: Secure API endpoints.
- Monstér Database: Access detailed information about hundreds of Monstér species.
- Search Functionality: Search for Pokémon by name, type, or other attributes.
- Catch & Release: User can catch monster and release.
## Tech Specifications

### Stack
- Golang(Echo)
- PostgreSQL
- Redis
- GCS(Storage)

### Architecture
The project follows a clean architecture pattern, separating concerns into different layers:

- **Handler Layer**: This layer is responsible for handling incoming requests.
- **Servic Layer**: Contains the business logic.
- **Repository Layer**: Deals with data storage and retrieval data from DB.

## Documentation
### API Specification
[API Specification](http://localhost:2000/swagger/index.html)

### Database Schema
![Database Schema](db_schema.png "DB Schema")

### Postman
[API Specification](https://documenter.getpostman.com/view/28576845/2s9YeK4Vfp)

## Quick Start
### Clone App
```
git clone https://github.com/DitoAdriel99/go-monsterdex.git
```
### Set ENV
```
Change .env.example to .env
```
### Running
if use docker
```
docker compose up -d
```
if use Makefile
```
make start
```