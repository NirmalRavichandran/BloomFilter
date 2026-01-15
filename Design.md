# Bloom Filter – Design Notes

## 1. Problem Statement
Why a Bloom Filter exists and what problem it solves.
- Bloom filter is a probabilistic data structure , rapidly fast , memory efficient and tells you that whether an element is present in the set or not .
- It solves the problem for applications where the amount of source data would require an impractical amount of memory .

## 2. Guarantees & Non-Guarantees
- No false negatives
- Possible false positives
- No deletions (Deletion of set bits can affect other elements also)

## 3. Core Components
### 3.1 Bit Array
- Size (m) : A larger bit array will have lesser false positives .
- Representation choice
- Memory considerations

### 3.2 Hash Functions
- Number of hashes (k) : Too many hash functions will slow down the filter ; Too few will have too many false positives .
- Hash strategy
- Uniformity assumptions 

## 4. Operations
### Insert
- Run the element through the k hash functions and set the bits of those resultant bits . O(k) Time complexity .
### Query
- Again run the element through the k hash functions and if all the resultant bits are set then it may be there ; if even a single bit is off , the element is not there . 

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
