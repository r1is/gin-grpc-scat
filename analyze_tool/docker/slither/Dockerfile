FROM trailofbits/eth-security-toolbox:latest
ENV LANG C.UTF-8
ENV TOOLNAME Slither
USER root
WORKDIR /home/slither/
ADD ./build/main .
EXPOSE $PORT
CMD ./main -p $PORT -t $TOOLNAME