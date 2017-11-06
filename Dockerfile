FROM alpine:3.6

RUN echo 'blurblurbot:x:1000:1000::/var/blurblurbot:/sbin/nologin' >> /etc/passwd \
    && echo 'blurblurbot:!::0:::::' >> /etc/shadow \
    && mkdir /var/blurblurbot \
&& chown blurblurbot /var/blurblurbot

ENV release=v0.1.2

RUN apk add --no-cache curl jq gnupg ca-certificates \
    && gpg --keyserver keyserver.ubuntu.com --recv-key F6037DD30859399B


RUN set -x \
    && mkdir /blurblurbot && cd /blurblurbot \
    && release=${release:-$(curl -s https://api.github.com/repos/dtchanpura/blurblurbot/releases/latest | jq -r .tag_name )} \
    && curl -sLO https://github.com/dtchanpura/blurblurbot/releases/download/${release}/blurblurbot-linux-amd64-${release}.tar.gz \
    && curl -sLO https://github.com/dtchanpura/blurblurbot/releases/download/${release}/sha256sum.txt.asc \
    && gpg2 --verify sha256sum.txt.asc \
    && grep blurblurbot-linux-amd64-${release} sha256sum.txt.asc | sha256sum \
    && tar -xf blurblurbot-linux-amd64-${release}.tar.gz \
    && rm -rf sha256sum.txt.asc blurblurbot-linux-amd64-${release}.tar.gz && apk del curl jq gnupg

USER blurblurbot

ENTRYPOINT ["/blurblurbot/blurblurbot"]
