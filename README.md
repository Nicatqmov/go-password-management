# 🔐 Password Management CLI Tool

A lightweight, secure, and efficient **Go-based command-line tool** for generating and managing strong passwords. Perfect for developers, system administrators, and security-conscious users who need quick, reliable password generation without compromising security.

---

## ✨ Features

- ✅ **Generate Secure Random Passwords** - Cryptographically secure password generation
- 🔢 **Custom Password Length** - Define password length from 8 to 128 characters
- 🔣 **Symbol Control** - Toggle inclusion of special characters and symbols
- 🏷️ **Service Naming** - Specify which app or service the password is for (optional)
- ⚡ **Lightning Fast** - Minimal dependencies, optimized for speed
- 🛡️ **Secure** - Uses Go's crypto/rand for true randomness
- 💻 **Cross-Platform** - Works on Windows, macOS, and Linux
- 🎯 **User-Friendly** - Simple and intuitive CLI interface

---

## 📋 System Requirements

- **Go 1.18** or higher
- **Operating System**: Windows, macOS, or Linux
- **RAM**: Minimal (< 10MB)
- **Disk Space**: < 5MB

---

## 📦 Installation

### Prerequisites
Ensure you have Go installed on your system. Download from [golang.org](https://golang.org/dl)

### Clone the Repository
```bash
git clone https://github.com/Nicatqmov/go-password-management.git
cd go-password-management
```

### Build from Source
```bash
# Install dependencies
go mod download

# Build the executable
go build -o password-cli main.go

# (Optional) Install globally
go install
```

### Quick Start
```bash
# Make it executable (Linux/macOS)
chmod +x password-cli

# Run the tool
./password-cli --help
```

---

## 🚀 Usage

### Basic Password Generation
```bash
# Generate a 16-character password with symbols
./password-cli --length 16 --symbols

# Generate a 12-character password without symbols
./password-cli --length 12
```

### With Service Name
```bash
# Generate password for Gmail with custom length
./password-cli --service gmail --length 20 --symbols

# Generate password for database with no symbols
./password-cli --service mydb --length 16
```

### Available Flags

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--length` | `-l` | int | 16 | Password length (8-128 characters) |
| `--symbols` | `-s` | bool | false | Include special characters (!@#$%^&*) |
| `--service` | `-svc` | string | - | Name of the service/app (optional) |
| `--help` | `-h` | - | - | Display help information |
| `--version` | `-v` | - | - | Show version number |

### Examples

```bash
# Generate 20-character password with symbols for GitHub
./password-cli --service github --length 20 --symbols

# Generate secure database password
./password-cli --service database --length 24 --symbols

# Generate simple alphanumeric password
./password-cli --length 12

# Generate strong password
./password-cli --length 32 --symbols
```

---

## 🏗️ Project Structure

```
go-password-management/
├── main.go              # Application entry point
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── README.md           # Project documentation
├── LICENSE             # License file
└── .gitignore          # Git ignore rules
```

### Key Components

- **main.go**: Contains CLI setup, flag parsing, and password generation logic
- **Flag Parsing**: Uses Go's built-in `flag` package for command-line arguments
- **Random Generation**: Implements `crypto/rand` for secure password generation
- **Character Sets**: Supports uppercase, lowercase, numbers, and optional symbols

---

## 💡 How It Works

1. **Parse Arguments**: The tool reads command-line flags for length, symbols, and service name
2. **Build Character Set**: Creates a pool of available characters based on user preferences
3. **Generate Password**: Uses cryptographic randomness to select characters from the pool
4. **Output Result**: Displays the generated password with optional service information

---

## 🔒 Security Considerations

- **Cryptographic Randomness**: Uses `crypto/rand` for true randomness, not pseudo-random
- **No Logging**: Passwords are never logged or stored
- **No Network**: Completely offline - no data transmitted
- **Open Source**: Code is transparent and auditable
- **Recommendations**:
  - Always use `--symbols` flag for maximum security
  - Use minimum 16-character length for important accounts
  - Never share generated passwords over unencrypted channels
  - Consider using a password manager to store generated passwords

---

## 🛠️ Development

### Building for Different Platforms

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o password-cli-linux main.go

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o password-cli-macos main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o password-cli.exe main.go
```

### Running Tests
```bash
go test -v
```

### Code Formatting
```bash
# Format code
go fmt ./...

# Lint code
golangci-lint run
```

---

## 📚 Dependencies

This project uses only Go standard library packages:
- `flag` - Command-line flag parsing
- `fmt` - Formatted I/O
- `crypto/rand` - Cryptographic random number generation
- `math/big` - Big integer arithmetic
- `strings` - String manipulation

No external dependencies required!

---

## 🤝 Contributing

Contributions are welcome! Here's how to contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines
- Follow Go code style guidelines
- Add tests for new features
- Update README for user-facing changes
- Keep commits atomic and well-described

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🐛 Troubleshooting

### Issue: `command not found: password-cli`
**Solution**: Make sure you're in the project directory or have installed globally with `go install`

### Issue: Permission denied
**Solution**: 
```bash
chmod +x password-cli
```

### Issue: Build errors
**Solution**: 
```bash
go mod tidy
go build -o password-cli main.go
```

---

## 📞 Support & Contact

- **GitHub Issues**: Report bugs and request features on [GitHub Issues](https://github.com/Nicatqmov/go-password-management/issues)
- **Author**: Nicatqmov
- **Repository**: [go-password-management](https://github.com/Nicatqmov/go-password-management)

---

## 🎯 Roadmap

- [ ] Add password strength indicator
- [ ] Support for custom character exclusions
- [ ] Password history/logging option
- [ ] Configuration file support
- [ ] Interactive mode
- [ ] Clipboard integration
- [ ] Generate multiple passwords at once

---

## ⭐ Show Your Support

If you find this tool helpful, please consider:
- Starring the repository ⭐
- Sharing with friends and colleagues
- Contributing improvements
- Reporting bugs and suggesting features

---

**Happy Password Generating! 🔐**