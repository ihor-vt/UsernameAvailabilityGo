🔍 Social Username Availability Checker

![Usage GIF](https://github.com/ihor-vt/UsernameAvailabilityGo/tree/main/assets/UsernameAvailabilityGo.gif)

Welcome to Social Username Availability Checker! This is a high-performance, command-line tool built in Go designed to quickly verify if a specific username is available across multiple social media platforms. Perfect for securing a consistent online identity or simply finding out if a username is taken, this tool is a must-have for brand managers, social media strategists, and developers.

🚀 Features

    •	Multi-Platform Search: Check availability on popular social platforms (e.g., Instagram, GitHub, X, LinkedIn, Pinterest) in seconds.
    •	Concurrency-Powered: Utilizes Go’s concurrency to query platforms in parallel, ensuring lightning-fast results.
    •	TUI (Text User Interface): Clean, intuitive interface for smooth user experience right in your terminal.
    •	Easily Extensible: Written in a modular format to support adding new platforms or adjusting search parameters with minimal effort.

💻 Installation

_Option 1: Build from Source_

Ensure you have Go 1.23.2 or later installed on your system.

1. Clone this repository:

```shell
git clone https://github.com/ihor-vt/UsernameAvailabilityGo.git
```

```shell
cd pro-go-checker
```

2. Build the application:

```shell
go build -o pro-go-checker
```

3. Run the application:

```shell
./pro-go-checker
```

_Option 2: Download Binary_

1. Download the latest binary for your system from the releases page.
2. Make the binary executable (if necessary):

```shell
chmod +x pro-go-checker
```

3. Run the application:

```shell
./pro-go-checker
```

🛠 Usage

Enter the desired username, and the tool will handle the rest! You’ll see a summary of results from each platform indicating if the username is available or already taken.

```shell
./pro-go-checker
```

📊 Supported Platforms

Currently, the application supports:
• GitHub
• X (formerly Twitter)
• LinkedIn
• Pinterest
• And more…

    Additional platforms can be easily integrated upon request or via contribution.

🤝 Contributing

Contributions are welcome! If you’d like to add support for more platforms, improve the interface, or enhance the functionality, feel free to open an issue or a pull request.

Thank you for using Social Username Availability Checker! If you find this tool helpful, feel free to ⭐ the repo.
