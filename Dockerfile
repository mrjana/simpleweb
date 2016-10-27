FROM alpine

RUN apk --update add curl

HEALTHCHECK --interval=3s --timeout=1s --retries=3 \
   CMD curl -f http://127.0.0.1:5000/health/ || exit 1

EXPOSE 5000
COPY simpleweb /usr/bin
