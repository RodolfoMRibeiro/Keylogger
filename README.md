# Windows Keylogger - Learn Windows API and DLL Functions

A keylogger written in Go that works exclusively on Windows. This project aims to help developers learn about the Windows API and its DLL functions. It logs all keystrokes on the system, creating a file to store all these logs. Additionally, it provides a terminal hide mode to prevent unauthorized individuals from detecting its presence.

Please note that the primary purpose of this project is educational, and any misuse or illegal activities are strongly discouraged and not supported by the author.

## Features
- Logs all keystrokes on a Windows system
- Save logs into a .txt file 
- Provides terminal hide mode to conceal the keylogger

## Prerequisites

To run this keylogger, ensure that you have the following installed on your system:

- Go programming language (version 1.19.1 or higher)
- Windows operating system (tested on Windows 10)

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/RodolfoMRibeiro/Keylogger.git
   ```
   
2. Change into the project directory:

   ```shell
   cd Keylogger
   ```
   
3. Build the executable:

   ```shell
   go build -o keylogger.exe main.go
   ```
## Usage
1. Run the project by executing the binary generated during the building process.
2. Monitor the console output and log file.
3. Terminate the program to stop the keylogger.

## Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow these steps:
1. Fork this repository.
2. Create a new branch: `git checkout -b feature/your-feature-name`.
3. Make your changes and commit them: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature/your-feature-name`.
5. Open a pull request.

Please ensure that your code adheres to the existing code style and includes appropriate tests.

## Disclaimer

This project is intended for educational purposes only. The author does not condone or support any misuse or illegal activities conducted with this software. The keylogger should only be used with proper authorization and consent from the system owner. It is the responsibility of the user to comply with all applicable laws and regulations regarding privacy and data protection.

The author is not liable for any damages or legal implications caused by the misuse or unethical use of this software. Any unauthorized or inappropriate usage of this keylogger is strictly discouraged and goes against the intended purpose of this project.

By using this software, you agree to take full responsibility for your actions and to use it solely for educational purposes and within the bounds of the law.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
