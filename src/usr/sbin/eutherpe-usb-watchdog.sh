#!/bin/bash

# Copyright (c) 2024, Rafael Santiago
# All rights reserved.
#
# This source code is licensed under the GPLv2 license found in the
# COPYING.GPLv2 file in the root directory of Eutherpe's source tree.

get_option() {
    needle="--"$1"="
    data=$2
    for option in $*
    do
        if [[ `echo $option | grep ^$needle | wc -l` == 1 ]] ; then
            data=$(echo $option | sed 's/'$needle'//g')
            break
        fi
    done
    echo $data
}

try_mount() {
    mount_point=$1
    mounted_device=""
    for device in $(find /dev/ -maxdepth 1 -type b | grep .*[0-9]$)
    do
        if [[ $(udevadm info --name=$device | grep ID_BUS=usb | wc -l) > 0 ]] ; then
            mount $device $mount_point -o umask=000 >/dev/null 2>&1
            echo *** info: $device mounted at $mount_point >&2
            mounted_device=$device
            break
        fi
    done
    sleep 5s
    echo $mounted_device
}


wait4mount() {
    mount_point=$1
    while [ 1 ] ;
    do
        echo *** info: waiting for an usb... >&2
        device=`try_mount $mount_point`
        if [ ! -z "$device" ] ; then
            `wait4umount $device` >/dev/null 2>&1
        fi
    done
}

wait4umount() {
    device=$1
    serial_number=`get_device_serial_number $device`
    while [ 1 ] ;
    do
        if [[ $(usb-devices | grep "SerialNumber=$serial_number" | wc -l) == 0
                || $(cat /proc/mounts | grep $device | wc -l) == 0 ]] ; then
            echo *** info: $device was dismounted >&2
            break
        fi
        sleep 5s
    done
}

get_device_serial_number() {
    device=`echo $1 | sed 's/\/dev\///g'`
    serial_number=`udevadm info --name=$device | grep ID_SERIAL_SHORT | sed s/.*ID_SERIAL_SHORT=//g`
    echo $serial_number
}

if [[ $# != 1 ]] ; then
    echo "use: $0 --mount-point=<directory>"
    exit 1
fi

mount_point=`get_option mount-point "" $*`

if [[ -z $mount_point ]] ; then
    echo "error: null mount point."
    exit 1
fi

`wait4mount $mount_point`
