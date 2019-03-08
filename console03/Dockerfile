# Simple Dockerfile to load up our go program
# in ubuntu and set it up as the running program.
FROM	ubuntu:latest
LABEL   maintainer="bob@2kranki.us"
RUN     mkdir -p /usr/local/bin
COPY	./console03_linux /usr/local/bin/
EXPOSE  9000/tcp
CMD	    "/usr/local/bin/console03_linux"
