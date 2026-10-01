# Dockerfile Generator

A simple command-line tool written in Go that uses a local or remote [Ollama](https://ollama.com/) LLM to generate Dockerfiles from a programming language.

You provide a programming language, and the application sends a prompt to Ollama and prints the generated Dockerfile.

## Features

- Written in Go
- Uses Ollama's HTTP API directly
- No external Go dependencies
- Supports local or remote Ollama servers
- Configurable Ollama model through environment variables
- Simple command-line interface
- Generates Dockerfiles using an LLM

## Requirements

- Go 1.20+
- Ollama
- An Ollama model

### Install Ollama

Install Ollama from:

https://ollama.com/

After installation, make sure Ollama is running.

### Pull a model

The default model used by this application is:

```text
llama3.1:8b
````

 Pull it with:

```
ollama pull llama3.1:8b
```

 You can verify that the model is available:

```
ollama list
```

 ## Getting Started

 Clone the repository:

```
git clone <your-repository-url>
cd Dockerfile-Generator
```

 Run the application:

```
go run main.go
```

 You will be prompted for a programming language:

```
Enter the programming language: golang
```

 The application sends the request to Ollama and prints the generated Dockerfile:

```
Generated Dockerfile:

FROM golang:alpine AS build

WORKDIR /app

RUN go get -d ./...

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]
```

 ## Configuration

 The application can be configured using environment variables.

 ### OLLAMA\_HOST

 Specifies the Ollama server.

 If not specified, the application uses:

```
http://localhost:11434
```

 Example:

```
export OLLAMA_HOST=http://localhost:11434
```

 For a remote Ollama server:

```
export OLLAMA_HOST=http://192.168.1.68:11434
```

 The application also accepts:

```
export OLLAMA_HOST=192.168.1.68:11434
```

 ### OLLAMA\_MODEL

 Specifies which Ollama model to use.

 Default:

```
llama3.1:8b
```

 Example:

```
export OLLAMA_MODEL=qwen3:8b
```

 ## Configuration Examples

 ### Local Ollama

```
export OLLAMA_HOST=http://localhost:11434
export OLLAMA_MODEL=llama3.1:8b

go run main.go
```

 ### Remote Ollama

```
export OLLAMA_HOST=http://192.168.1.68:11434
export OLLAMA_MODEL=llama3.1:8b

go run main.go
```

 ## Windows

 PowerShell:

```
$env:OLLAMA_HOST="http://localhost:11434"
$env:OLLAMA_MODEL="llama3.1:8b"

go run main.go
```

 For a remote Ollama server:

```
$env:OLLAMA_HOST="http://192.168.1.68:11434"
```

 ## How It Works

 The application follows a simple flow:

```
User
  │
  │ Enter programming language
  ▼
Go CLI
  │
  │ POST /api/chat
  ▼
Ollama
  │
  │ LLM generates Dockerfile
  ▼
Go CLI
  │
  ▼
Generated Dockerfile
```

 The application sends a request to Ollama's `/api/chat` endpoint.

 Example request:

```
{
  "model": "llama3.1:8b",
  "messages": [
    {
      "role": "user",
      "content": "ONLY Generate an ideal Dockerfile for golang..."
    }
  ],
  "stream": false
}
```

 ## Project Structure

 The project is intentionally small:

```
Dockerfile-Generator/
│
├── main.go
└── README.md
```

 The application currently uses only the Go standard library, so there is no need to install additional Go packages.

 ## Build

 Build a standalone executable:

```
go build -o dockerfile-generator main.go
```

 Run it:

```
./dockerfile-generator
```

 On Windows:

```
go build -o dockerfile-generator.exe main.go
```

 Then:

```
.\dockerfile-generator.exe
```

 ## Supported Ollama Models

 The application is not tied to a specific model.

 For example:

```
export OLLAMA_MODEL=llama3.1:8b
```

 or:

```
export OLLAMA_MODEL=qwen3:8b
```

 or any other model available in your Ollama installation.

 Check available models with:

```
ollama list
```

 ## Troubleshooting

 ### Cannot connect to Ollama

 If you see:

```
failed to connect to Ollama
```

 make sure Ollama is running.

 You can test the Ollama server with:

```
curl http://localhost:11434/api/tags
```

 If using a remote server, verify:

```
curl http://192.168.1.68:11434/api/tags
```

 and check your `OLLAMA_HOST` value.

 ### Model not found

 If Ollama reports that the model does not exist, pull it first:

```
ollama pull llama3.1:8b
```

 Or configure another installed model:

```
export OLLAMA_MODEL=qwen3:8b
```

 ### Empty language

 The application requires a programming language:

```
Enter the programming language:
Language cannot be empty.
```

 ## Notes

 The generated Dockerfile is produced by an LLM and should be reviewed before using it in production.

 Different models may produce different Dockerfiles for the same programming language.

 The quality of the generated Dockerfile depends on the selected Ollama model and the prompt.