# Graph Report - D:\dev\letcoode\src  (2026-07-26)

## Corpus Check
- 16 files · ~14,086 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 32 nodes · 43 edges · 7 communities detected
- Extraction: 100% EXTRACTED · 0% INFERRED · 0% AMBIGUOUS
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Community 0|Community 0]]
- [[_COMMUNITY_Community 1|Community 1]]
- [[_COMMUNITY_Community 2|Community 2]]
- [[_COMMUNITY_Community 3|Community 3]]
- [[_COMMUNITY_Community 4|Community 4]]
- [[_COMMUNITY_Community 5|Community 5]]
- [[_COMMUNITY_Community 6|Community 6]]

## God Nodes (most connected - your core abstractions)
1. `form` - 2 edges
2. `toggleBtn` - 2 edges
3. `input` - 2 edges
4. `current` - 2 edges
5. `next` - 2 edges
6. `genderBtns` - 2 edges
7. `genderInput` - 2 edges
8. `val` - 2 edges
9. `formData` - 2 edges
10. `data` - 2 edges

## Surprising Connections (you probably didn't know these)
- None detected - all connections are within the same source files.

## Communities

### Community 0 - "Community 0"

Cohesion: 0.2
Nodes (9): chartOutput, day, formData, gender, genderBtns, hour, main, next (+1 more)

### Community 1 - "Community 1"

Cohesion: 0.2
Nodes (9): current, data, existingChart, form, genderInput, input, month, toggleBtn (+1 more)

### Community 2 - "Community 2"

Cohesion: 0.29
Nodes (5): FlexibleInt, HoroscopePage, Place, Star, YearStar

### Community 3 - "Community 3"

Cohesion: 1.0
Nodes (1): InputData

### Community 4 - "Community 4"
_Unable to determine domain due to missing code entities._
Cohesion: 1.0
Nodes (0): 

### Community 5 - "Community 5"
_Unable to determine domain due to missing code entities._
Cohesion: 1.0
Nodes (0): 

### Community 6 - "Community 6"
_Unable to determine domain due to missing code entities._
Cohesion: 1.0
Nodes (0): 

## Knowledge Gaps
- **5 isolated node(s):** `HoroscopePage`, `Place`, `Star`, `YearStar`, `InputData`
  These have ≤1 connection - possible missing edges or undocumented components.
- **Thin community `Community 3`** (2 nodes): `input.go`, `InputData`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 4`** (1 nodes): `const.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 5`** (1 nodes): `form_scripts.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 6`** (1 nodes): `imports.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.