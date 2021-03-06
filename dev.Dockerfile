FROM golang:1.17.3-stretch

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

# Add a work directory
WORKDIR /app

# Install air for development
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# Start app
CMD air start

