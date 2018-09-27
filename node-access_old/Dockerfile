FROM golang
RUN mkdir /app
WORKDIR /app
COPY nodeaccess /app
CMD [ "/app/nodeaccess" ]
EXPOSE 9999
# ENTRYPOINT ["/bin/ash", "-c", "sleep 100000000"]