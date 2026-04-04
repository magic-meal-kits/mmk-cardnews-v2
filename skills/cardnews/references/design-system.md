# Design System

## Dimensions

| Output | Width | Height | Ratio |
|--------|-------|--------|-------|
| Card (default) | 1080 | 1350 | 4:5 |
| Vertical video | 1080 | 1920 | 9:16 |

## Typography

### Font Stack
- Title: `'Noto Serif KR', serif` — weight 900
- Body: `'Noto Sans KR', sans-serif` — weight 300-500
- UI elements: `'Noto Sans KR', sans-serif` — weight 300

### Sizes (1080px canvas)
| Element | Size | Weight | Line-height |
|---------|------|--------|------------|
| Cover title | 56px | 900 | 1.2 |
| Section title | 48px | 900 | 1.25 |
| Body text | 36px | 300 | 1.55 |
| Date/tag | 32px | 500 | 1.4 |
| Source/info | 30px | 300 | 1.6 |
| Page number | 28px | 300 | 1.0 |

### Korean Text
- `wordBreak: "keep-all"` — prevents mid-syllable breaks
- `letterSpacing`: -0.03em for titles, 0 for body
- Line-height must be generous (1.5+) for Korean readability

## Colors

### Base Palette
- Deep black: `#0A0A0A`
- Off-white: `#FAFAFA`
- Text on dark: `#FAFAFA` (primary), `rgba(250,250,250,0.7)` (secondary)
- Text on light: `#0A0A0A` (primary), `#555555` (secondary)
- Page numbers: `rgba(255,255,255,0.4)` (dark), `rgba(0,0,0,0.3)` (light)

### Gradient Overlays
- Image text readability: `linear-gradient(to bottom, transparent 40%, rgba(0,0,0,0.8) 100%)`
- Vertical video blur bg: `blur(36px) saturate(0.95) brightness(0.45)`

## Animation

### Standard Timings (30fps)
| Animation | Frames | Seconds |
|-----------|--------|---------|
| Background fade-in | 9 | 0.3s |
| Title entrance | 15 | 0.5s |
| Body entrance | 24 | 0.8s |
| Card fade-out | 15 | 0.5s |
| Audio delay | 20 | 0.67s |
| Subtitle fade | 8 | 0.27s |

### Interpolation
All animations use `interpolate()` with `extrapolateRight: "clamp"` / `extrapolateLeft: "clamp"`.

## Spacing
- Card padding: 80px horizontal
- Title-to-body gap: 16-32px
- Bottom margin for page number: 40px
- Split layout text padding: 72px left, 48px right
