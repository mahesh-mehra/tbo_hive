FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV CATBOOST_VERSION=1.2.5
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=1

# --------------------------------------------------
# System dependencies (x86_64)
# --------------------------------------------------
RUN apt-get update && apt-get install -y \
    ca-certificates \
    wget \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# --------------------------------------------------
# CatBoost inference runtime (x86_64 ONLY)
# --------------------------------------------------

# Stable C API header
RUN wget https://raw.githubusercontent.com/catboost/catboost/master/catboost/libs/model_interface/c_api.h \
    -O /usr/local/include/c_api.h

# Prebuilt CatBoost inference library (x86_64)
RUN wget https://github.com/catboost/catboost/releases/download/v${CATBOOST_VERSION}/libcatboostmodel.so \
    -O /usr/local/lib/libcatboostmodel.so

RUN ldconfig

# --------------------------------------------------
# Go (linux/amd64)
# --------------------------------------------------
RUN wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz && \
    rm go1.22.4.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

# --------------------------------------------------
# Application build
# --------------------------------------------------
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o tbo_backend

# --------------------------------------------------
# Run
# --------------------------------------------------
CMD ["./tbo_backend"]
