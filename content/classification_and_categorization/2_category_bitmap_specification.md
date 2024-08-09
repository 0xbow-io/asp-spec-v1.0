# 5.2 Category Bitmap Specification

## 5.2.1 Bitmap Vector

The bitmap is defined as:

$$B = (b_0, b_1, ..., b_i)\,\\ b_i \in \{0, 1\}\\ i \in \lbrack0,255\rbrack $$

Categories are mapped to a 256-bit vector where each bit represents a specific category
based on a predefined schema.
The bitmap can be efficiently stored as a 256-bit integer types or a byte array of size 32 (32 bytes).

Multiple bits indicates multiple categories, i.e.:

- Bits 64-127: Protocol specific Categories
- Bits 128-191: Network Specific Categories
- Bits 192-255: Cross-chain Categories

This compact representation allows for efficient storage and querying of categorized records.

## 5.2.1 Pointers:

Pointers facilitate bitmap to bitmap referencing and enable the creation of complex category structures.
Bitspaces can be reserved for pointers to other bitmaps as long as there is a clear schema for the pointer
and a mapping function between the pointer and the referenced bitmap.

## 5.2.2 Partitions & Reservations:

A parition is a logical grouping of categories within a specific range of bits.
Partitions are used to group categories based on their domain or ownership.

i.e. Bitspace 0-63 can be reserved for `AML` categories, while bitspace 64-127 can be reserved for `KYC` categories.

During the integration phase between an entity and an ASP, the entitity (i.e. a protocol)
can reserve a specific range of bits for it's categories, i.e. Bitspace 64-127 for `Protocol X` categories.
