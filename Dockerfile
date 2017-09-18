FROM golang
ADD public/system /www
ADD ilove /bin/
ADD admin_auth/views /admin_auth/views
CMD /bin/ilove
