# notifier
Send notifiations to Telegram

example file structure:

```
your-app/
├── cmd/
│   └── your-app/
│       └── main.go
├── internal/
│   ├── server/
│   │   ├── server.go
│   │   └── routes.go
│   ├── telegram/
│   │   ├── client.go
│   │   └── message.go
├── pkg/  (if you have libraries to be used in other services)
├── api/  (if you have API definitions and protocol files)
├── web/  (if you have frontend web UI)
├── scripts/
├── go.mod
├── go.sum
└── README.md
```

---

Example middleware

```go
// internal/server/user_routes.go

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to check if request is authenticated
		// ...

		if authenticated {
			c.Next()
		} else {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		}
	}
}

func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users", AuthRequired())
	{
		// Your user routes here
	}
}

```