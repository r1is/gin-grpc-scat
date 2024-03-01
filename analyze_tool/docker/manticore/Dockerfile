FROM trailofbits/manticore:latest
USER root
WORKDIR /home

ENV LANG C.UTF-8
ENV TOOLNAME Manticore

ADD ./build/main .
RUN chmod +x main
# 设置 LD_PRELOAD 环境变量，使程序在运行时使用 glibc 2.31
EXPOSE $PORT
CMD ./main -p $PORT -t $TOOLNAME