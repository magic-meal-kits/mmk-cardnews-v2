# Japanese Magazine Design Reference
> Card News Generation Guide for Remotion + Google Imagen 3 (Veo)
> Last updated: 2026-04-04

---

## Overview

This document defines the visual design system for each reference magazine.
Use this as a source of truth when prompting Google Imagen 3 for images,
and when building Remotion components for card news.

### Image Generation Platform
- **Platform**: Google Imagen 3 (via Vertex AI / AI Studio)
- **Anti-AI checklist** (append to every image prompt):
  ```
  shot on film, analog grain texture, editorial photography quality,
  no digital smoothing, no HDR, slightly desaturated, magazine scan feel,
  no AI artifacts, natural imperfections
  ```

### Remotion Usage Pattern
Each magazine style maps to a Remotion component variant.
Import the palette and typography constants from this file
to keep visual consistency across scenes.

---

## Magazine Index

| Magazine | Publisher | Founded | Genre | Style Keyword |
|----------|-----------|---------|-------|---------------|
| BRUTUS | Magazine House | 1980 | Pop Culture / Lifestyle | Bold, Intellectual |
| POPEYE | Magazine House | 1976 | Men's Fashion / City Boy | Relaxed Confidence |
| Casa BRUTUS | Magazine House | 2000 | Architecture / Design | Sophisticated Intelligence |
| OLIVE | Magazine House | 1982–2003 | Girl Culture (Vintage) | Handcrafted Intimacy |
| SWITCH | Switch Publishing | 1985 | Music / Film / Art Portrait | Portrait Authority |
| Pen | CCCメディアハウス | 1998 | Art / High Culture | Refined Minimalism |

---

## 1. BRUTUS

### Overall Tone
Bold, intellectual, provocative. Each issue is a completely different visual world —
the concept dictates everything. Cover design functions as a poster or standalone art piece.
Confident typographic aggression mixed with editorial restraint.
The reader is a design-conscious, trend-aware professional male aged 20–50 in Tokyo.

### Color Palette

```
PRIMARY:
  Brutus Red      #D62B2B   ← signature logo color, never changes
  Near Black      #1A1A1A
  Off White       #F5F0E8

ACCENT (per-issue, pick one):
  Warm Orange     #F5A623
  Deep Navy       #1C3A5E
  Mustard         #D4A017
  Forest Green    #2D5A3D

RULE: High contrast always. Base = Black + White + Red.
      Inject one accent per issue concept. Never soft.
```

### Typography

```
LOGO:
  - Custom lettering with jagged spiked strokes (references Brutus character's beard)
  - Always RED (#D62B2B)
  - Position: top-left or top-center, fixed
  - Never changes weight or position across issues

HEADLINE (Cover):
  - Style: ultra-bold condensed grotesque sans-serif
  - Weight: Black / ExtraBold (900)
  - Size: VERY large — 80–120pt equivalent on cover
  - Tracking: tight (-0.02em to -0.04em)
  - Transform: often ALL CAPS
  - Placement: overlaid on image, not separated
  - Axis: horizontal AND vertical mixing is acceptable
  - Font alternatives: Barlow Condensed Black, Anton, Bebas Neue, Impact

SUBHEADLINE:
  - Clean grotesque, medium weight
  - Smaller scale contrast against hero headline
  - Often bilingual (English title + Japanese subtitle)

BODY / INTERIOR:
  - Tategumi (vertical Japanese) mixed with horizontal Latin
  - Drop caps used for section openers
  - Strict column grid inside spreads
  - Tight leading
```

### Layout Pattern

```
COVER STRUCTURE:
  ┌─────────────────────────┐
  │ [BRUTUS]    [Issue No.] │  ← logo top-left, red
  │                         │
  │   FULL BLEED IMAGE      │  ← photo or illustration
  │   (bleeds past frame)   │
  │                         │
  │  HEADLINE OVERLAID      │  ← text on top of image
  │  LARGE & FEARLESS       │
  │                         │
  │  subhead  coverlines    │  ← minimal supporting text
  └─────────────────────────┘

RULES:
  - Text lives ON the image, not in a safe zone
  - Multi-axis text placement (diagonal, vertical allowed)
  - White space is deliberate choice, not default
  - Each issue completely reimagines the visual language
  - Concept-driven: the theme DICTATES the design system
```

### Google Imagen 3 Prompt

```
IMAGEN PROMPT TEMPLATE — BRUTUS STYLE:

[subject description], editorial magazine cover photography,
bold high-contrast lighting, BRUTUS magazine aesthetic,
Kodak Tri-X 400 pushed one stop, deep shadows, strong blacks,
graphic editorial composition, conceptual art direction,
Tokyo design culture, no commercial softness,
shot on film, grain texture, analog imperfections,
color palette: near black and off-white with red accent,
magazine quality, art director approved

NEGATIVE: soft lighting, pastel colors, stock photo, smooth AI skin,
HDR, over-processed, lifestyle commercial, gradient backgrounds
```

### Remotion Component Notes

```typescript
// BRUTUS palette constants
export const BRUTUS = {
  colors: {
    red: '#D62B2B',
    black: '#1A1A1A',
    offWhite: '#F5F0E8',
  },
  typography: {
    headline: { fontFamily: 'Barlow Condensed', fontWeight: 900, letterSpacing: '-0.03em' },
    logo: { fontFamily: 'Anton', color: '#D62B2B' },
  },
  layout: {
    textOverImage: true,          // headline overlaid on photo
    logoPosition: 'top-left',
    gridBreaking: true,           // diagonal/vertical text allowed
  }
}
```

---

## 2. POPEYE

### Overall Tone
Relaxed confidence. The "city boy" who doesn't try too hard.
Post-2012 redesign by studio Shiroi Rittai: radical white space,
clean geometric sans-serif, photography that feels casual but is perfectly composed.
Youthful but never immature. International sensibility with Tokyo DNA.
Target: young men who appreciate culture without performing it.

### Color Palette

```
CORE:
  Warm White      #F7F4EE   ← base background, warm not clinical
  Warm Cream      #E8E0D0
  Film Black      #1A1A1A
  Warm Gray       #B5A89A

ACCENT (sparingly — one per cover):
  City Blue       #4A7C9E
  Warm Red        #D4563A
  Sage Green      #7A9E7E
  Dusty Yellow    #D4B483

RULE: Warm neutral base always. One pop color max. Never cold.
      Film photography palette — slightly desaturated warm shadows.
```

### Typography

```
LOGO:
  - Custom bubble lettering (refined from 1976 original)
  - Similar to Frankfurter Highlight — rounded, friendly
  - Highlights at northeast position
  - Position: top-center, fixed
  - Color: black or reversed white depending on background

HEADLINE (Cover):
  - Style: geometric sans-serif (Futura, Nordvest)
  - Weight: Medium to Bold
  - Case: ALL CAPS
  - Size: medium — not overpowering
  - Tracking: tight to normal
  - Count: 2–4 words maximum on cover
  - Font alternatives: Futura PT, DM Sans, Neue Haas Grotesk

COVER LINES:
  - Extremely few (2–3 maximum)
  - Small point size
  - Placed with extreme care — positioned in quiet zones

WHITE SPACE:
  - Generous margins everywhere
  - Breathing room IS the design statement
  - Text elements are sparse by intention
```

### Layout Pattern

```
COVER STRUCTURE:
  ┌─────────────────────────┐
  │                         │
  │       [POPEYE]          │  ← logo top-center
  │                         │
  │    FULL FRAME PHOTO     │  ← single casual subject
  │    (natural, relaxed)   │
  │                         │
  │   HEADLINE              │  ← minimal, in clear zone
  │                         │
  │  · coverline  · coverline│  ← max 2-3 items
  └─────────────────────────┘

RULES:
  - Photography: single subject, casual energy, film feel
  - Text NEVER fights the image
  - "Unstudied" look is meticulously constructed
  - The "apparent" simplicity IS the sophistication
  - White space proportion: 30–40% of cover
```

### Google Imagen 3 Prompt

```
IMAGEN PROMPT TEMPLATE — POPEYE STYLE:

[subject description], editorial lifestyle photography,
POPEYE magazine city boy aesthetic, Tokyo casual fashion,
Kodak Portra 400, warm color grading, natural available light,
slightly underexposed, warm shadows, analog film quality,
relaxed candid energy but perfectly composed,
clean simple background or urban environment,
warm neutral tones, no commercial studio feel,
shot on film, grain texture, natural imperfections

NEGATIVE: cold tones, studio flash, commercial advertising,
perfect skin, oversaturated, HDR, stock photography feel,
busy composition, multiple subjects fighting for attention
```

### Remotion Component Notes

```typescript
// POPEYE palette constants
export const POPEYE = {
  colors: {
    warmWhite: '#F7F4EE',
    warmCream: '#E8E0D0',
    filmBlack: '#1A1A1A',
    warmGray: '#B5A89A',
  },
  typography: {
    headline: { fontFamily: 'Futura PT', fontWeight: 700, letterSpacing: '0.05em', textTransform: 'uppercase' },
    coverLine: { fontFamily: 'DM Sans', fontWeight: 400, fontSize: '0.7em' },
  },
  layout: {
    textOverImage: false,         // text in clear zones, not over image
    logoPosition: 'top-center',
    whiteSpaceRatio: 0.35,        // 35% of frame is breathing room
    maxCoverLines: 3,
  }
}
```

---

## 3. Casa BRUTUS

### Overall Tone
Sophisticated intelligence. Architecture and design treated as culture, not commerce.
Cover visuals carry the concept with minimal text intervention.
Headers and visuals are the story: "readers should understand the main point without reading every word."
Editor Ko Matsubara: "Key elements are visuals and headers."
Holistic — design + fashion + food + travel as unified lifestyle.

### Color Palette

```
ARCHITECTURAL BASE:
  Architectural White  #F2EDE5   ← warm paper white
  Charcoal             #2C2C2C
  Concrete Blue        #8B9EA8   ← warm concrete gray-blue
  Travertine           #C4B49A   ← stone warm neutral
  Plant Green          #6B8E7A   ← living material
  Blueprint Blue       #1C4E8A   ← technical accent

MATERIAL-WORLD RULE:
  Colors come from physical materials — concrete, stone, glass, wood, plant.
  Never sweet, never synthetic. Every palette reads like a material sample board.
  Per-issue color pulled from the featured architectural space.
```

### Typography

```
LOGO:
  - Clean roman lettering, balanced weight
  - Often reversed white on photography or solid black
  - Scale: smaller than BRUTUS logo — restrained
  - Never colorful — always black or white

HEADLINE (Cover):
  - Style: Mix of roman serif and clean grotesque sans (per issue)
  - Weight: Bold to Black
  - Size: Large, authoritative — commands the composition
  - Placement: clear zone, not buried in image complexity
  - Alignment: left-aligned to strong vertical axis
  - Kerning: generous on display type
  - Font alternatives: Freight Display, Canela, GT Super, Aktiv Grotesk

LAYOUT LOGIC:
  - Strict editorial grid
  - Headers align to vertical axis
  - Body text in clear column structure
  - Photography and type never compete — they share space clearly
```

### Layout Pattern

```
COVER STRUCTURE:
  ┌─────────────────────────┐
  │ [Casa BRUTUS]           │  ← logo top-left, small, white or black
  │                         │
  │  ARCHITECTURAL PHOTO    │  ← interior, building, or design object
  │  (full-bleed, quality)  │
  │                         │
  │                         │
  │ LARGE HEADLINE          │  ← in clear photo zone, not over texture
  │ Bold, authoritative     │
  │                         │
  │ Issue theme  Date       │
  └─────────────────────────┘

RULES:
  - Full-bleed architectural / interior photography
  - Text sits in clear zones — never over busy backgrounds
  - Strict column structure inside
  - Material textures must be visible in photography
  - Design object can be the entire cover (no person required)
```

### Google Imagen 3 Prompt

```
IMAGEN PROMPT TEMPLATE — CASA BRUTUS STYLE:

[architectural/interior/design subject description],
architectural photography, Casa BRUTUS magazine aesthetic,
4x5 large format quality, natural architectural lighting,
high dynamic range with restraint, material textures visible,
concrete, glass, wood, stone surfaces photographed with reverence,
editorial interior design photography, Japanese design culture,
warm but controlled color grading, no lifestyle staging,
shot on technical camera, architectural precision,
color palette: warm whites, concrete grays, travertine, plant green

NEGATIVE: people smiling, lifestyle staging, commercial interiors,
warm consumer photography, stock interior, AI-smoothed textures,
oversaturated colors, harsh flash, Instagram filter aesthetic
```

### Remotion Component Notes

```typescript
// Casa BRUTUS palette constants
export const CASA_BRUTUS = {
  colors: {
    architecturalWhite: '#F2EDE5',
    charcoal: '#2C2C2C',
    concreteBlue: '#8B9EA8',
    travertine: '#C4B49A',
    plantGreen: '#6B8E7A',
    blueprintBlue: '#1C4E8A',
  },
  typography: {
    headline: { fontFamily: 'Freight Display', fontWeight: 700, letterSpacing: '0.01em' },
    logo: { fontFamily: 'Aktiv Grotesk', fontWeight: 400, color: '#FFFFFF' },
  },
  layout: {
    textOverImage: false,
    logoPosition: 'top-left',
    clearZoneForText: true,       // never place text over complex backgrounds
    gridStrict: true,
  }
}
```

---

## 4. OLIVE

### Overall Tone
Charming, intimate, handcrafted. French-influenced, literary, anti-commercial.
Zine energy at magazine scale. The original "オリーブ少女" (Olive Girl) culture —
1982–2003 legacy with cult vintage status.
It felt personal, not polished. Like a letter from a friend with impeccable taste.
Now referenced globally as a benchmark for authentic editorial charm.

### Color Palette

```
VINTAGE ANALOG:
  Olive Green     #7B9E6B   ← signature color, logo color
  Aged Paper      #F0E8D5   ← warm paper base
  Warm Tan        #D4A574
  Skin Beige      #E8D5C4
  Rust            #B5704A
  Warm Brown      #8B7B6B

TONE RULE:
  Everything is slightly faded, warm, and analog.
  Like 35mm film left in the sun.
  Never fully saturated. Always a little nostalgic.
  Textures matter — paper grain, film grain.
```

### Typography

```
LOGO:
  - Italic serif with organic quality — script-adjacent
  - Color: Olive Green (#7B9E6B)
  - Weight: light to medium
  - Feel: feminine without being girlish, literary, personal

HEADLINE:
  - Varied across issues — sometimes handwritten, sometimes serif italic
  - Never aggressive, never bold-heavy
  - Mixed weights feel intentional, not inconsistent
  - Handwritten annotations acceptable
  - Decorative flourishes accepted (unlike other magazines in this list)
  - Font alternatives: Playfair Display Italic, EB Garamond Italic, Cormorant Italic

TEXT FEEL:
  - Like a personal letter
  - Text wraps around images freely (non-grid feel)
  - Cutout text effects, stamps, stickers all in vocabulary
  - Bilingual EN/FR/JP mixing is characteristic
```

### Layout Pattern

```
COVER STRUCTURE:
  ┌─────────────────────────┐
  │ Olive     [date]        │  ← logo italic serif, olive green
  │                         │
  │   PHOTO (warm, film)    │  ← personal, not fashion-editorial
  │   (girl, street, still) │
  │                         │
  │ Headline in italic      │  ← gentle, not commanding
  │ serif or handwritten    │
  │                         │
  │ ·short·items·in·dots·   │  ← charming small cover lines
  └─────────────────────────┘

RULES:
  - Non-mechanical grid — editorial collage feel
  - Images and text in conversation, not separation
  - Warm analog photography always
  - Content is personal, not aspirational
  - French magazine (Milk, Elle) influence visible
```

### Google Imagen 3 Prompt

```
IMAGEN PROMPT TEMPLATE — OLIVE MAGAZINE STYLE:

[subject description], vintage editorial photography,
OLIVE magazine 1990s Japanese girl culture aesthetic,
Fujifilm Superia 200, warm vintage color grading,
expired film warmth and color shift, 35mm grain,
personal snapshot energy, not fashion studio,
French editorial influence, quiet intimate mood,
slightly warm overexposed highlights, soft shadows,
analog imperfections, film borders optional,
color palette: olive green, aged paper, rust, warm tan

NEGATIVE: cold tones, digital precision, modern fashion,
commercial advertising, studio lighting, HDR,
perfect skin, saturated colors, Instagram aesthetic,
professional fashion photography look
```

### Remotion Component Notes

```typescript
// OLIVE palette constants
export const OLIVE = {
  colors: {
    oliveGreen: '#7B9E6B',
    agedPaper: '#F0E8D5',
    warmTan: '#D4A574',
    skinBeige: '#E8D5C4',
    rust: '#B5704A',
    warmBrown: '#8B7B6B',
  },
  typography: {
    headline: { fontFamily: 'Cormorant Garamond', fontStyle: 'italic', fontWeight: 500 },
    logo: { fontFamily: 'EB Garamond', fontStyle: 'italic', color: '#7B9E6B' },
    annotation: { fontFamily: 'Patrick Hand', fontWeight: 400 }, // handwritten feel
  },
  layout: {
    textOverImage: true,          // gentle overlay, not aggressive
    logoPosition: 'top-left',
    gridFreedom: 'high',          // collage/non-mechanical grid allowed
    texture: { grain: true, paper: true },
  }
}
```

---

## 5. SWITCH

### Overall Tone
Portrait-forward. The subject IS the design.
Everything — typography, color, layout — exists to serve the photograph.
Quiet authority. High-quality printing culture.
The cover functions as a gallery print, not a product.
Music, film, and art world figures treated with reverence and space.

### Color Palette

```
PORTRAIT-LED:
  Deep Black      #1A1A1A   ← primary
  Gallery White   #F5F2EC   ← warm white background
  Medium Gray     #6B6B6B
  Warm Mid        #A89880

RULE: Color is dictated entirely by the subject photograph.
      Monochrome (B&W) is common and respected.
      When color, it is filmic and purposeful.
      Text is ALWAYS white or black for legibility — never colored.
      No fixed accent color system — each issue follows the subject.
```

### Typography

```
LOGO:
  - Clean geometric sans-serif
  - Often white, reversed on photography
  - Scale: small — completely non-intrusive
  - Top edge of cover
  - Intentionally yields to the photograph

HEADLINE / SUBJECT NAME:
  - Clean sans-serif, medium weight
  - Subject name only — almost nothing else on the cover
  - Scale: moderate — does not compete with face
  - Color: white or black only
  - Font alternatives: Neue Haas Grotesk, Helvetica Neue, GT America

TEXT HIERARCHY (unbreakable rule):
  1. Portrait photograph
  2. Subject name
  3. Magazine title (logo)
  This order NEVER reverses.

TEXT DENSITY: Near zero.
  "If you can remove a word, remove it."
```

### Layout Pattern

```
COVER STRUCTURE:
  ┌─────────────────────────┐
  │ SWITCH                  │  ← logo: small, top, white
  │                         │
  │                         │
  │  FULL FRAME PORTRAIT    │  ← subject fills the frame
  │  (face, gaze, respect)  │
  │                         │
  │                         │
  │           Name          │  ← subject name only, minimal
  │                         │
  └─────────────────────────┘

RULES:
  - Photography fills the entire frame
  - Large empty space is correct and intentional
  - Never add elements that distract from the face
  - Eye contact with camera is preferred but not required
  - Controlled lighting — studio OR environmental, never snapshot
  - Medium to large format portrait quality
```

### Google Imagen 3 Prompt

```
IMAGEN PROMPT TEMPLATE — SWITCH MAGAZINE STYLE:

[subject person description], portrait editorial photography,
SWITCH magazine art direction, gallery print quality,
Hasselblad medium format portrait aesthetic,
controlled natural window light or studio,
shallow depth of field, subject is sole focus,
respectful intimate framing, gaze direct or thoughtful,
monochromatic or filmic color grading,
grain texture subtle, editorial portrait magazine quality,
no props competing with subject, clean simple background

NEGATIVE: busy composition, multiple subjects, candid snapshot,
commercial portrait, beauty retouching, dramatic makeup,
celebrity glam, wide angle distortion, action shots,
any background that competes with the subject's face
```

### Remotion Component Notes

```typescript
// SWITCH palette constants
export const SWITCH = {
  colors: {
    deepBlack: '#1A1A1A',
    galleryWhite: '#F5F2EC',
    medGray: '#6B6B6B',
    warmMid: '#A89880',
    // NOTE: accent colors derived from subject photo per issue
  },
  typography: {
    subjectName: { fontFamily: 'Neue Haas Grotesk', fontWeight: 400, color: '#FFFFFF' },
    logo: { fontFamily: 'GT America', fontWeight: 300, color: '#FFFFFF', fontSize: '0.6em' },
  },
  layout: {
    textOverImage: true,          // but minimal — name only
    logoPosition: 'top-left',
    maxTextElements: 2,           // logo + name. nothing else.
    portraitDominance: 1.0,       // 100% of frame is the photo
  }
}
```

---

## 6. Pen

### Overall Tone
Refined intelligence. Minimalism as philosophy, not trend.
Clean negative space used as a design element itself.
Swiss International Style typography meets Japanese negative space philosophy.
The cover breathes — 50–70% empty space is not a flaw, it is the design.
High culture subjects treated with museum-catalog gravity.

### Color Palette

```
SWISS MINIMAL:
  Pure White      #FAFAF7   ← near-white, not clinical
  True Black      #111111
  Paper White     #E8E4DC
  Mid Gray        #9C9C9C

ACCENT (one only, used with extreme restraint):
  Editorial Red   #C4381E   ← Swiss design red
  OR: single color pulled from the featured artwork

RULE: Black and white IS the system.
      One accent color maximum.
      Inspired by Swiss International Style (Helvetica era).
      If in doubt, remove the color.
```

### Typography

```
LOGO:
  - Thin lowercase serif — "Pen" — deliberate understatement
  - Position: top-left, very small
  - The restraint of the logo signals the magazine's confidence
  - Never bold, never shouting
  - Font: thin serif (Cormorant, Didot Light, Bodoni Light)

HEADLINE:
  - Binary scale choice: EITHER dominant (60%+ of cover frame)
    OR completely absent — no middle ground
  - When dominant: fills the frame, IS the cover
  - When absent: image or object alone carries the cover
  - Swiss grotesque or thin serif
  - Strict baseline grid alignment
  - Font alternatives: Helvetica Neue Light, Aktiv Grotesk Light,
    Apercu, Neue Haas Grotesk

GRID:
  - Swiss baseline grid — every element precisely positioned
  - Nothing arbitrary or decorative
  - Asymmetric layouts (not centered) — Müller-Brockmann influence
  - Negative space is primary, content is secondary
```

### Layout Pattern

```
COVER STRUCTURE — TYPE DOMINANT:
  ┌─────────────────────────┐
  │ pen                     │  ← logo: tiny, top-left, thin serif
  │                         │
  │                         │
  │  VERY LARGE             │  ← headline dominates (60%+ of frame)
  │  HEADLINE               │
  │  HERE                   │
  │                         │
  │            small image  │  ← isolated object, generous space
  └─────────────────────────┘

COVER STRUCTURE — IMAGE DOMINANT:
  ┌─────────────────────────┐
  │ pen                     │  ← logo: tiny, top-left
  │                         │
  │                         │
  │      [OBJECT]           │  ← isolated design/art object
  │      centered with      │
  │      massive margin     │
  │                         │
  │ Small text  Date        │  ← minimal bottom text
  └─────────────────────────┘

RULES:
  - 50–70% of the cover is empty — this is intentional
  - Every element aligned to invisible baseline grid
  - Swiss asymmetry: left-heavy or right-heavy, not centered
  - Single object or concept per cover — no visual complexity
  - Museum catalog gravity in every composition
```

### Google Imagen 3 Prompt

```
IMAGEN PROMPT TEMPLATE — PEN MAGAZINE STYLE:

[subject description: design object / artwork / single concept],
Pen magazine minimal editorial aesthetic,
Swiss design influence, museum photography quality,
object isolated on clean neutral background,
controlled studio lighting, even and precise,
product photography with editorial dignity,
no distracting context, pure subject focus,
color palette: warm white, neutral gray;
accent color only if part of the object itself,
technical photography precision,
shot on technical camera, no grain preferred (crisp),
archival quality

NEGATIVE: busy backgrounds, lifestyle context, people in frame,
warm vintage look, film grain, casual photography,
multiple objects competing, decorative props,
any visual complexity that reduces breathing room
```

### Remotion Component Notes

```typescript
// PEN palette constants
export const PEN = {
  colors: {
    pureWhite: '#FAFAF7',
    trueBlack: '#111111',
    paperWhite: '#E8E4DC',
    midGray: '#9C9C9C',
    editorialRed: '#C4381E', // use sparingly
  },
  typography: {
    headline: { fontFamily: 'Helvetica Neue', fontWeight: 300, letterSpacing: '-0.01em' },
    logo: { fontFamily: 'Cormorant Garamond', fontWeight: 300, fontSize: '0.55em', color: '#111111' },
  },
  layout: {
    textOverImage: false,
    logoPosition: 'top-left',
    negativeSpaceRatio: 0.6,      // 60% empty space target
    gridSystem: 'swiss-baseline',
    headlineScale: 'binary',      // either huge or absent — no middle
  }
}
```

---

## Quick Reference: Choosing a Magazine Style

```
CONTENT TYPE          → RECOMMENDED STYLE
──────────────────────────────────────────────────────
Concept / Theme piece → BRUTUS    (bold, concept-driven)
City lifestyle / food → POPEYE    (casual, warm, city)
Architecture / space  → Casa BRUTUS (material, precise)
Personal / nostalgic  → OLIVE     (warm, handcrafted, vintage)
Person / interview    → SWITCH    (portrait-first, minimal)
Art / high culture    → Pen       (minimal, swiss, breath)
```

---

## Remotion Integration Pattern

```typescript
// magazine-styles.ts
import { BRUTUS, POPEYE, CASA_BRUTUS, OLIVE, SWITCH, PEN } from './magazine-palettes'

export type MagazineStyle = 'brutus' | 'popeye' | 'casa' | 'olive' | 'switch' | 'pen'

export const MAGAZINE_STYLES = {
  brutus: BRUTUS,
  popeye: POPEYE,
  casa: CASA_BRUTUS,
  olive: OLIVE,
  switch: SWITCH,
  pen: PEN,
}

// Usage in Remotion component:
// const style = MAGAZINE_STYLES[props.magazineStyle]
// Apply style.colors, style.typography, style.layout
```

---

## Image Prompt Workflow (Imagen 3)

```
1. Pick magazine style based on content type
2. Use the magazine's IMAGEN PROMPT TEMPLATE from above
3. Replace [subject description] with your actual content
4. ALWAYS append the anti-AI base prompt:
   "shot on film, analog grain texture, editorial photography quality,
    no digital smoothing, no HDR, slightly desaturated, magazine scan feel,
    no AI artifacts, natural imperfections"
5. For PEN style: omit the anti-AI grain prompt (PEN prefers crisp technical)

POST-PROCESSING in Remotion:
  - Apply color grade via CSS filter matching the palette
  - Overlay subtle grain texture (PNG overlay at 5–8% opacity) for analog feel
  - Use mix-blend-mode: multiply for color tinting
```

---

*Reference compiled for Claude Code + Remotion card news pipeline*
*Magazine design analysis based on Typeroom, Fonts In Use, and design archives*

---

## Art Direction Prompting — Professional Composition Layer

> The difference between "a nice photo" and "a magazine-ready image" is **intent**.
> A professional art director doesn't describe what to photograph.
> They describe **how the image must function within a layout** before the shot is taken.
> Apply this layer on top of every magazine-specific prompt above.

---

### Core Principle: Design-First Image Generation

```
Amateur prompt:
  "a cup of coffee on a table"

Art directed prompt:
  "overhead composition, coffee cup positioned lower-left third,
  upper-right area intentionally empty for headline text overlay,
  surface texture visible for tactile depth, warm shadow falls left,
  color palette cream and deep brown, negative space is deliberate"
```

The image must already "know" where the text will live.

---

### Compositional Intent Framework

Every image prompt should answer these 4 questions before describing the subject:

```
1. TEXT ZONE — Where will the headline sit?
   → "upper third clear and uncluttered"
   → "left half open for text column"
   → "bottom quarter reserved for caption zone"

2. VISUAL ENTRY POINT — Where does the eye enter first?
   → "subject positioned at right third, eye directed left into open space"
   → "central subject with radial negative space outward"
   → "foreground element leads eye to mid-ground focus"

3. COLOR ZONE SEPARATION — Where are light vs. dark areas?
   → "bright upper area transitions to darker lower third"
   → "left side high-key, right side in shadow — natural text zone left"
   → "background deliberately underexposed for text legibility"

4. DEPTH LAYER — How many visual planes?
   → "sharp foreground, soft background — subject isolated"
   → "flat composition, no depth — graphic poster quality"
   → "two planes only: subject + background, no mid-ground clutter"
```

---

### Text Zone Templates

Use these phrases verbatim in Imagen 3 prompts:

```
TOP HEADLINE ZONE (Brutus / Casa-style):
  "upper 35% of frame intentionally minimal and low-detail,
  subject mass concentrated in lower two-thirds,
  bright or neutral upper area suitable for dark text overlay"

LEFT COLUMN ZONE (editorial spread style):
  "subject fills right half of frame,
  left half is clean negative space or shallow background,
  horizontal composition, landscape orientation implied"

BOTTOM CAPTION ZONE (Popeye / casual style):
  "subject fills upper 70% naturally,
  lower portion fades to neutral or shadow,
  bottom band suitable for small text"

CENTER STAGE — TEXT SURROUNDS (Pen / Switch style):
  "subject isolated center-frame with generous equal margins all sides,
  subject occupies maximum 40% of total frame area,
  surrounding space is clean, low-texture, single-tone"

FULL BLEED — TEXT ON TOP (Brutus aggressive style):
  "high contrast image — ensure areas of both deep shadow (for white text)
  and bright highlight (for black text exist in frame),
  avoid mid-tone flatness — push contrast so text has a landing zone
  regardless of placement position"
```

---

### Lighting Directives (by magazine style)

Lighting is art direction. Specify like a cinematographer.

```
BRUTUS — Dramatic Editorial:
  "single directional hard light source, cinematic contrast ratio,
  practical light motivation (window / lamp), deep unlit blacks,
  no fill light — let shadows fall dark, chiaroscuro quality,
  high contrast ratio 1:8 or deeper"

POPEYE — Casual Natural:
  "overcast diffused window light, soft even shadows, warm color temp 3200K,
  slight lens flare acceptable, no visible studio setup,
  feels like afternoon in someone's apartment or on a street,
  no artificial fill, what you see is what's there"

CASA BRUTUS — Architectural Precision:
  "even ambient light, no harsh directional shadows,
  all surfaces rendered with equal material clarity,
  light reveals texture not mood, color-accurate 5500K,
  as if photographed for an architecture firm's own portfolio"

OLIVE — Intimate Analog:
  "warm golden hour or indoor lamp quality, slightly overexposed highlights,
  shadows are warm not cool (shadow color temp 3000K),
  light wraps around subject softly, feels like the photographer is a friend,
  accidental quality — not studio-arranged"

SWITCH — Portrait Controlled:
  "single large soft source: north window or large beauty dish,
  even modeling light on face, deliberate shadow on one side,
  background falls 1-2 stops darker than subject,
  Rembrandt or split lighting — clear intentional craft"

PEN — Technical Studio:
  "perfectly even diffused light, zero visible shadows or hot spots,
  all surfaces equally illuminated, product photography precision,
  color-neutral daylight 5500K, absolute formal clarity of form"
```

---

### Camera & Lens Directives

```
ENVIRONMENTAL / WIDE:
  "shot on 35mm lens, moderate depth of field,
  slight environmental context visible but not dominant,
  natural perspective, no compression"

INTIMATE / STANDARD:
  "shot on 50mm lens, natural human-eye perspective,
  honest rendering, no distortion, what you see is what it is"

PORTRAIT / COMPRESSION:
  "shot on 85mm lens, background compression, subject isolated,
  background rendered soft, subject sharp, pleasing ratio"

DETAIL / TEXTURE:
  "macro focus, 100mm equivalent, extreme surface detail,
  shallow depth of field, background dissolved to soft tone"

ARCHITECTURAL:
  "tilt-shift 90mm equivalent, verticals perfectly straight,
  no keystoning, no perspective distortion, technical precision"

GRAPHIC / OVERHEAD:
  "directly overhead, flat-lay 90 degrees, no perspective,
  pure graphic quality, shadows short and symmetrical or eliminated"
```

---

### Color Pre-Visualization

Force the palette BEFORE Imagen generates:

```
BRUTUS HIGH CONTRAST:
  "desaturate all midtones, push blacks to near #1A1A1A,
  highlights to warm #F5F0E8, primary subject retains ONE saturated
  color as accent, everything else near-monochrome"

POPEYE WARM FILM:
  "Kodak Portra 400 color science: shadows shift warm/amber,
  highlights stay neutral-warm, skin tones golden,
  global -15% saturation from camera-accurate,
  no pure whites — all whites carry warm cream cast"

CASA BRUTUS MATERIAL:
  "palette from physical materials only: concrete gray, travertine warm stone,
  plant green, glass blue — no synthetic colors, no skin tones,
  no fashion colors, only what the architecture is made of"

OLIVE VINTAGE:
  "cross-process color shift: greens shift yellow-green,
  blues shift cyan, shadows warm orange-brown,
  overall saturation -25%, contrast -10%,
  Fujifilm slide film from 1994"

SWITCH MONOCHROME:
  "near-monochrome or full B&W: retain one hue channel maximum,
  tonal range from rich deep shadow to open bright highlight,
  no color noise, silver gelatin print quality"

PEN CLINICAL NEUTRAL:
  "camera-accurate, zero grading shift, true neutral midtones
  (no warm or cool cast), only the subject's inherent color visible,
  background: white #FAFAF7, gray #9C9C9C, black #111111"
```

---

### Full Art-Directed Prompt Structure

Combine all layers in this order:

```
[1. COMPOSITION INTENT]        ← layout function first
[2. TEXT ZONE DECLARATION]     ← where text will live
[3. LIGHTING DIRECTIVE]        ← how light behaves
[4. CAMERA / LENS]             ← focal length and perspective
[5. COLOR PRE-VISUALIZATION]   ← palette enforcement
[6. SUBJECT DESCRIPTION]       ← what is actually in the image
[7. MAGAZINE STYLE REFERENCE]  ← aesthetic DNA
[8. ANTI-AI BASE PROMPT]       ← analog/film quality layer
[9. NEGATIVE PROMPT]           ← explicit exclusions
```

#### Example — BRUTUS style / Card news: Tokyo coffee culture

```
Overhead flat-lay composition, strong diagonal tension in subject placement,
upper-left quadrant intentionally empty — reserved for headline text overlay,
subject mass pulls eye from upper-left to lower-right,
[LIGHTING] single hard directional light from upper-right, deep shadow lower-left,
push contrast: no mid-tone flatness, blacks deep, highlights punchy,
[CAMERA] 50mm overhead perspective, graphic quality, no depth distortion,
[COLOR] near-monochrome: coffee brown is the one saturated element,
everything else neutral: linen surface, dark shadow, ceramic white,
[SUBJECT] artisan ceramic coffee cup, crema surface with pattern, handmade texture,
few coffee grounds scattered deliberately, Japanese craft object quality,
[STYLE] BRUTUS magazine art direction, bold editorial graphic quality,
Tokyo design culture reference, concept-driven composition,
shot on film, editorial photography, analog grain subtle,
magazine scan quality, no digital smoothing,
NEGATIVE: lifestyle warmth, cozy morning feel, soft bokeh sentimentality,
Instagram aesthetic, stock photography composition,
centered symmetrical layout, all elements equal visual weight
```

#### Example — POPEYE style / Card news: Tokyo cycling

```
3/4 environmental portrait, subject positioned right-third of frame,
body language facing left — eye leads into open space on left,
left half of frame: open sky or clean building wall — reserved for text column,
[LIGHTING] afternoon overcast daylight, soft even shadows,
warm color temperature 3200K, no artificial light setup visible,
[CAMERA] 35mm, slight environmental context in frame, natural perspective,
[COLOR] Kodak Portra warm grade: asphalt grays shift slightly warm-amber,
subject clothing retains color, background desaturates naturally,
[SUBJECT] young man with bicycle, casual unstudied posture,
Tokyo street or park environment, architectural background, caught mid-movement,
not posed, real moment quality,
[STYLE] POPEYE magazine city boy editorial, relaxed confidence,
Shiroi Rittai studio art direction, Tokyo 2020s,
shot on 35mm film, natural grain, analog imperfections, no HDR,
NEGATIVE: fashion editorial posing, studio backdrop, smiling at camera,
commercial advertising, perfect grooming, stock photo energy,
busy distracting background, tourist photograph feel, overly saturated
```

---

### Photographer References by Magazine Style

Use as Imagen 3 style anchors — each carries strong compositional vocabulary:

```
BRUTUS:
  Daido Moriyama      — high contrast grain, Tokyo street, raw energy
  Issei Suda          — conceptual portrait, dark surreal quality
  Wolfgang Tillmans   — editorial conceptual, non-hierarchical composition

POPEYE:
  Rinko Kawauchi      — warm intimate, natural light, quiet moments
  Takashi Homma       — casual suburban Tokyo, honest film quality
  Jack Davison        — editorial portrait, playful wit, London/Tokyo aesthetic

CASA BRUTUS:
  Hiroshi Sugimoto    — architectural precision, long exposure, material focus
  Iwan Baan           — architecture in human scale, environmental context
  Hélène Binet        — material light, architectural shadow, texture reverence

OLIVE:
  Nan Goldin          — intimate snapshot, warm film, personal urgency
  Nobuyoshi Araki     — personal intimate Tokyo life, analog warmth
  Viviane Sassen      — warm color, poetic composition, soft femininity

SWITCH:
  Annie Leibovitz     — controlled portrait, subject reverence, narrative
  Peter Lindbergh     — B&W portrait, editorial gravity, anti-retouching
  Juergen Teller      — raw direct portrait, no retouching, confrontational honesty

PEN:
  Hiroshi Sugimoto    — minimal conceptual, precise formal study
  Wolfgang Tillmans   — object study, technical clarity, flat presentation
  Thomas Ruff         — neutral portrait / object study, zero affect
```

#### Usage syntax in Imagen 3:

```
"in the photographic style of Rinko Kawauchi,
warm intimate natural light, quiet compositional restraint"

→ Do NOT just write "Rinko Kawauchi style" — add 1-2 visual descriptors
  that explain WHY you're citing them. Imagen 3 needs the anchor + the reason.
```

---

### Imagen 3 Specific Syntax Notes

```
Responds well to:
  ✓ Specific color hex values: "background color #F5F0E8"
  ✓ Photography technical terms: f/2.8, ISO 400, 35mm, medium format
  ✓ Named film stocks: Kodak Portra 400, Fuji Superia 200, Ilford HP5
  ✓ Named photographers + visual descriptor (see above)
  ✓ Print medium reference: "as it would appear in a printed magazine"
  ✓ Ratio language: "subject occupies 40% of frame, 60% is negative space"
  ✓ Zone language: "upper-left reserved", "right third is subject zone"

Avoid:
  ✗ Abstract mood words without visual anchor: "melancholy", "hopeful"
  ✗ More than 3 photographer references (dilutes coherence)
  ✗ Contradictory directives: warm AND cold simultaneously
  ✗ Text-in-image requests (Imagen 3 handles text unreliably)
  ✗ Vague style words without visual consequence: "aesthetic", "vibe"
```

---

*Art direction layer added 2026-04-04*
*For use with Google Imagen 3 via Vertex AI / AI Studio*
