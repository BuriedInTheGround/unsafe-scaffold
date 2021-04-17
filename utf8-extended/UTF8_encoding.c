#include <errno.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

int main() {
    FILE *in = fopen("input.data", "rb");
    FILE *out = fopen("UTF8.data", "wb");
    if (in == NULL) {
        perror("fopen() for input file failed");
        exit(EXIT_FAILURE);
    }
    if (out == NULL) {
        perror("fopen() for output file failed");
        fclose(in);
        exit(EXIT_FAILURE);
    }

    uint32_t buf; // Buffer for reading the input file 4-bytes at a time.

    size_t ret = fread(&buf, sizeof(buf), 1, in); // Read 4-bytes of data.
    while (!feof(in)) {
        // Check that the expected number of bytes was read.
        if (ret != 1) {
            fprintf(stderr, "fread() failed: %zu\n", ret);
            fclose(in);
            fclose(out);
            exit(EXIT_FAILURE);
        }

        // Evaluate the number of bytes needed to store the codepoint in UTF-8.
        size_t utf8bytes = 1;
        if (buf >= 1<<31) {
            utf8bytes = 7;
        } else if (buf >= 1<<26) {
            utf8bytes = 6;
        } else if (buf >= 1<<21) {
            utf8bytes = 5;
        } else if (buf >= 1<<16) {
            utf8bytes = 4;
        } else if (buf >= 1<<11) {
            utf8bytes = 3;
        } else if (buf >= 1<<7) {
            utf8bytes = 2;
        }

        uint8_t utf8e_v[utf8bytes]; // Extended UTF-8 codeword vector.

        // Discern how to encode based on the number of UTF-8 bytes.
        if (utf8bytes > 1) {
            for (size_t i = utf8bytes-1; i >= 1; i--) {
                // Set utf8e_v[i] to (10 || xxxxxx) where xxxxxx are the 6
                // least significant bits of the buffer.
                utf8e_v[i] = (2 << 6) | (buf & 63);
                // Right shift the buffer by 6 positions to drop the 6 least
                // significant bits.
                buf >>= 6;
            }
            // Put into utf8e_v[0] utf8bytes bits worth of 1's followed by a 0,
            // then followed by the remaining bits from the buffer.
            utf8e_v[0] = (254 << (7-utf8bytes)) | (buf & (254 >> utf8bytes));
        } else {
            utf8e_v[0] = buf; // Put the entire buffer into utf8e_v[0].
        }

        // Write the resulting UTF-8 codeword into the output file.
        size_t wet = fwrite(utf8e_v, sizeof(utf8e_v), sizeof(*utf8e_v), out);
        if (wet != sizeof(*utf8e_v)) {
            fprintf(stderr, "fwrite() failed: %zu\n", wet);
            fclose(in);
            fclose(out);
            exit(EXIT_FAILURE);
        }

        ret = fread(&buf, sizeof(buf), 1, in); // Read the next 4-bytes.
    }

    fclose(in);
    fclose(out);

    exit(EXIT_SUCCESS);
}
