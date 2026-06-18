# ApkCheckPack

**APK Protection Detection Tool** — Detect Android APK protection features, third-party SDKs, security checks, and hardcoded information

## Features

- **Protection Detection** — Based on feature library to detect 40+ vendors' protection solutions (so library name/path, class name, regex matching)
- **Security Detection** — Detect anti-environment features such as ROOT, emulator, anti-debugging, proxy detection
- **Third-party SDK** — Identify common advertising, push, analytics SDKs
- **Hardcode Scanning** — Scan for sensitive information leakage (optional, requires `-hardcode`)
- **Certificate Scanning** — Detect certificate files in APK and output details
- **Embedded APK** — Support recursive scanning of embedded APK files like XAPK

## Build

```bash
go build -o ApkCheckPack.exe ./cmd/apkcheckpack
```

## Usage

```bash
ApkCheckPack.exe -f <APK file path or folder>
```

### Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `-f` | required | APK file path or folder path |
| `-root` | true | Detect ROOT features |
| `-emu` | true | Detect emulator features |
| `-debug` | true | Detect anti-debugging features |
| `-proxy` | true | Detect proxy features |
| `-sdk` | true | Detect third-party SDKs |
| `-hardcode` | false | Enable hardcode scanning (optional) |
| `-cert` | true | Detect certificate information |
| `-maxsize` | 500 | Max scan size per file (MB) |
| `-r` | true | Recursive scanning of embedded APKs |

### Examples

```bash
# Scan single APK
ApkCheckPack.exe -f test.apk

# Scan folder
ApkCheckPack.exe -f ./apks

# Enable hardcode scanning
ApkCheckPack.exe -f test.apk -hardcode true

# Limit scan file size
ApkCheckPack.exe -f test.apk -maxsize 100
```

![fun](doc/fun.png)

## Project Structure

```
ApkCheckPack/
├── cmd/apkcheckpack/main.go     # Main entry
├── pkg/
│   ├── sdk/                     # Third-party SDK detection
│   ├── pack/                    # Protection detection
│   ├── anti/                    # ROOT/emulator/anti-debugging/proxy detection
│   ├── hardcode/                # Hardcode detection
│   └── certificate/             # Certificate detection
└── go.mod
```

**Built-in Rules** — All rule data is compiled into the executable, no external configuration required.

```
sopath      Absolute path of feature so files
soname      Feature so file names only
other       Other feature files, strings
soregex     Regex matching for versioned feature so libraries
jclass      Java compiled code in dex, string matching
```

## Note

Tools are only auxiliary. New methods and vendors are constantly emerging, and feature detection may be overlooked. Do not rely solely on them.

Welcome to submit rules or provide samples of unrecognized protection apps for continuous updates.
