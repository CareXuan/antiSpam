FROM golang:1.15-alpine


ENV WORKING_DIR /www
COPY . $WORKING_DIR

RUN cd $WORKING_DIR
#    && go build
#
#CMD /www/antispam