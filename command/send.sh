#!/bin/bash
# source ./send.sh

echo "Press [CTRL+C] to stop.."
while :
do
	echo "command-client.main.forloop.increment:1|c" | nc -u -w 1 127.0.0.1 8125
	sleep 1
done
