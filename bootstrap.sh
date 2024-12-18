#!/usr/bin/bash

# Copyright (c) 2024, Rafael Santiago
# All rights reserved.
#
# This source code is licensed under the GPLv2 license found in the
# COPYING.GPLv2 file in the root directory of Eutherpe's source tree.

EUTHERPE_USER=eutherpe
EUTHERPE_PASSWD=eutherpe
SHOULD_SETUP_ETH_RESCUE_IFACE=0
EUTHERPE_DEFAULT_PORT=8080
SHOULD_BUILD_AND_INSTALL_BLUEZ_ALSA=1
SHOULD_REBOOT=1
EUTHERPE_WIFI_ESSID=""
EUTHERPE_WIFI_PASSPHRASE=""

has_internet_conectivity() {
    result=0
    ping -4 google.com -c 3 >/dev/null 2>&1
    if [[ $? == 0 ]] ; then
        result=1
    fi
    echo $result
}

is_active() {
    service=$1
    is=0
    if [[ $(systemctl is-active $service | grep ^active | wc -l) == 1 ]] ; then
        is=1
    fi
    echo $is
}

bootstrap_banner() {
    read -d '' data << "EOF"
#########################
\ \ \ \ \ \ \ ,|_|,   ,|_|,
\ \ \ \ \ \ \ |===|   |===|
\ \ \ \ \ \ \ |   |   |   |
\ \ \ \ \ \ \ /  &|   |&  \\
\ \ \_.-'`  , )* *( ,  `'-._ [ Eutherpe's Bootstrap ]
\ \ \ `"""""`"`   `"`"""""`
#########################
EOF
    echo -e "$data" >&2
    echo >&2
    echo "Hi there! I am the Eutherpe's bootstrap! What I am intending to do: " >&2
    echo >&2
    echo "- Create an user \"eutherpe\";" >&2
    echo "- Add it to sudo's group;" >&2
    echo "- Install some system dependencies required to you play your beloved tunes;" >&2
    echo "- Install Golang to actually build Eutherpe's app;" >&2
    echo "- Install kernel headers to make easy any specific system tune that you may want to do;" >&2
    echo "- Create the default's USB mount point in /media/USB;" >&2
    echo "- Build bluez-alsa from scratch and install it;" >&2
    echo "- Build up Eutherpe's app;" >&2
    echo "- Install whole Eutherpe's subsystem;" >&2
    echo "- Finally, reboot your system to finish applying all changes;" >&2
    echo >&2
}

areUroot() {
    yeah=1
    if [[ $USER != "root" ]] ; then
        echo "error: You must be root to bootstrap Eutherpe." >&2
        yeah=0
    fi
    echo $yeah
}

has_eutherpe_user() {
    echo $(id $EUTHERPE_USER 2>/dev/null | wc -l)
}

add_eutherpe_user() {
    useradd $EUTHERPE_USER >/dev/null 2>&1
    passwd $EUTHERPE_USER >/dev/null 2>&1 << EOF
$EUTHERPE_PASSWD
$EUTHERPE_PASSWD
EOF
    mkdir -p /home/$EUTHERPE_USER
    chown $EUTHERPE_USER:$EUTHERPE_USER /home/$EUTHERPE_USER >/dev/null 2>&1
    usermod -aG audio $EUTHERPE_USER
    echo $?
}

add_eutherpe_user_to_sudo_group() {
    usermod -aG sudo $EUTHERPE_USER
    echo $?
}

grant_eutherpe_user_nopasswd_privileges() {
    echo -e "$(cat /etc/sudoers | grep -v ^eutherpe)" > /etc/sudoers
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which ip)">>/etc/sudoers &&\
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which wpa_supplicant)">>/etc/sudoers &&\
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which dhclient)">>/etc/sudoers &&\
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which shutdown)">>/etc/sudoers &&\
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which systemctl) start bluealsa">>/etc/sudoers &&\
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which systemctl) stop bluealsa">>/etc/sudoers &&\
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which systemctl) restart bluealsa">>/etc/sudoers &&\
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which systemctl) start wpa_supplicant">>/etc/sudoers &&\
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which systemctl) stop wpa_supplicant">>/etc/sudoers &&\
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which systemctl) restart wpa_supplicant">>/etc/sudoers
    echo $?
}

is_sysdep_installed() {
    echo "$(dpkg -s $1 2>/dev/null | grep "Status: install" | wc -l)"
}

install_sysdeps() {
    exit_code=0
    apt-get update -y >/dev/null 2>&1
    for curr_dep in $(cat sysdeps.txt)
    do
        if [[ `is_sysdep_installed $curr_dep` == 0 ]] ; then
            echo "+-- installing $curr_dep..." >&2
            apt-get install $curr_dep -y >/dev/null 2>&1
            if [[ $? != 0 ]] ; then
                exit_code=$?
                break
            fi
            echo "*-- $curr_dep installed." >&2
        else
            echo "*-- $curr_dep already installed." >&2
        fi
    done
    echo $exit_code
}

get_arch() {
    arch_tag=$(uname -m)
    if [[ $(echo $arch_tag | grep x86_64 | wc -l)  == 1 ]] ; then
        arch_tag="amd64"
    elif [[ $(echo $arch_tag | grep ^aarch64 | wc -l) == 1 ]] ; then
        arch_tag="arm64"
    elif [[ $(echo $arch_tag | grep i686 | wc -l) == 1 ]] ; then
        arch_tag="386"
    else
        arch_tag=""
    fi
    echo $arch_tag
}

is_golang_installed() {
    if [[ -f /etc/profile.d/goenv.sh ]]; then
        source /etc/profile.d/goenv.sh
    fi
    go version >/dev/null 2>&1
    [[ $? == 0 ]] && echo "1" || echo "0"
}

install_golang() {
    arch_tag=`get_arch`
    if [[ -z $arch_tag ]] ; then
        echo "error: Your architecture $(uname -m) is not supported." >&2
        exit 1
    fi
    filename="go1.23.2.linux-"$arch_tag".tar.gz"
    download_uri="https://go.dev/dl"
    old_dir=$(pwd)
    cd /tmp
    wget $download_uri/$filename
    if [[ ! -f $filename ]] ; then
        cd $old_dir >/dev/null 2>&1
        echo "error: Golang's download has failed." >&2
        exit 1
    fi
    echo "+-- unpacking..." >&2
    tar -xvzf $filename -C /usr/local >/dev/null 2>&1
    exit_code=$?
    if [[ $exit_code == 0 ]] ; then
        echo "*-- unpacked." >&2
    else
        echo "x-- error during unpacking." >&2
    fi
    rm -f $filename >/dev/null 2>&1
    cd $old_dir >/dev/null 2>&1
    echo $exit_code
}

setup_goenv() {
    goenv_sh="/etc/profile.d/goenv.sh"
    echo -e "#!/usr/bin/bash\n\nexport GOROOT=/usr/local/go\nexport GOPATH=\$HOME/go\nexport PATH=\$GOPATH/bin:\$GOROOT/bin:\$PATH\n" > $goenv_sh
    chmod ugo+x $goenv_sh >/dev/null 2>&1
    echo $?
}

install_kernel_headers() {
    echo "+-- installing kernel headers..." >&2
    apt-get install linux-headers-$(uname -r) -y >/dev/null 2>&1
    echo "*-- linux-headers-$(uname -r) installed." >&2
    apt-get install gcc -y >/dev/null 2>&1
    echo "*-- gcc installed." >&2
    apt-get install make -y >/dev/null 2>&1
    echo "*-- make installed." >&2
    apt-get install perl -y >/dev/null 2>&1
    echo "*-- perl installed." >&2
    echo $?
}

create_usb_sto_mount_point() {
    mkdir -p /media/USB
    chmod ugo+rw /media/USB
    echo $?
}

build_eutherpe() {
    cd src
    go build -buildvcs=false >/dev/null 2>&1
    echo $?
}

setup_eutherpe_wifi_credentials() {
    player_cache_data=`cat /etc/eutherpe/player.cache | sed 's/"ESSID":".*",//'`
    patched_player_cache_data=`echo $player_cache_data | sed 's/^{/{"ESSID":"'$EUTHERPE_WIFI_ESSID'",/'`
    echo $patched_player_cache_data >/etc/eutherpe/player.cache
    psk=`wpa_passphrase $EUTHERPE_WIFI_ESSID $EUTHERPE_WIFI_PASSPHRASE | grep psk | grep -v '#psk'`
    EUTHERPE_WIFI_PASSPHRASE=""
    wpa_supplicant_conf="ctrl_interface=/run/wpa_supplicant
fast_reauth=1
#ap_scan=1
#update_config=1
#country=BR
network={
        scan_ssid=0
        proto=WPA
        key_mgmt=WPA-PSK
        auth_alg=OPEN
        ssid=\"$EUTHERPE_WIFI_ESSID\"
$psk
}"
    mkdir -p /etc/wpa_supplicant >/dev/null 2>&1
    echo "$wpa_supplicant_conf" >/etc/wpa_supplicant/wpa_supplicant.conf
    psk=""
    echo 0
}

install_eutherpe() {
    mkdir -p /etc/eutherpe >/dev/null 2>&1
    if [[ ! -f /etc/eutherpe/player.cache ]] ; then
        echo "{\"HostName\":\"eutherpe.local\"}" > /etc/eutherpe/player.cache
        chmod 755 /etc/eutherpe/player.cache
    fi
    if [[ ! -z $EUTHERPE_WIFI_ESSID ]] ; then
        if [[ `setup_eutherpe_wifi_credentials` != 0 ]] ; then
            echo "error: Unable to setup Wi-Fi credentials." >&2
            return 1
        fi
    fi
    chown -R $EUTHERPE_USER:$EUTHERPE_USER /etc/eutherpe >/dev/null 2>&1
    cp -rf src/web /etc/eutherpe/web >/dev/null 2>&1
    chmod -R 755 /etc/eutherpe
    cp src/eutherpe /usr/local/bin/ >/dev/null 2>&1
    cp src/usr/sbin/* /usr/sbin >/dev/null 2>&1
    chmod 777 /usr/sbin/eutherpe-usb-watchdog.sh
    chmod 777 /usr/sbin/run-eutherpe.sh
    cp src/etc/systemd/system/*.service /etc/systemd/system/ >/dev/null 2>&1
    if [[ $EUTHERPE_DEFAULT_PORT != 8080 ]]; then
        sed -i "s/^eutherpe.*$/eutherpe --listen-port=$EUTHERPE_DEFAULT_PORT/g" /usr/sbin/run-eutherpe.sh >/dev/null 2>&1
    fi
    systemctl start eutherpe-usb-watchdog eutherpe >/dev/null 2>&1
    systemctl enable eutherpe-usb-watchdog >/dev/null 2>&1
    systemctl enable eutherpe >/dev/null 2>&1
    echo $?
}

build_and_install_bluez_alsa() {
    echo "+-- cloning santiago's bluez-alsa fork..." >&2
    rm -rf bluez-alsa >&2
    git clone https://github.com/rafael-santiago/bluez-alsa -b v4.2.0 bluez-alsa >/dev/null 2>&1
    if [[ $? != 0 ]] ; then
        echo "error: while cloning bluez-alsa." >&2
        return 1
    fi
    echo "*-- cloned." >&2
    echo "+-- generating configure script..." >&2
    cd bluez-alsa >/dev/null 2>&1
    autoreconf --install >/dev/null 2>&1
    if [[ $? != 0 ]] ; then
        echo "error: while generating configure script." >&2
        return 1
    fi
    echo "*-- generated." >&2
    echo "+-- configuring build..." >&2
    mkdir build >/dev/null 2>&1
    cd build >/dev/null 2>&1
    ../configure --enable-mp3lame --enable-mpg123 --enable-rfcomm --enable-hcitop --enable-systemd --with-bluealsauser=$EUTHERPE_USER >/dev/null 2>&1
    if [[ $? != 0 ]] ; then
        echo "error: while configuring build." >&2
        return 1
    fi
    echo "*-- build configured." >&2
    echo "+-- building bluez-alsa..." >&2
    make >/dev/null 2>&1
    if [[ $? != 0 ]] ; then
        echo "error: while building bluez-alsa." >&2
        return 1
    fi
    echo "*-- built." >&2
    echo "+-- installing bluez-alsa..." >&2
    make install >/dev/null 2>&1
    if [[ $? != 0 ]] ; then
        echo "error: while installing bluez-alsa." >&2
        return 1
    fi
    echo "*-- installed." >&2
    echo "+-- configuring /var/lib directory and its permissions..." >&2
    mkdir /var/lib/bluealsa >/dev/null 2>&1
    chown $EUTHERPE_USER /var/lib/bluealsa >/dev/null 2>&1
    chmod 0700 /var/lib/bluealsa >/dev/null 2>&1
    if [[ $? != 0 ]] ; then
        echo "error: while configuring /var/lib directory." >&2
        return 1
    fi
    echo "*-- configured." >&2
    echo "+-- creating /etc/dbus-1/system.dbus/bluealsa.conf police file..." >&2
    mkdir -p /etc/dbus-1/system.dbus >/dev/null 2>&1
    cd ../.. >/dev/null 2>&1
    cp src/etc/dbus-1/system.dbus/bluealsa.conf /etc/dbus-1/system.dbus/bluealsa.conf >/dev/null 2>&1
    if [[ $? != 0 ]] ; then
        echo "error: while creating /etc/dbus-1/system.dbus/bluealsa.conf policy file." >&2
        return 1
    fi
    chmod ugo+rw /etc/dbus-1/system.dbus/bluealsa.conf >/dev/null 2>&1
    echo "*-- policy file created." >&2
    return 0
}

deactivate_avahi_daemon() {
    systemctl stop avahi-daemon >/dev/null 2>&1
    systemctl disable avahi-daemon >/dev/null 2>&1
    systemctl mask avahi-daemon >/dev/null 2>&1
    return 0
}

is_rescue_iface_set_already() {
    echo `cat /etc/network/interfaces | grep 'address 42.42.42.' | wc -l`
}

write_eth_config() {
    eth_iface=$1
    up_cmd=$2
    ip_addr=$3
    nt_mask=$4
    dname=$5
    cat /etc/network/interfaces > /etc/network/interfaces.stage
    echo "auto $eth_iface" >> /etc/network/interfaces.stage
    echo "iface $eth_iface inet static" >> /etc/network/interfaces.stage
    if [[ ! -z $up_cmd ]] ; then
        echo " up $up_cmd" >> /etc/network/interfaces.stage
    fi
    echo " address $ip_addr" >> /etc/network/interfaces.stage
    echo " netmask $nt_mask" >> /etc/network/interfaces.stage
    echo " dns-domain $dname" >> /etc/network/interfaces.stage
    mv /etc/network/interfaces.stage /etc/network/interfaces
}

get_eth_iface_pattern() {
    echo `ip link show | grep ': \(eth\|end\)' | tail -1 | awk '{ print $2 }' | sed 's/://g' | sed 's/[0-9]\+//g'`
}

setup_eth_rescue_iface() {
    exit_code=1
    eth_iface_pattern=`get_eth_iface_pattern`
    if [[ -z $eth_iface_pattern ]] ; then
        echo "error: Unable to find out a ethernet interface. It only supports 'ethN' or 'endN' interfaces." >&2
    elif [[ `is_rescue_iface_set_already` == 0 ]] ; then
        cp /etc/network/interfaces /etc/network/interfaces.bkp >/dev/null 2>&1
        `write_eth_config ${eth_iface_pattern}0 'ip route del 42.42.42.0/24 dev eth1' 42.42.42.1 255.255.255.0 euther-pi0.rescue`
        `write_eth_config ${eth_iface_pattern}1 'ip route del 42.42.42.0/24 dev eth0' 42.42.42.2 255.255.255.0 euther-pi1.rescue`
        exit_code=0
    else
        echo "=== bootstrap info: The ethernet rescue interfaces are already set." >&2
        exit_code=0
    fi
    echo $exit_code
}

set_wpa_supplicant_access_rights() {
    touch /etc/wpa_supplicant/wpa_supplicant.conf >/dev/null 2>&1
    chmod ugo+rw /etc/wpa_supplicant/wpa_supplicant.conf >/dev/null 2>&1
    echo $?
}

is_valid_number() {
    echo `echo $1 | grep "[0-9]\{1,\}$" | wc -l`
}

is_valid_port() {
    is_valid=`is_valid_number $1`
    if [[ $is_valid == 1 ]] && ! ([ "$1" -gt "0" ] && [ "$1" -le "65535" ]) ; then
        is_valid=0
    fi
    echo $is_valid
}

get_wifi_credentials() {
    read -p "Type the name of your Wi-Fi: " -r ESSID
    if [[ $? != 0 ]] ; then
        echo "error: Aborted." >&2
        return 1
    fi
    if [[ -z $ESSID ]] ; then
        echo "error: Null ESSID." >&2
        return 1
    fi
    read -sp "Type the Wi-Fi passphrase: " -r PASSPHRASE
    echo >&2
    read -sp "Confirm the Wi-Fi passpharse: " -r pass_confirm
    echo >&2
    if [[ $PASSPHRASE != $pass_confirm ]] ; then
        echo "error: The passphrase confirmation does not match." >&2
        return 1
    fi
    pass_confirm=""
    echo "$ESSID $PASSPHRASE"
}


`bootstrap_banner`

echo "=== Checking on your Internet conectivity..."

if [[ `has_internet_conectivity` != 1 ]] ; then
    echo "error: you are not connected to the Internet." >&2
    exit 1
fi

echo -e "=== Nice, you are connected to the Internet.\n"

answer="i"
while [[ ! $answer =~ ^[yYnN]$ ]]
do
    read -p "Do you want to bootstrap your system? [y/n] " -n 1 -r answer
    if [[  $answer =~ ^[nN]$ ]]; then
        echo
        exit 1
    fi
    echo
done

answer="i"
while [[ ! $answer =~ ^[yYnN]$ ]]
do
    read -p "Do you want to change Eutherpe's default listen port (8080)? [y/n] " -n 1 -r answer
    if [[ $answer =~ ^[yY]$ ]]; then
        echo
        read -p "Type the new listen port number: " EUTHERPE_DEFAULT_PORT
        if [[ `is_valid_port $EUTHERPE_DEFAULT_PORT` == 0 ]]; then
            echo "error: $EUTHERPE_DEFAULT_PORT is not a valid port."
            exit 1
        fi
    else
        echo
    fi
done

if [[ $(echo `get_arch` | grep ^arm | wc -l) == 1 ]] ; then
    answer="i"
    while [[ ! $answer =~ ^[yYnN]$ ]]
    do
        read -p "Do you want to set up a rescue ethernet interface? [y/n] " -n 1 -r answer
        if [[ $answer =~ ^[yY]$ ]]; then
            echo
            SHOULD_SETUP_ETH_RESCUE_IFACE=1
            break
        fi
        echo
    done
fi

if [[ $(echo `get_arch` | grep ^arm | wc -l) == 1 ]] ; then
    answer="i"
    while [[ ! $answer =~ ^[yYnN]$ ]]
    do
        read -p "Do you want to set up the Wi-Fi? [y/n] " -n 1 -r answer
        if [[ $answer =~ ^[yY]$ ]]; then
            echo
            wlan_creds=`get_wifi_credentials`
            if [[ -z $wlan_creds ]] ; then
                exit 1
            fi
            EUTHERPE_WIFI_ESSID=`echo $wlan_creds | awk '{ print $1 }'`
            EUTHERPE_WIFI_PASSPHRASE=`echo $wlan_creds | awk '{ print $2 }'`
            wlan_creds=""
            break
        fi
        echo
    done
fi

if [[ `areUroot` == 0 ]] ; then
    exit 1
else
    echo "=== Okay, you are root user :) let's start..."
fi

if [[ `is_active pulseaudio` == 1 ]]; then
    echo "bootstrap warn: Eutherpe uses ALSA and bluez-ALSA, you need to uninstall or at least deactivate pulseaudio before proceeding." >&2
    echo "aborted."
    exit 1
elif [[ `is_active pipewire` == 1 ]]; then
    echo "bootstrap warn: Eutherpe uses ALSA and bluez-ALSA, you need to uninstall or at least deactivate pipewire before proceeding." >&2
    echo "aborted."
    exit 1
elif [[ `is_active wireplumber` == 1 ]]; then
    echo "boostrap warn: Eutherpe uses ALSA and bluez-ALSA, you need to uninstall or at least deactivate wireplumber before proceeding." >&2
    echo "aborted."
    exit 1
fi

if [[ `is_active eutherpe` == 1 ]]; then
    echo "bootstrap info: An instance of eutherpe.service is running, let's stop it..."
    systemctl stop eutherpe
    if [[ `is_active eutherpe` == 1 ]]; then
        echo "error: Unable to stop eutherpe.service" >&2
        exit 1
    fi
    echo "bootstrap info: eutherpe.service was stopped."
fi

if [[ `is_active eutherpe-usb-watchdog` == 1 ]]; then
    echo "bootstrap info: An instance of eutherpe-usb-watchdog.service is running, let's stop it..."
    systemctl stop eutherpe-usb-watchdog
    if [[ `is_active eutherpe-usb-watchdog` == 1 ]]; then
        echo "error: Unable to stop eutherpe-usb-watchdog.service" >&2
        exit 1
    fi
    echo "bootstrap info: eutherpe-usb-watchdog.service was stopped."
fi

if [[ `is_active bluealsa` == 1 ]]; then
    SHOULD_BUILD_AND_INSTALL_BLUEZ_ALSA=0
    SHOULD_REBOOT=0
fi

if [[ `has_eutherpe_user` == 1 ]] ; then
    echo "=== Nice, $EUTHERPE_USER user already exists."
elif [[ `add_eutherpe_user` == 0 ]] ; then
    echo "=== The $EUTHERPE_USER user was added."
    SHOULD_REBOOT=1
fi

echo "=== bootstrap info: Adding $EUTHERPE_USER to sudo group..."

if [[ `add_eutherpe_user_to_sudo_group` != 0 ]] ; then
    echo "error: Unable to add user $EUTHERPE_USER to sudo group." >&2
    exit 1
fi

echo "=== bootstrap info: Done."
echo "=== bootstrap info: Installing system dependencies..."

if [[ `install_sysdeps` != 0 ]] ; then
    echo "error: Unable to install system dependencies." >&2
    exit 1
fi

echo "=== bootstrap info: Done."

echo "=== bootstrap info: granting $EUTHERPE_USER some nopasswd privileges..."

if [[ `grant_eutherpe_user_nopasswd_privileges` != 0 ]] ; then
    echo "error: Unable to grant nopasswd privileges to $EUTHERPE_USER." >&2
    exit 1
fi

echo "=== bootstrap info: Done."
echo "=== bootstrap info: Setting up /etc/wpa_supplicant/wpa_supplicant.conf access rights..."

if [[ `set_wpa_supplicant_access_rights` != 0 ]] ; then
    echo "error: Unable to set /etc/wpa_supplicant/wpa_supplicant.conf access rights." >&2
    exit 1
fi

echo "=== bootstrap info: Done."
echo "=== bootstrap info: Installing golang..."

if [[ `is_golang_installed` == 0 && `install_golang` != 0 ]] ; then
    echo "error: Unable to install golang." >&2
    exit 1
fi

echo "=== bootstrap info: Done."
echo "=== bootstrap info: Setting up golang environment..."

if [[ `setup_goenv` != 0 ]] ; then
    echo "error: Unable to set up golang environment." >&2
    exit 1
fi

source /etc/profile.d/goenv.sh

echo "=== bootstrap info: Done."
echo "=== bootstrap info: Installing kernel headers..."

if [[ `install_kernel_headers` != 0 ]] ; then
    echo "error: Unable to install kernel headers." >&2
    exit 1
fi

echo "=== bootstrap info: Done."
echo "=== bootstrap info: Creating USB storage mount point..."

if [[ `create_usb_sto_mount_point` != 0 ]] ; then
    echo "error: Unable to create USB storage mount point." >&2
    exit 1
fi

echo "=== bootstrap info: Done."

if [[ $SHOULD_SETUP_ETH_RESCUE_IFACE == 1 ]] ; then
    echo "=== bootstrap info: Setting up ethernet rescue interface 42.42.42.x..."
    if [[ `setup_eth_rescue_iface` != 0 ]] ; then
        echo "error: Unable to setup ethernet rescue interface." >&2
        exit 1
    fi
    echo "=== bootstrap info: Done."
fi

if [[ $SHOULD_BUILD_AND_INSTALL_BLUEZ_ALSA == 1 ]]; then
    echo "=== bootstrap info: Building and installing bluez-alsa..."
    `build_and_install_bluez_alsa`
    if [[ $? != 0 ]] ; then
        echo "error: Unable to install bluez-alsa." >&2
        exit 1
    fi
    echo "=== bootstrap info: Done."
else
    echo "=== bootstrap info: bluez-alsa is already installed."
fi
echo "=== bootstrap info: Now building Eutherpe..."

if [[ `build_eutherpe` != 0 ]] ; then
    echo "error: Unable to build Eutherpe." >&2
    exit 1
fi

echo "=== bootstrap info: Done."

if [[ `is_active avahi-daemon` != 0 ]] ; then
    echo "=== bootstrap info: Deactivating avahi-daemon..."
    `deactivate_avahi_daemon`
    echo "=== bootstrap info: Done."
else
    echo "=== bootstrap info: Nice, avahi-daemon is already deactivated."
fi

echo "=== bootstrap info: Now installing Eutherpe..."

if [[ `install_eutherpe` != 0 ]] ; then
    echo "error: Unable to install Eutherpe." >&2
    exit 1
fi

if [[ `is_active eutherpe` == 0 ]] ; then
    echo "error: eutherpe.service seems not to be active." >&2
    exit 1
fi

echo "=== bootstrap info: Nice, eutherpe.service is running."

if [[ `is_active eutherpe-usb-watchdog` == 0 ]] ; then
    echo "error: eutherpe-usb-watchdog.service seems not to be active." >&2
    exit 1
fi

echo "=== bootstrap info: Nice, eutherpe-usb-watchdog.service is running."
if [[ $SHOULD_REBOOT == 1 ]]; then
    echo "=== bootstrap info: Done. Reboot within 3 seconds..."
    sleep 3 && shutdown -r now
else
    echo "=== bootstrap info: Done."
fi
