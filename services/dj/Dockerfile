FROM qcast/base:latest as builder
WORKDIR /app
# allow us to only rebuild this image without building the base image
COPY . /app/services/dj
RUN buck build //services/dj --out ./dj


FROM debian:buster-slim  
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/dj /app/dj
CMD ["/app/dj"]  
