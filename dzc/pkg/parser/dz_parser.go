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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 72, 329,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 7, 4, 114, 10, 4, 12, 4, 14, 4, 117, 11, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 123, 10, 5, 3, 6, 3, 6, 5, 6, 127, 10, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 7,
	12, 152, 10, 12, 12, 12, 14, 12, 155, 11, 12, 5, 12, 157, 10, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 166, 10, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 7, 16, 175, 10, 16, 12, 16, 14,
	16, 178, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 19, 7, 19, 191, 10, 19, 12, 19, 14, 19, 194, 11, 19, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 7, 22, 205,
	10, 22, 12, 22, 14, 22, 208, 11, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3,
	25, 5, 25, 226, 10, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 250, 10, 32, 3, 33, 3, 33, 3, 34,
	3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	37, 3, 37, 5, 37, 267, 10, 37, 3, 38, 7, 38, 270, 10, 38, 12, 38, 14, 38,
	273, 11, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 281, 10,
	39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 7, 44, 298, 10, 44, 12, 44, 14, 44,
	301, 11, 44, 3, 44, 5, 44, 304, 10, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50,
	3, 50, 5, 50, 322, 10, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 2,
	2, 53, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 2, 5,
	3, 2, 70, 71, 3, 2, 21, 32, 3, 2, 53, 57, 2, 303, 2, 104, 3, 2, 2, 2, 4,
	108, 3, 2, 2, 2, 6, 115, 3, 2, 2, 2, 8, 122, 3, 2, 2, 2, 10, 126, 3, 2,
	2, 2, 12, 128, 3, 2, 2, 2, 14, 132, 3, 2, 2, 2, 16, 136, 3, 2, 2, 2, 18,
	139, 3, 2, 2, 2, 20, 144, 3, 2, 2, 2, 22, 156, 3, 2, 2, 2, 24, 158, 3,
	2, 2, 2, 26, 165, 3, 2, 2, 2, 28, 167, 3, 2, 2, 2, 30, 176, 3, 2, 2, 2,
	32, 179, 3, 2, 2, 2, 34, 183, 3, 2, 2, 2, 36, 192, 3, 2, 2, 2, 38, 195,
	3, 2, 2, 2, 40, 197, 3, 2, 2, 2, 42, 206, 3, 2, 2, 2, 44, 209, 3, 2, 2,
	2, 46, 213, 3, 2, 2, 2, 48, 225, 3, 2, 2, 2, 50, 227, 3, 2, 2, 2, 52, 229,
	3, 2, 2, 2, 54, 231, 3, 2, 2, 2, 56, 233, 3, 2, 2, 2, 58, 235, 3, 2, 2,
	2, 60, 241, 3, 2, 2, 2, 62, 249, 3, 2, 2, 2, 64, 251, 3, 2, 2, 2, 66, 253,
	3, 2, 2, 2, 68, 255, 3, 2, 2, 2, 70, 258, 3, 2, 2, 2, 72, 266, 3, 2, 2,
	2, 74, 271, 3, 2, 2, 2, 76, 280, 3, 2, 2, 2, 78, 282, 3, 2, 2, 2, 80, 287,
	3, 2, 2, 2, 82, 289, 3, 2, 2, 2, 84, 292, 3, 2, 2, 2, 86, 299, 3, 2, 2,
	2, 88, 305, 3, 2, 2, 2, 90, 309, 3, 2, 2, 2, 92, 312, 3, 2, 2, 2, 94, 314,
	3, 2, 2, 2, 96, 317, 3, 2, 2, 2, 98, 321, 3, 2, 2, 2, 100, 323, 3, 2, 2,
	2, 102, 325, 3, 2, 2, 2, 104, 105, 5, 4, 3, 2, 105, 106, 5, 6, 4, 2, 106,
	107, 7, 2, 2, 3, 107, 3, 3, 2, 2, 2, 108, 109, 7, 3, 2, 2, 109, 110, 7,
	67, 2, 2, 110, 111, 7, 42, 2, 2, 111, 5, 3, 2, 2, 2, 112, 114, 5, 8, 5,
	2, 113, 112, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115,
	116, 3, 2, 2, 2, 116, 7, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 123, 5,
	10, 6, 2, 119, 123, 5, 26, 14, 2, 120, 123, 5, 58, 30, 2, 121, 123, 5,
	46, 24, 2, 122, 118, 3, 2, 2, 2, 122, 119, 3, 2, 2, 2, 122, 120, 3, 2,
	2, 2, 122, 121, 3, 2, 2, 2, 123, 9, 3, 2, 2, 2, 124, 127, 5, 12, 7, 2,
	125, 127, 5, 14, 8, 2, 126, 124, 3, 2, 2, 2, 126, 125, 3, 2, 2, 2, 127,
	11, 3, 2, 2, 2, 128, 129, 7, 12, 2, 2, 129, 130, 5, 16, 9, 2, 130, 131,
	5, 60, 31, 2, 131, 13, 3, 2, 2, 2, 132, 133, 7, 13, 2, 2, 133, 134, 5,
	18, 10, 2, 134, 135, 5, 60, 31, 2, 135, 15, 3, 2, 2, 2, 136, 137, 7, 67,
	2, 2, 137, 138, 5, 20, 11, 2, 138, 17, 3, 2, 2, 2, 139, 140, 7, 67, 2,
	2, 140, 141, 5, 20, 11, 2, 141, 142, 7, 40, 2, 2, 142, 143, 5, 62, 32,
	2, 143, 19, 3, 2, 2, 2, 144, 145, 7, 33, 2, 2, 145, 146, 5, 22, 12, 2,
	146, 147, 7, 34, 2, 2, 147, 21, 3, 2, 2, 2, 148, 153, 5, 24, 13, 2, 149,
	150, 7, 41, 2, 2, 150, 152, 5, 24, 13, 2, 151, 149, 3, 2, 2, 2, 152, 155,
	3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 157, 3, 2,
	2, 2, 155, 153, 3, 2, 2, 2, 156, 148, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2,
	157, 23, 3, 2, 2, 2, 158, 159, 7, 67, 2, 2, 159, 160, 7, 40, 2, 2, 160,
	161, 5, 62, 32, 2, 161, 25, 3, 2, 2, 2, 162, 166, 5, 28, 15, 2, 163, 166,
	5, 34, 18, 2, 164, 166, 5, 40, 21, 2, 165, 162, 3, 2, 2, 2, 165, 163, 3,
	2, 2, 2, 165, 164, 3, 2, 2, 2, 166, 27, 3, 2, 2, 2, 167, 168, 7, 17, 2,
	2, 168, 169, 7, 66, 2, 2, 169, 170, 7, 35, 2, 2, 170, 171, 5, 30, 16, 2,
	171, 172, 7, 36, 2, 2, 172, 29, 3, 2, 2, 2, 173, 175, 5, 32, 17, 2, 174,
	173, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176, 177,
	3, 2, 2, 2, 177, 31, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 179, 180, 7, 67,
	2, 2, 180, 181, 7, 40, 2, 2, 181, 182, 5, 62, 32, 2, 182, 33, 3, 2, 2,
	2, 183, 184, 7, 18, 2, 2, 184, 185, 7, 66, 2, 2, 185, 186, 7, 35, 2, 2,
	186, 187, 5, 36, 19, 2, 187, 188, 7, 36, 2, 2, 188, 35, 3, 2, 2, 2, 189,
	191, 5, 38, 20, 2, 190, 189, 3, 2, 2, 2, 191, 194, 3, 2, 2, 2, 192, 190,
	3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 37, 3, 2, 2, 2, 194, 192, 3, 2,
	2, 2, 195, 196, 7, 65, 2, 2, 196, 39, 3, 2, 2, 2, 197, 198, 7, 19, 2, 2,
	198, 199, 7, 66, 2, 2, 199, 200, 7, 35, 2, 2, 200, 201, 5, 42, 22, 2, 201,
	202, 7, 36, 2, 2, 202, 41, 3, 2, 2, 2, 203, 205, 5, 44, 23, 2, 204, 203,
	3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2,
	2, 2, 207, 43, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 210, 7, 67, 2, 2,
	210, 211, 7, 40, 2, 2, 211, 212, 5, 62, 32, 2, 212, 45, 3, 2, 2, 2, 213,
	214, 7, 14, 2, 2, 214, 215, 7, 65, 2, 2, 215, 216, 7, 40, 2, 2, 216, 217,
	5, 64, 33, 2, 217, 218, 7, 53, 2, 2, 218, 219, 5, 48, 25, 2, 219, 220,
	7, 42, 2, 2, 220, 47, 3, 2, 2, 2, 221, 226, 5, 50, 26, 2, 222, 226, 5,
	52, 27, 2, 223, 226, 5, 54, 28, 2, 224, 226, 5, 56, 29, 2, 225, 221, 3,
	2, 2, 2, 225, 222, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 224, 3, 2, 2,
	2, 226, 49, 3, 2, 2, 2, 227, 228, 7, 65, 2, 2, 228, 51, 3, 2, 2, 2, 229,
	230, 7, 68, 2, 2, 230, 53, 3, 2, 2, 2, 231, 232, 7, 69, 2, 2, 232, 55,
	3, 2, 2, 2, 233, 234, 9, 2, 2, 2, 234, 57, 3, 2, 2, 2, 235, 236, 7, 5,
	2, 2, 236, 237, 7, 66, 2, 2, 237, 238, 7, 53, 2, 2, 238, 239, 5, 62, 32,
	2, 239, 240, 7, 42, 2, 2, 240, 59, 3, 2, 2, 2, 241, 242, 7, 35, 2, 2, 242,
	243, 5, 74, 38, 2, 243, 244, 7, 36, 2, 2, 244, 61, 3, 2, 2, 2, 245, 250,
	5, 64, 33, 2, 246, 250, 5, 66, 34, 2, 247, 250, 5, 68, 35, 2, 248, 250,
	5, 70, 36, 2, 249, 245, 3, 2, 2, 2, 249, 246, 3, 2, 2, 2, 249, 247, 3,
	2, 2, 2, 249, 248, 3, 2, 2, 2, 250, 63, 3, 2, 2, 2, 251, 252, 9, 3, 2,
	2, 252, 65, 3, 2, 2, 2, 253, 254, 7, 66, 2, 2, 254, 67, 3, 2, 2, 2, 255,
	256, 7, 64, 2, 2, 256, 257, 7, 66, 2, 2, 257, 69, 3, 2, 2, 2, 258, 259,
	7, 37, 2, 2, 259, 260, 7, 66, 2, 2, 260, 261, 7, 40, 2, 2, 261, 262, 5,
	72, 37, 2, 262, 263, 7, 38, 2, 2, 263, 71, 3, 2, 2, 2, 264, 267, 7, 68,
	2, 2, 265, 267, 7, 65, 2, 2, 266, 264, 3, 2, 2, 2, 266, 265, 3, 2, 2, 2,
	267, 73, 3, 2, 2, 2, 268, 270, 5, 76, 39, 2, 269, 268, 3, 2, 2, 2, 270,
	273, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 75, 3,
	2, 2, 2, 273, 271, 3, 2, 2, 2, 274, 281, 5, 78, 40, 2, 275, 281, 5, 82,
	42, 2, 276, 281, 5, 92, 47, 2, 277, 278, 5, 98, 50, 2, 278, 279, 7, 42,
	2, 2, 279, 281, 3, 2, 2, 2, 280, 274, 3, 2, 2, 2, 280, 275, 3, 2, 2, 2,
	280, 276, 3, 2, 2, 2, 280, 277, 3, 2, 2, 2, 281, 77, 3, 2, 2, 2, 282, 283,
	7, 67, 2, 2, 283, 284, 5, 80, 41, 2, 284, 285, 5, 96, 49, 2, 285, 286,
	7, 42, 2, 2, 286, 79, 3, 2, 2, 2, 287, 288, 9, 4, 2, 2, 288, 81, 3, 2,
	2, 2, 289, 290, 5, 84, 43, 2, 290, 291, 5, 86, 44, 2, 291, 83, 3, 2, 2,
	2, 292, 293, 7, 9, 2, 2, 293, 294, 5, 96, 49, 2, 294, 295, 5, 60, 31, 2,
	295, 85, 3, 2, 2, 2, 296, 298, 5, 88, 45, 2, 297, 296, 3, 2, 2, 2, 298,
	301, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 303,
	3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 302, 304, 5, 90, 46, 2, 303, 302, 3,
	2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 87, 3, 2, 2, 2, 305, 306, 7, 10, 2,
	2, 306, 307, 5, 96, 49, 2, 307, 308, 5, 60, 31, 2, 308, 89, 3, 2, 2, 2,
	309, 310, 7, 11, 2, 2, 310, 311, 5, 60, 31, 2, 311, 91, 3, 2, 2, 2, 312,
	313, 5, 94, 48, 2, 313, 93, 3, 2, 2, 2, 314, 315, 7, 8, 2, 2, 315, 316,
	5, 60, 31, 2, 316, 95, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 97, 3, 2,
	2, 2, 319, 322, 5, 100, 51, 2, 320, 322, 5, 102, 52, 2, 321, 319, 3, 2,
	2, 2, 321, 320, 3, 2, 2, 2, 322, 99, 3, 2, 2, 2, 323, 324, 7, 20, 2, 2,
	324, 101, 3, 2, 2, 2, 325, 326, 7, 20, 2, 2, 326, 327, 5, 96, 49, 2, 327,
	103, 3, 2, 2, 2, 19, 115, 122, 126, 153, 156, 165, 176, 192, 206, 225,
	249, 266, 271, 280, 299, 303, 321,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'pkg'", "'use'", "'type'", "'for'", "'while'", "'loop'", "'if'", "'elif'",
	"'else'", "'proc'", "'func'", "'const'", "'let'", "'mut'", "'struct'",
	"'enum'", "'union'", "'return'", "'i8_t'", "'u8_t'", "'i16_t'", "'u16_t'",
	"'i32_t'", "'u32_t'", "'i64_t'", "'u64_t'", "'char_t'", "'bool_t'", "'size_t'",
	"'ssize_t'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'::'", "':'", "','",
	"';'", "'+'", "'-'", "'*'", "'/'", "'%'", "'<<'", "'>>'", "'&'", "'|'",
	"'^'", "'='", "'+='", "'-='", "'*='", "'/='", "'%='", "'<<='", "'>>='",
	"'&='", "'|='", "'^='", "'@'", "", "", "", "", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "KW_PKG", "KW_USE", "KW_TYPE", "KW_FOR", "KW_WHILE", "KW_LOOP", "KW_IF",
	"KW_ELIF", "KW_ELSE", "KW_PROC", "KW_FUNC", "KW_CONST", "KW_LET", "KW_MUT",
	"KW_STRUCT", "KW_ENUM", "KW_UNION", "KW_RETURN", "I8_T", "U8_T", "I16_T",
	"U16_T", "I32_T", "U32_T", "I64_T", "U64_T", "CHAR_T", "BOOL_T", "SIZE_T",
	"SSIZE_T", "LEFT_PRT", "RIGHT_PRT", "LEFT_BRC", "RIGHT_BRC", "LEFT_BRK",
	"RIGHT_BRK", "DCOLON", "COLON", "COMMA", "SEMICOLON", "ADD", "SUB", "MUL",
	"DIV", "MOD", "SHL", "SHR", "AND", "OR", "XOR", "ASGN", "ADD_ASGN", "SUB_ASGN",
	"MUL_ASGN", "DIV_ASGN", "MOD_ASGN", "SHL_ASGN", "SHR_ASGN", "AND_ASGN",
	"OR_ASGN", "XOR_ASGN", "REF", "CONST", "TYPE", "IDENTIFIER", "INT_CONST",
	"FLOAT_CONST", "TRUE", "FALSE", "WHITESPACE",
}

var ruleNames = []string{
	"start", "pkg", "decls", "decl", "subdecl", "procdecl", "funcdecl", "procheader",
	"funcheader", "args", "argdecls", "argdecl", "complexdecl", "structdecl",
	"structattrs", "structattr", "enumdecl", "enumoptions", "enumoption", "uniondecl",
	"unionattrs", "unionattr", "constdecl", "constasgn", "casgn", "intasgn",
	"floatasgn", "boolconst", "typedecl", "block", "typespec", "basictypespec",
	"simpletypespec", "reftypespec", "arraytypespec", "sizespec", "statements",
	"statement", "assignment", "asgnop", "condition", "ifblock", "elseblocks",
	"elifblock", "elseblock", "loop", "trueloop", "expression", "retstatement",
	"procretstatement", "funcretstatement",
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
	DZParserKW_USE      = 2
	DZParserKW_TYPE     = 3
	DZParserKW_FOR      = 4
	DZParserKW_WHILE    = 5
	DZParserKW_LOOP     = 6
	DZParserKW_IF       = 7
	DZParserKW_ELIF     = 8
	DZParserKW_ELSE     = 9
	DZParserKW_PROC     = 10
	DZParserKW_FUNC     = 11
	DZParserKW_CONST    = 12
	DZParserKW_LET      = 13
	DZParserKW_MUT      = 14
	DZParserKW_STRUCT   = 15
	DZParserKW_ENUM     = 16
	DZParserKW_UNION    = 17
	DZParserKW_RETURN   = 18
	DZParserI8_T        = 19
	DZParserU8_T        = 20
	DZParserI16_T       = 21
	DZParserU16_T       = 22
	DZParserI32_T       = 23
	DZParserU32_T       = 24
	DZParserI64_T       = 25
	DZParserU64_T       = 26
	DZParserCHAR_T      = 27
	DZParserBOOL_T      = 28
	DZParserSIZE_T      = 29
	DZParserSSIZE_T     = 30
	DZParserLEFT_PRT    = 31
	DZParserRIGHT_PRT   = 32
	DZParserLEFT_BRC    = 33
	DZParserRIGHT_BRC   = 34
	DZParserLEFT_BRK    = 35
	DZParserRIGHT_BRK   = 36
	DZParserDCOLON      = 37
	DZParserCOLON       = 38
	DZParserCOMMA       = 39
	DZParserSEMICOLON   = 40
	DZParserADD         = 41
	DZParserSUB         = 42
	DZParserMUL         = 43
	DZParserDIV         = 44
	DZParserMOD         = 45
	DZParserSHL         = 46
	DZParserSHR         = 47
	DZParserAND         = 48
	DZParserOR          = 49
	DZParserXOR         = 50
	DZParserASGN        = 51
	DZParserADD_ASGN    = 52
	DZParserSUB_ASGN    = 53
	DZParserMUL_ASGN    = 54
	DZParserDIV_ASGN    = 55
	DZParserMOD_ASGN    = 56
	DZParserSHL_ASGN    = 57
	DZParserSHR_ASGN    = 58
	DZParserAND_ASGN    = 59
	DZParserOR_ASGN     = 60
	DZParserXOR_ASGN    = 61
	DZParserREF         = 62
	DZParserCONST       = 63
	DZParserTYPE        = 64
	DZParserIDENTIFIER  = 65
	DZParserINT_CONST   = 66
	DZParserFLOAT_CONST = 67
	DZParserTRUE        = 68
	DZParserFALSE       = 69
	DZParserWHITESPACE  = 70
)

// DZParser rules.
const (
	DZParserRULE_start            = 0
	DZParserRULE_pkg              = 1
	DZParserRULE_decls            = 2
	DZParserRULE_decl             = 3
	DZParserRULE_subdecl          = 4
	DZParserRULE_procdecl         = 5
	DZParserRULE_funcdecl         = 6
	DZParserRULE_procheader       = 7
	DZParserRULE_funcheader       = 8
	DZParserRULE_args             = 9
	DZParserRULE_argdecls         = 10
	DZParserRULE_argdecl          = 11
	DZParserRULE_complexdecl      = 12
	DZParserRULE_structdecl       = 13
	DZParserRULE_structattrs      = 14
	DZParserRULE_structattr       = 15
	DZParserRULE_enumdecl         = 16
	DZParserRULE_enumoptions      = 17
	DZParserRULE_enumoption       = 18
	DZParserRULE_uniondecl        = 19
	DZParserRULE_unionattrs       = 20
	DZParserRULE_unionattr        = 21
	DZParserRULE_constdecl        = 22
	DZParserRULE_constasgn        = 23
	DZParserRULE_casgn            = 24
	DZParserRULE_intasgn          = 25
	DZParserRULE_floatasgn        = 26
	DZParserRULE_boolconst        = 27
	DZParserRULE_typedecl         = 28
	DZParserRULE_block            = 29
	DZParserRULE_typespec         = 30
	DZParserRULE_basictypespec    = 31
	DZParserRULE_simpletypespec   = 32
	DZParserRULE_reftypespec      = 33
	DZParserRULE_arraytypespec    = 34
	DZParserRULE_sizespec         = 35
	DZParserRULE_statements       = 36
	DZParserRULE_statement        = 37
	DZParserRULE_assignment       = 38
	DZParserRULE_asgnop           = 39
	DZParserRULE_condition        = 40
	DZParserRULE_ifblock          = 41
	DZParserRULE_elseblocks       = 42
	DZParserRULE_elifblock        = 43
	DZParserRULE_elseblock        = 44
	DZParserRULE_loop             = 45
	DZParserRULE_trueloop         = 46
	DZParserRULE_expression       = 47
	DZParserRULE_retstatement     = 48
	DZParserRULE_procretstatement = 49
	DZParserRULE_funcretstatement = 50
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
		p.SetState(102)
		p.Pkg()
	}
	{
		p.SetState(103)
		p.Decls()
	}
	{
		p.SetState(104)
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
		p.SetState(106)
		p.Match(DZParserKW_PKG)
	}
	{
		p.SetState(107)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*PkgContext).name = _m
	}
	{
		p.SetState(108)
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DZParserKW_TYPE)|(1<<DZParserKW_PROC)|(1<<DZParserKW_FUNC)|(1<<DZParserKW_CONST)|(1<<DZParserKW_STRUCT)|(1<<DZParserKW_ENUM)|(1<<DZParserKW_UNION))) != 0 {
		{
			p.SetState(110)
			p.Decl()
		}

		p.SetState(115)
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

func (s *DeclContext) Subdecl() ISubdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubdeclContext)
}

func (s *DeclContext) Complexdecl() IComplexdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComplexdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComplexdeclContext)
}

func (s *DeclContext) Typedecl() ITypedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedeclContext)
}

func (s *DeclContext) Constdecl() IConstdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstdeclContext)
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

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_PROC, DZParserKW_FUNC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Subdecl()
		}

	case DZParserKW_STRUCT, DZParserKW_ENUM, DZParserKW_UNION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Complexdecl()
		}

	case DZParserKW_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Typedecl()
		}

	case DZParserKW_CONST:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(119)
			p.Constdecl()
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

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_PROC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Procdecl()
		}

	case DZParserKW_FUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
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
		p.SetState(126)
		p.Match(DZParserKW_PROC)
	}
	{
		p.SetState(127)
		p.Procheader()
	}
	{
		p.SetState(128)
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
		p.SetState(130)
		p.Match(DZParserKW_FUNC)
	}
	{
		p.SetState(131)
		p.Funcheader()
	}
	{
		p.SetState(132)
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
		p.SetState(134)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcheaderContext).id = _m
	}
	{
		p.SetState(135)
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

func (s *FuncheaderContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *FuncheaderContext) Typespec() ITypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypespecContext)
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
		p.SetState(137)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*FuncheaderContext).id = _m
	}
	{
		p.SetState(138)
		p.Args()
	}
	{
		p.SetState(139)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(140)
		p.Typespec()
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
		p.SetState(142)
		p.Match(DZParserLEFT_PRT)
	}
	{
		p.SetState(143)
		p.Argdecls()
	}
	{
		p.SetState(144)
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserIDENTIFIER {
		{
			p.SetState(146)
			p.Argdecl()
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserCOMMA {
			{
				p.SetState(147)
				p.Match(DZParserCOMMA)
			}
			{
				p.SetState(148)
				p.Argdecl()
			}

			p.SetState(153)
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
	GetT() ITypespecContext

	// SetT sets the t rule contexts.
	SetT(ITypespecContext)

	// IsArgdeclContext differentiates from other interfaces.
	IsArgdeclContext()
}

type ArgdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	t      ITypespecContext
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

func (s *ArgdeclContext) GetT() ITypespecContext { return s.t }

func (s *ArgdeclContext) SetT(v ITypespecContext) { s.t = v }

func (s *ArgdeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *ArgdeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *ArgdeclContext) Typespec() ITypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypespecContext)
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
		p.SetState(156)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ArgdeclContext).id = _m
	}
	{
		p.SetState(157)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(158)

		var _x = p.Typespec()

		localctx.(*ArgdeclContext).t = _x
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

func (s *ComplexdeclContext) Structdecl() IStructdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructdeclContext)
}

func (s *ComplexdeclContext) Enumdecl() IEnumdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumdeclContext)
}

func (s *ComplexdeclContext) Uniondecl() IUniondeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUniondeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUniondeclContext)
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
	p.EnterRule(localctx, 24, DZParserRULE_complexdecl)

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

	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_STRUCT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Structdecl()
		}

	case DZParserKW_ENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Enumdecl()
		}

	case DZParserKW_UNION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.Uniondecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 26, DZParserRULE_structdecl)

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
		p.SetState(165)
		p.Match(DZParserKW_STRUCT)
	}
	{
		p.SetState(166)

		var _m = p.Match(DZParserTYPE)

		localctx.(*StructdeclContext).id = _m
	}
	{
		p.SetState(167)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(168)
		p.Structattrs()
	}
	{
		p.SetState(169)
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
	p.EnterRule(localctx, 28, DZParserRULE_structattrs)
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
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserIDENTIFIER {
		{
			p.SetState(171)
			p.Structattr()
		}

		p.SetState(176)
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
	p.EnterRule(localctx, 30, DZParserRULE_structattr)

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
		p.SetState(177)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*StructattrContext).id = _m
	}
	{
		p.SetState(178)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(179)

		var _x = p.Typespec()

		localctx.(*StructattrContext).t = _x
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
	p.EnterRule(localctx, 32, DZParserRULE_enumdecl)

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
		p.SetState(181)
		p.Match(DZParserKW_ENUM)
	}
	{
		p.SetState(182)

		var _m = p.Match(DZParserTYPE)

		localctx.(*EnumdeclContext).id = _m
	}
	{
		p.SetState(183)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(184)
		p.Enumoptions()
	}
	{
		p.SetState(185)
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
	p.EnterRule(localctx, 34, DZParserRULE_enumoptions)
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
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserCONST {
		{
			p.SetState(187)
			p.Enumoption()
		}

		p.SetState(192)
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
	p.EnterRule(localctx, 36, DZParserRULE_enumoption)

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
		p.SetState(193)

		var _m = p.Match(DZParserCONST)

		localctx.(*EnumoptionContext).id = _m
	}

	return localctx
}

// IUniondeclContext is an interface to support dynamic dispatch.
type IUniondeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsUniondeclContext differentiates from other interfaces.
	IsUniondeclContext()
}

type UniondeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyUniondeclContext() *UniondeclContext {
	var p = new(UniondeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_uniondecl
	return p
}

func (*UniondeclContext) IsUniondeclContext() {}

func NewUniondeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UniondeclContext {
	var p = new(UniondeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_uniondecl

	return p
}

func (s *UniondeclContext) GetParser() antlr.Parser { return s.parser }

func (s *UniondeclContext) GetId() antlr.Token { return s.id }

func (s *UniondeclContext) SetId(v antlr.Token) { s.id = v }

func (s *UniondeclContext) KW_UNION() antlr.TerminalNode {
	return s.GetToken(DZParserKW_UNION, 0)
}

func (s *UniondeclContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *UniondeclContext) Unionattrs() IUnionattrsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionattrsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionattrsContext)
}

func (s *UniondeclContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *UniondeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
}

func (s *UniondeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UniondeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UniondeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterUniondecl(s)
	}
}

func (s *UniondeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitUniondecl(s)
	}
}

func (p *DZParser) Uniondecl() (localctx IUniondeclContext) {
	localctx = NewUniondeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DZParserRULE_uniondecl)

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
		p.SetState(195)
		p.Match(DZParserKW_UNION)
	}
	{
		p.SetState(196)

		var _m = p.Match(DZParserTYPE)

		localctx.(*UniondeclContext).id = _m
	}
	{
		p.SetState(197)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(198)
		p.Unionattrs()
	}
	{
		p.SetState(199)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// IUnionattrsContext is an interface to support dynamic dispatch.
type IUnionattrsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionattrsContext differentiates from other interfaces.
	IsUnionattrsContext()
}

type UnionattrsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionattrsContext() *UnionattrsContext {
	var p = new(UnionattrsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_unionattrs
	return p
}

func (*UnionattrsContext) IsUnionattrsContext() {}

func NewUnionattrsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionattrsContext {
	var p = new(UnionattrsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_unionattrs

	return p
}

func (s *UnionattrsContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionattrsContext) AllUnionattr() []IUnionattrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnionattrContext)(nil)).Elem())
	var tst = make([]IUnionattrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnionattrContext)
		}
	}

	return tst
}

func (s *UnionattrsContext) Unionattr(i int) IUnionattrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionattrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnionattrContext)
}

func (s *UnionattrsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionattrsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionattrsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterUnionattrs(s)
	}
}

func (s *UnionattrsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitUnionattrs(s)
	}
}

func (p *DZParser) Unionattrs() (localctx IUnionattrsContext) {
	localctx = NewUnionattrsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DZParserRULE_unionattrs)
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
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserIDENTIFIER {
		{
			p.SetState(201)
			p.Unionattr()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnionattrContext is an interface to support dynamic dispatch.
type IUnionattrContext interface {
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

	// IsUnionattrContext differentiates from other interfaces.
	IsUnionattrContext()
}

type UnionattrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	t      ITypespecContext
}

func NewEmptyUnionattrContext() *UnionattrContext {
	var p = new(UnionattrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_unionattr
	return p
}

func (*UnionattrContext) IsUnionattrContext() {}

func NewUnionattrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionattrContext {
	var p = new(UnionattrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_unionattr

	return p
}

func (s *UnionattrContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionattrContext) GetId() antlr.Token { return s.id }

func (s *UnionattrContext) SetId(v antlr.Token) { s.id = v }

func (s *UnionattrContext) GetT() ITypespecContext { return s.t }

func (s *UnionattrContext) SetT(v ITypespecContext) { s.t = v }

func (s *UnionattrContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *UnionattrContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *UnionattrContext) Typespec() ITypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypespecContext)
}

func (s *UnionattrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionattrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionattrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterUnionattr(s)
	}
}

func (s *UnionattrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitUnionattr(s)
	}
}

func (p *DZParser) Unionattr() (localctx IUnionattrContext) {
	localctx = NewUnionattrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DZParserRULE_unionattr)

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
		p.SetState(207)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*UnionattrContext).id = _m
	}
	{
		p.SetState(208)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(209)

		var _x = p.Typespec()

		localctx.(*UnionattrContext).t = _x
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

	// IsConstdeclContext differentiates from other interfaces.
	IsConstdeclContext()
}

type ConstdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
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

func (s *ConstdeclContext) KW_CONST() antlr.TerminalNode {
	return s.GetToken(DZParserKW_CONST, 0)
}

func (s *ConstdeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *ConstdeclContext) Basictypespec() IBasictypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasictypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasictypespecContext)
}

func (s *ConstdeclContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *ConstdeclContext) Constasgn() IConstasgnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstasgnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstasgnContext)
}

func (s *ConstdeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *ConstdeclContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
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
	p.EnterRule(localctx, 44, DZParserRULE_constdecl)

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
		p.SetState(211)
		p.Match(DZParserKW_CONST)
	}
	{
		p.SetState(212)

		var _m = p.Match(DZParserCONST)

		localctx.(*ConstdeclContext).id = _m
	}
	{
		p.SetState(213)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(214)
		p.Basictypespec()
	}
	{
		p.SetState(215)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(216)
		p.Constasgn()
	}
	{
		p.SetState(217)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// IConstasgnContext is an interface to support dynamic dispatch.
type IConstasgnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstasgnContext differentiates from other interfaces.
	IsConstasgnContext()
}

type ConstasgnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstasgnContext() *ConstasgnContext {
	var p = new(ConstasgnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_constasgn
	return p
}

func (*ConstasgnContext) IsConstasgnContext() {}

func NewConstasgnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstasgnContext {
	var p = new(ConstasgnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_constasgn

	return p
}

func (s *ConstasgnContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstasgnContext) Casgn() ICasgnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasgnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICasgnContext)
}

func (s *ConstasgnContext) Intasgn() IIntasgnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntasgnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntasgnContext)
}

func (s *ConstasgnContext) Floatasgn() IFloatasgnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatasgnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatasgnContext)
}

func (s *ConstasgnContext) Boolconst() IBoolconstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolconstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolconstContext)
}

func (s *ConstasgnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstasgnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstasgnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterConstasgn(s)
	}
}

func (s *ConstasgnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitConstasgn(s)
	}
}

func (p *DZParser) Constasgn() (localctx IConstasgnContext) {
	localctx = NewConstasgnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DZParserRULE_constasgn)

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

	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserCONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.Casgn()
		}

	case DZParserINT_CONST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.Intasgn()
		}

	case DZParserFLOAT_CONST:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.Floatasgn()
		}

	case DZParserTRUE, DZParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(222)
			p.Boolconst()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICasgnContext is an interface to support dynamic dispatch.
type ICasgnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsCasgnContext differentiates from other interfaces.
	IsCasgnContext()
}

type CasgnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyCasgnContext() *CasgnContext {
	var p = new(CasgnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_casgn
	return p
}

func (*CasgnContext) IsCasgnContext() {}

func NewCasgnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasgnContext {
	var p = new(CasgnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_casgn

	return p
}

func (s *CasgnContext) GetParser() antlr.Parser { return s.parser }

func (s *CasgnContext) GetId() antlr.Token { return s.id }

func (s *CasgnContext) SetId(v antlr.Token) { s.id = v }

func (s *CasgnContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
}

func (s *CasgnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasgnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasgnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterCasgn(s)
	}
}

func (s *CasgnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitCasgn(s)
	}
}

func (p *DZParser) Casgn() (localctx ICasgnContext) {
	localctx = NewCasgnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DZParserRULE_casgn)

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
		p.SetState(225)

		var _m = p.Match(DZParserCONST)

		localctx.(*CasgnContext).id = _m
	}

	return localctx
}

// IIntasgnContext is an interface to support dynamic dispatch.
type IIntasgnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsIntasgnContext differentiates from other interfaces.
	IsIntasgnContext()
}

type IntasgnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyIntasgnContext() *IntasgnContext {
	var p = new(IntasgnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_intasgn
	return p
}

func (*IntasgnContext) IsIntasgnContext() {}

func NewIntasgnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntasgnContext {
	var p = new(IntasgnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_intasgn

	return p
}

func (s *IntasgnContext) GetParser() antlr.Parser { return s.parser }

func (s *IntasgnContext) GetV() antlr.Token { return s.v }

func (s *IntasgnContext) SetV(v antlr.Token) { s.v = v }

func (s *IntasgnContext) INT_CONST() antlr.TerminalNode {
	return s.GetToken(DZParserINT_CONST, 0)
}

func (s *IntasgnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntasgnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntasgnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterIntasgn(s)
	}
}

func (s *IntasgnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitIntasgn(s)
	}
}

func (p *DZParser) Intasgn() (localctx IIntasgnContext) {
	localctx = NewIntasgnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DZParserRULE_intasgn)

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
		p.SetState(227)

		var _m = p.Match(DZParserINT_CONST)

		localctx.(*IntasgnContext).v = _m
	}

	return localctx
}

// IFloatasgnContext is an interface to support dynamic dispatch.
type IFloatasgnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsFloatasgnContext differentiates from other interfaces.
	IsFloatasgnContext()
}

type FloatasgnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyFloatasgnContext() *FloatasgnContext {
	var p = new(FloatasgnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_floatasgn
	return p
}

func (*FloatasgnContext) IsFloatasgnContext() {}

func NewFloatasgnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatasgnContext {
	var p = new(FloatasgnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_floatasgn

	return p
}

func (s *FloatasgnContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatasgnContext) GetV() antlr.Token { return s.v }

func (s *FloatasgnContext) SetV(v antlr.Token) { s.v = v }

func (s *FloatasgnContext) FLOAT_CONST() antlr.TerminalNode {
	return s.GetToken(DZParserFLOAT_CONST, 0)
}

func (s *FloatasgnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatasgnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatasgnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterFloatasgn(s)
	}
}

func (s *FloatasgnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitFloatasgn(s)
	}
}

func (p *DZParser) Floatasgn() (localctx IFloatasgnContext) {
	localctx = NewFloatasgnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DZParserRULE_floatasgn)

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
		p.SetState(229)

		var _m = p.Match(DZParserFLOAT_CONST)

		localctx.(*FloatasgnContext).v = _m
	}

	return localctx
}

// IBoolconstContext is an interface to support dynamic dispatch.
type IBoolconstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsBoolconstContext differentiates from other interfaces.
	IsBoolconstContext()
}

type BoolconstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyBoolconstContext() *BoolconstContext {
	var p = new(BoolconstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_boolconst
	return p
}

func (*BoolconstContext) IsBoolconstContext() {}

func NewBoolconstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolconstContext {
	var p = new(BoolconstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_boolconst

	return p
}

func (s *BoolconstContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolconstContext) GetV() antlr.Token { return s.v }

func (s *BoolconstContext) SetV(v antlr.Token) { s.v = v }

func (s *BoolconstContext) TRUE() antlr.TerminalNode {
	return s.GetToken(DZParserTRUE, 0)
}

func (s *BoolconstContext) FALSE() antlr.TerminalNode {
	return s.GetToken(DZParserFALSE, 0)
}

func (s *BoolconstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolconstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolconstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterBoolconst(s)
	}
}

func (s *BoolconstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitBoolconst(s)
	}
}

func (p *DZParser) Boolconst() (localctx IBoolconstContext) {
	localctx = NewBoolconstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DZParserRULE_boolconst)
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
		p.SetState(231)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*BoolconstContext).v = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == DZParserTRUE || _la == DZParserFALSE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*BoolconstContext).v = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

	// IsTypedeclContext differentiates from other interfaces.
	IsTypedeclContext()
}

type TypedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
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

func (s *TypedeclContext) KW_TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_TYPE, 0)
}

func (s *TypedeclContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *TypedeclContext) Typespec() ITypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypespecContext)
}

func (s *TypedeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *TypedeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
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
	p.EnterRule(localctx, 56, DZParserRULE_typedecl)

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
		p.Match(DZParserKW_TYPE)
	}
	{
		p.SetState(234)

		var _m = p.Match(DZParserTYPE)

		localctx.(*TypedeclContext).id = _m
	}
	{
		p.SetState(235)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(236)
		p.Typespec()
	}
	{
		p.SetState(237)
		p.Match(DZParserSEMICOLON)
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

func (s *BlockContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *BlockContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
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
	p.EnterRule(localctx, 58, DZParserRULE_block)

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
		p.SetState(239)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(240)
		p.Statements()
	}
	{
		p.SetState(241)
		p.Match(DZParserRIGHT_BRC)
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

func (s *TypespecContext) Basictypespec() IBasictypespecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasictypespecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasictypespecContext)
}

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
	p.EnterRule(localctx, 60, DZParserRULE_typespec)

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

	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserI8_T, DZParserU8_T, DZParserI16_T, DZParserU16_T, DZParserI32_T, DZParserU32_T, DZParserI64_T, DZParserU64_T, DZParserCHAR_T, DZParserBOOL_T, DZParserSIZE_T, DZParserSSIZE_T:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.Basictypespec()
		}

	case DZParserTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.Simpletypespec()
		}

	case DZParserREF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(245)
			p.Reftypespec()
		}

	case DZParserLEFT_BRK:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)
			p.Arraytypespec()
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
	p.EnterRule(localctx, 62, DZParserRULE_basictypespec)
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
		p.SetState(249)

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

// ISimpletypespecContext is an interface to support dynamic dispatch.
type ISimpletypespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsSimpletypespecContext differentiates from other interfaces.
	IsSimpletypespecContext()
}

type SimpletypespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
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

func (s *SimpletypespecContext) GetId() antlr.Token { return s.id }

func (s *SimpletypespecContext) SetId(v antlr.Token) { s.id = v }

func (s *SimpletypespecContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
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
	p.EnterRule(localctx, 64, DZParserRULE_simpletypespec)

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
		p.SetState(251)

		var _m = p.Match(DZParserTYPE)

		localctx.(*SimpletypespecContext).id = _m
	}

	return localctx
}

// IReftypespecContext is an interface to support dynamic dispatch.
type IReftypespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsReftypespecContext differentiates from other interfaces.
	IsReftypespecContext()
}

type ReftypespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
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

func (s *ReftypespecContext) GetId() antlr.Token { return s.id }

func (s *ReftypespecContext) SetId(v antlr.Token) { s.id = v }

func (s *ReftypespecContext) REF() antlr.TerminalNode {
	return s.GetToken(DZParserREF, 0)
}

func (s *ReftypespecContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
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
	p.EnterRule(localctx, 66, DZParserRULE_reftypespec)

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
		p.SetState(253)
		p.Match(DZParserREF)
	}
	{
		p.SetState(254)

		var _m = p.Match(DZParserTYPE)

		localctx.(*ReftypespecContext).id = _m
	}

	return localctx
}

// IArraytypespecContext is an interface to support dynamic dispatch.
type IArraytypespecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetSize returns the size rule contexts.
	GetSize() ISizespecContext

	// SetSize sets the size rule contexts.
	SetSize(ISizespecContext)

	// IsArraytypespecContext differentiates from other interfaces.
	IsArraytypespecContext()
}

type ArraytypespecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
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

func (s *ArraytypespecContext) GetId() antlr.Token { return s.id }

func (s *ArraytypespecContext) SetId(v antlr.Token) { s.id = v }

func (s *ArraytypespecContext) GetSize() ISizespecContext { return s.size }

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

func (s *ArraytypespecContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
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
	p.EnterRule(localctx, 68, DZParserRULE_arraytypespec)

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
		p.SetState(256)
		p.Match(DZParserLEFT_BRK)
	}
	{
		p.SetState(257)

		var _m = p.Match(DZParserTYPE)

		localctx.(*ArraytypespecContext).id = _m
	}
	{
		p.SetState(258)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(259)

		var _x = p.Sizespec()

		localctx.(*ArraytypespecContext).size = _x
	}
	{
		p.SetState(260)
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
	p.EnterRule(localctx, 70, DZParserRULE_sizespec)

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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserINT_CONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Match(DZParserINT_CONST)
		}

	case DZParserCONST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)

			var _m = p.Match(DZParserCONST)

			localctx.(*SizespecContext).name = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *DZParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, DZParserRULE_statements)
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
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DZParserKW_LOOP)|(1<<DZParserKW_IF)|(1<<DZParserKW_RETURN))) != 0) || _la == DZParserIDENTIFIER {
		{
			p.SetState(266)
			p.Statement()
		}

		p.SetState(271)
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

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

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

func (s *StatementContext) Retstatement() IRetstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetstatementContext)
}

func (s *StatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
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
	p.EnterRule(localctx, 74, DZParserRULE_statement)

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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Assignment()
		}

	case DZParserKW_IF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.Condition()
		}

	case DZParserKW_LOOP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(274)
			p.Loop()
		}

	case DZParserKW_RETURN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(275)
			p.Retstatement()
		}
		{
			p.SetState(276)
			p.Match(DZParserSEMICOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) GetId() antlr.Token { return s.id }

func (s *AssignmentContext) SetId(v antlr.Token) { s.id = v }

func (s *AssignmentContext) Asgnop() IAsgnopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsgnopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsgnopContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *AssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *DZParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, DZParserRULE_assignment)

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

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*AssignmentContext).id = _m
	}
	{
		p.SetState(281)
		p.Asgnop()
	}
	{
		p.SetState(282)
		p.Expression()
	}
	{
		p.SetState(283)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// IAsgnopContext is an interface to support dynamic dispatch.
type IAsgnopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsAsgnopContext differentiates from other interfaces.
	IsAsgnopContext()
}

type AsgnopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyAsgnopContext() *AsgnopContext {
	var p = new(AsgnopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_asgnop
	return p
}

func (*AsgnopContext) IsAsgnopContext() {}

func NewAsgnopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsgnopContext {
	var p = new(AsgnopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_asgnop

	return p
}

func (s *AsgnopContext) GetParser() antlr.Parser { return s.parser }

func (s *AsgnopContext) GetId() antlr.Token { return s.id }

func (s *AsgnopContext) SetId(v antlr.Token) { s.id = v }

func (s *AsgnopContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *AsgnopContext) ADD_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserADD_ASGN, 0)
}

func (s *AsgnopContext) SUB_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserSUB_ASGN, 0)
}

func (s *AsgnopContext) MUL_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserMUL_ASGN, 0)
}

func (s *AsgnopContext) DIV_ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserDIV_ASGN, 0)
}

func (s *AsgnopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsgnopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsgnopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterAsgnop(s)
	}
}

func (s *AsgnopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitAsgnop(s)
	}
}

func (p *DZParser) Asgnop() (localctx IAsgnopContext) {
	localctx = NewAsgnopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, DZParserRULE_asgnop)
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
		p.SetState(285)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AsgnopContext).id = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(DZParserASGN-51))|(1<<(DZParserADD_ASGN-51))|(1<<(DZParserSUB_ASGN-51))|(1<<(DZParserMUL_ASGN-51))|(1<<(DZParserDIV_ASGN-51)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AsgnopContext).id = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

func (s *ConditionContext) Ifblock() IIfblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfblockContext)
}

func (s *ConditionContext) Elseblocks() IElseblocksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseblocksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseblocksContext)
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
	p.EnterRule(localctx, 80, DZParserRULE_condition)

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
		p.SetState(287)
		p.Ifblock()
	}
	{
		p.SetState(288)
		p.Elseblocks()
	}

	return localctx
}

// IIfblockContext is an interface to support dynamic dispatch.
type IIfblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfblockContext differentiates from other interfaces.
	IsIfblockContext()
}

type IfblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfblockContext() *IfblockContext {
	var p = new(IfblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_ifblock
	return p
}

func (*IfblockContext) IsIfblockContext() {}

func NewIfblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfblockContext {
	var p = new(IfblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_ifblock

	return p
}

func (s *IfblockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfblockContext) KW_IF() antlr.TerminalNode {
	return s.GetToken(DZParserKW_IF, 0)
}

func (s *IfblockContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfblockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterIfblock(s)
	}
}

func (s *IfblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitIfblock(s)
	}
}

func (p *DZParser) Ifblock() (localctx IIfblockContext) {
	localctx = NewIfblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, DZParserRULE_ifblock)

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
		p.SetState(290)
		p.Match(DZParserKW_IF)
	}
	{
		p.SetState(291)
		p.Expression()
	}
	{
		p.SetState(292)
		p.Block()
	}

	return localctx
}

// IElseblocksContext is an interface to support dynamic dispatch.
type IElseblocksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseblocksContext differentiates from other interfaces.
	IsElseblocksContext()
}

type ElseblocksContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseblocksContext() *ElseblocksContext {
	var p = new(ElseblocksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_elseblocks
	return p
}

func (*ElseblocksContext) IsElseblocksContext() {}

func NewElseblocksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseblocksContext {
	var p = new(ElseblocksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_elseblocks

	return p
}

func (s *ElseblocksContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseblocksContext) AllElifblock() []IElifblockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElifblockContext)(nil)).Elem())
	var tst = make([]IElifblockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElifblockContext)
		}
	}

	return tst
}

func (s *ElseblocksContext) Elifblock(i int) IElifblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElifblockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElifblockContext)
}

func (s *ElseblocksContext) Elseblock() IElseblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseblockContext)
}

func (s *ElseblocksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseblocksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseblocksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterElseblocks(s)
	}
}

func (s *ElseblocksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitElseblocks(s)
	}
}

func (p *DZParser) Elseblocks() (localctx IElseblocksContext) {
	localctx = NewElseblocksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, DZParserRULE_elseblocks)
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
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserKW_ELIF {
		{
			p.SetState(294)
			p.Elifblock()
		}

		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserKW_ELSE {
		{
			p.SetState(300)
			p.Elseblock()
		}

	}

	return localctx
}

// IElifblockContext is an interface to support dynamic dispatch.
type IElifblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElifblockContext differentiates from other interfaces.
	IsElifblockContext()
}

type ElifblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElifblockContext() *ElifblockContext {
	var p = new(ElifblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_elifblock
	return p
}

func (*ElifblockContext) IsElifblockContext() {}

func NewElifblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifblockContext {
	var p = new(ElifblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_elifblock

	return p
}

func (s *ElifblockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifblockContext) KW_ELIF() antlr.TerminalNode {
	return s.GetToken(DZParserKW_ELIF, 0)
}

func (s *ElifblockContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElifblockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElifblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElifblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterElifblock(s)
	}
}

func (s *ElifblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitElifblock(s)
	}
}

func (p *DZParser) Elifblock() (localctx IElifblockContext) {
	localctx = NewElifblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, DZParserRULE_elifblock)

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
		p.SetState(303)
		p.Match(DZParserKW_ELIF)
	}
	{
		p.SetState(304)
		p.Expression()
	}
	{
		p.SetState(305)
		p.Block()
	}

	return localctx
}

// IElseblockContext is an interface to support dynamic dispatch.
type IElseblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseblockContext differentiates from other interfaces.
	IsElseblockContext()
}

type ElseblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseblockContext() *ElseblockContext {
	var p = new(ElseblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_elseblock
	return p
}

func (*ElseblockContext) IsElseblockContext() {}

func NewElseblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseblockContext {
	var p = new(ElseblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_elseblock

	return p
}

func (s *ElseblockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseblockContext) KW_ELSE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_ELSE, 0)
}

func (s *ElseblockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterElseblock(s)
	}
}

func (s *ElseblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitElseblock(s)
	}
}

func (p *DZParser) Elseblock() (localctx IElseblockContext) {
	localctx = NewElseblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, DZParserRULE_elseblock)

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
		p.SetState(307)
		p.Match(DZParserKW_ELSE)
	}
	{
		p.SetState(308)
		p.Block()
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

func (s *LoopContext) Trueloop() ITrueloopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrueloopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrueloopContext)
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
	p.EnterRule(localctx, 90, DZParserRULE_loop)

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
		p.SetState(310)
		p.Trueloop()
	}

	return localctx
}

// ITrueloopContext is an interface to support dynamic dispatch.
type ITrueloopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrueloopContext differentiates from other interfaces.
	IsTrueloopContext()
}

type TrueloopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrueloopContext() *TrueloopContext {
	var p = new(TrueloopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_trueloop
	return p
}

func (*TrueloopContext) IsTrueloopContext() {}

func NewTrueloopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrueloopContext {
	var p = new(TrueloopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_trueloop

	return p
}

func (s *TrueloopContext) GetParser() antlr.Parser { return s.parser }

func (s *TrueloopContext) KW_LOOP() antlr.TerminalNode {
	return s.GetToken(DZParserKW_LOOP, 0)
}

func (s *TrueloopContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TrueloopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueloopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrueloopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterTrueloop(s)
	}
}

func (s *TrueloopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitTrueloop(s)
	}
}

func (p *DZParser) Trueloop() (localctx ITrueloopContext) {
	localctx = NewTrueloopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, DZParserRULE_trueloop)

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
		p.Match(DZParserKW_LOOP)
	}
	{
		p.SetState(313)
		p.Block()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, DZParserRULE_expression)

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

	return localctx
}

// IRetstatementContext is an interface to support dynamic dispatch.
type IRetstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetstatementContext differentiates from other interfaces.
	IsRetstatementContext()
}

type RetstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetstatementContext() *RetstatementContext {
	var p = new(RetstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_retstatement
	return p
}

func (*RetstatementContext) IsRetstatementContext() {}

func NewRetstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetstatementContext {
	var p = new(RetstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_retstatement

	return p
}

func (s *RetstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RetstatementContext) Procretstatement() IProcretstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcretstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcretstatementContext)
}

func (s *RetstatementContext) Funcretstatement() IFuncretstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncretstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncretstatementContext)
}

func (s *RetstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterRetstatement(s)
	}
}

func (s *RetstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitRetstatement(s)
	}
}

func (p *DZParser) Retstatement() (localctx IRetstatementContext) {
	localctx = NewRetstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, DZParserRULE_retstatement)

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

	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.Procretstatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(318)
			p.Funcretstatement()
		}

	}

	return localctx
}

// IProcretstatementContext is an interface to support dynamic dispatch.
type IProcretstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcretstatementContext differentiates from other interfaces.
	IsProcretstatementContext()
}

type ProcretstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcretstatementContext() *ProcretstatementContext {
	var p = new(ProcretstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_procretstatement
	return p
}

func (*ProcretstatementContext) IsProcretstatementContext() {}

func NewProcretstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcretstatementContext {
	var p = new(ProcretstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_procretstatement

	return p
}

func (s *ProcretstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcretstatementContext) KW_RETURN() antlr.TerminalNode {
	return s.GetToken(DZParserKW_RETURN, 0)
}

func (s *ProcretstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcretstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcretstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterProcretstatement(s)
	}
}

func (s *ProcretstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitProcretstatement(s)
	}
}

func (p *DZParser) Procretstatement() (localctx IProcretstatementContext) {
	localctx = NewProcretstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, DZParserRULE_procretstatement)

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
		p.Match(DZParserKW_RETURN)
	}

	return localctx
}

// IFuncretstatementContext is an interface to support dynamic dispatch.
type IFuncretstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v rule contexts.
	GetV() IExpressionContext

	// SetV sets the v rule contexts.
	SetV(IExpressionContext)

	// IsFuncretstatementContext differentiates from other interfaces.
	IsFuncretstatementContext()
}

type FuncretstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      IExpressionContext
}

func NewEmptyFuncretstatementContext() *FuncretstatementContext {
	var p = new(FuncretstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_funcretstatement
	return p
}

func (*FuncretstatementContext) IsFuncretstatementContext() {}

func NewFuncretstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncretstatementContext {
	var p = new(FuncretstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_funcretstatement

	return p
}

func (s *FuncretstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncretstatementContext) GetV() IExpressionContext { return s.v }

func (s *FuncretstatementContext) SetV(v IExpressionContext) { s.v = v }

func (s *FuncretstatementContext) KW_RETURN() antlr.TerminalNode {
	return s.GetToken(DZParserKW_RETURN, 0)
}

func (s *FuncretstatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FuncretstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncretstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncretstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterFuncretstatement(s)
	}
}

func (s *FuncretstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitFuncretstatement(s)
	}
}

func (p *DZParser) Funcretstatement() (localctx IFuncretstatementContext) {
	localctx = NewFuncretstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, DZParserRULE_funcretstatement)

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
		p.SetState(323)
		p.Match(DZParserKW_RETURN)
	}
	{
		p.SetState(324)

		var _x = p.Expression()

		localctx.(*FuncretstatementContext).v = _x
	}

	return localctx
}
