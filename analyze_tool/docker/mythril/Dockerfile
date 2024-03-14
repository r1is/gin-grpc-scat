FROM mythril/myth:0.24.5
ENV LANG C.UTF-8
ENV TOOLNAME Mythril

WORKDIR /home/mythril
ADD ./build/main .
USER root
RUN chmod +x main
USER mythril

EXPOSE $PORT
CMD ./main -p $PORT -t $TOOLNAME
