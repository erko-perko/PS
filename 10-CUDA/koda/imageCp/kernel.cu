#ifdef __cplusplus
extern "C" {
#endif

// V primeru robnih slikovnih točk, kjer nimamo na voljo vseh sosedov, obstaja več možnih pristopov.
// Najbolj običajen je, da za vrednosti slikovnih točk zunaj robov slike uporabimo kar najbližjo vrednost znotraj slike.
// Izberemo velikost okna (3x3), ki določa območje sosednjih slikovnih točk okoli ciljne slikovne točke.
// Preberemo vse vrednosti slikovnih točk znotraj okna in jih uredimo po velikosti.
// Mediano izračunamo tako, da vzamemo srednjo vrednost urejenega seznama. V primeru okna velikosti 3x3, imamo v seznamu 9 vrednosti. Srednja vrednost se nahaja na indeksu 4.
// Vrednost centralnega piksla nadomestimo z mediano.
__global__ void process(unsigned char *img_in, unsigned char *img_out, int width, int height) {
    // row
    int j = threadIdx.x + blockIdx.x * blockDim.x;
    // col
    int i = threadIdx.y + blockIdx.y * blockDim.y;
    int ipx = i * width + j;
    while (ipx < width * height) {
        if(i == 0) {
            if(j == 0) {
                img_out[ipx] = img_in[width + 1];
            } else if(j == width - 1) {
                img_out[ipx] = img_in[width + j - 1];
            } else {
                img_out[ipx] = img_in[width + j];
            }
        } else if(i == height - 1) {
            if(j == 0) {
                img_out[ipx] = img_in[(i - 1) * width + 1];
            } else if(j == width - 1) {
                img_out[ipx] = img_in[(i - 1) * width + j - 1];
            } else {
                img_out[ipx] = img_in[(i - 1) * width + j];
            }
        } else if(j == 0) {
            img_out[ipx] = img_in[i * width + 1];
        } else if(j == width - 1) {
            img_out[ipx] = img_in[i * width + j - 1];
        } else {
            // collect values
            unsigned char window[9];
            window[0] = img_in[(i - 1) * width + (j - 1)];
            window[1] = img_in[(i - 1) * width + j];
            window[2] = img_in[(i - 1) * width + (j + 1)];
            window[3] = img_in[i * width + (j - 1)];
            window[4] = img_in[i * width + j];
            window[5] = img_in[i * width + (j + 1)];
            window[6] = img_in[(i + 1) * width + (j - 1)];
            window[7] = img_in[(i + 1) * width + j];
            window[8] = img_in[(i + 1) * width + (j + 1)];
            // sort values (bubble sort)
            for(int a = 0; a < 9 - 1; a++) {
                for(int b = 0; b < 9 - a - 1; b++) {
                    if(window[b] > window[b + 1]) {
                        unsigned char temp = window[b];
                        window[b] = window[b + 1];
                        window[b + 1] = temp;
                    }
                }
            }
            // get median
            img_out[ipx] = window[4];
        }
        ipx += blockDim.x * gridDim.x * blockDim.y * gridDim.y;
        if (ipx < width * height) {
            i = ipx / width;
            j = ipx % width;
        }
    }
}
#ifdef __cplusplus
}
#endif