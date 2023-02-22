[supervisord]
nodaemon=true

[program:omunganet-api]
command=/app/air -c .air.toml
directory=/app
autostart=true
autorestart=true
stdout_logfile=/app/omungaNet-api.log
stderr_logfile=/app/omungaNet-api.err.log
user=root
stopsignal=INT
stopasgroup=true
killasgroup=true
priority=1

[program:omunganet-redis]
command=redis-server
directory=/app
autostart=true
autorestart=true
user=root
priority=2
