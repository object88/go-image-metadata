# JPEG encoding

A JPEG consists of a series of `marker segments`, or just `segments`.  Each `marker segment` has a specific purpose, mostly to do with the metadata, encoding, and encoded image data.

For the purpose of gathering JPEG metadata, most `marker segments` are not interesting.

The first segment is the header, and contains the byte sequence that indicates this is a JPEG, `[0xff, 0xd8]`.
