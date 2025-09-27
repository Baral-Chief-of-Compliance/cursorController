# Activity Monitor - Browser Auto Killer

Программа для автоматического закрытия браузера при неактивности пользователя (отсутствии действий с мышкой и клавиатурой).

## Описание

Программа работает в фоновом режиме и отслеживает активность пользователя. Если в течение заданного времени не происходит действий с мышкой или клавиатурой, программа автоматически закрывает указанный браузер.

## Функциональность

- 📊 Мониторинг активности мыши и клавиатуры
- ⏰ Настраиваемое время бездействия (по умолчанию 5 минут)
- 🌐 Поддержка различных браузеров (по умолчанию chromium-browser)
- 🔧 Простая настройка через флаги командной строки
- 🐧 Оптимизировано для Linux

## Требования

### Системные зависимости (Ubuntu/Debian)

Перед сборкой установите необходимые зависимости:

```bash
sudo apt update
sudo apt install gcc libc6-dev
sudo apt install libx11-dev xorg-dev libxtst-dev
sudo apt install xsel xclip
sudo apt install libpng++-dev
sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev libxkbcommon-dev
```

### Go

Требуется Go версии 1.24.6 или выше.

## Установка и сборка

1. **Скачайте исходный код:**
```bash
git clone https://github.com/Baral-Chief-of-Compliance/cursorController.git
cd cursorController
```

2. **Соберите программу:**
```bash
go build -o activity-monitor
```

## Использование

### Базовое использование (с значениями по умолчанию):
```bash
./activity-monitor
```
- Время бездействия: 5 минут
- Браузер для закрытия: chromium-browser

### С кастомными настройки:
```bash
./activity-monitor -delta-time=10 -browser-name="firefox"
```

### Параметры командной строки:

| Флаг | Описание | По умолчанию |
|------|----------|-------------|
| `-delta-time` | Время бездействия в минутах | `5` |
| `-browser-name` | Имя процесса браузера | `chromium-browser` |
| `-help` | Показать справку | - |

### Примеры использования:

```bash
# Закрывать Firefox после 3 минут бездействия
./activity-monitor -delta-time=3 -browser-name="firefox"

# Закрывать Google Chrome после 10 минут бездействия
./activity-monitor -delta-time=10 -browser-name="chrome"
```

## Запуск в фоновом режиме

Для автоматического запуска при старте системы через Openbox, добавьте в файл автозапуска:

```bash
sudo nano /etc/xdg/openbox/autostart
```

Добавьте строку (укажите правильный путь к файлу):
```bash
~/cursorController/activity-monitor &
```

Или если программа находится в домашней директории:
```bash
~/activity-monitor &
```

После перезагрузки система будет автоматически запускать монитор активности.

## Структура проекта

```
cursorController/
├── go.mod              # Go модули
├── go.sum              # Зависимости
├── cursorListener.go   # Основной код программы
└── README.md          # Этот файл
```
