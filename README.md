# mmk-cardnews-v2

Card news generator for Instagram/LinkedIn — powered by Remotion React + Gemini AI images.

A Claude Code plugin that generates professional card news (카드뉴스) from any topic.

## Quick Start

### As Claude Code Plugin

```bash
claude plugin add github.com/pureugong/mmk-cardnews-v2
```

Then use the `/cardnews` skill in Claude Code.

### Manual Setup

```bash
git clone https://github.com/pureugong/mmk-cardnews-v2.git
cd mmk-cardnews-v2
./setup.sh
```

## CLI Commands

```bash
# Create new card news project
mmk-cn new "Topic" --slides 7

# Generate AI images (requires GEMINI_API_KEY)
mmk-cn imagegen -p "prompt" -o public/card-01.png

# Live preview in Remotion Studio
mmk-cn preview ./output/project/

# Export PNG stills
mmk-cn still ./output/project/

# Export MP4 video
mmk-cn render ./output/project/

# Generate PDF carousel
mmk-cn pdf ./output/project/out/stills/
```

## Pipeline

```
Topic
  → mmk-cn new (scaffold Remotion project)
  → mmk-cn imagegen (AI images via Gemini)
  → Write Card01.tsx ~ CardN.tsx (React components)
  → mmk-cn still / render / pdf (export)
```

## Architecture

- **Go CLI** (`mmk-cn`) — orchestration, scaffolding, image generation
- **Remotion React** — visual rendering (PNG stills, MP4 video)
- **Gemini API** (`gemini-3.1-flash-image-preview`) — AI image generation
- **No Python. No chromedp.**

## Requirements

- Node.js 18+
- Go 1.23+ (for building from source) or download prebuilt binary
- Gemini API key (optional, for AI image generation)

## Configuration

```bash
cp .env.example .env
# Set GEMINI_API_KEY for image generation
```

## License

MIT
