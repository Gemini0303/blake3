package compress

import (
	"fmt"
	"github.com/zeebo/blake3/internal/alg/compress/compress_pure"
)

func Compress(chain *[8]uint32, block *[16]uint32, counter uint64, blen uint32, flags uint32, out *[16]uint32) {
	fmt.Printf("use blake3 with mavericks \n")
	compress_pure.Compress(chain, block, counter, blen, flags, out)
}
