FROM alpine
COPY README.md README.md
RUN echo "standard docker build" > /build.txt