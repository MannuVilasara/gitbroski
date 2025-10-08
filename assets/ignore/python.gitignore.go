// Package ignore provides language-specific .gitignore templates.
package ignore

func Python() string {
	return `# Byte-compiled / optimized / DLL files
__pycache__/
*.py[cod]
*$py.class

# C extensions
*.so

# Distribution / packaging
.Python
build/
develop-eggs/
dist/
downloads/
eggs/
.eggs/
lib/
lib64/
parts/
sdist/
var/
wheels/
share/python-wheels/
*.egg-info/
.installed.cfg
*.egg
MANIFEST

# Installer logs
pip-log.txt
pip-delete-this-directory.txt

# Unit test / coverage reports
htmlcov/
.tox/
.coverage
.coverage.*
.cache
nosetests.xml
coverage.xml
*.cover
.hypothesis/
.pytest_cache/

# Translations
*.mo
*.pot

# Django Stuff
*.log
local_settings.py
db.sqlite3

# Flask stuff
instance/
.webassets-cache

# Jupyter Notebook stuff
.ipynb_checkpoints

# Virtual environment
.venv/
venv/
env/
/venv
/env

# Secrets/Configuration (should almost always be ignored)
.env
.flaskenv
settings.ini

# IDE-specific directories and files
# VS Code
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
.history/

# JetBrains IDEs (PyCharm, IntelliJ, etc.)
.idea/

# Windows
Thumbs.db
ehthumbs.db`
}
