# DG Mascot V1 Image Prompts

## Recommendation First

Yes. Start with **one strong front-facing idle master image first**, then use remix / reference-image generation to derive the other poses.

That is the best V1 strategy if consistency matters.

### Why this is the right order

- it locks the mascot identity early
- it stabilizes silhouette, proportions, face, and color palette
- it reduces the chance that each pose looks like a different creature
- it makes later cleanup easier if you redraw or vectorize by hand

### Best workflow

1. Generate a **single polished master idle image**
2. Pick the best one and lock it as the visual anchor
3. Use that idle image as the reference/remix source for:
   - `wave`
   - `nod`
   - `think`
   - `rest`
   - `spark`
4. After that, clean or redraw the final set into SVG if needed

### What to keep fixed across every remix

- body silhouette
- face shape
- eye size and placement
- spark nub shape
- main palette
- overall camera angle

### What can change across remixes

- eye expression
- mouth shape
- slight body tilt
- accent emphasis
- tiny implied gesture shape

## Global Generation Rules

Use these rules for every image:

- transparent background
- centered composition
- front-facing or near-front-facing
- compact rounded tool spirit
- dark teal body
- warm gold spark nub
- simple clean illustration
- readable at small size
- not humanoid
- not anime
- not corporate mascot
- not hyper-rendered
- no props
- no text
- no scene background

## Global Negative Prompt

Use this negative prompt for every generation:

> humanoid, human body, detailed hands, realistic anatomy, anime character, corporate assistant mascot, heavy armor, robot, mechanical limbs, over-detailed shading, background scene, text, speech bubble, photo realism, horror, creepy eyes, complex accessories, clutter, side profile, extreme perspective

## Master Prompt: Idle Anchor

This is the most important image in the set.

### Prompt

> A small rounded mascot spirit for a desktop learning app, front-facing, shaped like a compact droplet-like tool sprite, dark teal body, warm gold spark nub on top, large readable eyes, tiny expressive mouth, clever and slightly mischievous but warm, calm and observant, simple vector-friendly silhouette, clean flat illustration, transparent background, centered composition, highly readable at small size, not humanoid, not anime, charming and compact.

### Optional style suffix

> soft flat shading, subtle clean outline separation, polished app mascot design, minimal but expressive face

### Goal

This image should define:

- final silhouette
- final face proportions
- final spark nub shape
- final palette
- final vibe

## Prompt 02: Wave

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, keep the exact same silhouette, face proportions, spark nub, and color palette. Show a welcoming wave pose with a slight body tilt, brighter eyes, a small friendly smile, and one side subtly lifted as if waving. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- keep identity nearly identical to idle
- only change expression and pose energy
- do not redesign the character

## Prompt 03: Nod

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, preserve the exact silhouette, spark nub, facial proportions, and colors. Show an approving nod pose with slight downward compression, pleased eyes, and a confident small smile. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- should feel approving, not overexcited
- keep the body compact

## Prompt 04: Think

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, preserve the exact silhouette, spark nub, facial proportions, and palette. Show a thoughtful pose with a slight lean, mildly asymmetrical eyes, a tiny puzzled mouth, and a subtle questioning energy. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- thoughtful, not sad
- curious, not scolding

## Prompt 05: Rest

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, preserve the exact silhouette, spark nub, face proportions, and color palette. Show a calm rest pose with slightly lower posture, softer eyes, a relaxed mouth, and a cozy paused energy. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- should look restful, not sleepy or sad
- keep the silhouette recognizable

## Prompt 06: Spark

Use the chosen idle image as the reference/remix source.

### Prompt

> Same mascot character as the reference image, preserve the exact silhouette, face proportions, spark nub identity, and palette. Show a celebratory spark pose with brighter expression, open posture, energized eyes, a proud smile, and slightly emphasized spark details. Transparent background, centered composition, same camera angle, simple clean flat illustration, readable at small size.

### Extra guidance

- special, but still suitable for a calm productivity app
- can include tiny surrounding spark accents if they stay minimal

## Optional Prompt 07: Collapsed Badge

If the full character reads poorly when the bubble is collapsed, generate a simplified badge version.

### Prompt

> Same mascot character as the reference image, simplified into a compact face-first badge icon, keeping the same eyes, spark nub, and character identity. Very clear and readable at tiny size, transparent background, centered composition, simple clean flat illustration.

## Prompting Notes For Better Consistency

### If your tool supports reference strength

Use:

- medium-high identity preservation
- low-medium pose deviation for most poses
- medium deviation for `spark`

### If your tool supports seed reuse

Reuse the same seed from the chosen idle image for all derivative poses when possible.

### If your tool supports character reference

Always use the selected idle anchor image as the single source of truth.

Do not mix multiple reference images in V1 unless the tool absolutely requires it.

## Evaluation Checklist

After generation, check each pose:

- does it still clearly look like the same mascot?
- does the spark nub stay recognizable?
- do the eyes keep the same identity?
- does the palette remain stable?
- can the emotion be understood at a small size?
- does it still fit the app's calm dark UI?

## Production Advice

Best practical approach:

1. generate `idle` until it is excellent
2. freeze it
3. generate all other poses by remix/reference
4. choose the best variant of each
5. clean them into final SVG assets later

This is safer than generating six unrelated prompts from scratch.
