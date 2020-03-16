#!/bin/sh
sslocal -s c32s1.jamjams.net -p 27176 -k 9YMkGGv4Fg -b 0.0.0.0 -l 1080 -m aes-256-gcm -d start
sleep 3

curl --connect-timeout 5 -m 5 -s -w "%{http_code}" "www.google.com" -o /dev/null --socks5 127.0.0.1:1080
killall sslocal
