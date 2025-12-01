## What is Session-Based Authentication?

Session-based authentication is a stateful authentication technique where we use sessions to keep track of the authenticated user. Here is how Session Based Authentication works:

- User submits the login request for authentication.

- Server validates the credentials. If the credentials are valid, the server initiates a session and stores some information about the client. This information can be stored in memory, file system, or database. The server also generates a unique identifier that it can later use to retrieve this session information from the storage. Server sends this unique session identifier to the client.

- Client saves the session id in a cookie and this cookie is sent to the server in each request made after the authentication.

- Server, upon receiving a request, checks if the session id is present in the request and uses this session id to get information about the client.

And that is how session-based authentication works.

[font](https://roadmap.sh/guides/session-based-authentication)


## How to run

First of all, you have to clone the repository

```bash
git clone git@github.com:cpared/session-based-auth.git
```

Then you have to run this REST API with the following command in a terminal

```bash
go run ./cmd/api
```

This will run the server in port 8080

## API DOC

### Login

You just only send a request to the login path with a user and password. Here is an example with a valid user and password that is previous hardcode in the code

##### Curl

```bash
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "user": "test_user",
    "password": "12345"
}'
```

##### Body

```JSON
{
    "user": "test_user",
    "password": "12345"
}
```

##### Response

- 200 OK 
- 401 Unauthorized

##### Set-Cookie

```
session_id=<token>; HttpOnly; Path=/; SameSite=Strict; Expires=<Date>
```

### Logout

This endpoint needs the sessionID that was sent in the cookie when you login

##### Curl

```bash
curl --location --request POST 'http://localhost:8080/logout' \
--header 'Cookie: sessionID=5a17e271-a848-4b65-bb0f-873445f20f35'
```

##### Response

- 200 OK 
- 401 Unauthorized

### Pokemon Types

This is only endpoint that wrappe a pokemon API but you need to be login first (thats why I create this repo)

```bash
curl --location 'http://localhost:8080/types/pokemons/1' \
--header 'Cookie: sessionID=5a17e271-a848-4b65-bb0f-873445f20f35'
```

[PokeAPI docs](https://pokeapi.co/docs/v2#types)