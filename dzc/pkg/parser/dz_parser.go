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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 78, 384,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2, 3, 2,
	7, 2, 89, 10, 2, 12, 2, 14, 2, 92, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 106, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 126, 10, 7, 3, 8, 3, 8, 5, 8, 130, 10, 8,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 7, 14, 155, 10, 14, 12, 14, 14, 14, 158, 11, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 168, 10, 16, 12, 16,
	14, 16, 171, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 186, 10, 18, 12, 18, 14,
	18, 189, 11, 18, 5, 18, 191, 10, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 206, 10,
	20, 12, 20, 14, 20, 209, 11, 20, 5, 20, 211, 10, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 7, 22, 224, 10,
	22, 12, 22, 14, 22, 227, 11, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 243,
	10, 23, 3, 24, 3, 24, 7, 24, 247, 10, 24, 12, 24, 14, 24, 250, 11, 24,
	3, 24, 5, 24, 253, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 5, 28, 268, 10, 28, 3, 29,
	3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 7, 34, 295, 10, 34, 12, 34, 14, 34, 298, 11, 34, 5,
	34, 300, 10, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 7, 36, 323, 10, 36, 12, 36, 14, 36, 326, 11, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 333, 10, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 7, 36, 342, 10, 36, 12, 36, 14, 36, 345, 11, 36,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 7, 39, 359, 10, 39, 12, 39, 14, 39, 362, 11, 39, 5, 39, 364,
	10, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42,
	3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 382, 10, 43, 3,
	43, 2, 3, 70, 44, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 72, 74, 76, 78, 80, 82, 84, 2, 8, 4, 2, 20, 21, 73, 75, 3, 2, 22,
	33, 3, 2, 73, 74, 4, 2, 20, 21, 74, 74, 4, 2, 45, 45, 49, 49, 5, 2, 44,
	48, 51, 56, 58, 72, 2, 383, 2, 86, 3, 2, 2, 2, 4, 95, 3, 2, 2, 2, 6, 105,
	3, 2, 2, 2, 8, 107, 3, 2, 2, 2, 10, 115, 3, 2, 2, 2, 12, 125, 3, 2, 2,
	2, 14, 129, 3, 2, 2, 2, 16, 131, 3, 2, 2, 2, 18, 133, 3, 2, 2, 2, 20, 135,
	3, 2, 2, 2, 22, 138, 3, 2, 2, 2, 24, 144, 3, 2, 2, 2, 26, 148, 3, 2, 2,
	2, 28, 161, 3, 2, 2, 2, 30, 163, 3, 2, 2, 2, 32, 174, 3, 2, 2, 2, 34, 179,
	3, 2, 2, 2, 36, 195, 3, 2, 2, 2, 38, 199, 3, 2, 2, 2, 40, 217, 3, 2, 2,
	2, 42, 221, 3, 2, 2, 2, 44, 242, 3, 2, 2, 2, 46, 244, 3, 2, 2, 2, 48, 254,
	3, 2, 2, 2, 50, 258, 3, 2, 2, 2, 52, 262, 3, 2, 2, 2, 54, 267, 3, 2, 2,
	2, 56, 269, 3, 2, 2, 2, 58, 272, 3, 2, 2, 2, 60, 276, 3, 2, 2, 2, 62, 279,
	3, 2, 2, 2, 64, 282, 3, 2, 2, 2, 66, 289, 3, 2, 2, 2, 68, 304, 3, 2, 2,
	2, 70, 332, 3, 2, 2, 2, 72, 346, 3, 2, 2, 2, 74, 351, 3, 2, 2, 2, 76, 353,
	3, 2, 2, 2, 78, 367, 3, 2, 2, 2, 80, 371, 3, 2, 2, 2, 82, 373, 3, 2, 2,
	2, 84, 381, 3, 2, 2, 2, 86, 90, 5, 4, 3, 2, 87, 89, 5, 6, 4, 2, 88, 87,
	3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2,
	91, 93, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 94, 7, 2, 2, 3, 94, 3, 3, 2,
	2, 2, 95, 96, 7, 3, 2, 2, 96, 97, 7, 77, 2, 2, 97, 98, 7, 43, 2, 2, 98,
	5, 3, 2, 2, 2, 99, 106, 5, 8, 5, 2, 100, 106, 5, 10, 6, 2, 101, 106, 5,
	26, 14, 2, 102, 106, 5, 30, 16, 2, 103, 106, 5, 34, 18, 2, 104, 106, 5,
	38, 20, 2, 105, 99, 3, 2, 2, 2, 105, 100, 3, 2, 2, 2, 105, 101, 3, 2, 2,
	2, 105, 102, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 104, 3, 2, 2, 2, 106,
	7, 3, 2, 2, 2, 107, 108, 7, 15, 2, 2, 108, 109, 7, 73, 2, 2, 109, 110,
	7, 40, 2, 2, 110, 111, 5, 16, 9, 2, 111, 112, 7, 62, 2, 2, 112, 113, 9,
	2, 2, 2, 113, 114, 7, 43, 2, 2, 114, 9, 3, 2, 2, 2, 115, 116, 7, 4, 2,
	2, 116, 117, 7, 76, 2, 2, 117, 118, 7, 62, 2, 2, 118, 119, 5, 12, 7, 2,
	119, 120, 7, 43, 2, 2, 120, 11, 3, 2, 2, 2, 121, 126, 5, 14, 8, 2, 122,
	126, 5, 20, 11, 2, 123, 126, 5, 22, 12, 2, 124, 126, 5, 24, 13, 2, 125,
	121, 3, 2, 2, 2, 125, 122, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 124,
	3, 2, 2, 2, 126, 13, 3, 2, 2, 2, 127, 130, 5, 16, 9, 2, 128, 130, 5, 18,
	10, 2, 129, 127, 3, 2, 2, 2, 129, 128, 3, 2, 2, 2, 130, 15, 3, 2, 2, 2,
	131, 132, 9, 3, 2, 2, 132, 17, 3, 2, 2, 2, 133, 134, 7, 76, 2, 2, 134,
	19, 3, 2, 2, 2, 135, 136, 7, 50, 2, 2, 136, 137, 5, 14, 8, 2, 137, 21,
	3, 2, 2, 2, 138, 139, 7, 38, 2, 2, 139, 140, 5, 14, 8, 2, 140, 141, 7,
	40, 2, 2, 141, 142, 9, 4, 2, 2, 142, 143, 7, 39, 2, 2, 143, 23, 3, 2, 2,
	2, 144, 145, 7, 38, 2, 2, 145, 146, 5, 14, 8, 2, 146, 147, 7, 39, 2, 2,
	147, 25, 3, 2, 2, 2, 148, 149, 7, 18, 2, 2, 149, 150, 7, 76, 2, 2, 150,
	151, 7, 36, 2, 2, 151, 156, 5, 28, 15, 2, 152, 153, 7, 41, 2, 2, 153, 155,
	5, 28, 15, 2, 154, 152, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3,
	2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158, 156, 3, 2, 2,
	2, 159, 160, 7, 37, 2, 2, 160, 27, 3, 2, 2, 2, 161, 162, 7, 73, 2, 2, 162,
	29, 3, 2, 2, 2, 163, 164, 7, 17, 2, 2, 164, 165, 7, 76, 2, 2, 165, 169,
	7, 36, 2, 2, 166, 168, 5, 32, 17, 2, 167, 166, 3, 2, 2, 2, 168, 171, 3,
	2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 172, 3, 2, 2,
	2, 171, 169, 3, 2, 2, 2, 172, 173, 7, 37, 2, 2, 173, 31, 3, 2, 2, 2, 174,
	175, 7, 77, 2, 2, 175, 176, 7, 40, 2, 2, 176, 177, 5, 12, 7, 2, 177, 178,
	7, 43, 2, 2, 178, 33, 3, 2, 2, 2, 179, 180, 7, 13, 2, 2, 180, 181, 7, 77,
	2, 2, 181, 190, 7, 34, 2, 2, 182, 187, 5, 36, 19, 2, 183, 184, 7, 41, 2,
	2, 184, 186, 5, 36, 19, 2, 185, 183, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2,
	187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 191, 3, 2, 2, 2, 189,
	187, 3, 2, 2, 2, 190, 182, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192,
	3, 2, 2, 2, 192, 193, 7, 35, 2, 2, 193, 194, 5, 42, 22, 2, 194, 35, 3,
	2, 2, 2, 195, 196, 7, 77, 2, 2, 196, 197, 7, 40, 2, 2, 197, 198, 5, 12,
	7, 2, 198, 37, 3, 2, 2, 2, 199, 200, 7, 14, 2, 2, 200, 201, 7, 77, 2, 2,
	201, 210, 7, 34, 2, 2, 202, 207, 5, 40, 21, 2, 203, 204, 7, 41, 2, 2, 204,
	206, 5, 40, 21, 2, 205, 203, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2, 207, 205,
	3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 211, 3, 2, 2, 2, 209, 207, 3, 2,
	2, 2, 210, 202, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2,
	212, 213, 7, 35, 2, 2, 213, 214, 7, 40, 2, 2, 214, 215, 5, 12, 7, 2, 215,
	216, 5, 42, 22, 2, 216, 39, 3, 2, 2, 2, 217, 218, 7, 77, 2, 2, 218, 219,
	7, 40, 2, 2, 219, 220, 5, 12, 7, 2, 220, 41, 3, 2, 2, 2, 221, 225, 7, 36,
	2, 2, 222, 224, 5, 44, 23, 2, 223, 222, 3, 2, 2, 2, 224, 227, 3, 2, 2,
	2, 225, 223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 228, 3, 2, 2, 2, 227,
	225, 3, 2, 2, 2, 228, 229, 7, 37, 2, 2, 229, 43, 3, 2, 2, 2, 230, 243,
	5, 46, 24, 2, 231, 243, 5, 54, 28, 2, 232, 243, 5, 60, 31, 2, 233, 243,
	5, 62, 32, 2, 234, 243, 5, 66, 34, 2, 235, 236, 5, 64, 33, 2, 236, 237,
	7, 43, 2, 2, 237, 243, 3, 2, 2, 2, 238, 239, 5, 70, 36, 2, 239, 240, 7,
	43, 2, 2, 240, 243, 3, 2, 2, 2, 241, 243, 5, 84, 43, 2, 242, 230, 3, 2,
	2, 2, 242, 231, 3, 2, 2, 2, 242, 232, 3, 2, 2, 2, 242, 233, 3, 2, 2, 2,
	242, 234, 3, 2, 2, 2, 242, 235, 3, 2, 2, 2, 242, 238, 3, 2, 2, 2, 242,
	241, 3, 2, 2, 2, 243, 45, 3, 2, 2, 2, 244, 248, 5, 48, 25, 2, 245, 247,
	5, 50, 26, 2, 246, 245, 3, 2, 2, 2, 247, 250, 3, 2, 2, 2, 248, 246, 3,
	2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 252, 3, 2, 2, 2, 250, 248, 3, 2, 2,
	2, 251, 253, 5, 52, 27, 2, 252, 251, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2,
	253, 47, 3, 2, 2, 2, 254, 255, 7, 10, 2, 2, 255, 256, 5, 70, 36, 2, 256,
	257, 5, 42, 22, 2, 257, 49, 3, 2, 2, 2, 258, 259, 7, 11, 2, 2, 259, 260,
	5, 70, 36, 2, 260, 261, 5, 42, 22, 2, 261, 51, 3, 2, 2, 2, 262, 263, 7,
	12, 2, 2, 263, 264, 5, 42, 22, 2, 264, 53, 3, 2, 2, 2, 265, 268, 5, 56,
	29, 2, 266, 268, 5, 58, 30, 2, 267, 265, 3, 2, 2, 2, 267, 266, 3, 2, 2,
	2, 268, 55, 3, 2, 2, 2, 269, 270, 7, 7, 2, 2, 270, 271, 5, 42, 22, 2, 271,
	57, 3, 2, 2, 2, 272, 273, 7, 6, 2, 2, 273, 274, 5, 70, 36, 2, 274, 275,
	5, 42, 22, 2, 275, 59, 3, 2, 2, 2, 276, 277, 7, 8, 2, 2, 277, 278, 7, 43,
	2, 2, 278, 61, 3, 2, 2, 2, 279, 280, 7, 9, 2, 2, 280, 281, 7, 43, 2, 2,
	281, 63, 3, 2, 2, 2, 282, 283, 7, 16, 2, 2, 283, 284, 7, 77, 2, 2, 284,
	285, 7, 40, 2, 2, 285, 286, 5, 12, 7, 2, 286, 287, 7, 62, 2, 2, 287, 288,
	5, 70, 36, 2, 288, 65, 3, 2, 2, 2, 289, 290, 7, 77, 2, 2, 290, 299, 7,
	34, 2, 2, 291, 296, 5, 68, 35, 2, 292, 293, 7, 41, 2, 2, 293, 295, 5, 68,
	35, 2, 294, 292, 3, 2, 2, 2, 295, 298, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2,
	296, 297, 3, 2, 2, 2, 297, 300, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 299,
	291, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 302,
	7, 35, 2, 2, 302, 303, 7, 43, 2, 2, 303, 67, 3, 2, 2, 2, 304, 305, 7, 77,
	2, 2, 305, 306, 7, 62, 2, 2, 306, 307, 5, 70, 36, 2, 307, 69, 3, 2, 2,
	2, 308, 309, 8, 36, 1, 2, 309, 333, 7, 77, 2, 2, 310, 333, 5, 74, 38, 2,
	311, 333, 7, 73, 2, 2, 312, 313, 7, 34, 2, 2, 313, 314, 5, 70, 36, 2, 314,
	315, 7, 35, 2, 2, 315, 333, 3, 2, 2, 2, 316, 317, 7, 38, 2, 2, 317, 318,
	5, 70, 36, 2, 318, 319, 7, 39, 2, 2, 319, 333, 3, 2, 2, 2, 320, 324, 7,
	36, 2, 2, 321, 323, 5, 72, 37, 2, 322, 321, 3, 2, 2, 2, 323, 326, 3, 2,
	2, 2, 324, 322, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 327, 3, 2, 2, 2,
	326, 324, 3, 2, 2, 2, 327, 333, 7, 37, 2, 2, 328, 333, 5, 76, 39, 2, 329,
	330, 5, 80, 41, 2, 330, 331, 5, 70, 36, 3, 331, 333, 3, 2, 2, 2, 332, 308,
	3, 2, 2, 2, 332, 310, 3, 2, 2, 2, 332, 311, 3, 2, 2, 2, 332, 312, 3, 2,
	2, 2, 332, 316, 3, 2, 2, 2, 332, 320, 3, 2, 2, 2, 332, 328, 3, 2, 2, 2,
	332, 329, 3, 2, 2, 2, 333, 343, 3, 2, 2, 2, 334, 335, 12, 4, 2, 2, 335,
	336, 5, 82, 42, 2, 336, 337, 5, 70, 36, 5, 337, 342, 3, 2, 2, 2, 338, 339,
	12, 6, 2, 2, 339, 340, 7, 42, 2, 2, 340, 342, 7, 77, 2, 2, 341, 334, 3,
	2, 2, 2, 341, 338, 3, 2, 2, 2, 342, 345, 3, 2, 2, 2, 343, 341, 3, 2, 2,
	2, 343, 344, 3, 2, 2, 2, 344, 71, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 346,
	347, 7, 42, 2, 2, 347, 348, 7, 77, 2, 2, 348, 349, 7, 62, 2, 2, 349, 350,
	5, 70, 36, 2, 350, 73, 3, 2, 2, 2, 351, 352, 9, 5, 2, 2, 352, 75, 3, 2,
	2, 2, 353, 354, 7, 77, 2, 2, 354, 363, 7, 34, 2, 2, 355, 360, 5, 78, 40,
	2, 356, 357, 7, 41, 2, 2, 357, 359, 5, 78, 40, 2, 358, 356, 3, 2, 2, 2,
	359, 362, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361,
	364, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 363, 355, 3, 2, 2, 2, 363, 364,
	3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 366, 7, 35, 2, 2, 366, 77, 3, 2,
	2, 2, 367, 368, 7, 77, 2, 2, 368, 369, 7, 62, 2, 2, 369, 370, 5, 70, 36,
	2, 370, 79, 3, 2, 2, 2, 371, 372, 9, 6, 2, 2, 372, 81, 3, 2, 2, 2, 373,
	374, 9, 7, 2, 2, 374, 83, 3, 2, 2, 2, 375, 376, 7, 19, 2, 2, 376, 382,
	7, 43, 2, 2, 377, 378, 7, 19, 2, 2, 378, 379, 5, 70, 36, 2, 379, 380, 7,
	43, 2, 2, 380, 382, 3, 2, 2, 2, 381, 375, 3, 2, 2, 2, 381, 377, 3, 2, 2,
	2, 382, 85, 3, 2, 2, 2, 26, 90, 105, 125, 129, 156, 169, 187, 190, 207,
	210, 225, 242, 248, 252, 267, 296, 299, 324, 332, 341, 343, 360, 363, 381,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'pkg'", "'type'", "'for'", "'while'", "'loop'", "'break'", "'continue'",
	"'if'", "'elif'", "'else'", "'proc'", "'func'", "'const'", "'let'", "'struct'",
	"'enum'", "'return'", "'true'", "'false'", "'i8_t'", "'u8_t'", "'i16_t'",
	"'u16_t'", "'i32_t'", "'u32_t'", "'i64_t'", "'u64_t'", "'char_t'", "'bool_t'",
	"'size_t'", "'ssize_t'", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'",
	"','", "'.'", "';'", "'+'", "'-'", "'*'", "'/'", "'%'", "'~'", "'@'", "'<<'",
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
	"LEFT_BRK", "RIGHT_BRK", "COLON", "COMMA", "DOT", "SEMICOLON", "ADD", "SUB",
	"MUL", "DIV", "MOD", "NOT", "REF", "SHL", "SHR", "AND", "OR", "XOR", "EQL",
	"NOT_EQL", "GRT", "GRT_EQL", "LESS", "LESS_EQL", "ASGN", "ADD_ASGN", "SUB_ASGN",
	"MUL_ASGN", "DIV_ASGN", "MOD_ASGN", "SHL_ASGN", "SHR_ASGN", "AND_ASGN",
	"OR_ASGN", "XOR_ASGN", "CONST", "INT_VALUE", "CHAR_VALUE", "TYPE", "IDENTIFIER",
	"WHITESPACE",
}

var ruleNames = []string{
	"start", "pkg", "decl", "constDecl", "typeDecl", "typeSpec", "simpleTypeSpec",
	"basicTypeSpec", "namedTypeSpec", "refTypeSpec", "arrayTypeSpec", "sliceTypeSpec",
	"enumDecl", "enumOption", "structDecl", "structAttr", "procDecl", "procArg",
	"funcDecl", "funcArg", "block", "statement", "condition", "ifConditionBranch",
	"elifConditionBranch", "elseConditionBranch", "loop", "trueLoop", "whileLoop",
	"breakStatement", "continueStatement", "varDecl", "procCall", "procParam",
	"expression", "attrAsgn", "constValue", "funcCall", "funcParam", "unaryOperator",
	"binaryOperator", "returnStatement",
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
	DZParserDOT         = 40
	DZParserSEMICOLON   = 41
	DZParserADD         = 42
	DZParserSUB         = 43
	DZParserMUL         = 44
	DZParserDIV         = 45
	DZParserMOD         = 46
	DZParserNOT         = 47
	DZParserREF         = 48
	DZParserSHL         = 49
	DZParserSHR         = 50
	DZParserAND         = 51
	DZParserOR          = 52
	DZParserXOR         = 53
	DZParserEQL         = 54
	DZParserNOT_EQL     = 55
	DZParserGRT         = 56
	DZParserGRT_EQL     = 57
	DZParserLESS        = 58
	DZParserLESS_EQL    = 59
	DZParserASGN        = 60
	DZParserADD_ASGN    = 61
	DZParserSUB_ASGN    = 62
	DZParserMUL_ASGN    = 63
	DZParserDIV_ASGN    = 64
	DZParserMOD_ASGN    = 65
	DZParserSHL_ASGN    = 66
	DZParserSHR_ASGN    = 67
	DZParserAND_ASGN    = 68
	DZParserOR_ASGN     = 69
	DZParserXOR_ASGN    = 70
	DZParserCONST       = 71
	DZParserINT_VALUE   = 72
	DZParserCHAR_VALUE  = 73
	DZParserTYPE        = 74
	DZParserIDENTIFIER  = 75
	DZParserWHITESPACE  = 76
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
	DZParserRULE_attrAsgn            = 35
	DZParserRULE_constValue          = 36
	DZParserRULE_funcCall            = 37
	DZParserRULE_funcParam           = 38
	DZParserRULE_unaryOperator       = 39
	DZParserRULE_binaryOperator      = 40
	DZParserRULE_returnStatement     = 41
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
		p.SetState(84)
		p.Pkg()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DZParserKW_TYPE)|(1<<DZParserKW_PROC)|(1<<DZParserKW_FUNC)|(1<<DZParserKW_CONST)|(1<<DZParserKW_STRUCT)|(1<<DZParserKW_ENUM))) != 0 {
		{
			p.SetState(85)
			p.Decl()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
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
		p.SetState(93)
		p.Match(DZParserKW_PKG)
	}
	{
		p.SetState(94)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*PkgContext).name = _m
	}
	{
		p.SetState(95)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_CONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.ConstDecl()
		}

	case DZParserKW_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.TypeDecl()
		}

	case DZParserKW_ENUM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.EnumDecl()
		}

	case DZParserKW_STRUCT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(100)
			p.StructDecl()
		}

	case DZParserKW_PROC:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(101)
			p.ProcDecl()
		}

	case DZParserKW_FUNC:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(102)
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

func (s *ConstDeclContext) CHAR_VALUE() antlr.TerminalNode {
	return s.GetToken(DZParserCHAR_VALUE, 0)
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
		p.SetState(105)
		p.Match(DZParserKW_CONST)
	}
	{
		p.SetState(106)

		var _m = p.Match(DZParserCONST)

		localctx.(*ConstDeclContext).name = _m
	}
	{
		p.SetState(107)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(108)

		var _x = p.BasicTypeSpec()

		localctx.(*ConstDeclContext).tName = _x
	}
	{
		p.SetState(109)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(110)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ConstDeclContext).value = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == DZParserKW_TRUE || _la == DZParserKW_FALSE || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(DZParserCONST-71))|(1<<(DZParserINT_VALUE-71))|(1<<(DZParserCHAR_VALUE-71)))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ConstDeclContext).value = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(111)
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
		p.SetState(113)
		p.Match(DZParserKW_TYPE)
	}
	{
		p.SetState(114)

		var _m = p.Match(DZParserTYPE)

		localctx.(*TypeDeclContext).name = _m
	}
	{
		p.SetState(115)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(116)

		var _x = p.TypeSpec()

		localctx.(*TypeDeclContext).tName = _x
	}
	{
		p.SetState(117)
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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.SimpleTypeSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.RefTypeSpec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.ArrayTypeSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserI8_T, DZParserU8_T, DZParserI16_T, DZParserU16_T, DZParserI32_T, DZParserU32_T, DZParserI64_T, DZParserU64_T, DZParserCHAR_T, DZParserBOOL_T, DZParserSIZE_T, DZParserSSIZE_T:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.BasicTypeSpec()
		}

	case DZParserTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
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
		p.SetState(129)
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
		p.SetState(131)

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
		p.SetState(133)
		p.Match(DZParserREF)
	}
	{
		p.SetState(134)

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
		p.SetState(136)
		p.Match(DZParserLEFT_BRK)
	}
	{
		p.SetState(137)

		var _x = p.SimpleTypeSpec()

		localctx.(*ArrayTypeSpecContext).tName = _x
	}
	{
		p.SetState(138)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(139)

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
		p.SetState(140)
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
		p.SetState(142)
		p.Match(DZParserLEFT_BRK)
	}
	{
		p.SetState(143)

		var _x = p.SimpleTypeSpec()

		localctx.(*SliceTypeSpecContext).tName = _x
	}
	{
		p.SetState(144)
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
		p.SetState(146)
		p.Match(DZParserKW_ENUM)
	}
	{
		p.SetState(147)

		var _m = p.Match(DZParserTYPE)

		localctx.(*EnumDeclContext).name = _m
	}
	{
		p.SetState(148)
		p.Match(DZParserLEFT_BRC)
	}
	{
		p.SetState(149)
		p.EnumOption()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserCOMMA {
		{
			p.SetState(150)
			p.Match(DZParserCOMMA)
		}
		{
			p.SetState(151)
			p.EnumOption()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
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
		p.SetState(159)

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
		p.SetState(161)
		p.Match(DZParserKW_STRUCT)
	}
	{
		p.SetState(162)

		var _m = p.Match(DZParserTYPE)

		localctx.(*StructDeclContext).name = _m
	}
	{
		p.SetState(163)
		p.Match(DZParserLEFT_BRC)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserIDENTIFIER {
		{
			p.SetState(164)
			p.StructAttr()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(170)
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
		p.SetState(172)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*StructAttrContext).name = _m
	}
	{
		p.SetState(173)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(174)

		var _x = p.TypeSpec()

		localctx.(*StructAttrContext).tName = _x
	}
	{
		p.SetState(175)
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
		p.SetState(177)
		p.Match(DZParserKW_PROC)
	}
	{
		p.SetState(178)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcDeclContext).name = _m
	}
	{
		p.SetState(179)
		p.Match(DZParserLEFT_PRT)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserIDENTIFIER {
		{
			p.SetState(180)
			p.ProcArg()
		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserCOMMA {
			{
				p.SetState(181)
				p.Match(DZParserCOMMA)
			}
			{
				p.SetState(182)
				p.ProcArg()
			}

			p.SetState(187)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(190)
		p.Match(DZParserRIGHT_PRT)
	}
	{
		p.SetState(191)

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
		p.SetState(193)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcArgContext).name = _m
	}
	{
		p.SetState(194)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(195)

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

func (s *FuncDeclContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_PRT, 0)
}

func (s *FuncDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
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
		p.SetState(197)
		p.Match(DZParserKW_FUNC)
	}
	{
		p.SetState(198)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*FuncDeclContext).name = _m
	}
	{
		p.SetState(199)
		p.Match(DZParserLEFT_PRT)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserIDENTIFIER {
		{
			p.SetState(200)
			p.FuncArg()
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserCOMMA {
			{
				p.SetState(201)
				p.Match(DZParserCOMMA)
			}
			{
				p.SetState(202)
				p.FuncArg()
			}

			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(210)
		p.Match(DZParserRIGHT_PRT)
	}
	{
		p.SetState(211)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(212)

		var _x = p.TypeSpec()

		localctx.(*FuncDeclContext).tName = _x
	}
	{
		p.SetState(213)

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
		p.SetState(215)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*FuncArgContext).name = _m
	}
	{
		p.SetState(216)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(217)

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

func (s *BlockContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *BlockContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

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
	{
		p.SetState(219)
		p.Match(DZParserLEFT_BRC)
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DZParserKW_WHILE)|(1<<DZParserKW_LOOP)|(1<<DZParserKW_BREAK)|(1<<DZParserKW_CONTINUE)|(1<<DZParserKW_IF)|(1<<DZParserKW_LET)|(1<<DZParserKW_RETURN)|(1<<DZParserKW_TRUE)|(1<<DZParserKW_FALSE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(DZParserLEFT_PRT-32))|(1<<(DZParserLEFT_BRC-32))|(1<<(DZParserLEFT_BRK-32))|(1<<(DZParserSUB-32))|(1<<(DZParserNOT-32)))) != 0) || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(DZParserCONST-71))|(1<<(DZParserINT_VALUE-71))|(1<<(DZParserIDENTIFIER-71)))) != 0) {
		{
			p.SetState(220)
			p.Statement()
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(226)
		p.Match(DZParserRIGHT_BRC)
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

	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Condition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Loop()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.BreakStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
			p.ContinueStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(232)
			p.ProcCall()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(233)
			p.VarDecl()
		}
		{
			p.SetState(234)
			p.Match(DZParserSEMICOLON)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(236)
			p.expression(0)
		}
		{
			p.SetState(237)
			p.Match(DZParserSEMICOLON)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(239)
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
		p.SetState(242)
		p.IfConditionBranch()
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DZParserKW_ELIF {
		{
			p.SetState(243)
			p.ElifConditionBranch()
		}

		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserKW_ELSE {
		{
			p.SetState(249)
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
		p.SetState(252)
		p.Match(DZParserKW_IF)
	}
	{
		p.SetState(253)

		var _x = p.expression(0)

		localctx.(*IfConditionBranchContext).cond = _x
	}
	{
		p.SetState(254)

		var _x = p.Block()

		localctx.(*IfConditionBranchContext).body = _x
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
		p.SetState(256)
		p.Match(DZParserKW_ELIF)
	}
	{
		p.SetState(257)

		var _x = p.expression(0)

		localctx.(*ElifConditionBranchContext).cond = _x
	}
	{
		p.SetState(258)

		var _x = p.Block()

		localctx.(*ElifConditionBranchContext).body = _x
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
		p.SetState(260)
		p.Match(DZParserKW_ELSE)
	}
	{
		p.SetState(261)

		var _x = p.Block()

		localctx.(*ElseConditionBranchContext).body = _x
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

	p.SetState(265)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DZParserKW_LOOP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.TrueLoop()
		}

	case DZParserKW_WHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
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
		p.SetState(267)
		p.Match(DZParserKW_LOOP)
	}
	{
		p.SetState(268)

		var _x = p.Block()

		localctx.(*TrueLoopContext).body = _x
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
		p.SetState(270)
		p.Match(DZParserKW_WHILE)
	}
	{
		p.SetState(271)

		var _x = p.expression(0)

		localctx.(*WhileLoopContext).cond = _x
	}
	{
		p.SetState(272)

		var _x = p.Block()

		localctx.(*WhileLoopContext).body = _x
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

func (s *BreakStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
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
		p.SetState(274)
		p.Match(DZParserKW_BREAK)
	}
	{
		p.SetState(275)
		p.Match(DZParserSEMICOLON)
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

func (s *ContinueStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(DZParserSEMICOLON, 0)
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
	{
		p.SetState(278)
		p.Match(DZParserSEMICOLON)
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

	// GetTName returns the tName rule contexts.
	GetTName() ITypeSpecContext

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetTName sets the tName rule contexts.
	SetTName(ITypeSpecContext)

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	tName  ITypeSpecContext
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

func (s *VarDeclContext) GetTName() ITypeSpecContext { return s.tName }

func (s *VarDeclContext) GetValue() IExpressionContext { return s.value }

func (s *VarDeclContext) SetTName(v ITypeSpecContext) { s.tName = v }

func (s *VarDeclContext) SetValue(v IExpressionContext) { s.value = v }

func (s *VarDeclContext) KW_LET() antlr.TerminalNode {
	return s.GetToken(DZParserKW_LET, 0)
}

func (s *VarDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(DZParserCOLON, 0)
}

func (s *VarDeclContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *VarDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *VarDeclContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
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
		p.SetState(280)
		p.Match(DZParserKW_LET)
	}
	{
		p.SetState(281)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*VarDeclContext).name = _m
	}
	{
		p.SetState(282)
		p.Match(DZParserCOLON)
	}
	{
		p.SetState(283)

		var _x = p.TypeSpec()

		localctx.(*VarDeclContext).tName = _x
	}
	{
		p.SetState(284)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(285)

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
		p.SetState(287)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcCallContext).procName = _m
	}
	{
		p.SetState(288)
		p.Match(DZParserLEFT_PRT)
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserIDENTIFIER {
		{
			p.SetState(289)
			p.ProcParam()
		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserCOMMA {
			{
				p.SetState(290)
				p.Match(DZParserCOMMA)
			}
			{
				p.SetState(291)
				p.ProcParam()
			}

			p.SetState(296)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(299)
		p.Match(DZParserRIGHT_PRT)
	}
	{
		p.SetState(300)
		p.Match(DZParserSEMICOLON)
	}

	return localctx
}

// IProcParamContext is an interface to support dynamic dispatch.
type IProcParamContext interface {
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

	// IsProcParamContext differentiates from other interfaces.
	IsProcParamContext()
}

type ProcParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
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

func (s *ProcParamContext) GetName() antlr.Token { return s.name }

func (s *ProcParamContext) SetName(v antlr.Token) { s.name = v }

func (s *ProcParamContext) GetValue() IExpressionContext { return s.value }

func (s *ProcParamContext) SetValue(v IExpressionContext) { s.value = v }

func (s *ProcParamContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *ProcParamContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

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
		p.SetState(302)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*ProcParamContext).name = _m
	}
	{
		p.SetState(303)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(304)

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

	// GetConstName returns the constName token.
	GetConstName() antlr.Token

	// GetAttrName returns the attrName token.
	GetAttrName() antlr.Token

	// SetVarName sets the varName token.
	SetVarName(antlr.Token)

	// SetConstName sets the constName token.
	SetConstName(antlr.Token)

	// SetAttrName sets the attrName token.
	SetAttrName(antlr.Token)

	// GetVarExpr returns the varExpr rule contexts.
	GetVarExpr() IExpressionContext

	// GetLBinExpr returns the lBinExpr rule contexts.
	GetLBinExpr() IExpressionContext

	// GetConstVal returns the constVal rule contexts.
	GetConstVal() IConstValueContext

	// GetPrtExpr returns the prtExpr rule contexts.
	GetPrtExpr() IExpressionContext

	// GetBrkExpr returns the brkExpr rule contexts.
	GetBrkExpr() IExpressionContext

	// GetBrcExpr returns the brcExpr rule contexts.
	GetBrcExpr() IAttrAsgnContext

	// GetFuncCallValue returns the funcCallValue rule contexts.
	GetFuncCallValue() IFuncCallContext

	// GetUnOp returns the unOp rule contexts.
	GetUnOp() IUnaryOperatorContext

	// GetUnExpr returns the unExpr rule contexts.
	GetUnExpr() IExpressionContext

	// GetBinOp returns the binOp rule contexts.
	GetBinOp() IBinaryOperatorContext

	// GetRBinExpr returns the rBinExpr rule contexts.
	GetRBinExpr() IExpressionContext

	// SetVarExpr sets the varExpr rule contexts.
	SetVarExpr(IExpressionContext)

	// SetLBinExpr sets the lBinExpr rule contexts.
	SetLBinExpr(IExpressionContext)

	// SetConstVal sets the constVal rule contexts.
	SetConstVal(IConstValueContext)

	// SetPrtExpr sets the prtExpr rule contexts.
	SetPrtExpr(IExpressionContext)

	// SetBrkExpr sets the brkExpr rule contexts.
	SetBrkExpr(IExpressionContext)

	// SetBrcExpr sets the brcExpr rule contexts.
	SetBrcExpr(IAttrAsgnContext)

	// SetFuncCallValue sets the funcCallValue rule contexts.
	SetFuncCallValue(IFuncCallContext)

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
	parser        antlr.Parser
	varExpr       IExpressionContext
	lBinExpr      IExpressionContext
	varName       antlr.Token
	constVal      IConstValueContext
	constName     antlr.Token
	prtExpr       IExpressionContext
	brkExpr       IExpressionContext
	brcExpr       IAttrAsgnContext
	funcCallValue IFuncCallContext
	unOp          IUnaryOperatorContext
	unExpr        IExpressionContext
	binOp         IBinaryOperatorContext
	rBinExpr      IExpressionContext
	attrName      antlr.Token
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

func (s *ExpressionContext) GetConstName() antlr.Token { return s.constName }

func (s *ExpressionContext) GetAttrName() antlr.Token { return s.attrName }

func (s *ExpressionContext) SetVarName(v antlr.Token) { s.varName = v }

func (s *ExpressionContext) SetConstName(v antlr.Token) { s.constName = v }

func (s *ExpressionContext) SetAttrName(v antlr.Token) { s.attrName = v }

func (s *ExpressionContext) GetVarExpr() IExpressionContext { return s.varExpr }

func (s *ExpressionContext) GetLBinExpr() IExpressionContext { return s.lBinExpr }

func (s *ExpressionContext) GetConstVal() IConstValueContext { return s.constVal }

func (s *ExpressionContext) GetPrtExpr() IExpressionContext { return s.prtExpr }

func (s *ExpressionContext) GetBrkExpr() IExpressionContext { return s.brkExpr }

func (s *ExpressionContext) GetBrcExpr() IAttrAsgnContext { return s.brcExpr }

func (s *ExpressionContext) GetFuncCallValue() IFuncCallContext { return s.funcCallValue }

func (s *ExpressionContext) GetUnOp() IUnaryOperatorContext { return s.unOp }

func (s *ExpressionContext) GetUnExpr() IExpressionContext { return s.unExpr }

func (s *ExpressionContext) GetBinOp() IBinaryOperatorContext { return s.binOp }

func (s *ExpressionContext) GetRBinExpr() IExpressionContext { return s.rBinExpr }

func (s *ExpressionContext) SetVarExpr(v IExpressionContext) { s.varExpr = v }

func (s *ExpressionContext) SetLBinExpr(v IExpressionContext) { s.lBinExpr = v }

func (s *ExpressionContext) SetConstVal(v IConstValueContext) { s.constVal = v }

func (s *ExpressionContext) SetPrtExpr(v IExpressionContext) { s.prtExpr = v }

func (s *ExpressionContext) SetBrkExpr(v IExpressionContext) { s.brkExpr = v }

func (s *ExpressionContext) SetBrcExpr(v IAttrAsgnContext) { s.brcExpr = v }

func (s *ExpressionContext) SetFuncCallValue(v IFuncCallContext) { s.funcCallValue = v }

func (s *ExpressionContext) SetUnOp(v IUnaryOperatorContext) { s.unOp = v }

func (s *ExpressionContext) SetUnExpr(v IExpressionContext) { s.unExpr = v }

func (s *ExpressionContext) SetBinOp(v IBinaryOperatorContext) { s.binOp = v }

func (s *ExpressionContext) SetRBinExpr(v IExpressionContext) { s.rBinExpr = v }

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *ExpressionContext) ConstValue() IConstValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstValueContext)
}

func (s *ExpressionContext) CONST() antlr.TerminalNode {
	return s.GetToken(DZParserCONST, 0)
}

func (s *ExpressionContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_PRT, 0)
}

func (s *ExpressionContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_PRT, 0)
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

func (s *ExpressionContext) LEFT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRK, 0)
}

func (s *ExpressionContext) RIGHT_BRK() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRK, 0)
}

func (s *ExpressionContext) LEFT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_BRC, 0)
}

func (s *ExpressionContext) RIGHT_BRC() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_BRC, 0)
}

func (s *ExpressionContext) AllAttrAsgn() []IAttrAsgnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttrAsgnContext)(nil)).Elem())
	var tst = make([]IAttrAsgnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttrAsgnContext)
		}
	}

	return tst
}

func (s *ExpressionContext) AttrAsgn(i int) IAttrAsgnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrAsgnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttrAsgnContext)
}

func (s *ExpressionContext) FuncCall() IFuncCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
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

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(DZParserDOT, 0)
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(307)

			var _m = p.Match(DZParserIDENTIFIER)

			localctx.(*ExpressionContext).varName = _m
		}

	case 2:
		{
			p.SetState(308)

			var _x = p.ConstValue()

			localctx.(*ExpressionContext).constVal = _x
		}

	case 3:
		{
			p.SetState(309)

			var _m = p.Match(DZParserCONST)

			localctx.(*ExpressionContext).constName = _m
		}

	case 4:
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

	case 5:
		{
			p.SetState(314)
			p.Match(DZParserLEFT_BRK)
		}
		{
			p.SetState(315)

			var _x = p.expression(0)

			localctx.(*ExpressionContext).brkExpr = _x
		}
		{
			p.SetState(316)
			p.Match(DZParserRIGHT_BRK)
		}

	case 6:
		{
			p.SetState(318)
			p.Match(DZParserLEFT_BRC)
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DZParserDOT {
			{
				p.SetState(319)

				var _x = p.AttrAsgn()

				localctx.(*ExpressionContext).brcExpr = _x
			}

			p.SetState(324)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(325)
			p.Match(DZParserRIGHT_BRC)
		}

	case 7:
		{
			p.SetState(326)

			var _x = p.FuncCall()

			localctx.(*ExpressionContext).funcCallValue = _x
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
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(339)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
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

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).varExpr = _prevctx
				p.PushNewRecursionContext(localctx, _startState, DZParserRULE_expression)
				p.SetState(336)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(337)
					p.Match(DZParserDOT)
				}
				{
					p.SetState(338)

					var _m = p.Match(DZParserIDENTIFIER)

					localctx.(*ExpressionContext).attrName = _m
				}

			}

		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IAttrAsgnContext is an interface to support dynamic dispatch.
type IAttrAsgnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAttrName returns the attrName token.
	GetAttrName() antlr.Token

	// SetAttrName sets the attrName token.
	SetAttrName(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsAttrAsgnContext differentiates from other interfaces.
	IsAttrAsgnContext()
}

type AttrAsgnContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	attrName antlr.Token
	value    IExpressionContext
}

func NewEmptyAttrAsgnContext() *AttrAsgnContext {
	var p = new(AttrAsgnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_attrAsgn
	return p
}

func (*AttrAsgnContext) IsAttrAsgnContext() {}

func NewAttrAsgnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrAsgnContext {
	var p = new(AttrAsgnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_attrAsgn

	return p
}

func (s *AttrAsgnContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrAsgnContext) GetAttrName() antlr.Token { return s.attrName }

func (s *AttrAsgnContext) SetAttrName(v antlr.Token) { s.attrName = v }

func (s *AttrAsgnContext) GetValue() IExpressionContext { return s.value }

func (s *AttrAsgnContext) SetValue(v IExpressionContext) { s.value = v }

func (s *AttrAsgnContext) DOT() antlr.TerminalNode {
	return s.GetToken(DZParserDOT, 0)
}

func (s *AttrAsgnContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *AttrAsgnContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *AttrAsgnContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AttrAsgnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrAsgnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrAsgnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterAttrAsgn(s)
	}
}

func (s *AttrAsgnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitAttrAsgn(s)
	}
}

func (p *DZParser) AttrAsgn() (localctx IAttrAsgnContext) {
	localctx = NewAttrAsgnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, DZParserRULE_attrAsgn)

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
		p.SetState(344)
		p.Match(DZParserDOT)
	}
	{
		p.SetState(345)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*AttrAsgnContext).attrName = _m
	}
	{
		p.SetState(346)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(347)

		var _x = p.expression(0)

		localctx.(*AttrAsgnContext).value = _x
	}

	return localctx
}

// IConstValueContext is an interface to support dynamic dispatch.
type IConstValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstValueContext differentiates from other interfaces.
	IsConstValueContext()
}

type ConstValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstValueContext() *ConstValueContext {
	var p = new(ConstValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_constValue
	return p
}

func (*ConstValueContext) IsConstValueContext() {}

func NewConstValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstValueContext {
	var p = new(ConstValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_constValue

	return p
}

func (s *ConstValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstValueContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(DZParserINT_VALUE, 0)
}

func (s *ConstValueContext) KW_TRUE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_TRUE, 0)
}

func (s *ConstValueContext) KW_FALSE() antlr.TerminalNode {
	return s.GetToken(DZParserKW_FALSE, 0)
}

func (s *ConstValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterConstValue(s)
	}
}

func (s *ConstValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitConstValue(s)
	}
}

func (p *DZParser) ConstValue() (localctx IConstValueContext) {
	localctx = NewConstValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, DZParserRULE_constValue)
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
		p.SetState(349)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DZParserKW_TRUE || _la == DZParserKW_FALSE || _la == DZParserINT_VALUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFuncName returns the funcName token.
	GetFuncName() antlr.Token

	// SetFuncName sets the funcName token.
	SetFuncName(antlr.Token)

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	funcName antlr.Token
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DZParserRULE_funcCall
	return p
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DZParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) GetFuncName() antlr.Token { return s.funcName }

func (s *FuncCallContext) SetFuncName(v antlr.Token) { s.funcName = v }

func (s *FuncCallContext) LEFT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserLEFT_PRT, 0)
}

func (s *FuncCallContext) RIGHT_PRT() antlr.TerminalNode {
	return s.GetToken(DZParserRIGHT_PRT, 0)
}

func (s *FuncCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

func (s *FuncCallContext) AllFuncParam() []IFuncParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncParamContext)(nil)).Elem())
	var tst = make([]IFuncParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncParamContext)
		}
	}

	return tst
}

func (s *FuncCallContext) FuncParam(i int) IFuncParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncParamContext)
}

func (s *FuncCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DZParserCOMMA)
}

func (s *FuncCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DZParserCOMMA, i)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DZListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (p *DZParser) FuncCall() (localctx IFuncCallContext) {
	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, DZParserRULE_funcCall)
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
		p.SetState(351)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*FuncCallContext).funcName = _m
	}
	{
		p.SetState(352)
		p.Match(DZParserLEFT_PRT)
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DZParserIDENTIFIER {
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

	return localctx
}

// IFuncParamContext is an interface to support dynamic dispatch.
type IFuncParamContext interface {
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

	// IsFuncParamContext differentiates from other interfaces.
	IsFuncParamContext()
}

type FuncParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
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

func (s *FuncParamContext) GetName() antlr.Token { return s.name }

func (s *FuncParamContext) SetName(v antlr.Token) { s.name = v }

func (s *FuncParamContext) GetValue() IExpressionContext { return s.value }

func (s *FuncParamContext) SetValue(v IExpressionContext) { s.value = v }

func (s *FuncParamContext) ASGN() antlr.TerminalNode {
	return s.GetToken(DZParserASGN, 0)
}

func (s *FuncParamContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DZParserIDENTIFIER, 0)
}

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
	p.EnterRule(localctx, 76, DZParserRULE_funcParam)

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
		p.SetState(365)

		var _m = p.Match(DZParserIDENTIFIER)

		localctx.(*FuncParamContext).name = _m
	}
	{
		p.SetState(366)
		p.Match(DZParserASGN)
	}
	{
		p.SetState(367)

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
	p.EnterRule(localctx, 78, DZParserRULE_unaryOperator)
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
		p.SetState(369)
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
	p.EnterRule(localctx, 80, DZParserRULE_binaryOperator)
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
		p.SetState(371)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(DZParserADD-42))|(1<<(DZParserSUB-42))|(1<<(DZParserMUL-42))|(1<<(DZParserDIV-42))|(1<<(DZParserMOD-42))|(1<<(DZParserSHL-42))|(1<<(DZParserSHR-42))|(1<<(DZParserAND-42))|(1<<(DZParserOR-42))|(1<<(DZParserXOR-42))|(1<<(DZParserEQL-42))|(1<<(DZParserGRT-42))|(1<<(DZParserGRT_EQL-42))|(1<<(DZParserLESS-42))|(1<<(DZParserLESS_EQL-42))|(1<<(DZParserASGN-42))|(1<<(DZParserADD_ASGN-42))|(1<<(DZParserSUB_ASGN-42))|(1<<(DZParserMUL_ASGN-42))|(1<<(DZParserDIV_ASGN-42))|(1<<(DZParserMOD_ASGN-42))|(1<<(DZParserSHL_ASGN-42))|(1<<(DZParserSHR_ASGN-42))|(1<<(DZParserAND_ASGN-42))|(1<<(DZParserOR_ASGN-42))|(1<<(DZParserXOR_ASGN-42)))) != 0) {
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
	p.EnterRule(localctx, 82, DZParserRULE_returnStatement)

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

	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(373)
			p.Match(DZParserKW_RETURN)
		}
		{
			p.SetState(374)
			p.Match(DZParserSEMICOLON)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Match(DZParserKW_RETURN)
		}
		{
			p.SetState(376)

			var _x = p.expression(0)

			localctx.(*ReturnStatementContext).value = _x
		}
		{
			p.SetState(377)
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

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
