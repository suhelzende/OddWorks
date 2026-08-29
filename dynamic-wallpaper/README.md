# Dynamic Wallpaper

A tiny Windows utility written in Go that changes your desktop wallpaper based on the time of day.

## Config

Edit `config.json`:

```json
{
  "wallpapers": [
    {
      "time": "05:30",
      "file": "wallpapers/6am.png"
    },
    {
      "time": "12:30",
      "file": "wallpapers/11am.png"
    },
    {
      "time": "16:00",
      "file": "wallpapers/4pm.png"
    },
    {
      "time": "19:00",
      "file": "wallpapers/7pm.png"
    },
    {
      "time": "20:00",
      "file": "wallpapers/8pm.png"
    }
  ]
}
```

## Setup

### 1. Create the folder

```text
C:\DynamicWallpaper\
```

### 2. Keep the files in this structure

```text
C:\DynamicWallpaper\
├── dynamic-wallpaper.exe
├── config.json
└── wallpapers\
    ├── 6am.png
    ├── 11am.png
    ├── 4pm.png
    ├── 7pm.png
    └── 8pm.png
```

### 3. Add wallpapers

Put your wallpaper images inside the `wallpapers` folder.

### 4. Configure the schedule

Edit `config.json` and set the desired time and wallpaper.

Use `HH:MM` 24-hour format.

### 5. Run

Open PowerShell:

```powershell
cd C:\DynamicWallpaper
.\dynamic-wallpaper.exe
```

## Start Automatically on Windows

### 1. Open the Startup folder

Press `Win + R` and enter:

```text
shell:startup
```

### 2. Create a shortcut

Create a shortcut to:

```text
C:\DynamicWallpaper\dynamic-wallpaper.exe
```

Place the shortcut in the Startup folder.

### 3. Set the working directory

Open the shortcut's **Properties** and set **Start in** to:

```text
C:\DynamicWallpaper
```

Windows will now start the utility automatically when you log in.

## Build

### Requirements

- Windows
- Go

### Install dependency

```bash
go get golang.org/x/sys/windows
```

### Build

```bash
go build -ldflags="-H=windowsgui" -o dynamic-wallpaper.exe
```

## How It Works

The utility checks the current time periodically.

When the current time matches a configured time, the corresponding wallpaper is applied.

The configuration is re-read periodically, so changes to `config.json` take effect without restarting the application.

## Stop

Open **Task Manager → Details**, find `dynamic-wallpaper.exe`, and select **End task**.

To prevent it from starting with Windows, remove its shortcut from the Startup folder.
