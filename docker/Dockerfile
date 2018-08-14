FROM golang:1.9

RUN apt-get update -y && apt-get install lsb-release -y

RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
    echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    apt-get update -y && apt-get install google-cloud-sdk google-cloud-sdk-app-engine-go -y

RUN go get github.com/golang/dep/cmd/dep

WORKDIR /workspace/src/app
