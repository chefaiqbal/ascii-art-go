# ASCII Art

ASCII Art is a program that generates graphic representations of input strings using ASCII characters.

## Usage

1. Ensure you have Go installed on your machine.
2. Clone this repository:

```bash
git clone https://github.com/your-username/ascii-art.git
Navigate to the project directory:
bash
Copy code
cd ascii-art
Run the program with an input string:
bash
Copy code
go run . "Hello, World!"
The program will generate the ASCII art representation of the input string and display it in the console.

Features
Supports input with numbers, letters, spaces, special characters, and line breaks.
Each character has a height of 8 lines.
Characters are separated by a new line (\n).
Uses a standard.txt file to store the ASCII representations of characters.
Customization
If you want to customize the ASCII characters or add more characters, you can edit the standard.txt file. Each character's ASCII representation should be separated by an empty line. Ensure the width and height of each character match the constant values defined in the code.

Example
Here is an example of generating ASCII art for the input string "Hello, World!":

 _    _          _   _                    __          __                 _       _   _  
| |  | |        | | | |                   \ \        / /                | |     | | | | 
| |__| |   ___  | | | |   ___              \ \  /\  / /    ___    _ __  | |   __| | | | 
|  __  |  / _ \ | | | |  / _ \              \ \/  \/ /    / _ \  | '__| | |  / _` | | | 
| |  | | |  __/ | | | | | (_) |  _           \  /\  /    | (_) | | |    | | | (_| | |_| 
|_|  |_|  \___| |_| |_|  \___/  ( )           \/  \/      \___/  |_|    |_|  \__,_| (_) 
                                |/                                                      

                                