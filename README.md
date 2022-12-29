# Game Service

This is a simple game service that allows you to play a game of Rock, Paper,
Scissors, Lizard, Spock against computer.

## Features
* Minimal web service application.
* Database support using sqlite.
* Possibility to database migrations and database seeding.
* Logging of errors and success operations.
* Incoming request validation.
* Health check endpoint.

## Project dependencies
To run project locally you will need:
* Installed Go on your machine.

## Services
The following services are used with the scope of the project:
* **app/tooling/admin**: command line tool to help manage database (migrations)
* **app/service/api**: service, that handles incoming request.

*_If you are user  Jetbrains IDE `.http` folder will be useful  to work with 
running service locally. If you run service locally you can specify `host` in 
`http-client-env.json` file and use `choices.http` or `play.http` to make 
request to service._*

## How to play

#### To play the game you need to send a POST request to the `/play` endpoint with  the following payload:

    {
        "player": 1
    }

where `player` is an integer between 1 and 5, inclusive, that represents the
id of possible choices.

The response will be a JSON object with the following structure:

    {
        "player": 1,
        "computer": 2,
        "result": "win"
    }

where `player` and `computer` are integers between 1 and 5, inclusive, that
represent the ids of the choices made by the player and the computer,
respectively, and `result` is a string that represents the result of the game.
Computer always makes a random choice on each request.

#### To get possible choices you need to make a GET request to the `/choices` endpoint.

The response will be a JSON object with the following structure:

    {
        "choices": [
            {
                "id": 1,
                "name": "rock"
            },
            {
                "id": 2,
                "name": "paper"
            },
            {
                "id": 3,
                "name": "scissors"
            },
            {
                "id": 4,
                "name": "lizard"
            },
            {
                "id": 5,
                "name": "spock"
            }
        ]
    }

where `choices` is an array of objects that represent possible choices. Each
object has `id` and `name` fields.

##### To generate a random number between 1 and 5, inclusive, by making a GET request to the `/choice` endpoint.

The response will be a JSON object with the following structure:

    {
        "id": 1,
        "name": "rock"
    }

where `id` is an integer between 1 and 5, inclusive, that represents the id of
selected choice and `name` is a string that represents the name of choice.

##### To get scoreboard of last 10 games you need to make a GET request to the `/scoreboard` endpoint.

The response will be a JSON object with the following structure:

    [
        {
            "results": "tie",
            "player": 2,
            "computer": 2
        },
        {
            "results": "win",
            "player": 2,
            "computer": 1
        },
        {
            "results": "win",
            "player": 2,
            "computer": 1
        },
        {
            "results": "tie",
            "player": 2,
            "computer": 2
        },
        {
            "results": "tie",
            "player": 2,
            "computer": 2
        },
        {
            "results": "tie",
            "player": 2,
            "computer": 2
        },
        {   
            "results": "lose",
            "player": 2,
            "computer": 3
        },
        {
            "results": "win",
            "player": 2,
            "computer": 1
        },
        {
            "results": "tie",
            "player": 2,
            "computer": 2
        },
        {
            "results": "tie",
            "player": 2,
            "computer": 2
        }
    ]

which returns  an array of objects that represent the last 10 round result.

##### To restore scoreboard you need to make a DELETE request to the `/scoreboard` endpoint.
If everything goes well, application will respond with HTTP status code 204 
without content.

##### Possible error responses
If you send a request to the `/play` endpoint with an invalid payload, or 
`choice_id` the response will be a JSON object with the following structure:
_*HTTP status code 400*_

    {
        "request_id": "C9UOV4XQ",
        "status": "Bad request."
    }

If any other error occurs, the response will be a JSON object with the following
_*HTTP status code 500*_

    {
        "request_id": "C9UOV4XQ",
        "status": "Internal server error."
    }

## How to run application locally
Because of application is storing data in database, to successfully use 
application it is better to run database migration.

* **make migrate** : it will create a database with proper schema.

To download all related dependencies:

* **make deps** : it will download all project dependencies. 

To run application you can use:

* **make run** : it will run an application on host and port which is provided 
in **etc/config.yaml**.

To build application you can use:
* **make build** 

To run tests you can use:
* **make test**

