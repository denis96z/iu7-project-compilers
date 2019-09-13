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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 76, 357,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 3, 2, 3, 2, 7, 2, 83, 10, 2, 12, 2, 14, 2, 86, 11, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 100, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 120, 10, 7,
	3, 8, 3, 8, 5, 8, 124, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 149, 10, 14, 12, 14, 14,
	14, 152, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	7, 16, 162, 10, 16, 12, 16, 14, 16, 165, 11, 16, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18,
	180, 10, 18, 12, 18, 14, 18, 183, 11, 18, 5, 18, 185, 10, 18, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 7, 20, 200, 10, 20, 12, 20, 14, 20, 203, 11, 20, 5, 20, 205, 10,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22,
	7, 22, 217, 10, 22, 12, 22, 14, 22, 220, 11, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 234,
	10, 23, 3, 24, 3, 24, 7, 24, 238, 10, 24, 12, 24, 14, 24, 241, 11, 24,
	3, 24, 5, 24, 244, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 5, 28, 265, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 292,
	10, 34, 12, 34, 14, 34, 295, 11, 34, 5, 34, 297, 10, 34, 3, 34, 3, 34,
	3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	7, 36, 322, 10, 36, 12, 36, 14, 36, 325, 11, 36, 5, 36, 327, 10, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 5, 36, 333, 10, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	7, 36, 339, 10, 36, 12, 36, 14, 36, 342, 11, 36, 3, 37, 3, 37, 3, 38, 3,
	38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 355, 10, 40,
	3, 40, 2, 3, 70, 41, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
	66, 68, 70, 72, 74, 76, 78, 2, 8, 4, 2, 20, 21, 72, 73, 3, 2, 22, 33, 3,
	2, 72, 73, 3, 2, 20, 21, 4, 2, 44, 44, 48, 48, 5, 2, 43, 47, 50, 55, 57,
	71, 2, 357, 2, 80, 3, 2, 2, 2, 4, 89, 3, 2, 2, 2, 6, 99, 3, 2, 2, 2, 8,
	101, 3, 2, 2, 2, 10, 109, 3, 2, 2, 2, 12, 119, 3, 2, 2, 2, 14, 123, 3,
	2, 2, 2, 16, 125, 3, 2, 2, 2, 18, 127, 3, 2, 2, 2, 20, 129, 3, 2, 2, 2,
	22, 132, 3, 2, 2, 2, 24, 138, 3, 2, 2, 2, 26, 142, 3, 2, 2, 2, 28, 155,
	3, 2, 2, 2, 30, 157, 3, 2, 2, 2, 32, 168, 3, 2, 2, 2, 34, 173, 3, 2, 2,
	2, 36, 189, 3, 2, 2, 2, 38, 193, 3, 2, 2, 2, 40, 211, 3, 2, 2, 2, 42, 218,
	3, 2, 2, 2, 44, 233, 3, 2, 2, 2, 46, 235, 3, 2, 2, 2, 48, 245, 3, 2, 2,
	2, 50, 251, 3, 2, 2, 2, 52, 257, 3, 2, 2, 2, 54, 264, 3, 2, 2, 2, 56, 266,
	3, 2, 2, 2, 58, 271, 3, 2, 2, 2, 60, 277, 3, 2, 2, 2, 62, 279, 3, 2, 2,
	2, 64, 281, 3, 2, 2, 2, 66, 286, 3, 2, 2, 2, 68, 301, 3, 2, 2, 2, 70, 332,
	3, 2, 2, 2, 72, 343, 3, 2, 2, 2, 74, 345, 3, 2, 2, 2, 76, 347, 3, 2, 2,
	2, 78, 354, 3, 2, 2, 2, 80, 84, 5, 4, 3, 2, 81, 83, 5, 6, 4, 2, 82, 81,
	3, 2, 2, 2, 83, 86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2,
	85, 87, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 87, 88, 7, 2, 2, 3, 88, 3, 3, 2,
	2, 2, 89, 90, 7, 3, 2, 2, 90, 91, 7, 75, 2, 2, 91, 92, 7, 42, 2, 2, 92,
	5, 3, 2, 2, 2, 93, 100, 5, 8, 5, 2, 94, 100, 5, 10, 6, 2, 95, 100, 5, 26,
	14, 2, 96, 100, 5, 30, 16, 2, 97, 100, 5, 34, 18, 2, 98, 100, 5, 38, 20,
	2, 99, 93, 3, 2, 2, 2, 99, 94, 3, 2, 2, 2, 99, 95, 3, 2, 2, 2, 99, 96,
	3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 98, 3, 2, 2, 2, 100, 7, 3, 2, 2, 2,
	101, 102, 7, 15, 2, 2, 102, 103, 7, 72, 2, 2, 103, 104, 7, 40, 2, 2, 104,
	105, 5, 16, 9, 2, 105, 106, 7, 61, 2, 2, 106, 107, 9, 2, 2, 2, 107, 108,
	7, 42, 2, 2, 108, 9, 3, 2, 2, 2, 109, 110, 7, 4, 2, 2, 110, 111, 7, 74,
	2, 2, 111, 112, 7, 61, 2, 2, 112, 113, 5, 12, 7, 2, 113, 114, 7, 42, 2,
	2, 114, 11, 3, 2, 2, 2, 115, 120, 5, 14, 8, 2, 116, 120, 5, 20, 11, 2,
	117, 120, 5, 22, 12, 2, 118, 120, 5, 24, 13, 2, 119, 115, 3, 2, 2, 2, 119,
	116, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 13, 3,
	2, 2, 2, 121, 124, 5, 16, 9, 2, 122, 124, 5, 18, 10, 2, 123, 121, 3, 2,
	2, 2, 123, 122, 3, 2, 2, 2, 124, 15, 3, 2, 2, 2, 125, 126, 9, 3, 2, 2,
	126, 17, 3, 2, 2, 2, 127, 128, 7, 74, 2, 2, 128, 19, 3, 2, 2, 2, 129, 130,
	7, 49, 2, 2, 130, 131, 5, 14, 8, 2, 131, 21, 3, 2, 2, 2, 132, 133, 7, 38,
	2, 2, 133, 134, 5, 14, 8, 2, 134, 135, 7, 40, 2, 2, 135, 136, 9, 4, 2,
	2, 136, 137, 7, 39, 2, 2, 137, 23, 3, 2, 2, 2, 138, 139, 7, 38, 2, 2, 139,
	140, 5, 14, 8, 2, 140, 141, 7, 39, 2, 2, 141, 25, 3, 2, 2, 2, 142, 143,
	7, 18, 2, 2, 143, 144, 7, 74, 2, 2, 144, 145, 7, 36, 2, 2, 145, 150, 5,
	28, 15, 2, 146, 147, 7, 41, 2, 2, 147, 149, 5, 28, 15, 2, 148, 146, 3,
	2, 2, 2, 149, 152, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2,
	2, 151, 153, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 153, 154, 7, 37, 2, 2, 154,
	27, 3, 2, 2, 2, 155, 156, 7, 72, 2, 2, 156, 29, 3, 2, 2, 2, 157, 158, 7,
	17, 2, 2, 158, 159, 7, 74, 2, 2, 159, 163, 7, 36, 2, 2, 160, 162, 5, 32,
	17, 2, 161, 160, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2,
	163, 164, 3, 2, 2, 2, 164, 166, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166,
	167, 7, 37, 2, 2, 167, 31, 3, 2, 2, 2, 168, 169, 7, 75, 2, 2, 169, 170,
	7, 40, 2, 2, 170, 171, 5, 12, 7, 2, 171, 172, 7, 42, 2, 2, 172, 33, 3,
	2, 2, 2, 173, 174, 7, 13, 2, 2, 174, 175, 7, 75, 2, 2, 175, 184, 7, 34,
	2, 2, 176, 181, 5, 36, 19, 2, 177, 178, 7, 41, 2, 2, 178, 180, 5, 36, 19,
	2, 179, 177, 3, 2, 2, 2, 180, 183, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181,
	182, 3, 2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184, 176,
	3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 7, 35,
	2, 2, 187, 188, 5, 42, 22, 2, 188, 35, 3, 2, 2, 2, 189, 190, 7, 75, 2,
	2, 190, 191, 7, 40, 2, 2, 191, 192, 5, 12, 7, 2, 192, 37, 3, 2, 2, 2, 193,
	194, 7, 14, 2, 2, 194, 195, 7, 75, 2, 2, 195, 204, 7, 34, 2, 2, 196, 201,
	5, 40, 21, 2, 197, 198, 7, 41, 2, 2, 198, 200, 5, 40, 21, 2, 199, 197,
	3, 2, 2, 2, 200, 203, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2,
	2, 2, 202, 205, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 204, 196, 3, 2, 2, 2,
	204, 205, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 7, 40, 2, 2, 207,
	208, 5, 12, 7, 2, 208, 209, 7, 35, 2, 2, 209, 210, 5, 42, 22, 2, 210, 39,
	3, 2, 2, 2, 211, 212, 7, 75, 2, 2, 212, 213, 7, 40, 2, 2, 213, 214, 5,
	12, 7, 2, 214, 41, 3, 2, 2, 2, 215, 217, 5, 44, 23, 2, 216, 215, 3, 2,
	2, 2, 217, 220, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2,
	219, 43, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 234, 5, 46, 24, 2, 222,
	234, 5, 54, 28, 2, 223, 234, 5, 60, 31, 2, 224, 234, 5, 62, 32, 2, 225,
	234, 5, 66, 34, 2, 226, 227, 5, 64, 33, 2, 227, 228, 7, 42, 2, 2, 228,
	234, 3, 2, 2, 2, 229, 230, 5, 70, 36, 2, 230, 231, 7, 42, 2, 2, 231, 234,
	3, 2, 2, 2, 232, 234, 5, 78, 40, 2, 233, 221, 3, 2, 2, 2, 233, 222, 3,
	2, 2, 2, 233, 223, 3, 2, 2, 2, 233, 224, 3, 2, 2, 2, 233, 225, 3, 2, 2,
	2, 233, 226, 3, 2, 2, 2, 233, 229, 3, 2, 2, 2, 233, 232, 3, 2, 2, 2, 234,
	45, 3, 2, 2, 2, 235, 239, 5, 48, 25, 2, 236, 238, 5, 50, 26, 2, 237, 236,
	3, 2, 2, 2, 238, 241, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239, 240, 3, 2,
	2, 2, 240, 243, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 242, 244, 5, 52, 27,
	2, 243, 242, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 47, 3, 2, 2, 2, 245,
	246, 7, 10, 2, 2, 246, 247, 5, 70, 36, 2, 247, 248, 7, 36, 2, 2, 248, 249,
	5, 42, 22, 2, 249, 250, 7, 37, 2, 2, 250, 49, 3, 2, 2, 2, 251, 252, 7,
	11, 2, 2, 252, 253, 5, 70, 36, 2, 253, 254, 7, 36, 2, 2, 254, 255, 5, 42,
	22, 2, 255, 256, 7, 37, 2, 2, 256, 51, 3, 2, 2, 2, 257, 258, 7, 12, 2,
	2, 258, 259, 7, 36, 2, 2, 259, 260, 5, 42, 22, 2, 260, 261, 7, 37, 2, 2,
	261, 53, 3, 2, 2, 2, 262, 265, 5, 56, 29, 2, 263, 265, 5, 58, 30, 2, 264,
	262, 3, 2, 2, 2, 264, 263, 3, 2, 2, 2, 265, 55, 3, 2, 2, 2, 266, 267, 7,
	7, 2, 2, 267, 268, 7, 36, 2, 2, 268, 269, 5, 42, 22, 2, 269, 270, 7, 37,
	2, 2, 270, 57, 3, 2, 2, 2, 271, 272, 7, 6, 2, 2, 272, 273, 5, 70, 36, 2,
	273, 274, 7, 36, 2, 2, 274, 275, 5, 42, 22, 2, 275, 276, 7, 37, 2, 2, 276,
	59, 3, 2, 2, 2, 277, 278, 7, 8, 2, 2, 278, 61, 3, 2, 2, 2, 279, 280, 7,
	9, 2, 2, 280, 63, 3, 2, 2, 2, 281, 282, 7, 16, 2, 2, 282, 283, 7, 75, 2,
	2, 283, 284, 7, 61, 2, 2, 284, 285, 5, 70, 36, 2, 285, 65, 3, 2, 2, 2,
	286, 287, 7, 75, 2, 2, 287, 296, 7, 34, 2, 2, 288, 293, 5, 68, 35, 2, 289,
	290, 7, 41, 2, 2, 290, 292, 5, 68, 35, 2, 291, 289, 3, 2, 2, 2, 292, 295,
	3, 2, 2, 2, 293, 291, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 297, 3, 2,
	2, 2, 295, 293, 3, 2, 2, 2, 296, 288, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2,
	297, 298, 3, 2, 2, 2, 298, 299, 7, 35, 2, 2, 299, 300, 7, 42, 2, 2, 300,
	67, 3, 2, 2, 2, 301, 302, 5, 70, 36, 2, 302, 69, 3, 2, 2, 2, 303, 304,
	8, 36, 1, 2, 304, 333, 7, 75, 2, 2, 305, 333, 7, 73, 2, 2, 306, 333, 9,
	5, 2, 2, 307, 333, 7, 72, 2, 2, 308, 309, 7, 38, 2, 2, 309, 310, 5, 70,
	36, 2, 310, 311, 7, 39, 2, 2, 311, 333, 3, 2, 2, 2, 312, 313, 7, 34, 2,
	2, 313, 314, 5, 70, 36, 2, 314, 315, 7, 35, 2, 2, 315, 333, 3, 2, 2, 2,
	316, 317, 7, 75, 2, 2, 317, 326, 7, 34, 2, 2, 318, 323, 5, 72, 37, 2, 319,
	320, 7, 41, 2, 2, 320, 322, 5, 72, 37, 2, 321, 319, 3, 2, 2, 2, 322, 325,
	3, 2, 2, 2, 323, 321, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 327, 3, 2,
	2, 2, 325, 323, 3, 2, 2, 2, 326, 318, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2,
	327, 328, 3, 2, 2, 2, 328, 333, 7, 35, 2, 2, 329, 330, 5, 74, 38, 2, 330,
	331, 5, 70, 36, 3, 331, 333, 3, 2, 2, 2, 332, 303, 3, 2, 2, 2, 332, 305,
	3, 2, 2, 2, 332, 306, 3, 2, 2, 2, 332, 307, 3, 2, 2, 2, 332, 308, 3, 2,
	2, 2, 332, 312, 3, 2, 2, 2, 332, 316, 3, 2, 2, 2, 332, 329, 3, 2, 2, 2,
	333, 340, 3, 2, 2, 2, 334, 335, 12, 4, 2, 2, 335, 336, 5, 76, 39, 2, 336,
	337, 5, 70, 36, 5, 337, 339, 3, 2, 2, 2, 338, 334, 3, 2, 2, 2, 339, 342,
	3, 2, 2, 2, 340, 338, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 71, 3, 2,
	2, 2, 342, 340, 3, 2, 2, 2, 343, 344, 5, 70, 36, 2, 344, 73, 3, 2, 2, 2,
	345, 346, 9, 6, 2, 2, 346, 75, 3, 2, 2, 2, 347, 348, 9, 7, 2, 2, 348, 77,
	3, 2, 2, 2, 349, 355, 7, 19, 2, 2, 350, 351, 7, 19, 2, 2, 351, 352, 5,
	70, 36, 2, 352, 353, 7, 42, 2, 2, 353, 355, 3, 2, 2, 2, 354, 349, 3, 2,
	2, 2, 354, 350, 3, 2, 2, 2, 355, 79, 3, 2, 2, 2, 24, 84, 99, 119, 123,
	150, 163, 181, 184, 201, 204, 218, 233, 239, 243, 264, 293, 296, 323, 326,
	332, 340, 354,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'pkg'", "'type'", "'for'", "'while'", "'loop'", "'break'", "'continue'",
	"'if'", "'elif'", "'else'", "'proc'", "'func'", "'const'", "'let'", "'struct'",
	"'enum'", "'return'", "'true'", "'false'", "'i8_t'", "'u8_t'", "'i16_t'",
	"'u16_t'", "'i32_t'", "'u32_t'", "'i64_t'", "'u64_t'", "'char_t'", "'bool_t'",
	"'size_t'", "'ssize_t'", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'",
	"','", "';'", "'+'", "'-'", "'*'", "'/'", "'%'", "'~'", "'@'", "'<<'",
	"'>>'", "'&'", "'|'", "'^'", "'=='", "'!='", "'>'", "'>='", "'<'", "'<='",
	"'='", "'+='", "'-='", "'*='", "'/='", "'%='", "'<<='", "'>>='", "'&='",
	"'|='", "'^='",
}
var symbolicNames = []string{
	"", "KW_PKG", "KW_TYPE", "KW_FOR", "KW_WHILE", "KW_LOOP", "KW_BREAK", "KW_CONTINUE",
	"KW_IF", "KW_ELIF", "KW_ELSE", "KW_PROC", "KW_FUNC", "KW_CONST", "KW_LET",
	"KW_STRUCT", "KW_ENUM", "KW_RETURN", "KW_TRUE", "KW_FALSE", "I8_T", "U8_T",
	"I16_T", "U16_T", "I32_T", "U32_T", "I64_T", "U64_T", "CHAR_T", "BOOL_T",
	"SIZE_T", "SSIZE_T", "LEFT_PRT", "RIGHT_PRT", "LEFT_BRC", "RIGHT_BRC",
	"LEFT_BRK", "RIGHT_BRK", "COLON", "COMMA", "SEMICOLON", "ADD", "SUB", "MUL",
	"DIV", "MOD", "NOT", "REF", "SHL", "SHR", "AND", "OR", "XOR", "EQL", "NOT_EQL",
	"GRT", "GRT_EQL", "LESS", "LESS_EQL", "ASGN", "ADD_ASGN", "SUB_ASGN", "MUL_ASGN",
	"DIV_ASGN", "MOD_ASGN", "SHL_ASGN", "SHR_ASGN", "AND_ASGN", "OR_ASGN",
	"XOR_ASGN", "CONST", "INT_VALUE", "TYPE", "IDENTIFIER", "WHITESPACE",
}

var ruleNames = []string{
	"start", "pkg", "decl", "constDecl", "typeDecl", "typeSpec", "simpleTypeSpec",
	"basicTypeSpec", "namedTypeSpec", "refTypeSpec", "arrayTypeSpec", "sliceTypeSpec",
	"enumDecl", "enumOption", "structDecl", "structAttr", "procDecl", "procArg",
	"funcDecl", "funcArg", "block", "statement", "condition", "ifConditionBranch",
	"elifConditionBranch", "elseConditionBranch", "loop", "trueLoop", "whileLoop",
	"breakStatement", "continueStatement", "varDecl", "procCall", "procParam",
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
	DZParserKW_TRUE     = 18
	DZParserKW_FALSE    = 19
	DZParserI8_T        = 20
	DZParserU8_T        = 21
	DZParserI16_T       = 22
	DZParserU16_T       = 23
	DZParserI32_T       = 24
	DZParserU32_T       = 25
	DZParserI64_T       = 26
	DZParserU64_T       = 27
	DZParserCHAR_T      = 28
	DZParserBOOL_T      = 29
	DZParserSIZE_T      = 30
	DZParserSSIZE_T     = 31
	DZParserLEFT_PRT    = 32
	DZParserRIGHT_PRT   = 33
	DZParserLEFT_BRC    = 34
	DZParserRIGHT_BRC   = 35
	DZParserLEFT_BRK    = 36
	DZParserRIGHT_BRK   = 37
	DZParserCOLON       = 38
	DZParserCOMMA       = 39
	DZParserSEMICOLON   = 40
	DZParserADD         = 41
	DZParserSUB         = 42
	DZParserMUL         = 43
	DZParserDIV         = 44
	DZParserMOD         = 45
	DZParserNOT         = 46
	DZParserREF         = 47
	DZParserSHL         = 48
	DZParserSHR         = 49
	DZParserAND         = 50
	DZParserOR          = 51
	DZParserXOR         = 52
	DZParserEQL         = 53
	DZParserNOT_EQL     = 54
	DZParserGRT         = 55
	DZParserGRT_EQL     = 56
	DZParserLESS        = 57
	DZParserLESS_EQL    = 58
	DZParserASGN        = 59
	DZParserADD_ASGN    = 60
	DZParserSUB_ASGN    = 61
	DZParserMUL_ASGN    = 62
	DZParserDIV_ASGN    = 63
	DZParserMOD_ASGN    = 64
	DZParserSHL_ASGN    = 65
	DZParserSHR_ASGN    = 66
	DZParserAND_ASGN    = 67
	DZParserOR_ASGN     = 68
	DZParserXOR_ASGN    = 69
	DZParserCONST       = 70
	DZParserINT_VALUE   = 71
	DZParserTYPE        = 72
	DZParserIDENTIFIER  = 73
	DZParserWHITESPACE  = 74
)

// DZParser rules.
const (
	DZParserRULE_start               = 0
	DZParserRULE_pkg                 = 1
	DZParserRULE_decl                = 2
	DZParserRULE_constDecl           = 3
	DZParserRULE_typeDecl            = 4
	DZParserRULE_typeSpec            = 5
	DZParserRULE_simpleTypeSpec      = 6
	DZParserRULE_basicTypeSpec       = 7
	DZParserRULE_namedTypeSpec       = 8
	DZParserRULE_refTypeSpec         = 9
	DZParserRULE_arrayTypeSpec       = 10
	DZParserRULE_sliceTypeSpec       = 11
	DZParserRULE_enumDecl            = 12
	DZParserRULE_enumOption          = 13
	DZParserRULE_structDecl          = 14
	DZParserRULE_structAttr          = 15
	DZParserRULE_procDecl            = 16
	DZParserRULE_procArg             = 17
	DZParserRULE_funcDecl            = 18
	DZParserRULE_funcArg             = 19
	DZParserRULE_block               = 20
	DZParserRULE_statement           = 21
	DZParserRULE_condition           = 22
	DZParserRULE_ifConditionBranch   = 23
	DZParserRULE_elifConditionBranch = 24
	DZParserRULE_elseConditionBranch = 25
	DZParserRULE_loop                = 26
	DZParserRULE_trueLoop            = 27
	DZParserRULE_whileLoop           = 28
	DZParserRULE_breakStatement      = 29
	DZParserRULE_continueStatement   = 30
	DZParserRULE_varDecl             = 31
	DZParserRULE_procCall            = 32
	DZParserRULE_procParam           = 33
	DZParserRULE_expression          = 34
	DZParserRULE_funcParam           = 35
	DZParserRULE_unaryOperator       = 36
	DZParserRULE_binaryOperator      = 37
	DZParserRULE_returnStatement     = 38
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

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(DZParserEOF, 0)
}

func (s *StartContext) AllDecl() []IDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclContext)(nil)).Elem())
	var tst = make([]IDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclContext)
		}
	}

	return tst
}

func (s *StartContext) Decl(i int) IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
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
		p.SetState(78)
		p.Pkg()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DZParserKW_TYPE)|(1<<DZParserKW_PROC)|(1<<DZParserKW_FUNC)|(1<<DZParserKW_CONST)|(1<<DZParserKW_STRUCT)|(1<<DZParserKW_ENUM))) != 0 {
		{
			p.SetState(79)
			p.Decl()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
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
		p.SetState(87)
		p.Match(DZParserKW_PKG)
	}
	{
		p.SetState(88)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*PkgContext).name = _m
	}
	{
		p.SetState(89)
		p.Match(DZParserSEMICOLON)
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

func (s *DeclContext) ConstDecl() IConstDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (s *DeclContext) TypeDecl() ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *DeclContext) EnumDecl() IEnumDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDeclContext)
}

func (s *DeclContext) StructDecl() IStructDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *DeclContext) ProcDecl() IProcDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcDeclContext)
}

func (s *DeclContext) FuncDecl() IFuncDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
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
	p.EnterRule(localctx, 4, DZParserRULE_decl)

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

	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_CONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.ConstDecl()
		}

	case DZParserKW_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.TypeDecl()
		}

	case DZParserKW_ENUM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.EnumDecl()
		}

	case DZParserKW_STRUCT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(94)
			p.StructDecl()
		}

	case DZParserKW_PROC:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(95)
			p.ProcDecl()
		}

	case DZParserKW_FUNC:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(96)
			p.FuncDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstDeclContext is an interface to support dynamic dispatch.
type IConstDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// GetTName returns the tName rule contexts.
	GetTName() IBasicTypeSpecContext

	// SetTName sets the tName rule contexts.
	SetTName(IBasicTypeSpecContext)

	// IsConstDeclContext differentiates from other interfaces.
	IsConstDeclContext()
}

type ConstDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	tName  IBasicTypeSpecContext
	value  antlr.Token
}

func NewEmptyConstDeclContext() *ConstDeclContext {
	var p = new(ConstDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_constDecl
	return p
}

func (*ConstDeclContext) IsConstDeclContext() {}

func NewConstDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclContext {
	var p = new(ConstDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_constDecl

	return p
}

func (s *ConstDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDeclContext) GetName() antlr.Token { return s.name }

func (s *ConstDeclContext) GetValue() antlr.Token { return s.value }

func (s *ConstDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *ConstDeclContext) SetValue(v antlr.Token) { s.value = v }

func (s *ConstDeclContext) GetTName() IBasicTypeSpecContext { return s.tName }

func (s *ConstDeclContext) SetTName(v IBasicTypeSpecContext) { s.tName = v }

func (s *ConstDeclContext) KW_CONST() antlr.TerminalNode {
	return s.GetToken(DZParserKW_CONST, 0)
}

func (s *ConstDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *ConstDeclContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *ConstDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *ConstDeclContext) AllCONST() []antlr.TerminalNode {
	return s.GetTokens(DZParserCONST)
}

func (s *ConstDeclContext) CONST(i int) antlr.TerminalNode {
	return s.GetToken(DZParserCONST, i)
}

func (s *ConstDeclContext) BasicTypeSpec() IBasicTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicTypeSpecContext)
}

func (s *ConstDeclContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(DZParserINT_VALUE, 0)
}

func (s *ConstDeclContext) KW_TRUE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_TRUE, 0)
}

func (s *ConstDeclContext) KW_FALSE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_FALSE, 0)
}

func (s *ConstDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterConstDecl(s)
	}
}

func (s *ConstDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitConstDecl(s)
	}
}

func (p *DZParser) ConstDecl() (localctx IConstDeclContext) {
	localctx = NewConstDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DZParserRULE_constDecl)
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
		p.SetState(99)
		p.Match(DZParserKW_CONST)
	}
	{
		p.SetState(100)

		var _m = p.Match(DZParserCONST)

		localctx.(*ConstDeclContext).name = _m
	}
	{
		p.SetState(101)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(102)

		var _x = p.BasicTypeSpec()

		localctx.(*ConstDeclContext).tName = _x
	}
	{
		p.SetState(103)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(104)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ConstDeclContext).value = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == DZParserKW_TRUE || _la == DZParserKW_FALSE || _la == DZParserCONST || _la == DZParserINT_VALUE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ConstDeclContext).value = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(105)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetTName returns the tName rule contexts.
	GetTName() ITypeSpecContext

	// SetTName sets the tName rule contexts.
	SetTName(ITypeSpecContext)

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	tName  ITypeSpecContext
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) GetName() antlr.Token { return s.name }

func (s *TypeDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *TypeDeclContext) GetTName() ITypeSpecContext { return s.tName }

func (s *TypeDeclContext) SetTName(v ITypeSpecContext) { s.tName = v }

func (s *TypeDeclContext) KW_TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_TYPE, 0)
}

func (s *TypeDeclContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *TypeDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *TypeDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
}

func (s *TypeDeclContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *DZParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DZParserRULE_typeDecl)

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
		p.SetState(107)
		p.Match(DZParserKW_TYPE)
	}
	{
		p.SetState(108)

		var _m = p.Match(DZParserTYPE)

		localctx.(*TypeDeclContext).name = _m
	}
	{
		p.SetState(109)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(110)

		var _x = p.TypeSpec()

		localctx.(*TypeDeclContext).tName = _x
	}
	{
		p.SetState(111)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) SimpleTypeSpec() ISimpleTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleTypeSpecContext)
}

func (s *TypeSpecContext) RefTypeSpec() IRefTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRefTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRefTypeSpecContext)
}

func (s *TypeSpecContext) ArrayTypeSpec() IArrayTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeSpecContext)
}

func (s *TypeSpecContext) SliceTypeSpec() ISliceTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISliceTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISliceTypeSpecContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (p *DZParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DZParserRULE_typeSpec)

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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.SimpleTypeSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.RefTypeSpec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.ArrayTypeSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)
			p.SliceTypeSpec()
		}

	}

	return localctx
}

// ISimpleTypeSpecContext is an interface to support dynamic dispatch.
type ISimpleTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleTypeSpecContext differentiates from other interfaces.
	IsSimpleTypeSpecContext()
}

type SimpleTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleTypeSpecContext() *SimpleTypeSpecContext {
	var p = new(SimpleTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_simpleTypeSpec
	return p
}

func (*SimpleTypeSpecContext) IsSimpleTypeSpecContext() {}

func NewSimpleTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTypeSpecContext {
	var p = new(SimpleTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_simpleTypeSpec

	return p
}

func (s *SimpleTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTypeSpecContext) BasicTypeSpec() IBasicTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicTypeSpecContext)
}

func (s *SimpleTypeSpecContext) NamedTypeSpec() INamedTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeSpecContext)
}

func (s *SimpleTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterSimpleTypeSpec(s)
	}
}

func (s *SimpleTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitSimpleTypeSpec(s)
	}
}

func (p *DZParser) SimpleTypeSpec() (localctx ISimpleTypeSpecContext) {
	localctx = NewSimpleTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DZParserRULE_simpleTypeSpec)

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

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserI8_T, DZParserU8_T, DZParserI16_T, DZParserU16_T, DZParserI32_T, DZParserU32_T, DZParserI64_T, DZParserU64_T, DZParserCHAR_T, DZParserBOOL_T, DZParserSIZE_T, DZParserSSIZE_T:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.BasicTypeSpec()
		}

	case DZParserTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.NamedTypeSpec()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBasicTypeSpecContext is an interface to support dynamic dispatch.
type IBasicTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicTypeSpecContext differentiates from other interfaces.
	IsBasicTypeSpecContext()
}

type BasicTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicTypeSpecContext() *BasicTypeSpecContext {
	var p = new(BasicTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_basicTypeSpec
	return p
}

func (*BasicTypeSpecContext) IsBasicTypeSpecContext() {}

func NewBasicTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicTypeSpecContext {
	var p = new(BasicTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_basicTypeSpec

	return p
}

func (s *BasicTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicTypeSpecContext) I8_T() antlr.TerminalNode {
	return s.GetToken(DZParserI8_T, 0)
}

func (s *BasicTypeSpecContext) U8_T() antlr.TerminalNode {
	return s.GetToken(DZParserU8_T, 0)
}

func (s *BasicTypeSpecContext) I16_T() antlr.TerminalNode {
	return s.GetToken(DZParserI16_T, 0)
}

func (s *BasicTypeSpecContext) U16_T() antlr.TerminalNode {
	return s.GetToken(DZParserU16_T, 0)
}

func (s *BasicTypeSpecContext) I32_T() antlr.TerminalNode {
	return s.GetToken(DZParserI32_T, 0)
}

func (s *BasicTypeSpecContext) U32_T() antlr.TerminalNode {
	return s.GetToken(DZParserU32_T, 0)
}

func (s *BasicTypeSpecContext) I64_T() antlr.TerminalNode {
	return s.GetToken(DZParserI64_T, 0)
}

func (s *BasicTypeSpecContext) U64_T() antlr.TerminalNode {
	return s.GetToken(DZParserU64_T, 0)
}

func (s *BasicTypeSpecContext) SIZE_T() antlr.TerminalNode {
	return s.GetToken(DZParserSIZE_T, 0)
}

func (s *BasicTypeSpecContext) SSIZE_T() antlr.TerminalNode {
	return s.GetToken(DZParserSSIZE_T, 0)
}

func (s *BasicTypeSpecContext) CHAR_T() antlr.TerminalNode {
	return s.GetToken(DZParserCHAR_T, 0)
}

func (s *BasicTypeSpecContext) BOOL_T() antlr.TerminalNode {
	return s.GetToken(DZParserBOOL_T, 0)
}

func (s *BasicTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterBasicTypeSpec(s)
	}
}

func (s *BasicTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitBasicTypeSpec(s)
	}
}

func (p *DZParser) BasicTypeSpec() (localctx IBasicTypeSpecContext) {
	localctx = NewBasicTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DZParserRULE_basicTypeSpec)
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
		p.SetState(123)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DZParserI8_T)|(1<<DZParserU8_T)|(1<<DZParserI16_T)|(1<<DZParserU16_T)|(1<<DZParserI32_T)|(1<<DZParserU32_T)|(1<<DZParserI64_T)|(1<<DZParserU64_T)|(1<<DZParserCHAR_T)|(1<<DZParserBOOL_T)|(1<<DZParserSIZE_T)|(1<<DZParserSSIZE_T))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INamedTypeSpecContext is an interface to support dynamic dispatch.
type INamedTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsNamedTypeSpecContext differentiates from other interfaces.
	IsNamedTypeSpecContext()
}

type NamedTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyNamedTypeSpecContext() *NamedTypeSpecContext {
	var p = new(NamedTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_namedTypeSpec
	return p
}

func (*NamedTypeSpecContext) IsNamedTypeSpecContext() {}

func NewNamedTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedTypeSpecContext {
	var p = new(NamedTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_namedTypeSpec

	return p
}

func (s *NamedTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedTypeSpecContext) GetName() antlr.Token { return s.name }

func (s *NamedTypeSpecContext) SetName(v antlr.Token) { s.name = v }

func (s *NamedTypeSpecContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
}

func (s *NamedTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterNamedTypeSpec(s)
	}
}

func (s *NamedTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitNamedTypeSpec(s)
	}
}

func (p *DZParser) NamedTypeSpec() (localctx INamedTypeSpecContext) {
	localctx = NewNamedTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DZParserRULE_namedTypeSpec)

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
		p.SetState(125)

		var _m = p.Match(DZParserTYPE)

		localctx.(*NamedTypeSpecContext).name = _m
	}

	return localctx
}

// IRefTypeSpecContext is an interface to support dynamic dispatch.
type IRefTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTName returns the tName rule contexts.
	GetTName() ISimpleTypeSpecContext

	// SetTName sets the tName rule contexts.
	SetTName(ISimpleTypeSpecContext)

	// IsRefTypeSpecContext differentiates from other interfaces.
	IsRefTypeSpecContext()
}

type RefTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tName  ISimpleTypeSpecContext
}

func NewEmptyRefTypeSpecContext() *RefTypeSpecContext {
	var p = new(RefTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_refTypeSpec
	return p
}

func (*RefTypeSpecContext) IsRefTypeSpecContext() {}

func NewRefTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RefTypeSpecContext {
	var p = new(RefTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_refTypeSpec

	return p
}

func (s *RefTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *RefTypeSpecContext) GetTName() ISimpleTypeSpecContext { return s.tName }

func (s *RefTypeSpecContext) SetTName(v ISimpleTypeSpecContext) { s.tName = v }

func (s *RefTypeSpecContext) REF() antlr.TerminalNode {
	return s.GetToken(DZParserREF, 0)
}

func (s *RefTypeSpecContext) SimpleTypeSpec() ISimpleTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleTypeSpecContext)
}

func (s *RefTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RefTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterRefTypeSpec(s)
	}
}

func (s *RefTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitRefTypeSpec(s)
	}
}

func (p *DZParser) RefTypeSpec() (localctx IRefTypeSpecContext) {
	localctx = NewRefTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DZParserRULE_refTypeSpec)

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
		p.Match(DZParserREF)
	}
	{
		p.SetState(128)

		var _x = p.SimpleTypeSpec()

		localctx.(*RefTypeSpecContext).tName = _x
	}

	return localctx
}

// IArrayTypeSpecContext is an interface to support dynamic dispatch.
type IArrayTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSize returns the size token.
	GetSize() antlr.Token

	// SetSize sets the size token.
	SetSize(antlr.Token)

	// GetTName returns the tName rule contexts.
	GetTName() ISimpleTypeSpecContext

	// SetTName sets the tName rule contexts.
	SetTName(ISimpleTypeSpecContext)

	// IsArrayTypeSpecContext differentiates from other interfaces.
	IsArrayTypeSpecContext()
}

type ArrayTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tName  ISimpleTypeSpecContext
	size   antlr.Token
}

func NewEmptyArrayTypeSpecContext() *ArrayTypeSpecContext {
	var p = new(ArrayTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_arrayTypeSpec
	return p
}

func (*ArrayTypeSpecContext) IsArrayTypeSpecContext() {}

func NewArrayTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeSpecContext {
	var p = new(ArrayTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_arrayTypeSpec

	return p
}

func (s *ArrayTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeSpecContext) GetSize() antlr.Token { return s.size }

func (s *ArrayTypeSpecContext) SetSize(v antlr.Token) { s.size = v }

func (s *ArrayTypeSpecContext) GetTName() ISimpleTypeSpecContext { return s.tName }

func (s *ArrayTypeSpecContext) SetTName(v ISimpleTypeSpecContext) { s.tName = v }

func (s *ArrayTypeSpecContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRK, 0)
}

func (s *ArrayTypeSpecContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *ArrayTypeSpecContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRK, 0)
}

func (s *ArrayTypeSpecContext) SimpleTypeSpec() ISimpleTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleTypeSpecContext)
}

func (s *ArrayTypeSpecContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(DZParserINT_VALUE, 0)
}

func (s *ArrayTypeSpecContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
}

func (s *ArrayTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterArrayTypeSpec(s)
	}
}

func (s *ArrayTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitArrayTypeSpec(s)
	}
}

func (p *DZParser) ArrayTypeSpec() (localctx IArrayTypeSpecContext) {
	localctx = NewArrayTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DZParserRULE_arrayTypeSpec)
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
		p.SetState(130)
		p.Match(DZParserLEFT_BRK)
	}
	{
		p.SetState(131)

		var _x = p.SimpleTypeSpec()

		localctx.(*ArrayTypeSpecContext).tName = _x
	}
	{
		p.SetState(132)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(133)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ArrayTypeSpecContext).size = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == DZParserCONST || _la == DZParserINT_VALUE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ArrayTypeSpecContext).size = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(134)
		p.Match(DZParserRIGHT_BRK)
	}

	return localctx
}

// ISliceTypeSpecContext is an interface to support dynamic dispatch.
type ISliceTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTName returns the tName rule contexts.
	GetTName() ISimpleTypeSpecContext

	// SetTName sets the tName rule contexts.
	SetTName(ISimpleTypeSpecContext)

	// IsSliceTypeSpecContext differentiates from other interfaces.
	IsSliceTypeSpecContext()
}

type SliceTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tName  ISimpleTypeSpecContext
}

func NewEmptySliceTypeSpecContext() *SliceTypeSpecContext {
	var p = new(SliceTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_sliceTypeSpec
	return p
}

func (*SliceTypeSpecContext) IsSliceTypeSpecContext() {}

func NewSliceTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceTypeSpecContext {
	var p = new(SliceTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_sliceTypeSpec

	return p
}

func (s *SliceTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceTypeSpecContext) GetTName() ISimpleTypeSpecContext { return s.tName }

func (s *SliceTypeSpecContext) SetTName(v ISimpleTypeSpecContext) { s.tName = v }

func (s *SliceTypeSpecContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRK, 0)
}

func (s *SliceTypeSpecContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRK, 0)
}

func (s *SliceTypeSpecContext) SimpleTypeSpec() ISimpleTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleTypeSpecContext)
}

func (s *SliceTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterSliceTypeSpec(s)
	}
}

func (s *SliceTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitSliceTypeSpec(s)
	}
}

func (p *DZParser) SliceTypeSpec() (localctx ISliceTypeSpecContext) {
	localctx = NewSliceTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DZParserRULE_sliceTypeSpec)

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
		p.Match(DZParserLEFT_BRK)
	}
	{
		p.SetState(137)

		var _x = p.SimpleTypeSpec()

		localctx.(*SliceTypeSpecContext).tName = _x
	}
	{
		p.SetState(138)
		p.Match(DZParserRIGHT_BRK)
	}

	return localctx
}

// IEnumDeclContext is an interface to support dynamic dispatch.
type IEnumDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsEnumDeclContext differentiates from other interfaces.
	IsEnumDeclContext()
}

type EnumDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyEnumDeclContext() *EnumDeclContext {
	var p = new(EnumDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_enumDecl
	return p
}

func (*EnumDeclContext) IsEnumDeclContext() {}

func NewEnumDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclContext {
	var p = new(EnumDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_enumDecl

	return p
}

func (s *EnumDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDeclContext) GetName() antlr.Token { return s.name }

func (s *EnumDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *EnumDeclContext) KW_ENUM() antlr.TerminalNode {
	return s.GetToken(DZParserKW_ENUM, 0)
}

func (s *EnumDeclContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *EnumDeclContext) AllEnumOption() []IEnumOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumOptionContext)(nil)).Elem())
	var tst = make([]IEnumOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumOptionContext)
		}
	}

	return tst
}

func (s *EnumDeclContext) EnumOption(i int) IEnumOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumOptionContext)
}

func (s *EnumDeclContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *EnumDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
}

func (s *EnumDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DZParserCOMMA)
}

func (s *EnumDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DZParserCOMMA, i)
}

func (s *EnumDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterEnumDecl(s)
	}
}

func (s *EnumDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitEnumDecl(s)
	}
}

func (p *DZParser) EnumDecl() (localctx IEnumDeclContext) {
	localctx = NewEnumDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DZParserRULE_enumDecl)
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
		p.SetState(140)
		p.Match(DZParserKW_ENUM)
	}
	{
		p.SetState(141)

		var _m = p.Match(DZParserTYPE)

		localctx.(*EnumDeclContext).name = _m
	}
	{
		p.SetState(142)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(143)
		p.EnumOption()
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserCOMMA {
		{
			p.SetState(144)
			p.Match(DZParserCOMMA)
		}
		{
			p.SetState(145)
			p.EnumOption()
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// IEnumOptionContext is an interface to support dynamic dispatch.
type IEnumOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsEnumOptionContext differentiates from other interfaces.
	IsEnumOptionContext()
}

type EnumOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyEnumOptionContext() *EnumOptionContext {
	var p = new(EnumOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_enumOption
	return p
}

func (*EnumOptionContext) IsEnumOptionContext() {}

func NewEnumOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumOptionContext {
	var p = new(EnumOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_enumOption

	return p
}

func (s *EnumOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumOptionContext) GetName() antlr.Token { return s.name }

func (s *EnumOptionContext) SetName(v antlr.Token) { s.name = v }

func (s *EnumOptionContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
}

func (s *EnumOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterEnumOption(s)
	}
}

func (s *EnumOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitEnumOption(s)
	}
}

func (p *DZParser) EnumOption() (localctx IEnumOptionContext) {
	localctx = NewEnumOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DZParserRULE_enumOption)

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

		var _m = p.Match(DZParserCONST)

		localctx.(*EnumOptionContext).name = _m
	}

	return localctx
}

// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_structDecl
	return p
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) GetName() antlr.Token { return s.name }

func (s *StructDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *StructDeclContext) KW_STRUCT() antlr.TerminalNode {
	return s.GetToken(DZParserKW_STRUCT, 0)
}

func (s *StructDeclContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *StructDeclContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *StructDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(DZParserTYPE, 0)
}

func (s *StructDeclContext) AllStructAttr() []IStructAttrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructAttrContext)(nil)).Elem())
	var tst = make([]IStructAttrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructAttrContext)
		}
	}

	return tst
}

func (s *StructDeclContext) StructAttr(i int) IStructAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructAttrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructAttrContext)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (p *DZParser) StructDecl() (localctx IStructDeclContext) {
	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DZParserRULE_structDecl)
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
		p.Match(DZParserKW_STRUCT)
	}
	{
		p.SetState(156)

		var _m = p.Match(DZParserTYPE)

		localctx.(*StructDeclContext).name = _m
	}
	{
		p.SetState(157)
		p.Match(DZParserLEFT_BRC)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserIDENTIFIER {
		{
			p.SetState(158)
			p.StructAttr()
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(164)
		p.Match(DZParserRIGHT_BRC)
	}

	return localctx
}

// IStructAttrContext is an interface to support dynamic dispatch.
type IStructAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetTName returns the tName rule contexts.
	GetTName() ITypeSpecContext

	// SetTName sets the tName rule contexts.
	SetTName(ITypeSpecContext)

	// IsStructAttrContext differentiates from other interfaces.
	IsStructAttrContext()
}

type StructAttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	tName  ITypeSpecContext
}

func NewEmptyStructAttrContext() *StructAttrContext {
	var p = new(StructAttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_structAttr
	return p
}

func (*StructAttrContext) IsStructAttrContext() {}

func NewStructAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructAttrContext {
	var p = new(StructAttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_structAttr

	return p
}

func (s *StructAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *StructAttrContext) GetName() antlr.Token { return s.name }

func (s *StructAttrContext) SetName(v antlr.Token) { s.name = v }

func (s *StructAttrContext) GetTName() ITypeSpecContext { return s.tName }

func (s *StructAttrContext) SetTName(v ITypeSpecContext) { s.tName = v }

func (s *StructAttrContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *StructAttrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
}

func (s *StructAttrContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *StructAttrContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *StructAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterStructAttr(s)
	}
}

func (s *StructAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitStructAttr(s)
	}
}

func (p *DZParser) StructAttr() (localctx IStructAttrContext) {
	localctx = NewStructAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DZParserRULE_structAttr)

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
		p.SetState(166)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*StructAttrContext).name = _m
	}
	{
		p.SetState(167)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(168)

		var _x = p.TypeSpec()

		localctx.(*StructAttrContext).tName = _x
	}
	{
		p.SetState(169)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// IProcDeclContext is an interface to support dynamic dispatch.
type IProcDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetBody returns the body rule contexts.
	GetBody() IBlockContext

	// SetBody sets the body rule contexts.
	SetBody(IBlockContext)

	// IsProcDeclContext differentiates from other interfaces.
	IsProcDeclContext()
}

type ProcDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	body   IBlockContext
}

func NewEmptyProcDeclContext() *ProcDeclContext {
	var p = new(ProcDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_procDecl
	return p
}

func (*ProcDeclContext) IsProcDeclContext() {}

func NewProcDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcDeclContext {
	var p = new(ProcDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_procDecl

	return p
}

func (s *ProcDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcDeclContext) GetName() antlr.Token { return s.name }

func (s *ProcDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *ProcDeclContext) GetBody() IBlockContext { return s.body }

func (s *ProcDeclContext) SetBody(v IBlockContext) { s.body = v }

func (s *ProcDeclContext) KW_PROC() antlr.TerminalNode {
	return s.GetToken(DZParserKW_PROC, 0)
}

func (s *ProcDeclContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_PRT, 0)
}

func (s *ProcDeclContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_PRT, 0)
}

func (s *ProcDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *ProcDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProcDeclContext) AllProcArg() []IProcArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProcArgContext)(nil)).Elem())
	var tst = make([]IProcArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProcArgContext)
		}
	}

	return tst
}

func (s *ProcDeclContext) ProcArg(i int) IProcArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProcArgContext)
}

func (s *ProcDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DZParserCOMMA)
}

func (s *ProcDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DZParserCOMMA, i)
}

func (s *ProcDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterProcDecl(s)
	}
}

func (s *ProcDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitProcDecl(s)
	}
}

func (p *DZParser) ProcDecl() (localctx IProcDeclContext) {
	localctx = NewProcDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DZParserRULE_procDecl)
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
		p.SetState(171)
		p.Match(DZParserKW_PROC)
	}
	{
		p.SetState(172)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcDeclContext).name = _m
	}
	{
		p.SetState(173)
		p.Match(DZParserLEFT_PRT)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserIDENTIFIER {
		{
			p.SetState(174)
			p.ProcArg()
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserCOMMA {
			{
				p.SetState(175)
				p.Match(DZParserCOMMA)
			}
			{
				p.SetState(176)
				p.ProcArg()
			}

			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(184)
		p.Match(DZParserRIGHT_PRT)
	}
	{
		p.SetState(185)

		var _x = p.Block()

		localctx.(*ProcDeclContext).body = _x
	}

	return localctx
}

// IProcArgContext is an interface to support dynamic dispatch.
type IProcArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetTName returns the tName rule contexts.
	GetTName() ITypeSpecContext

	// SetTName sets the tName rule contexts.
	SetTName(ITypeSpecContext)

	// IsProcArgContext differentiates from other interfaces.
	IsProcArgContext()
}

type ProcArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	tName  ITypeSpecContext
}

func NewEmptyProcArgContext() *ProcArgContext {
	var p = new(ProcArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_procArg
	return p
}

func (*ProcArgContext) IsProcArgContext() {}

func NewProcArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcArgContext {
	var p = new(ProcArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_procArg

	return p
}

func (s *ProcArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcArgContext) GetName() antlr.Token { return s.name }

func (s *ProcArgContext) SetName(v antlr.Token) { s.name = v }

func (s *ProcArgContext) GetTName() ITypeSpecContext { return s.tName }

func (s *ProcArgContext) SetTName(v ITypeSpecContext) { s.tName = v }

func (s *ProcArgContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *ProcArgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *ProcArgContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *ProcArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterProcArg(s)
	}
}

func (s *ProcArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitProcArg(s)
	}
}

func (p *DZParser) ProcArg() (localctx IProcArgContext) {
	localctx = NewProcArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DZParserRULE_procArg)

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
		p.SetState(187)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcArgContext).name = _m
	}
	{
		p.SetState(188)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(189)

		var _x = p.TypeSpec()

		localctx.(*ProcArgContext).tName = _x
	}

	return localctx
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetTName returns the tName rule contexts.
	GetTName() ITypeSpecContext

	// GetBody returns the body rule contexts.
	GetBody() IBlockContext

	// SetTName sets the tName rule contexts.
	SetTName(ITypeSpecContext)

	// SetBody sets the body rule contexts.
	SetBody(IBlockContext)

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	tName  ITypeSpecContext
	body   IBlockContext
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_funcDecl
	return p
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) GetName() antlr.Token { return s.name }

func (s *FuncDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *FuncDeclContext) GetTName() ITypeSpecContext { return s.tName }

func (s *FuncDeclContext) GetBody() IBlockContext { return s.body }

func (s *FuncDeclContext) SetTName(v ITypeSpecContext) { s.tName = v }

func (s *FuncDeclContext) SetBody(v IBlockContext) { s.body = v }

func (s *FuncDeclContext) KW_FUNC() antlr.TerminalNode {
	return s.GetToken(DZParserKW_FUNC, 0)
}

func (s *FuncDeclContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_PRT, 0)
}

func (s *FuncDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *FuncDeclContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_PRT, 0)
}

func (s *FuncDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *FuncDeclContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *FuncDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclContext) AllFuncArg() []IFuncArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncArgContext)(nil)).Elem())
	var tst = make([]IFuncArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncArgContext)
		}
	}

	return tst
}

func (s *FuncDeclContext) FuncArg(i int) IFuncArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncArgContext)
}

func (s *FuncDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DZParserCOMMA)
}

func (s *FuncDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DZParserCOMMA, i)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (p *DZParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DZParserRULE_funcDecl)
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
		p.SetState(191)
		p.Match(DZParserKW_FUNC)
	}
	{
		p.SetState(192)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*FuncDeclContext).name = _m
	}
	{
		p.SetState(193)
		p.Match(DZParserLEFT_PRT)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserIDENTIFIER {
		{
			p.SetState(194)
			p.FuncArg()
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserCOMMA {
			{
				p.SetState(195)
				p.Match(DZParserCOMMA)
			}
			{
				p.SetState(196)
				p.FuncArg()
			}

			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(204)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(205)

		var _x = p.TypeSpec()

		localctx.(*FuncDeclContext).tName = _x
	}
	{
		p.SetState(206)
		p.Match(DZParserRIGHT_PRT)
	}
	{
		p.SetState(207)

		var _x = p.Block()

		localctx.(*FuncDeclContext).body = _x
	}

	return localctx
}

// IFuncArgContext is an interface to support dynamic dispatch.
type IFuncArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetTName returns the tName rule contexts.
	GetTName() ITypeSpecContext

	// SetTName sets the tName rule contexts.
	SetTName(ITypeSpecContext)

	// IsFuncArgContext differentiates from other interfaces.
	IsFuncArgContext()
}

type FuncArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	tName  ITypeSpecContext
}

func NewEmptyFuncArgContext() *FuncArgContext {
	var p = new(FuncArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_funcArg
	return p
}

func (*FuncArgContext) IsFuncArgContext() {}

func NewFuncArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgContext {
	var p = new(FuncArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_funcArg

	return p
}

func (s *FuncArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgContext) GetName() antlr.Token { return s.name }

func (s *FuncArgContext) SetName(v antlr.Token) { s.name = v }

func (s *FuncArgContext) GetTName() ITypeSpecContext { return s.tName }

func (s *FuncArgContext) SetTName(v ITypeSpecContext) { s.tName = v }

func (s *FuncArgContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *FuncArgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *FuncArgContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *FuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterFuncArg(s)
	}
}

func (s *FuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitFuncArg(s)
	}
}

func (p *DZParser) FuncArg() (localctx IFuncArgContext) {
	localctx = NewFuncArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DZParserRULE_funcArg)

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

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*FuncArgContext).name = _m
	}
	{
		p.SetState(210)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(211)

		var _x = p.TypeSpec()

		localctx.(*FuncArgContext).tName = _x
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
	p.EnterRule(localctx, 40, DZParserRULE_block)
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
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DZParserKW_WHILE)|(1<<DZParserKW_LOOP)|(1<<DZParserKW_BREAK)|(1<<DZParserKW_CONTINUE)|(1<<DZParserKW_IF)|(1<<DZParserKW_LET)|(1<<DZParserKW_RETURN)|(1<<DZParserKW_TRUE)|(1<<DZParserKW_FALSE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(DZParserLEFT_PRT-32))|(1<<(DZParserLEFT_BRK-32))|(1<<(DZParserSUB-32))|(1<<(DZParserNOT-32)))) != 0) || (((_la-70)&-(0x1f+1)) == 0 && ((1<<uint((_la-70)))&((1<<(DZParserCONST-70))|(1<<(DZParserINT_VALUE-70))|(1<<(DZParserIDENTIFIER-70)))) != 0) {
		{
			p.SetState(213)
			p.Statement()
		}

		p.SetState(218)
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

func (s *StatementContext) VarDecl() IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
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
	p.EnterRule(localctx, 42, DZParserRULE_statement)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.Condition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.Loop()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.BreakStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(222)
			p.ContinueStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(223)
			p.ProcCall()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(224)
			p.VarDecl()
		}
		{
			p.SetState(225)
			p.Match(DZParserSEMICOLON)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(227)
			p.expression(0)
		}
		{
			p.SetState(228)
			p.Match(DZParserSEMICOLON)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(230)
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
	p.EnterRule(localctx, 44, DZParserRULE_condition)
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
		p.IfConditionBranch()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserKW_ELIF {
		{
			p.SetState(234)
			p.ElifConditionBranch()
		}

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserKW_ELSE {
		{
			p.SetState(240)
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
	p.EnterRule(localctx, 46, DZParserRULE_ifConditionBranch)

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
		p.SetState(243)
		p.Match(DZParserKW_IF)
	}
	{
		p.SetState(244)

		var _x = p.expression(0)

		localctx.(*IfConditionBranchContext).cond = _x
	}
	{
		p.SetState(245)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(246)

		var _x = p.Block()

		localctx.(*IfConditionBranchContext).body = _x
	}
	{
		p.SetState(247)
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
	p.EnterRule(localctx, 48, DZParserRULE_elifConditionBranch)

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
		p.Match(DZParserKW_ELIF)
	}
	{
		p.SetState(250)

		var _x = p.expression(0)

		localctx.(*ElifConditionBranchContext).cond = _x
	}
	{
		p.SetState(251)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(252)

		var _x = p.Block()

		localctx.(*ElifConditionBranchContext).body = _x
	}
	{
		p.SetState(253)
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
	p.EnterRule(localctx, 50, DZParserRULE_elseConditionBranch)

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
		p.Match(DZParserKW_ELSE)
	}
	{
		p.SetState(256)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(257)

		var _x = p.Block()

		localctx.(*ElseConditionBranchContext).body = _x
	}
	{
		p.SetState(258)
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
	p.EnterRule(localctx, 52, DZParserRULE_loop)

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

	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_LOOP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.TrueLoop()
		}

	case DZParserKW_WHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
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
	p.EnterRule(localctx, 54, DZParserRULE_trueLoop)

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
		p.SetState(264)
		p.Match(DZParserKW_LOOP)
	}
	{
		p.SetState(265)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(266)

		var _x = p.Block()

		localctx.(*TrueLoopContext).body = _x
	}
	{
		p.SetState(267)
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
	p.EnterRule(localctx, 56, DZParserRULE_whileLoop)

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
		p.Match(DZParserKW_WHILE)
	}
	{
		p.SetState(270)

		var _x = p.expression(0)

		localctx.(*WhileLoopContext).cond = _x
	}
	{
		p.SetState(271)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(272)

		var _x = p.Block()

		localctx.(*WhileLoopContext).body = _x
	}
	{
		p.SetState(273)
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
	p.EnterRule(localctx, 58, DZParserRULE_breakStatement)

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
	p.EnterRule(localctx, 60, DZParserRULE_continueStatement)

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
		p.SetState(277)
		p.Match(DZParserKW_CONTINUE)
	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
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

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	value  IExpressionContext
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) GetName() antlr.Token { return s.name }

func (s *VarDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *VarDeclContext) GetValue() IExpressionContext { return s.value }

func (s *VarDeclContext) SetValue(v IExpressionContext) { s.value = v }

func (s *VarDeclContext) KW_LET() antlr.TerminalNode {
	return s.GetToken(DZParserKW_LET, 0)
}

func (s *VarDeclContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *VarDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *VarDeclContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (p *DZParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DZParserRULE_varDecl)

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
		p.SetState(279)
		p.Match(DZParserKW_LET)
	}
	{
		p.SetState(280)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*VarDeclContext).name = _m
	}
	{
		p.SetState(281)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(282)

		var _x = p.expression(0)

		localctx.(*VarDeclContext).value = _x
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
	p.EnterRule(localctx, 64, DZParserRULE_procCall)
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
		p.SetState(284)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcCallContext).procName = _m
	}
	{
		p.SetState(285)
		p.Match(DZParserLEFT_PRT)
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(DZParserKW_TRUE-18))|(1<<(DZParserKW_FALSE-18))|(1<<(DZParserLEFT_PRT-18))|(1<<(DZParserLEFT_BRK-18))|(1<<(DZParserSUB-18))|(1<<(DZParserNOT-18)))) != 0) || (((_la-70)&-(0x1f+1)) == 0 && ((1<<uint((_la-70)))&((1<<(DZParserCONST-70))|(1<<(DZParserINT_VALUE-70))|(1<<(DZParserIDENTIFIER-70)))) != 0) {
		{
			p.SetState(286)
			p.ProcParam()
		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserCOMMA {
			{
				p.SetState(287)
				p.Match(DZParserCOMMA)
			}
			{
				p.SetState(288)
				p.ProcParam()
			}

			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(296)
		p.Match(DZParserRIGHT_PRT)
	}
	{
		p.SetState(297)
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
	p.EnterRule(localctx, 66, DZParserRULE_procParam)

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
		p.SetState(299)

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

func (s *ExpressionContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(DZParserINT_VALUE, 0)
}

func (s *ExpressionContext) KW_TRUE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_TRUE, 0)
}

func (s *ExpressionContext) KW_FALSE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_FALSE, 0)
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
	_startState := 68
	p.EnterRecursionRule(localctx, 68, DZParserRULE_expression, _p)
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
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(302)

			var _m = p.Match(DZParserIDENTIFIER)

			localctx.(*ExpressionContext).varName = _m
		}

	case 2:
		{
			p.SetState(303)

			var _m = p.Match(DZParserINT_VALUE)

			localctx.(*ExpressionContext).intValue = _m
		}

	case 3:
		{
			p.SetState(304)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).boolValue = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DZParserKW_TRUE || _la == DZParserKW_FALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).boolValue = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		{
			p.SetState(305)

			var _m = p.Match(DZParserCONST)

			localctx.(*ExpressionContext).constName = _m
		}

	case 5:
		{
			p.SetState(306)
			p.Match(DZParserLEFT_BRK)
		}
		{
			p.SetState(307)

			var _x = p.expression(0)

			localctx.(*ExpressionContext).brkExpr = _x
		}
		{
			p.SetState(308)
			p.Match(DZParserRIGHT_BRK)
		}

	case 6:
		{
			p.SetState(310)
			p.Match(DZParserLEFT_PRT)
		}
		{
			p.SetState(311)

			var _x = p.expression(0)

			localctx.(*ExpressionContext).prtExpr = _x
		}
		{
			p.SetState(312)
			p.Match(DZParserRIGHT_PRT)
		}

	case 7:
		{
			p.SetState(314)

			var _m = p.Match(DZParserIDENTIFIER)

			localctx.(*ExpressionContext).funcName = _m
		}
		{
			p.SetState(315)
			p.Match(DZParserLEFT_PRT)
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(DZParserKW_TRUE-18))|(1<<(DZParserKW_FALSE-18))|(1<<(DZParserLEFT_PRT-18))|(1<<(DZParserLEFT_BRK-18))|(1<<(DZParserSUB-18))|(1<<(DZParserNOT-18)))) != 0) || (((_la-70)&-(0x1f+1)) == 0 && ((1<<uint((_la-70)))&((1<<(DZParserCONST-70))|(1<<(DZParserINT_VALUE-70))|(1<<(DZParserIDENTIFIER-70)))) != 0) {
			{
				p.SetState(316)
				p.FuncParam()
			}
			p.SetState(321)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == DZParserCOMMA {
				{
					p.SetState(317)
					p.Match(DZParserCOMMA)
				}
				{
					p.SetState(318)
					p.FuncParam()
				}

				p.SetState(323)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(326)
			p.Match(DZParserRIGHT_PRT)
		}

	case 8:
		{
			p.SetState(327)

			var _x = p.UnaryOperator()

			localctx.(*ExpressionContext).unOp = _x
		}
		{
			p.SetState(328)

			var _x = p.expression(1)

			localctx.(*ExpressionContext).unExpr = _x
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			localctx.(*ExpressionContext).lBinExpr = _prevctx
			p.PushNewRecursionContext(localctx, _startState, DZParserRULE_expression)
			p.SetState(332)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(333)

				var _x = p.BinaryOperator()

				localctx.(*ExpressionContext).binOp = _x
			}
			{
				p.SetState(334)

				var _x = p.expression(3)

				localctx.(*ExpressionContext).rBinExpr = _x
			}

		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 70, DZParserRULE_funcParam)

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
		p.SetState(341)

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
	p.EnterRule(localctx, 72, DZParserRULE_unaryOperator)
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
		p.SetState(343)
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

func (s *BinaryOperatorContext) EQL() antlr.TerminalNode {
	return s.GetToken(DZParserEQL, 0)
}

func (s *BinaryOperatorContext) GRT() antlr.TerminalNode {
	return s.GetToken(DZParserGRT, 0)
}

func (s *BinaryOperatorContext) LESS() antlr.TerminalNode {
	return s.GetToken(DZParserLESS, 0)
}

func (s *BinaryOperatorContext) GRT_EQL() antlr.TerminalNode {
	return s.GetToken(DZParserGRT_EQL, 0)
}

func (s *BinaryOperatorContext) LESS_EQL() antlr.TerminalNode {
	return s.GetToken(DZParserLESS_EQL, 0)
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
	p.EnterRule(localctx, 74, DZParserRULE_binaryOperator)
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
		p.SetState(345)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(DZParserADD-41))|(1<<(DZParserSUB-41))|(1<<(DZParserMUL-41))|(1<<(DZParserDIV-41))|(1<<(DZParserMOD-41))|(1<<(DZParserSHL-41))|(1<<(DZParserSHR-41))|(1<<(DZParserAND-41))|(1<<(DZParserOR-41))|(1<<(DZParserXOR-41))|(1<<(DZParserEQL-41))|(1<<(DZParserGRT-41))|(1<<(DZParserGRT_EQL-41))|(1<<(DZParserLESS-41))|(1<<(DZParserLESS_EQL-41))|(1<<(DZParserASGN-41))|(1<<(DZParserADD_ASGN-41))|(1<<(DZParserSUB_ASGN-41))|(1<<(DZParserMUL_ASGN-41))|(1<<(DZParserDIV_ASGN-41))|(1<<(DZParserMOD_ASGN-41))|(1<<(DZParserSHL_ASGN-41))|(1<<(DZParserSHR_ASGN-41))|(1<<(DZParserAND_ASGN-41))|(1<<(DZParserOR_ASGN-41))|(1<<(DZParserXOR_ASGN-41)))) != 0) {
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
	p.EnterRule(localctx, 76, DZParserRULE_returnStatement)

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

	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(347)
			p.Match(DZParserKW_RETURN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(348)
			p.Match(DZParserKW_RETURN)
		}
		{
			p.SetState(349)

			var _x = p.expression(0)

			localctx.(*ReturnStatementContext).value = _x
		}
		{
			p.SetState(350)
			p.Match(DZParserSEMICOLON)
		}

	}

	return localctx
}

func (p *DZParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 34:
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
