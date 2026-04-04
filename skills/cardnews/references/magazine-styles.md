# Magazine Style Guide

> Design system based on 6 Japanese magazines. Select a style matching your content type when planning card news.

---

## Style Matching Rule

| Content Type | Style | Why |
|-------------|-------|-----|
| Concept / theme (IT terms, trends, social issues) | **BRUTUS** | Bold, concept-driven, each issue is a different visual world |
| Lifestyle / food / city / travel | **POPEYE** | Casual confidence, warm film, city boy sensibility |
| Architecture / space / interior / design objects | **Casa BRUTUS** | Material reverence, architectural precision |
| Personal / emotional / nostalgic / vintage | **OLIVE** | Handcrafted intimacy, analog warmth, feels like a letter |
| Portrait / interview / profile | **SWITCH** | Portrait-first, the photo IS the design |
| Art / minimal / high culture / single object | **Pen** | Swiss minimalism, white space IS the design |

---

## 1. BRUTUS

### Palette
```
Primary:     #1A1A1A  (near black)
Background:  #F5F0E8  (off white — warm paper)
Accent:      #D62B2B  (brutus red — signature, never changes)
Per-issue accent (pick ONE): #F5A623 | #1C3A5E | #D4A017 | #2D5A3D
```

### Typography
- **Headline**: Barlow Condensed Black (900), tight tracking (-0.03em), ALL CAPS allowed
- **Subheadline**: clean grotesque, medium weight, bilingual (EN + KR)
- **Body**: Noto Sans KR 400, tight leading

### Layout
- Text lives ON the image — never in a safe zone
- Multi-axis text (diagonal, vertical) allowed
- White space = deliberate choice
- Visual language reinvented per issue

### Overlay
- Cover: `linear-gradient(to bottom, rgba(0,0,0,0.2) 0%, rgba(0,0,0,0.6) 100%)`
- Content: `linear-gradient(to bottom, rgba(0,0,0,0.4) 0%, rgba(0,0,0,0.85) 100%)`
- High contrast always — never soft

### Remotion Constants
```typescript
export const BRUTUS = {
  colors: {
    primary: '#1A1A1A',
    background: '#F5F0E8',
    accent: '#D62B2B',
  },
  typography: {
    headline: { fontFamily: "'Barlow Condensed', sans-serif", fontWeight: 900, letterSpacing: '-0.03em', textTransform: 'uppercase' as const },
    body: { fontFamily: "'Noto Sans KR', sans-serif", fontWeight: 400 },
  },
  overlay: {
    cover: 'linear-gradient(to bottom, rgba(0,0,0,0.2) 0%, rgba(0,0,0,0.6) 100%)',
    content: 'linear-gradient(to bottom, rgba(0,0,0,0.4) 0%, rgba(0,0,0,0.85) 100%)',
    closing: 'linear-gradient(to bottom, rgba(0,0,0,0.2) 0%, rgba(0,0,0,0.6) 100%)',
  },
}
```

---

## 2. POPEYE

### Palette
```
Primary:     #1A1A1A  (film black)
Background:  #F7F4EE  (warm white — not clinical)
Secondary:   #E8E0D0  (warm cream)
Muted:       #B5A89A  (warm gray)
Accent (ONE): #4A7C9E | #D4563A | #7A9E7E | #D4B483
```

### Typography
- **Headline**: Futura PT Bold (700), ALL CAPS, 2-4 words max
- **Cover lines**: DM Sans 400, small, sparse (max 3)
- **Body**: Noto Sans KR 400

### Layout
- Text NEVER fights the image
- White space 30-40% of frame — breathing room IS the design
- Photography: single subject, casual energy, film feel
- "Unstudied" look is meticulously constructed

### Overlay
- Cover: `linear-gradient(to bottom, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0.45) 100%)` (gentle)
- Content: `linear-gradient(to bottom, rgba(0,0,0,0.3) 0%, rgba(0,0,0,0.75) 100%)`
- Always warm, never cold

### Remotion Constants
```typescript
export const POPEYE = {
  colors: {
    primary: '#1A1A1A',
    background: '#F7F4EE',
    cream: '#E8E0D0',
    muted: '#B5A89A',
  },
  typography: {
    headline: { fontFamily: "'DM Sans', sans-serif", fontWeight: 700, letterSpacing: '0.05em', textTransform: 'uppercase' as const },
    body: { fontFamily: "'Noto Sans KR', sans-serif", fontWeight: 400 },
  },
  overlay: {
    cover: 'linear-gradient(to bottom, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0.45) 100%)',
    content: 'linear-gradient(to bottom, rgba(0,0,0,0.3) 0%, rgba(0,0,0,0.75) 100%)',
    closing: 'linear-gradient(to bottom, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0.45) 100%)',
  },
}
```

---

## 3. Casa BRUTUS

### Palette
```
Primary:     #2C2C2C  (charcoal)
Background:  #F2EDE5  (architectural white — warm paper)
Concrete:    #8B9EA8  (concrete blue)
Stone:       #C4B49A  (travertine)
Plant:       #6B8E7A  (plant green)
Technical:   #1C4E8A  (blueprint blue)
```

### Typography
- **Headline**: serif + grotesque mix per issue, Bold-Black, left-aligned
- **Logo**: small, restrained, white or black only
- **Body**: strict column structure

### Layout
- Full-bleed architectural/interior photography
- Text in clear zones — never over busy backgrounds
- Strict editorial grid
- Material textures must be visible

### Overlay
- Cover: `linear-gradient(to bottom, rgba(0,0,0,0.15) 0%, rgba(0,0,0,0.55) 100%)`
- Content: `linear-gradient(to bottom, rgba(0,0,0,0.35) 0%, rgba(0,0,0,0.8) 100%)`

### Remotion Constants
```typescript
export const CASA_BRUTUS = {
  colors: {
    primary: '#2C2C2C',
    background: '#F2EDE5',
    concrete: '#8B9EA8',
    stone: '#C4B49A',
    plant: '#6B8E7A',
    blueprint: '#1C4E8A',
  },
  typography: {
    headline: { fontFamily: "'Noto Serif KR', serif", fontWeight: 700, letterSpacing: '0.01em' },
    body: { fontFamily: "'Noto Sans KR', sans-serif", fontWeight: 400 },
  },
  overlay: {
    cover: 'linear-gradient(to bottom, rgba(0,0,0,0.15) 0%, rgba(0,0,0,0.55) 100%)',
    content: 'linear-gradient(to bottom, rgba(0,0,0,0.35) 0%, rgba(0,0,0,0.8) 100%)',
    closing: 'linear-gradient(to bottom, rgba(0,0,0,0.15) 0%, rgba(0,0,0,0.55) 100%)',
  },
}
```

---

## 4. OLIVE

### Palette
```
Primary:     #8B7B6B  (warm brown)
Background:  #F0E8D5  (aged paper)
Accent:      #7B9E6B  (olive green — signature)
Warm:        #D4A574  (warm tan)
Skin:        #E8D5C4  (skin beige)
Rust:        #B5704A  (rust)
```

### Typography
- **Headline**: Cormorant Garamond Italic 500 — literary, gentle
- **Logo**: EB Garamond Italic, olive green
- **Annotations**: handwritten feel allowed
- Mixed EN/FR/KR bilingual

### Layout
- Non-mechanical grid — editorial collage feel
- Images and text in conversation
- Warm analog photography always
- Personal, not aspirational

### Overlay
- Cover: warm tint `linear-gradient(to bottom, rgba(64,48,32,0.15) 0%, rgba(64,48,32,0.5) 100%)`
- Content: `linear-gradient(to bottom, rgba(64,48,32,0.3) 0%, rgba(64,48,32,0.75) 100%)`

### Remotion Constants
```typescript
export const OLIVE = {
  colors: {
    primary: '#8B7B6B',
    background: '#F0E8D5',
    accent: '#7B9E6B',
    warm: '#D4A574',
    skin: '#E8D5C4',
    rust: '#B5704A',
  },
  typography: {
    headline: { fontFamily: "'Cormorant Garamond', serif", fontWeight: 500, fontStyle: 'italic' as const },
    body: { fontFamily: "'Noto Sans KR', sans-serif", fontWeight: 400 },
  },
  overlay: {
    cover: 'linear-gradient(to bottom, rgba(64,48,32,0.15) 0%, rgba(64,48,32,0.5) 100%)',
    content: 'linear-gradient(to bottom, rgba(64,48,32,0.3) 0%, rgba(64,48,32,0.75) 100%)',
    closing: 'linear-gradient(to bottom, rgba(64,48,32,0.15) 0%, rgba(64,48,32,0.5) 100%)',
  },
}
```

---

## 5. SWITCH

### Palette
```
Primary:     #1A1A1A  (deep black)
Background:  #F5F2EC  (gallery white)
Secondary:   #6B6B6B  (medium gray)
Muted:       #A89880  (warm mid)
Text:        always white or black — NEVER colored
```

### Typography
- **Subject name**: Neue Haas Grotesk 400, white — almost nothing else
- **Logo**: GT America 300, small, non-intrusive
- Text hierarchy: 1) Portrait → 2) Name → 3) Logo (never reverses)
- Near-zero text density

### Layout
- Photography fills entire frame
- Large empty space is correct and intentional
- Never add elements that distract from the subject
- Medium to large format portrait quality

### Overlay
- Cover: `linear-gradient(to bottom, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0.4) 100%)` (minimal)
- Content: `linear-gradient(to bottom, rgba(0,0,0,0.25) 0%, rgba(0,0,0,0.7) 100%)`

### Remotion Constants
```typescript
export const SWITCH_MAG = {
  colors: {
    primary: '#1A1A1A',
    background: '#F5F2EC',
    secondary: '#6B6B6B',
    muted: '#A89880',
  },
  typography: {
    headline: { fontFamily: "'Noto Sans KR', sans-serif", fontWeight: 400, color: '#FFFFFF' },
    body: { fontFamily: "'Noto Sans KR', sans-serif", fontWeight: 300, color: '#FFFFFF' },
  },
  overlay: {
    cover: 'linear-gradient(to bottom, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0.4) 100%)',
    content: 'linear-gradient(to bottom, rgba(0,0,0,0.25) 0%, rgba(0,0,0,0.7) 100%)',
    closing: 'linear-gradient(to bottom, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0.4) 100%)',
  },
}
```

---

## 6. Pen

### Palette
```
Primary:     #111111  (true black)
Background:  #FAFAF7  (pure white — near-white, not clinical)
Paper:       #E8E4DC  (paper white)
Gray:        #9C9C9C  (mid gray)
Accent:      #C4381E  (editorial red — extreme restraint, one only)
```

### Typography
- **Logo**: Cormorant Garamond 300, tiny, top-left — deliberate understatement
- **Headline**: binary scale — EITHER dominant (60%+ of frame) OR completely absent
- Swiss grotesque or thin serif
- Strict baseline grid, asymmetric (Müller-Brockmann)

### Layout
- 50-70% of frame is empty — intentional
- Every element aligned to invisible baseline grid
- Swiss asymmetry: left-heavy or right-heavy, not centered
- Single object/concept per cover — no visual complexity
- Museum catalog gravity

### Overlay
- Cover: `linear-gradient(to bottom, rgba(0,0,0,0.05) 0%, rgba(0,0,0,0.35) 100%)` (barely there)
- Content: `linear-gradient(to bottom, rgba(0,0,0,0.2) 0%, rgba(0,0,0,0.65) 100%)`

### Remotion Constants
```typescript
export const PEN = {
  colors: {
    primary: '#111111',
    background: '#FAFAF7',
    paper: '#E8E4DC',
    gray: '#9C9C9C',
    accent: '#C4381E',
  },
  typography: {
    headline: { fontFamily: "'Noto Sans KR', sans-serif", fontWeight: 300, letterSpacing: '-0.01em' },
    body: { fontFamily: "'Noto Sans KR', sans-serif", fontWeight: 300 },
  },
  overlay: {
    cover: 'linear-gradient(to bottom, rgba(0,0,0,0.05) 0%, rgba(0,0,0,0.35) 100%)',
    content: 'linear-gradient(to bottom, rgba(0,0,0,0.2) 0%, rgba(0,0,0,0.65) 100%)',
    closing: 'linear-gradient(to bottom, rgba(0,0,0,0.05) 0%, rgba(0,0,0,0.35) 100%)',
  },
}
```
