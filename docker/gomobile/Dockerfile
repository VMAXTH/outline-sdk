FROM ubuntu:22.04

ARG GO_VERSION=1.22.5

ENV DEBIAN_FRONTEND=noninteractive \
    ANDROID_HOME=/usr/local/lib/android/sdk \
    ANDROID_SDK_ROOT=/usr/local/lib/android/sdk \
    ANDROID_NDK_HOME=/usr/local/lib/android/sdk/ndk/26.1.10909125 \
    GOROOT=/usr/local/go \
    GOPATH=/go \
    PATH=/usr/local/go/bin:/usr/local/lib/android/sdk/cmdline-tools/latest/bin:/usr/local/lib/android/sdk/platform-tools:/go/bin:$PATH

# Базовые утилиты + JDK
RUN apt-get update && apt-get install -y --no-install-recommends \
      ca-certificates curl unzip git openjdk-17-jdk \
    && rm -rf /var/lib/apt/lists/*

# Go
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/go.tgz \
    && tar -C /usr/local -xzf /tmp/go.tgz \
    && rm /tmp/go.tgz \
    && mkdir -p /go

# Android cmdline-tools (с резервными версиями, чтобы не ловить 404)
RUN mkdir -p ${ANDROID_HOME}/cmdline-tools/latest \
    && cd /tmp \
    && for u in \
         https://dl.google.com/android/repository/commandlinetools-linux-11076702_latest.zip \
         https://dl.google.com/android/repository/commandlinetools-linux-11079508_latest.zip \
         https://dl.google.com/android/repository/commandlinetools-linux-11085808_latest.zip ; do \
         echo "try $u"; curl -fsSL "$u" -o cmdtools.zip && break || true; \
       done \
    && unzip -q cmdtools.zip -d ${ANDROID_HOME}/cmdline-tools/latest \
    && rm -f cmdtools.zip

# Лицензии и пакеты SDK/NDK
RUN yes | sdkmanager --licenses || true \
 && yes | sdkmanager \
      "platform-tools" \
      "platforms;android-34" \
      "build-tools;34.0.0" \
      "ndk;26.1.10909125"

# gomobile/gobind и инициализация (БЕЗ -ndk — берётся из ANDROID_NDK_HOME)
RUN /usr/local/go/bin/go install golang.org/x/mobile/cmd/gomobile@latest \
 && /usr/local/go/bin/go install golang.org/x/mobile/cmd/gobind@latest \
 && ANDROID_NDK_HOME=${ANDROID_NDK_HOME} /go/bin/gomobile init

WORKDIR /work
