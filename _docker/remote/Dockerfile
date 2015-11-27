FROM centos
MAINTAINER hondata <takahiro_honda@unirita.co.jp>

RUN yum install -y passwd openssh-server
RUN sed -ri 's/UsePAM yes/#UsePAM yes/g' /etc/ssh/sshd_config
RUN sed -ri 's/#UsePAM no/UsePAM no/g' /etc/ssh/sshd_config
RUN sed -i -e 's/^\(session.*pam_loginuid.so\)/#\1/g' /etc/pam.d/sshd
RUN mkdir -p /var/run/sshd && chmod 755 /var/run/sshd
RUN /usr/bin/ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key -C '' -N ''
RUN /usr/bin/ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -C '' -N ''
RUN /usr/bin/ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key -C '' -N ''
RUN useradd passuser
RUN echo -e "testpass\ntestpass" | passwd passuser
ADD scripts/test.sh /home/passuser/test.sh
RUN chmod 755 /home/passuser/test.sh
RUN useradd keyuser
RUN echo -e "testpass2\ntestpass2" | passwd keyuser
RUN mkdir -p /home/keyuser/tmp
RUN chmod 777 /home/keyuser/tmp
RUN mkdir -p /home/keyuser/.ssh
ADD keys/id_rsa.pub /home/keyuser/.ssh/authorized_keys
RUN chown -R keyuser:keyuser /home/keyuser/.ssh
RUN chmod 400 /home/keyuser/.ssh/authorized_keys

EXPOSE 22

CMD /usr/sbin/sshd -D
