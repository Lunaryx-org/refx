# refx

A CLI tool to replace Go import paths across your entire project.

## Installation
```bash
go install github.com/Lunaryx-org/refx@latest
```

## Usage
```bash
refx <old-import-path> <new-import-path>
```

### Example

When moving a project from personal account to organization:
```bash
refx github.com/private/myproject github.com/myorg/myproject
```

This will:
- Find all `.go` files in the current directory and subdirectories
- Replace old import paths with new ones
- Safely update files using atomic file operations

### Use Cases

- Moving repositories between GitHub accounts
- Renaming your GitHub username
- Migrating from GitLab to GitHub (or vice versa)
- Refactoring internal import paths

## How it Works

refx scans all `.go` files recursively, finds lines containing the old import path, replaces them with the new path, and atomically updates each file to ensure data integrity.

## Safety

Each file is written to a temporary file first, then atomically renamed to replace the original. This ensures that if anything goes wrong, your original files remain intact.

## Roadmap
- [ ] Add a verbose flag --verbose for more explained output
- [x] Better output formatting with progress indicators
- [ ] Automatic backup before making changes
- [x] Ignore golang keywords  
- [ ] `--dry-run` flag to preview changes
- [ ] Color-coded output
- [ ] Statistics summary
- [ ] Only work on files that have that import path / ignore irevelant golang files
- [ ] Config file to apply rules if specified

## License

MIT

## Author

Gustavo Pereira (lunaryx.org@gmail.com)
