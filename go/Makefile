define USAGAE
make $${Directory you want to run.}
example:
	make log # it will run the make all under the log/
	make log.glog # it will equal to run the command `make glog` under log

debug with:
	1. inspect VAR: `echo 'print: ; @echo "$$(VAR)"' | make -f Makefile -f - print`
	2. bash debug: make SHELL="/bin/bash -vx"
	3. makefile debug: make --debug

dry run:
	make -n
endef

.PHONY: help
help:
	@$(info $(USAGAE)) # can not use echo $$(USAGAE), it will treat the multi line var as command
	@echo

subdirs = $(shell ls -l |egrep '^d'| awk '{print $$9}'|egrep '[^.]+'|xargs echo) # using $$() will not be expand in include
include $(foreach a,$(subdirs),$(a)/Makefile)

## make the first target of sub directory
.PHONY: $(subdirs)
$(subdirs):
	@echo make $@/Makefile
	@$(MAKE) -C $@

## S{Subdir}.${Make file target}
$(foreach a,$(subdirs),$(a).%):
	@echo make $* under $(subst .$*, ,$@) 
	@$(MAKE) -C $(subst .$*, ,$@) $*
