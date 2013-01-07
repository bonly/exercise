package main
import (
    "os"
    "io"
    "bytes"
    "encoding/base64"
    "compress/zlib"
)
var  tryme  =  `eJxslTGy4zAMQ/
ucIr00PoDUuWPhRoUkHoE97z8LULI3+f9zJj92dvwMghD3dV53nazrbGcbc4zza
qKzHlH5OJJ3sSo1v3Pvs+to/5+8K2D7E7DrnH3OcY4pqYpWcIJWrKubFNx00ISv
+wV7VMXXdQ0VNxNUKrmY8mnAck4+3Zy3Apr0v5TdsI0capZqMpmaSinV03G836T
VIt0Nt7kSpm09891m9BhusbonsqrpmFMFNm1WqZYEqvMBvR0tjz9gyzEC8aqzKX
RVrxWiUrLklgkDqpprHzABJhZc9vm7Tyij6RQXMChDa2AVOJbM8pusyt5RSTBnt
1Lc+/w9z9fZQhU/
Dbgh6KZm2gUtafVItGEerOqK7sWRG8z8T89IXI1OceVjJTlGB1SMMptHSHhJAxw
xgVCP5748W/m6Liq7WtfOLuE3WlyZBcM0hTDC+I9W4439SxtCC6co6WYi/
jClQsjDYqfw3BEWGEc3V8tIzg8YK7xnOqLZhmmmxcITQBEKt2hVJWSrxYBkftJe
ULXND3nhHQ9SJD9DAWFPu7hcN2Gdi/
TPuBG2Z0B1y9DJ0CJm9KWAtgEEbdhCMbvAPRF50X8GI7rdvw9iENkI15L2UQvFDKsV2FF9fsMoip1G01fD6O5MxMWHpm0Wc+alJKTvSI8yOh/
5D9RSBonO9j7qE7VkzVSgGoOVH8cp2iRrwbDRYMdS911MGXcKTuZUr8WKjud0s8
19kO6lhrWoOJ08NpGDr0LOiMLBFHOc
+zW1e6DcZytne6ADS8uwhjBIBAtQHsz4w8CGWYpjiU1lZrhobTxnPpbjksa/
QC0Wk47pY1diva2NgccdyWK2FL/wVrG8sWiE38DFQV
+yGmxQwfv4UjyK4wzv1lXocBSGOBoFoxz/EcA7bBDBz4S9/gUAAP//yf/Baw==`
func main(){    
    r0 := bytes.NewBufferString(tryme)
    r1 := base64.NewDecoder(base64.StdEncoding, r0)
    r2, _ := zlib.NewReader(r1)
    io.Copy(os.Stdout, r2)
    r2.Close()
}