FROM alpine
COPY README.md README.md
RUN echo "standard docker built" > /built.txt