FROM registry.access.redhat.com/ubi8-minimal

#RUN microdnf install /usr/sbin/ip

#CMD exec /bin/bash -c "trap : TERM INT; sleep infinity; wait"
CMD exec echo "FOO, BAR!"
