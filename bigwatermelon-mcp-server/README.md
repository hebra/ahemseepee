# Big Watermelon MCP Server

Big Watermelon is fruit & veggies wholesale store in the east of Melbourne.
The daily deals are announced on their [website](https://www.bigwatermelon.com.au/) 
each morning as a set of images.

This service will download the images ones per day, extract the deal details and makes
them available as a MCP server for AI agents to consume.

## Running

The service uses Google Gemini for image analysis and expects a valid Gemini API key
via the environment variable `GEMINI_API_KEY`, for example:

```shell
export GEMINI_API_KEY=ABCDE-XXXX-YYY-ETC
```

```shekk
curl -X POST http://localhost:8080/mcp \
-H "Content-Type: application/json" \
-d '{
    "action": "get-big-watermelon-deals",
    "parameters": {},
    "request_id": "123"
}'
```

Basic build and run:

```shell
make build
make run
```

Development workflow:
```shell
1. make dev        # Start development server with hot reload
make fmt        # Format code
make lint       # Run linter
```

Testing:
```shell
make test
make coverage
```

Docker operations:
```shell
make docker-build
make docker-run
```
