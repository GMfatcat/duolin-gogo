# AI Card Authoring Prompt

## Purpose

This file defines a reusable prompt for generating `duolin-gogo` knowledge cards with an LLM.

The goal is:

- generate Markdown cards that already match our schema
- reduce cleanup work after AI generation
- keep bilingual content and hook metadata consistent
- make AI output safe to validate through the existing importer

## Recommended Workflow

1. give the model one topic or one command at a time
2. ask for exactly one card per response
3. run `Validate knowledge` after saving the file
4. review the generated card manually before trusting it

Do not ask the model to generate many loosely related cards in one response unless you are ready to review each card carefully.

## Prompt Template

Use the following prompt as the base authoring prompt for any LLM.

```text
You are generating one Markdown study card for duolin-gogo.

Output requirements:
- Return only the final Markdown card.
- Do not add explanations outside the card.
- The card must be valid Markdown with YAML frontmatter.
- The body must contain exactly two sections:
  - ## zh-TW
  - ## en

Schema requirements:
- One file = one card.
- Required fields:
  - id
  - title_zh
  - title_en
  - type
  - answer
- Strongly recommended bilingual fields:
  - question_zh
  - question_en
  - choices_zh
  - choices_en
  - clickbait_zh
  - clickbait_en
  - review_hint_zh
  - review_hint_en
- Optional but preferred metadata:
  - tags
  - difficulty
  - confusion_with
  - metaphor_seed
  - hook_style_tags
  - enabled

Allowed question types:
- single-choice
- true-false

Rules:
- If type is single-choice:
  - provide 2 to 4 answer choices
  - answer must be a zero-based integer index
- If type is true-false:
  - answer must be true or false
- id must be lowercase kebab-case and Git-topic specific
- tags should be lowercase
- enabled should be true
- body_format should be bilingual-section
- The zh-TW explanation should be concise, clear Traditional Chinese.
- The en explanation should be concise, clear English.
- The zh-TW and en sections must explain the same concept.
- The card should focus on one concrete Git concept only.
- The question should test conceptual understanding, not trivia.
- The wrong choices should be plausible but clearly wrong.
- clickbait should create curiosity without becoming unrelated nonsense.
- review_hint should be short and memorable.

Quality bar:
- Prefer practical developer knowledge.
- Avoid vague wording.
- Avoid unsupported claims.
- Do not invent flags or behaviors.
- Keep explanations readable in under 90 seconds.

Now generate exactly one card for this topic:
TOPIC_PLACEHOLDER
```

## Example Topic Inputs

Good inputs:

- `git squash merge`
- `git reflog`
- `difference between git merge and git rebase`
- `what git remote set-url does`
- `how git stash pop behaves`

Less ideal inputs:

- `teach me git`
- `all about branching`
- `advanced git stuff`

The more specific the topic, the better the generated card will be.

## Recommended Follow-Up Prompt

After the model produces a card, you can use a second pass like this:

```text
Review the following duolin-gogo card for schema compliance and concept accuracy.

Check:
- frontmatter completeness
- bilingual consistency between zh-TW and en
- whether answer matches the choices
- whether confusion_with is reasonable
- whether clickbait is interesting but still relevant
- whether review_hint is short and useful

If the card is good, return the corrected final Markdown only.
If the card has problems, fix them and return the corrected final Markdown only.
Do not add commentary outside the Markdown.
```

## Card Writing Heuristics

Use these heuristics when prompting or reviewing AI output:

- one card should teach one mental model
- prefer command intent over memorizing syntax
- prefer common mistakes over obscure edge cases
- if a flag is dangerous, emphasize consequence in `clickbait` or `review_hint`
- if two commands are commonly confused, use `confusion_with`
- if a metaphor helps, keep `metaphor_seed` short and concrete

## Good Metadata Patterns

### Safe comparison card

```yaml
confusion_with: [git-fetch-basic, git-pull-composition]
metaphor_seed: [先看, 再決定, 觀察]
hook_style_tags: [comparison, safer-first, misunderstood]
```

### Risky command card

```yaml
confusion_with: [git-revert-safe-undo]
metaphor_seed: [後悔藥, 倒帶, 危險]
hook_style_tags: [fear_of_mistake, consequence, misunderstood]
```

### History inspection card

```yaml
metaphor_seed: [時間軸, 歷史書, 線索]
hook_style_tags: [serious, safer-first]
```

## Acceptance Checklist

Before saving an AI-generated card:

- `id` looks stable and unique
- `type` is valid
- `answer` matches the question format
- bilingual sections both exist
- zh-TW is actually Traditional Chinese
- explanations are not excessively long
- `clickbait` is relevant to the concept
- `review_hint` is short
- `Validate knowledge` passes without errors
- any warnings are intentional and acceptable

## Recommendation

For now, use AI as:

- a first-draft generator
- a structure filler
- a phrasing assistant

Do not use AI as the only accuracy check.
Always run validation and do a quick human review before treating generated cards as trusted study material.
