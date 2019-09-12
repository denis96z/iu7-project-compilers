// Code generated from /home/denis/Documents/iu7-project-compilers/dzc/DZ.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // DZ

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 70, 394,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 7, 4, 116, 10, 4, 12, 4, 14, 4, 119,
	11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 125, 10, 5, 3, 6, 3, 6, 5, 6, 129,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 7, 12, 153, 10, 12, 12, 12, 14, 12, 156, 11, 12, 5, 12, 158, 10, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 169,
	10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 178, 10,
	17, 13, 17, 14, 17, 179, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 20, 7, 20, 192, 10, 20, 12, 20, 14, 20, 195, 11, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 5, 23, 213, 10, 23, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 28, 5, 28, 230, 10, 28, 3, 29, 3, 29, 5, 29, 234, 10, 29, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 34, 3, 34, 5, 34, 251, 10, 34, 3, 35, 7, 35, 254, 10,
	35, 12, 35, 14, 35, 257, 11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 271, 10, 36, 3, 37,
	3, 37, 7, 37, 275, 10, 37, 12, 37, 14, 37, 278, 11, 37, 3, 37, 5, 37, 281,
	10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 5,
	41, 302, 10, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 7, 47, 329, 10, 47, 12, 47,
	14, 47, 332, 11, 47, 5, 47, 334, 10, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3,
	48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 7, 49, 359, 10,
	49, 12, 49, 14, 49, 362, 11, 49, 5, 49, 364, 10, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 5, 49, 370, 10, 49, 3, 49, 3, 49, 3, 49, 3, 49, 7, 49, 376, 10,
	49, 12, 49, 14, 49, 379, 11, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3,
	52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 5, 53, 392, 10, 53, 3, 53, 2, 3,
	96, 54, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
	2, 6, 3, 2, 65, 66, 3, 2, 20, 31, 4, 2, 42, 42, 46, 46, 4, 2, 41, 45, 47,
	62, 2, 381, 2, 106, 3, 2, 2, 2, 4, 110, 3, 2, 2, 2, 6, 117, 3, 2, 2, 2,
	8, 124, 3, 2, 2, 2, 10, 128, 3, 2, 2, 2, 12, 130, 3, 2, 2, 2, 14, 134,
	3, 2, 2, 2, 16, 138, 3, 2, 2, 2, 18, 141, 3, 2, 2, 2, 20, 145, 3, 2, 2,
	2, 22, 157, 3, 2, 2, 2, 24, 159, 3, 2, 2, 2, 26, 163, 3, 2, 2, 2, 28, 168,
	3, 2, 2, 2, 30, 170, 3, 2, 2, 2, 32, 177, 3, 2, 2, 2, 34, 181, 3, 2, 2,
	2, 36, 184, 3, 2, 2, 2, 38, 193, 3, 2, 2, 2, 40, 196, 3, 2, 2, 2, 42, 201,
	3, 2, 2, 2, 44, 212, 3, 2, 2, 2, 46, 214, 3, 2, 2, 2, 48, 216, 3, 2, 2,
	2, 50, 218, 3, 2, 2, 2, 52, 220, 3, 2, 2, 2, 54, 229, 3, 2, 2, 2, 56, 233,
	3, 2, 2, 2, 58, 235, 3, 2, 2, 2, 60, 237, 3, 2, 2, 2, 62, 239, 3, 2, 2,
	2, 64, 242, 3, 2, 2, 2, 66, 250, 3, 2, 2, 2, 68, 255, 3, 2, 2, 2, 70, 270,
	3, 2, 2, 2, 72, 272, 3, 2, 2, 2, 74, 282, 3, 2, 2, 2, 76, 288, 3, 2, 2,
	2, 78, 294, 3, 2, 2, 2, 80, 301, 3, 2, 2, 2, 82, 303, 3, 2, 2, 2, 84, 308,
	3, 2, 2, 2, 86, 314, 3, 2, 2, 2, 88, 316, 3, 2, 2, 2, 90, 318, 3, 2, 2,
	2, 92, 323, 3, 2, 2, 2, 94, 338, 3, 2, 2, 2, 96, 369, 3, 2, 2, 2, 98, 380,
	3, 2, 2, 2, 100, 382, 3, 2, 2, 2, 102, 384, 3, 2, 2, 2, 104, 391, 3, 2,
	2, 2, 106, 107, 5, 4, 3, 2, 107, 108, 5, 6, 4, 2, 108, 109, 7, 2, 2, 3,
	109, 3, 3, 2, 2, 2, 110, 111, 7, 3, 2, 2, 111, 112, 7, 69, 2, 2, 112, 113,
	7, 40, 2, 2, 113, 5, 3, 2, 2, 2, 114, 116, 5, 8, 5, 2, 115, 114, 3, 2,
	2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2,
	118, 7, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 125, 5, 42, 22, 2, 121,
	125, 5, 52, 27, 2, 122, 125, 5, 28, 15, 2, 123, 125, 5, 10, 6, 2, 124,
	120, 3, 2, 2, 2, 124, 121, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 123,
	3, 2, 2, 2, 125, 9, 3, 2, 2, 2, 126, 129, 5, 12, 7, 2, 127, 129, 5, 14,
	8, 2, 128, 126, 3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 11, 3, 2, 2, 2,
	130, 131, 7, 13, 2, 2, 131, 132, 5, 16, 9, 2, 132, 133, 5, 68, 35, 2, 133,
	13, 3, 2, 2, 2, 134, 135, 7, 14, 2, 2, 135, 136, 5, 18, 10, 2, 136, 137,
	5, 68, 35, 2, 137, 15, 3, 2, 2, 2, 138, 139, 7, 69, 2, 2, 139, 140, 5,
	20, 11, 2, 140, 17, 3, 2, 2, 2, 141, 142, 7, 69, 2, 2, 142, 143, 5, 20,
	11, 2, 143, 144, 5, 26, 14, 2, 144, 19, 3, 2, 2, 2, 145, 146, 7, 32, 2,
	2, 146, 147, 5, 22, 12, 2, 147, 148, 7, 33, 2, 2, 148, 21, 3, 2, 2, 2,
	149, 154, 5, 24, 13, 2, 150, 151, 7, 39, 2, 2, 151, 153, 5, 24, 13, 2,
	152, 150, 3, 2, 2, 2, 153, 156, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154,
	155, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 157, 149,
	3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 23, 3, 2, 2, 2, 159, 160, 7, 69,
	2, 2, 160, 161, 7, 38, 2, 2, 161, 162, 5, 56, 29, 2, 162, 25, 3, 2, 2,
	2, 163, 164, 7, 38, 2, 2, 164, 165, 5, 56, 29, 2, 165, 27, 3, 2, 2, 2,
	166, 169, 5, 30, 16, 2, 167, 169, 5, 36, 19, 2, 168, 166, 3, 2, 2, 2, 168,
	167, 3, 2, 2, 2, 169, 29, 3, 2, 2, 2, 170, 171, 7, 18, 2, 2, 171, 172,
	7, 68, 2, 2, 172, 173, 7, 34, 2, 2, 173, 174, 5, 32, 17, 2, 174, 175, 7,
	35, 2, 2, 175, 31, 3, 2, 2, 2, 176, 178, 5, 34, 18, 2, 177, 176, 3, 2,
	2, 2, 178, 179, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2,
	180, 33, 3, 2, 2, 2, 181, 182, 7, 67, 2, 2, 182, 183, 7, 39, 2, 2, 183,
	35, 3, 2, 2, 2, 184, 185, 7, 17, 2, 2, 185, 186, 7, 68, 2, 2, 186, 187,
	7, 34, 2, 2, 187, 188, 5, 38, 20, 2, 188, 189, 7, 35, 2, 2, 189, 37, 3,
	2, 2, 2, 190, 192, 5, 40, 21, 2, 191, 190, 3, 2, 2, 2, 192, 195, 3, 2,
	2, 2, 193, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 39, 3, 2, 2, 2,
	195, 193, 3, 2, 2, 2, 196, 197, 7, 69, 2, 2, 197, 198, 7, 38, 2, 2, 198,
	199, 5, 54, 28, 2, 199, 200, 7, 40, 2, 2, 200, 41, 3, 2, 2, 2, 201, 202,
	7, 15, 2, 2, 202, 203, 7, 67, 2, 2, 203, 204, 7, 38, 2, 2, 204, 205, 5,
	58, 30, 2, 205, 206, 7, 52, 2, 2, 206, 207, 5, 44, 23, 2, 207, 208, 7,
	40, 2, 2, 208, 43, 3, 2, 2, 2, 209, 213, 5, 46, 24, 2, 210, 213, 5, 48,
	25, 2, 211, 213, 5, 50, 26, 2, 212, 209, 3, 2, 2, 2, 212, 210, 3, 2, 2,
	2, 212, 211, 3, 2, 2, 2, 213, 45, 3, 2, 2, 2, 214, 215, 7, 64, 2, 2, 215,
	47, 3, 2, 2, 2, 216, 217, 9, 2, 2, 2, 217, 49, 3, 2, 2, 2, 218, 219, 7,
	67, 2, 2, 219, 51, 3, 2, 2, 2, 220, 221, 7, 4, 2, 2, 221, 222, 7, 68, 2,
	2, 222, 223, 7, 52, 2, 2, 223, 224, 5, 54, 28, 2, 224, 225, 7, 40, 2, 2,
	225, 53, 3, 2, 2, 2, 226, 230, 5, 56, 29, 2, 227, 230, 5, 62, 32, 2, 228,
	230, 5, 64, 33, 2, 229, 226, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229, 228,
	3, 2, 2, 2, 230, 55, 3, 2, 2, 2, 231, 234, 5, 58, 30, 2, 232, 234, 5, 60,
	31, 2, 233, 231, 3, 2, 2, 2, 233, 232, 3, 2, 2, 2, 234, 57, 3, 2, 2, 2,
	235, 236, 9, 3, 2, 2, 236, 59, 3, 2, 2, 2, 237, 238, 7, 68, 2, 2, 238,
	61, 3, 2, 2, 2, 239, 240, 7, 63, 2, 2, 240, 241, 5, 56, 29, 2, 241, 63,
	3, 2, 2, 2, 242, 243, 7, 36, 2, 2, 243, 244, 5, 56, 29, 2, 244, 245, 7,
	38, 2, 2, 245, 246, 5, 66, 34, 2, 246, 247, 7, 37, 2, 2, 247, 65, 3, 2,
	2, 2, 248, 251, 7, 64, 2, 2, 249, 251, 7, 67, 2, 2, 250, 248, 3, 2, 2,
	2, 250, 249, 3, 2, 2, 2, 251, 67, 3, 2, 2, 2, 252, 254, 5, 70, 36, 2, 253,
	252, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255, 256,
	3, 2, 2, 2, 256, 69, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 258, 271, 5, 72,
	37, 2, 259, 271, 5, 80, 41, 2, 260, 271, 5, 86, 44, 2, 261, 271, 5, 88,
	45, 2, 262, 271, 5, 92, 47, 2, 263, 264, 5, 90, 46, 2, 264, 265, 7, 40,
	2, 2, 265, 271, 3, 2, 2, 2, 266, 267, 5, 96, 49, 2, 267, 268, 7, 40, 2,
	2, 268, 271, 3, 2, 2, 2, 269, 271, 5, 104, 53, 2, 270, 258, 3, 2, 2, 2,
	270, 259, 3, 2, 2, 2, 270, 260, 3, 2, 2, 2, 270, 261, 3, 2, 2, 2, 270,
	262, 3, 2, 2, 2, 270, 263, 3, 2, 2, 2, 270, 266, 3, 2, 2, 2, 270, 269,
	3, 2, 2, 2, 271, 71, 3, 2, 2, 2, 272, 276, 5, 74, 38, 2, 273, 275, 5, 76,
	39, 2, 274, 273, 3, 2, 2, 2, 275, 278, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2,
	276, 277, 3, 2, 2, 2, 277, 280, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 279,
	281, 5, 78, 40, 2, 280, 279, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 73,
	3, 2, 2, 2, 282, 283, 7, 10, 2, 2, 283, 284, 5, 96, 49, 2, 284, 285, 7,
	34, 2, 2, 285, 286, 5, 68, 35, 2, 286, 287, 7, 35, 2, 2, 287, 75, 3, 2,
	2, 2, 288, 289, 7, 11, 2, 2, 289, 290, 5, 96, 49, 2, 290, 291, 7, 34, 2,
	2, 291, 292, 5, 68, 35, 2, 292, 293, 7, 35, 2, 2, 293, 77, 3, 2, 2, 2,
	294, 295, 7, 12, 2, 2, 295, 296, 7, 34, 2, 2, 296, 297, 5, 68, 35, 2, 297,
	298, 7, 35, 2, 2, 298, 79, 3, 2, 2, 2, 299, 302, 5, 82, 42, 2, 300, 302,
	5, 84, 43, 2, 301, 299, 3, 2, 2, 2, 301, 300, 3, 2, 2, 2, 302, 81, 3, 2,
	2, 2, 303, 304, 7, 7, 2, 2, 304, 305, 7, 34, 2, 2, 305, 306, 5, 68, 35,
	2, 306, 307, 7, 35, 2, 2, 307, 83, 3, 2, 2, 2, 308, 309, 7, 6, 2, 2, 309,
	310, 5, 96, 49, 2, 310, 311, 7, 34, 2, 2, 311, 312, 5, 68, 35, 2, 312,
	313, 7, 35, 2, 2, 313, 85, 3, 2, 2, 2, 314, 315, 7, 8, 2, 2, 315, 87, 3,
	2, 2, 2, 316, 317, 7, 9, 2, 2, 317, 89, 3, 2, 2, 2, 318, 319, 7, 16, 2,
	2, 319, 320, 7, 69, 2, 2, 320, 321, 7, 52, 2, 2, 321, 322, 5, 96, 49, 2,
	322, 91, 3, 2, 2, 2, 323, 324, 7, 69, 2, 2, 324, 333, 7, 32, 2, 2, 325,
	330, 5, 94, 48, 2, 326, 327, 7, 39, 2, 2, 327, 329, 5, 94, 48, 2, 328,
	326, 3, 2, 2, 2, 329, 332, 3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 330, 331,
	3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 333, 325, 3, 2,
	2, 2, 333, 334, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336, 7, 33, 2, 2,
	336, 337, 7, 40, 2, 2, 337, 93, 3, 2, 2, 2, 338, 339, 5, 96, 49, 2, 339,
	95, 3, 2, 2, 2, 340, 341, 8, 49, 1, 2, 341, 370, 7, 69, 2, 2, 342, 370,
	7, 64, 2, 2, 343, 370, 9, 2, 2, 2, 344, 370, 7, 67, 2, 2, 345, 346, 7,
	36, 2, 2, 346, 347, 5, 96, 49, 2, 347, 348, 7, 37, 2, 2, 348, 370, 3, 2,
	2, 2, 349, 350, 7, 32, 2, 2, 350, 351, 5, 96, 49, 2, 351, 352, 7, 33, 2,
	2, 352, 370, 3, 2, 2, 2, 353, 354, 7, 69, 2, 2, 354, 363, 7, 32, 2, 2,
	355, 360, 5, 98, 50, 2, 356, 357, 7, 39, 2, 2, 357, 359, 5, 98, 50, 2,
	358, 356, 3, 2, 2, 2, 359, 362, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360,
	361, 3, 2, 2, 2, 361, 364, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 363, 355,
	3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 370, 7, 33,
	2, 2, 366, 367, 5, 100, 51, 2, 367, 368, 5, 96, 49, 3, 368, 370, 3, 2,
	2, 2, 369, 340, 3, 2, 2, 2, 369, 342, 3, 2, 2, 2, 369, 343, 3, 2, 2, 2,
	369, 344, 3, 2, 2, 2, 369, 345, 3, 2, 2, 2, 369, 349, 3, 2, 2, 2, 369,
	353, 3, 2, 2, 2, 369, 366, 3, 2, 2, 2, 370, 377, 3, 2, 2, 2, 371, 372,
	12, 4, 2, 2, 372, 373, 5, 102, 52, 2, 373, 374, 5, 96, 49, 5, 374, 376,
	3, 2, 2, 2, 375, 371, 3, 2, 2, 2, 376, 379, 3, 2, 2, 2, 377, 375, 3, 2,
	2, 2, 377, 378, 3, 2, 2, 2, 378, 97, 3, 2, 2, 2, 379, 377, 3, 2, 2, 2,
	380, 381, 5, 96, 49, 2, 381, 99, 3, 2, 2, 2, 382, 383, 9, 4, 2, 2, 383,
	101, 3, 2, 2, 2, 384, 385, 9, 5, 2, 2, 385, 103, 3, 2, 2, 2, 386, 392,
	7, 19, 2, 2, 387, 388, 7, 19, 2, 2, 388, 389, 5, 96, 49, 2, 389, 390, 7,
	40, 2, 2, 390, 392, 3, 2, 2, 2, 391, 386, 3, 2, 2, 2, 391, 387, 3, 2, 2,
	2, 392, 105, 3, 2, 2, 2, 26, 117, 124, 128, 154, 157, 168, 179, 193, 212,
	229, 233, 250, 255, 270, 276, 280, 301, 330, 333, 360, 363, 369, 377, 391,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'pkg'", "'type'", "'for'", "'while'", "'loop'", "'break'", "'continue'",
	"'if'", "'elif'", "'else'", "'proc'", "'func'", "'const'", "'let'", "'struct'",
	"'enum'", "'return'", "'i8_t'", "'u8_t'", "'i16_t'", "'u16_t'", "'i32_t'",
	"'u32_t'", "'i64_t'", "'u64_t'", "'char_t'", "'bool_t'", "'size_t'", "'ssize_t'",
	"'('", "')'", "'{'", "'}'", "'['", "']'", "':'", "','", "';'", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'~'", "'<<'", "'>>'", "'&'", "'|'", "'^'", "'='",
	"'+='", "'-='", "'*='", "'/='", "'%='", "'<<='", "'>>='", "'&='", "'|='",
	"'^='", "'@'", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "KW_PKG", "KW_TYPE", "KW_FOR", "KW_WHILE", "KW_LOOP", "KW_BREAK", "KW_CONTINUE",
	"KW_IF", "KW_ELIF", "KW_ELSE", "KW_PROC", "KW_FUNC", "KW_CONST", "KW_LET",
	"KW_STRUCT", "KW_ENUM", "KW_RETURN", "I8_T", "U8_T", "I16_T", "U16_T",
	"I32_T", "U32_T", "I64_T", "U64_T", "CHAR_T", "BOOL_T", "SIZE_T", "SSIZE_T",
	"LEFT_PRT", "RIGHT_PRT", "LEFT_BRC", "RIGHT_BRC", "LEFT_BRK", "RIGHT_BRK",
	"COLON", "COMMA", "SEMICOLON", "ADD", "SUB", "MUL", "DIV", "MOD", "NOT",
	"SHL", "SHR", "AND", "OR", "XOR", "ASGN", "ADD_ASGN", "SUB_ASGN", "MUL_ASGN",
	"DIV_ASGN", "MOD_ASGN", "SHL_ASGN", "SHR_ASGN", "AND_ASGN", "OR_ASGN",
	"XOR_ASGN", "REF", "INT_CONST", "TRUE", "FALSE", "CONST", "TYPE", "IDENTIFIER",
	"WHITESPACE",
}

var ruleNames = []string{
	"start", "pkg", "decls", "decl", "subdecl", "procdecl", "funcdecl", "procheader",
	"funcheader", "args", "argdecls", "argdecl", "funcret", "complexdecl",
	"enumdecl", "enumoptions", "enumoption", "structdecl", "structattrs", "structattr",
	"constdecl", "constexpr", "intexpr", "boolexpr", "constval", "typedecl",
	"typespec", "simpletypespec", "basictypespec", "namedtypespec", "reftypespec",
	"arraytypespec", "sizespec", "block", "statement", "condition", "ifConditionBranch",
	"elifConditionBranch", "elseConditionBranch", "loop", "trueLoop", "whileLoop",
	"breakStatement", "continueStatement", "declaration", "procCall", "procParam",
	"expression", "funcParam", "unaryOperator", "binaryOperator", "returnStatement",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DZParser struct {
	*antlr.BaseParser
}

func NewDZParser(input antlr.TokenStream) *DZParser {
	this := new(DZParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DZ.g4"

	return this
}

// DZParser tokens.
const (
	DZParserEOF         = antlr.TokenEOF
	DZParserKW_PKG      = 1
	DZParserKW_TYPE     = 2
	DZParserKW_FOR      = 3
	DZParserKW_WHILE    = 4
	DZParserKW_LOOP     = 5
	DZParserKW_BREAK    = 6
	DZParserKW_CONTINUE = 7
	DZParserKW_IF       = 8
	DZParserKW_ELIF     = 9
	DZParserKW_ELSE     = 10
	DZParserKW_PROC     = 11
	DZParserKW_FUNC     = 12
	DZParserKW_CONST    = 13
	DZParserKW_LET      = 14
	DZParserKW_STRUCT   = 15
	DZParserKW_ENUM     = 16
	DZParserKW_RETURN   = 17
	DZParserI8_T        = 18
	DZParserU8_T        = 19
	DZParserI16_T       = 20
	DZParserU16_T       = 21
	DZParserI32_T       = 22
	DZParserU32_T       = 23
	DZParserI64_T       = 24
	DZParserU64_T       = 25
	DZParserCHAR_T      = 26
	DZParserBOOL_T      = 27
	DZParserSIZE_T      = 28
	DZParserSSIZE_T     = 29
	DZParserLEFT_PRT    = 30
	DZParserRIGHT_PRT   = 31
	DZParserLEFT_BRC    = 32
	DZParserRIGHT_BRC   = 33
	DZParserLEFT_BRK    = 34
	DZParserRIGHT_BRK   = 35
	DZParserCOLON       = 36
	DZParserCOMMA       = 37
	DZParserSEMICOLON   = 38
	DZParserADD         = 39
	DZParserSUB         = 40
	DZParserMUL         = 41
	DZParserDIV         = 42
	DZParserMOD         = 43
	DZParserNOT         = 44
	DZParserSHL         = 45
	DZParserSHR         = 46
	DZParserAND         = 47
	DZParserOR          = 48
	DZParserXOR         = 49
	DZParserASGN        = 50
	DZParserADD_ASGN    = 51
	DZParserSUB_ASGN    = 52
	DZParserMUL_ASGN    = 53
	DZParserDIV_ASGN    = 54
	DZParserMOD_ASGN    = 55
	DZParserSHL_ASGN    = 56
	DZParserSHR_ASGN    = 57
	DZParserAND_ASGN    = 58
	DZParserOR_ASGN     = 59
	DZParserXOR_ASGN    = 60
	DZParserREF         = 61
	DZParserINT_CONST   = 62
	DZParserTRUE        = 63
	DZParserFALSE       = 64
	DZParserCONST       = 65
	DZParserTYPE        = 66
	DZParserIDENTIFIER  = 67
	DZParserWHITESPACE  = 68
)

// DZParser rules.
const (
	DZParserRULE_start               = 0
	DZParserRULE_pkg                 = 1
	DZParserRULE_decls               = 2
	DZParserRULE_decl                = 3
	DZParserRULE_subdecl             = 4
	DZParserRULE_procdecl            = 5
	DZParserRULE_funcdecl            = 6
	DZParserRULE_procheader          = 7
	DZParserRULE_funcheader          = 8
	DZParserRULE_args                = 9
	DZParserRULE_argdecls            = 10
	DZParserRULE_argdecl             = 11
	DZParserRULE_funcret             = 12
	DZParserRULE_complexdecl         = 13
	DZParserRULE_enumdecl            = 14
	DZParserRULE_enumoptions         = 15
	DZParserRULE_enumoption          = 16
	DZParserRULE_structdecl          = 17
	DZParserRULE_structattrs         = 18
	DZParserRULE_structattr          = 19
	DZParserRULE_constdecl           = 20
	DZParserRULE_constexpr           = 21
	DZParserRULE_intexpr             = 22
	DZParserRULE_boolexpr            = 23
	DZParserRULE_constval            = 24
	DZParserRULE_typedecl            = 25
	DZParserRULE_typespec            = 26
	DZParserRULE_simpletypespec      = 27
	DZParserRULE_basictypespec       = 28
	DZParserRULE_namedtypespec       = 29
	DZParserRULE_reftypespec         = 30
	DZParserRULE_arraytypespec       = 31
	DZParserRULE_sizespec            = 32
	DZParserRULE_block               = 33
	DZParserRULE_statement           = 34
	DZParserRULE_condition           = 35
	DZParserRULE_ifConditionBranch   = 36
	DZParserRULE_elifConditionBranch = 37
	DZParserRULE_elseConditionBranch = 38
	DZParserRULE_loop                = 39
	DZParserRULE_trueLoop            = 40
	DZParserRULE_whileLoop           = 41
	DZParserRULE_breakStatement      = 42
	DZParserRULE_continueStatement   = 43
	DZParserRULE_declaration         = 44
	DZParserRULE_procCall            = 45
	DZParserRULE_procParam           = 46
	DZParserRULE_expression          = 47
	DZParserRULE_funcParam           = 48
	DZParserRULE_unaryOperator       = 49
	DZParserRULE_binaryOperator      = 50
	DZParserRULE_returnStatement     = 51
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Pkg() IPkgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPkgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPkgContext)
}

func (s *StartContext) Decls() IDeclsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclsContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(DZParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *DZParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DZParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Pkg()
	}
	{
		p.SetState(105)
		p.Decls()
	}
	{
		p.SetState(106)
		p.Match(DZParserEOF)
	}

	return localctx
}

// IPkgContext is an interface to support dynamic dispatch.
type IPkgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsPkgContext differentiates from other interfaces.
	IsPkgContext()
}

type PkgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyPkgContext() *PkgContext {
	var p = new(PkgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_pkg
	return p
}

func (*PkgContext) IsPkgContext() {}

func NewPkgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PkgContext {
	var p = new(PkgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_pkg

	return p
}

func (s *PkgContext) GetParser() antlr.Parser { return s.parser }

func (s *PkgContext) GetName() antlr.Token { return s.name }

func (s *PkgContext) SetName(v antlr.Token) { s.name = v }

func (s *PkgContext) KW_PKG() antlr.TerminalNode {
	return s.GetToken(DZParserKW_PKG, 0)
}

func (s *PkgContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *PkgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *PkgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PkgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PkgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterPkg(s)
	}
}

func (s *PkgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitPkg(s)
	}
}

func (p *DZParser) Pkg() (localctx IPkgContext) {
	localctx = NewPkgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DZParserRULE_pkg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(DZParserKW_PKG)
	}
	{
		p.SetState(109)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*PkgContext).name = _m
	}
	{
		p.SetState(110)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// IDeclsContext is an interface to support dynamic dispatch.
type IDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclsContext differentiates from other interfaces.
	IsDeclsContext()
}

type DeclsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclsContext() *DeclsContext {
	var p = new(DeclsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_decls
	return p
}

func (*DeclsContext) IsDeclsContext() {}

func NewDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclsContext {
	var p = new(DeclsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_decls

	return p
}

func (s *DeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclsContext) AllDecl() []IDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclContext)(nil)).Elem())
	var tst = make([]IDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclContext)
		}
	}

	return tst
}

func (s *DeclsContext) Decl(i int) IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *DeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterDecls(s)
	}
}

func (s *DeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitDecls(s)
	}
}

func (p *DZParser) Decls() (localctx IDeclsContext) {
	localctx = NewDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DZParserRULE_decls)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DZParserKW_TYPE)|(1<<DZParserKW_PROC)|(1<<DZParserKW_FUNC)|(1<<DZParserKW_CONST)|(1<<DZParserKW_STRUCT)|(1<<DZParserKW_ENUM))) != 0 {
		{
			p.SetState(112)
			p.Decl()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) Constdecl() IConstdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstdeclContext)
}

func (s *DeclContext) Typedecl() ITypedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedeclContext)
}

func (s *DeclContext) Complexdecl() IComplexdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComplexdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComplexdeclContext)
}

func (s *DeclContext) Subdecl() ISubdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubdeclContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *DZParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DZParserRULE_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(122)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_CONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.Constdecl()
		}

	case DZParserKW_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.Typedecl()
		}

	case DZParserKW_STRUCT, DZParserKW_ENUM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.Complexdecl()
		}

	case DZParserKW_PROC, DZParserKW_FUNC:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(121)
			p.Subdecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISubdeclContext is an interface to support dynamic dispatch.
type ISubdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubdeclContext differentiates from other interfaces.
	IsSubdeclContext()
}

type SubdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubdeclContext() *SubdeclContext {
	var p = new(SubdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_subdecl
	return p
}

func (*SubdeclContext) IsSubdeclContext() {}

func NewSubdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubdeclContext {
	var p = new(SubdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_subdecl

	return p
}

func (s *SubdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SubdeclContext) Procdecl() IProcdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcdeclContext)
}

func (s *SubdeclContext) Funcdecl() IFuncdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncdeclContext)
}

func (s *SubdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterSubdecl(s)
	}
}

func (s *SubdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitSubdecl(s)
	}
}

func (p *DZParser) Subdecl() (localctx ISubdeclContext) {
	localctx = NewSubdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DZParserRULE_subdecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_PROC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Procdecl()
		}

	case DZParserKW_FUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Funcdecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IProcdeclContext is an interface to support dynamic dispatch.
type IProcdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcdeclContext differentiates from other interfaces.
	IsProcdeclContext()
}

type ProcdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcdeclContext() *ProcdeclContext {
	var p = new(ProcdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_procdecl
	return p
}

func (*ProcdeclContext) IsProcdeclContext() {}

func NewProcdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcdeclContext {
	var p = new(ProcdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_procdecl

	return p
}

func (s *ProcdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcdeclContext) KW_PROC() antlr.TerminalNode {
	return s.GetToken(DZParserKW_PROC, 0)
}

func (s *ProcdeclContext) Procheader() IProcheaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcheaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcheaderContext)
}

func (s *ProcdeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProcdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterProcdecl(s)
	}
}

func (s *ProcdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitProcdecl(s)
	}
}

func (p *DZParser) Procdecl() (localctx IProcdeclContext) {
	localctx = NewProcdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DZParserRULE_procdecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(DZParserKW_PROC)
	}
	{
		p.SetState(129)
		p.Procheader()
	}
	{
		p.SetState(130)
		p.Block()
	}

	return localctx
}

// IFuncdeclContext is an interface to support dynamic dispatch.
type IFuncdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncdeclContext differentiates from other interfaces.
	IsFuncdeclContext()
}

type FuncdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncdeclContext() *FuncdeclContext {
	var p = new(FuncdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_funcdecl
	return p
}

func (*FuncdeclContext) IsFuncdeclContext() {}

func NewFuncdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncdeclContext {
	var p = new(FuncdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_funcdecl

	return p
}

func (s *FuncdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncdeclContext) KW_FUNC() antlr.TerminalNode {
	return s.GetToken(DZParserKW_FUNC, 0)
}

func (s *FuncdeclContext) Funcheader() IFuncheaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncheaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncheaderContext)
}

func (s *FuncdeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterFuncdecl(s)
	}
}

func (s *FuncdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitFuncdecl(s)
	}
}

func (p *DZParser) Funcdecl() (localctx IFuncdeclContext) {
	localctx = NewFuncdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DZParserRULE_funcdecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(DZParserKW_FUNC)
	}
	{
		p.SetState(133)
		p.Funcheader()
	}
	{
		p.SetState(134)
		p.Block()
	}

	return localctx
}

// IProcheaderContext is an interface to support dynamic dispatch.
type IProcheaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsProcheaderContext differentiates from other interfaces.
	IsProcheaderContext()
}

type ProcheaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyProcheaderContext() *ProcheaderContext {
	var p = new(ProcheaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_procheader
	return p
}

func (*ProcheaderContext) IsProcheaderContext() {}

func NewProcheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcheaderContext {
	var p = new(ProcheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_procheader

	return p
}

func (s *ProcheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcheaderContext) GetId() antlr.Token { return s.id }

func (s *ProcheaderContext) SetId(v antlr.Token) { s.id = v }

func (s *ProcheaderContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *ProcheaderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *ProcheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterProcheader(s)
	}
}

func (s *ProcheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitProcheader(s)
	}
}

func (p *DZParser) Procheader() (localctx IProcheaderContext) {
	localctx = NewProcheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DZParserRULE_procheader)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcheaderContext).id = _m
	}
	{
		p.SetState(137)
		p.Args()
	}

	return localctx
}

// IFuncheaderContext is an interface to support dynamic dispatch.
type IFuncheaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsFuncheaderContext differentiates from other interfaces.
	IsFuncheaderContext()
}

type FuncheaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyFuncheaderContext() *FuncheaderContext {
	var p = new(FuncheaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_funcheader
	return p
}

func (*FuncheaderContext) IsFuncheaderContext() {}

func NewFuncheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncheaderContext {
	var p = new(FuncheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_funcheader

	return p
}

func (s *FuncheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncheaderContext) GetId() antlr.Token { return s.id }

func (s *FuncheaderContext) SetId(v antlr.Token) { s.id = v }

func (s *FuncheaderContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *FuncheaderContext) Funcret() IFuncretContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncretContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncretContext)
}

func (s *FuncheaderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *FuncheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterFuncheader(s)
	}
}

func (s *FuncheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitFuncheader(s)
	}
}

func (p *DZParser) Funcheader() (localctx IFuncheaderContext) {
	localctx = NewFuncheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DZParserRULE_funcheader)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*FuncheaderContext).id = _m
	}
	{
		p.SetState(140)
		p.Args()
	}
	{
		p.SetState(141)
		p.Funcret()
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_PRT, 0)
}

func (s *ArgsContext) Argdecls() IArgdeclsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgdeclsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgdeclsContext)
}

func (s *ArgsContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_PRT, 0)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *DZParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DZParserRULE_args)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(DZParserLEFT_PRT)
	}
	{
		p.SetState(144)
		p.Argdecls()
	}
	{
		p.SetState(145)
		p.Match(DZParserRIGHT_PRT)
	}

	return localctx
}

// IArgdeclsContext is an interface to support dynamic dispatch.
type IArgdeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgdeclsContext differentiates from other interfaces.
	IsArgdeclsContext()
}

type ArgdeclsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgdeclsContext() *ArgdeclsContext {
	var p = new(ArgdeclsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_argdecls
	return p
}

func (*ArgdeclsContext) IsArgdeclsContext() {}

func NewArgdeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgdeclsContext {
	var p = new(ArgdeclsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_argdecls

	return p
}

func (s *ArgdeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgdeclsContext) AllArgdecl() []IArgdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgdeclContext)(nil)).Elem())
	var tst = make([]IArgdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgdeclContext)
		}
	}

	return tst
}

func (s *ArgdeclsContext) Argdecl(i int) IArgdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgdeclContext)
}

func (s *ArgdeclsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DZParserCOMMA)
}

func (s *ArgdeclsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DZParserCOMMA, i)
}

func (s *ArgdeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgdeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgdeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterArgdecls(s)
	}
}

func (s *ArgdeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitArgdecls(s)
	}
}

func (p *DZParser) Argdecls() (localctx IArgdeclsContext) {
	localctx = NewArgdeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DZParserRULE_argdecls)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserIDENTIFIER {
		{
			p.SetState(147)
			p.Argdecl()
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserCOMMA {
			{
				p.SetState(148)
				p.Match(DZParserCOMMA)
			}
			{
				p.SetState(149)
				p.Argdecl()
			}

			p.SetState(154)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IArgdeclContext is an interface to support dynamic dispatch.
type IArgdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() ISimpletypespecContext

	// SetT sets the t rule contexts.
	SetT(ISimpletypespecContext)

	// IsArgdeclContext differentiates from other interfaces.
	IsArgdeclContext()
}

type ArgdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	t      ISimpletypespecContext
}

func NewEmptyArgdeclContext() *ArgdeclContext {
	var p = new(ArgdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_argdecl
	return p
}

func (*ArgdeclContext) IsArgdeclContext() {}

func NewArgdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgdeclContext {
	var p = new(ArgdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_argdecl

	return p
}

func (s *ArgdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgdeclContext) GetId() antlr.Token { return s.id }

func (s *ArgdeclContext) SetId(v antlr.Token) { s.id = v }

func (s *ArgdeclContext) GetT() ISimpletypespecContext { return s.t }

func (s *ArgdeclContext) SetT(v ISimpletypespecContext) { s.t = v }

func (s *ArgdeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *ArgdeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *ArgdeclContext) Simpletypespec() ISimpletypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpletypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpletypespecContext)
}

func (s *ArgdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterArgdecl(s)
	}
}

func (s *ArgdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitArgdecl(s)
	}
}

func (p *DZParser) Argdecl() (localctx IArgdeclContext) {
	localctx = NewArgdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DZParserRULE_argdecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ArgdeclContext).id = _m
	}
	{
		p.SetState(158)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(159)

		var _x = p.Simpletypespec()

		localctx.(*ArgdeclContext).t = _x
	}

	return localctx
}

// IFuncretContext is an interface to support dynamic dispatch.
type IFuncretContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t rule contexts.
	GetT() ISimpletypespecContext

	// SetT sets the t rule contexts.
	SetT(ISimpletypespecContext)

	// IsFuncretContext differentiates from other interfaces.
	IsFuncretContext()
}

type FuncretContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t      ISimpletypespecContext
}

func NewEmptyFuncretContext() *FuncretContext {
	var p = new(FuncretContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_funcret
	return p
}

func (*FuncretContext) IsFuncretContext() {}

func NewFuncretContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncretContext {
	var p = new(FuncretContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_funcret

	return p
}

func (s *FuncretContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncretContext) GetT() ISimpletypespecContext { return s.t }

func (s *FuncretContext) SetT(v ISimpletypespecContext) { s.t = v }

func (s *FuncretContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *FuncretContext) Simpletypespec() ISimpletypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpletypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpletypespecContext)
}

func (s *FuncretContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncretContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncretContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterFuncret(s)
	}
}

func (s *FuncretContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitFuncret(s)
	}
}

func (p *DZParser) Funcret() (localctx IFuncretContext) {
	localctx = NewFuncretContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DZParserRULE_funcret)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(162)

		var _x = p.Simpletypespec()

		localctx.(*FuncretContext).t = _x
	}

	return localctx
}

// IComplexdeclContext is an interface to support dynamic dispatch.
type IComplexdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComplexdeclContext differentiates from other interfaces.
	IsComplexdeclContext()
}

type ComplexdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComplexdeclContext() *ComplexdeclContext {
	var p = new(ComplexdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_complexdecl
	return p
}

func (*ComplexdeclContext) IsComplexdeclContext() {}

func NewComplexdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComplexdeclContext {
	var p = new(ComplexdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_complexdecl

	return p
}

func (s *ComplexdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ComplexdeclContext) Enumdecl() IEnumdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumdeclContext)
}

func (s *ComplexdeclContext) Structdecl() IStructdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructdeclContext)
}

func (s *ComplexdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComplexdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComplexdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterComplexdecl(s)
	}
}

func (s *ComplexdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitComplexdecl(s)
	}
}

func (p *DZParser) Complexdecl() (localctx IComplexdeclContext) {
	localctx = NewComplexdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DZParserRULE_complexdecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_ENUM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Enumdecl()
		}

	case DZParserKW_STRUCT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Structdecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumdeclContext is an interface to support dynamic dispatch.
type IEnumdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsEnumdeclContext differentiates from other interfaces.
	IsEnumdeclContext()
}

type EnumdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyEnumdeclContext() *EnumdeclContext {
	var p = new(EnumdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_enumdecl
	return p
}

func (*EnumdeclContext) IsEnumdeclContext() {}

func NewEnumdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumdeclContext {
	var p = new(EnumdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_enumdecl

	return p
}

func (s *EnumdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumdeclContext) GetId() antlr.Token { return s.id }

func (s *EnumdeclContext) SetId(v antlr.Token) { s.id = v }

func (s *EnumdeclContext) KW_ENUM() antlr.TerminalNode {
	return s.GetToken(DZParserKW_ENUM, 0)
}

func (s *EnumdeclContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *EnumdeclContext) Enumoptions() IEnumoptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumoptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumoptionsContext)
}

func (s *EnumdeclContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *EnumdeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
}

func (s *EnumdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterEnumdecl(s)
	}
}

func (s *EnumdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitEnumdecl(s)
	}
}

func (p *DZParser) Enumdecl() (localctx IEnumdeclContext) {
	localctx = NewEnumdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DZParserRULE_enumdecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(DZParserKW_ENUM)
	}
	{
		p.SetState(169)

		var _m = p.Match(DZParserTYPE)

		localctx.(*EnumdeclContext).id = _m
	}
	{
		p.SetState(170)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(171)
		p.Enumoptions()
	}
	{
		p.SetState(172)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// IEnumoptionsContext is an interface to support dynamic dispatch.
type IEnumoptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumoptionsContext differentiates from other interfaces.
	IsEnumoptionsContext()
}

type EnumoptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumoptionsContext() *EnumoptionsContext {
	var p = new(EnumoptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_enumoptions
	return p
}

func (*EnumoptionsContext) IsEnumoptionsContext() {}

func NewEnumoptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumoptionsContext {
	var p = new(EnumoptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_enumoptions

	return p
}

func (s *EnumoptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumoptionsContext) AllEnumoption() []IEnumoptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumoptionContext)(nil)).Elem())
	var tst = make([]IEnumoptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumoptionContext)
		}
	}

	return tst
}

func (s *EnumoptionsContext) Enumoption(i int) IEnumoptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumoptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumoptionContext)
}

func (s *EnumoptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumoptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumoptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterEnumoptions(s)
	}
}

func (s *EnumoptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitEnumoptions(s)
	}
}

func (p *DZParser) Enumoptions() (localctx IEnumoptionsContext) {
	localctx = NewEnumoptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DZParserRULE_enumoptions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DZParserCONST {
		{
			p.SetState(174)
			p.Enumoption()
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEnumoptionContext is an interface to support dynamic dispatch.
type IEnumoptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsEnumoptionContext differentiates from other interfaces.
	IsEnumoptionContext()
}

type EnumoptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyEnumoptionContext() *EnumoptionContext {
	var p = new(EnumoptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_enumoption
	return p
}

func (*EnumoptionContext) IsEnumoptionContext() {}

func NewEnumoptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumoptionContext {
	var p = new(EnumoptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_enumoption

	return p
}

func (s *EnumoptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumoptionContext) GetId() antlr.Token { return s.id }

func (s *EnumoptionContext) SetId(v antlr.Token) { s.id = v }

func (s *EnumoptionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(DZParserCOMMA, 0)
}

func (s *EnumoptionContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
}

func (s *EnumoptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumoptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumoptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterEnumoption(s)
	}
}

func (s *EnumoptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitEnumoption(s)
	}
}

func (p *DZParser) Enumoption() (localctx IEnumoptionContext) {
	localctx = NewEnumoptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DZParserRULE_enumoption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)

		var _m = p.Match(DZParserCONST)

		localctx.(*EnumoptionContext).id = _m
	}
	{
		p.SetState(180)
		p.Match(DZParserCOMMA)
	}

	return localctx
}

// IStructdeclContext is an interface to support dynamic dispatch.
type IStructdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsStructdeclContext differentiates from other interfaces.
	IsStructdeclContext()
}

type StructdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyStructdeclContext() *StructdeclContext {
	var p = new(StructdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_structdecl
	return p
}

func (*StructdeclContext) IsStructdeclContext() {}

func NewStructdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructdeclContext {
	var p = new(StructdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_structdecl

	return p
}

func (s *StructdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructdeclContext) GetId() antlr.Token { return s.id }

func (s *StructdeclContext) SetId(v antlr.Token) { s.id = v }

func (s *StructdeclContext) KW_STRUCT() antlr.TerminalNode {
	return s.GetToken(DZParserKW_STRUCT, 0)
}

func (s *StructdeclContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *StructdeclContext) Structattrs() IStructattrsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructattrsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructattrsContext)
}

func (s *StructdeclContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *StructdeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
}

func (s *StructdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterStructdecl(s)
	}
}

func (s *StructdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitStructdecl(s)
	}
}

func (p *DZParser) Structdecl() (localctx IStructdeclContext) {
	localctx = NewStructdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DZParserRULE_structdecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(DZParserKW_STRUCT)
	}
	{
		p.SetState(183)

		var _m = p.Match(DZParserTYPE)

		localctx.(*StructdeclContext).id = _m
	}
	{
		p.SetState(184)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(185)
		p.Structattrs()
	}
	{
		p.SetState(186)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// IStructattrsContext is an interface to support dynamic dispatch.
type IStructattrsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructattrsContext differentiates from other interfaces.
	IsStructattrsContext()
}

type StructattrsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructattrsContext() *StructattrsContext {
	var p = new(StructattrsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_structattrs
	return p
}

func (*StructattrsContext) IsStructattrsContext() {}

func NewStructattrsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructattrsContext {
	var p = new(StructattrsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_structattrs

	return p
}

func (s *StructattrsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructattrsContext) AllStructattr() []IStructattrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructattrContext)(nil)).Elem())
	var tst = make([]IStructattrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructattrContext)
		}
	}

	return tst
}

func (s *StructattrsContext) Structattr(i int) IStructattrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructattrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructattrContext)
}

func (s *StructattrsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructattrsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructattrsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterStructattrs(s)
	}
}

func (s *StructattrsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitStructattrs(s)
	}
}

func (p *DZParser) Structattrs() (localctx IStructattrsContext) {
	localctx = NewStructattrsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DZParserRULE_structattrs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserIDENTIFIER {
		{
			p.SetState(188)
			p.Structattr()
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStructattrContext is an interface to support dynamic dispatch.
type IStructattrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() ITypespecContext

	// SetT sets the t rule contexts.
	SetT(ITypespecContext)

	// IsStructattrContext differentiates from other interfaces.
	IsStructattrContext()
}

type StructattrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	t      ITypespecContext
}

func NewEmptyStructattrContext() *StructattrContext {
	var p = new(StructattrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_structattr
	return p
}

func (*StructattrContext) IsStructattrContext() {}

func NewStructattrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructattrContext {
	var p = new(StructattrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_structattr

	return p
}

func (s *StructattrContext) GetParser() antlr.Parser { return s.parser }

func (s *StructattrContext) GetId() antlr.Token { return s.id }

func (s *StructattrContext) SetId(v antlr.Token) { s.id = v }

func (s *StructattrContext) GetT() ITypespecContext { return s.t }

func (s *StructattrContext) SetT(v ITypespecContext) { s.t = v }

func (s *StructattrContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *StructattrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *StructattrContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *StructattrContext) Typespec() ITypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypespecContext)
}

func (s *StructattrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructattrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructattrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterStructattr(s)
	}
}

func (s *StructattrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitStructattr(s)
	}
}

func (p *DZParser) Structattr() (localctx IStructattrContext) {
	localctx = NewStructattrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DZParserRULE_structattr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*StructattrContext).id = _m
	}
	{
		p.SetState(195)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(196)

		var _x = p.Typespec()

		localctx.(*StructattrContext).t = _x
	}
	{
		p.SetState(197)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// IConstdeclContext is an interface to support dynamic dispatch.
type IConstdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() IBasictypespecContext

	// GetV returns the v rule contexts.
	GetV() IConstexprContext

	// SetT sets the t rule contexts.
	SetT(IBasictypespecContext)

	// SetV sets the v rule contexts.
	SetV(IConstexprContext)

	// IsConstdeclContext differentiates from other interfaces.
	IsConstdeclContext()
}

type ConstdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	t      IBasictypespecContext
	v      IConstexprContext
}

func NewEmptyConstdeclContext() *ConstdeclContext {
	var p = new(ConstdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_constdecl
	return p
}

func (*ConstdeclContext) IsConstdeclContext() {}

func NewConstdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstdeclContext {
	var p = new(ConstdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_constdecl

	return p
}

func (s *ConstdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstdeclContext) GetId() antlr.Token { return s.id }

func (s *ConstdeclContext) SetId(v antlr.Token) { s.id = v }

func (s *ConstdeclContext) GetT() IBasictypespecContext { return s.t }

func (s *ConstdeclContext) GetV() IConstexprContext { return s.v }

func (s *ConstdeclContext) SetT(v IBasictypespecContext) { s.t = v }

func (s *ConstdeclContext) SetV(v IConstexprContext) { s.v = v }

func (s *ConstdeclContext) KW_CONST() antlr.TerminalNode {
	return s.GetToken(DZParserKW_CONST, 0)
}

func (s *ConstdeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *ConstdeclContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *ConstdeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *ConstdeclContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
}

func (s *ConstdeclContext) Basictypespec() IBasictypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasictypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasictypespecContext)
}

func (s *ConstdeclContext) Constexpr() IConstexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstexprContext)
}

func (s *ConstdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterConstdecl(s)
	}
}

func (s *ConstdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitConstdecl(s)
	}
}

func (p *DZParser) Constdecl() (localctx IConstdeclContext) {
	localctx = NewConstdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DZParserRULE_constdecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(DZParserKW_CONST)
	}
	{
		p.SetState(200)

		var _m = p.Match(DZParserCONST)

		localctx.(*ConstdeclContext).id = _m
	}
	{
		p.SetState(201)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(202)

		var _x = p.Basictypespec()

		localctx.(*ConstdeclContext).t = _x
	}
	{
		p.SetState(203)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(204)

		var _x = p.Constexpr()

		localctx.(*ConstdeclContext).v = _x
	}
	{
		p.SetState(205)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// IConstexprContext is an interface to support dynamic dispatch.
type IConstexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstexprContext differentiates from other interfaces.
	IsConstexprContext()
}

type ConstexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstexprContext() *ConstexprContext {
	var p = new(ConstexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_constexpr
	return p
}

func (*ConstexprContext) IsConstexprContext() {}

func NewConstexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstexprContext {
	var p = new(ConstexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_constexpr

	return p
}

func (s *ConstexprContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstexprContext) Intexpr() IIntexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntexprContext)
}

func (s *ConstexprContext) Boolexpr() IBoolexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolexprContext)
}

func (s *ConstexprContext) Constval() IConstvalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstvalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstvalContext)
}

func (s *ConstexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterConstexpr(s)
	}
}

func (s *ConstexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitConstexpr(s)
	}
}

func (p *DZParser) Constexpr() (localctx IConstexprContext) {
	localctx = NewConstexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DZParserRULE_constexpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(210)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserINT_CONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Intexpr()
		}

	case DZParserTRUE, DZParserFALSE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Boolexpr()
		}

	case DZParserCONST:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(209)
			p.Constval()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntexprContext is an interface to support dynamic dispatch.
type IIntexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsIntexprContext differentiates from other interfaces.
	IsIntexprContext()
}

type IntexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyIntexprContext() *IntexprContext {
	var p = new(IntexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_intexpr
	return p
}

func (*IntexprContext) IsIntexprContext() {}

func NewIntexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntexprContext {
	var p = new(IntexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_intexpr

	return p
}

func (s *IntexprContext) GetParser() antlr.Parser { return s.parser }

func (s *IntexprContext) GetV() antlr.Token { return s.v }

func (s *IntexprContext) SetV(v antlr.Token) { s.v = v }

func (s *IntexprContext) INT_CONST() antlr.TerminalNode {
	return s.GetToken(DZParserINT_CONST, 0)
}

func (s *IntexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterIntexpr(s)
	}
}

func (s *IntexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitIntexpr(s)
	}
}

func (p *DZParser) Intexpr() (localctx IIntexprContext) {
	localctx = NewIntexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DZParserRULE_intexpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)

		var _m = p.Match(DZParserINT_CONST)

		localctx.(*IntexprContext).v = _m
	}

	return localctx
}

// IBoolexprContext is an interface to support dynamic dispatch.
type IBoolexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsBoolexprContext differentiates from other interfaces.
	IsBoolexprContext()
}

type BoolexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyBoolexprContext() *BoolexprContext {
	var p = new(BoolexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_boolexpr
	return p
}

func (*BoolexprContext) IsBoolexprContext() {}

func NewBoolexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolexprContext {
	var p = new(BoolexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_boolexpr

	return p
}

func (s *BoolexprContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolexprContext) GetV() antlr.Token { return s.v }

func (s *BoolexprContext) SetV(v antlr.Token) { s.v = v }

func (s *BoolexprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(DZParserTRUE, 0)
}

func (s *BoolexprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(DZParserFALSE, 0)
}

func (s *BoolexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterBoolexpr(s)
	}
}

func (s *BoolexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitBoolexpr(s)
	}
}

func (p *DZParser) Boolexpr() (localctx IBoolexprContext) {
	localctx = NewBoolexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DZParserRULE_boolexpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*BoolexprContext).v = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == DZParserTRUE || _la == DZParserFALSE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*BoolexprContext).v = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConstvalContext is an interface to support dynamic dispatch.
type IConstvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsConstvalContext differentiates from other interfaces.
	IsConstvalContext()
}

type ConstvalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyConstvalContext() *ConstvalContext {
	var p = new(ConstvalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_constval
	return p
}

func (*ConstvalContext) IsConstvalContext() {}

func NewConstvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstvalContext {
	var p = new(ConstvalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_constval

	return p
}

func (s *ConstvalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstvalContext) GetId() antlr.Token { return s.id }

func (s *ConstvalContext) SetId(v antlr.Token) { s.id = v }

func (s *ConstvalContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
}

func (s *ConstvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterConstval(s)
	}
}

func (s *ConstvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitConstval(s)
	}
}

func (p *DZParser) Constval() (localctx IConstvalContext) {
	localctx = NewConstvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DZParserRULE_constval)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)

		var _m = p.Match(DZParserCONST)

		localctx.(*ConstvalContext).id = _m
	}

	return localctx
}

// ITypedeclContext is an interface to support dynamic dispatch.
type ITypedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() ITypespecContext

	// SetT sets the t rule contexts.
	SetT(ITypespecContext)

	// IsTypedeclContext differentiates from other interfaces.
	IsTypedeclContext()
}

type TypedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	t      ITypespecContext
}

func NewEmptyTypedeclContext() *TypedeclContext {
	var p = new(TypedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_typedecl
	return p
}

func (*TypedeclContext) IsTypedeclContext() {}

func NewTypedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedeclContext {
	var p = new(TypedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_typedecl

	return p
}

func (s *TypedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedeclContext) GetId() antlr.Token { return s.id }

func (s *TypedeclContext) SetId(v antlr.Token) { s.id = v }

func (s *TypedeclContext) GetT() ITypespecContext { return s.t }

func (s *TypedeclContext) SetT(v ITypespecContext) { s.t = v }

func (s *TypedeclContext) KW_TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_TYPE, 0)
}

func (s *TypedeclContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *TypedeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *TypedeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
}

func (s *TypedeclContext) Typespec() ITypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypespecContext)
}

func (s *TypedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterTypedecl(s)
	}
}

func (s *TypedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitTypedecl(s)
	}
}

func (p *DZParser) Typedecl() (localctx ITypedeclContext) {
	localctx = NewTypedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DZParserRULE_typedecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(DZParserKW_TYPE)
	}
	{
		p.SetState(219)

		var _m = p.Match(DZParserTYPE)

		localctx.(*TypedeclContext).id = _m
	}
	{
		p.SetState(220)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(221)

		var _x = p.Typespec()

		localctx.(*TypedeclContext).t = _x
	}
	{
		p.SetState(222)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// ITypespecContext is an interface to support dynamic dispatch.
type ITypespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypespecContext differentiates from other interfaces.
	IsTypespecContext()
}

type TypespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypespecContext() *TypespecContext {
	var p = new(TypespecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_typespec
	return p
}

func (*TypespecContext) IsTypespecContext() {}

func NewTypespecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypespecContext {
	var p = new(TypespecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_typespec

	return p
}

func (s *TypespecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypespecContext) Simpletypespec() ISimpletypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpletypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpletypespecContext)
}

func (s *TypespecContext) Reftypespec() IReftypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReftypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReftypespecContext)
}

func (s *TypespecContext) Arraytypespec() IArraytypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArraytypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArraytypespecContext)
}

func (s *TypespecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypespecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypespecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterTypespec(s)
	}
}

func (s *TypespecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitTypespec(s)
	}
}

func (p *DZParser) Typespec() (localctx ITypespecContext) {
	localctx = NewTypespecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DZParserRULE_typespec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserI8_T, DZParserU8_T, DZParserI16_T, DZParserU16_T, DZParserI32_T, DZParserU32_T, DZParserI64_T, DZParserU64_T, DZParserCHAR_T, DZParserBOOL_T, DZParserSIZE_T, DZParserSSIZE_T, DZParserTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(224)
			p.Simpletypespec()
		}

	case DZParserREF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(225)
			p.Reftypespec()
		}

	case DZParserLEFT_BRK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(226)
			p.Arraytypespec()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimpletypespecContext is an interface to support dynamic dispatch.
type ISimpletypespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpletypespecContext differentiates from other interfaces.
	IsSimpletypespecContext()
}

type SimpletypespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpletypespecContext() *SimpletypespecContext {
	var p = new(SimpletypespecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_simpletypespec
	return p
}

func (*SimpletypespecContext) IsSimpletypespecContext() {}

func NewSimpletypespecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpletypespecContext {
	var p = new(SimpletypespecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_simpletypespec

	return p
}

func (s *SimpletypespecContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpletypespecContext) Basictypespec() IBasictypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasictypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasictypespecContext)
}

func (s *SimpletypespecContext) Namedtypespec() INamedtypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedtypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedtypespecContext)
}

func (s *SimpletypespecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpletypespecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpletypespecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterSimpletypespec(s)
	}
}

func (s *SimpletypespecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitSimpletypespec(s)
	}
}

func (p *DZParser) Simpletypespec() (localctx ISimpletypespecContext) {
	localctx = NewSimpletypespecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DZParserRULE_simpletypespec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(231)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserI8_T, DZParserU8_T, DZParserI16_T, DZParserU16_T, DZParserI32_T, DZParserU32_T, DZParserI64_T, DZParserU64_T, DZParserCHAR_T, DZParserBOOL_T, DZParserSIZE_T, DZParserSSIZE_T:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Basictypespec()
		}

	case DZParserTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Namedtypespec()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBasictypespecContext is an interface to support dynamic dispatch.
type IBasictypespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsBasictypespecContext differentiates from other interfaces.
	IsBasictypespecContext()
}

type BasictypespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyBasictypespecContext() *BasictypespecContext {
	var p = new(BasictypespecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_basictypespec
	return p
}

func (*BasictypespecContext) IsBasictypespecContext() {}

func NewBasictypespecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasictypespecContext {
	var p = new(BasictypespecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_basictypespec

	return p
}

func (s *BasictypespecContext) GetParser() antlr.Parser { return s.parser }

func (s *BasictypespecContext) GetId() antlr.Token { return s.id }

func (s *BasictypespecContext) SetId(v antlr.Token) { s.id = v }

func (s *BasictypespecContext) I8_T() antlr.TerminalNode {
	return s.GetToken(DZParserI8_T, 0)
}

func (s *BasictypespecContext) U8_T() antlr.TerminalNode {
	return s.GetToken(DZParserU8_T, 0)
}

func (s *BasictypespecContext) I16_T() antlr.TerminalNode {
	return s.GetToken(DZParserI16_T, 0)
}

func (s *BasictypespecContext) U16_T() antlr.TerminalNode {
	return s.GetToken(DZParserU16_T, 0)
}

func (s *BasictypespecContext) I32_T() antlr.TerminalNode {
	return s.GetToken(DZParserI32_T, 0)
}

func (s *BasictypespecContext) U32_T() antlr.TerminalNode {
	return s.GetToken(DZParserU32_T, 0)
}

func (s *BasictypespecContext) I64_T() antlr.TerminalNode {
	return s.GetToken(DZParserI64_T, 0)
}

func (s *BasictypespecContext) U64_T() antlr.TerminalNode {
	return s.GetToken(DZParserU64_T, 0)
}

func (s *BasictypespecContext) CHAR_T() antlr.TerminalNode {
	return s.GetToken(DZParserCHAR_T, 0)
}

func (s *BasictypespecContext) BOOL_T() antlr.TerminalNode {
	return s.GetToken(DZParserBOOL_T, 0)
}

func (s *BasictypespecContext) SIZE_T() antlr.TerminalNode {
	return s.GetToken(DZParserSIZE_T, 0)
}

func (s *BasictypespecContext) SSIZE_T() antlr.TerminalNode {
	return s.GetToken(DZParserSSIZE_T, 0)
}

func (s *BasictypespecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasictypespecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasictypespecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterBasictypespec(s)
	}
}

func (s *BasictypespecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitBasictypespec(s)
	}
}

func (p *DZParser) Basictypespec() (localctx IBasictypespecContext) {
	localctx = NewBasictypespecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DZParserRULE_basictypespec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*BasictypespecContext).id = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DZParserI8_T)|(1<<DZParserU8_T)|(1<<DZParserI16_T)|(1<<DZParserU16_T)|(1<<DZParserI32_T)|(1<<DZParserU32_T)|(1<<DZParserI64_T)|(1<<DZParserU64_T)|(1<<DZParserCHAR_T)|(1<<DZParserBOOL_T)|(1<<DZParserSIZE_T)|(1<<DZParserSSIZE_T))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*BasictypespecContext).id = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INamedtypespecContext is an interface to support dynamic dispatch.
type INamedtypespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsNamedtypespecContext differentiates from other interfaces.
	IsNamedtypespecContext()
}

type NamedtypespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyNamedtypespecContext() *NamedtypespecContext {
	var p = new(NamedtypespecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_namedtypespec
	return p
}

func (*NamedtypespecContext) IsNamedtypespecContext() {}

func NewNamedtypespecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedtypespecContext {
	var p = new(NamedtypespecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_namedtypespec

	return p
}

func (s *NamedtypespecContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedtypespecContext) GetId() antlr.Token { return s.id }

func (s *NamedtypespecContext) SetId(v antlr.Token) { s.id = v }

func (s *NamedtypespecContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
}

func (s *NamedtypespecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedtypespecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedtypespecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterNamedtypespec(s)
	}
}

func (s *NamedtypespecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitNamedtypespec(s)
	}
}

func (p *DZParser) Namedtypespec() (localctx INamedtypespecContext) {
	localctx = NewNamedtypespecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DZParserRULE_namedtypespec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)

		var _m = p.Match(DZParserTYPE)

		localctx.(*NamedtypespecContext).id = _m
	}

	return localctx
}

// IReftypespecContext is an interface to support dynamic dispatch.
type IReftypespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id rule contexts.
	GetId() ISimpletypespecContext

	// SetId sets the id rule contexts.
	SetId(ISimpletypespecContext)

	// IsReftypespecContext differentiates from other interfaces.
	IsReftypespecContext()
}

type ReftypespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     ISimpletypespecContext
}

func NewEmptyReftypespecContext() *ReftypespecContext {
	var p = new(ReftypespecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_reftypespec
	return p
}

func (*ReftypespecContext) IsReftypespecContext() {}

func NewReftypespecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReftypespecContext {
	var p = new(ReftypespecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_reftypespec

	return p
}

func (s *ReftypespecContext) GetParser() antlr.Parser { return s.parser }

func (s *ReftypespecContext) GetId() ISimpletypespecContext { return s.id }

func (s *ReftypespecContext) SetId(v ISimpletypespecContext) { s.id = v }

func (s *ReftypespecContext) REF() antlr.TerminalNode {
	return s.GetToken(DZParserREF, 0)
}

func (s *ReftypespecContext) Simpletypespec() ISimpletypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpletypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpletypespecContext)
}

func (s *ReftypespecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReftypespecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReftypespecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterReftypespec(s)
	}
}

func (s *ReftypespecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitReftypespec(s)
	}
}

func (p *DZParser) Reftypespec() (localctx IReftypespecContext) {
	localctx = NewReftypespecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DZParserRULE_reftypespec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(DZParserREF)
	}
	{
		p.SetState(238)

		var _x = p.Simpletypespec()

		localctx.(*ReftypespecContext).id = _x
	}

	return localctx
}

// IArraytypespecContext is an interface to support dynamic dispatch.
type IArraytypespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id rule contexts.
	GetId() ISimpletypespecContext

	// GetSize returns the size rule contexts.
	GetSize() ISizespecContext

	// SetId sets the id rule contexts.
	SetId(ISimpletypespecContext)

	// SetSize sets the size rule contexts.
	SetSize(ISizespecContext)

	// IsArraytypespecContext differentiates from other interfaces.
	IsArraytypespecContext()
}

type ArraytypespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     ISimpletypespecContext
	size   ISizespecContext
}

func NewEmptyArraytypespecContext() *ArraytypespecContext {
	var p = new(ArraytypespecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_arraytypespec
	return p
}

func (*ArraytypespecContext) IsArraytypespecContext() {}

func NewArraytypespecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArraytypespecContext {
	var p = new(ArraytypespecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_arraytypespec

	return p
}

func (s *ArraytypespecContext) GetParser() antlr.Parser { return s.parser }

func (s *ArraytypespecContext) GetId() ISimpletypespecContext { return s.id }

func (s *ArraytypespecContext) GetSize() ISizespecContext { return s.size }

func (s *ArraytypespecContext) SetId(v ISimpletypespecContext) { s.id = v }

func (s *ArraytypespecContext) SetSize(v ISizespecContext) { s.size = v }

func (s *ArraytypespecContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRK, 0)
}

func (s *ArraytypespecContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *ArraytypespecContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRK, 0)
}

func (s *ArraytypespecContext) Simpletypespec() ISimpletypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpletypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpletypespecContext)
}

func (s *ArraytypespecContext) Sizespec() ISizespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISizespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISizespecContext)
}

func (s *ArraytypespecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraytypespecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArraytypespecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterArraytypespec(s)
	}
}

func (s *ArraytypespecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitArraytypespec(s)
	}
}

func (p *DZParser) Arraytypespec() (localctx IArraytypespecContext) {
	localctx = NewArraytypespecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DZParserRULE_arraytypespec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(DZParserLEFT_BRK)
	}
	{
		p.SetState(241)

		var _x = p.Simpletypespec()

		localctx.(*ArraytypespecContext).id = _x
	}
	{
		p.SetState(242)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(243)

		var _x = p.Sizespec()

		localctx.(*ArraytypespecContext).size = _x
	}
	{
		p.SetState(244)
		p.Match(DZParserRIGHT_BRK)
	}

	return localctx
}

// ISizespecContext is an interface to support dynamic dispatch.
type ISizespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsSizespecContext differentiates from other interfaces.
	IsSizespecContext()
}

type SizespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptySizespecContext() *SizespecContext {
	var p = new(SizespecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_sizespec
	return p
}

func (*SizespecContext) IsSizespecContext() {}

func NewSizespecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizespecContext {
	var p = new(SizespecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_sizespec

	return p
}

func (s *SizespecContext) GetParser() antlr.Parser { return s.parser }

func (s *SizespecContext) GetName() antlr.Token { return s.name }

func (s *SizespecContext) SetName(v antlr.Token) { s.name = v }

func (s *SizespecContext) INT_CONST() antlr.TerminalNode {
	return s.GetToken(DZParserINT_CONST, 0)
}

func (s *SizespecContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
}

func (s *SizespecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizespecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SizespecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterSizespec(s)
	}
}

func (s *SizespecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitSizespec(s)
	}
}

func (p *DZParser) Sizespec() (localctx ISizespecContext) {
	localctx = NewSizespecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, DZParserRULE_sizespec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserINT_CONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			p.Match(DZParserINT_CONST)
		}

	case DZParserCONST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)

			var _m = p.Match(DZParserCONST)

			localctx.(*SizespecContext).name = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *DZParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, DZParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(DZParserKW_WHILE-4))|(1<<(DZParserKW_LOOP-4))|(1<<(DZParserKW_BREAK-4))|(1<<(DZParserKW_CONTINUE-4))|(1<<(DZParserKW_IF-4))|(1<<(DZParserKW_LET-4))|(1<<(DZParserKW_RETURN-4))|(1<<(DZParserLEFT_PRT-4))|(1<<(DZParserLEFT_BRK-4)))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(DZParserSUB-40))|(1<<(DZParserNOT-40))|(1<<(DZParserINT_CONST-40))|(1<<(DZParserTRUE-40))|(1<<(DZParserFALSE-40))|(1<<(DZParserCONST-40))|(1<<(DZParserIDENTIFIER-40)))) != 0) {
		{
			p.SetState(250)
			p.Statement()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinueStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) ProcCall() IProcCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcCallContext)
}

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *DZParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, DZParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Condition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.Loop()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(258)
			p.BreakStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(259)
			p.ContinueStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(260)
			p.ProcCall()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(261)
			p.Declaration()
		}
		{
			p.SetState(262)
			p.Match(DZParserSEMICOLON)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(264)
			p.expression(0)
		}
		{
			p.SetState(265)
			p.Match(DZParserSEMICOLON)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(267)
			p.ReturnStatement()
		}

	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) IfConditionBranch() IIfConditionBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfConditionBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfConditionBranchContext)
}

func (s *ConditionContext) AllElifConditionBranch() []IElifConditionBranchContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElifConditionBranchContext)(nil)).Elem())
	var tst = make([]IElifConditionBranchContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElifConditionBranchContext)
		}
	}

	return tst
}

func (s *ConditionContext) ElifConditionBranch(i int) IElifConditionBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElifConditionBranchContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElifConditionBranchContext)
}

func (s *ConditionContext) ElseConditionBranch() IElseConditionBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseConditionBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseConditionBranchContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *DZParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, DZParserRULE_condition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.IfConditionBranch()
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserKW_ELIF {
		{
			p.SetState(271)
			p.ElifConditionBranch()
		}

		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserKW_ELSE {
		{
			p.SetState(277)
			p.ElseConditionBranch()
		}

	}

	return localctx
}

// IIfConditionBranchContext is an interface to support dynamic dispatch.
type IIfConditionBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// GetBody returns the body rule contexts.
	GetBody() IBlockContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// SetBody sets the body rule contexts.
	SetBody(IBlockContext)

	// IsIfConditionBranchContext differentiates from other interfaces.
	IsIfConditionBranchContext()
}

type IfConditionBranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   IExpressionContext
	body   IBlockContext
}

func NewEmptyIfConditionBranchContext() *IfConditionBranchContext {
	var p = new(IfConditionBranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_ifConditionBranch
	return p
}

func (*IfConditionBranchContext) IsIfConditionBranchContext() {}

func NewIfConditionBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfConditionBranchContext {
	var p = new(IfConditionBranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_ifConditionBranch

	return p
}

func (s *IfConditionBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *IfConditionBranchContext) GetCond() IExpressionContext { return s.cond }

func (s *IfConditionBranchContext) GetBody() IBlockContext { return s.body }

func (s *IfConditionBranchContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *IfConditionBranchContext) SetBody(v IBlockContext) { s.body = v }

func (s *IfConditionBranchContext) KW_IF() antlr.TerminalNode {
	return s.GetToken(DZParserKW_IF, 0)
}

func (s *IfConditionBranchContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *IfConditionBranchContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *IfConditionBranchContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfConditionBranchContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfConditionBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfConditionBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfConditionBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterIfConditionBranch(s)
	}
}

func (s *IfConditionBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitIfConditionBranch(s)
	}
}

func (p *DZParser) IfConditionBranch() (localctx IIfConditionBranchContext) {
	localctx = NewIfConditionBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, DZParserRULE_ifConditionBranch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Match(DZParserKW_IF)
	}
	{
		p.SetState(281)

		var _x = p.expression(0)

		localctx.(*IfConditionBranchContext).cond = _x
	}
	{
		p.SetState(282)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(283)

		var _x = p.Block()

		localctx.(*IfConditionBranchContext).body = _x
	}
	{
		p.SetState(284)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// IElifConditionBranchContext is an interface to support dynamic dispatch.
type IElifConditionBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// GetBody returns the body rule contexts.
	GetBody() IBlockContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// SetBody sets the body rule contexts.
	SetBody(IBlockContext)

	// IsElifConditionBranchContext differentiates from other interfaces.
	IsElifConditionBranchContext()
}

type ElifConditionBranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   IExpressionContext
	body   IBlockContext
}

func NewEmptyElifConditionBranchContext() *ElifConditionBranchContext {
	var p = new(ElifConditionBranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_elifConditionBranch
	return p
}

func (*ElifConditionBranchContext) IsElifConditionBranchContext() {}

func NewElifConditionBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifConditionBranchContext {
	var p = new(ElifConditionBranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_elifConditionBranch

	return p
}

func (s *ElifConditionBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifConditionBranchContext) GetCond() IExpressionContext { return s.cond }

func (s *ElifConditionBranchContext) GetBody() IBlockContext { return s.body }

func (s *ElifConditionBranchContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *ElifConditionBranchContext) SetBody(v IBlockContext) { s.body = v }

func (s *ElifConditionBranchContext) KW_ELIF() antlr.TerminalNode {
	return s.GetToken(DZParserKW_ELIF, 0)
}

func (s *ElifConditionBranchContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *ElifConditionBranchContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *ElifConditionBranchContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElifConditionBranchContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElifConditionBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifConditionBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElifConditionBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterElifConditionBranch(s)
	}
}

func (s *ElifConditionBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitElifConditionBranch(s)
	}
}

func (p *DZParser) ElifConditionBranch() (localctx IElifConditionBranchContext) {
	localctx = NewElifConditionBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, DZParserRULE_elifConditionBranch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(DZParserKW_ELIF)
	}
	{
		p.SetState(287)

		var _x = p.expression(0)

		localctx.(*ElifConditionBranchContext).cond = _x
	}
	{
		p.SetState(288)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(289)

		var _x = p.Block()

		localctx.(*ElifConditionBranchContext).body = _x
	}
	{
		p.SetState(290)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// IElseConditionBranchContext is an interface to support dynamic dispatch.
type IElseConditionBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBody returns the body rule contexts.
	GetBody() IBlockContext

	// SetBody sets the body rule contexts.
	SetBody(IBlockContext)

	// IsElseConditionBranchContext differentiates from other interfaces.
	IsElseConditionBranchContext()
}

type ElseConditionBranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	body   IBlockContext
}

func NewEmptyElseConditionBranchContext() *ElseConditionBranchContext {
	var p = new(ElseConditionBranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_elseConditionBranch
	return p
}

func (*ElseConditionBranchContext) IsElseConditionBranchContext() {}

func NewElseConditionBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseConditionBranchContext {
	var p = new(ElseConditionBranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_elseConditionBranch

	return p
}

func (s *ElseConditionBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseConditionBranchContext) GetBody() IBlockContext { return s.body }

func (s *ElseConditionBranchContext) SetBody(v IBlockContext) { s.body = v }

func (s *ElseConditionBranchContext) KW_ELSE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_ELSE, 0)
}

func (s *ElseConditionBranchContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *ElseConditionBranchContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *ElseConditionBranchContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseConditionBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseConditionBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseConditionBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterElseConditionBranch(s)
	}
}

func (s *ElseConditionBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitElseConditionBranch(s)
	}
}

func (p *DZParser) ElseConditionBranch() (localctx IElseConditionBranchContext) {
	localctx = NewElseConditionBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, DZParserRULE_elseConditionBranch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(DZParserKW_ELSE)
	}
	{
		p.SetState(293)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(294)

		var _x = p.Block()

		localctx.(*ElseConditionBranchContext).body = _x
	}
	{
		p.SetState(295)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) TrueLoop() ITrueLoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrueLoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrueLoopContext)
}

func (s *LoopContext) WhileLoop() IWhileLoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileLoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileLoopContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *DZParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, DZParserRULE_loop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(299)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_LOOP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.TrueLoop()
		}

	case DZParserKW_WHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.WhileLoop()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITrueLoopContext is an interface to support dynamic dispatch.
type ITrueLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBody returns the body rule contexts.
	GetBody() IBlockContext

	// SetBody sets the body rule contexts.
	SetBody(IBlockContext)

	// IsTrueLoopContext differentiates from other interfaces.
	IsTrueLoopContext()
}

type TrueLoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	body   IBlockContext
}

func NewEmptyTrueLoopContext() *TrueLoopContext {
	var p = new(TrueLoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_trueLoop
	return p
}

func (*TrueLoopContext) IsTrueLoopContext() {}

func NewTrueLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrueLoopContext {
	var p = new(TrueLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_trueLoop

	return p
}

func (s *TrueLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *TrueLoopContext) GetBody() IBlockContext { return s.body }

func (s *TrueLoopContext) SetBody(v IBlockContext) { s.body = v }

func (s *TrueLoopContext) KW_LOOP() antlr.TerminalNode {
	return s.GetToken(DZParserKW_LOOP, 0)
}

func (s *TrueLoopContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *TrueLoopContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *TrueLoopContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TrueLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrueLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterTrueLoop(s)
	}
}

func (s *TrueLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitTrueLoop(s)
	}
}

func (p *DZParser) TrueLoop() (localctx ITrueLoopContext) {
	localctx = NewTrueLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, DZParserRULE_trueLoop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(DZParserKW_LOOP)
	}
	{
		p.SetState(302)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(303)

		var _x = p.Block()

		localctx.(*TrueLoopContext).body = _x
	}
	{
		p.SetState(304)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// IWhileLoopContext is an interface to support dynamic dispatch.
type IWhileLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// GetBody returns the body rule contexts.
	GetBody() IBlockContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// SetBody sets the body rule contexts.
	SetBody(IBlockContext)

	// IsWhileLoopContext differentiates from other interfaces.
	IsWhileLoopContext()
}

type WhileLoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   IExpressionContext
	body   IBlockContext
}

func NewEmptyWhileLoopContext() *WhileLoopContext {
	var p = new(WhileLoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_whileLoop
	return p
}

func (*WhileLoopContext) IsWhileLoopContext() {}

func NewWhileLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoopContext {
	var p = new(WhileLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_whileLoop

	return p
}

func (s *WhileLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileLoopContext) GetCond() IExpressionContext { return s.cond }

func (s *WhileLoopContext) GetBody() IBlockContext { return s.body }

func (s *WhileLoopContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *WhileLoopContext) SetBody(v IBlockContext) { s.body = v }

func (s *WhileLoopContext) KW_WHILE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_WHILE, 0)
}

func (s *WhileLoopContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *WhileLoopContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *WhileLoopContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileLoopContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterWhileLoop(s)
	}
}

func (s *WhileLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitWhileLoop(s)
	}
}

func (p *DZParser) WhileLoop() (localctx IWhileLoopContext) {
	localctx = NewWhileLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, DZParserRULE_whileLoop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(DZParserKW_WHILE)
	}
	{
		p.SetState(307)

		var _x = p.expression(0)

		localctx.(*WhileLoopContext).cond = _x
	}
	{
		p.SetState(308)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(309)

		var _x = p.Block()

		localctx.(*WhileLoopContext).body = _x
	}
	{
		p.SetState(310)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) KW_BREAK() antlr.TerminalNode {
	return s.GetToken(DZParserKW_BREAK, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (p *DZParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, DZParserRULE_breakStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(DZParserKW_BREAK)
	}

	return localctx
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_continueStatement
	return p
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) KW_CONTINUE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_CONTINUE, 0)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (p *DZParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, DZParserRULE_continueStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(DZParserKW_CONTINUE)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	value  IExpressionContext
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) GetName() antlr.Token { return s.name }

func (s *DeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *DeclarationContext) GetValue() IExpressionContext { return s.value }

func (s *DeclarationContext) SetValue(v IExpressionContext) { s.value = v }

func (s *DeclarationContext) KW_LET() antlr.TerminalNode {
	return s.GetToken(DZParserKW_LET, 0)
}

func (s *DeclarationContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *DeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *DeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *DZParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, DZParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Match(DZParserKW_LET)
	}
	{
		p.SetState(317)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*DeclarationContext).name = _m
	}
	{
		p.SetState(318)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(319)

		var _x = p.expression(0)

		localctx.(*DeclarationContext).value = _x
	}

	return localctx
}

// IProcCallContext is an interface to support dynamic dispatch.
type IProcCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetProcName returns the procName token.
	GetProcName() antlr.Token

	// SetProcName sets the procName token.
	SetProcName(antlr.Token)

	// IsProcCallContext differentiates from other interfaces.
	IsProcCallContext()
}

type ProcCallContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	procName antlr.Token
}

func NewEmptyProcCallContext() *ProcCallContext {
	var p = new(ProcCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_procCall
	return p
}

func (*ProcCallContext) IsProcCallContext() {}

func NewProcCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcCallContext {
	var p = new(ProcCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_procCall

	return p
}

func (s *ProcCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcCallContext) GetProcName() antlr.Token { return s.procName }

func (s *ProcCallContext) SetProcName(v antlr.Token) { s.procName = v }

func (s *ProcCallContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_PRT, 0)
}

func (s *ProcCallContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_PRT, 0)
}

func (s *ProcCallContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *ProcCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *ProcCallContext) AllProcParam() []IProcParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProcParamContext)(nil)).Elem())
	var tst = make([]IProcParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProcParamContext)
		}
	}

	return tst
}

func (s *ProcCallContext) ProcParam(i int) IProcParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProcParamContext)
}

func (s *ProcCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DZParserCOMMA)
}

func (s *ProcCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DZParserCOMMA, i)
}

func (s *ProcCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterProcCall(s)
	}
}

func (s *ProcCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitProcCall(s)
	}
}

func (p *DZParser) ProcCall() (localctx IProcCallContext) {
	localctx = NewProcCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, DZParserRULE_procCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcCallContext).procName = _m
	}
	{
		p.SetState(322)
		p.Match(DZParserLEFT_PRT)
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(DZParserLEFT_PRT-30))|(1<<(DZParserLEFT_BRK-30))|(1<<(DZParserSUB-30))|(1<<(DZParserNOT-30)))) != 0) || (((_la-62)&-(0x1f+1)) == 0 && ((1<<uint((_la-62)))&((1<<(DZParserINT_CONST-62))|(1<<(DZParserTRUE-62))|(1<<(DZParserFALSE-62))|(1<<(DZParserCONST-62))|(1<<(DZParserIDENTIFIER-62)))) != 0) {
		{
			p.SetState(323)
			p.ProcParam()
		}
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserCOMMA {
			{
				p.SetState(324)
				p.Match(DZParserCOMMA)
			}
			{
				p.SetState(325)
				p.ProcParam()
			}

			p.SetState(330)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(333)
		p.Match(DZParserRIGHT_PRT)
	}
	{
		p.SetState(334)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// IProcParamContext is an interface to support dynamic dispatch.
type IProcParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsProcParamContext differentiates from other interfaces.
	IsProcParamContext()
}

type ProcParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  IExpressionContext
}

func NewEmptyProcParamContext() *ProcParamContext {
	var p = new(ProcParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_procParam
	return p
}

func (*ProcParamContext) IsProcParamContext() {}

func NewProcParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcParamContext {
	var p = new(ProcParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_procParam

	return p
}

func (s *ProcParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcParamContext) GetValue() IExpressionContext { return s.value }

func (s *ProcParamContext) SetValue(v IExpressionContext) { s.value = v }

func (s *ProcParamContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProcParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterProcParam(s)
	}
}

func (s *ProcParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitProcParam(s)
	}
}

func (p *DZParser) ProcParam() (localctx IProcParamContext) {
	localctx = NewProcParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, DZParserRULE_procParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)

		var _x = p.expression(0)

		localctx.(*ProcParamContext).value = _x
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVarName returns the varName token.
	GetVarName() antlr.Token

	// GetIntValue returns the intValue token.
	GetIntValue() antlr.Token

	// GetBoolValue returns the boolValue token.
	GetBoolValue() antlr.Token

	// GetConstName returns the constName token.
	GetConstName() antlr.Token

	// GetFuncName returns the funcName token.
	GetFuncName() antlr.Token

	// SetVarName sets the varName token.
	SetVarName(antlr.Token)

	// SetIntValue sets the intValue token.
	SetIntValue(antlr.Token)

	// SetBoolValue sets the boolValue token.
	SetBoolValue(antlr.Token)

	// SetConstName sets the constName token.
	SetConstName(antlr.Token)

	// SetFuncName sets the funcName token.
	SetFuncName(antlr.Token)

	// GetLBinExpr returns the lBinExpr rule contexts.
	GetLBinExpr() IExpressionContext

	// GetBrkExpr returns the brkExpr rule contexts.
	GetBrkExpr() IExpressionContext

	// GetPrtExpr returns the prtExpr rule contexts.
	GetPrtExpr() IExpressionContext

	// GetUnOp returns the unOp rule contexts.
	GetUnOp() IUnaryOperatorContext

	// GetUnExpr returns the unExpr rule contexts.
	GetUnExpr() IExpressionContext

	// GetBinOp returns the binOp rule contexts.
	GetBinOp() IBinaryOperatorContext

	// GetRBinExpr returns the rBinExpr rule contexts.
	GetRBinExpr() IExpressionContext

	// SetLBinExpr sets the lBinExpr rule contexts.
	SetLBinExpr(IExpressionContext)

	// SetBrkExpr sets the brkExpr rule contexts.
	SetBrkExpr(IExpressionContext)

	// SetPrtExpr sets the prtExpr rule contexts.
	SetPrtExpr(IExpressionContext)

	// SetUnOp sets the unOp rule contexts.
	SetUnOp(IUnaryOperatorContext)

	// SetUnExpr sets the unExpr rule contexts.
	SetUnExpr(IExpressionContext)

	// SetBinOp sets the binOp rule contexts.
	SetBinOp(IBinaryOperatorContext)

	// SetRBinExpr sets the rBinExpr rule contexts.
	SetRBinExpr(IExpressionContext)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	lBinExpr  IExpressionContext
	varName   antlr.Token
	intValue  antlr.Token
	boolValue antlr.Token
	constName antlr.Token
	brkExpr   IExpressionContext
	prtExpr   IExpressionContext
	funcName  antlr.Token
	unOp      IUnaryOperatorContext
	unExpr    IExpressionContext
	binOp     IBinaryOperatorContext
	rBinExpr  IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetVarName() antlr.Token { return s.varName }

func (s *ExpressionContext) GetIntValue() antlr.Token { return s.intValue }

func (s *ExpressionContext) GetBoolValue() antlr.Token { return s.boolValue }

func (s *ExpressionContext) GetConstName() antlr.Token { return s.constName }

func (s *ExpressionContext) GetFuncName() antlr.Token { return s.funcName }

func (s *ExpressionContext) SetVarName(v antlr.Token) { s.varName = v }

func (s *ExpressionContext) SetIntValue(v antlr.Token) { s.intValue = v }

func (s *ExpressionContext) SetBoolValue(v antlr.Token) { s.boolValue = v }

func (s *ExpressionContext) SetConstName(v antlr.Token) { s.constName = v }

func (s *ExpressionContext) SetFuncName(v antlr.Token) { s.funcName = v }

func (s *ExpressionContext) GetLBinExpr() IExpressionContext { return s.lBinExpr }

func (s *ExpressionContext) GetBrkExpr() IExpressionContext { return s.brkExpr }

func (s *ExpressionContext) GetPrtExpr() IExpressionContext { return s.prtExpr }

func (s *ExpressionContext) GetUnOp() IUnaryOperatorContext { return s.unOp }

func (s *ExpressionContext) GetUnExpr() IExpressionContext { return s.unExpr }

func (s *ExpressionContext) GetBinOp() IBinaryOperatorContext { return s.binOp }

func (s *ExpressionContext) GetRBinExpr() IExpressionContext { return s.rBinExpr }

func (s *ExpressionContext) SetLBinExpr(v IExpressionContext) { s.lBinExpr = v }

func (s *ExpressionContext) SetBrkExpr(v IExpressionContext) { s.brkExpr = v }

func (s *ExpressionContext) SetPrtExpr(v IExpressionContext) { s.prtExpr = v }

func (s *ExpressionContext) SetUnOp(v IUnaryOperatorContext) { s.unOp = v }

func (s *ExpressionContext) SetUnExpr(v IExpressionContext) { s.unExpr = v }

func (s *ExpressionContext) SetBinOp(v IBinaryOperatorContext) { s.binOp = v }

func (s *ExpressionContext) SetRBinExpr(v IExpressionContext) { s.rBinExpr = v }

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *ExpressionContext) INT_CONST() antlr.TerminalNode {
	return s.GetToken(DZParserINT_CONST, 0)
}

func (s *ExpressionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(DZParserTRUE, 0)
}

func (s *ExpressionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(DZParserFALSE, 0)
}

func (s *ExpressionContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
}

func (s *ExpressionContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRK, 0)
}

func (s *ExpressionContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRK, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_PRT, 0)
}

func (s *ExpressionContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_PRT, 0)
}

func (s *ExpressionContext) AllFuncParam() []IFuncParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncParamContext)(nil)).Elem())
	var tst = make([]IFuncParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncParamContext)
		}
	}

	return tst
}

func (s *ExpressionContext) FuncParam(i int) IFuncParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncParamContext)
}

func (s *ExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DZParserCOMMA)
}

func (s *ExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DZParserCOMMA, i)
}

func (s *ExpressionContext) UnaryOperator() IUnaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *ExpressionContext) BinaryOperator() IBinaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DZParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *DZParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 94
	p.EnterRecursionRule(localctx, 94, DZParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(339)

			var _m = p.Match(DZParserIDENTIFIER)

			localctx.(*ExpressionContext).varName = _m
		}

	case 2:
		{
			p.SetState(340)

			var _m = p.Match(DZParserINT_CONST)

			localctx.(*ExpressionContext).intValue = _m
		}

	case 3:
		{
			p.SetState(341)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).boolValue = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DZParserTRUE || _la == DZParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).boolValue = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		{
			p.SetState(342)

			var _m = p.Match(DZParserCONST)

			localctx.(*ExpressionContext).constName = _m
		}

	case 5:
		{
			p.SetState(343)
			p.Match(DZParserLEFT_BRK)
		}
		{
			p.SetState(344)

			var _x = p.expression(0)

			localctx.(*ExpressionContext).brkExpr = _x
		}
		{
			p.SetState(345)
			p.Match(DZParserRIGHT_BRK)
		}

	case 6:
		{
			p.SetState(347)
			p.Match(DZParserLEFT_PRT)
		}
		{
			p.SetState(348)

			var _x = p.expression(0)

			localctx.(*ExpressionContext).prtExpr = _x
		}
		{
			p.SetState(349)
			p.Match(DZParserRIGHT_PRT)
		}

	case 7:
		{
			p.SetState(351)

			var _m = p.Match(DZParserIDENTIFIER)

			localctx.(*ExpressionContext).funcName = _m
		}
		{
			p.SetState(352)
			p.Match(DZParserLEFT_PRT)
		}
		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(DZParserLEFT_PRT-30))|(1<<(DZParserLEFT_BRK-30))|(1<<(DZParserSUB-30))|(1<<(DZParserNOT-30)))) != 0) || (((_la-62)&-(0x1f+1)) == 0 && ((1<<uint((_la-62)))&((1<<(DZParserINT_CONST-62))|(1<<(DZParserTRUE-62))|(1<<(DZParserFALSE-62))|(1<<(DZParserCONST-62))|(1<<(DZParserIDENTIFIER-62)))) != 0) {
			{
				p.SetState(353)
				p.FuncParam()
			}
			p.SetState(358)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DZParserCOMMA {
				{
					p.SetState(354)
					p.Match(DZParserCOMMA)
				}
				{
					p.SetState(355)
					p.FuncParam()
				}

				p.SetState(360)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(363)
			p.Match(DZParserRIGHT_PRT)
		}

	case 8:
		{
			p.SetState(364)

			var _x = p.UnaryOperator()

			localctx.(*ExpressionContext).unOp = _x
		}
		{
			p.SetState(365)

			var _x = p.expression(1)

			localctx.(*ExpressionContext).unExpr = _x
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			localctx.(*ExpressionContext).lBinExpr = _prevctx
			p.PushNewRecursionContext(localctx, _startState, DZParserRULE_expression)
			p.SetState(369)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(370)

				var _x = p.BinaryOperator()

				localctx.(*ExpressionContext).binOp = _x
			}
			{
				p.SetState(371)

				var _x = p.expression(3)

				localctx.(*ExpressionContext).rBinExpr = _x
			}

		}
		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncParamContext is an interface to support dynamic dispatch.
type IFuncParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsFuncParamContext differentiates from other interfaces.
	IsFuncParamContext()
}

type FuncParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  IExpressionContext
}

func NewEmptyFuncParamContext() *FuncParamContext {
	var p = new(FuncParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_funcParam
	return p
}

func (*FuncParamContext) IsFuncParamContext() {}

func NewFuncParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamContext {
	var p = new(FuncParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_funcParam

	return p
}

func (s *FuncParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamContext) GetValue() IExpressionContext { return s.value }

func (s *FuncParamContext) SetValue(v IExpressionContext) { s.value = v }

func (s *FuncParamContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterFuncParam(s)
	}
}

func (s *FuncParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitFuncParam(s)
	}
}

func (p *DZParser) FuncParam() (localctx IFuncParamContext) {
	localctx = NewFuncParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, DZParserRULE_funcParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)

		var _x = p.expression(0)

		localctx.(*FuncParamContext).value = _x
	}

	return localctx
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(DZParserSUB, 0)
}

func (s *UnaryOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(DZParserNOT, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (p *DZParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, DZParserRULE_unaryOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DZParserSUB || _la == DZParserNOT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryOperatorContext is an interface to support dynamic dispatch.
type IBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOperatorContext differentiates from other interfaces.
	IsBinaryOperatorContext()
}

type BinaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperatorContext() *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_binaryOperator
	return p
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperatorContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *BinaryOperatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(DZParserADD, 0)
}

func (s *BinaryOperatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(DZParserSUB, 0)
}

func (s *BinaryOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(DZParserMUL, 0)
}

func (s *BinaryOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(DZParserDIV, 0)
}

func (s *BinaryOperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(DZParserMOD, 0)
}

func (s *BinaryOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(DZParserAND, 0)
}

func (s *BinaryOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(DZParserOR, 0)
}

func (s *BinaryOperatorContext) XOR() antlr.TerminalNode {
	return s.GetToken(DZParserXOR, 0)
}

func (s *BinaryOperatorContext) SHL() antlr.TerminalNode {
	return s.GetToken(DZParserSHL, 0)
}

func (s *BinaryOperatorContext) SHR() antlr.TerminalNode {
	return s.GetToken(DZParserSHR, 0)
}

func (s *BinaryOperatorContext) ADD_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserADD_ASGN, 0)
}

func (s *BinaryOperatorContext) SUB_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserSUB_ASGN, 0)
}

func (s *BinaryOperatorContext) MUL_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserMUL_ASGN, 0)
}

func (s *BinaryOperatorContext) DIV_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserDIV_ASGN, 0)
}

func (s *BinaryOperatorContext) MOD_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserMOD_ASGN, 0)
}

func (s *BinaryOperatorContext) AND_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserAND_ASGN, 0)
}

func (s *BinaryOperatorContext) OR_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserOR_ASGN, 0)
}

func (s *BinaryOperatorContext) XOR_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserXOR_ASGN, 0)
}

func (s *BinaryOperatorContext) SHL_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserSHL_ASGN, 0)
}

func (s *BinaryOperatorContext) SHR_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserSHR_ASGN, 0)
}

func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterBinaryOperator(s)
	}
}

func (s *BinaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitBinaryOperator(s)
	}
}

func (p *DZParser) BinaryOperator() (localctx IBinaryOperatorContext) {
	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, DZParserRULE_binaryOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(DZParserADD-39))|(1<<(DZParserSUB-39))|(1<<(DZParserMUL-39))|(1<<(DZParserDIV-39))|(1<<(DZParserMOD-39))|(1<<(DZParserSHL-39))|(1<<(DZParserSHR-39))|(1<<(DZParserAND-39))|(1<<(DZParserOR-39))|(1<<(DZParserXOR-39))|(1<<(DZParserASGN-39))|(1<<(DZParserADD_ASGN-39))|(1<<(DZParserSUB_ASGN-39))|(1<<(DZParserMUL_ASGN-39))|(1<<(DZParserDIV_ASGN-39))|(1<<(DZParserMOD_ASGN-39))|(1<<(DZParserSHL_ASGN-39))|(1<<(DZParserSHR_ASGN-39))|(1<<(DZParserAND_ASGN-39))|(1<<(DZParserOR_ASGN-39))|(1<<(DZParserXOR_ASGN-39)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  IExpressionContext
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) GetValue() IExpressionContext { return s.value }

func (s *ReturnStatementContext) SetValue(v IExpressionContext) { s.value = v }

func (s *ReturnStatementContext) KW_RETURN() antlr.TerminalNode {
	return s.GetToken(DZParserKW_RETURN, 0)
}

func (s *ReturnStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *DZParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, DZParserRULE_returnStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)
			p.Match(DZParserKW_RETURN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.Match(DZParserKW_RETURN)
		}
		{
			p.SetState(386)

			var _x = p.expression(0)

			localctx.(*ReturnStatementContext).value = _x
		}
		{
			p.SetState(387)
			p.Match(DZParserSEMICOLON)
		}

	}

	return localctx
}

func (p *DZParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 47:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DZParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
