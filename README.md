# Raspberry Blink Container
![LicenseBadge](https://img.shields.io/github/license/YuhriBernardes/raspberry-blink-container?logoColor=green&style=for-the-badge)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/YuhriBernardes/raspberry-blink-container?logo=go&style=for-the-badge)
![Docker Image Version (latest semver)](https://img.shields.io/docker/v/yuhribernardes/rasp-blink?color=blue&logo=Docker&style=for-the-badge)

A POC to validate how to access, manage and control raspberry's GPIO from a docker container

# Build container

```sh
docker build --no-cache -t=yuhribernardes/rasp-blink:<tag> .
```

> You can give the image the name you want

# Run on you Raspberry

Notice that you need to have installed Docker on it

```sh
docker run --device /dev/gpiomem -d yuhribernardes/rasp-blink:<tag>
```