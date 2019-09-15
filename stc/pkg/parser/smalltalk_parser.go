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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 310,
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
	3, 5, 3, 110, 10, 3, 3, 3, 3, 3, 5, 3, 114, 10, 3, 3, 3, 5, 3, 117, 10,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 6, 5, 125, 10, 5, 13, 5, 14, 5,
	126, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 6, 7, 134, 10, 7, 13, 7, 14, 7, 135,
	3, 8, 3, 8, 5, 8, 140, 10, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 6, 11, 150, 10, 11, 13, 11, 14, 11, 151, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 5, 13, 160, 10, 13, 3, 13, 3, 13, 3, 13, 3, 14, 6,
	14, 166, 10, 14, 13, 14, 14, 14, 167, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	3, 16, 5, 16, 176, 10, 16, 3, 16, 5, 16, 179, 10, 16, 3, 17, 3, 17, 3,
	17, 7, 17, 184, 10, 17, 12, 17, 14, 17, 187, 11, 17, 3, 17, 5, 17, 190,
	10, 17, 3, 18, 3, 18, 5, 18, 194, 10, 18, 3, 19, 3, 19, 7, 19, 198, 10,
	19, 12, 19, 14, 19, 201, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25,
	218, 10, 25, 3, 26, 3, 26, 5, 26, 222, 10, 26, 3, 26, 7, 26, 225, 10, 26,
	12, 26, 14, 26, 228, 11, 26, 3, 27, 3, 27, 7, 27, 232, 10, 27, 12, 27,
	14, 27, 235, 11, 27, 3, 28, 3, 28, 6, 28, 239, 10, 28, 13, 28, 14, 28,
	240, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 250, 10, 30,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 270, 10, 35,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3,
	40, 3, 40, 6, 40, 284, 10, 40, 13, 40, 14, 40, 285, 3, 40, 3, 40, 3, 41,
	3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3,
	46, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 306, 10, 47, 3, 48, 3, 48, 3, 48,
	2, 2, 49, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 2, 6, 3, 2, 31, 32, 3,
	2, 34, 36, 3, 2, 9, 10, 3, 2, 3, 7, 2, 295, 2, 99, 3, 2, 2, 2, 4, 105,
	3, 2, 2, 2, 6, 120, 3, 2, 2, 2, 8, 122, 3, 2, 2, 2, 10, 130, 3, 2, 2, 2,
	12, 133, 3, 2, 2, 2, 14, 139, 3, 2, 2, 2, 16, 141, 3, 2, 2, 2, 18, 144,
	3, 2, 2, 2, 20, 147, 3, 2, 2, 2, 22, 155, 3, 2, 2, 2, 24, 157, 3, 2, 2,
	2, 26, 165, 3, 2, 2, 2, 28, 171, 3, 2, 2, 2, 30, 175, 3, 2, 2, 2, 32, 180,
	3, 2, 2, 2, 34, 193, 3, 2, 2, 2, 36, 195, 3, 2, 2, 2, 38, 204, 3, 2, 2,
	2, 40, 207, 3, 2, 2, 2, 42, 209, 3, 2, 2, 2, 44, 211, 3, 2, 2, 2, 46, 213,
	3, 2, 2, 2, 48, 215, 3, 2, 2, 2, 50, 219, 3, 2, 2, 2, 52, 229, 3, 2, 2,
	2, 54, 236, 3, 2, 2, 2, 56, 242, 3, 2, 2, 2, 58, 249, 3, 2, 2, 2, 60, 251,
	3, 2, 2, 2, 62, 253, 3, 2, 2, 2, 64, 255, 3, 2, 2, 2, 66, 257, 3, 2, 2,
	2, 68, 269, 3, 2, 2, 2, 70, 271, 3, 2, 2, 2, 72, 273, 3, 2, 2, 2, 74, 275,
	3, 2, 2, 2, 76, 277, 3, 2, 2, 2, 78, 280, 3, 2, 2, 2, 80, 289, 3, 2, 2,
	2, 82, 291, 3, 2, 2, 2, 84, 293, 3, 2, 2, 2, 86, 295, 3, 2, 2, 2, 88, 297,
	3, 2, 2, 2, 90, 299, 3, 2, 2, 2, 92, 305, 3, 2, 2, 2, 94, 307, 3, 2, 2,
	2, 96, 98, 5, 4, 3, 2, 97, 96, 3, 2, 2, 2, 98, 101, 3, 2, 2, 2, 99, 97,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 3, 2, 2, 2, 101, 99, 3, 2, 2,
	2, 102, 103, 5, 6, 4, 2, 103, 104, 7, 2, 2, 3, 104, 3, 3, 2, 2, 2, 105,
	106, 7, 11, 2, 2, 106, 109, 7, 31, 2, 2, 107, 108, 7, 19, 2, 2, 108, 110,
	7, 31, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 3, 2,
	2, 2, 111, 113, 7, 22, 2, 2, 112, 114, 5, 8, 5, 2, 113, 112, 3, 2, 2, 2,
	113, 114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2, 115, 117, 5, 12, 7, 2, 116,
	115, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119,
	7, 23, 2, 2, 119, 5, 3, 2, 2, 2, 120, 121, 5, 30, 16, 2, 121, 7, 3, 2,
	2, 2, 122, 124, 7, 17, 2, 2, 123, 125, 5, 10, 6, 2, 124, 123, 3, 2, 2,
	2, 125, 126, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127,
	128, 3, 2, 2, 2, 128, 129, 7, 17, 2, 2, 129, 9, 3, 2, 2, 2, 130, 131, 7,
	31, 2, 2, 131, 11, 3, 2, 2, 2, 132, 134, 5, 14, 8, 2, 133, 132, 3, 2, 2,
	2, 134, 135, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	13, 3, 2, 2, 2, 137, 140, 5, 16, 9, 2, 138, 140, 5, 18, 10, 2, 139, 137,
	3, 2, 2, 2, 139, 138, 3, 2, 2, 2, 140, 15, 3, 2, 2, 2, 141, 142, 7, 31,
	2, 2, 142, 143, 5, 24, 13, 2, 143, 17, 3, 2, 2, 2, 144, 145, 7, 32, 2,
	2, 145, 146, 5, 24, 13, 2, 146, 19, 3, 2, 2, 2, 147, 149, 7, 17, 2, 2,
	148, 150, 5, 22, 12, 2, 149, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151,
	149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154,
	7, 17, 2, 2, 154, 21, 3, 2, 2, 2, 155, 156, 7, 31, 2, 2, 156, 23, 3, 2,
	2, 2, 157, 159, 7, 22, 2, 2, 158, 160, 5, 26, 14, 2, 159, 158, 3, 2, 2,
	2, 159, 160, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 5, 30, 16, 2,
	162, 163, 7, 23, 2, 2, 163, 25, 3, 2, 2, 2, 164, 166, 5, 28, 15, 2, 165,
	164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168,
	3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 7, 17, 2, 2, 170, 27, 3, 2,
	2, 2, 171, 172, 7, 19, 2, 2, 172, 173, 7, 31, 2, 2, 173, 29, 3, 2, 2, 2,
	174, 176, 5, 20, 11, 2, 175, 174, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176,
	178, 3, 2, 2, 2, 177, 179, 5, 32, 17, 2, 178, 177, 3, 2, 2, 2, 178, 179,
	3, 2, 2, 2, 179, 31, 3, 2, 2, 2, 180, 185, 5, 34, 18, 2, 181, 182, 7, 16,
	2, 2, 182, 184, 5, 34, 18, 2, 183, 181, 3, 2, 2, 2, 184, 187, 3, 2, 2,
	2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187,
	185, 3, 2, 2, 2, 188, 190, 7, 16, 2, 2, 189, 188, 3, 2, 2, 2, 189, 190,
	3, 2, 2, 2, 190, 33, 3, 2, 2, 2, 191, 194, 5, 36, 19, 2, 192, 194, 5, 40,
	21, 2, 193, 191, 3, 2, 2, 2, 193, 192, 3, 2, 2, 2, 194, 35, 3, 2, 2, 2,
	195, 199, 5, 38, 20, 2, 196, 198, 5, 38, 20, 2, 197, 196, 3, 2, 2, 2, 198,
	201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 202,
	3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202, 203, 5, 44, 23, 2, 203, 37, 3, 2,
	2, 2, 204, 205, 7, 31, 2, 2, 205, 206, 7, 21, 2, 2, 206, 39, 3, 2, 2, 2,
	207, 208, 5, 42, 22, 2, 208, 41, 3, 2, 2, 2, 209, 210, 5, 44, 23, 2, 210,
	43, 3, 2, 2, 2, 211, 212, 5, 46, 24, 2, 212, 45, 3, 2, 2, 2, 213, 214,
	5, 48, 25, 2, 214, 47, 3, 2, 2, 2, 215, 217, 5, 54, 28, 2, 216, 218, 5,
	50, 26, 2, 217, 216, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 49, 3, 2, 2,
	2, 219, 226, 5, 52, 27, 2, 220, 222, 7, 20, 2, 2, 221, 220, 3, 2, 2, 2,
	221, 222, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 225, 5, 52, 27, 2, 224,
	221, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227,
	3, 2, 2, 2, 227, 51, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229, 233, 9, 2,
	2, 2, 230, 232, 5, 54, 28, 2, 231, 230, 3, 2, 2, 2, 232, 235, 3, 2, 2,
	2, 233, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 53, 3, 2, 2, 2, 235,
	233, 3, 2, 2, 2, 236, 238, 5, 58, 30, 2, 237, 239, 5, 56, 29, 2, 238, 237,
	3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2,
	2, 2, 241, 55, 3, 2, 2, 2, 242, 243, 5, 92, 47, 2, 243, 244, 5, 58, 30,
	2, 244, 57, 3, 2, 2, 2, 245, 250, 5, 60, 31, 2, 246, 250, 5, 62, 32, 2,
	247, 250, 5, 64, 33, 2, 248, 250, 5, 66, 34, 2, 249, 245, 3, 2, 2, 2, 249,
	246, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 249, 248, 3, 2, 2, 2, 250, 59, 3,
	2, 2, 2, 251, 252, 5, 68, 35, 2, 252, 61, 3, 2, 2, 2, 253, 254, 5, 90,
	46, 2, 254, 63, 3, 2, 2, 2, 255, 256, 5, 24, 13, 2, 256, 65, 3, 2, 2, 2,
	257, 258, 7, 24, 2, 2, 258, 259, 5, 34, 18, 2, 259, 260, 7, 25, 2, 2, 260,
	67, 3, 2, 2, 2, 261, 270, 5, 70, 36, 2, 262, 270, 5, 72, 37, 2, 263, 270,
	5, 82, 42, 2, 264, 270, 5, 74, 38, 2, 265, 270, 5, 78, 40, 2, 266, 270,
	5, 84, 43, 2, 267, 270, 5, 86, 44, 2, 268, 270, 5, 88, 45, 2, 269, 261,
	3, 2, 2, 2, 269, 262, 3, 2, 2, 2, 269, 263, 3, 2, 2, 2, 269, 264, 3, 2,
	2, 2, 269, 265, 3, 2, 2, 2, 269, 266, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2,
	269, 268, 3, 2, 2, 2, 270, 69, 3, 2, 2, 2, 271, 272, 9, 3, 2, 2, 272, 71,
	3, 2, 2, 2, 273, 274, 7, 38, 2, 2, 274, 73, 3, 2, 2, 2, 275, 276, 5, 76,
	39, 2, 276, 75, 3, 2, 2, 2, 277, 278, 7, 37, 2, 2, 278, 279, 7, 31, 2,
	2, 279, 77, 3, 2, 2, 2, 280, 281, 7, 37, 2, 2, 281, 283, 7, 24, 2, 2, 282,
	284, 5, 80, 41, 2, 283, 282, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 283,
	3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 288, 7, 25,
	2, 2, 288, 79, 3, 2, 2, 2, 289, 290, 5, 68, 35, 2, 290, 81, 3, 2, 2, 2,
	291, 292, 7, 39, 2, 2, 292, 83, 3, 2, 2, 2, 293, 294, 7, 8, 2, 2, 294,
	85, 3, 2, 2, 2, 295, 296, 7, 12, 2, 2, 296, 87, 3, 2, 2, 2, 297, 298, 9,
	4, 2, 2, 298, 89, 3, 2, 2, 2, 299, 300, 7, 31, 2, 2, 300, 91, 3, 2, 2,
	2, 301, 306, 5, 94, 48, 2, 302, 303, 5, 94, 48, 2, 303, 304, 5, 94, 48,
	2, 304, 306, 3, 2, 2, 2, 305, 301, 3, 2, 2, 2, 305, 302, 3, 2, 2, 2, 306,
	93, 3, 2, 2, 2, 307, 308, 9, 5, 2, 2, 308, 95, 3, 2, 2, 2, 27, 99, 109,
	113, 116, 126, 135, 139, 151, 159, 167, 175, 178, 185, 189, 193, 199, 217,
	221, 226, 233, 240, 249, 269, 285, 305,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'%'", "'nil'", "'true'", "'false'", "'class'",
	"'self'", "'super'", "", "", "'.'", "'|'", "'^'", "':'", "';'", "':='",
	"'['", "']'", "'('", "')'", "'{'", "'}'", "'<'", "'>'", "'primitive:'",
	"", "", "'16r'", "", "", "", "'#'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "NIL", "TRUE", "FALSE", "CLASS", "SELF", "SUPER",
	"COMMENT", "WHITESPACE", "DOT", "PIPE", "CARROT", "COLON", "SEMICOLON",
	"ASSIGNMENT", "LEFT_BRK", "RIGHT_BRK", "LEFT_PRT", "RIGHT_PRT", "LEFT_BRC",
	"RIGHT_BRC", "LEFT_GLM", "RIGHT_GLM", "PRIMITIVE", "IDENTIFIER", "KEYWORD",
	"HEX", "INT_NUMBER", "HEX_INT_NUMBER", "FLOAT_NUMBER", "SYMBOL_PREFIX",
	"CHAR", "STRING",
}

var ruleNames = []string{
	"script", "classDef", "main", "instanceVars", "instanceVar", "methods",
	"method", "namedMethod", "keywordMethod", "localVars", "localVar", "block",
	"blockArgs", "blockArg", "body", "statements", "statement", "assignmentStatement",
	"assignmentItem", "messageStatements", "messageStatement", "messageExpression",
	"methodExpression", "methodSend", "methodSendRange", "methodSendItem",
	"binaryExpression", "binExprTailItem", "unaryExpression", "primaryLiteral",
	"primaryIdentifier", "primaryBlock", "primaryExpression", "literal", "literalNumber",
	"literalChar", "literalSymbol", "symbol", "literalArray", "literalArrayItem",
	"literalString", "literalNil", "literalSelf", "literalBool", "identifier",
	"binaryOperator", "opchar",
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
	SmallTalkParserT__4           = 5
	SmallTalkParserNIL            = 6
	SmallTalkParserTRUE           = 7
	SmallTalkParserFALSE          = 8
	SmallTalkParserCLASS          = 9
	SmallTalkParserSELF           = 10
	SmallTalkParserSUPER          = 11
	SmallTalkParserCOMMENT        = 12
	SmallTalkParserWHITESPACE     = 13
	SmallTalkParserDOT            = 14
	SmallTalkParserPIPE           = 15
	SmallTalkParserCARROT         = 16
	SmallTalkParserCOLON          = 17
	SmallTalkParserSEMICOLON      = 18
	SmallTalkParserASSIGNMENT     = 19
	SmallTalkParserLEFT_BRK       = 20
	SmallTalkParserRIGHT_BRK      = 21
	SmallTalkParserLEFT_PRT       = 22
	SmallTalkParserRIGHT_PRT      = 23
	SmallTalkParserLEFT_BRC       = 24
	SmallTalkParserRIGHT_BRC      = 25
	SmallTalkParserLEFT_GLM       = 26
	SmallTalkParserRIGHT_GLM      = 27
	SmallTalkParserPRIMITIVE      = 28
	SmallTalkParserIDENTIFIER     = 29
	SmallTalkParserKEYWORD        = 30
	SmallTalkParserHEX            = 31
	SmallTalkParserINT_NUMBER     = 32
	SmallTalkParserHEX_INT_NUMBER = 33
	SmallTalkParserFLOAT_NUMBER   = 34
	SmallTalkParserSYMBOL_PREFIX  = 35
	SmallTalkParserCHAR           = 36
	SmallTalkParserSTRING         = 37
)

// SmallTalkParser rules.
const (
	SmallTalkParserRULE_script              = 0
	SmallTalkParserRULE_classDef            = 1
	SmallTalkParserRULE_main                = 2
	SmallTalkParserRULE_instanceVars        = 3
	SmallTalkParserRULE_instanceVar         = 4
	SmallTalkParserRULE_methods             = 5
	SmallTalkParserRULE_method              = 6
	SmallTalkParserRULE_namedMethod         = 7
	SmallTalkParserRULE_keywordMethod       = 8
	SmallTalkParserRULE_localVars           = 9
	SmallTalkParserRULE_localVar            = 10
	SmallTalkParserRULE_block               = 11
	SmallTalkParserRULE_blockArgs           = 12
	SmallTalkParserRULE_blockArg            = 13
	SmallTalkParserRULE_body                = 14
	SmallTalkParserRULE_statements          = 15
	SmallTalkParserRULE_statement           = 16
	SmallTalkParserRULE_assignmentStatement = 17
	SmallTalkParserRULE_assignmentItem      = 18
	SmallTalkParserRULE_messageStatements   = 19
	SmallTalkParserRULE_messageStatement    = 20
	SmallTalkParserRULE_messageExpression   = 21
	SmallTalkParserRULE_methodExpression    = 22
	SmallTalkParserRULE_methodSend          = 23
	SmallTalkParserRULE_methodSendRange     = 24
	SmallTalkParserRULE_methodSendItem      = 25
	SmallTalkParserRULE_binaryExpression    = 26
	SmallTalkParserRULE_binExprTailItem     = 27
	SmallTalkParserRULE_unaryExpression     = 28
	SmallTalkParserRULE_primaryLiteral      = 29
	SmallTalkParserRULE_primaryIdentifier   = 30
	SmallTalkParserRULE_primaryBlock        = 31
	SmallTalkParserRULE_primaryExpression   = 32
	SmallTalkParserRULE_literal             = 33
	SmallTalkParserRULE_literalNumber       = 34
	SmallTalkParserRULE_literalChar         = 35
	SmallTalkParserRULE_literalSymbol       = 36
	SmallTalkParserRULE_symbol              = 37
	SmallTalkParserRULE_literalArray        = 38
	SmallTalkParserRULE_literalArrayItem    = 39
	SmallTalkParserRULE_literalString       = 40
	SmallTalkParserRULE_literalNil          = 41
	SmallTalkParserRULE_literalSelf         = 42
	SmallTalkParserRULE_literalBool         = 43
	SmallTalkParserRULE_identifier          = 44
	SmallTalkParserRULE_binaryOperator      = 45
	SmallTalkParserRULE_opchar              = 46
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

func (s *ClassDefContext) InstanceVars() IInstanceVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstanceVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstanceVarsContext)
}

func (s *ClassDefContext) Methods() IMethodsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodsContext)
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

	if _la == SmallTalkParserPIPE {
		{
			p.SetState(110)
			p.InstanceVars()
		}

	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SmallTalkParserIDENTIFIER || _la == SmallTalkParserKEYWORD {
		{
			p.SetState(113)
			p.Methods()
		}

	}
	{
		p.SetState(116)
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
		p.SetState(118)
		p.Body()
	}

	return localctx
}

// IInstanceVarsContext is an interface to support dynamic dispatch.
type IInstanceVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstanceVarsContext differentiates from other interfaces.
	IsInstanceVarsContext()
}

type InstanceVarsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstanceVarsContext() *InstanceVarsContext {
	var p = new(InstanceVarsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_instanceVars
	return p
}

func (*InstanceVarsContext) IsInstanceVarsContext() {}

func NewInstanceVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceVarsContext {
	var p = new(InstanceVarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_instanceVars

	return p
}

func (s *InstanceVarsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceVarsContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(SmallTalkParserPIPE)
}

func (s *InstanceVarsContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(SmallTalkParserPIPE, i)
}

func (s *InstanceVarsContext) AllInstanceVar() []IInstanceVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstanceVarContext)(nil)).Elem())
	var tst = make([]IInstanceVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstanceVarContext)
		}
	}

	return tst
}

func (s *InstanceVarsContext) InstanceVar(i int) IInstanceVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstanceVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstanceVarContext)
}

func (s *InstanceVarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceVarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceVarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterInstanceVars(s)
	}
}

func (s *InstanceVarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitInstanceVars(s)
	}
}

func (p *SmallTalkParser) InstanceVars() (localctx IInstanceVarsContext) {
	localctx = NewInstanceVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SmallTalkParserRULE_instanceVars)
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
		p.SetState(120)
		p.Match(SmallTalkParserPIPE)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SmallTalkParserIDENTIFIER {
		{
			p.SetState(121)
			p.InstanceVar()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(SmallTalkParserPIPE)
	}

	return localctx
}

// IInstanceVarContext is an interface to support dynamic dispatch.
type IInstanceVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsInstanceVarContext differentiates from other interfaces.
	IsInstanceVarContext()
}

type InstanceVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyInstanceVarContext() *InstanceVarContext {
	var p = new(InstanceVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_instanceVar
	return p
}

func (*InstanceVarContext) IsInstanceVarContext() {}

func NewInstanceVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceVarContext {
	var p = new(InstanceVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_instanceVar

	return p
}

func (s *InstanceVarContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceVarContext) GetName() antlr.Token { return s.name }

func (s *InstanceVarContext) SetName(v antlr.Token) { s.name = v }

func (s *InstanceVarContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *InstanceVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterInstanceVar(s)
	}
}

func (s *InstanceVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitInstanceVar(s)
	}
}

func (p *SmallTalkParser) InstanceVar() (localctx IInstanceVarContext) {
	localctx = NewInstanceVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SmallTalkParserRULE_instanceVar)

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

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*InstanceVarContext).name = _m
	}

	return localctx
}

// IMethodsContext is an interface to support dynamic dispatch.
type IMethodsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodsContext differentiates from other interfaces.
	IsMethodsContext()
}

type MethodsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodsContext() *MethodsContext {
	var p = new(MethodsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_methods
	return p
}

func (*MethodsContext) IsMethodsContext() {}

func NewMethodsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodsContext {
	var p = new(MethodsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_methods

	return p
}

func (s *MethodsContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodsContext) AllMethod() []IMethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodContext)(nil)).Elem())
	var tst = make([]IMethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodContext)
		}
	}

	return tst
}

func (s *MethodsContext) Method(i int) IMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *MethodsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMethods(s)
	}
}

func (s *MethodsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMethods(s)
	}
}

func (p *SmallTalkParser) Methods() (localctx IMethodsContext) {
	localctx = NewMethodsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SmallTalkParserRULE_methods)
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
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SmallTalkParserIDENTIFIER || _la == SmallTalkParserKEYWORD {
		{
			p.SetState(130)
			p.Method()
		}

		p.SetState(133)
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

func (s *MethodContext) NamedMethod() INamedMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedMethodContext)
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
	p.EnterRule(localctx, 12, SmallTalkParserRULE_method)

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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SmallTalkParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.NamedMethod()
		}

	case SmallTalkParserKEYWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.KeywordMethod()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INamedMethodContext is an interface to support dynamic dispatch.
type INamedMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsNamedMethodContext differentiates from other interfaces.
	IsNamedMethodContext()
}

type NamedMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyNamedMethodContext() *NamedMethodContext {
	var p = new(NamedMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_namedMethod
	return p
}

func (*NamedMethodContext) IsNamedMethodContext() {}

func NewNamedMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedMethodContext {
	var p = new(NamedMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_namedMethod

	return p
}

func (s *NamedMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedMethodContext) GetName() antlr.Token { return s.name }

func (s *NamedMethodContext) SetName(v antlr.Token) { s.name = v }

func (s *NamedMethodContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *NamedMethodContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *NamedMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterNamedMethod(s)
	}
}

func (s *NamedMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitNamedMethod(s)
	}
}

func (p *SmallTalkParser) NamedMethod() (localctx INamedMethodContext) {
	localctx = NewNamedMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SmallTalkParserRULE_namedMethod)

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

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*NamedMethodContext).name = _m
	}
	{
		p.SetState(140)
		p.Block()
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

func (s *KeywordMethodContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
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
	p.EnterRule(localctx, 16, SmallTalkParserRULE_keywordMethod)

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

		var _m = p.Match(SmallTalkParserKEYWORD)

		localctx.(*KeywordMethodContext).name = _m
	}
	{
		p.SetState(143)
		p.Block()
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
	p.EnterRule(localctx, 18, SmallTalkParserRULE_localVars)
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
		p.SetState(145)
		p.Match(SmallTalkParserPIPE)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SmallTalkParserIDENTIFIER {
		{
			p.SetState(146)
			p.LocalVar()
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)
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
	p.EnterRule(localctx, 20, SmallTalkParserRULE_localVar)

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
		p.SetState(153)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*LocalVarContext).name = _m
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
	p.RuleIndex = SmallTalkParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserLEFT_BRK, 0)
}

func (s *BlockContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *BlockContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserRIGHT_BRK, 0)
}

func (s *BlockContext) BlockArgs() IBlockArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockArgsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SmallTalkParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SmallTalkParserRULE_block)
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
		p.Match(SmallTalkParserLEFT_BRK)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SmallTalkParserCOLON {
		{
			p.SetState(156)
			p.BlockArgs()
		}

	}
	{
		p.SetState(159)
		p.Body()
	}
	{
		p.SetState(160)
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
	p.EnterRule(localctx, 24, SmallTalkParserRULE_blockArgs)
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
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SmallTalkParserCOLON {
		{
			p.SetState(162)
			p.BlockArg()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
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
	p.EnterRule(localctx, 26, SmallTalkParserRULE_blockArg)

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
		p.SetState(169)
		p.Match(SmallTalkParserCOLON)
	}
	{
		p.SetState(170)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*BlockArgContext).name = _m
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
	p.EnterRule(localctx, 28, SmallTalkParserRULE_body)
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
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SmallTalkParserPIPE {
		{
			p.SetState(172)
			p.LocalVars()
		}

	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(SmallTalkParserNIL-6))|(1<<(SmallTalkParserTRUE-6))|(1<<(SmallTalkParserFALSE-6))|(1<<(SmallTalkParserSELF-6))|(1<<(SmallTalkParserLEFT_BRK-6))|(1<<(SmallTalkParserLEFT_PRT-6))|(1<<(SmallTalkParserIDENTIFIER-6))|(1<<(SmallTalkParserINT_NUMBER-6))|(1<<(SmallTalkParserHEX_INT_NUMBER-6))|(1<<(SmallTalkParserFLOAT_NUMBER-6))|(1<<(SmallTalkParserSYMBOL_PREFIX-6))|(1<<(SmallTalkParserCHAR-6))|(1<<(SmallTalkParserSTRING-6)))) != 0 {
		{
			p.SetState(175)
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
	p.EnterRule(localctx, 30, SmallTalkParserRULE_statements)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Statement()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(179)
				p.Match(SmallTalkParserDOT)
			}
			{
				p.SetState(180)
				p.Statement()
			}

		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SmallTalkParserDOT {
		{
			p.SetState(186)
			p.Match(SmallTalkParserDOT)
		}

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

func (s *StatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementContext) MessageStatements() IMessageStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageStatementsContext)
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
	p.EnterRule(localctx, 32, SmallTalkParserRULE_statement)

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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.AssignmentStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.MessageStatements()
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
	p.EnterRule(localctx, 34, SmallTalkParserRULE_assignmentStatement)

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
		p.SetState(193)
		p.AssignmentItem()
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(194)
				p.AssignmentItem()
			}

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	{
		p.SetState(200)
		p.MessageExpression()
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
	p.EnterRule(localctx, 36, SmallTalkParserRULE_assignmentItem)

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
		p.SetState(202)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*AssignmentItemContext).name = _m
	}
	{
		p.SetState(203)
		p.Match(SmallTalkParserASSIGNMENT)
	}

	return localctx
}

// IMessageStatementsContext is an interface to support dynamic dispatch.
type IMessageStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageStatementsContext differentiates from other interfaces.
	IsMessageStatementsContext()
}

type MessageStatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageStatementsContext() *MessageStatementsContext {
	var p = new(MessageStatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_messageStatements
	return p
}

func (*MessageStatementsContext) IsMessageStatementsContext() {}

func NewMessageStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageStatementsContext {
	var p = new(MessageStatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_messageStatements

	return p
}

func (s *MessageStatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageStatementsContext) MessageStatement() IMessageStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageStatementContext)
}

func (s *MessageStatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageStatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageStatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMessageStatements(s)
	}
}

func (s *MessageStatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMessageStatements(s)
	}
}

func (p *SmallTalkParser) MessageStatements() (localctx IMessageStatementsContext) {
	localctx = NewMessageStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SmallTalkParserRULE_messageStatements)

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
		p.SetState(205)
		p.MessageStatement()
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
	p.EnterRule(localctx, 40, SmallTalkParserRULE_messageStatement)

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
		p.MessageExpression()
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

func (s *MessageExpressionContext) MethodExpression() IMethodExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodExpressionContext)
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
	localctx = NewMessageExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SmallTalkParserRULE_messageExpression)

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
		p.SetState(209)
		p.MethodExpression()
	}

	return localctx
}

// IMethodExpressionContext is an interface to support dynamic dispatch.
type IMethodExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodExpressionContext differentiates from other interfaces.
	IsMethodExpressionContext()
}

type MethodExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodExpressionContext() *MethodExpressionContext {
	var p = new(MethodExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_methodExpression
	return p
}

func (*MethodExpressionContext) IsMethodExpressionContext() {}

func NewMethodExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodExpressionContext {
	var p = new(MethodExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_methodExpression

	return p
}

func (s *MethodExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodExpressionContext) MethodSend() IMethodSendContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSendContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodSendContext)
}

func (s *MethodExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMethodExpression(s)
	}
}

func (s *MethodExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMethodExpression(s)
	}
}

func (p *SmallTalkParser) MethodExpression() (localctx IMethodExpressionContext) {
	localctx = NewMethodExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SmallTalkParserRULE_methodExpression)

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
		p.MethodSend()
	}

	return localctx
}

// IMethodSendContext is an interface to support dynamic dispatch.
type IMethodSendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRecv returns the recv rule contexts.
	GetRecv() IBinaryExpressionContext

	// SetRecv sets the recv rule contexts.
	SetRecv(IBinaryExpressionContext)

	// IsMethodSendContext differentiates from other interfaces.
	IsMethodSendContext()
}

type MethodSendContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	recv   IBinaryExpressionContext
}

func NewEmptyMethodSendContext() *MethodSendContext {
	var p = new(MethodSendContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_methodSend
	return p
}

func (*MethodSendContext) IsMethodSendContext() {}

func NewMethodSendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodSendContext {
	var p = new(MethodSendContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_methodSend

	return p
}

func (s *MethodSendContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodSendContext) GetRecv() IBinaryExpressionContext { return s.recv }

func (s *MethodSendContext) SetRecv(v IBinaryExpressionContext) { s.recv = v }

func (s *MethodSendContext) BinaryExpression() IBinaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryExpressionContext)
}

func (s *MethodSendContext) MethodSendRange() IMethodSendRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSendRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodSendRangeContext)
}

func (s *MethodSendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodSendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodSendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMethodSend(s)
	}
}

func (s *MethodSendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMethodSend(s)
	}
}

func (p *SmallTalkParser) MethodSend() (localctx IMethodSendContext) {
	localctx = NewMethodSendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SmallTalkParserRULE_methodSend)
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
		p.SetState(213)

		var _x = p.BinaryExpression()

		localctx.(*MethodSendContext).recv = _x
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SmallTalkParserIDENTIFIER || _la == SmallTalkParserKEYWORD {
		{
			p.SetState(214)
			p.MethodSendRange()
		}

	}

	return localctx
}

// IMethodSendRangeContext is an interface to support dynamic dispatch.
type IMethodSendRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodSendRangeContext differentiates from other interfaces.
	IsMethodSendRangeContext()
}

type MethodSendRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodSendRangeContext() *MethodSendRangeContext {
	var p = new(MethodSendRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_methodSendRange
	return p
}

func (*MethodSendRangeContext) IsMethodSendRangeContext() {}

func NewMethodSendRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodSendRangeContext {
	var p = new(MethodSendRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_methodSendRange

	return p
}

func (s *MethodSendRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodSendRangeContext) AllMethodSendItem() []IMethodSendItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodSendItemContext)(nil)).Elem())
	var tst = make([]IMethodSendItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodSendItemContext)
		}
	}

	return tst
}

func (s *MethodSendRangeContext) MethodSendItem(i int) IMethodSendItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSendItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodSendItemContext)
}

func (s *MethodSendRangeContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(SmallTalkParserSEMICOLON)
}

func (s *MethodSendRangeContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(SmallTalkParserSEMICOLON, i)
}

func (s *MethodSendRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodSendRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodSendRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMethodSendRange(s)
	}
}

func (s *MethodSendRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMethodSendRange(s)
	}
}

func (p *SmallTalkParser) MethodSendRange() (localctx IMethodSendRangeContext) {
	localctx = NewMethodSendRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SmallTalkParserRULE_methodSendRange)
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
		p.SetState(217)
		p.MethodSendItem()
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SmallTalkParserSEMICOLON)|(1<<SmallTalkParserIDENTIFIER)|(1<<SmallTalkParserKEYWORD))) != 0 {
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SmallTalkParserSEMICOLON {
			{
				p.SetState(218)
				p.Match(SmallTalkParserSEMICOLON)
			}

		}
		{
			p.SetState(221)
			p.MethodSendItem()
		}

		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodSendItemContext is an interface to support dynamic dispatch.
type IMethodSendItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetArg returns the arg rule contexts.
	GetArg() IBinaryExpressionContext

	// SetArg sets the arg rule contexts.
	SetArg(IBinaryExpressionContext)

	// IsMethodSendItemContext differentiates from other interfaces.
	IsMethodSendItemContext()
}

type MethodSendItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	arg    IBinaryExpressionContext
}

func NewEmptyMethodSendItemContext() *MethodSendItemContext {
	var p = new(MethodSendItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_methodSendItem
	return p
}

func (*MethodSendItemContext) IsMethodSendItemContext() {}

func NewMethodSendItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodSendItemContext {
	var p = new(MethodSendItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_methodSendItem

	return p
}

func (s *MethodSendItemContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodSendItemContext) GetName() antlr.Token { return s.name }

func (s *MethodSendItemContext) SetName(v antlr.Token) { s.name = v }

func (s *MethodSendItemContext) GetArg() IBinaryExpressionContext { return s.arg }

func (s *MethodSendItemContext) SetArg(v IBinaryExpressionContext) { s.arg = v }

func (s *MethodSendItemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *MethodSendItemContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserKEYWORD, 0)
}

func (s *MethodSendItemContext) AllBinaryExpression() []IBinaryExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBinaryExpressionContext)(nil)).Elem())
	var tst = make([]IBinaryExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBinaryExpressionContext)
		}
	}

	return tst
}

func (s *MethodSendItemContext) BinaryExpression(i int) IBinaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBinaryExpressionContext)
}

func (s *MethodSendItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodSendItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodSendItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterMethodSendItem(s)
	}
}

func (s *MethodSendItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitMethodSendItem(s)
	}
}

func (p *SmallTalkParser) MethodSendItem() (localctx IMethodSendItemContext) {
	localctx = NewMethodSendItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SmallTalkParserRULE_methodSendItem)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*MethodSendItemContext).name = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SmallTalkParserIDENTIFIER || _la == SmallTalkParserKEYWORD) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*MethodSendItemContext).name = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(228)

				var _x = p.BinaryExpression()

				localctx.(*MethodSendItemContext).arg = _x
			}

		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IBinaryExpressionContext is an interface to support dynamic dispatch.
type IBinaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryExpressionContext differentiates from other interfaces.
	IsBinaryExpressionContext()
}

type BinaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryExpressionContext() *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_binaryExpression
	return p
}

func (*BinaryExpressionContext) IsBinaryExpressionContext() {}

func NewBinaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_binaryExpression

	return p
}

func (s *BinaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *BinaryExpressionContext) AllBinExprTailItem() []IBinExprTailItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBinExprTailItemContext)(nil)).Elem())
	var tst = make([]IBinExprTailItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBinExprTailItemContext)
		}
	}

	return tst
}

func (s *BinaryExpressionContext) BinExprTailItem(i int) IBinExprTailItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinExprTailItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBinExprTailItemContext)
}

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

func (p *SmallTalkParser) BinaryExpression() (localctx IBinaryExpressionContext) {
	localctx = NewBinaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SmallTalkParserRULE_binaryExpression)
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
		p.SetState(234)
		p.UnaryExpression()
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SmallTalkParserT__0)|(1<<SmallTalkParserT__1)|(1<<SmallTalkParserT__2)|(1<<SmallTalkParserT__3)|(1<<SmallTalkParserT__4))) != 0) {
		{
			p.SetState(235)
			p.BinExprTailItem()
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBinExprTailItemContext is an interface to support dynamic dispatch.
type IBinExprTailItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op rule contexts.
	GetOp() IBinaryOperatorContext

	// SetOp sets the op rule contexts.
	SetOp(IBinaryOperatorContext)

	// IsBinExprTailItemContext differentiates from other interfaces.
	IsBinExprTailItemContext()
}

type BinExprTailItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     IBinaryOperatorContext
}

func NewEmptyBinExprTailItemContext() *BinExprTailItemContext {
	var p = new(BinExprTailItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_binExprTailItem
	return p
}

func (*BinExprTailItemContext) IsBinExprTailItemContext() {}

func NewBinExprTailItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinExprTailItemContext {
	var p = new(BinExprTailItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_binExprTailItem

	return p
}

func (s *BinExprTailItemContext) GetParser() antlr.Parser { return s.parser }

func (s *BinExprTailItemContext) GetOp() IBinaryOperatorContext { return s.op }

func (s *BinExprTailItemContext) SetOp(v IBinaryOperatorContext) { s.op = v }

func (s *BinExprTailItemContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *BinExprTailItemContext) BinaryOperator() IBinaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *BinExprTailItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinExprTailItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinExprTailItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterBinExprTailItem(s)
	}
}

func (s *BinExprTailItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitBinExprTailItem(s)
	}
}

func (p *SmallTalkParser) BinExprTailItem() (localctx IBinExprTailItemContext) {
	localctx = NewBinExprTailItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SmallTalkParserRULE_binExprTailItem)

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

		var _x = p.BinaryOperator()

		localctx.(*BinExprTailItemContext).op = _x
	}
	{
		p.SetState(241)
		p.UnaryExpression()
	}

	return localctx
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_unaryExpression
	return p
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PrimaryLiteral() IPrimaryLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryLiteralContext)
}

func (s *UnaryExpressionContext) PrimaryIdentifier() IPrimaryIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryIdentifierContext)
}

func (s *UnaryExpressionContext) PrimaryBlock() IPrimaryBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryBlockContext)
}

func (s *UnaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (p *SmallTalkParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SmallTalkParserRULE_unaryExpression)

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
	case SmallTalkParserNIL, SmallTalkParserTRUE, SmallTalkParserFALSE, SmallTalkParserSELF, SmallTalkParserINT_NUMBER, SmallTalkParserHEX_INT_NUMBER, SmallTalkParserFLOAT_NUMBER, SmallTalkParserSYMBOL_PREFIX, SmallTalkParserCHAR, SmallTalkParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.PrimaryLiteral()
		}

	case SmallTalkParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.PrimaryIdentifier()
		}

	case SmallTalkParserLEFT_BRK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(245)
			p.PrimaryBlock()
		}

	case SmallTalkParserLEFT_PRT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)
			p.PrimaryExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimaryLiteralContext is an interface to support dynamic dispatch.
type IPrimaryLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryLiteralContext differentiates from other interfaces.
	IsPrimaryLiteralContext()
}

type PrimaryLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryLiteralContext() *PrimaryLiteralContext {
	var p = new(PrimaryLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_primaryLiteral
	return p
}

func (*PrimaryLiteralContext) IsPrimaryLiteralContext() {}

func NewPrimaryLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryLiteralContext {
	var p = new(PrimaryLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_primaryLiteral

	return p
}

func (s *PrimaryLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryLiteralContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterPrimaryLiteral(s)
	}
}

func (s *PrimaryLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitPrimaryLiteral(s)
	}
}

func (p *SmallTalkParser) PrimaryLiteral() (localctx IPrimaryLiteralContext) {
	localctx = NewPrimaryLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SmallTalkParserRULE_primaryLiteral)

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
		p.Literal()
	}

	return localctx
}

// IPrimaryIdentifierContext is an interface to support dynamic dispatch.
type IPrimaryIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryIdentifierContext differentiates from other interfaces.
	IsPrimaryIdentifierContext()
}

type PrimaryIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryIdentifierContext() *PrimaryIdentifierContext {
	var p = new(PrimaryIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_primaryIdentifier
	return p
}

func (*PrimaryIdentifierContext) IsPrimaryIdentifierContext() {}

func NewPrimaryIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryIdentifierContext {
	var p = new(PrimaryIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_primaryIdentifier

	return p
}

func (s *PrimaryIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryIdentifierContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PrimaryIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterPrimaryIdentifier(s)
	}
}

func (s *PrimaryIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitPrimaryIdentifier(s)
	}
}

func (p *SmallTalkParser) PrimaryIdentifier() (localctx IPrimaryIdentifierContext) {
	localctx = NewPrimaryIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SmallTalkParserRULE_primaryIdentifier)

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
		p.Identifier()
	}

	return localctx
}

// IPrimaryBlockContext is an interface to support dynamic dispatch.
type IPrimaryBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryBlockContext differentiates from other interfaces.
	IsPrimaryBlockContext()
}

type PrimaryBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryBlockContext() *PrimaryBlockContext {
	var p = new(PrimaryBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_primaryBlock
	return p
}

func (*PrimaryBlockContext) IsPrimaryBlockContext() {}

func NewPrimaryBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryBlockContext {
	var p = new(PrimaryBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_primaryBlock

	return p
}

func (s *PrimaryBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryBlockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PrimaryBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterPrimaryBlock(s)
	}
}

func (s *PrimaryBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitPrimaryBlock(s)
	}
}

func (p *SmallTalkParser) PrimaryBlock() (localctx IPrimaryBlockContext) {
	localctx = NewPrimaryBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SmallTalkParserRULE_primaryBlock)

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
		p.Block()
	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserLEFT_PRT, 0)
}

func (s *PrimaryExpressionContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *PrimaryExpressionContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserRIGHT_PRT, 0)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (p *SmallTalkParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SmallTalkParserRULE_primaryExpression)

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
		p.Match(SmallTalkParserLEFT_PRT)
	}
	{
		p.SetState(256)
		p.Statement()
	}
	{
		p.SetState(257)
		p.Match(SmallTalkParserRIGHT_PRT)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) LiteralNumber() ILiteralNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralNumberContext)
}

func (s *LiteralContext) LiteralChar() ILiteralCharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralCharContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralCharContext)
}

func (s *LiteralContext) LiteralString() ILiteralStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralStringContext)
}

func (s *LiteralContext) LiteralSymbol() ILiteralSymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralSymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralSymbolContext)
}

func (s *LiteralContext) LiteralArray() ILiteralArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralArrayContext)
}

func (s *LiteralContext) LiteralNil() ILiteralNilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralNilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralNilContext)
}

func (s *LiteralContext) LiteralSelf() ILiteralSelfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralSelfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralSelfContext)
}

func (s *LiteralContext) LiteralBool() ILiteralBoolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralBoolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralBoolContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *SmallTalkParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SmallTalkParserRULE_literal)

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

	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.LiteralNumber()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.LiteralChar()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(261)
			p.LiteralString()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(262)
			p.LiteralSymbol()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(263)
			p.LiteralArray()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(264)
			p.LiteralNil()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(265)
			p.LiteralSelf()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(266)
			p.LiteralBool()
		}

	}

	return localctx
}

// ILiteralNumberContext is an interface to support dynamic dispatch.
type ILiteralNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsLiteralNumberContext differentiates from other interfaces.
	IsLiteralNumberContext()
}

type LiteralNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyLiteralNumberContext() *LiteralNumberContext {
	var p = new(LiteralNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literalNumber
	return p
}

func (*LiteralNumberContext) IsLiteralNumberContext() {}

func NewLiteralNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralNumberContext {
	var p = new(LiteralNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literalNumber

	return p
}

func (s *LiteralNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralNumberContext) GetValue() antlr.Token { return s.value }

func (s *LiteralNumberContext) SetValue(v antlr.Token) { s.value = v }

func (s *LiteralNumberContext) INT_NUMBER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserINT_NUMBER, 0)
}

func (s *LiteralNumberContext) HEX_INT_NUMBER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserHEX_INT_NUMBER, 0)
}

func (s *LiteralNumberContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserFLOAT_NUMBER, 0)
}

func (s *LiteralNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteralNumber(s)
	}
}

func (s *LiteralNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteralNumber(s)
	}
}

func (p *SmallTalkParser) LiteralNumber() (localctx ILiteralNumberContext) {
	localctx = NewLiteralNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SmallTalkParserRULE_literalNumber)
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
		p.SetState(269)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*LiteralNumberContext).value = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(SmallTalkParserINT_NUMBER-32))|(1<<(SmallTalkParserHEX_INT_NUMBER-32))|(1<<(SmallTalkParserFLOAT_NUMBER-32)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*LiteralNumberContext).value = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralCharContext is an interface to support dynamic dispatch.
type ILiteralCharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsLiteralCharContext differentiates from other interfaces.
	IsLiteralCharContext()
}

type LiteralCharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyLiteralCharContext() *LiteralCharContext {
	var p = new(LiteralCharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literalChar
	return p
}

func (*LiteralCharContext) IsLiteralCharContext() {}

func NewLiteralCharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralCharContext {
	var p = new(LiteralCharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literalChar

	return p
}

func (s *LiteralCharContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralCharContext) GetValue() antlr.Token { return s.value }

func (s *LiteralCharContext) SetValue(v antlr.Token) { s.value = v }

func (s *LiteralCharContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserCHAR, 0)
}

func (s *LiteralCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralCharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteralChar(s)
	}
}

func (s *LiteralCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteralChar(s)
	}
}

func (p *SmallTalkParser) LiteralChar() (localctx ILiteralCharContext) {
	localctx = NewLiteralCharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SmallTalkParserRULE_literalChar)

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
		p.SetState(271)

		var _m = p.Match(SmallTalkParserCHAR)

		localctx.(*LiteralCharContext).value = _m
	}

	return localctx
}

// ILiteralSymbolContext is an interface to support dynamic dispatch.
type ILiteralSymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralSymbolContext differentiates from other interfaces.
	IsLiteralSymbolContext()
}

type LiteralSymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralSymbolContext() *LiteralSymbolContext {
	var p = new(LiteralSymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literalSymbol
	return p
}

func (*LiteralSymbolContext) IsLiteralSymbolContext() {}

func NewLiteralSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralSymbolContext {
	var p = new(LiteralSymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literalSymbol

	return p
}

func (s *LiteralSymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralSymbolContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *LiteralSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralSymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteralSymbol(s)
	}
}

func (s *LiteralSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteralSymbol(s)
	}
}

func (p *SmallTalkParser) LiteralSymbol() (localctx ILiteralSymbolContext) {
	localctx = NewLiteralSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SmallTalkParserRULE_literalSymbol)

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
		p.SetState(273)
		p.Symbol()
	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) GetName() antlr.Token { return s.name }

func (s *SymbolContext) SetName(v antlr.Token) { s.name = v }

func (s *SymbolContext) SYMBOL_PREFIX() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserSYMBOL_PREFIX, 0)
}

func (s *SymbolContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserIDENTIFIER, 0)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *SmallTalkParser) Symbol() (localctx ISymbolContext) {
	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SmallTalkParserRULE_symbol)

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
		p.Match(SmallTalkParserSYMBOL_PREFIX)
	}
	{
		p.SetState(276)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*SymbolContext).name = _m
	}

	return localctx
}

// ILiteralArrayContext is an interface to support dynamic dispatch.
type ILiteralArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralArrayContext differentiates from other interfaces.
	IsLiteralArrayContext()
}

type LiteralArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralArrayContext() *LiteralArrayContext {
	var p = new(LiteralArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literalArray
	return p
}

func (*LiteralArrayContext) IsLiteralArrayContext() {}

func NewLiteralArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralArrayContext {
	var p = new(LiteralArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literalArray

	return p
}

func (s *LiteralArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralArrayContext) SYMBOL_PREFIX() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserSYMBOL_PREFIX, 0)
}

func (s *LiteralArrayContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserLEFT_PRT, 0)
}

func (s *LiteralArrayContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserRIGHT_PRT, 0)
}

func (s *LiteralArrayContext) AllLiteralArrayItem() []ILiteralArrayItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralArrayItemContext)(nil)).Elem())
	var tst = make([]ILiteralArrayItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralArrayItemContext)
		}
	}

	return tst
}

func (s *LiteralArrayContext) LiteralArrayItem(i int) ILiteralArrayItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralArrayItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralArrayItemContext)
}

func (s *LiteralArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteralArray(s)
	}
}

func (s *LiteralArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteralArray(s)
	}
}

func (p *SmallTalkParser) LiteralArray() (localctx ILiteralArrayContext) {
	localctx = NewLiteralArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SmallTalkParserRULE_literalArray)
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
		p.SetState(278)
		p.Match(SmallTalkParserSYMBOL_PREFIX)
	}
	{
		p.SetState(279)
		p.Match(SmallTalkParserLEFT_PRT)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(SmallTalkParserNIL-6))|(1<<(SmallTalkParserTRUE-6))|(1<<(SmallTalkParserFALSE-6))|(1<<(SmallTalkParserSELF-6))|(1<<(SmallTalkParserINT_NUMBER-6))|(1<<(SmallTalkParserHEX_INT_NUMBER-6))|(1<<(SmallTalkParserFLOAT_NUMBER-6))|(1<<(SmallTalkParserSYMBOL_PREFIX-6))|(1<<(SmallTalkParserCHAR-6))|(1<<(SmallTalkParserSTRING-6)))) != 0) {
		{
			p.SetState(280)
			p.LiteralArrayItem()
		}

		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(285)
		p.Match(SmallTalkParserRIGHT_PRT)
	}

	return localctx
}

// ILiteralArrayItemContext is an interface to support dynamic dispatch.
type ILiteralArrayItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value rule contexts.
	GetValue() ILiteralContext

	// SetValue sets the value rule contexts.
	SetValue(ILiteralContext)

	// IsLiteralArrayItemContext differentiates from other interfaces.
	IsLiteralArrayItemContext()
}

type LiteralArrayItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  ILiteralContext
}

func NewEmptyLiteralArrayItemContext() *LiteralArrayItemContext {
	var p = new(LiteralArrayItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literalArrayItem
	return p
}

func (*LiteralArrayItemContext) IsLiteralArrayItemContext() {}

func NewLiteralArrayItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralArrayItemContext {
	var p = new(LiteralArrayItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literalArrayItem

	return p
}

func (s *LiteralArrayItemContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralArrayItemContext) GetValue() ILiteralContext { return s.value }

func (s *LiteralArrayItemContext) SetValue(v ILiteralContext) { s.value = v }

func (s *LiteralArrayItemContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralArrayItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralArrayItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralArrayItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteralArrayItem(s)
	}
}

func (s *LiteralArrayItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteralArrayItem(s)
	}
}

func (p *SmallTalkParser) LiteralArrayItem() (localctx ILiteralArrayItemContext) {
	localctx = NewLiteralArrayItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SmallTalkParserRULE_literalArrayItem)

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

		var _x = p.Literal()

		localctx.(*LiteralArrayItemContext).value = _x
	}

	return localctx
}

// ILiteralStringContext is an interface to support dynamic dispatch.
type ILiteralStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsLiteralStringContext differentiates from other interfaces.
	IsLiteralStringContext()
}

type LiteralStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyLiteralStringContext() *LiteralStringContext {
	var p = new(LiteralStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literalString
	return p
}

func (*LiteralStringContext) IsLiteralStringContext() {}

func NewLiteralStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralStringContext {
	var p = new(LiteralStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literalString

	return p
}

func (s *LiteralStringContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralStringContext) GetValue() antlr.Token { return s.value }

func (s *LiteralStringContext) SetValue(v antlr.Token) { s.value = v }

func (s *LiteralStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserSTRING, 0)
}

func (s *LiteralStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteralString(s)
	}
}

func (s *LiteralStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteralString(s)
	}
}

func (p *SmallTalkParser) LiteralString() (localctx ILiteralStringContext) {
	localctx = NewLiteralStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SmallTalkParserRULE_literalString)

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
		p.SetState(289)

		var _m = p.Match(SmallTalkParserSTRING)

		localctx.(*LiteralStringContext).value = _m
	}

	return localctx
}

// ILiteralNilContext is an interface to support dynamic dispatch.
type ILiteralNilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralNilContext differentiates from other interfaces.
	IsLiteralNilContext()
}

type LiteralNilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralNilContext() *LiteralNilContext {
	var p = new(LiteralNilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literalNil
	return p
}

func (*LiteralNilContext) IsLiteralNilContext() {}

func NewLiteralNilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralNilContext {
	var p = new(LiteralNilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literalNil

	return p
}

func (s *LiteralNilContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralNilContext) NIL() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserNIL, 0)
}

func (s *LiteralNilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralNilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralNilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteralNil(s)
	}
}

func (s *LiteralNilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteralNil(s)
	}
}

func (p *SmallTalkParser) LiteralNil() (localctx ILiteralNilContext) {
	localctx = NewLiteralNilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SmallTalkParserRULE_literalNil)

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
		p.SetState(291)
		p.Match(SmallTalkParserNIL)
	}

	return localctx
}

// ILiteralSelfContext is an interface to support dynamic dispatch.
type ILiteralSelfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralSelfContext differentiates from other interfaces.
	IsLiteralSelfContext()
}

type LiteralSelfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralSelfContext() *LiteralSelfContext {
	var p = new(LiteralSelfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literalSelf
	return p
}

func (*LiteralSelfContext) IsLiteralSelfContext() {}

func NewLiteralSelfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralSelfContext {
	var p = new(LiteralSelfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literalSelf

	return p
}

func (s *LiteralSelfContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralSelfContext) SELF() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserSELF, 0)
}

func (s *LiteralSelfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralSelfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralSelfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteralSelf(s)
	}
}

func (s *LiteralSelfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteralSelf(s)
	}
}

func (p *SmallTalkParser) LiteralSelf() (localctx ILiteralSelfContext) {
	localctx = NewLiteralSelfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SmallTalkParserRULE_literalSelf)

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
		p.SetState(293)
		p.Match(SmallTalkParserSELF)
	}

	return localctx
}

// ILiteralBoolContext is an interface to support dynamic dispatch.
type ILiteralBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsLiteralBoolContext differentiates from other interfaces.
	IsLiteralBoolContext()
}

type LiteralBoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyLiteralBoolContext() *LiteralBoolContext {
	var p = new(LiteralBoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_literalBool
	return p
}

func (*LiteralBoolContext) IsLiteralBoolContext() {}

func NewLiteralBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralBoolContext {
	var p = new(LiteralBoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_literalBool

	return p
}

func (s *LiteralBoolContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralBoolContext) GetValue() antlr.Token { return s.value }

func (s *LiteralBoolContext) SetValue(v antlr.Token) { s.value = v }

func (s *LiteralBoolContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserTRUE, 0)
}

func (s *LiteralBoolContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SmallTalkParserFALSE, 0)
}

func (s *LiteralBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralBoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterLiteralBool(s)
	}
}

func (s *LiteralBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitLiteralBool(s)
	}
}

func (p *SmallTalkParser) LiteralBool() (localctx ILiteralBoolContext) {
	localctx = NewLiteralBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SmallTalkParserRULE_literalBool)
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
		p.SetState(295)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*LiteralBoolContext).value = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SmallTalkParserTRUE || _la == SmallTalkParserFALSE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*LiteralBoolContext).value = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
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

func (s *IdentifierContext) GetName() antlr.Token { return s.name }

func (s *IdentifierContext) SetName(v antlr.Token) { s.name = v }

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
	p.EnterRule(localctx, 88, SmallTalkParserRULE_identifier)

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
		p.SetState(297)

		var _m = p.Match(SmallTalkParserIDENTIFIER)

		localctx.(*IdentifierContext).name = _m
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
	p.RuleIndex = SmallTalkParserRULE_binaryOperator
	return p
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperatorContext) AllOpchar() []IOpcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOpcharContext)(nil)).Elem())
	var tst = make([]IOpcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOpcharContext)
		}
	}

	return tst
}

func (s *BinaryOperatorContext) Opchar(i int) IOpcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOpcharContext)
}

func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterBinaryOperator(s)
	}
}

func (s *BinaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitBinaryOperator(s)
	}
}

func (p *SmallTalkParser) BinaryOperator() (localctx IBinaryOperatorContext) {
	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SmallTalkParserRULE_binaryOperator)

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

	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(299)
			p.Opchar()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.Opchar()
		}
		{
			p.SetState(301)
			p.Opchar()
		}

	}

	return localctx
}

// IOpcharContext is an interface to support dynamic dispatch.
type IOpcharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpcharContext differentiates from other interfaces.
	IsOpcharContext()
}

type OpcharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcharContext() *OpcharContext {
	var p = new(OpcharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmallTalkParserRULE_opchar
	return p
}

func (*OpcharContext) IsOpcharContext() {}

func NewOpcharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcharContext {
	var p = new(OpcharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmallTalkParserRULE_opchar

	return p
}

func (s *OpcharContext) GetParser() antlr.Parser { return s.parser }
func (s *OpcharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.EnterOpchar(s)
	}
}

func (s *OpcharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmallTalkListener); ok {
		listenerT.ExitOpchar(s)
	}
}

func (p *SmallTalkParser) Opchar() (localctx IOpcharContext) {
	localctx = NewOpcharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SmallTalkParserRULE_opchar)
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
		p.SetState(305)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SmallTalkParserT__0)|(1<<SmallTalkParserT__1)|(1<<SmallTalkParserT__2)|(1<<SmallTalkParserT__3)|(1<<SmallTalkParserT__4))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
