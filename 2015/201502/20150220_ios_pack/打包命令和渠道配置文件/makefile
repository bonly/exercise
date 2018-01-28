#K_DIR :=$(shell date -d now "+%Y%m%d_%H%M") 
K_DIR :=$(shell date  "+%Y%m%d_%H%M") 
F_DIR :=/home/share/1-驯龙英雄/游戏安装包/
R_DIR :=/home/share/1-驯龙英雄/游戏安装包/Last_Release
T_DIR :=/home/share/1-驯龙英雄/游戏安装包/Last_Test

CLEAR=\x1b[0m
GREEN=\x1b[32;01m
RED=\x1b[31;01m
YELLOW=\x1b[33;01m
BLUE=\x1b[34;01m
CYAN=\x1b[36;01m
BOLD=\033[1m
NORMAL=\033[0m

ifndef rev
rev:=Debug
endif

all:
	@echo  "$(BOLD)$(RED)build_test   $(GREEN)连测试服务器的版本"
	@echo  "$(BOLD)$(RED)build_release   $(GREEN)连生产服务器的版本"
	@echo  "$(BOLD)$(RED)ipa rev=Release|Debug  $(GREEN)连生产服务器的发布版本"

build_ipa:
	rm -rf ../bin/*
	rm -rf build.log
	rm -rf err.log
	@if [ -d $(dir) ]; then \
	echo "processing single dir $(dir) for $(rev) " $(IP); \
	$(MAKE) -C $(dir) ipa rev=$(rev) 2>>err.log |tee build.log; \
	else \
	for a in $$(ls); do \
		if [ -d $$a ]; then \
			echo "processing folder $$a for $(rev) " $(IP); \
			$(MAKE) -C $$a ipa rev=$(rev) 2>>err.log |tee build.log; \
		fi; \
	done; \
	fi;
	@echo "build Done!"

build_bdp:
	rm -rf ../bin/*
	rm -rf err.log
	@for a in $$(ls); do \
		if [ -d $$a ]; then \
			echo "processing folder $$a for release " ${PROD_IP}; \
			export PLATFORM_IP=${PROD_IP} \
			$(MAKE) -C $$a bdp 2>>err.log; \
			$(MAKE) -C $$a pkg>>err.log; \
		fi; \
	done;
	@echo "build Done!"

build_bdt:
	rm -rf ../bin/*
	rm -rf err.log
	@for a in $$(ls); do \
		if [ -d $$a ]; then \
			echo "processing folder $$a for test " ${TEST_IP}; \
			export PLATFORM_IP=${TEST_IP} \
			$(MAKE) -C $$a bdt 2>>err.log; \
			$(MAKE) -C $$a pkg>>err.log; \
		fi; \
	done;
	@echo "build Done!"

send_release:
	ssh 100 mkdir $(F_DIR)/R_$(K_DIR)
	touch ../bin/release_package_$(rev)
	scp ../bin/* 100:$(F_DIR)/R_$(K_DIR)
	ssh 100 rm -rf $(R_DIR)
	ssh 100 ln -s $(F_DIR)/R_$(K_DIR) $(R_DIR)

send_test:
	ssh 100 mkdir $(F_DIR)/T_$(K_DIR)
	touch ../bin/test_package
	scp ../bin/* 100:$(F_DIR)/T_$(K_DIR)
	ssh 100 rm -rf $(T_DIR)
	ssh 100 ln -s $(F_DIR)/T_$(K_DIR) $(T_DIR)

send_pkg:
	ssh 100 mkdir $(F_DIR)/$(rev)_$(K_DIR)
	touch ../bin/$(rev)_package
	scp ../bin/* 100:$(F_DIR)/$(rev)_$(K_DIR)
	ssh 100 rm -rf $(F_DIR)/Last_$(rev)
	ssh 100 ln -s $(F_DIR)/$(rev)_$(K_DIR) $(F_DIR)/Last_$(rev)

build_test: build_bdt send_test

build_release: build_bdp send_release

ipa: build_ipa 


test:
	@echo $(F_DIR)/T_$(K_DIR)


