# Webgen - A web app generator tool
**webgen** is a tool written in the **Go** programming language to help me personally to generate web applications like a react app.

# Install
Check if `$ go env GOPATH` is on your `$PATH`.

**if not**, follow the steps for your operating system
- (Windows: https://windowsloop.com/how-to-add-to-windows-path/)
- (Linux/Mac: https://linuxize.com/post/how-to-add-directory-to-path-in-linux/)


## Build the program
```sh
$ go build -o bin/
```
## Install to `$GOPATH`
```sh
$ go install
```


# Usage
- Show the help message
```sh
$ webgen --help
```
- Generate a web app (check templates)
```sh
$ webgen gen [[TEMPLATE]] [[APP NAME]]
```

# Templates
Just replace `[[TEMPLATE]]` with the name of any of the following:
-  **default**  - Simple website (HTML/CSS/JS)
-  **react**  - JavaScript React boilterplate
-  **react-ts**  - TypeScript React boilterplate

## Example
```sh
$ webgen gen react-ts my_react_ts_app
```