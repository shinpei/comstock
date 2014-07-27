#!/bin/bash
# Comstock - stock your command to the cloud
# comstock shell scripts
# shinpei(c)2014

saveBashHistory() {
    fc -l -n | tail -n 2 | head -n 1 | sed -e 's/^[[:space:]]*//'  | comstock --shell bash save
}

# entry point
if [ -z "$BASH_VERSION" ] || [ "${BASH##*/}" != "bash" ]; then 
    echo "not bash"
else 
    saveBashHistory
fi

