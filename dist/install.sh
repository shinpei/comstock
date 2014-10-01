set -e

COMSTOCK_GITHUB_RELEASE_URL=https://github.com/shinpei/comstock/releases/download
COMSTOCK_VERSION=0.1.8
COMSTOCK_ARCH=amd64

if [ `uname` = "Darwin" ]; then
    COMSTOCK_OS=darwin
elif [ `uname` = "Linux" ]; then
    COMSTOCK_OS=linux
else
    echo "Binary release are only available for Mac(Darwin) or Linux, contribute us for your distribution!"
    exit
fi

COMSTOCK_BINNAME=comstock_${COMSTOCK_VERSION}_${COMSTOCK_OS}_${COMSTOCK_ARCH}.zip


if [ ! -n "$COMSTOCK_DIR" ]; then
    COMSTOCK_DIR=~/comstock
fi

if [ -d "$COMSTOCK_DIR" ]; then
    echo "You already have comstock dir, you'll need to remove '$COMSTOCK_DIR' if you want to install"
    exit
fi

mkdir -p $COMSTOCK_DIR || {
    echo "cannot make $COMSTOCK_DIR, check permission"
    exit
}

## check unzip
hash unzip > /dev/null 2>&1 || {
    echo "You'll need 'unzip' for installing comstock"
    exit
}

## check downloader
hash curl > /dev/null 2>&1  && {
    COMSTOCK_DOWNLOADER="curl -L "
}

if [ -z "$COMSTOCK_DOWNLOADER" ]; then
    hash wget > /dev/null 2>&1 && {
	COMSTOCK_DOWNLOADER="wget --no-check-certificate -O "
    }
    if [ -z "$COMSTOCK_DOWNLOADER" ]; then
	echo "You'll need either 'curl' or 'wget' for downloading comstock"
	exit
    fi
fi

COMSTOCK_URL=$COMSTOCK_GITHUB_RELEASE_URL/$COMSTOCK_VERSION/$COMSTOCK_BINNAME
echo "fetching $COMSTOCK_URL ..."

$COMSTOCK_DOWNLOADER $COMSTOCK_URL -o $COMSTOCK_DIR/$COMSTOCK_BINNAME || {
    echo "Download failed somehow, please start again"
    exit
}

DESTDIR=/usr/local/bin

cd $COMSTOCK_DIR
unzip $COMSTOCK_DIR/$COMSTOCK_BINNAME
install -m 755 comstock-cli ${DESTDIR}
install -m 755 coms_save_previous ${DESTDIR}
install -m 755 comstock ${DESTDIR}
cd -

###
rm -rf $COMSTOCK_DIR
###
echo "comstock installed!"
