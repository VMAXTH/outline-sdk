FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV ANDROID_SDK_ROOT=/opt/android-sdk
ENV ANDROID_HOME=/opt/android-sdk
ENV PATH=$PATH:/usr/local/go/bin:/go/bin:$ANDROID_SDK_ROOT/cmdline-tools/latest/bin:$ANDROID_SDK_ROOT/platform-tools

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl unzip git ca-certificates build-essential openjdk-17-jdk \
 && rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://go.dev/dl/go1.22.7.linux-amd64.tar.gz | tar -C /usr/local -xz

RUN mkdir -p $ANDROID_SDK_ROOT/cmdline-tools/latest \
 && curl -fsSL https://dl.google.com/android/repository/commandlinetools-linux-11076708_latest.zip -o /tmp/cmdtools.zip \
 && unzip -q /tmp/cmdtools.zip -d $ANDROID_SDK_ROOT/cmdline-tools/latest \
 && rm /tmp/cmdtools.zip

RUN yes | sdkmanager --licenses || true \
 && yes | sdkmanager \
      "platform-tools" \
      "platforms;android-34" \
      "build-tools;34.0.0" \
      "ndk;26.1.10909125"

RUN go install golang.org/x/mobile/cmd/gomobile@latest \
 && go install golang.org/x/mobile/cmd/gobind@latest \
 && gomobile version
