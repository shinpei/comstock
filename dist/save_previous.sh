#!/bin/bash
# Comstock - stock your command to the cloud
# comstock shell scripts
# shinpei(c)2014

saveBashHistory() {
    fc -l -n 2 | tac |  comstock --shell bash save
}

# entry point
if [ -z "$BASH_VERSION" ] || [ "${BASH##*/}" != "bash" ]; then 
    saveBashHistory
fi

