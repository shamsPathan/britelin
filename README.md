# britelin

This repository contains a Go program that adjusts the brightness of the backlight on your system.

## Prerequisites

- Go installed on your machine ([Go Installation Guide](https://golang.org/doc/install))
- Linux system with sysfs support for backlight control

## Clone the Repository

```bash
git clone https://github.com/shamspathan/britelin.git
cd britelin
```

## Installation

Compile the program using `go build`:

```bash
go build -p=2 -ldflags="-s -w" -gcflags="all=-N -l" britelin.go 
```

## Usage

Run the program with the value you want to add or subtract from the current brightness:

```bash
sudo ./britelin <value_to_add>
```

- Replace `<value_to_add>` with the numeric value by which you want to adjust the brightness.
- Use a positive value to increase brightness and a negative value to decrease brightness.

Example:

```bash
# Increase brightness by 1
sudo ./britelin 1

# Decrease brightness by 1
sudo ./britelin -1
```

### Note:

- Ensure the program is executed with appropriate permissions to write to the backlight file (`/sys/class/backlight/intel_backlight/brightness`). You might need to use `sudo` or adjust permissions accordingly.

### Integrating with i3wm (i3 Window Manager)

#### Step 1: Build and Install britelin

Ensure you have already built the `britelin` program as described in the previous sections. If not, follow these steps:

```bash
git clone https://github.com/shamspathan/britelin.git
cd britelin
go build -p=2 -ldflags="-s -w" -gcflags="all=-N -l" britelin.go 
sudo mv britelin /usr/local/bin
```

#### Step 2: Configure i3wm

1. Open the i3wm configuration file (`~/.config/i3/config`) in a text editor.

2. Add the following lines to the configuration file:

```bash
bindsym XF86MonBrightnessDown exec britelin -1
bindsym XF86MonBrightnessUp exec britelin 1
```

These lines bind the `XF86MonBrightnessDown` and `XF86MonBrightnessUp` keys (typically used for brightness adjustment on keyboards) to execute the `britelin` program with the specified brightness adjustments (`-1` for decrease, `1` for increase).

#### Step 3: Reload i3wm Configuration

Save the changes to the i3wm configuration file and reload i3wm to apply the changes:

Press `Mod+Shift+R` to reload the i3wm configuration.

#### Usage:

After integrating `britelin` with i3wm, you can now use the designated brightness adjustment keys (usually the function keys related to brightness on your keyboard) to adjust the brightness level within the i3wm environment.

For example:

- Press `XF86MonBrightnessDown` to decrease brightness.
- Press `XF86MonBrightnessUp` to increase brightness.

#### Notes:

- Ensure the `britelin` program is in the system's path (`/usr/local/bin`) for i3wm to execute it correctly.
- Adjust the `-1` and `1` values in the i3wm configuration to change the brightness adjustment step size according to your preference.


## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the [MIT License](LICENSE).

