FROM golang:1.22

LABEL description="Use this Dockerfile to build a container for running tests against the symwalker package."
LABEL version="1.0"
LABEL author="github.com/orme292"

WORKDIR /app

ADD . /app

RUN /bin/bash -c 'mkdir -p /tests/users/{andrew,brian,carolyn,david,erin,frank}/{downloads,documents,pictures}'
RUN /bin/bash -c 'for i in {1..5}; do touch /tests/users/{andrew,brian,carolyn,david,erin,frank}/pictures/$i.jpg; done'
RUN /bin/bash -c 'for i in {1..6}; do touch /tests/users/{andrew,brian,carolyn,david,erin,frank}/documents/$i-report.doc; done'
RUN /bin/bash -c 'for i in a b c d e f g; do touch /tests/users/{andrew,brian,carolyn,david,erin,frank}/downloads/$i.part; done'

RUN /bin/bash -c 'ln -s /tests/users /tests/users/frank/documents/rogue'
RUN ls -ahl /tests/users/andrew/downloads
RUN ls -ahl /tests/users/frank/documents

#RUN mkdir -p /tests/start
#RUN mkdir -p /tests/start/symlink_target_dir
#RUN touch /tests/start/unreadable.file && chmod 000 /tests/start/unreadable.file
#RUN mkdir -p /tests/start/unreadable_dir && chmod 000 /tests/start/unreadable_dir
#RUN /bin/bash -c 'for i in {1..5}; do touch /tests/start/readable_$i.file; done'
#
#RUN mkdir -p /tests/files/{1,2,3,4,5,6,7,8,9,10}/{a,b,c,d,e,f,g,h,i,j}
#RUN mkdir -p /tests/symlink/files
#RUN ln -s /tests/start/readable_1.file /tests/symlink/files/link_to_readable_1.file
#RUN mkdir -p /tests/symlink/dirs
#RUN ln -s /tests/start/symlink_target_dir /tests/symlink/dirs/link_to_readable_dir_1

RUN find /tests -type d

ENV GOPATH /app
CMD ["go", "test", "./..."]
