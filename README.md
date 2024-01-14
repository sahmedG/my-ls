My-LS
![image](https://parsifar.com/wp-content/uploads/2021/10/ls-command.jpg)
Overview
My-LS is a GoLang program that emulates the functionality of the ls command, providing a simple and customizable way to list files and directories in a specified location. The program supports various flags, including -r, -R, -t, -a, and -l, enabling users to tailor the output to their specific needs.

Features

Directory Listing: Lists files and directories in the specified location.
Flag Support: Supports the following flags:

-r: Reverse the order of the listing.
-R: Recursively list subdirectories encountered.
-t: Sort files by modification time, newest first.
-a: Include entries that start with a dot (hidden files).
-l: List in long format, providing additional information.

Customizable Output: Users can combine flags to customize the output according to their preferences.
Usage
bash
Copy code
go run main.go [flags] [directory]

Flags
-r: Reverse order.
-R: Recursively list subdirectories.
-t: Sort by modification time.
-a: Include hidden files.
-l: List in long format.

Example

bash
Copy code
go run main.go -lrt /path/to/directory
Output
The output is a formatted list of files and directories based on the specified flags. For example:

bash

my-ls
├── main.go
├── file1.txt
├── folder1
│   ├── file2.txt
│   └── file3.txt
├── folder2
│   ├── file4.txt
│   └── file5.txt
└── .hidden_file

Installation
To install My-LS, clone the repository and build the executable:

bash

git clone https://github.com/yourusername/my-ls.git
cd my-ls
go build

Contributing
Contributions are welcome! Please follow the contribution guidelines when submitting pull requests.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments
Inspired by the ls command in Unix-like operating systems.
Special thanks to the Go community for their valuable contributions.
Contact
For issues or suggestions, please open an issue.

Happy listing! 📂

# Contributers
sahmedG (Sameer Goumaa)
MSK17A (Mohammed Alsammak)
