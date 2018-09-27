FROM eglass/goruntime:alpine
RUN mkdir /app
WORKDIR /app
COPY data-panel /app
CMD [ "/app/data-panel" ]
EXPOSE 8999
# ENTRYPOINT ["/bin/ash", "-c", "sleep 100000000"]