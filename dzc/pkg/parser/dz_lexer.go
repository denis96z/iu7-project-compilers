// Code generated from /home/denis/Documents/iu7-project-compilers/dzc/DZ.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 70, 451,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3,
	42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46,
	3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3,
	51, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 55,
	3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3,
	58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61,
	3, 62, 5, 62, 378, 10, 62, 3, 62, 6, 62, 381, 10, 62, 13, 62, 14, 62, 382,
	3, 63, 5, 63, 386, 10, 63, 3, 63, 6, 63, 389, 10, 63, 13, 63, 14, 63, 390,
	3, 63, 3, 63, 6, 63, 395, 10, 63, 13, 63, 14, 63, 396, 5, 63, 399, 10,
	63, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65,
	3, 65, 3, 66, 3, 66, 5, 66, 414, 10, 66, 3, 66, 7, 66, 417, 10, 66, 12,
	66, 14, 66, 420, 11, 66, 3, 67, 3, 67, 5, 67, 424, 10, 67, 3, 67, 7, 67,
	427, 10, 67, 12, 67, 14, 67, 430, 11, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3,
	68, 5, 68, 437, 10, 68, 3, 68, 7, 68, 440, 10, 68, 12, 68, 14, 68, 443,
	11, 68, 3, 69, 6, 69, 446, 10, 69, 13, 69, 14, 69, 447, 3, 69, 3, 69, 2,
	2, 70, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57,
	30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75,
	39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93,
	48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56,
	111, 57, 113, 58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64,
	127, 65, 129, 66, 131, 67, 133, 68, 135, 69, 137, 70, 3, 2, 12, 3, 2, 47,
	47, 3, 2, 50, 59, 3, 2, 48, 48, 3, 2, 67, 92, 3, 2, 97, 97, 4, 2, 50, 59,
	67, 92, 3, 2, 99, 124, 4, 2, 50, 59, 99, 124, 3, 2, 118, 118, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 463, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2,
	2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3,
	2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61,
	3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2,
	69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2,
	2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2,
	2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2,
	2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3,
	2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2,
	107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2,
	2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121,
	3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2,
	2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3,
	2, 2, 2, 2, 137, 3, 2, 2, 2, 3, 139, 3, 2, 2, 2, 5, 143, 3, 2, 2, 2, 7,
	148, 3, 2, 2, 2, 9, 152, 3, 2, 2, 2, 11, 158, 3, 2, 2, 2, 13, 163, 3, 2,
	2, 2, 15, 166, 3, 2, 2, 2, 17, 171, 3, 2, 2, 2, 19, 176, 3, 2, 2, 2, 21,
	181, 3, 2, 2, 2, 23, 186, 3, 2, 2, 2, 25, 192, 3, 2, 2, 2, 27, 196, 3,
	2, 2, 2, 29, 200, 3, 2, 2, 2, 31, 207, 3, 2, 2, 2, 33, 212, 3, 2, 2, 2,
	35, 218, 3, 2, 2, 2, 37, 225, 3, 2, 2, 2, 39, 230, 3, 2, 2, 2, 41, 235,
	3, 2, 2, 2, 43, 241, 3, 2, 2, 2, 45, 247, 3, 2, 2, 2, 47, 253, 3, 2, 2,
	2, 49, 259, 3, 2, 2, 2, 51, 265, 3, 2, 2, 2, 53, 271, 3, 2, 2, 2, 55, 278,
	3, 2, 2, 2, 57, 285, 3, 2, 2, 2, 59, 292, 3, 2, 2, 2, 61, 300, 3, 2, 2,
	2, 63, 302, 3, 2, 2, 2, 65, 304, 3, 2, 2, 2, 67, 306, 3, 2, 2, 2, 69, 308,
	3, 2, 2, 2, 71, 310, 3, 2, 2, 2, 73, 312, 3, 2, 2, 2, 75, 314, 3, 2, 2,
	2, 77, 316, 3, 2, 2, 2, 79, 318, 3, 2, 2, 2, 81, 320, 3, 2, 2, 2, 83, 322,
	3, 2, 2, 2, 85, 324, 3, 2, 2, 2, 87, 326, 3, 2, 2, 2, 89, 328, 3, 2, 2,
	2, 91, 331, 3, 2, 2, 2, 93, 334, 3, 2, 2, 2, 95, 336, 3, 2, 2, 2, 97, 338,
	3, 2, 2, 2, 99, 340, 3, 2, 2, 2, 101, 342, 3, 2, 2, 2, 103, 345, 3, 2,
	2, 2, 105, 348, 3, 2, 2, 2, 107, 351, 3, 2, 2, 2, 109, 354, 3, 2, 2, 2,
	111, 357, 3, 2, 2, 2, 113, 361, 3, 2, 2, 2, 115, 365, 3, 2, 2, 2, 117,
	368, 3, 2, 2, 2, 119, 371, 3, 2, 2, 2, 121, 374, 3, 2, 2, 2, 123, 377,
	3, 2, 2, 2, 125, 385, 3, 2, 2, 2, 127, 400, 3, 2, 2, 2, 129, 405, 3, 2,
	2, 2, 131, 411, 3, 2, 2, 2, 133, 421, 3, 2, 2, 2, 135, 434, 3, 2, 2, 2,
	137, 445, 3, 2, 2, 2, 139, 140, 7, 114, 2, 2, 140, 141, 7, 109, 2, 2, 141,
	142, 7, 105, 2, 2, 142, 4, 3, 2, 2, 2, 143, 144, 7, 118, 2, 2, 144, 145,
	7, 123, 2, 2, 145, 146, 7, 114, 2, 2, 146, 147, 7, 103, 2, 2, 147, 6, 3,
	2, 2, 2, 148, 149, 7, 104, 2, 2, 149, 150, 7, 113, 2, 2, 150, 151, 7, 116,
	2, 2, 151, 8, 3, 2, 2, 2, 152, 153, 7, 121, 2, 2, 153, 154, 7, 106, 2,
	2, 154, 155, 7, 107, 2, 2, 155, 156, 7, 110, 2, 2, 156, 157, 7, 103, 2,
	2, 157, 10, 3, 2, 2, 2, 158, 159, 7, 110, 2, 2, 159, 160, 7, 113, 2, 2,
	160, 161, 7, 113, 2, 2, 161, 162, 7, 114, 2, 2, 162, 12, 3, 2, 2, 2, 163,
	164, 7, 107, 2, 2, 164, 165, 7, 104, 2, 2, 165, 14, 3, 2, 2, 2, 166, 167,
	7, 103, 2, 2, 167, 168, 7, 110, 2, 2, 168, 169, 7, 107, 2, 2, 169, 170,
	7, 104, 2, 2, 170, 16, 3, 2, 2, 2, 171, 172, 7, 103, 2, 2, 172, 173, 7,
	110, 2, 2, 173, 174, 7, 117, 2, 2, 174, 175, 7, 103, 2, 2, 175, 18, 3,
	2, 2, 2, 176, 177, 7, 114, 2, 2, 177, 178, 7, 116, 2, 2, 178, 179, 7, 113,
	2, 2, 179, 180, 7, 101, 2, 2, 180, 20, 3, 2, 2, 2, 181, 182, 7, 104, 2,
	2, 182, 183, 7, 119, 2, 2, 183, 184, 7, 112, 2, 2, 184, 185, 7, 101, 2,
	2, 185, 22, 3, 2, 2, 2, 186, 187, 7, 101, 2, 2, 187, 188, 7, 113, 2, 2,
	188, 189, 7, 112, 2, 2, 189, 190, 7, 117, 2, 2, 190, 191, 7, 118, 2, 2,
	191, 24, 3, 2, 2, 2, 192, 193, 7, 110, 2, 2, 193, 194, 7, 103, 2, 2, 194,
	195, 7, 118, 2, 2, 195, 26, 3, 2, 2, 2, 196, 197, 7, 111, 2, 2, 197, 198,
	7, 119, 2, 2, 198, 199, 7, 118, 2, 2, 199, 28, 3, 2, 2, 2, 200, 201, 7,
	117, 2, 2, 201, 202, 7, 118, 2, 2, 202, 203, 7, 116, 2, 2, 203, 204, 7,
	119, 2, 2, 204, 205, 7, 101, 2, 2, 205, 206, 7, 118, 2, 2, 206, 30, 3,
	2, 2, 2, 207, 208, 7, 103, 2, 2, 208, 209, 7, 112, 2, 2, 209, 210, 7, 119,
	2, 2, 210, 211, 7, 111, 2, 2, 211, 32, 3, 2, 2, 2, 212, 213, 7, 119, 2,
	2, 213, 214, 7, 112, 2, 2, 214, 215, 7, 107, 2, 2, 215, 216, 7, 113, 2,
	2, 216, 217, 7, 112, 2, 2, 217, 34, 3, 2, 2, 2, 218, 219, 7, 116, 2, 2,
	219, 220, 7, 103, 2, 2, 220, 221, 7, 118, 2, 2, 221, 222, 7, 119, 2, 2,
	222, 223, 7, 116, 2, 2, 223, 224, 7, 112, 2, 2, 224, 36, 3, 2, 2, 2, 225,
	226, 7, 107, 2, 2, 226, 227, 7, 58, 2, 2, 227, 228, 7, 97, 2, 2, 228, 229,
	7, 118, 2, 2, 229, 38, 3, 2, 2, 2, 230, 231, 7, 119, 2, 2, 231, 232, 7,
	58, 2, 2, 232, 233, 7, 97, 2, 2, 233, 234, 7, 118, 2, 2, 234, 40, 3, 2,
	2, 2, 235, 236, 7, 107, 2, 2, 236, 237, 7, 51, 2, 2, 237, 238, 7, 56, 2,
	2, 238, 239, 7, 97, 2, 2, 239, 240, 7, 118, 2, 2, 240, 42, 3, 2, 2, 2,
	241, 242, 7, 119, 2, 2, 242, 243, 7, 51, 2, 2, 243, 244, 7, 56, 2, 2, 244,
	245, 7, 97, 2, 2, 245, 246, 7, 118, 2, 2, 246, 44, 3, 2, 2, 2, 247, 248,
	7, 107, 2, 2, 248, 249, 7, 53, 2, 2, 249, 250, 7, 52, 2, 2, 250, 251, 7,
	97, 2, 2, 251, 252, 7, 118, 2, 2, 252, 46, 3, 2, 2, 2, 253, 254, 7, 119,
	2, 2, 254, 255, 7, 53, 2, 2, 255, 256, 7, 52, 2, 2, 256, 257, 7, 97, 2,
	2, 257, 258, 7, 118, 2, 2, 258, 48, 3, 2, 2, 2, 259, 260, 7, 107, 2, 2,
	260, 261, 7, 56, 2, 2, 261, 262, 7, 54, 2, 2, 262, 263, 7, 97, 2, 2, 263,
	264, 7, 118, 2, 2, 264, 50, 3, 2, 2, 2, 265, 266, 7, 119, 2, 2, 266, 267,
	7, 56, 2, 2, 267, 268, 7, 54, 2, 2, 268, 269, 7, 97, 2, 2, 269, 270, 7,
	118, 2, 2, 270, 52, 3, 2, 2, 2, 271, 272, 7, 101, 2, 2, 272, 273, 7, 106,
	2, 2, 273, 274, 7, 99, 2, 2, 274, 275, 7, 116, 2, 2, 275, 276, 7, 97, 2,
	2, 276, 277, 7, 118, 2, 2, 277, 54, 3, 2, 2, 2, 278, 279, 7, 100, 2, 2,
	279, 280, 7, 113, 2, 2, 280, 281, 7, 113, 2, 2, 281, 282, 7, 110, 2, 2,
	282, 283, 7, 97, 2, 2, 283, 284, 7, 118, 2, 2, 284, 56, 3, 2, 2, 2, 285,
	286, 7, 117, 2, 2, 286, 287, 7, 107, 2, 2, 287, 288, 7, 124, 2, 2, 288,
	289, 7, 103, 2, 2, 289, 290, 7, 97, 2, 2, 290, 291, 7, 118, 2, 2, 291,
	58, 3, 2, 2, 2, 292, 293, 7, 117, 2, 2, 293, 294, 7, 117, 2, 2, 294, 295,
	7, 107, 2, 2, 295, 296, 7, 124, 2, 2, 296, 297, 7, 103, 2, 2, 297, 298,
	7, 97, 2, 2, 298, 299, 7, 118, 2, 2, 299, 60, 3, 2, 2, 2, 300, 301, 7,
	42, 2, 2, 301, 62, 3, 2, 2, 2, 302, 303, 7, 43, 2, 2, 303, 64, 3, 2, 2,
	2, 304, 305, 7, 125, 2, 2, 305, 66, 3, 2, 2, 2, 306, 307, 7, 127, 2, 2,
	307, 68, 3, 2, 2, 2, 308, 309, 7, 93, 2, 2, 309, 70, 3, 2, 2, 2, 310, 311,
	7, 95, 2, 2, 311, 72, 3, 2, 2, 2, 312, 313, 7, 60, 2, 2, 313, 74, 3, 2,
	2, 2, 314, 315, 7, 46, 2, 2, 315, 76, 3, 2, 2, 2, 316, 317, 7, 61, 2, 2,
	317, 78, 3, 2, 2, 2, 318, 319, 7, 45, 2, 2, 319, 80, 3, 2, 2, 2, 320, 321,
	7, 47, 2, 2, 321, 82, 3, 2, 2, 2, 322, 323, 7, 44, 2, 2, 323, 84, 3, 2,
	2, 2, 324, 325, 7, 49, 2, 2, 325, 86, 3, 2, 2, 2, 326, 327, 7, 39, 2, 2,
	327, 88, 3, 2, 2, 2, 328, 329, 7, 62, 2, 2, 329, 330, 7, 62, 2, 2, 330,
	90, 3, 2, 2, 2, 331, 332, 7, 64, 2, 2, 332, 333, 7, 64, 2, 2, 333, 92,
	3, 2, 2, 2, 334, 335, 7, 40, 2, 2, 335, 94, 3, 2, 2, 2, 336, 337, 7, 126,
	2, 2, 337, 96, 3, 2, 2, 2, 338, 339, 7, 96, 2, 2, 339, 98, 3, 2, 2, 2,
	340, 341, 7, 63, 2, 2, 341, 100, 3, 2, 2, 2, 342, 343, 7, 45, 2, 2, 343,
	344, 7, 63, 2, 2, 344, 102, 3, 2, 2, 2, 345, 346, 7, 47, 2, 2, 346, 347,
	7, 63, 2, 2, 347, 104, 3, 2, 2, 2, 348, 349, 7, 44, 2, 2, 349, 350, 7,
	63, 2, 2, 350, 106, 3, 2, 2, 2, 351, 352, 7, 49, 2, 2, 352, 353, 7, 63,
	2, 2, 353, 108, 3, 2, 2, 2, 354, 355, 7, 39, 2, 2, 355, 356, 7, 63, 2,
	2, 356, 110, 3, 2, 2, 2, 357, 358, 7, 62, 2, 2, 358, 359, 7, 62, 2, 2,
	359, 360, 7, 63, 2, 2, 360, 112, 3, 2, 2, 2, 361, 362, 7, 64, 2, 2, 362,
	363, 7, 64, 2, 2, 363, 364, 7, 63, 2, 2, 364, 114, 3, 2, 2, 2, 365, 366,
	7, 40, 2, 2, 366, 367, 7, 63, 2, 2, 367, 116, 3, 2, 2, 2, 368, 369, 7,
	126, 2, 2, 369, 370, 7, 63, 2, 2, 370, 118, 3, 2, 2, 2, 371, 372, 7, 96,
	2, 2, 372, 373, 7, 63, 2, 2, 373, 120, 3, 2, 2, 2, 374, 375, 7, 66, 2,
	2, 375, 122, 3, 2, 2, 2, 376, 378, 9, 2, 2, 2, 377, 376, 3, 2, 2, 2, 377,
	378, 3, 2, 2, 2, 378, 380, 3, 2, 2, 2, 379, 381, 9, 3, 2, 2, 380, 379,
	3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 382, 383, 3, 2,
	2, 2, 383, 124, 3, 2, 2, 2, 384, 386, 9, 2, 2, 2, 385, 384, 3, 2, 2, 2,
	385, 386, 3, 2, 2, 2, 386, 388, 3, 2, 2, 2, 387, 389, 9, 3, 2, 2, 388,
	387, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 390, 391,
	3, 2, 2, 2, 391, 398, 3, 2, 2, 2, 392, 394, 9, 4, 2, 2, 393, 395, 9, 3,
	2, 2, 394, 393, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 394, 3, 2, 2, 2,
	396, 397, 3, 2, 2, 2, 397, 399, 3, 2, 2, 2, 398, 392, 3, 2, 2, 2, 398,
	399, 3, 2, 2, 2, 399, 126, 3, 2, 2, 2, 400, 401, 7, 118, 2, 2, 401, 402,
	7, 116, 2, 2, 402, 403, 7, 119, 2, 2, 403, 404, 7, 103, 2, 2, 404, 128,
	3, 2, 2, 2, 405, 406, 7, 104, 2, 2, 406, 407, 7, 99, 2, 2, 407, 408, 7,
	110, 2, 2, 408, 409, 7, 117, 2, 2, 409, 410, 7, 103, 2, 2, 410, 130, 3,
	2, 2, 2, 411, 418, 9, 5, 2, 2, 412, 414, 9, 6, 2, 2, 413, 412, 3, 2, 2,
	2, 413, 414, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 417, 9, 7, 2, 2, 416,
	413, 3, 2, 2, 2, 417, 420, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 418, 419,
	3, 2, 2, 2, 419, 132, 3, 2, 2, 2, 420, 418, 3, 2, 2, 2, 421, 428, 9, 8,
	2, 2, 422, 424, 9, 6, 2, 2, 423, 422, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2,
	424, 425, 3, 2, 2, 2, 425, 427, 9, 9, 2, 2, 426, 423, 3, 2, 2, 2, 427,
	430, 3, 2, 2, 2, 428, 426, 3, 2, 2, 2, 428, 429, 3, 2, 2, 2, 429, 431,
	3, 2, 2, 2, 430, 428, 3, 2, 2, 2, 431, 432, 9, 6, 2, 2, 432, 433, 9, 10,
	2, 2, 433, 134, 3, 2, 2, 2, 434, 441, 9, 8, 2, 2, 435, 437, 9, 6, 2, 2,
	436, 435, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437, 438, 3, 2, 2, 2, 438,
	440, 9, 9, 2, 2, 439, 436, 3, 2, 2, 2, 440, 443, 3, 2, 2, 2, 441, 439,
	3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 136, 3, 2, 2, 2, 443, 441, 3, 2,
	2, 2, 444, 446, 9, 11, 2, 2, 445, 444, 3, 2, 2, 2, 446, 447, 3, 2, 2, 2,
	447, 445, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 449, 3, 2, 2, 2, 449,
	450, 8, 69, 2, 2, 450, 138, 3, 2, 2, 2, 16, 2, 377, 382, 385, 390, 396,
	398, 413, 418, 423, 428, 436, 441, 447, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'pkg'", "'type'", "'for'", "'while'", "'loop'", "'if'", "'elif'",
	"'else'", "'proc'", "'func'", "'const'", "'let'", "'mut'", "'struct'",
	"'enum'", "'union'", "'return'", "'i8_t'", "'u8_t'", "'i16_t'", "'u16_t'",
	"'i32_t'", "'u32_t'", "'i64_t'", "'u64_t'", "'char_t'", "'bool_t'", "'size_t'",
	"'ssize_t'", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'", "','", "';'",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'<<'", "'>>'", "'&'", "'|'", "'^'",
	"'='", "'+='", "'-='", "'*='", "'/='", "'%='", "'<<='", "'>>='", "'&='",
	"'|='", "'^='", "'@'", "", "", "'true'", "'false'",
}

var lexerSymbolicNames = []string{
	"", "KW_PKG", "KW_TYPE", "KW_FOR", "KW_WHILE", "KW_LOOP", "KW_IF", "KW_ELIF",
	"KW_ELSE", "KW_PROC", "KW_FUNC", "KW_CONST", "KW_LET", "KW_MUT", "KW_STRUCT",
	"KW_ENUM", "KW_UNION", "KW_RETURN", "I8_T", "U8_T", "I16_T", "U16_T", "I32_T",
	"U32_T", "I64_T", "U64_T", "CHAR_T", "BOOL_T", "SIZE_T", "SSIZE_T", "LEFT_PRT",
	"RIGHT_PRT", "LEFT_BRC", "RIGHT_BRC", "LEFT_BRK", "RIGHT_BRK", "COLON",
	"COMMA", "SEMICOLON", "ADD", "SUB", "MUL", "DIV", "MOD", "SHL", "SHR",
	"AND", "OR", "XOR", "ASGN", "ADD_ASGN", "SUB_ASGN", "MUL_ASGN", "DIV_ASGN",
	"MOD_ASGN", "SHL_ASGN", "SHR_ASGN", "AND_ASGN", "OR_ASGN", "XOR_ASGN",
	"REF", "INT_CONST", "FLOAT_CONST", "TRUE", "FALSE", "CONST", "TYPE", "IDENTIFIER",
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"KW_PKG", "KW_TYPE", "KW_FOR", "KW_WHILE", "KW_LOOP", "KW_IF", "KW_ELIF",
	"KW_ELSE", "KW_PROC", "KW_FUNC", "KW_CONST", "KW_LET", "KW_MUT", "KW_STRUCT",
	"KW_ENUM", "KW_UNION", "KW_RETURN", "I8_T", "U8_T", "I16_T", "U16_T", "I32_T",
	"U32_T", "I64_T", "U64_T", "CHAR_T", "BOOL_T", "SIZE_T", "SSIZE_T", "LEFT_PRT",
	"RIGHT_PRT", "LEFT_BRC", "RIGHT_BRC", "LEFT_BRK", "RIGHT_BRK", "COLON",
	"COMMA", "SEMICOLON", "ADD", "SUB", "MUL", "DIV", "MOD", "SHL", "SHR",
	"AND", "OR", "XOR", "ASGN", "ADD_ASGN", "SUB_ASGN", "MUL_ASGN", "DIV_ASGN",
	"MOD_ASGN", "SHL_ASGN", "SHR_ASGN", "AND_ASGN", "OR_ASGN", "XOR_ASGN",
	"REF", "INT_CONST", "FLOAT_CONST", "TRUE", "FALSE", "CONST", "TYPE", "IDENTIFIER",
	"WHITESPACE",
}

type DZLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewDZLexer(input antlr.CharStream) *DZLexer {

	l := new(DZLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "DZ.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DZLexer tokens.
const (
	DZLexerKW_PKG      = 1
	DZLexerKW_TYPE     = 2
	DZLexerKW_FOR      = 3
	DZLexerKW_WHILE    = 4
	DZLexerKW_LOOP     = 5
	DZLexerKW_IF       = 6
	DZLexerKW_ELIF     = 7
	DZLexerKW_ELSE     = 8
	DZLexerKW_PROC     = 9
	DZLexerKW_FUNC     = 10
	DZLexerKW_CONST    = 11
	DZLexerKW_LET      = 12
	DZLexerKW_MUT      = 13
	DZLexerKW_STRUCT   = 14
	DZLexerKW_ENUM     = 15
	DZLexerKW_UNION    = 16
	DZLexerKW_RETURN   = 17
	DZLexerI8_T        = 18
	DZLexerU8_T        = 19
	DZLexerI16_T       = 20
	DZLexerU16_T       = 21
	DZLexerI32_T       = 22
	DZLexerU32_T       = 23
	DZLexerI64_T       = 24
	DZLexerU64_T       = 25
	DZLexerCHAR_T      = 26
	DZLexerBOOL_T      = 27
	DZLexerSIZE_T      = 28
	DZLexerSSIZE_T     = 29
	DZLexerLEFT_PRT    = 30
	DZLexerRIGHT_PRT   = 31
	DZLexerLEFT_BRC    = 32
	DZLexerRIGHT_BRC   = 33
	DZLexerLEFT_BRK    = 34
	DZLexerRIGHT_BRK   = 35
	DZLexerCOLON       = 36
	DZLexerCOMMA       = 37
	DZLexerSEMICOLON   = 38
	DZLexerADD         = 39
	DZLexerSUB         = 40
	DZLexerMUL         = 41
	DZLexerDIV         = 42
	DZLexerMOD         = 43
	DZLexerSHL         = 44
	DZLexerSHR         = 45
	DZLexerAND         = 46
	DZLexerOR          = 47
	DZLexerXOR         = 48
	DZLexerASGN        = 49
	DZLexerADD_ASGN    = 50
	DZLexerSUB_ASGN    = 51
	DZLexerMUL_ASGN    = 52
	DZLexerDIV_ASGN    = 53
	DZLexerMOD_ASGN    = 54
	DZLexerSHL_ASGN    = 55
	DZLexerSHR_ASGN    = 56
	DZLexerAND_ASGN    = 57
	DZLexerOR_ASGN     = 58
	DZLexerXOR_ASGN    = 59
	DZLexerREF         = 60
	DZLexerINT_CONST   = 61
	DZLexerFLOAT_CONST = 62
	DZLexerTRUE        = 63
	DZLexerFALSE       = 64
	DZLexerCONST       = 65
	DZLexerTYPE        = 66
	DZLexerIDENTIFIER  = 67
	DZLexerWHITESPACE  = 68
)
