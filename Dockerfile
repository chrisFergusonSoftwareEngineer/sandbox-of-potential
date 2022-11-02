FROM ubuntu:22.04

RUN apt-get update
RUN apt-get -y upgrade

RUN echo eval $(ssh-agent -s)" >> /root/.bashrc

RUN echo "Y" | apt install git
RUN echo "Y" | apt install curl
RUN echo "Y" | apt install unzip

CMD ["bash"]