package flowMap

const fnvBasis = 14695981039346656037
const fnvPrime = 1099511628211

// FnvHash is used by our FastHash functions, and implements the FNV hash
// created by Glenn Fowler, Landon Curt Noll, and Phong Vo.
// See http://isthe.com/chongo/tech/comp/fnv/.
func FnvHash(s []byte) (h uint64) {
	h = fnvBasis
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= fnvPrime
	}
	return
}

// It is guaranteed to collide with its reverse flow.
// IE: the flow A->B will have the same hash as the flow B->A.
func FastTwoHash(IP []byte, UE []byte) (h uint64) {
	h = FnvHash(IP) + FnvHash(UE)
	h ^= uint64(4) //type
	h *= fnvPrime
	return
}

func FastFourHash(srcIP string, dstIP string, protocol string, dstPort string) (h uint64) {
	h = FnvHash([]byte(dstPort)) + FnvHash([]byte(dstIP+srcIP)) + FnvHash([]byte(protocol))
	h ^= uint64(4) //type
	h *= fnvPrime
	return
}

func FastFiveHash(srcIP string, dstIP string, srcPort string, dstPort string, protocol string) (h uint64) {
	h = FnvHash([]byte(srcIP+srcPort)) + FnvHash([]byte(dstIP+dstPort)) + FnvHash([]byte(protocol))
	h ^= uint64(4) //type
	h *= fnvPrime
	return
}
