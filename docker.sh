#!/bin/bash
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

docker run -it --rm -v "$(pwd):/twgeocode" -v /etc/passwd:/etc/passwd:ro -v /etc/group:/etc/group:ro -e "HOME=/tmp" --user "$(id -u):$(id -g)" --workdir /twgeocode golang /twgeocode/gen.sh
