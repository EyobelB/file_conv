# file_conv
Multithreaded File Converter Practice for Learning Go \

## Current progress:
    1. Created constructors for data frames, metadata, and the flac audio object\
    2. Isolated the magic number and displayed it\

## To-do:
    a. Make existing code OS agnostic (for example, use OS file separators instead of defaulting to forward slashes)\
    b. Print out any Vorbis comment strings/metadata\
    c. Print out a frame formatted in a readable way (to confirm ability)\
    c. Create frames (via multithreading) and store them in the frame array object\
    d. Start creating the constructs for MP3 files in mp3.go\
    e. Repeat steps a-c for mp3.go\
    e. Create a conversion.go file to run the conversion process\
    f. Create a main.go file to run the user interface and reference other files\
    g. Update the Makefile so that each file is individually compiled, and then linked in final step (to minimize compile time)\
