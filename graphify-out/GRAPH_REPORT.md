# Graph Report - D:\dev\letcoode\src  (2026-07-30)

## Corpus Check
- 15 files · ~14,224 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 22 nodes · 16 edges · 6 communities detected
- Extraction: 100% EXTRACTED · 0% INFERRED · 0% AMBIGUOUS
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Community 0|Community 0]]
- [[_COMMUNITY_Community 1|Community 1]]
- [[_COMMUNITY_Community 2|Community 2]]
- [[_COMMUNITY_Community 3|Community 3]]
- [[_COMMUNITY_Community 4|Community 4]]
- [[_COMMUNITY_Community 5|Community 5]]

## God Nodes (most connected - your core abstractions)
1. `FlexibleInt` - 2 edges
2. `HoroscopePage` - 1 edges
3. `Place` - 1 edges
4. `Star` - 1 edges
5. `YearStar` - 1 edges
6. `InputData` - 1 edges
7. `form` - 1 edges
8. `formData` - 1 edges
9. `data` - 1 edges
10. `day` - 1 edges

## Toxic Hotspots (high risk + high activity)
1. `client.js` - Risk Score: 80% (FUNCTION)
2. `earth_table.go` - Risk Score: 71% (FUNCTION)

## Surprising Connections (you probably didn't know these)
- None detected - all connections are within the same source files.

## Communities

### Community 0 - "Community 0"

Cohesion: 0.2
Nodes (9): chartOutput, data, day, form, formData, gender, hour, month (+1 more)

### Community 1 - "Community 1"

Cohesion: 0.29
Nodes (5): FlexibleInt, HoroscopePage, Place, Star, YearStar

### Community 2 - "Community 2"

Cohesion: 1.0
Nodes (1): InputData

### Community 3 - "Community 3"

Cohesion: 1.0
Nodes (0): 

### Community 4 - "Community 4"
_Unable to determine domain due to missing code entities._
Cohesion: 1.0
Nodes (0): 

### Community 5 - "Community 5"
_Unable to determine domain due to missing code entities._
Cohesion: 1.0
Nodes (0): 

## Knowledge Gaps
- **14 isolated node(s):** `HoroscopePage`, `Place`, `Star`, `YearStar`, `InputData` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **Thin community `Community 2`** (2 nodes): `input.go`, `InputData`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 3`** (1 nodes): `const.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 4`** (1 nodes): `form_scripts.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 5`** (1 nodes): `imports.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.