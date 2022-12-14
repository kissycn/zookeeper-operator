FROM openjdk:11-jre-slim

ENV DT_ROOT_DIR=/opt/dtweave \
    ZOO_BASE_DIR=/opt/dtweave/zookeeper \
    ZOO_CONF_DIR=/opt/dtweave/zookeeper/conf

ARG GPG_KEY=DFF24FB8323ADAC90E3CF36F729E61230EA917E9
ARG SHORT_DISTRO_NAME=zookeeper-3.7.1
ARG DISTRO_NAME=apache-zookeeper-3.7.1-bin

# Add a user with an explicit UID/GID and create necessary directories
RUN set -eux; \
    groupadd -r zookeeper --gid=1000; \
    useradd -r -g zookeeper --uid=1000 zookeeper; \
    mkdir -p "$ZOO_BASE_DIR" "$ZOO_CONF_DIR"; \
    chown zookeeper:zookeeper "$DT_ROOT_DIR" "$ZOO_BASE_DIR" "$ZOO_CONF_DIR"

# Install required packges
RUN set -eux; \
    apt-get update; \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        dirmngr \
        gosu \
        gnupg \
        netcat \
        wget; \
    rm -rf /var/lib/apt/lists/*; \
    # Verify that gosu binary works
    gosu nobody true \

# Download Apache Zookeeper, verify its PGP signature, untar and clean up
RUN set -eux; \
    ddist() { \
        local f="$1"; shift; \
        local distFile="$1"; shift; \
        local success=; \
        local distUrl=; \
        for distUrl in \
            https://www.apache.org/dist/ \
            https://archive.apache.org/dist/ \
            'https://www.apache.org/dyn/closer.cgi?action=download&filename=' \
        ; do \
            if wget -q -O "$f" "$distUrl$distFile" && [ -s "$f" ]; then \
                success=1; \
                break; \
            fi; \
        done; \
        [ -n "$success" ]; \
    }; \
    ddist "$DISTRO_NAME.tar.gz" "zookeeper/$SHORT_DISTRO_NAME/$DISTRO_NAME.tar.gz"; \
    ddist "$DISTRO_NAME.tar.gz.asc" "zookeeper/$SHORT_DISTRO_NAME/$DISTRO_NAME.tar.gz.asc"; \
    export GNUPGHOME="$(mktemp -d)"; \
    gpg --keyserver hkps://keyserver.pgp.com --recv-key "$GPG_KEY" || \
    gpg --keyserver hkps://keyserver.ubuntu.com --recv-keys "$GPG_KEY" || \
    gpg --keyserver hkps://pgp.mit.edu --recv-keys "$GPG_KEY"; \
    gpg --batch --verify "$DISTRO_NAME.tar.gz.asc" "$DISTRO_NAME.tar.gz"; \
    tar -zxf "$DISTRO_NAME.tar.gz"; \
    mv "$DISTRO_NAME/"* "$ZOO_BASE_DIR"; \
    rm -rf "$GNUPGHOME" "$DISTRO_NAME.tar.gz" "$DISTRO_NAME.tar.gz.asc";

COPY prebuild /
RUN chown -R zookeeper:zookeeper /opt/dtweave/; \
    chmod -R +x /opt/dtweave/scripts/

WORKDIR $ZOO_BASE_DIR
EXPOSE 2181 2888 3888 8080

ENV PATH=$PATH:$ZOO_BASE_DIR/bin \
    ZOOCFGDIR=$ZOO_CONF_DIR

USER 1000
ENTRYPOINT ["/opt/dtweave/scripts/entrypoint.sh"]
CMD ["zkServer.sh", "start-foreground"]