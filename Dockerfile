
FROM golang@sha256:8cc1c0f534c0fef088f8fe09edc404f6ff4f729745b85deae5510bfd4c157fb2 as builder

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git ca-certificates curl tzdata && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

WORKDIR /src/
COPY . .

# Fetch dependencies.
RUN go get -d -v

RUN mkdir -p locdata && \
   curl https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz | tar --strip-components=1 -zxvf - -C locdata/


# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /app .

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable
COPY --from=builder /app  /app

COPY --from=builder /src/*.png /

COPY --from=builder /src/browserconfig.xml .

COPY --from=builder /src/favicon.ico .

COPY --from=builder /src/manifest.json .

COPY --from=builder /src/locdata/GeoLite2-City.mmdb /locdata/GeoLite2-City.mmdb

# Use an unprivileged user.
USER appuser

EXPOSE 3000

# Run the hello binary.
ENTRYPOINT ["/app"]




# FROM golang:alpine AS builder

# ENV GO111MODULE=on

# RUN mkdir /user && \
#     echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
#     echo 'nobody:x:65534:' > /user/group

# RUN apk add --no-cache ca-certificates git curl

# WORKDIR /src

# COPY ./go.mod ./go.sum ./
# RUN go mod download && \
#     mkdir locdata && \
#     curl https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz -o locdata/GeoLite2-City.tar.gz

# RUN go get -u github.com/kevinburke/go-bindata && \
#     ls -al / && \
#     ls -al /src && \
#     env && \
#     find / | grep bindata && \
#     go-bindata --nocompress locdata/

# COPY ./ ./

# RUN CGO_ENABLED=0 go build \
#     -installsuffix 'static' \
#     -o /app .

# FROM scratch AS final

# WORKDIR /

# COPY --from=builder /user/group /user/passwd /etc/

# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# COPY --from=builder /app /app

# COPY --from=builder /src/*.png /

# COPY --from=builder /src/browserconfig.xml .

# COPY --from=builder /src/favicon.ico .

# COPY --from=builder /src/manifest.json .

# EXPOSE 3000

# # Perform any further action as an unprivileged user.
# USER nobody:nobody

# # Run the compiled binary.
# ENTRYPOINT ["/app"]