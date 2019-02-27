FROM ubuntu:18.10
MAINTAINER Alex Zoer  <r5am@mail.ru>

# Рабочая директория приложения внутри контейнера
WORKDIR /usr/local/sbin

# Копировать в образ файлы приложения
COPY fast_fake_rest .
RUN mkdir -p /usr/local/etc/fast_fake_rest
COPY config.yaml /usr/local/etc/fast_fake_rest/config.yaml
RUN mkdir -p /var/log/fast_fake_rest

# Открыть порт чтобы он был доступен снаружи контейнера
EXPOSE 8083

# Запустить приложение внутри контейнера
CMD fast_fake_rest

