#!/usr/bin/bash

EUTHERPE_USER=eutherpe
EUTHERPE_PASSWD=eutherpe

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
    echo "- Build up Eutherpe's app;" >&2
    echo "- Install whole Eutherpe's package it;" >&2
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
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which dhclient)">>/etc/sudoers
    echo "$EUTHERPE_USER        ALL=(ALL:ALL)   NOPASSWD: $(which shutdown)">>/etc/sudoers
    echo $?
}

is_sysdep_installed() {
    echo "$(dpkg -s $1 2>/dev/null | grep "Status: install" | wc -l)"
}

install_sysdeps() {
    exit_code=0
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
    go version >/dev/null 2>&1
    [[ $? == 0 ]] && echo "1" || echo "0"
}

install_golang() {
    arch_tag=`get_arch`
    if [[ -z $arch_tag ]] ; then
        echo "error: Your architecture $(uname -m) is not supported." >&2
        exit 1
    fi
    filename="go1.19.linux-"$arch_tag".tar.gz"
    download_uri="https://go.dev/dl"
    old_dir=$(pwd)
    cd /tmp
    wget $download_uri/$filename
    if [[ ! -f $filename ]] ; then
        cd $old_dir >/dev/null 2>&1
        echo "error: Golang's download has failed." >&2
        exit 1
    fi
    echo "*-- unpacking..." >&2
    tar -xvzf $filename -C /usr/local >/dev/null 2>&1
    exit_code=$?
    if [[ $exit_code == 0 ]] ; then
        echo "+-- unpacked." >&2
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
    apt-get install linux-headers-$(uname -r) -y >/dev/null 2>&1
    echo "+-- linux-headers-$(uname -r) installed." >&2
    apt-get install gcc -y >/dev/null 2>&1
    echo "+-- gcc installed." >&2
    apt-get install make -y >/dev/null 2>&1
    echo "+-- make installed." >&2
    apt-get install perl -y >/dev/null 2>&1
    echo "+-- perl installed." >&2
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

install_eutherpe() {
    mkdir -p /etc/eutherpe >/dev/null 2>&1
    if [[ ! -f /etc/eutherpe/player.cache ]] ; then
        echo "{\"HostName\":\"eutherpe.local\"}" > /etc/eutherpe/player.cache
        chmod 755 /etc/eutherpe/player.cache
    fi
    chown -R $EUTHERPE_USER:$EUTHERPE_USER /etc/eutherpe >/dev/null 2>&1
    cp -rf src/web /etc/eutherpe/web >/dev/null 2>&1
    chmod -R 755 /etc/eutherpe
    cp src/eutherpe /usr/local/bin/ >/dev/null 2>&1
    cp src/usr/sbin/* /usr/sbin >/dev/null 2>&1
    chmod 777 /usr/sbin/eutherpe-usb-watchdog.sh
    chmod 777 /usr/sbin/run-eutherpe.sh
    cp src/etc/systemd/system/*.service /etc/systemd/system/ >/dev/null 2>&1
    systemctl start eutherpe-usb-watchdog eutherpe >/dev/null 2>&1
    systemctl enable eutherpe-usb-watchdog >/dev/null 2>&1
    systemctl enable eutherpe >/dev/null 2>&1
    echo $?
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

if [[ `areUroot` == 0 ]] ; then
    exit 1
else
    echo "=== Okay, you are root user :) let's start..."
fi

if [[ `has_eutherpe_user` == 1 ]] ; then
    echo "=== Nice, $EUTHERPE_USER user already exists."
elif [[ `add_eutherpe_user` == 0 ]] ; then
    echo "=== The $EUTHERPE_USER user was added."
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
echo "=== bootstrap info: Installing golang..."

if [[ `is_golang_installed` == 0 && `install_golang` != 0 ]] ; then
    echo "error: Unable to install golang. ($x)" >&2
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
echo "=== bootstrap info: Now building Eutherpe..."

if [[ `build_eutherpe` != 0 ]] ; then
    echo "error: Unable to build Eutherpe." >&2
    exit 1
fi

echo "=== bootstrap info: Done."
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

echo "=== bootstrap info: Done."

exit 0
