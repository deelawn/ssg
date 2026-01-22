package embedtheme

import (
	"os"
	"path/filepath"
)

// ThemeFiles contains the default theme file contents
var ThemeFiles = map[string]string{
	"templates/base.html": `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{if .Title}}{{.Title}} - {{end}}{{.Site.Title}}</title>
    <meta name="description" content="{{if .Description}}{{.Description}}{{else}}{{.Site.Description}}{{end}}">
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <header>
        <nav>
            <h1><a href="/">{{.Site.Title}}</a></h1>
            <ul>
                <li><a href="/">Home</a></li>
                <li><a href="/pages/about.html">About</a></li>
            </ul>
        </nav>
    </header>

    <main>
        {{template "content" .}}
    </main>

    <footer>
        <p>&copy; {{.Year}} {{.Site.Author}}{{if .Site.Author}} - {{end}}{{.Site.Title}}</p>
        <p>Built with SSG</p>
    </footer>
</body>
</html>
`,
	"templates/post.html": `{{define "content"}}
<article class="post">
    <header class="post-header">
        <h1>{{.Title}}</h1>
        {{if .Date}}
        <time datetime="{{.Date}}">{{.Date}}</time>
        {{end}}
        {{if .Author}}
        <p class="author">By {{.Author}}</p>
        {{end}}
    </header>
    <div class="post-content">
        {{.Content}}
    </div>
</article>
{{end}}
`,
	"templates/page.html": `{{define "content"}}
<article class="page">
    <header class="page-header">
        <h1>{{.Title}}</h1>
    </header>
    <div class="page-content">
        {{.Content}}
    </div>
</article>
{{end}}
`,
	"templates/index.html": `{{define "content"}}
<div class="home">
    <h1>{{.Site.Title}}</h1>
    <p>{{.Site.Description}}</p>

    {{if .Posts}}
    <section class="posts">
        <h2>Recent Posts</h2>
        <ul class="post-list">
            {{range .Posts}}
            <li class="post-item">
                <h3><a href="{{.URL}}">{{.Title}}</a></h3>
                {{if .Date}}
                <time datetime="{{.Date}}">{{.Date}}</time>
                {{end}}
                {{if .Description}}
                <p>{{.Description}}</p>
                {{end}}
            </li>
            {{end}}
        </ul>
    </section>
    {{end}}
</div>
{{end}}
`,
	"static/style.css": `/* Reset and base styles */
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

:root {
    --primary-color: #2c3e50;
    --secondary-color: #3498db;
    --text-color: #333;
    --bg-color: #fff;
    --border-color: #e1e8ed;
    --max-width: 800px;
}

body {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
    line-height: 1.6;
    color: var(--text-color);
    background-color: var(--bg-color);
}

/* Header and Navigation */
body > header {
    background-color: var(--primary-color);
    color: white;
    padding: 1rem 0;
    border-bottom: 3px solid var(--secondary-color);
}

nav {
    max-width: var(--max-width);
    margin: 0 auto;
    padding: 0 1rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

nav h1 {
    font-size: 1.5rem;
}

nav h1 a {
    color: white;
    text-decoration: none;
}

nav ul {
    list-style: none;
    display: flex;
    gap: 1.5rem;
}

nav a {
    color: white;
    text-decoration: none;
    transition: color 0.3s;
}

nav a:hover {
    color: var(--secondary-color);
}

/* Main content */
main {
    max-width: var(--max-width);
    margin: 2rem auto;
    padding: 0 1rem;
    min-height: calc(100vh - 200px);
}

/* Articles */
article {
    margin-bottom: 2rem;
}

.post-header, .page-header {
    background-color: var(--bg-color);
    margin-bottom: 1.5rem;
    padding-bottom: 1rem;
    border-bottom: 3px solid var(--border-color);
}

.post-header h1, .page-header h1 {
    color: var(--text-color);
    margin-bottom: 0.5rem;
}

.post-header time, .page-header time {
    color: #666;
    font-size: 0.9rem;
    display: block;
    margin-bottom: 0.25rem;
}

.post-header .author, .page-header .author {
    color: #666;
    font-style: italic;
}

/* Content */
.post-content, .page-content {
    line-height: 1.8;
}

.post-content h1,
.post-content h2,
.post-content h3,
.page-content h1,
.page-content h2,
.page-content h3 {
    margin-top: 1.5rem;
    margin-bottom: 0.75rem;
    color: var(--primary-color);
}

.post-content p,
.page-content p {
    margin-bottom: 1rem;
}

.post-content code,
.page-content code {
    background-color: #f4f4f4;
    padding: 0.2rem 0.4rem;
    border-radius: 3px;
    font-family: "Courier New", monospace;
}

.post-content pre,
.page-content pre {
    background-color: #f4f4f4;
    padding: 1rem;
    border-radius: 5px;
    overflow-x: auto;
    margin-bottom: 1rem;
}

.post-content pre code,
.page-content pre code {
    background-color: transparent;
    padding: 0;
}

/* Home page */
.home h1 {
    color: var(--primary-color);
    margin-bottom: 0.5rem;
}

.home > p {
    color: #666;
    font-size: 1.1rem;
    margin-bottom: 2rem;
}

.posts {
    margin-top: 2rem;
}

.posts h2 {
    color: var(--primary-color);
    margin-bottom: 1rem;
    padding-bottom: 0.5rem;
    border-bottom: 2px solid var(--border-color);
}

.post-list {
    list-style: none;
}

.post-item {
    padding: 1.5rem 0;
    border-bottom: 1px solid var(--border-color);
}

.post-item:last-child {
    border-bottom: none;
}

.post-item h3 {
    margin-bottom: 0.5rem;
}

.post-item h3 a {
    color: var(--primary-color);
    text-decoration: none;
    transition: color 0.3s;
}

.post-item h3 a:hover {
    color: var(--secondary-color);
}

.post-item time {
    color: #666;
    font-size: 0.9rem;
}

.post-item p {
    margin-top: 0.5rem;
    color: #666;
}

/* Footer */
footer {
    background-color: #f8f9fa;
    padding: 2rem 0;
    margin-top: 3rem;
    border-top: 1px solid var(--border-color);
    text-align: center;
}

footer p {
    color: #666;
    font-size: 0.9rem;
    margin: 0.25rem 0;
}

/* Responsive */
@media (max-width: 768px) {
    nav {
        flex-direction: column;
        gap: 1rem;
    }

    nav ul {
        gap: 1rem;
    }
}
`,
}

// WriteThemeFiles writes the embedded theme files to the specified directory
func WriteThemeFiles(themeDir string) error {
	for relPath, content := range ThemeFiles {
		fullPath := filepath.Join(themeDir, relPath)

		// Ensure directory exists
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return err
		}

		// Write file
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return err
		}
	}
	return nil
}
