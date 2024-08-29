#
# Copyright (c) 2024, Rafael Santiago
# All rights reserved.
#
# This source code is licensed under the GPLv2 license found in the
# COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
#

define do_update
	@if [ "$(USER)" != "root" ]; then\
	    echo error: You are not root.;\
	    exit 1;\
	fi
	@if [ ! -d /etc/eutherpe ]; then\
	    echo error: You have not ran bootstrap yet.;\
	    exit 1;\
	fi
	@if [ "$$(systemctl status eutherpe | grep 'active (' | wc -l)" = "1" ]; then\
	    systemctl stop eutherpe;\
	    while [ "$$(systemctl status eutherpe | grep 'inactive (' | wc -l)" != "1" ]; do\
		sleep 1s;\
	    done;\
	fi
	@cp eutherpe $$(which eutherpe) && \
	cp -rf web/ /etc/eutherpe/web/ && \
	systemctl start eutherpe && \
	echo info: Your Eutherpe\'s core was updated.
endef

eutherpe:
	@ if [ -f /etc/profile.d/goenv.sh ] ; then\
	    . /etc/profile.d/goenv.sh ;\
	    go build ;\
	else \
	    go build ;\
	fi

tests:
	@go clean -testcache
	@go test internal/... -v

bootstrap:
	@$(shell cd .. && ./bootstrap.sh)

update: eutherpe
	@$(call do_update)

.PHONY: bootstrap eutherpe tests update
