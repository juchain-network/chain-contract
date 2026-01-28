FROM ubuntu:24.04

RUN apt-get update && apt-get install -y \
    ca-certificates \
    jq \
    curl \
    iproute2 \
    && rm -rf /var/lib/apt/lists/*

# Copy juchain binary
COPY juchain /usr/local/bin/juchain
COPY start.sh /usr/local/bin/start.sh
RUN chmod +x /usr/local/bin/start.sh

# Set working directory
WORKDIR /data

# Default command
ENTRYPOINT ["/usr/local/bin/start.sh"]