ARG BASE=centos:7
FROM $BASE

LABEL maintainer="Mikhail Buslaev (buslaevnmh@yandex.ru)"

RUN cd /home
    # update yum utilities
RUN yum -y update
    # inst