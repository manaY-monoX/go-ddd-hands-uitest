# Golang 1.24.2とDebianをベースイメージとして使用
FROM golang:1.24.2

# 必要なパッケージをインストール
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    bash \
    wget \
    git \
    curl \
    ca-certificates \
    libc6 \
    libnss3 \
    libnspr4 \
    alsa-utils \
    libatk1.0-0 \
    libatspi2.0-0 \
    libcairo2 \
    libcups2 \
    dbus \
    libexpat1 \
    libflac12 \
    libgdk-pixbuf-2.0-0 \
    libglib2.0-0 \
    libjpeg62-turbo \
    libpng16-16 \
    libwebp7 \
    libxcomposite1 \
    libxdamage1 \
    libxext6 \
    libxfixes3 \
    libxrandr2 \
    mesa-utils \
    libpango-1.0-0 \
    libsnappy1v5 \
    libstdc++6 \
    xvfb \
    fonts-freefont-ttf \
    npm && \
    rm -rf /var/lib/apt/lists/*

# Playwrightのインストール
RUN npm install -g playwright && \
    npx playwright install --with-deps

# アプリケーションディレクトリの作成と設定
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app

# デフォルトのコマンド（必要に応じて変更）
CMD ["bash"]