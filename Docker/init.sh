#!/bin/bash

# BSD 2-Clause License

# Copyright (c) 2020, Supreeth Herle
# All rights reserved.

# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions are met:

# 1. Redistributions of source code must retain the above copyright notice, this
#    list of conditions and the following disclaimer.

# 2. Redistributions in binary form must reproduce the above copyright notice,
#    this list of conditions and the following disclaimer in the documentation
#    and/or other materials provided with the distribution.

# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
# SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
# OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

cp /mnt/enb.yml examples/gw-tester/enb/enb.yml
cp /mnt/mme.yml examples/gw-tester/mme/mme.yml

sed -i 's|MCC|'$MCC'|g' examples/gw-tester/enb/enb.yml
sed -i 's|MNC|'$MNC'|g' examples/gw-tester/enb/enb.yml
sed -i 's|GTP_DOCKER_IP|'$GTP_DOCKER_IP'|g' examples/gw-tester/enb/enb.yml
sed -i 's|UE1_IMSI|'$UE1_IMSI'|g' examples/gw-tester/enb/enb.yml
sed -i 's|UE1_IMEISV|'$UE1_IMEISV'|g' examples/gw-tester/enb/enb.yml
sed -i 's|UE_IPV4_INTERNET|'$UE_IPV4_INTERNET'|g' examples/gw-tester/enb/enb.yml

sed -i 's|MCC|'$MCC'|g' /go-gtp/examples/gw-tester/mme/mme.yml
sed -i 's|MNC|'$MNC'|g' /go-gtp/examples/gw-tester/mme/mme.yml
sed -i 's|SGWC_IP|'$SGWC_IP'|g' /go-gtp/examples/gw-tester/mme/mme.yml
sed -i 's|SMF_IP|'$SMF_IP'|g' /go-gtp/examples/gw-tester/mme/mme.yml
sed -i 's|GTP_DOCKER_IP|'$GTP_DOCKER_IP'|g' /go-gtp/examples/gw-tester/mme/mme.yml

cd /go-gtp/examples/gw-tester/mme/
./mme &  

cd /go-gtp/examples/gw-tester/enb/
./enb

# Sync docker time
#ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
