# set base image for build stage
FROM golang:1.15.2-alpine AS build

# set working directory for subsequent instructions
WORKDIR /app

# copy module dependencies and checksums
COPY go.mod go.sum ./

# install dependencies
RUN go mod download

# copy remaining files (i.e. source code)
COPY . .

# build Go binary
RUN GOOS=linux CGO_ENABLED=0 go build -o api cmd/go-api-starter/main.go

# set base image for run stage
FROM scratch

# copy executable from build stage
COPY --from=build /app/api ./

# define entry command for image
ENTRYPOINT [ "./api" ]
