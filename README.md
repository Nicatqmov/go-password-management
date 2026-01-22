# 🔐 Password CLI Tool

A lightweight and secure **Go-based CLI tool** for generating strong passwords.  
You can control password length, choose whether to include symbols, and specify the app or service name the password is for.

---

## ✨ Features

- ✅ Generate secure random passwords
- 🔢 Custom password length
- 🔣 Optionally include symbols
- 🏷️ Specify app/service name
- ⚡ Fast and simple CLI interface
- 🛠 Built in Go (Golang)

---

## 📦 Installation

### Clone the repository
```bash
git clone https://github.com/yourusername/password-cli.git
cd password-cli
Build the binary
go build -o password-cli
(Optional) Install globally
sudo mv password-cli /usr/local/bin/
🚀 Usage
password-cli [flags]
🧰 Available Flags
Flag	Description	Example
-length	Length of the password	-length 16
-symbols	Include symbols (true or false)	-symbols false
-app	Name of the app/service	-app github
📌 Examples
Generate a 16-character password with symbols
password-cli -length 16 -symbols true -app github
Generate a 12-character password without symbols
password-cli -length 12 -symbols false -app gmail
Minimal usage (defaults applied)
password-cli
⚙️ Default Configuration
If no flags are provided:

Length: 12

Symbols: true

App Name: unknown

🧪 Sample Output
App Name   : github
Password   : gA#9Kp!xQ2@Lm
Length     : 16
Symbols    : enabled
🔐 Security Notes
Passwords are generated using cryptographically secure randomness

Passwords are not stored by default

Suitable for local usage, scripting, and automation

🛠 Tech Stack
Language: Go (Golang)

Libraries: Standard library (crypto/rand, flag, fmt)

🗺 Roadmap (Optional)
 Save passwords securely

 Encrypt stored passwords

 Clipboard support

 Config file support

 Export passwords

🤝 Contributing
Contributions are welcome!
Feel free to open issues or submit pull requests.

📄 License
MIT License © 2026


---

This is **all properly formatted in Markdown**, ready to copy directly into a `README.md` file.  

If you want, I can now **write the full Go CLI code** that matches this README — with all flags, options for symbols, length, and app name — fully working and ready to run.  

Do you want me to do that next?