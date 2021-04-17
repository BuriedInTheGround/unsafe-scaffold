#include <errno.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

int main() {
    FILE *in = fopen("UTF8.data", "rb");
    FILE *out = fopen("output.data", "wb");
    if (in == NULL) {
        perror("fopen() for input file failed");
        exit(EXIT_FAILURE);
    }
    if (out == NULL) {
        perror("fopen() for output file failed");
        fclose(in);
        exit(EXIT_FAILURE);
    }

    uint8_t buf_s; // One-byte buffer.

    size_t ret = fread(&buf_s, sizeof(buf_s), 1, in); // Read 1-byte of data.
    size_t idx = 1;
    while (!feof(in)) {
        // Check that the expected number of bytes was read.
        if (ret != 1) {
            fprintf(stderr, "fread() failed: %zu\n", ret);
            fclose(in);
            fclose(out);
            exit(EXIT_FAILURE);
        }

        // Bring the seek cursor back by 1 byte.
        int e = fseek(in, -1L, SEEK_CUR);
        if (e) {
            perror("fseek() failed");
            fclose(in);
            fclose(out);
            exit(EXIT_FAILURE);
        }

        // Evaluate the number of bytes of the UTF-8 codeword.
        size_t utf8bytes = 1;
        if ((buf_s ^ 254) == 0) {
            utf8bytes = 7;
        } else if ((buf_s ^ 252) < 1<<1) {
            utf8bytes = 6;
        } else if ((buf_s ^ 248) < 1<<2) {
            utf8bytes = 5;
        } else if ((buf_s ^ 240) < 1<<3) {
            utf8bytes = 4;
        } else if ((buf_s ^ 224) < 1<<4) {
            utf8bytes = 3;
        } else if ((buf_s ^ 192) < 1<<5) {
            utf8bytes = 2;
        }

        // Buffer for reading the codeword from the input file, utf8bytes-bytes
        // at a time.
        uint8_t buf[utf8bytes];

        ret = fread(buf, sizeof(buf), 1, in); // Read utf8bytes-bytes of data.
        if (ret != 1) {
            fprintf(stderr, "fread() failed: %zu\n", ret);
            fclose(in);
            fclose(out);
            exit(EXIT_FAILURE);
        }

        uint32_t codepoint = 0; // Initialize the codepoint.

        // Discern how to decode based on the number of UTF-8 bytes.
        if (utf8bytes > 1) {
            // Put the least significant (7-utf8bytes) bits of the first UTF-8
            // byte into the codepoint.
            codepoint |= buf[0] & (254 >> (utf8bytes+1));
            // Left shift the codepoint by 6 positions to make room for the
            // next 6 bits from the UTF-8 codeword.
            codepoint <<= 6;
            for (size_t i = 1; i < utf8bytes-1; i++) {
                // Check that the byte starts with the 2-bit string 10. The
                // UTF-8 codeword is otherwise invalid.
                if ((buf[i] & (3 << 6)) != (2 << 6)) {
                    fprintf(stderr, "bad utf-8 codeword #%zu, byte #%zu: %#02x\n", idx, i+1, buf[i]);
                    fclose(in);
                    fclose(out);
                    exit(EXIT_FAILURE);
                }
                // Put the least significant 6 bits of the i-th UTF-8 byte into
                // the codepoint and do the left shift by 6 positions.
                codepoint |= buf[i] ^ (2 << 6);
                codepoint <<= 6;
            }
            // Put the least significant 6 bits of the last UTF-8 byte into the
            // codepoint.
            codepoint |= buf[utf8bytes-1] ^ (2 << 6);
        } else {
            // Put the first UTF-8 byte into the codepoint.
            codepoint |= buf[0];
        }

        // Write the resulting codepoint into the output file.
        size_t wet = fwrite(&codepoint, sizeof(codepoint), 1, out);
        if (wet != 1) {
            fprintf(stderr, "fwrite() failed: %zu\n", wet);
            fclose(in);
            fclose(out);
            exit(EXIT_FAILURE);
        }

        ret = fread(&buf_s, sizeof(buf_s), 1, in); // Read the next 1-byte.
        idx++;
    }

    fclose(in);
    fclose(out);

    exit(EXIT_SUCCESS);
}
