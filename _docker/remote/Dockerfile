FROM centos
MAINTAINER hondata <takahiro_honda@unirita.co.jp>

RUN yum install -y openssh-server
RUN sed -ri 's/UsePAM yes/UsePAM no/g' /etc/ssh/sshd_config
RUN mkdir -p /var/run/sshd && chmod 755 /var/run/sshd
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
RUN /etc/init.d/sshd start

EXPOSE 22

CMD /sbin/init
