# This docker file is purely for CI on Travis so we can ssh into it
FROM ubuntu:latest

# I know this shouldn't be root but seriously, this container has one purpose
# and it's short lived.
ADD id_rsa.pub /root/.ssh/authorized_keys

# SO INSECURE NEVER USE THIS IMAGE. SERIOUSLY I KNOW I'M A BAD PERSON.
RUN apt-get update && apt-get install -y openssh-server 

# ONE LAST TIME NEVER USE THIS IMAGE FOR THE LOVE OF RICHARD STALLMAN
RUN mkdir /var/run/sshd
CMD ["/usr/sbin/sshd", "-D"]
