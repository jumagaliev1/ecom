#   Progress Report 3

In this period we're done:
## Registration:

```
curl -i 
-d "{"last_name":"Zhumagaliyev","first_name":"Alibi","email":"alibi@gmail.com","password":"pa$$word","role":"Client"}" 
localhost:4000/v1/users
```
## Authorization:

`curl -d '{"email": "alibi@gmail.com", "password": "pa$$word"}' localhost:4000/v1/tokens/authentication`

```
"authentication_token": {
    "token": "HNI2U2Y3PLMZNNASDHGWUNTN7U",
    "expiry": "2023-02-26T22:48:33.263054979+06:00"
}
```
### Authenticate example
`curl -H "Authorization: Bearer VLZOZ6NHE47EZUR7RAM36AZMXQ" localhost:4000/v1/healthcheck`
## Dynamic filtering:
#### By title (Full-Text Search):

`curl -i localhost:4000/v1/products?title=col`

```
 "products": [
    {
        "id": 7,
        "category": 1,
        "user": 1,
        "title": "Coca cola",
        "description": "Black Drink",
        "price": "460 тг",
        "rating": 2,
        "stock": 1,
        "images": [
            "image1",
            "image2"
        ]
    }
]

```
#### By category:

`curl -i localhost:4000/v1/products?category=meal`

```
 "products": [
    {
        "id": 6,
        "category": 1,
        "user": 1,
        "title": "Milk",
        "description": "White Drink",
        "price": "1490 тг",
        "stock": 1,
        "images": [
            "image1, image2"
        ]
    },
    {
        "id": 7,
        "category": 1,
        "user": 1,
        "title": "Coca cola",
        "description": "Black Drink",
        "price": "460 тг",
        "rating": 2,
        "stock": 1,
        "images": [
            "image1",
            "image2"
        ]
    }
]
```
#### Sort:
`curl -i localhost:4000/v1/products?sort=-price`

```
 "products": [
    {
        "id": 6,
        "category": 1,
        "user": 1,
        "title": "Milk",
        "description": "White Drink",
        "price": "1490 тг",
        "stock": 1,
        "images": [
            "image1, image2"
        ]
    },
    {
        "id": 7,
        "category": 1,
        "user": 1,
        "title": "Coca cola",
        "description": "Black Drink",
        "price": "460 тг",
        "rating": 2,
        "stock": 1,
        "images": [
            "image1",
            "image2"
        ]
    }
]
```

## Pagination (Metadata):
```
"metadata": {
    "current_page": 1,
    "page_size": 20,
    "first_page": 1,
    "last_page": 1,
    "total_records": 2
},
```

## Custom logger:
```
{"level":"INFO","time":"2023-02-25T17:00:18Z","message":"database connection pool established"}
{"level":"INFO","time":"2023-02-25T17:00:18Z","message":"starting server","properties":{"addr":":4000","env":"development"}}

{"level":"INFO","time":"2023-02-25T16:48:27Z","message":"caught signal","properties":{"signal":"interrupt"}}
{"level":"INFO","time":"2023-02-25T16:48:27Z","message":"stopped server","properties":{"addr":":4000"}}
```

## Panic Recovery:
```go
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
```

## Graceful Shutdown:

```go
shutdownError := make(chan error)

go func() {
    quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit

	app.logger.PrintInfo("caught signal", map[string]string{
		"signal": s.String(),
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	shutdownError <- srv.Shutdown(ctx)
}()
```
