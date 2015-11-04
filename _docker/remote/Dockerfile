FROM centos:centos6
MAINTAINER hondata <takahiro_honda@unirita.co.jp>

RUN yum install -y openssh-server
RUN sed -ri 's/UsePAM yes/UsePAM no/g' /etc/ssh/sshd_config
RUN mkdir -p /var/run/sshd && chmod 755 /var/run/sshd
RUN useradd passuser
RUN echo -e "testpass\ntestpass" | passwd passuser
RUN useradd keyuser
RUN echo -e "testpass2\ntestpass2" | passwd keyuser
RUN mkdir -p /home/keyuser/.ssh
ADD keys/id_rsa.pub /home/keyuser/.ssh/authorized_keys
RUN chown -R keyuser:keyuser /home/keyuser/.ssh
RUN chmod 400 /home/keyuser/.ssh/authorized_keys
RUN /etc/init.d/sshd start

EXPOSE 22

CMD /sbin/init
