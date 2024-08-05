# 5.2 Category Bitmap Specification

The category bitmap is a 256-bit vector where each bit represents a specific category. This compact representation allows for efficient storage and querying of categorized records.

The bitmap is defined as:

$$B = (b_0, b_1, ..., b_i)\,\\ b_i \in \{0, 1\}\\ i \in \lbrack0,255\rbrack $$

Categories are assigned to specific bits based on a predefined schema.
Multiple bits indicates multiple categories.

For example:

- Bits 0-63: Transaction types
- Bits 64-127: Value ranges
- Bits 128-191: Protocol-specific categories
- Bits 192-255: Cross-protocol interactions

The bitmap can be efficiently stored and manipulated using 256-bit integer types or a byte array of size 32 (32 bytes).
