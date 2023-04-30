# Petssenger

Resolution of [Hash back-end challenge](https://github.com/hashlab/hiring/blob/5ae82743d1afd7f741d59ee63ffa8149ffa12660/challenges/pt-br/backend-finance-challenge.md)

## Getting Started

- Make sure [docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/install/) are both installed
- In the application root run the `docker-compose up --build -d` command
- With **all** containers online, run the `yarn populate` command to populate the databases with the initial data
- Finally, test the application by running the `yarn test` command

## HTTP Routes

- POST http://localhost:3000/ride/estimate - Estimates the price of a race based on the chosen city, distance, and travel time. The estimate is only valid for **20 seconds** due to dynamic city tax. The estimation can be confirmed, or just expire.

```sh
$ curl --location --request POST 'http://localhost:3000/ride/estimate' --header 'Content-Type: application/json' --header 'X-User-ID: 08842beb-a4fc-4cbb2-9f87-d80f1a2d5045' --data-raw '{ "city": "RIO_DE_JANEIRO", "distance": 7.23, "time": 19.6 }'
```

- POST http://localhost:3000/ride/perform - Confirms the last trip estimated by the user. Upon confirming, the city's dynamic rate is increased and only returns to its previous value after **5 minutes**.

```sh
$ curl --location --request POST 'http://localhost:3000/ride/perform' --header 'X-User-ID: 08842beb-a4fc-4cbb2-9f87-d80f1a2d5045'
```

- POST http://localhost:3002/user - Creates a user by a given email address.

```sh
$ curl --location --request POST 'http://localhost:3002/user' --header 'Content-Type: application/json' --data-raw '{ "email": "next@petssenger.com" }'
```

## gRPC Functions

- **GetPricingFeesByCity** - Returns the pricing data (base rate, distance rate, time rate and service rate) of a race for a given city

- **GetDynamicFeesByCity** - Returns the dynamic rate for a given city. Not merged into the above function for caching reasons

- **IncreaseDynamicFeesByCity** - Increments the dynamic rate of a given city

- **AuthUser** - Determines whether a user exists or not by its UUID (X-User-ID)

## License

This project is distributed under the [MIT license](LICENSE)
