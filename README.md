# SSG - Simple Static Site Generator

A lightweight, easy-to-use static site generator written in Go. Convert your markdown files into a beautiful static website with minimal configuration.

## Features

- **Simple CLI** - Easy-to-use command-line interface
- **Markdown Support** - Write content in markdown with TOML frontmatter
- **Theme System** - Customizable themes with Go templates
- **Default Theme** - Clean, responsive theme included out of the box
- **Fast Builds** - Lightning-fast site generation
- **Zero Config** - Works out of the box with sensible defaults

## Installation

```bash
go install github.com/deelawn/ssg/cmd/ssg@latest
```

Or build from source:

```bash
git clone https://github.com/deelawn/ssg.git
cd ssg
go build -o ssg ./cmd/ssg
```

## Quick Start

### 1. Initialize a New Site

```bash
ssg init my-site
cd my-site
```

This creates:
- `config.toml` - Site configuration
- `content/posts/` - Blog posts directory
- `content/pages/` - Static pages directory
- `themes/default/` - Default theme
- `public/` - Generated output directory

### 2. Add Content

Create a new blog post:

```bash
ssg add post "My First Post"
```

Create a new page:

```bash
ssg add page about
```

### 3. Build Your Site

```bash
ssg build
```

Your static site will be generated in the `public/` directory.

## Configuration

Edit `config.toml` to customize your site:

```toml
[site]
title = "My Static Site"
description = "A site generated with SSG"
domain = "example.com"
author = "Your Name"

[build]
content_dir = "content"
output_dir = "public"

[theme]
name = "default"
dir = "themes"
```

## Content Format

Content files use TOML frontmatter:

```markdown
+++
title = "My Post Title"
date = 2024-01-15
draft = false
author = "Your Name"
description = "Post description"
+++

# Your Content Here

Write your content in markdown...
```

## Commands

### init

Initialize a new static site:

```bash
ssg init [directory]
```

### add

Add new content:

```bash
ssg add post <name>    # Create a new blog post
ssg add page <name>    # Create a new page
```

### build

Build the static site:

```bash
ssg build
```

## Themes

### Default Theme

SSG comes with a clean, responsive default theme that includes:
- Responsive navigation
- Blog post listing on homepage
- Individual post and page layouts
- Modern, minimal styling

### Custom Themes

Create custom themes in the `themes/` directory:

```
themes/
  your-theme/
    templates/
      base.html     # Base layout
      index.html    # Homepage
      post.html     # Blog post layout
      page.html     # Page layout
    static/
      style.css     # Your styles
```

Templates use Go's `html/template` syntax. The base template defines the overall structure, while content templates define the main content area using `{{define "content"}}`.

## Project Structure

```
my-site/
├── config.toml
├── content/
│   ├── posts/
│   │   └── my-post.md
│   └── pages/
│       └── about.md
├── themes/
│   └── default/
│       ├── templates/
│       │   ├── base.html
│       │   ├── index.html
│       │   ├── post.html
│       │   └── page.html
│       └── static/
│           └── style.css
└── public/              # Generated site
    ├── index.html
    ├── posts/
    │   └── my-post.html
    ├── pages/
    │   └── about.html
    └── static/
        └── style.css
```

## Deployment

After building your site, deploy the `public/` directory to any static hosting service:

- **GitHub Pages**: Push to your repository and enable Pages
- **Netlify**: Drag and drop the `public/` folder
- **Vercel**: Deploy via CLI or web interface
- **AWS S3**: Upload to an S3 bucket with static website hosting enabled

## License

MIT License - See LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
