FROM python:3.9.0a3-alpine3.10
LABEL maintainer="weslenng"

WORKDIR /petssenger/playground
COPY playground/src/. .

EXPOSE 3000
CMD ["python3", "-m", "http.server", "3000"]
