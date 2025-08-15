# Программа для убийства браузера при бездействии
### Написана под linux(ubuntu)

#### Подготовка:

- sudo apt install gcc libc6-dev
- sudo apt install libx11-dev xorg-dev libxtst-dev
- sudo apt install xsel xclip
- sudo apt install libpng++-dev
- sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev libxkbcommon-dev

#### Сборка проекта:
- go mod tidy
- go build