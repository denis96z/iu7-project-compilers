// Code generated from /home/denis/Documents/iu7-project-compilers/stc/SmallTalk.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SmallTalk

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 317,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 7, 2, 98,
	10, 2, 12, 2, 14, 2, 101, 11, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 110, 10, 3, 3, 3, 3, 3, 5, 3, 114, 10, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 5, 6, 5, 121, 10, 5, 13, 5, 14, 5, 122, 3, 6, 3, 6, 3, 6, 5, 6, 128,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 6, 10, 149, 10, 10, 13,
	10, 14, 10, 150, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 6, 12,
	160, 10, 12, 13, 12, 14, 12, 161, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 5,
	14, 169, 10, 14, 3, 14, 5, 14, 172, 10, 14, 3, 15, 3, 15, 3, 15, 7, 15,
	177, 10, 15, 12, 15, 14, 15, 180, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16, 5,
	16, 186, 10, 16, 3, 17, 3, 17, 7, 17, 190, 10, 17, 12, 17, 14, 17, 193,
	11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 5, 20, 205, 10, 20, 3, 20, 3, 20, 7, 20, 209, 10, 20, 12, 20, 14,
	20, 212, 11, 20, 3, 21, 3, 21, 3, 21, 5, 21, 217, 10, 21, 3, 22, 3, 22,
	3, 22, 7, 22, 222, 10, 22, 12, 22, 14, 22, 225, 11, 22, 3, 23, 3, 23, 3,
	24, 3, 24, 5, 24, 231, 10, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26,
	6, 26, 239, 10, 26, 13, 26, 14, 26, 240, 3, 27, 3, 27, 5, 27, 245, 10,
	27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29,
	256, 10, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 5,
	33, 266, 10, 33, 3, 33, 3, 33, 3, 33, 3, 34, 6, 34, 272, 10, 34, 13, 34,
	14, 34, 273, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 5, 36, 286, 10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37,
	293, 10, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3,
	42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47,
	3, 47, 3, 48, 3, 48, 3, 48, 2, 3, 38, 49, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90,
	92, 94, 2, 4, 3, 2, 8, 9, 3, 2, 3, 6, 2, 304, 2, 99, 3, 2, 2, 2, 4, 105,
	3, 2, 2, 2, 6, 117, 3, 2, 2, 2, 8, 120, 3, 2, 2, 2, 10, 127, 3, 2, 2, 2,
	12, 129, 3, 2, 2, 2, 14, 134, 3, 2, 2, 2, 16, 141, 3, 2, 2, 2, 18, 148,
	3, 2, 2, 2, 20, 154, 3, 2, 2, 2, 22, 157, 3, 2, 2, 2, 24, 165, 3, 2, 2,
	2, 26, 168, 3, 2, 2, 2, 28, 173, 3, 2, 2, 2, 30, 185, 3, 2, 2, 2, 32, 187,
	3, 2, 2, 2, 34, 196, 3, 2, 2, 2, 36, 199, 3, 2, 2, 2, 38, 204, 3, 2, 2,
	2, 40, 216, 3, 2, 2, 2, 42, 218, 3, 2, 2, 2, 44, 226, 3, 2, 2, 2, 46, 230,
	3, 2, 2, 2, 48, 234, 3, 2, 2, 2, 50, 238, 3, 2, 2, 2, 52, 244, 3, 2, 2,
	2, 54, 246, 3, 2, 2, 2, 56, 255, 3, 2, 2, 2, 58, 257, 3, 2, 2, 2, 60, 259,
	3, 2, 2, 2, 62, 261, 3, 2, 2, 2, 64, 263, 3, 2, 2, 2, 66, 271, 3, 2, 2,
	2, 68, 277, 3, 2, 2, 2, 70, 285, 3, 2, 2, 2, 72, 292, 3, 2, 2, 2, 74, 294,
	3, 2, 2, 2, 76, 296, 3, 2, 2, 2, 78, 298, 3, 2, 2, 2, 80, 300, 3, 2, 2,
	2, 82, 302, 3, 2, 2, 2, 84, 304, 3, 2, 2, 2, 86, 306, 3, 2, 2, 2, 88, 308,
	3, 2, 2, 2, 90, 310, 3, 2, 2, 2, 92, 312, 3, 2, 2, 2, 94, 314, 3, 2, 2,
	2, 96, 98, 5, 4, 3, 2, 97, 96, 3, 2, 2, 2, 98, 101, 3, 2, 2, 2, 99, 97,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 3, 2, 2, 2, 101, 99, 3, 2, 2,
	2, 102, 103, 5, 6, 4, 2, 103, 104, 7, 2, 2, 3, 104, 3, 3, 2, 2, 2, 105,
	106, 7, 10, 2, 2, 106, 109, 7, 26, 2, 2, 107, 108, 7, 17, 2, 2, 108, 110,
	7, 26, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 3, 2,
	2, 2, 111, 113, 7, 20, 2, 2, 112, 114, 5, 8, 5, 2, 113, 112, 3, 2, 2, 2,
	113, 114, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 7, 21, 2, 2, 116,
	5, 3, 2, 2, 2, 117, 118, 5, 26, 14, 2, 118, 7, 3, 2, 2, 2, 119, 121, 5,
	10, 6, 2, 120, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 120, 3, 2, 2,
	2, 122, 123, 3, 2, 2, 2, 123, 9, 3, 2, 2, 2, 124, 128, 5, 12, 7, 2, 125,
	128, 5, 14, 8, 2, 126, 128, 5, 16, 9, 2, 127, 124, 3, 2, 2, 2, 127, 125,
	3, 2, 2, 2, 127, 126, 3, 2, 2, 2, 128, 11, 3, 2, 2, 2, 129, 130, 7, 26,
	2, 2, 130, 131, 7, 20, 2, 2, 131, 132, 5, 26, 14, 2, 132, 133, 7, 21, 2,
	2, 133, 13, 3, 2, 2, 2, 134, 135, 7, 26, 2, 2, 135, 136, 7, 20, 2, 2, 136,
	137, 5, 20, 11, 2, 137, 138, 7, 16, 2, 2, 138, 139, 5, 26, 14, 2, 139,
	140, 7, 21, 2, 2, 140, 15, 3, 2, 2, 2, 141, 142, 7, 27, 2, 2, 142, 143,
	7, 20, 2, 2, 143, 144, 5, 18, 10, 2, 144, 145, 5, 26, 14, 2, 145, 146,
	7, 21, 2, 2, 146, 17, 3, 2, 2, 2, 147, 149, 5, 20, 11, 2, 148, 147, 3,
	2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2,
	2, 151, 152, 3, 2, 2, 2, 152, 153, 7, 16, 2, 2, 153, 19, 3, 2, 2, 2, 154,
	155, 7, 17, 2, 2, 155, 156, 7, 26, 2, 2, 156, 21, 3, 2, 2, 2, 157, 159,
	7, 16, 2, 2, 158, 160, 5, 24, 13, 2, 159, 158, 3, 2, 2, 2, 160, 161, 3,
	2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 3, 2, 2,
	2, 163, 164, 7, 16, 2, 2, 164, 23, 3, 2, 2, 2, 165, 166, 7, 26, 2, 2, 166,
	25, 3, 2, 2, 2, 167, 169, 5, 22, 12, 2, 168, 167, 3, 2, 2, 2, 168, 169,
	3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 172, 5, 28, 15, 2, 171, 170, 3,
	2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 27, 3, 2, 2, 2, 173, 178, 5, 30, 16,
	2, 174, 175, 7, 15, 2, 2, 175, 177, 5, 30, 16, 2, 176, 174, 3, 2, 2, 2,
	177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179,
	181, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 181, 182, 7, 15, 2, 2, 182, 29,
	3, 2, 2, 2, 183, 186, 5, 36, 19, 2, 184, 186, 5, 32, 17, 2, 185, 183, 3,
	2, 2, 2, 185, 184, 3, 2, 2, 2, 186, 31, 3, 2, 2, 2, 187, 191, 5, 34, 18,
	2, 188, 190, 5, 34, 18, 2, 189, 188, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2,
	191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 194, 3, 2, 2, 2, 193,
	191, 3, 2, 2, 2, 194, 195, 5, 38, 20, 2, 195, 33, 3, 2, 2, 2, 196, 197,
	7, 26, 2, 2, 197, 198, 7, 19, 2, 2, 198, 35, 3, 2, 2, 2, 199, 200, 5, 38,
	20, 2, 200, 37, 3, 2, 2, 2, 201, 202, 8, 20, 1, 2, 202, 205, 5, 54, 28,
	2, 203, 205, 5, 56, 29, 2, 204, 201, 3, 2, 2, 2, 204, 203, 3, 2, 2, 2,
	205, 210, 3, 2, 2, 2, 206, 207, 12, 5, 2, 2, 207, 209, 5, 42, 22, 2, 208,
	206, 3, 2, 2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211,
	3, 2, 2, 2, 211, 39, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 217, 5, 48,
	25, 2, 214, 217, 5, 46, 24, 2, 215, 217, 5, 44, 23, 2, 216, 213, 3, 2,
	2, 2, 216, 214, 3, 2, 2, 2, 216, 215, 3, 2, 2, 2, 217, 41, 3, 2, 2, 2,
	218, 223, 5, 40, 21, 2, 219, 220, 7, 18, 2, 2, 220, 222, 5, 40, 21, 2,
	221, 219, 3, 2, 2, 2, 222, 225, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223,
	224, 3, 2, 2, 2, 224, 43, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 226, 227, 7,
	26, 2, 2, 227, 45, 3, 2, 2, 2, 228, 231, 5, 92, 47, 2, 229, 231, 5, 94,
	48, 2, 230, 228, 3, 2, 2, 2, 230, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2,
	232, 233, 5, 52, 27, 2, 233, 47, 3, 2, 2, 2, 234, 235, 7, 27, 2, 2, 235,
	236, 5, 50, 26, 2, 236, 49, 3, 2, 2, 2, 237, 239, 5, 52, 27, 2, 238, 237,
	3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2,
	2, 2, 241, 51, 3, 2, 2, 2, 242, 245, 5, 56, 29, 2, 243, 245, 5, 54, 28,
	2, 244, 242, 3, 2, 2, 2, 244, 243, 3, 2, 2, 2, 245, 53, 3, 2, 2, 2, 246,
	247, 7, 22, 2, 2, 247, 248, 5, 30, 16, 2, 248, 249, 7, 23, 2, 2, 249, 55,
	3, 2, 2, 2, 250, 256, 5, 58, 30, 2, 251, 256, 5, 60, 31, 2, 252, 256, 5,
	62, 32, 2, 253, 256, 5, 64, 33, 2, 254, 256, 5, 70, 36, 2, 255, 250, 3,
	2, 2, 2, 255, 251, 3, 2, 2, 2, 255, 252, 3, 2, 2, 2, 255, 253, 3, 2, 2,
	2, 255, 254, 3, 2, 2, 2, 256, 57, 3, 2, 2, 2, 257, 258, 7, 26, 2, 2, 258,
	59, 3, 2, 2, 2, 259, 260, 7, 11, 2, 2, 260, 61, 3, 2, 2, 2, 261, 262, 7,
	12, 2, 2, 262, 63, 3, 2, 2, 2, 263, 265, 7, 20, 2, 2, 264, 266, 5, 66,
	34, 2, 265, 264, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2,
	267, 268, 5, 26, 14, 2, 268, 269, 7, 21, 2, 2, 269, 65, 3, 2, 2, 2, 270,
	272, 5, 68, 35, 2, 271, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 271,
	3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 276, 7, 16,
	2, 2, 276, 67, 3, 2, 2, 2, 277, 278, 7, 17, 2, 2, 278, 279, 7, 26, 2, 2,
	279, 69, 3, 2, 2, 2, 280, 286, 5, 72, 37, 2, 281, 286, 5, 84, 43, 2, 282,
	286, 5, 86, 44, 2, 283, 286, 5, 88, 45, 2, 284, 286, 5, 90, 46, 2, 285,
	280, 3, 2, 2, 2, 285, 281, 3, 2, 2, 2, 285, 282, 3, 2, 2, 2, 285, 283,
	3, 2, 2, 2, 285, 284, 3, 2, 2, 2, 286, 71, 3, 2, 2, 2, 287, 293, 5, 74,
	38, 2, 288, 293, 5, 76, 39, 2, 289, 293, 5, 78, 40, 2, 290, 293, 5, 80,
	41, 2, 291, 293, 5, 82, 42, 2, 292, 287, 3, 2, 2, 2, 292, 288, 3, 2, 2,
	2, 292, 289, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292, 291, 3, 2, 2, 2, 293,
	73, 3, 2, 2, 2, 294, 295, 7, 31, 2, 2, 295, 75, 3, 2, 2, 2, 296, 297, 7,
	32, 2, 2, 297, 77, 3, 2, 2, 2, 298, 299, 7, 33, 2, 2, 299, 79, 3, 2, 2,
	2, 300, 301, 7, 34, 2, 2, 301, 81, 3, 2, 2, 2, 302, 303, 7, 35, 2, 2, 303,
	83, 3, 2, 2, 2, 304, 305, 7, 36, 2, 2, 305, 85, 3, 2, 2, 2, 306, 307, 7,
	37, 2, 2, 307, 87, 3, 2, 2, 2, 308, 309, 9, 2, 2, 2, 309, 89, 3, 2, 2,
	2, 310, 311, 7, 7, 2, 2, 311, 91, 3, 2, 2, 2, 312, 313, 7, 26, 2, 2, 313,
	93, 3, 2, 2, 2, 314, 315, 9, 3, 2, 2, 315, 95, 3, 2, 2, 2, 26, 99, 109,
	113, 122, 127, 150, 161, 168, 171, 178, 185, 191, 204, 210, 216, 223, 230,
	240, 244, 255, 265, 273, 285, 292,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'nil'", "'true'", "'false'", "'class'",
	"'self'", "'super'", "", "", "'.'", "'|'", "':'", "';'", "':='", "'['",
	"']'", "'('", "')'", "'{'", "'}'", "", "", "'2r'", "'8r'", "'16r'",
}
var symbolicNames = []string{
	"", "", "", "", "", "NIL", "TRUE", "FALSE", "CLASS", "SELF", "SUPER", "COMMENT",
	"WHITESPACE", "DOT", "PIPE", "COLON", "SEMICOLON", "ASSIGNMENT", "LEFT_BRK",
	"RIGHT_BRK", "LEFT_PRT", "RIGHT_PRT", "LEFT_BRC", "RIGHT_BRC", "IDENTIFIER",
	"KEYWORD", "BIN", "OCT", "HEX", "BIN_INT_NUMBER", "OCT_INT_NUMBER", "DEC_INT_NUMBER",
	"HEX_INT_NUMBER", "FLOAT_NUMBER", "CHAR", "STRING",
}

var ruleNames = []string{
	"script", "classDef", "main", "instanceMethods", "method", "unaryMethod",
	"binaryMethod", "keywordMethod", "methodArgs", "methodArg", "localVars",
	"localVar", "body", "statements", "statement", "assignmentStatement", "assignmentItem",
	"messageStatement", "messageExpression", "message", "messageRange", "unaryMessage",
	"binaryMessage", "keywordMessage", "keywordMessageArgs", "messageArg",
	"prtStatement", "value", "valueVar", "valueSelf", "valueSuper", "valueBlock",
	"blockArgs", "blockArg", "valueConst", "valueConstNum", "valueConstBinInt",
	"valueConstOctInt", "valueConstDecInt", "valueConstHexInt", "valueConstFloat",
	"valueConstChar", "valueConstString", "valueConstBool", "valueNil", "identifier",
	"operator",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SmallTalkParser struct {
	*antlr.BaseParser
}

func NewSmallTalkParser(input antlr.TokenStream) *SmallTalkParser {
	this := new(SmallTalkParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SmallTalk.g4"

	return this
}

// SmallTalkParser tokens.
const (
	SmallTalkParserEOF            = antlr.TokenEOF
	SmallTalkParserT__0           = 1
	SmallTalkParserT__1           = 2
	SmallTalkParserT__2           = 3
	SmallTalkParserT__3           = 4
	SmallTalkParserNIL            = 5
	SmallTalkParserTRUE           = 6
	SmallTalkParserFALSE          = 7
	SmallTalkParserCLASS          = 8
	SmallTalkParserSELF           = 9
	SmallTalkParserSUPER          = 10
	SmallTalkParserCOMMENT        = 11
	SmallTalkParserWHITESPACE     = 12
	SmallTalkParserDOT            = 13
	SmallTalkParserPIPE           = 14
	SmallTalkParserCOLON          = 15
	SmallTalkParserSEMICOLON      = 16
	SmallTalkParserASSIGNMENT     = 17
	SmallTalkParserLEFT_BRK       = 18
	SmallTalkParserRIGHT_BRK      = 19
	SmallTalkParserLEFT_PRT       = 20
	SmallTalkParserRIGHT_PRT      = 21
	SmallTalkParserLEFT_BRC       = 22
	SmallTalkParserRIGHT_BRC      = 23
	SmallTalkParserIDENTIFIER     = 24
	SmallTalkParserKEYWORD        = 25
	SmallTalkParserBIN            = 26
	SmallTalkParserOCT            = 27
	SmallTalkParserHEX            = 28
	SmallTalkParserBIN_INT_NUMBER = 29
	SmallTalkParserOCT_INT_NUMBER = 30
	SmallTalkParserDEC_INT_NUMBER = 31
	SmallTalkParserHEX_INT_NUMBER = 32
	SmallTalkParserFLOAT_NUMBER   = 33
	SmallTalkParserCHAR           = 34
	SmallTalkParserSTRING         = 35
)

// SmallTalkParser rules.
const (
	SmallTalkParserRULE_script              = 0
	SmallTalkParserRULE_classDef            = 1
	SmallTalkParserRULE_main                = 2
	SmallTalkParserRULE_instanceMethods     = 3
	SmallTalkParserRULE_method              = 4
	SmallTalkParserRULE_unaryMethod         = 5
	SmallTalkParserRULE_binaryMethod        = 6
	SmallTalkParserRULE_keywordMethod       = 7
	SmallTalkParserRULE_methodArgs          = 8
	SmallTalkParserRULE_methodArg           = 9
	SmallTalkParserRULE_localVars           = 10
	SmallTalkParserRULE_localVar            = 11
	SmallTalkParserRULE_body                = 12
	SmallTalkParserRULE_statements          = 13
	SmallTalkParserRULE_statement           = 14
	SmallTalkParserRULE_assignmentStatement = 15
	SmallTalkParserRULE_assignmentItem      = 16
	SmallTalkParserRULE_messageStatement    = 17
	SmallTalkParserRULE_messageExpression   = 18
	SmallTalkParserRULE_message             = 19
	SmallTalkParserRULE_messageRange        = 20
	SmallTalkParserRULE_unaryMessage        = 21
	SmallTalkParserRULE_binaryMessage       = 22
	SmallTalkParserRULE_keywordMessage      = 23
	SmallTalkParserRULE_keywordMessageArgs  = 24
	SmallTalkParserRULE_messageArg          = 25
	SmallTalkParserRULE_prtStatement        = 26
	SmallTalkParserRULE_value               = 27
	SmallTalkParserRULE_valueVar            = 28
	SmallTalkParserRULE_valueSelf           = 29
	SmallTalkParserRULE_valueSuper          = 30
	SmallTalkParserRULE_valueBlock          = 31
	SmallTalkParserRULE_blockArgs           = 32
	SmallTalkParserRULE_blockArg            = 33
	SmallTalkParserRULE_valueConst          = 34
	SmallTalkParserRULE_valueConstNum       = 35
	SmallTalkParserRULE_valueConstBinInt    = 36
	SmallTalkParserRULE_valueConstOctInt    = 37
	SmallTalkParserRULE_valueConstDecInt    = 38
	SmallTalkParserRULE_valueConstHexInt    = 39
	SmallTalkParserRULE_valueConstFloat     = 40
	SmallTalkParserRULE_valueConstChar      = 41
	SmallTalkParserRULE_valueConstString    = 42
	SmallTalkParserRULE_valueConstBool      = 43
	SmallTalkParserRULE_valueNil            = 44
	SmallTalkParserRULE_identifier          = 45
	SmallTalkParserRULE_operator            = 46
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) Main() IMainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMainContext)
}

func (s *ScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserEOF, 0)
}

func (s *ScriptContext) AllClassDef() []IClassDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassDefContext)(nil)).Elem())
	var tst = make([]IClassDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassDefContext)
		}
	}

	return tst
}

func (s *ScriptContext) ClassDef(i int) IClassDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassDefContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *SmallTalkParser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SmallTalkParserRULE_script)
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
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SmallTalkParserCLASS {
		{
			p.SetState(94)
			p.ClassDef()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
		p.Main()
	}
	{
		p.SetState(101)
		p.Match(SmallTalkParserEOF)
	}

	return localctx
}

// IClassDefContext is an interface to support dynamic dispatch.
type IClassDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetSup returns the sup token.
	GetSup() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetSup sets the sup token.
	SetSup(antlr.Token)

	// IsClassDefContext differentiates from other interfaces.
	IsClassDefContext()
}

type ClassDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	sup    antlr.Token
}

func NewEmptyClassDefContext() *ClassDefContext {
	var p = new(ClassDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_classDef
	return p
}

func (*ClassDefContext) IsClassDefContext() {}

func NewClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefContext {
	var p = new(ClassDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_classDef

	return p
}

func (s *ClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefContext) GetName() antlr.Token { return s.name }

func (s *ClassDefContext) GetSup() antlr.Token { return s.sup }

func (s *ClassDefContext) SetName(v antlr.Token) { s.name = v }

func (s *ClassDefContext) SetSup(v antlr.Token) { s.sup = v }

func (s *ClassDefContext) CLASS() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserCLASS, 0)
}

func (s *ClassDefContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserLEFT_BRK, 0)
}

func (s *ClassDefContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserRIGHT_BRK, 0)
}

func (s *ClassDefContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(SmallTalkParserIDENTIFIER)
}

func (s *ClassDefContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, i)
}

func (s *ClassDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserCOLON, 0)
}

func (s *ClassDefContext) InstanceMethods() IInstanceMethodsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstanceMethodsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstanceMethodsContext)
}

func (s *ClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterClassDef(s)
	}
}

func (s *ClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitClassDef(s)
	}
}

func (p *SmallTalkParser) ClassDef() (localctx IClassDefContext) {
	localctx = NewClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SmallTalkParserRULE_classDef)
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
		p.SetState(103)
		p.Match(SmallTalkParserCLASS)
	}
	{
		p.SetState(104)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*ClassDefContext).name = _m
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SmallTalkParserCOLON {
		{
			p.SetState(105)
			p.Match(SmallTalkParserCOLON)
		}
		{
			p.SetState(106)

			var _m = p.Match(SmallTalkParserIDENTIFIER)

			localctx.(*ClassDefContext).sup = _m
		}

	}
	{
		p.SetState(109)
		p.Match(SmallTalkParserLEFT_BRK)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SmallTalkParserIDENTIFIER || _la == SmallTalkParserKEYWORD {
		{
			p.SetState(110)
			p.InstanceMethods()
		}

	}
	{
		p.SetState(113)
		p.Match(SmallTalkParserRIGHT_BRK)
	}

	return localctx
}

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *SmallTalkParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SmallTalkParserRULE_main)

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
		p.SetState(115)
		p.Body()
	}

	return localctx
}

// IInstanceMethodsContext is an interface to support dynamic dispatch.
type IInstanceMethodsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstanceMethodsContext differentiates from other interfaces.
	IsInstanceMethodsContext()
}

type InstanceMethodsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstanceMethodsContext() *InstanceMethodsContext {
	var p = new(InstanceMethodsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_instanceMethods
	return p
}

func (*InstanceMethodsContext) IsInstanceMethodsContext() {}

func NewInstanceMethodsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceMethodsContext {
	var p = new(InstanceMethodsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_instanceMethods

	return p
}

func (s *InstanceMethodsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceMethodsContext) AllMethod() []IMethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodContext)(nil)).Elem())
	var tst = make([]IMethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodContext)
		}
	}

	return tst
}

func (s *InstanceMethodsContext) Method(i int) IMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *InstanceMethodsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceMethodsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceMethodsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterInstanceMethods(s)
	}
}

func (s *InstanceMethodsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitInstanceMethods(s)
	}
}

func (p *SmallTalkParser) InstanceMethods() (localctx IInstanceMethodsContext) {
	localctx = NewInstanceMethodsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SmallTalkParserRULE_instanceMethods)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SmallTalkParserIDENTIFIER || _la == SmallTalkParserKEYWORD {
		{
			p.SetState(117)
			p.Method()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) UnaryMethod() IUnaryMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryMethodContext)
}

func (s *MethodContext) BinaryMethod() IBinaryMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryMethodContext)
}

func (s *MethodContext) KeywordMethod() IKeywordMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordMethodContext)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *SmallTalkParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SmallTalkParserRULE_method)

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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.UnaryMethod()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.BinaryMethod()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.KeywordMethod()
		}

	}

	return localctx
}

// IUnaryMethodContext is an interface to support dynamic dispatch.
type IUnaryMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsUnaryMethodContext differentiates from other interfaces.
	IsUnaryMethodContext()
}

type UnaryMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyUnaryMethodContext() *UnaryMethodContext {
	var p = new(UnaryMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_unaryMethod
	return p
}

func (*UnaryMethodContext) IsUnaryMethodContext() {}

func NewUnaryMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryMethodContext {
	var p = new(UnaryMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_unaryMethod

	return p
}

func (s *UnaryMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryMethodContext) GetName() antlr.Token { return s.name }

func (s *UnaryMethodContext) SetName(v antlr.Token) { s.name = v }

func (s *UnaryMethodContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserLEFT_BRK, 0)
}

func (s *UnaryMethodContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *UnaryMethodContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserRIGHT_BRK, 0)
}

func (s *UnaryMethodContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *UnaryMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterUnaryMethod(s)
	}
}

func (s *UnaryMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitUnaryMethod(s)
	}
}

func (p *SmallTalkParser) UnaryMethod() (localctx IUnaryMethodContext) {
	localctx = NewUnaryMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SmallTalkParserRULE_unaryMethod)

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
		p.SetState(127)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*UnaryMethodContext).name = _m
	}
	{
		p.SetState(128)
		p.Match(SmallTalkParserLEFT_BRK)
	}
	{
		p.SetState(129)
		p.Body()
	}
	{
		p.SetState(130)
		p.Match(SmallTalkParserRIGHT_BRK)
	}

	return localctx
}

// IBinaryMethodContext is an interface to support dynamic dispatch.
type IBinaryMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsBinaryMethodContext differentiates from other interfaces.
	IsBinaryMethodContext()
}

type BinaryMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyBinaryMethodContext() *BinaryMethodContext {
	var p = new(BinaryMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_binaryMethod
	return p
}

func (*BinaryMethodContext) IsBinaryMethodContext() {}

func NewBinaryMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryMethodContext {
	var p = new(BinaryMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_binaryMethod

	return p
}

func (s *BinaryMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryMethodContext) GetName() antlr.Token { return s.name }

func (s *BinaryMethodContext) SetName(v antlr.Token) { s.name = v }

func (s *BinaryMethodContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserLEFT_BRK, 0)
}

func (s *BinaryMethodContext) MethodArg() IMethodArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodArgContext)
}

func (s *BinaryMethodContext) PIPE() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserPIPE, 0)
}

func (s *BinaryMethodContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *BinaryMethodContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserRIGHT_BRK, 0)
}

func (s *BinaryMethodContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *BinaryMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterBinaryMethod(s)
	}
}

func (s *BinaryMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitBinaryMethod(s)
	}
}

func (p *SmallTalkParser) BinaryMethod() (localctx IBinaryMethodContext) {
	localctx = NewBinaryMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SmallTalkParserRULE_binaryMethod)

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

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*BinaryMethodContext).name = _m
	}
	{
		p.SetState(133)
		p.Match(SmallTalkParserLEFT_BRK)
	}
	{
		p.SetState(134)
		p.MethodArg()
	}
	{
		p.SetState(135)
		p.Match(SmallTalkParserPIPE)
	}
	{
		p.SetState(136)
		p.Body()
	}
	{
		p.SetState(137)
		p.Match(SmallTalkParserRIGHT_BRK)
	}

	return localctx
}

// IKeywordMethodContext is an interface to support dynamic dispatch.
type IKeywordMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsKeywordMethodContext differentiates from other interfaces.
	IsKeywordMethodContext()
}

type KeywordMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyKeywordMethodContext() *KeywordMethodContext {
	var p = new(KeywordMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_keywordMethod
	return p
}

func (*KeywordMethodContext) IsKeywordMethodContext() {}

func NewKeywordMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordMethodContext {
	var p = new(KeywordMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_keywordMethod

	return p
}

func (s *KeywordMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordMethodContext) GetName() antlr.Token { return s.name }

func (s *KeywordMethodContext) SetName(v antlr.Token) { s.name = v }

func (s *KeywordMethodContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserLEFT_BRK, 0)
}

func (s *KeywordMethodContext) MethodArgs() IMethodArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodArgsContext)
}

func (s *KeywordMethodContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *KeywordMethodContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserRIGHT_BRK, 0)
}

func (s *KeywordMethodContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserKEYWORD, 0)
}

func (s *KeywordMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterKeywordMethod(s)
	}
}

func (s *KeywordMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitKeywordMethod(s)
	}
}

func (p *SmallTalkParser) KeywordMethod() (localctx IKeywordMethodContext) {
	localctx = NewKeywordMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SmallTalkParserRULE_keywordMethod)

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

		var _m = p.Match(SmallTalkParserKEYWORD)

		localctx.(*KeywordMethodContext).name = _m
	}
	{
		p.SetState(140)
		p.Match(SmallTalkParserLEFT_BRK)
	}
	{
		p.SetState(141)
		p.MethodArgs()
	}
	{
		p.SetState(142)
		p.Body()
	}
	{
		p.SetState(143)
		p.Match(SmallTalkParserRIGHT_BRK)
	}

	return localctx
}

// IMethodArgsContext is an interface to support dynamic dispatch.
type IMethodArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodArgsContext differentiates from other interfaces.
	IsMethodArgsContext()
}

type MethodArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodArgsContext() *MethodArgsContext {
	var p = new(MethodArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_methodArgs
	return p
}

func (*MethodArgsContext) IsMethodArgsContext() {}

func NewMethodArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodArgsContext {
	var p = new(MethodArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_methodArgs

	return p
}

func (s *MethodArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodArgsContext) PIPE() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserPIPE, 0)
}

func (s *MethodArgsContext) AllMethodArg() []IMethodArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodArgContext)(nil)).Elem())
	var tst = make([]IMethodArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodArgContext)
		}
	}

	return tst
}

func (s *MethodArgsContext) MethodArg(i int) IMethodArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodArgContext)
}

func (s *MethodArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMethodArgs(s)
	}
}

func (s *MethodArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMethodArgs(s)
	}
}

func (p *SmallTalkParser) MethodArgs() (localctx IMethodArgsContext) {
	localctx = NewMethodArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SmallTalkParserRULE_methodArgs)
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
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SmallTalkParserCOLON {
		{
			p.SetState(145)
			p.MethodArg()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(SmallTalkParserPIPE)
	}

	return localctx
}

// IMethodArgContext is an interface to support dynamic dispatch.
type IMethodArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsMethodArgContext differentiates from other interfaces.
	IsMethodArgContext()
}

type MethodArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyMethodArgContext() *MethodArgContext {
	var p = new(MethodArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_methodArg
	return p
}

func (*MethodArgContext) IsMethodArgContext() {}

func NewMethodArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodArgContext {
	var p = new(MethodArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_methodArg

	return p
}

func (s *MethodArgContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodArgContext) GetName() antlr.Token { return s.name }

func (s *MethodArgContext) SetName(v antlr.Token) { s.name = v }

func (s *MethodArgContext) COLON() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserCOLON, 0)
}

func (s *MethodArgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *MethodArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMethodArg(s)
	}
}

func (s *MethodArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMethodArg(s)
	}
}

func (p *SmallTalkParser) MethodArg() (localctx IMethodArgContext) {
	localctx = NewMethodArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SmallTalkParserRULE_methodArg)

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
		p.SetState(152)
		p.Match(SmallTalkParserCOLON)
	}
	{
		p.SetState(153)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*MethodArgContext).name = _m
	}

	return localctx
}

// ILocalVarsContext is an interface to support dynamic dispatch.
type ILocalVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalVarsContext differentiates from other interfaces.
	IsLocalVarsContext()
}

type LocalVarsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalVarsContext() *LocalVarsContext {
	var p = new(LocalVarsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_localVars
	return p
}

func (*LocalVarsContext) IsLocalVarsContext() {}

func NewLocalVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalVarsContext {
	var p = new(LocalVarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_localVars

	return p
}

func (s *LocalVarsContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalVarsContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(SmallTalkParserPIPE)
}

func (s *LocalVarsContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(SmallTalkParserPIPE, i)
}

func (s *LocalVarsContext) AllLocalVar() []ILocalVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILocalVarContext)(nil)).Elem())
	var tst = make([]ILocalVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILocalVarContext)
		}
	}

	return tst
}

func (s *LocalVarsContext) LocalVar(i int) ILocalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILocalVarContext)
}

func (s *LocalVarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalVarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalVarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLocalVars(s)
	}
}

func (s *LocalVarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLocalVars(s)
	}
}

func (p *SmallTalkParser) LocalVars() (localctx ILocalVarsContext) {
	localctx = NewLocalVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SmallTalkParserRULE_localVars)
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
		p.SetState(155)
		p.Match(SmallTalkParserPIPE)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SmallTalkParserIDENTIFIER {
		{
			p.SetState(156)
			p.LocalVar()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(161)
		p.Match(SmallTalkParserPIPE)
	}

	return localctx
}

// ILocalVarContext is an interface to support dynamic dispatch.
type ILocalVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsLocalVarContext differentiates from other interfaces.
	IsLocalVarContext()
}

type LocalVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyLocalVarContext() *LocalVarContext {
	var p = new(LocalVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_localVar
	return p
}

func (*LocalVarContext) IsLocalVarContext() {}

func NewLocalVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalVarContext {
	var p = new(LocalVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_localVar

	return p
}

func (s *LocalVarContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalVarContext) GetName() antlr.Token { return s.name }

func (s *LocalVarContext) SetName(v antlr.Token) { s.name = v }

func (s *LocalVarContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *LocalVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLocalVar(s)
	}
}

func (s *LocalVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLocalVar(s)
	}
}

func (p *SmallTalkParser) LocalVar() (localctx ILocalVarContext) {
	localctx = NewLocalVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SmallTalkParserRULE_localVar)

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
		p.SetState(163)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*LocalVarContext).name = _m
	}

	return localctx
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) LocalVars() ILocalVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalVarsContext)
}

func (s *BodyContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitBody(s)
	}
}

func (p *SmallTalkParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SmallTalkParserRULE_body)
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
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SmallTalkParserPIPE {
		{
			p.SetState(165)
			p.LocalVars()
		}

	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(SmallTalkParserNIL-5))|(1<<(SmallTalkParserTRUE-5))|(1<<(SmallTalkParserFALSE-5))|(1<<(SmallTalkParserSELF-5))|(1<<(SmallTalkParserSUPER-5))|(1<<(SmallTalkParserLEFT_BRK-5))|(1<<(SmallTalkParserLEFT_PRT-5))|(1<<(SmallTalkParserIDENTIFIER-5))|(1<<(SmallTalkParserBIN_INT_NUMBER-5))|(1<<(SmallTalkParserOCT_INT_NUMBER-5))|(1<<(SmallTalkParserDEC_INT_NUMBER-5))|(1<<(SmallTalkParserHEX_INT_NUMBER-5))|(1<<(SmallTalkParserFLOAT_NUMBER-5))|(1<<(SmallTalkParserCHAR-5))|(1<<(SmallTalkParserSTRING-5)))) != 0 {
		{
			p.SetState(168)
			p.Statements()
		}

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
	p.RuleIndex = SmallTalkParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_statements

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

func (s *StatementsContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SmallTalkParserDOT)
}

func (s *StatementsContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SmallTalkParserDOT, i)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *SmallTalkParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SmallTalkParserRULE_statements)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Statement()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(172)
				p.Match(SmallTalkParserDOT)
			}
			{
				p.SetState(173)
				p.Statement()
			}

		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	{
		p.SetState(179)
		p.Match(SmallTalkParserDOT)
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
	p.RuleIndex = SmallTalkParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) MessageStatement() IMessageStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageStatementContext)
}

func (s *StatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *SmallTalkParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SmallTalkParserRULE_statement)

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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.MessageStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.AssignmentStatement()
		}

	}

	return localctx
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_assignmentStatement
	return p
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) AllAssignmentItem() []IAssignmentItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignmentItemContext)(nil)).Elem())
	var tst = make([]IAssignmentItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignmentItemContext)
		}
	}

	return tst
}

func (s *AssignmentStatementContext) AssignmentItem(i int) IAssignmentItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignmentItemContext)
}

func (s *AssignmentStatementContext) MessageExpression() IMessageExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageExpressionContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (p *SmallTalkParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SmallTalkParserRULE_assignmentStatement)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.AssignmentItem()
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(186)
				p.AssignmentItem()
			}

		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}
	{
		p.SetState(192)
		p.messageExpression(0)
	}

	return localctx
}

// IAssignmentItemContext is an interface to support dynamic dispatch.
type IAssignmentItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsAssignmentItemContext differentiates from other interfaces.
	IsAssignmentItemContext()
}

type AssignmentItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyAssignmentItemContext() *AssignmentItemContext {
	var p = new(AssignmentItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_assignmentItem
	return p
}

func (*AssignmentItemContext) IsAssignmentItemContext() {}

func NewAssignmentItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentItemContext {
	var p = new(AssignmentItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_assignmentItem

	return p
}

func (s *AssignmentItemContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentItemContext) GetName() antlr.Token { return s.name }

func (s *AssignmentItemContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignmentItemContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserASSIGNMENT, 0)
}

func (s *AssignmentItemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *AssignmentItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterAssignmentItem(s)
	}
}

func (s *AssignmentItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitAssignmentItem(s)
	}
}

func (p *SmallTalkParser) AssignmentItem() (localctx IAssignmentItemContext) {
	localctx = NewAssignmentItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SmallTalkParserRULE_assignmentItem)

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

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*AssignmentItemContext).name = _m
	}
	{
		p.SetState(195)
		p.Match(SmallTalkParserASSIGNMENT)
	}

	return localctx
}

// IMessageStatementContext is an interface to support dynamic dispatch.
type IMessageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageStatementContext differentiates from other interfaces.
	IsMessageStatementContext()
}

type MessageStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageStatementContext() *MessageStatementContext {
	var p = new(MessageStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_messageStatement
	return p
}

func (*MessageStatementContext) IsMessageStatementContext() {}

func NewMessageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageStatementContext {
	var p = new(MessageStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_messageStatement

	return p
}

func (s *MessageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageStatementContext) MessageExpression() IMessageExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageExpressionContext)
}

func (s *MessageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMessageStatement(s)
	}
}

func (s *MessageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMessageStatement(s)
	}
}

func (p *SmallTalkParser) MessageStatement() (localctx IMessageStatementContext) {
	localctx = NewMessageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SmallTalkParserRULE_messageStatement)

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
		p.SetState(197)
		p.messageExpression(0)
	}

	return localctx
}

// IMessageExpressionContext is an interface to support dynamic dispatch.
type IMessageExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageExpressionContext differentiates from other interfaces.
	IsMessageExpressionContext()
}

type MessageExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageExpressionContext() *MessageExpressionContext {
	var p = new(MessageExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_messageExpression
	return p
}

func (*MessageExpressionContext) IsMessageExpressionContext() {}

func NewMessageExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageExpressionContext {
	var p = new(MessageExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_messageExpression

	return p
}

func (s *MessageExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageExpressionContext) PrtStatement() IPrtStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrtStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrtStatementContext)
}

func (s *MessageExpressionContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *MessageExpressionContext) MessageExpression() IMessageExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageExpressionContext)
}

func (s *MessageExpressionContext) MessageRange() IMessageRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageRangeContext)
}

func (s *MessageExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMessageExpression(s)
	}
}

func (s *MessageExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMessageExpression(s)
	}
}

func (p *SmallTalkParser) MessageExpression() (localctx IMessageExpressionContext) {
	return p.messageExpression(0)
}

func (p *SmallTalkParser) messageExpression(_p int) (localctx IMessageExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMessageExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMessageExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, SmallTalkParserRULE_messageExpression, _p)

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
	p.SetState(202)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SmallTalkParserLEFT_PRT:
		{
			p.SetState(200)
			p.PrtStatement()
		}

	case SmallTalkParserNIL, SmallTalkParserTRUE, SmallTalkParserFALSE, SmallTalkParserSELF, SmallTalkParserSUPER, SmallTalkParserLEFT_BRK, SmallTalkParserIDENTIFIER, SmallTalkParserBIN_INT_NUMBER, SmallTalkParserOCT_INT_NUMBER, SmallTalkParserDEC_INT_NUMBER, SmallTalkParserHEX_INT_NUMBER, SmallTalkParserFLOAT_NUMBER, SmallTalkParserCHAR, SmallTalkParserSTRING:
		{
			p.SetState(201)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMessageExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SmallTalkParserRULE_messageExpression)
			p.SetState(204)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(205)
				p.MessageRange()
			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) KeywordMessage() IKeywordMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordMessageContext)
}

func (s *MessageContext) BinaryMessage() IBinaryMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryMessageContext)
}

func (s *MessageContext) UnaryMessage() IUnaryMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryMessageContext)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (p *SmallTalkParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SmallTalkParserRULE_message)

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

	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.KeywordMessage()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.BinaryMessage()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(213)
			p.UnaryMessage()
		}

	}

	return localctx
}

// IMessageRangeContext is an interface to support dynamic dispatch.
type IMessageRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageRangeContext differentiates from other interfaces.
	IsMessageRangeContext()
}

type MessageRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageRangeContext() *MessageRangeContext {
	var p = new(MessageRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_messageRange
	return p
}

func (*MessageRangeContext) IsMessageRangeContext() {}

func NewMessageRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageRangeContext {
	var p = new(MessageRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_messageRange

	return p
}

func (s *MessageRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageRangeContext) AllMessage() []IMessageContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageContext)(nil)).Elem())
	var tst = make([]IMessageContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageContext)
		}
	}

	return tst
}

func (s *MessageRangeContext) Message(i int) IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *MessageRangeContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(SmallTalkParserSEMICOLON)
}

func (s *MessageRangeContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(SmallTalkParserSEMICOLON, i)
}

func (s *MessageRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMessageRange(s)
	}
}

func (s *MessageRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMessageRange(s)
	}
}

func (p *SmallTalkParser) MessageRange() (localctx IMessageRangeContext) {
	localctx = NewMessageRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SmallTalkParserRULE_messageRange)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Message()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(217)
				p.Match(SmallTalkParserSEMICOLON)
			}
			{
				p.SetState(218)
				p.Message()
			}

		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryMessageContext is an interface to support dynamic dispatch.
type IUnaryMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMethodName returns the methodName token.
	GetMethodName() antlr.Token

	// SetMethodName sets the methodName token.
	SetMethodName(antlr.Token)

	// IsUnaryMessageContext differentiates from other interfaces.
	IsUnaryMessageContext()
}

type UnaryMessageContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	methodName antlr.Token
}

func NewEmptyUnaryMessageContext() *UnaryMessageContext {
	var p = new(UnaryMessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_unaryMessage
	return p
}

func (*UnaryMessageContext) IsUnaryMessageContext() {}

func NewUnaryMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryMessageContext {
	var p = new(UnaryMessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_unaryMessage

	return p
}

func (s *UnaryMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryMessageContext) GetMethodName() antlr.Token { return s.methodName }

func (s *UnaryMessageContext) SetMethodName(v antlr.Token) { s.methodName = v }

func (s *UnaryMessageContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *UnaryMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterUnaryMessage(s)
	}
}

func (s *UnaryMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitUnaryMessage(s)
	}
}

func (p *SmallTalkParser) UnaryMessage() (localctx IUnaryMessageContext) {
	localctx = NewUnaryMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SmallTalkParserRULE_unaryMessage)

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
		p.SetState(224)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*UnaryMessageContext).methodName = _m
	}

	return localctx
}

// IBinaryMessageContext is an interface to support dynamic dispatch.
type IBinaryMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdName returns the idName rule contexts.
	GetIdName() IIdentifierContext

	// GetOpName returns the opName rule contexts.
	GetOpName() IOperatorContext

	// SetIdName sets the idName rule contexts.
	SetIdName(IIdentifierContext)

	// SetOpName sets the opName rule contexts.
	SetOpName(IOperatorContext)

	// IsBinaryMessageContext differentiates from other interfaces.
	IsBinaryMessageContext()
}

type BinaryMessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	idName IIdentifierContext
	opName IOperatorContext
}

func NewEmptyBinaryMessageContext() *BinaryMessageContext {
	var p = new(BinaryMessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_binaryMessage
	return p
}

func (*BinaryMessageContext) IsBinaryMessageContext() {}

func NewBinaryMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryMessageContext {
	var p = new(BinaryMessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_binaryMessage

	return p
}

func (s *BinaryMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryMessageContext) GetIdName() IIdentifierContext { return s.idName }

func (s *BinaryMessageContext) GetOpName() IOperatorContext { return s.opName }

func (s *BinaryMessageContext) SetIdName(v IIdentifierContext) { s.idName = v }

func (s *BinaryMessageContext) SetOpName(v IOperatorContext) { s.opName = v }

func (s *BinaryMessageContext) MessageArg() IMessageArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageArgContext)
}

func (s *BinaryMessageContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *BinaryMessageContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *BinaryMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterBinaryMessage(s)
	}
}

func (s *BinaryMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitBinaryMessage(s)
	}
}

func (p *SmallTalkParser) BinaryMessage() (localctx IBinaryMessageContext) {
	localctx = NewBinaryMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SmallTalkParserRULE_binaryMessage)

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
	p.SetState(228)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SmallTalkParserIDENTIFIER:
		{
			p.SetState(226)

			var _x = p.Identifier()

			localctx.(*BinaryMessageContext).idName = _x
		}

	case SmallTalkParserT__0, SmallTalkParserT__1, SmallTalkParserT__2, SmallTalkParserT__3:
		{
			p.SetState(227)

			var _x = p.Operator()

			localctx.(*BinaryMessageContext).opName = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(230)
		p.MessageArg()
	}

	return localctx
}

// IKeywordMessageContext is an interface to support dynamic dispatch.
type IKeywordMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMethodName returns the methodName token.
	GetMethodName() antlr.Token

	// SetMethodName sets the methodName token.
	SetMethodName(antlr.Token)

	// IsKeywordMessageContext differentiates from other interfaces.
	IsKeywordMessageContext()
}

type KeywordMessageContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	methodName antlr.Token
}

func NewEmptyKeywordMessageContext() *KeywordMessageContext {
	var p = new(KeywordMessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_keywordMessage
	return p
}

func (*KeywordMessageContext) IsKeywordMessageContext() {}

func NewKeywordMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordMessageContext {
	var p = new(KeywordMessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_keywordMessage

	return p
}

func (s *KeywordMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordMessageContext) GetMethodName() antlr.Token { return s.methodName }

func (s *KeywordMessageContext) SetMethodName(v antlr.Token) { s.methodName = v }

func (s *KeywordMessageContext) KeywordMessageArgs() IKeywordMessageArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordMessageArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordMessageArgsContext)
}

func (s *KeywordMessageContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserKEYWORD, 0)
}

func (s *KeywordMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterKeywordMessage(s)
	}
}

func (s *KeywordMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitKeywordMessage(s)
	}
}

func (p *SmallTalkParser) KeywordMessage() (localctx IKeywordMessageContext) {
	localctx = NewKeywordMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SmallTalkParserRULE_keywordMessage)

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
		p.SetState(232)

		var _m = p.Match(SmallTalkParserKEYWORD)

		localctx.(*KeywordMessageContext).methodName = _m
	}
	{
		p.SetState(233)
		p.KeywordMessageArgs()
	}

	return localctx
}

// IKeywordMessageArgsContext is an interface to support dynamic dispatch.
type IKeywordMessageArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordMessageArgsContext differentiates from other interfaces.
	IsKeywordMessageArgsContext()
}

type KeywordMessageArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordMessageArgsContext() *KeywordMessageArgsContext {
	var p = new(KeywordMessageArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_keywordMessageArgs
	return p
}

func (*KeywordMessageArgsContext) IsKeywordMessageArgsContext() {}

func NewKeywordMessageArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordMessageArgsContext {
	var p = new(KeywordMessageArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_keywordMessageArgs

	return p
}

func (s *KeywordMessageArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordMessageArgsContext) AllMessageArg() []IMessageArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageArgContext)(nil)).Elem())
	var tst = make([]IMessageArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageArgContext)
		}
	}

	return tst
}

func (s *KeywordMessageArgsContext) MessageArg(i int) IMessageArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageArgContext)
}

func (s *KeywordMessageArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordMessageArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordMessageArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterKeywordMessageArgs(s)
	}
}

func (s *KeywordMessageArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitKeywordMessageArgs(s)
	}
}

func (p *SmallTalkParser) KeywordMessageArgs() (localctx IKeywordMessageArgsContext) {
	localctx = NewKeywordMessageArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SmallTalkParserRULE_keywordMessageArgs)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(235)
				p.MessageArg()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IMessageArgContext is an interface to support dynamic dispatch.
type IMessageArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageArgContext differentiates from other interfaces.
	IsMessageArgContext()
}

type MessageArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageArgContext() *MessageArgContext {
	var p = new(MessageArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_messageArg
	return p
}

func (*MessageArgContext) IsMessageArgContext() {}

func NewMessageArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageArgContext {
	var p = new(MessageArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_messageArg

	return p
}

func (s *MessageArgContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageArgContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *MessageArgContext) PrtStatement() IPrtStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrtStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrtStatementContext)
}

func (s *MessageArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMessageArg(s)
	}
}

func (s *MessageArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMessageArg(s)
	}
}

func (p *SmallTalkParser) MessageArg() (localctx IMessageArgContext) {
	localctx = NewMessageArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SmallTalkParserRULE_messageArg)

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

	p.SetState(242)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SmallTalkParserNIL, SmallTalkParserTRUE, SmallTalkParserFALSE, SmallTalkParserSELF, SmallTalkParserSUPER, SmallTalkParserLEFT_BRK, SmallTalkParserIDENTIFIER, SmallTalkParserBIN_INT_NUMBER, SmallTalkParserOCT_INT_NUMBER, SmallTalkParserDEC_INT_NUMBER, SmallTalkParserHEX_INT_NUMBER, SmallTalkParserFLOAT_NUMBER, SmallTalkParserCHAR, SmallTalkParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.Value()
		}

	case SmallTalkParserLEFT_PRT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.PrtStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrtStatementContext is an interface to support dynamic dispatch.
type IPrtStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrtStatementContext differentiates from other interfaces.
	IsPrtStatementContext()
}

type PrtStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrtStatementContext() *PrtStatementContext {
	var p = new(PrtStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_prtStatement
	return p
}

func (*PrtStatementContext) IsPrtStatementContext() {}

func NewPrtStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrtStatementContext {
	var p = new(PrtStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_prtStatement

	return p
}

func (s *PrtStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrtStatementContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserLEFT_PRT, 0)
}

func (s *PrtStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *PrtStatementContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserRIGHT_PRT, 0)
}

func (s *PrtStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrtStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrtStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterPrtStatement(s)
	}
}

func (s *PrtStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitPrtStatement(s)
	}
}

func (p *SmallTalkParser) PrtStatement() (localctx IPrtStatementContext) {
	localctx = NewPrtStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SmallTalkParserRULE_prtStatement)

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
		p.SetState(244)
		p.Match(SmallTalkParserLEFT_PRT)
	}
	{
		p.SetState(245)
		p.Statement()
	}
	{
		p.SetState(246)
		p.Match(SmallTalkParserRIGHT_PRT)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) ValueVar() IValueVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueVarContext)
}

func (s *ValueContext) ValueSelf() IValueSelfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueSelfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueSelfContext)
}

func (s *ValueContext) ValueSuper() IValueSuperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueSuperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueSuperContext)
}

func (s *ValueContext) ValueBlock() IValueBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueBlockContext)
}

func (s *ValueContext) ValueConst() IValueConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *SmallTalkParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SmallTalkParserRULE_value)

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

	p.SetState(253)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SmallTalkParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.ValueVar()
		}

	case SmallTalkParserSELF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.ValueSelf()
		}

	case SmallTalkParserSUPER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(250)
			p.ValueSuper()
		}

	case SmallTalkParserLEFT_BRK:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(251)
			p.ValueBlock()
		}

	case SmallTalkParserNIL, SmallTalkParserTRUE, SmallTalkParserFALSE, SmallTalkParserBIN_INT_NUMBER, SmallTalkParserOCT_INT_NUMBER, SmallTalkParserDEC_INT_NUMBER, SmallTalkParserHEX_INT_NUMBER, SmallTalkParserFLOAT_NUMBER, SmallTalkParserCHAR, SmallTalkParserSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(252)
			p.ValueConst()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueVarContext is an interface to support dynamic dispatch.
type IValueVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueVarContext differentiates from other interfaces.
	IsValueVarContext()
}

type ValueVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueVarContext() *ValueVarContext {
	var p = new(ValueVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueVar
	return p
}

func (*ValueVarContext) IsValueVarContext() {}

func NewValueVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueVarContext {
	var p = new(ValueVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueVar

	return p
}

func (s *ValueVarContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueVarContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *ValueVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueVar(s)
	}
}

func (s *ValueVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueVar(s)
	}
}

func (p *SmallTalkParser) ValueVar() (localctx IValueVarContext) {
	localctx = NewValueVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SmallTalkParserRULE_valueVar)

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
		p.SetState(255)
		p.Match(SmallTalkParserIDENTIFIER)
	}

	return localctx
}

// IValueSelfContext is an interface to support dynamic dispatch.
type IValueSelfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueSelfContext differentiates from other interfaces.
	IsValueSelfContext()
}

type ValueSelfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueSelfContext() *ValueSelfContext {
	var p = new(ValueSelfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueSelf
	return p
}

func (*ValueSelfContext) IsValueSelfContext() {}

func NewValueSelfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueSelfContext {
	var p = new(ValueSelfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueSelf

	return p
}

func (s *ValueSelfContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueSelfContext) SELF() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserSELF, 0)
}

func (s *ValueSelfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueSelfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueSelfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueSelf(s)
	}
}

func (s *ValueSelfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueSelf(s)
	}
}

func (p *SmallTalkParser) ValueSelf() (localctx IValueSelfContext) {
	localctx = NewValueSelfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SmallTalkParserRULE_valueSelf)

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
		p.SetState(257)
		p.Match(SmallTalkParserSELF)
	}

	return localctx
}

// IValueSuperContext is an interface to support dynamic dispatch.
type IValueSuperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueSuperContext differentiates from other interfaces.
	IsValueSuperContext()
}

type ValueSuperContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueSuperContext() *ValueSuperContext {
	var p = new(ValueSuperContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueSuper
	return p
}

func (*ValueSuperContext) IsValueSuperContext() {}

func NewValueSuperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueSuperContext {
	var p = new(ValueSuperContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueSuper

	return p
}

func (s *ValueSuperContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueSuperContext) SUPER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserSUPER, 0)
}

func (s *ValueSuperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueSuperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueSuperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueSuper(s)
	}
}

func (s *ValueSuperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueSuper(s)
	}
}

func (p *SmallTalkParser) ValueSuper() (localctx IValueSuperContext) {
	localctx = NewValueSuperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SmallTalkParserRULE_valueSuper)

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
		p.SetState(259)
		p.Match(SmallTalkParserSUPER)
	}

	return localctx
}

// IValueBlockContext is an interface to support dynamic dispatch.
type IValueBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueBlockContext differentiates from other interfaces.
	IsValueBlockContext()
}

type ValueBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueBlockContext() *ValueBlockContext {
	var p = new(ValueBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueBlock
	return p
}

func (*ValueBlockContext) IsValueBlockContext() {}

func NewValueBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueBlockContext {
	var p = new(ValueBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueBlock

	return p
}

func (s *ValueBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueBlockContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserLEFT_BRK, 0)
}

func (s *ValueBlockContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ValueBlockContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserRIGHT_BRK, 0)
}

func (s *ValueBlockContext) BlockArgs() IBlockArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockArgsContext)
}

func (s *ValueBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueBlock(s)
	}
}

func (s *ValueBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueBlock(s)
	}
}

func (p *SmallTalkParser) ValueBlock() (localctx IValueBlockContext) {
	localctx = NewValueBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SmallTalkParserRULE_valueBlock)
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
		p.SetState(261)
		p.Match(SmallTalkParserLEFT_BRK)
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SmallTalkParserCOLON {
		{
			p.SetState(262)
			p.BlockArgs()
		}

	}
	{
		p.SetState(265)
		p.Body()
	}
	{
		p.SetState(266)
		p.Match(SmallTalkParserRIGHT_BRK)
	}

	return localctx
}

// IBlockArgsContext is an interface to support dynamic dispatch.
type IBlockArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockArgsContext differentiates from other interfaces.
	IsBlockArgsContext()
}

type BlockArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockArgsContext() *BlockArgsContext {
	var p = new(BlockArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_blockArgs
	return p
}

func (*BlockArgsContext) IsBlockArgsContext() {}

func NewBlockArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockArgsContext {
	var p = new(BlockArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_blockArgs

	return p
}

func (s *BlockArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockArgsContext) PIPE() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserPIPE, 0)
}

func (s *BlockArgsContext) AllBlockArg() []IBlockArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockArgContext)(nil)).Elem())
	var tst = make([]IBlockArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockArgContext)
		}
	}

	return tst
}

func (s *BlockArgsContext) BlockArg(i int) IBlockArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockArgContext)
}

func (s *BlockArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterBlockArgs(s)
	}
}

func (s *BlockArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitBlockArgs(s)
	}
}

func (p *SmallTalkParser) BlockArgs() (localctx IBlockArgsContext) {
	localctx = NewBlockArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SmallTalkParserRULE_blockArgs)
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

	for ok := true; ok; ok = _la == SmallTalkParserCOLON {
		{
			p.SetState(268)
			p.BlockArg()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(273)
		p.Match(SmallTalkParserPIPE)
	}

	return localctx
}

// IBlockArgContext is an interface to support dynamic dispatch.
type IBlockArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsBlockArgContext differentiates from other interfaces.
	IsBlockArgContext()
}

type BlockArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyBlockArgContext() *BlockArgContext {
	var p = new(BlockArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_blockArg
	return p
}

func (*BlockArgContext) IsBlockArgContext() {}

func NewBlockArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockArgContext {
	var p = new(BlockArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_blockArg

	return p
}

func (s *BlockArgContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockArgContext) GetName() antlr.Token { return s.name }

func (s *BlockArgContext) SetName(v antlr.Token) { s.name = v }

func (s *BlockArgContext) COLON() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserCOLON, 0)
}

func (s *BlockArgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *BlockArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterBlockArg(s)
	}
}

func (s *BlockArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitBlockArg(s)
	}
}

func (p *SmallTalkParser) BlockArg() (localctx IBlockArgContext) {
	localctx = NewBlockArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SmallTalkParserRULE_blockArg)

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
		p.SetState(275)
		p.Match(SmallTalkParserCOLON)
	}
	{
		p.SetState(276)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*BlockArgContext).name = _m
	}

	return localctx
}

// IValueConstContext is an interface to support dynamic dispatch.
type IValueConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstContext differentiates from other interfaces.
	IsValueConstContext()
}

type ValueConstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstContext() *ValueConstContext {
	var p = new(ValueConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConst
	return p
}

func (*ValueConstContext) IsValueConstContext() {}

func NewValueConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstContext {
	var p = new(ValueConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConst

	return p
}

func (s *ValueConstContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstContext) ValueConstNum() IValueConstNumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstNumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstNumContext)
}

func (s *ValueConstContext) ValueConstChar() IValueConstCharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstCharContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstCharContext)
}

func (s *ValueConstContext) ValueConstString() IValueConstStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstStringContext)
}

func (s *ValueConstContext) ValueConstBool() IValueConstBoolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstBoolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstBoolContext)
}

func (s *ValueConstContext) ValueNil() IValueNilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueNilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueNilContext)
}

func (s *ValueConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConst(s)
	}
}

func (s *ValueConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConst(s)
	}
}

func (p *SmallTalkParser) ValueConst() (localctx IValueConstContext) {
	localctx = NewValueConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SmallTalkParserRULE_valueConst)

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

	p.SetState(283)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SmallTalkParserBIN_INT_NUMBER, SmallTalkParserOCT_INT_NUMBER, SmallTalkParserDEC_INT_NUMBER, SmallTalkParserHEX_INT_NUMBER, SmallTalkParserFLOAT_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(278)
			p.ValueConstNum()
		}

	case SmallTalkParserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(279)
			p.ValueConstChar()
		}

	case SmallTalkParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(280)
			p.ValueConstString()
		}

	case SmallTalkParserTRUE, SmallTalkParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(281)
			p.ValueConstBool()
		}

	case SmallTalkParserNIL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(282)
			p.ValueNil()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueConstNumContext is an interface to support dynamic dispatch.
type IValueConstNumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstNumContext differentiates from other interfaces.
	IsValueConstNumContext()
}

type ValueConstNumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstNumContext() *ValueConstNumContext {
	var p = new(ValueConstNumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConstNum
	return p
}

func (*ValueConstNumContext) IsValueConstNumContext() {}

func NewValueConstNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstNumContext {
	var p = new(ValueConstNumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConstNum

	return p
}

func (s *ValueConstNumContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstNumContext) ValueConstBinInt() IValueConstBinIntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstBinIntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstBinIntContext)
}

func (s *ValueConstNumContext) ValueConstOctInt() IValueConstOctIntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstOctIntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstOctIntContext)
}

func (s *ValueConstNumContext) ValueConstDecInt() IValueConstDecIntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstDecIntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstDecIntContext)
}

func (s *ValueConstNumContext) ValueConstHexInt() IValueConstHexIntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstHexIntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstHexIntContext)
}

func (s *ValueConstNumContext) ValueConstFloat() IValueConstFloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueConstFloatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueConstFloatContext)
}

func (s *ValueConstNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstNumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConstNum(s)
	}
}

func (s *ValueConstNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConstNum(s)
	}
}

func (p *SmallTalkParser) ValueConstNum() (localctx IValueConstNumContext) {
	localctx = NewValueConstNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SmallTalkParserRULE_valueConstNum)

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

	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SmallTalkParserBIN_INT_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.ValueConstBinInt()
		}

	case SmallTalkParserOCT_INT_NUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.ValueConstOctInt()
		}

	case SmallTalkParserDEC_INT_NUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(287)
			p.ValueConstDecInt()
		}

	case SmallTalkParserHEX_INT_NUMBER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(288)
			p.ValueConstHexInt()
		}

	case SmallTalkParserFLOAT_NUMBER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(289)
			p.ValueConstFloat()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueConstBinIntContext is an interface to support dynamic dispatch.
type IValueConstBinIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstBinIntContext differentiates from other interfaces.
	IsValueConstBinIntContext()
}

type ValueConstBinIntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstBinIntContext() *ValueConstBinIntContext {
	var p = new(ValueConstBinIntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConstBinInt
	return p
}

func (*ValueConstBinIntContext) IsValueConstBinIntContext() {}

func NewValueConstBinIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstBinIntContext {
	var p = new(ValueConstBinIntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConstBinInt

	return p
}

func (s *ValueConstBinIntContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstBinIntContext) BIN_INT_NUMBER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserBIN_INT_NUMBER, 0)
}

func (s *ValueConstBinIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstBinIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstBinIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConstBinInt(s)
	}
}

func (s *ValueConstBinIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConstBinInt(s)
	}
}

func (p *SmallTalkParser) ValueConstBinInt() (localctx IValueConstBinIntContext) {
	localctx = NewValueConstBinIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SmallTalkParserRULE_valueConstBinInt)

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
		p.Match(SmallTalkParserBIN_INT_NUMBER)
	}

	return localctx
}

// IValueConstOctIntContext is an interface to support dynamic dispatch.
type IValueConstOctIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstOctIntContext differentiates from other interfaces.
	IsValueConstOctIntContext()
}

type ValueConstOctIntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstOctIntContext() *ValueConstOctIntContext {
	var p = new(ValueConstOctIntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConstOctInt
	return p
}

func (*ValueConstOctIntContext) IsValueConstOctIntContext() {}

func NewValueConstOctIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstOctIntContext {
	var p = new(ValueConstOctIntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConstOctInt

	return p
}

func (s *ValueConstOctIntContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstOctIntContext) OCT_INT_NUMBER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserOCT_INT_NUMBER, 0)
}

func (s *ValueConstOctIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstOctIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstOctIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConstOctInt(s)
	}
}

func (s *ValueConstOctIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConstOctInt(s)
	}
}

func (p *SmallTalkParser) ValueConstOctInt() (localctx IValueConstOctIntContext) {
	localctx = NewValueConstOctIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SmallTalkParserRULE_valueConstOctInt)

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
		p.SetState(294)
		p.Match(SmallTalkParserOCT_INT_NUMBER)
	}

	return localctx
}

// IValueConstDecIntContext is an interface to support dynamic dispatch.
type IValueConstDecIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstDecIntContext differentiates from other interfaces.
	IsValueConstDecIntContext()
}

type ValueConstDecIntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstDecIntContext() *ValueConstDecIntContext {
	var p = new(ValueConstDecIntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConstDecInt
	return p
}

func (*ValueConstDecIntContext) IsValueConstDecIntContext() {}

func NewValueConstDecIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstDecIntContext {
	var p = new(ValueConstDecIntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConstDecInt

	return p
}

func (s *ValueConstDecIntContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstDecIntContext) DEC_INT_NUMBER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserDEC_INT_NUMBER, 0)
}

func (s *ValueConstDecIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstDecIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstDecIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConstDecInt(s)
	}
}

func (s *ValueConstDecIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConstDecInt(s)
	}
}

func (p *SmallTalkParser) ValueConstDecInt() (localctx IValueConstDecIntContext) {
	localctx = NewValueConstDecIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SmallTalkParserRULE_valueConstDecInt)

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
		p.SetState(296)
		p.Match(SmallTalkParserDEC_INT_NUMBER)
	}

	return localctx
}

// IValueConstHexIntContext is an interface to support dynamic dispatch.
type IValueConstHexIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstHexIntContext differentiates from other interfaces.
	IsValueConstHexIntContext()
}

type ValueConstHexIntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstHexIntContext() *ValueConstHexIntContext {
	var p = new(ValueConstHexIntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConstHexInt
	return p
}

func (*ValueConstHexIntContext) IsValueConstHexIntContext() {}

func NewValueConstHexIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstHexIntContext {
	var p = new(ValueConstHexIntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConstHexInt

	return p
}

func (s *ValueConstHexIntContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstHexIntContext) HEX_INT_NUMBER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserHEX_INT_NUMBER, 0)
}

func (s *ValueConstHexIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstHexIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstHexIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConstHexInt(s)
	}
}

func (s *ValueConstHexIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConstHexInt(s)
	}
}

func (p *SmallTalkParser) ValueConstHexInt() (localctx IValueConstHexIntContext) {
	localctx = NewValueConstHexIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SmallTalkParserRULE_valueConstHexInt)

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
		p.SetState(298)
		p.Match(SmallTalkParserHEX_INT_NUMBER)
	}

	return localctx
}

// IValueConstFloatContext is an interface to support dynamic dispatch.
type IValueConstFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstFloatContext differentiates from other interfaces.
	IsValueConstFloatContext()
}

type ValueConstFloatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstFloatContext() *ValueConstFloatContext {
	var p = new(ValueConstFloatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConstFloat
	return p
}

func (*ValueConstFloatContext) IsValueConstFloatContext() {}

func NewValueConstFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstFloatContext {
	var p = new(ValueConstFloatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConstFloat

	return p
}

func (s *ValueConstFloatContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstFloatContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserFLOAT_NUMBER, 0)
}

func (s *ValueConstFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstFloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConstFloat(s)
	}
}

func (s *ValueConstFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConstFloat(s)
	}
}

func (p *SmallTalkParser) ValueConstFloat() (localctx IValueConstFloatContext) {
	localctx = NewValueConstFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SmallTalkParserRULE_valueConstFloat)

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
		p.SetState(300)
		p.Match(SmallTalkParserFLOAT_NUMBER)
	}

	return localctx
}

// IValueConstCharContext is an interface to support dynamic dispatch.
type IValueConstCharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstCharContext differentiates from other interfaces.
	IsValueConstCharContext()
}

type ValueConstCharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstCharContext() *ValueConstCharContext {
	var p = new(ValueConstCharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConstChar
	return p
}

func (*ValueConstCharContext) IsValueConstCharContext() {}

func NewValueConstCharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstCharContext {
	var p = new(ValueConstCharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConstChar

	return p
}

func (s *ValueConstCharContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstCharContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserCHAR, 0)
}

func (s *ValueConstCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstCharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConstChar(s)
	}
}

func (s *ValueConstCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConstChar(s)
	}
}

func (p *SmallTalkParser) ValueConstChar() (localctx IValueConstCharContext) {
	localctx = NewValueConstCharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SmallTalkParserRULE_valueConstChar)

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
		p.SetState(302)
		p.Match(SmallTalkParserCHAR)
	}

	return localctx
}

// IValueConstStringContext is an interface to support dynamic dispatch.
type IValueConstStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstStringContext differentiates from other interfaces.
	IsValueConstStringContext()
}

type ValueConstStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstStringContext() *ValueConstStringContext {
	var p = new(ValueConstStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConstString
	return p
}

func (*ValueConstStringContext) IsValueConstStringContext() {}

func NewValueConstStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstStringContext {
	var p = new(ValueConstStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConstString

	return p
}

func (s *ValueConstStringContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserSTRING, 0)
}

func (s *ValueConstStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConstString(s)
	}
}

func (s *ValueConstStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConstString(s)
	}
}

func (p *SmallTalkParser) ValueConstString() (localctx IValueConstStringContext) {
	localctx = NewValueConstStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SmallTalkParserRULE_valueConstString)

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
		p.SetState(304)
		p.Match(SmallTalkParserSTRING)
	}

	return localctx
}

// IValueConstBoolContext is an interface to support dynamic dispatch.
type IValueConstBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueConstBoolContext differentiates from other interfaces.
	IsValueConstBoolContext()
}

type ValueConstBoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueConstBoolContext() *ValueConstBoolContext {
	var p = new(ValueConstBoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueConstBool
	return p
}

func (*ValueConstBoolContext) IsValueConstBoolContext() {}

func NewValueConstBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueConstBoolContext {
	var p = new(ValueConstBoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueConstBool

	return p
}

func (s *ValueConstBoolContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueConstBoolContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserTRUE, 0)
}

func (s *ValueConstBoolContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserFALSE, 0)
}

func (s *ValueConstBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueConstBoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueConstBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueConstBool(s)
	}
}

func (s *ValueConstBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueConstBool(s)
	}
}

func (p *SmallTalkParser) ValueConstBool() (localctx IValueConstBoolContext) {
	localctx = NewValueConstBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SmallTalkParserRULE_valueConstBool)
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
		p.SetState(306)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SmallTalkParserTRUE || _la == SmallTalkParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueNilContext is an interface to support dynamic dispatch.
type IValueNilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueNilContext differentiates from other interfaces.
	IsValueNilContext()
}

type ValueNilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueNilContext() *ValueNilContext {
	var p = new(ValueNilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_valueNil
	return p
}

func (*ValueNilContext) IsValueNilContext() {}

func NewValueNilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueNilContext {
	var p = new(ValueNilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_valueNil

	return p
}

func (s *ValueNilContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueNilContext) NIL() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserNIL, 0)
}

func (s *ValueNilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueNilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueNilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterValueNil(s)
	}
}

func (s *ValueNilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitValueNil(s)
	}
}

func (p *SmallTalkParser) ValueNil() (localctx IValueNilContext) {
	localctx = NewValueNilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SmallTalkParserRULE_valueNil)

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
		p.SetState(308)
		p.Match(SmallTalkParserNIL)
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *SmallTalkParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SmallTalkParserRULE_identifier)

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
		p.Match(SmallTalkParserIDENTIFIER)
	}

	return localctx
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *SmallTalkParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SmallTalkParserRULE_operator)
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
		p.SetState(312)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SmallTalkParserT__0)|(1<<SmallTalkParserT__1)|(1<<SmallTalkParserT__2)|(1<<SmallTalkParserT__3))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *SmallTalkParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *MessageExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MessageExpressionContext)
		}
		return p.MessageExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SmallTalkParser) MessageExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
