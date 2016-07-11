FROM merxer/docker-demo:base

RUN mkdir /app

COPY src/ /app

WORKDIR /app

CMD  /app/server

EXPOSE 8888
