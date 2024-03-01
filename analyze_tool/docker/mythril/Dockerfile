FROM mythril/myth:0.24.5
USER root
ENV LANG C.UTF-8
ENV TOOLNAME Mythril

WORKDIR /home/mythril
ADD ./build/main .
RUN chmod +x main

EXPOSE $PORT
CMD ./main -p $PORT -t $TOOLNAME
