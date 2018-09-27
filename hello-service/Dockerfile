FROM alpine
RUN mkdir /app
WORKDIR /app
COPY hello-service /app
CMD [ "/app/hello-service" ]
EXPOSE 8080
# ENV =     添加环境变量
# ENTRYPOINT ["/bin/ash", "-c", "sleep 100000000"]