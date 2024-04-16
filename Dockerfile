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
RUN /bin/bash -c 'mkdir -p /tests2/more/directories/to/find'
RUN /bin/bash -c 'mkdir -p /tests3/more'
RUN /bin/bash -c 'for i in {1..6}; do touch /tests3/doc-$i.pdf; done'
RUN /bin/bash -c 'for i in {1..6}; do touch /tests3/more/img-$i.jpg; done'
RUN /bin/bash -c 'touch /file.txt'
RUN /bin/bash -c 'ln -s /tests/users /tests/users/frank/documents/rogue'
RUN /bin/bash -c 'ln -s /tests/users/andrew /tests/users/david/documents/rogue'
RUN /bin/bash -c 'ln -s /tests2/ /tests/users/andrew/others'
RUN /bin/bash -c 'ln -s /file.txt /tests/users/andrew/linkedfile'
RUN /bin/bash -c 'ln -s /tests3 /tests/users/docs1'
RUN /bin/bash -c 'ln -s /tests3 /tests/users/docs2'
RUN /bin/bash -c 'ln -s /tests/users/docs1 /tests3/more/rogue'

RUN ls -ahl /tests/users/andrew/downloads
RUN ls -ahl /tests/users/frank/documents

RUN ls -ahl /tests/users

RUN find /tests -type d

ENV GOPATH /app
CMD ["go", "test", "./..."]
