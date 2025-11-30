# ---- Build stage ----
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Copy root go.mod & go.sum (monorepo support)
COPY ../../go.mod ../../go.sum ./

# Copy the entire repository (monorepo)
COPY ../../ ./

# Now CD into the service folder
WORKDIR /app/services/file-validation

RUN go mod tidy
RUN go build -o file-validation

# ---- Run stage ----
FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/services/file-validation/file-validation .

EXPOSE 8010
CMD ["./file-validation"]
