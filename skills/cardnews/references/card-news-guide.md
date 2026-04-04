# Card News Design Guide

## Card Types

| Type | Use Case | Structure |
|------|----------|-----------|
| **List** (나열형) | Rankings, tips, places | Cover → Items 1~N → Closing |
| **Storytelling** (스토리텔링형) | Process, timeline | Cover → Story flow → Conclusion |
| **Focus** (집중형) | Single topic deep dive | Cover → Problem → Analysis → Solution → Closing |
| **Q&A** (문답형) | FAQ, misconceptions | Cover → Q1/A1 → Q2/A2 → Summary |

## Design Principles

### Layout
- **Size**: 1080 x 1350px (4:5 portrait, Instagram standard)
- **Background**: Alternate #0A0A0A (deep black) / #FAFAFA (off-white)
- **Images**: Full color, editorial magazine photography (NOT B&W)
- **Layout variety**: Each card uses different pattern (FullBleed/Split/TextOnly)
- **Page counter**: Bottom center, "3/13" format

### Typography
- **Title font**: Serif (Noto Serif KR) — bold, 48px+
- **Body font**: Sans-serif (Noto Sans KR) — light/regular, 36px+
- **Korean text**: `word-break: keep-all` (no mid-word breaks)
- **Line-height**: Title 1.25, Body 1.55, Lists 1.6

### Color
- Auto-determine from topic mood:
  - Nature/spring → warm pastels, pink, green
  - Tech → dark + neon accents (blue, purple)
  - Food → warm tones (orange, brown)
  - Business → neutral + one accent color

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
- Korean (한국어), polite form (존댓말: ~해요, ~합니다)
- Max 3-4 bullet points per slide
- Each bullet: under 25 characters
- Body text: 2-3 short sentences max
- Cover title: max 20 characters
- Content title: max 30 characters
