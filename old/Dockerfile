FROM golang:1.22

LABEL description="Use this Dockerfile to build a container for running tests against the symwalker package."
LABEL version="1.0"
LABEL author="github.com/orme292"

WORKDIR /app

ADD . /app

RUN mkdir -p /tests/start
RUN mkdir -p /tests/start/symlink_target_dir
RUN touch /tests/start/unreadable.file && chmod 000 /tests/start/unreadable.file
RUN mkdir -p /tests/start/unreadable_dir && chmod 000 /tests/start/unreadable_dir
RUN /bin/bash -c 'for i in {1..5}; do touch /tests/start/readable_$i.file; done'

RUN mkdir -p /tests/files/{1,2,3,4,5,6,7,8,9,10}/{a,b,c,d,e,f,g,h,i,j}
RUN mkdir -p /tests/symlink/files
RUN ln -s /tests/start/readable_1.file /tests/symlink/files/link_to_readable_1.file
RUN mkdir -p /tests/symlink/dirs
RUN ln -s /tests/start/symlink_target_dir /tests/symlink/dirs/link_to_readable_dir_1


RUN ls -ahl /tests/start
RUN ls -ahl /tests/symlink
RUN ls -ahl /tests/symlink/files
RUN ls -ahl /tests/symlink/dirs

ENV GOPATH /app
CMD ["go", "test", "./..."]
