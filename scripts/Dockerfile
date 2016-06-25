# This docker file is purely for CI on Travis so we can ssh into it
FROM ubuntu:latest

# I know this shouldn't be root but seriously, this container has one purpose
# and it's short lived.
ADD "$HOME/.ssh/id_rsa.pub" /root/.ssh/authorized_keys
RUN "systemctl enable ssh"
