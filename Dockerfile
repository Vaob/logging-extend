ARG BASE=centos:7
FROM $BASE

LABEL maintainer="Mikhail Buslaev (buslaevnmh@yandex.ru)"

RUN cd /home
    # update yum utilities
RUN yum -y update
    # install epel(need for easy golang installation)
RUN yum -y install epel-release
    #install git
RUN yum -y install git
    # install 