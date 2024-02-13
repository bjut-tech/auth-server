FROM node:lts as web-builder

WORKDIR /app

COPY ./web/package.json ./web/yarn.lock ./

RUN yarn install --frozen-lockfile

COPY ./web .

RUN yarn build

FROM golang:1.21 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /go/src/github.com/bjut-tech/auth-server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY --from=web-builder /app/dist ./web/dist

RUN go build -a -installsuffix cgo -o /go/bin/auth-server .

FROM gcr.io/distroless/static-debian12:nonroot AS runner

ENV APP_ENV=production

COPY --from=builder --chown=nonroot:nonroot /go/bin/auth-server /usr/bin/auth-server

EXPOSE 8080

ENTRYPOINT ["/usr/bin/auth-server"]
