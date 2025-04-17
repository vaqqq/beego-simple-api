<img src="static/img/title.png" alt="Title Image" width="500" style="display:block; margin: 0 auto;"/>

# Beego Simple API

A simple web API built with [Beego v2](https://github.com/beego/beego) and Go 1.20+.  
It includes rate limiting, XSRF protection, and a styled frontend for testing API endpoints.


## Features
- REST Endpoints (`/hello`, `/time`, `/post`)
- CSRF/XSRF Protection
- Custom Rate Limiting (1 request/second per IP, burst capacity of 5)
- Styled Frontend Interface
- Security Headers (Content Security Policy, HSTS, etc.)

## Get Started
Install [Go](https://go.dev/dl/) before you get started.
```bash
# Clone the repository
git clone https://github.com/vaqqq/beego-simple-api.git

# Install dependencies
go mod tidy

# Run using the Bee tool (if installed)
bee run

# OR run manually
go run main.go
```

Access the API in your browser:  
`http://localhost:8080`

---

## Available Endpoints

| Method | Endpoint   | Description                 |
|--------|------------|-----------------------------|
| GET    | `/hello`   | Returns a hello message     |
| GET    | `/time`    | Returns the server time     |
| POST   | `/post`    | Accepts a message via POST  |

**POST Body Example:**

```
message=HelloFromFrontend
```


## Frontend Web Interface

You can test all API endpoints through the built-in web interface:  
`http://localhost:8080/`


## Built With
- [Beego v2](https://beego.wiki/)
- [Go 1.20+](https://golang.org/dl/)
- [Rate Limiting (golang.org/x/time/rate)](https://pkg.go.dev/golang.org/x/time/rate)
- [Beego XSRF Protection](https://beego.wiki/docs/mvc/controller/xsrf/)
