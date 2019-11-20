#!/bin/sh

killall -9 "qmoon_server"

sleep 1

ATM_IPLimit=5 MailUser=admin@qoschain.io MailPassword=Qoschain@16 MailSmtpServer=smtp.mxhichina.com ATM_KEY=DhNuqyAErrY70gRQwHTBcOzw/BQG1QNMJYcvyY7wgohV9vsfRq0jxHEWzCRy57v6Zb7DAuL1bFmUkitqQ0LhkA== ATM_AMOUNT=1000000000 ./qmoon_server server  > access.log 2>&1 &

