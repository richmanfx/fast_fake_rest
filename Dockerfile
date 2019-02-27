FROM ubuntu:18.10
MAINTAINER Alex Zoer  <r5am@mail.ru>

# Рабочая директория приложения внутри контейнера
WORKDIR /usr/local/sbin

# Копировать в образ файлы приложения
COPY fast_fake_rest .

# Открыть порт чтобы он был доступен снаружи контейнера
EXPOSE 8083

# Запустить приложение внутри контейнера
CMD fast_fake_rest

