# Git-Broski 

https://gratis-neon-644.notion.site/GitBroski-Broski-for-your-Github-Workflow-286e41747653800a9cc4f7b014c6cf51

**Broski for your Git!**  
A CLI tool to perform various manual tasks with single commands

## Usage
1. Open currently working directory in github/gitlab
```bash
gitbroski open
```
2. Auto create .gitignore files
```bash
gitbroski ignore <Language>
```
Currently supports python<br>
Issues open for node.js, golang and many more language<br>

3. Easy Empty commit
```bash
gitbroski empty commit <your-Message>
```


## Installation

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/gitbroski.git
cd gitbroski
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Build the Project
```bash
go build -o gitbroski ./cmd
```

### 4. (Optional) Create a Symlink for Global Use
```bash
sudo ln -s /full/path/to/gitbroski /usr/local/bin/gitbroski
```
> This allows you to run `gitbroski` from **any directory**.

- Opens the **current Git repository** in your default browser.

## Contribution Guide
1. **Fork** and **clone** the repository.
2. Install dependencies:
```bash
go mod tidy
```
3. Build the project:
```bash
go build -o gitbroski ./cmd
```
4. Create a symlink for global use (optional):
```bash
sudo ln -s /full/path/to/gitbroski /usr/local/bin/gitbroski
```
5. Run `gitbroski` from **any folder** and contribute!

## License
MIT
