#!/usr/bin/bash

# INFO(Rafael): Running pulseaudio as root is a source of problems, headaches and
#               poor sound quality. This is the way I have found to avoid them.

su - eutherpe <<EOF
eutherpe
/usr/local/bin/eutherpe
EOF
