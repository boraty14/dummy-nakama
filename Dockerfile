FROM heroiclabs/nakama-pluginbuilder:3.22.0 AS builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend
COPY . .

RUN go mod download
RUN go build --trimpath --buildmode=plugin -o ./backend.so

FROM heroiclabs/nakama:3.22.0

COPY --from=builder /backend/backend.so /nakama/data/modules
COPY --from=builder /backend/local.yml /nakama/data/
COPY --from=builder /backend/*.json /nakama/data/modules