# docker/gomobile/Dockerfile
# Образ для сборки AAR через gomobile bind

FROM ubuntu:22.04

# ---------- Базовые пакеты ----------
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates curl unzip zip xz-utils \
    git build-essential make \
    openjdk-17-jdk \
    python3 \
    && rm -rf /var/lib/apt/lists/*

# ---------- Пути и переменные окружения ----------
# Android SDK
ENV ANDROID_SDK_ROOT=/usr/local/lib/android/sdk
ENV ANDROID_HOME=${ANDROID_SDK_ROOT}
# Конкретная версия NDK, которая ставится ниже
ENV ANDROID_NDK_VERSION=26.1.10909125
ENV ANDROID_NDK_HOME=${ANDROID_SDK_ROOT}/ndk/${ANDROID_NDK_VERSION}

# Go
ENV GOPATH=/go
ENV GOROOT=/usr/local/go
ENV GO_VERSION=1.22.10

# PATH
ENV PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin:${ANDROID_SDK_ROOT}/platform-tools:${ANDROID_SDK_ROOT}/cmdline-tools/latest/bin

# ---------- Установка Go ----------
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/go.tgz \
 && rm -rf /usr/local/go && tar -C /usr/local -xzf /tmp/go.tgz \
 && rm -f /tmp/go.tgz \
 && mkdir -p ${GOPATH}/pkg ${GOPATH}/bin

# ---------- Установка Android cmdline-tools ----------
RUN mkdir -p ${ANDROID_SDK_ROOT}/cmdline-tools \
 && curl -fsSL https://dl.google.com/android/repository/commandlinetools-linux-11076708_latest.zip -o /tmp/cmdtools.zip \
 && unzip -q /tmp/cmdtools.zip -d ${ANDROID_SDK_ROOT}/cmdline-tools \
 # архив раскладывается как cmdline-tools/cmdline-tools — переносим в expected "latest"
 && mv ${ANDROID_SDK_ROOT}/cmdline-tools/cmdline-tools ${ANDROID_SDK_ROOT}/cmdline-tools/latest \
 && rm -f /tmp/cmdtools.zip \
 # иногда полезно положить прямую ссылку на sdkmanager
 && ln -sf ${ANDROID_SDK_ROOT}/cmdline-tools/latest/bin/sdkmanager /usr/local/bin/sdkmanager

# ---------- Принятие лицензий и установка SDK/NDK пакетов ----------
RUN yes | ${ANDROID_SDK_ROOT}/cmdline-tools/latest/bin/sdkmanager --licenses || true \
 && yes | ${ANDROID_SDK_ROOT}/cmdline-tools/latest/bin/sdkmanager \
      "platform-tools" \
      "platforms;android-34" \
      "build-tools;34.0.0" \
      "ndk;${ANDROID_NDK_VERSION}"

# ---------- Установка gomobile/gobind и инициализация ----------
# При необходимости можно зафиксировать версии, сейчас берём актуальные
RUN go install golang.org/x/mobile/cmd/gomobile@latest \
 && go install golang.org/x/mobile/cmd/gobind@latest \
 # инициализация gomobile с указанием установленного NDK
 && gomobile init -ndk ${ANDROID_NDK_HOME}

# Рабочая директория для сборок (будет монтироваться из GH Actions)
WORKDIR /workspace
