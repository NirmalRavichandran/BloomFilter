# Bloom Filter – Design Notes

## 1. Problem Statement
Why a Bloom Filter exists and what problem it solves.

## 2. Guarantees & Non-Guarantees
- No false negatives
- Possible false positives
- No deletions (classic BF)

## 3. Core Components
### 3.1 Bit Array
- Size (m)
- Representation choice
- Memory considerations

### 3.2 Hash Functions
- Number of hashes (k)
- Hash strategy
- Uniformity assumptions

## 4. Operations
### Insert
### Query

## 5. Probability Model
- False positive derivation
- Relationship between m, n, k

## 6. Design Decisions
- Why fixed size
- Why chosen hash approach
- Why no dynamic resizing

## 7. Limitations
- Saturation
- No deletions
- False positives increase over time

## 8. Possible Extensions
- Counting Bloom Filter
- Scalable Bloom Filter
- Thread-safe variant
