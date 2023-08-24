# Golang API Course - Final Project (@ Estiam)

## Introduction - 2 Approaches

This repository describes the project assignment that students are required to deliver as part of the assessment of the course.

This project will require the student to **fill up missing code** on a webapp written in golang.

This web-app allows clients to Create, Update, Delete or Read trades. Students are required to complete the CRUD pattern on the webapp.

### The suggested problem

Imagine you're tasked to develop a web-server in golang for a fintech company. The main product at the company allows investors to **buy/sell assets in the stock exchange.**

Assets are described by a name and a sequence of **candlesticks** describing the price of that asset through time.

![Candlestick](/images/candles.png)

[In here you can learn more about candlesticks.](https://en.wikipedia.org/wiki/Candlestick_chart)

A candlestick describes all the transactions of a given asset in a given time. The candlestick can encompass all transactions within 1 minute, 5 minutes or even 1 hour. So candlesticks can come in any periodicity, although some periodicities are more common than others.

The color of a candle indicates how the price changed during the period. If the price decreased, the color of the candle is **red**, and if the price increased, the color is **green**.

* High: the highest price sold during the period
* Low: the lowest price sold during the period
* Open: the first price sold during the period
* Close: the last price sold during the period

The chart below exemplifies how candlesticks compose a candlestick chart. In this example, candlesticks have a 30 minute periodicity.

![Chart](/images/chart.png)

In JSON, a candlestick can be similar to this structure:

```json
{
    "asset": "NASDAQ: AMZN",
    "token": "AMZN",
    "high": "134.69",
    "low": "133.99",
    "open": "134.20",
    "close": "134.62",
    "period": "1m"
}
```

In order to form these candlesticks, however, hundreds of trades happen within 1 minute, and the candlestick is formed after the period of 1 minute ends (in the case of a 1 minute candle).

A **trade** is a simple transaction. A **trade** can be described in JSON like this:

```json
{
    "asset": "NASDAQ: AMZN",
    "token": "AMZN",
    "price": "134.51",
    "maker": "Robert",
    "taker": "Alissa",
    "time": "2023-08-22T12:40:55"
}
```

To decide whether to buy or to sell, investors need some [**technical indicators**](https://en.wikipedia.org/wiki/Technical_indicator). The simplest technical indicators there is is the [**moving average**](https://en.wikipedia.org/wiki/Moving_average).

While trades happen, the moving average keeps constanly updating. It is an average that takes into account only trades in a given period (for example, the moving average of the last 15 minutes of trades).

Knowing all of this, you are tasked to develop a webapp that does the following.

1. Receives trades as HTTP requests
2. Form candlesticks for a given asset and a given periodicity
3. Calculate the moving average for a given period (for example, last 15 minutes) for a given asset (AMZN, for example).
4. Based on the moving average value and the candle, decide whether the user should buy, sell or do nothing with the asset. The user will prompt the criteria for operating in the stock.

### Requirements

#### 1. GET endpoint for Trades

Your webapp should have GET endpoint for trades.

It should receive `token` param to describe the asset you want to get trades from. 

A **suggestion** is to pass token as a **query param**.

```
GET /trades?token=AMZN
```
```
[
    {"id": "x", "asset": "NASDAQ: AMZN", "token": "AMZN" , "price": "134.55", "maker": "Robert", "taker": "Alissa", "time": "2023-08-22T12:40:59"},
    {"id": "y", "asset": "NASDAQ: AMZN", "token": "AMZN" , "price": "134.51", "maker": "John", "taker": "Alissa", "time": "2023-08-22T12:40:55"},
    {"id": "z", "asset": "NASDAQ: AMZN", "token": "AMZN" , "price": "134.49", "maker": "Alissa", "taker": "John", "time": "2023-08-22T12:40:51"},
    ...
]
```

This implementation is only a suggestion. You are free to implement the way you'd like.

Some decisions you have to make:
* If the `token` doesn't exist, should you return an error or simply an empty array?
* What are the status codes for this endpoin? In case of sucess? And in case of empty array?
* Note that trades have an ID. This ID can be created by either the database when storing the trade, or in the code with [uuid library](https://pkg.go.dev/github.com/google/uuid). 

#### 2. POST endpoint for Trades

There should be an endpoint to **store trades** into the database. These trades are going to be used to form candles and calculate the moving average.

```
POST /trades
```
```json
{
    "asset": "NASDAQ: AMZN",
    "token": "AMZN",
    "price": "134.51",
    "maker": "Robert",
    "taker": "Alissa",
    "time": "2023-08-22T12:40:55"
}
```

Some of the decisions you have to make:
* Create an ID for each trade. It can be done by the DB or through code using [uuid library](https://pkg.go.dev/github.com/google/uuid).
* If time or price are malformed, you should respond with a **400 - bad request**.
* All fields are mandatory. If the request is incomplete, respond with **400 - bad request**

#### 3. DELETE endpoint for trades

There should be an endpoint for deleting trades.

```
DELETE /trades/[id]
```
```
[]
```

The endpoint should receive `id` as a **path param**. 

Some of the decisions you have to make:
* What status code should you return in case of success?
* What happens if `id` is not found?

### 4. PATCH endpoint for trades

There should be an patch endpoint for trades, that allows you to change only some fields of the trade object.

```
PATCH /trades/[id]
```
```
{
    "price": "134.51",
    "maker": "Robert",
    "taker": "Alissa",
}
```

You are not allowed to patch `asset`, `token` or `time` in this endpoint.

Some decisions:
* What status code should be returned in case of success?
* If `id` is not found, what is the status code returned?

#### 5. LOGIN endpoint

Let's not forget to add a security layer to our backend.

```
POST /login
```
Body:
```json
{
    "username": "pedro",
    "password": "estiam2023"
}
```

Response:
```json
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNDU2Nzg5LCJuYW1lIjoiSm9zZXBoIn0.OpOSSw7e485LOP5PrzScxHb7SR6sAOMRckfFwi4rp7o
```

This endpoint should create a [JWT](https://github.com/golang-jwt/jwt) based on the user's info and return it to the user.

Some decisions
* How long [will the JWT last](https://jwt.io/)?
You can set the expiration time when creating the JWT.
* What should be returned in case the user doest not exist?
* What should be returned in case the password is wrong and generating the JWT wasn't possible?

#### 6. USER endpoint

The user should be able to create new users on the DB.

```
POST /user
```
Body:
```json
{
    "username": "pedro",
    "password": "estiam2023"
}
```

Some decisions
* Attention when storing the password on the DB. The password should be encrypted. You can use [the bcrypt package](https://pkg.go.dev/golang.org/x/crypto/bcrypt).

#### 6.A - Extra endpoints for USER

You can implement extra endpoints for USER such as `PATCH` and `DELETE` in order to make your web-app more robust.

#### 7. Extra: Middlewares

##### 7.1 Recover middleware

You should add a **recover middleware** to your web-app. When panic ocurrs, to avoid breaking the program, you should be able to recover from the panic in the request pipeline.

A practical example:

```golang
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
```

##### 7.2 Logging middleware

There should be a logger in your application. It is going to be a middleware that's gonna log important info from the request. 

##### 7.3 Authorization middleware

There should be an authorization middleware that expects a `JWT` in the header of every request except `POST /login`.

Every request should be accompanied by a `header` containing the authorization token received in the `/login` endpoint.

For example, GET /trades would be something like this:
```
GET /trades 
    -H "authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
```

The header should be called `authorization`.

Some decisions:
* In case of a bad JWT, you should return a 401 - Unauthorized status code as well as the error message.

#### 7.4 Other middlewares

Feel free to explore and add new middlewares to your application. There are some common ones, such as [CORS middleware](https://www.stackhawk.com/blog/configuring-cors-for-go/).

### 8. Extra: Implementing the business logic (Moving averages)

In this requirement we implement the business logic of our application.

#### 8.1 GET candles endpoint

There should be an endpoint that forms candles and return these to the user based on the trades that were already informed.

In order to form a candle, you should implement the business logic behind it. 

:warning: **Trades should be aggregated to form candlesticks for a given periodicity**.

The request should have the `periodicity` and the `token` describing what trades you want to form a candle from.

`open` will describe the first price of all trades in the period.
`close` will describe the last price of all trades in the period.
`high` is the highest price of all trades in the period.
`low` is the lowest price of all trades in the period.

There should also be a limit to **how many candles you can process on a request**. 1 hour of trades is reasonable.

```
GET /candles?token=AMZN&period=1m
```
Response:
```json
[
    {"id": "x", ..., "period": "1m", "time": "2023-08-22T12:40:00" },
    {"id": "x", ..., "period": "1m", "time": "2023-08-22T12:41:00" },
    {"id": "x", ..., "period": "1m", "time": "2023-08-22T12:42:00" },
...
]
```

_Notice in the example above that the **candle time** conforms to the periodicity of **1 minute**._

Constraints:
* Only these periodicities are allowed:
    - 1m
    - 5m
    - 10m
    - 15m

Some decisions
* You should limit the amount of candles to be generated. 1 hour of trades is enough. For example, 1 hour of trades will have 60 candlesticks of 1 minute each or 4 candlesticks of 15 minutes each.
* You should respond with **400 - bad request** if the periodicity is not one of the 4 periodicities possible (1m, 5m, 10m, 15m).
* Feel free to make other decisions that are not described in here.

#### 8.2 GET moving average endpoint

You also have to implement the logic behind the moving average. It's basically an average that changes through time. When new trades arrive, the average "moves". Old trades are discarded and new trades are integrated into the moving average.

Moving average is, commonly, about the last 15 minutes.

:warning: **For this requirement, you only need to implement moving average for the last 15 minutes.** Feel free to add different time ranges for the moving average.

```
GET /movingaverage?token=AMZN
```
Response:
```json
{
    "token": "AMZN",
    "movingAvg": 134.61
}
```

:warning: **CHALLENGE**

The moving average can be optimized. Trades arrive constanly; each trade that arrives, moves the average to the right. This means an old trade is discarded and a new one is computed into the average.

How would you optimize this to avoid having to calculate the average for all the trades every time a new trade arrives?

Some decisions:
* Respond **400 - bad request** when Token doesn't exist in your DB.
* How to calculate the moving average without doing repetitive computation for every new trade that arrives?

### 9. EXTRA: Redis for caching

As an extra requirement, you can add Redis to your webapp. Redis is a caching mechanism that allows you to respond request fastly by avoiding DB lookups. This is essential on a fintech business, because the investor wants a quick response to the market fluctuations. **A good example of Redis use would be storing the moving average there.** But you can also keep trades in the redis cache.

### 10. Checklist

This is a checklist that will guide the assessment of the final project. Every feature must be working as expected.

- [ ] CRUD for Trades (GET, POST, PATCH, DELETE)
- [ ] POST endpoint for User
- [ ] POST endpoint for Login

Extra points if you implement the logic and some middlewares:
- [ ] GET endpoint for Candles
- [ ] GET endpoint for Moving Average
- [ ] MIDDLEWARES: Logging, Authorization, Recovery.

### 10. The database

We will be using PostgreSQL as our main database. We won't be learning it deeply, but the student is permitted to choose any Database for their project.

### Assessment

:warning: **The project should be submitted as a compressed file (zip/tar/rar) on Microsoft Teams. It can be done in groups of up to 5 people. Only one submission per group.**

:warning: **The names of the group members should be inside the code in commentary on the main module.**