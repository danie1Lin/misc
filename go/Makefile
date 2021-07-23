TAR = $(shell ls -l |egrep '^d'| awk '{print $$9}'|egrep '[^.]+'|xargs echo) # using $$() will not be expand in include
include $(foreach a,$(TAR),$(a)/Makefile)

.PHONY: $(TAR)
$(TAR):
	@echo make $@/Makefile
	@$(MAKE) -C $@

$(foreach a,$(TAR),$(a).%):
	@echo make $* under $(subst .$*, ,$@) 
	@$(MAKE) -C $(subst .$*, ,$@) $*


# make Dirtory you want to run
# example: make log, it will run the make all under the log/
# debug with 
# 1. `echo 'print: ; @echo "$(TAR)"' | make -f Makefile -f - print`
# 2. make SHELL="/bin/bash -vx"
# 3, make --debug
# dry run with make -n
