<div align="center">

<h1>
🛸 RepoRover 🚀
</h1>

<p>
<strong>Your Smart Terminal Companion for GitHub Exploration</strong>
</p>

<p>
<a href="https://goreportcard.com/report/github.com/sidharthchauhan/reporover">
<img src="https://goreportcard.com/badge/github.com/sidharthchauhan/reporover" alt="Go Report Card">
</a>
<a href="https://golang.org/">
<img src="https://img.shields.io/github/go-mod/go-version/sidharthchauhan/reporover" alt="Go Version">
</a>
<a href="LICENSE">
<img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License">
</a>
<a href="https://github.com/sidharthchauhan/reporover/issues">
<img src="https://img.shields.io/github/issues/sidharthchauhan/reporover" alt="Issues">
</a>
</p>

<p>
<b>
<a href="#-features">Features</a> •
<a href="#%EF%B8%8F-quick-start">Quick Start</a> •
<a href="#-usage-examples">Usage</a> •
<a href="#-contributing">Contributing</a> •
<a href="#-license">License</a>
</b>
</p>

<pre>
🌟 Explore GitHub Repositories
📊 View Detailed Statistics
🌲 Browse Code Structure
🔍 Find Trending Projects
</pre>

</div>

## ✨ Features

> Explore GitHub like never before, right from your terminal!

- 🔍 **Smart Repository Discovery**
  - Find trending repositories by language
  - Get instant insights with rich summaries
  - Star counts and popularity metrics

- 📊 **Detailed Analytics**
  - Comprehensive repository statistics
  - Fork and issue tracking
  - Activity trends

- 🌲 **Code Navigation**
  - Browse repository structure
  - Quick file content preview
  - Easy directory traversal

## ⚡️ Quick Start

```bash
# Clone and build
git clone https://github.com/sidharthchauhan/reporover.git
cd reporover
go build -o rover

# Make it available system-wide
sudo cp rover /usr/local/bin/

# Start exploring!
rover explore go
```

## 🎯 Commands

### 🔍 Explore Repositories
```bash
# Find trending Go repositories
rover explore go

# Explore Python projects
rover explore python

# Discover JavaScript libraries
rover explore javascript
```

### 📊 Repository Statistics
```bash
# Get detailed stats for a repository
rover stats kubernetes/kubernetes

# View quick stats with --brief flag
rover stats golang/go --brief

# Check repository health
rover stats facebook/react
```

### 🌲 Browse Code Structure
```bash
# View full repository structure
rover tree golang/go

# Limit depth to 2 levels
rover tree nodejs/node --depth 2

# Focus on specific directory
rover tree kubernetes/kubernetes --path cmd/
```

### 🐛 Track Issues
```bash
# View all open issues
rover issues moby/moby

# See recent issues (last 7 days)
rover issues golang/go --recent

# Filter by label
rover issues kubernetes/kubernetes --label bug
```

### 👤 User Profiles
```bash
# View user profile
rover profile torvalds

# Get user's top repositories
rover profile sidharthchauhan --repos

# Show contribution activity
rover profile gopher --activity
```

### 🔄 Common Patterns
```bash
# Workflow 1: Explore and analyze
rover explore go
rover stats <interesting-repo>
rover tree <interesting-repo>

# Workflow 2: Issue tracking
rover issues kubernetes/kubernetes
rover stats kubernetes/kubernetes

# Workflow 3: User discovery
rover profile gopher
rover explore go --author gopher
```

### 💡 Tips
- Use `--help` with any command for detailed options
- All commands support JSON output with `--json` flag
- Press `q` to exit any paginated output
- Use `--limit N` to control number of results

## 🛠 Technical Stack

- **Language**: Go 1.21+
- **Core Dependencies**:
  - `github.com/spf13/cobra` - CLI framework
  - `github.com/fatih/color` - Terminal coloring
  - `net/http` - HTTP client for GitHub API
  - `encoding/json` - JSON parsing

## 🔧 Architecture

```
reporover/
├── cmd/           # Command implementations
├── github/        # GitHub API client
├── bin/           # Compiled binaries
└── main.go        # Entry point
```

### API Integration
- RESTful GitHub API v3
- Rate limiting: 60 requests/hour (unauthenticated)
- Optional: Personal access token support

### Error Handling
- Graceful error recovery
- User-friendly error messages
- Rate limit warnings

## 🤝 Contributing

We love your input! We want to make contributing to RepoRover as easy and transparent as possible. Here's how you can help:

### 🐛 Report Issues
- Check if the issue has already been reported
- Open a new issue with a clear title and description
- Include code samples and error messages if applicable

### 🔧 Submit Changes
1. Fork the repo and create your branch:
   ```bash
   git checkout -b feature/amazing-feature
   ```

2. Make your changes:
   - Follow the existing code style
   - Add comments for explain complex logic
   - Update documentation if needed

3. Test your changes:
   ```bash
   go test ./...
   ```

4. Commit your updates:
   ```bash
   git commit -m 'Add amazing feature'
   ```

5. Push and create a pull request:
   ```bash
   git push origin feature/amazing-feature
   ```

## 📄 License

MIT License - See [LICENSE](LICENSE) for details.

---

<div align="center">
<p>
<strong>Made with ❤️ by <a href="https://github.com/sidharthchauhan">Sidharth Chauhan</a></strong>
</p>
<p>
<a href="https://github.com/sidharthchauhan">
<img src="https://img.shields.io/github/followers/sidharthchauhan?label=Follow&style=social" alt="Follow on GitHub">
</a>
</p>
</div>
