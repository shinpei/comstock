#!/bin/bash
# Comstock - stock your command to the cloud
# comstock shell scripts
# shinpei(c)2014

saveBashHistory() {
    fc -ln -2 > hoge
    cat hoge | comstock save --shell bash
}

# entry point
if [ -z "$BASH_VERSION" ] || [ "${BASH##*/}" != "bash" ]; then 
    saveBashHistory
fi

