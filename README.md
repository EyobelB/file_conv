# file_conv
Multithreaded File Converter Practice for Learning Go

## Current progress:
    1. Created constructors for data frames, metadata, and the flac audio object
    2. Isolated the magic number and displayed it
    3. Created code to display raw binary
    4. Finalized topology of codebase
    5. Finalized the core algorithm of conversion to be implemented

## To-do:
    a. Make existing code OS agnostic (for example, use OS file separators instead of defaulting to forward slashes)
    b. Print out any Vorbis comment strings/metadata
    c. Print out a frame formatted in a readable way (to confirm ability)
    c. Create frames (via multithreading) and store them in the frame array object
    d. Start creating the constructs for MP3 files in mp3.go
    e. Repeat steps a-c for mp3.go
    e. Create a conversion.go file to run the conversion process
    f. Create a main.go file to run the user interface and reference other files
    g. Update the Makefile so that each file is individually compiled, and then linked in final step (to minimize compile time)

## Methodology:
To convert files, the general "algorithm" this converter will use is as follows....
    1. Decode the file into it's bare components via a Struct
    2. Convert said file into a raw format, whether that's image, text, audio, video, etc.
    3. Use decoded information from the original file and the raw data to create the desired format

This methodology allows us to minimize the functions to write, where instead of
creating code for direct conversions between every format, we can simply
follow these three steps for every file format....
    1. Decode
    2. Convert to raw
    3. Convert from raw

This conversion to raw process may be a bit complicated for rasterization/vectorization of images,
    but we'll cross that bridge when we get there.