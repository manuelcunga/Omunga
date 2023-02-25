# Imagem para desenvolvimento
FROM golang:1.19 AS dev

WORKDIR /app

# Instalação do nodejs e nodemon
RUN apt-get update && apt-get install -y nodejs npm
RUN npm install -g nodemon

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["nodemon", "-L", "--exec", "go", "run", "main.go"]

# Imagem para produção
# FROM golang:1.19 AS prod

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# CMD ["main"]
