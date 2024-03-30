#!/bin/bash
# This is a bash script used as an entrypoint for a Docker container.
set -e
# The 'set -e' command causes the shell to exit if any invoked command fails.



chown -R shortify:shortify /logs
# The 'chown -R shortify:shortify $LOG_DIR' command changes the owner and group of the LOG_DIR directory (and its subdirectories) to 'shortify'.

exec sudo -u shortify "$@"
# The 'exec sudo -u shortify "$@"' command replaces the current shell with a new one where the following command (passed as arguments) is executed as the 'shortify' user.