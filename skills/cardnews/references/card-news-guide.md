# Card News Design Guide

## Card Types

| Type | Use Case | Structure |
|------|----------|-----------|
| **List** | Rankings, tips, places | Cover → Items 1~N → Closing |
| **Storytelling** | Process, timeline | Cover → Story flow → Conclusion |
| **Focus** | Single topic deep dive | Cover → Problem → Analysis → Solution → Closing |
| **Q&A** | FAQ, misconceptions | Cover → Q1/A1 → Q2/A2 → Summary |

## Design Principles

### Magazine Style Selection

Select a magazine style matching your content type first (see `magazine-styles.md`):

| Content Type | Style | Vibe |
|-------------|-------|------|
| Concept / theme / IT / trends | BRUTUS | Bold, high-contrast, conceptual |
| Lifestyle / food / city | POPEYE | Warm film, casual, city boy |
| Architecture / space / interior | Casa BRUTUS | Material, precise, architectural |
| Personal / nostalgic / vintage | OLIVE | Handcrafted, analog warmth |
| Portrait / interview | SWITCH | Portrait-first, minimal text |
| Art / minimal | Pen | Swiss grid, extreme white space |

### Layout
- **Size**: 1080 x 1350px (4:5 portrait, Instagram standard)
- **All cards use FullBleed pattern** — background image + dark overlay + text
- **No gradient-only backgrounds** — always use AI-generated images
- **Layout variety**: vary overlay intensity + text position for visual rhythm
- **Page counter**: Bottom center, "3/13" format

### Overlay Pattern
| Card Type | Overlay | Text Position |
|-----------|---------|---------------|
| Cover | Light (0.1~0.5) | Center or bottom |
| Content | Heavy (0.3~0.85) | Bottom (bottom: 100px) |
| Closing | Light (0.1~0.5) | Center |

### Typography
- **Title/Body font**: use typography constants from selected magazine style
- **Fallback**: Noto Sans KR (sans), Noto Serif KR (serif)
- **Korean text**: `word-break: keep-all` (no mid-word breaks)
- **Line-height**: Title 1.25-1.35, Body 1.55-1.8, Lists 1.6
- **Accent color**: use only for keyword emphasis (per-style accent color)

### Color
- **Auto-determined from selected magazine style palette**
- Text: always white (on dark overlay) or accent color (for emphasis)
- No solid background or gradient-only backgrounds — overlay on images only

### Anti-AI Design Principles
- Minimize emoji (0-1 on cover, none on content slides)
- Avoid perfect left-right symmetry
- Text ON image — no separate white background zones
- Bullet dots should be small and subtle (8px max)
- Avoid excessive box-shadow / border-radius

### Mobile Readability (Critical)
1080px canvas → mobile ~375px = ~1/2.88 scale.

| Element | Canvas min | Mobile equiv |
|---------|-----------|-------------|
| Title | 48px | ~17px |
| Body | 36px | ~12px |
| Date/tag | 32px | ~11px |
| Source/info | 30px | ~10px |
| Page number | 28px | ~10px |
| **Absolute minimum** | **24px** | ~8px |

### Image Rules
- Generate with NO text in images (pure visual only)
- Full color, not black and white
- Occupy 60%+ of card area in FullBleed layouts
- Use gradient mask for text readability over images

## Content Guidelines

### Writing Rules
- Default output language: Korean (polite form). User can specify other languages.
- Max 3-4 bullet points per slide
- Each bullet: under 25 characters
- Body text: 2-3 short sentences max
- Cover title: max 20 characters
- Content title: max 30 characters
