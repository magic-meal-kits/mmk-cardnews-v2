---
name: cardnews
description: This skill should be used when the user asks to "카드뉴스 만들어줘", "card news", "create card news", "카드뉴스 생성", "make slides", "인스타 카드뉴스", "LinkedIn carousel". Generates Instagram/LinkedIn card news from a topic using AI images and Remotion React rendering.
---

# Card News Generator (카드뉴스)

> Topic → Research → Plan → AI Images → React Card Components → PNG/MP4/PDF

---

## Prerequisites

- **Node.js** 18+ (for Remotion)
- **mmk-cn** CLI (built from this project)
- **Gemini API key** (for AI image generation, optional)

## Output Path

```
{cwd}/{topic-slug}/
├── src/cards/Card01.tsx ~ CardNN.tsx   # React card components
├── src/data.ts                         # Timing + subtitle data
├── public/card-*.png                   # AI-generated images
├── out/stills/card-*.png               # Exported PNG stills
├── out/output.mp4                      # Exported video
└── out/carousel.pdf                    # LinkedIn PDF carousel
```

`SKILL_DIR` = directory where this SKILL.md lives (for references).

---

## Workflow

### Step 0: Environment Check

1. Check Node.js:
```bash
node --version
```

2. Check mmk-cn binary:
```bash
mmk-cn --version
```
If missing: `cd {SKILL_DIR}/../.. && make build`

3. Check Gemini API key (optional):
```bash
echo $GEMINI_API_KEY
```
No key → skip image generation, use text-only cards.

Report available features:
```
환경 확인 완료!
✅ 카드뉴스 생성 (Remotion React)
✅ PNG/PDF 내보내기
✅/❌ AI 이미지 생성 (Gemini)
✅/❌ 영상 렌더링 (MP4)
```

### Step 1: Research

**Type**: prompt (WebSearch)

Use WebSearch to gather 5-10 sources on the topic.

Output: `{output}/research.md`

Structure:
- Topic overview
- 5-10 key points with sources
- Synthesis
- Card news message candidates

**Checkpoint** — AskUserQuestion:
```
"리서치 결과예요. 확인해주세요."
Options: 진행 / 수정 요청 / 여기서 종료
```

### Step 2: Card News Planning

**Type**: prompt + reference

Read research.md + `{SKILL_DIR}/references/card-news-guide.md` to create a plan.

Output: `{output}/plan.md`

Plan includes:
1. Card type (list/storytelling/focus/Q&A)
2. Total slides (7-13)
3. Per-card table: # | Part (cover/body/closing) | Message | Image description | Layout pattern (FullBleed/Split/TextOnly)
4. Color mood based on topic

**Checkpoint** — AskUserQuestion:
```
"기획서예요. 장수나 내용 흐름을 확인해주세요."
Options: 진행 / 수정 요청 / 여기서 종료
```

### Step 3: Scaffold Project

```bash
mmk-cn new "{title}" --slides {N} --output "{output}"
```

Then install dependencies:
```bash
cd "{output}" && npm install
```

This creates the Remotion project with skeleton card components.

### Step 4: AI Image Generation (Optional)

**Requires**: `GEMINI_API_KEY` environment variable or `.env` file.

Model: `gemini-3.1-flash-image-preview` (configurable via `GEMINI_MODEL`).

1. Extract image descriptions from plan.md
2. Create `{output}/batch.json`:
```json
[
  {"prompt": "DO NOT include any text... {english prompt}", "output": "public/card-01.png"},
  {"prompt": "...", "output": "public/card-02.png"}
]
```

3. Run batch generation:
```bash
mmk-cn imagegen --batch "{output}/batch.json" --aspect-ratio 3:4 --image-size 2K --env-file .env
```

**Prompt rules**:
- Start with: "DO NOT include any text, letters, words, numbers, watermarks, logos, or typography in the image. Pure visual only."
- Full color editorial magazine photography
- 200-400 word English prompt

**Checkpoint** — AskUserQuestion:
```
"이미지들이에요. 마음에 안 드는 건 재생성할 수 있어요."
Options: 진행 / 일부 재생성 / 텍스트 온리로 진행
```

### Step 5: Write Card Components

This is the core creative step. Read `{SKILL_DIR}/references/card-template-guide.md` for patterns.

For each card, write `src/cards/CardNN.tsx`:
- Use `FullBleedCard`, `SplitCard`, or `TextOnlyCard` pattern from `CardTemplate.tsx`
- Fill in actual content from plan.md
- Reference images via `staticFile("card-NN.png")`
- Include fade-in/fade-out animations via `interpolate()`
- Page numbers: `N / {total}` at bottom center

**Font size rules (mandatory)**:
| Element | Canvas min | Mobile equiv |
|---------|-----------|-------------|
| Title (h2) | **48px** | ~17px |
| Body (desc) | **36px** | ~12px |
| Date/tag | **32px** | ~11px |
| Info/source | **30px** | ~10px |
| Page number | **28px** | ~10px |
| **Absolute min** | **24px** | ~8px |

Update `src/cards/index.ts` to import all card components.

### Step 6: Preview

```bash
mmk-cn preview "{output}"
```

Opens Remotion Studio in browser for live preview and editing.

**Checkpoint** — AskUserQuestion:
```
"Remotion Studio에서 미리보기를 확인해주세요."
Options: 진행 / 수정 요청
```

### Step 7: Export

Export based on user needs:

**PNG stills (Instagram carousel)**:
```bash
mmk-cn still "{output}"
```

**MP4 video**:
```bash
mmk-cn render "{output}"
```

**PDF carousel (LinkedIn)**:
```bash
mmk-cn pdf "{output}/out/stills"
```

**Checkpoint** — AskUserQuestion:
```
"완성! 확인해주세요."
Options: 완료 / 수정 요청
```

---

## Failure Recovery

Report completed steps clearly:
```
완성된 결과물:
✅ Step 1: 리서치 → research.md
✅ Step 2: 기획서 → plan.md
✅ Step 3: 프로젝트 생성 → {output}/
✅ Step 4: 이미지 → public/card-*.png
✅ Step 5: 카드 컴포넌트 → src/cards/
❌ Step 6-7: (건너뜀)
```

## References

- `references/card-news-guide.md` — Design rules, card types, color themes
- `references/card-template-guide.md` — How to write Card TSX components
- `references/design-system.md` — Typography, spacing, mobile readability

## Settings

| Setting | Default | Override |
|---------|---------|----------|
| Gemini API key | (optional) | `GEMINI_API_KEY` env or `.env` |
| Gemini model | gemini-3.1-flash-image-preview | `GEMINI_MODEL` env or `.env` |
| Image aspect ratio | 3:4 | `--aspect-ratio` flag |
| Image size | 2K | `--image-size` flag |
| Slides | 7 | `--slides` flag |
| Card size | 1080x1350 | Fixed (4:5 portrait) |
| FPS | 30 | Fixed |
