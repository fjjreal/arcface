FROM golang:1.14.7

LABEL maintainer="face"

# set timezome
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ADD ./ /data/src
RUN cp /data/src/libs/hr/libarcsoft_face_engine.so /usr/lib/libarcsoft_face_engine.so
RUN cp /data/src/libs/hr/libarcsoft_face.so /usr/lib/libarcsoft_face.so

ADD ./conf /data/app/conf

WORKDIR /data/src
RUN go env -w GOPROXY=https://goproxy.cn
RUN go build -o /data/app/main ./main.go

RUN rm /data/src/conf/arc.json
RUN cp /data/src/conf/example.arc.json /data/src/conf/arc.json
#product delete
#RUN rm -rf /data/src
RUN mkdir /data/app/temp

WORKDIR /data/app

RUN rm /data/app/conf/arc.json
RUN cp /data/app/conf/example.arc.json /data/app/conf/arc.json

EXPOSE 8080
ENTRYPOINT ["./main"]

# docker build -t "hrface:latest" -f dev.Dockerfile .