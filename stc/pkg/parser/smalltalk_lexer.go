// Code generated from /home/denis/Documents/iu7-project-compilers/stc/SmallTalk.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 37, 239,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 118, 10, 12, 12, 12, 14, 12, 121,
	11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 6, 13, 128, 10, 13, 13, 13,
	14, 13, 129, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 7, 25, 159, 10,
	25, 12, 25, 14, 25, 162, 11, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 5, 30,
	179, 10, 30, 3, 30, 6, 30, 182, 10, 30, 13, 30, 14, 30, 183, 3, 31, 3,
	31, 5, 31, 188, 10, 31, 3, 31, 6, 31, 191, 10, 31, 13, 31, 14, 31, 192,
	3, 32, 5, 32, 196, 10, 32, 3, 32, 6, 32, 199, 10, 32, 13, 32, 14, 32, 200,
	3, 33, 3, 33, 5, 33, 205, 10, 33, 3, 33, 6, 33, 208, 10, 33, 13, 33, 14,
	33, 209, 3, 34, 5, 34, 213, 10, 34, 3, 34, 6, 34, 216, 10, 34, 13, 34,
	14, 34, 217, 3, 34, 3, 34, 6, 34, 222, 10, 34, 13, 34, 14, 34, 223, 3,
	35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 7, 36, 233, 10, 36, 12, 36,
	14, 36, 236, 11, 36, 3, 36, 3, 36, 2, 2, 37, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25,
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34,
	67, 35, 69, 36, 71, 37, 3, 2, 13, 3, 2, 36, 36, 5, 2, 11, 12, 15, 15, 34,
	34, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2,
	47, 47, 3, 2, 50, 51, 3, 2, 50, 57, 3, 2, 50, 59, 4, 2, 50, 59, 67, 72,
	5, 2, 11, 12, 34, 34, 66, 66, 3, 2, 41, 41, 2, 255, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2,
	2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3,
	2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51,
	3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2,
	59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2,
	2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 3, 73, 3, 2, 2,
	2, 5, 75, 3, 2, 2, 2, 7, 77, 3, 2, 2, 2, 9, 79, 3, 2, 2, 2, 11, 81, 3,
	2, 2, 2, 13, 85, 3, 2, 2, 2, 15, 90, 3, 2, 2, 2, 17, 96, 3, 2, 2, 2, 19,
	102, 3, 2, 2, 2, 21, 107, 3, 2, 2, 2, 23, 113, 3, 2, 2, 2, 25, 127, 3,
	2, 2, 2, 27, 133, 3, 2, 2, 2, 29, 135, 3, 2, 2, 2, 31, 137, 3, 2, 2, 2,
	33, 139, 3, 2, 2, 2, 35, 141, 3, 2, 2, 2, 37, 144, 3, 2, 2, 2, 39, 146,
	3, 2, 2, 2, 41, 148, 3, 2, 2, 2, 43, 150, 3, 2, 2, 2, 45, 152, 3, 2, 2,
	2, 47, 154, 3, 2, 2, 2, 49, 156, 3, 2, 2, 2, 51, 163, 3, 2, 2, 2, 53, 166,
	3, 2, 2, 2, 55, 169, 3, 2, 2, 2, 57, 172, 3, 2, 2, 2, 59, 176, 3, 2, 2,
	2, 61, 185, 3, 2, 2, 2, 63, 195, 3, 2, 2, 2, 65, 202, 3, 2, 2, 2, 67, 212,
	3, 2, 2, 2, 69, 225, 3, 2, 2, 2, 71, 228, 3, 2, 2, 2, 73, 74, 7, 45, 2,
	2, 74, 4, 3, 2, 2, 2, 75, 76, 7, 47, 2, 2, 76, 6, 3, 2, 2, 2, 77, 78, 7,
	44, 2, 2, 78, 8, 3, 2, 2, 2, 79, 80, 7, 49, 2, 2, 80, 10, 3, 2, 2, 2, 81,
	82, 7, 112, 2, 2, 82, 83, 7, 107, 2, 2, 83, 84, 7, 110, 2, 2, 84, 12, 3,
	2, 2, 2, 85, 86, 7, 118, 2, 2, 86, 87, 7, 116, 2, 2, 87, 88, 7, 119, 2,
	2, 88, 89, 7, 103, 2, 2, 89, 14, 3, 2, 2, 2, 90, 91, 7, 104, 2, 2, 91,
	92, 7, 99, 2, 2, 92, 93, 7, 110, 2, 2, 93, 94, 7, 117, 2, 2, 94, 95, 7,
	103, 2, 2, 95, 16, 3, 2, 2, 2, 96, 97, 7, 101, 2, 2, 97, 98, 7, 110, 2,
	2, 98, 99, 7, 99, 2, 2, 99, 100, 7, 117, 2, 2, 100, 101, 7, 117, 2, 2,
	101, 18, 3, 2, 2, 2, 102, 103, 7, 117, 2, 2, 103, 104, 7, 103, 2, 2, 104,
	105, 7, 110, 2, 2, 105, 106, 7, 104, 2, 2, 106, 20, 3, 2, 2, 2, 107, 108,
	7, 117, 2, 2, 108, 109, 7, 119, 2, 2, 109, 110, 7, 114, 2, 2, 110, 111,
	7, 103, 2, 2, 111, 112, 7, 116, 2, 2, 112, 22, 3, 2, 2, 2, 113, 119, 7,
	36, 2, 2, 114, 115, 7, 36, 2, 2, 115, 118, 7, 36, 2, 2, 116, 118, 10, 2,
	2, 2, 117, 114, 3, 2, 2, 2, 117, 116, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2,
	119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 122, 3, 2, 2, 2, 121,
	119, 3, 2, 2, 2, 122, 123, 7, 36, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125,
	8, 12, 2, 2, 125, 24, 3, 2, 2, 2, 126, 128, 9, 3, 2, 2, 127, 126, 3, 2,
	2, 2, 128, 129, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2,
	130, 131, 3, 2, 2, 2, 131, 132, 8, 13, 2, 2, 132, 26, 3, 2, 2, 2, 133,
	134, 7, 48, 2, 2, 134, 28, 3, 2, 2, 2, 135, 136, 7, 126, 2, 2, 136, 30,
	3, 2, 2, 2, 137, 138, 7, 60, 2, 2, 138, 32, 3, 2, 2, 2, 139, 140, 7, 61,
	2, 2, 140, 34, 3, 2, 2, 2, 141, 142, 7, 60, 2, 2, 142, 143, 7, 63, 2, 2,
	143, 36, 3, 2, 2, 2, 144, 145, 7, 93, 2, 2, 145, 38, 3, 2, 2, 2, 146, 147,
	7, 95, 2, 2, 147, 40, 3, 2, 2, 2, 148, 149, 7, 42, 2, 2, 149, 42, 3, 2,
	2, 2, 150, 151, 7, 43, 2, 2, 151, 44, 3, 2, 2, 2, 152, 153, 7, 125, 2,
	2, 153, 46, 3, 2, 2, 2, 154, 155, 7, 127, 2, 2, 155, 48, 3, 2, 2, 2, 156,
	160, 9, 4, 2, 2, 157, 159, 9, 5, 2, 2, 158, 157, 3, 2, 2, 2, 159, 162,
	3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 50, 3, 2,
	2, 2, 162, 160, 3, 2, 2, 2, 163, 164, 5, 49, 25, 2, 164, 165, 5, 31, 16,
	2, 165, 52, 3, 2, 2, 2, 166, 167, 7, 52, 2, 2, 167, 168, 7, 116, 2, 2,
	168, 54, 3, 2, 2, 2, 169, 170, 7, 58, 2, 2, 170, 171, 7, 116, 2, 2, 171,
	56, 3, 2, 2, 2, 172, 173, 7, 51, 2, 2, 173, 174, 7, 56, 2, 2, 174, 175,
	7, 116, 2, 2, 175, 58, 3, 2, 2, 2, 176, 178, 5, 53, 27, 2, 177, 179, 9,
	6, 2, 2, 178, 177, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 3, 2, 2,
	2, 180, 182, 9, 7, 2, 2, 181, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183,
	181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 60, 3, 2, 2, 2, 185, 187, 5,
	55, 28, 2, 186, 188, 9, 6, 2, 2, 187, 186, 3, 2, 2, 2, 187, 188, 3, 2,
	2, 2, 188, 190, 3, 2, 2, 2, 189, 191, 9, 8, 2, 2, 190, 189, 3, 2, 2, 2,
	191, 192, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193,
	62, 3, 2, 2, 2, 194, 196, 9, 6, 2, 2, 195, 194, 3, 2, 2, 2, 195, 196, 3,
	2, 2, 2, 196, 198, 3, 2, 2, 2, 197, 199, 9, 9, 2, 2, 198, 197, 3, 2, 2,
	2, 199, 200, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201,
	64, 3, 2, 2, 2, 202, 204, 5, 57, 29, 2, 203, 205, 9, 6, 2, 2, 204, 203,
	3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 207, 3, 2, 2, 2, 206, 208, 9, 10,
	2, 2, 207, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2,
	209, 210, 3, 2, 2, 2, 210, 66, 3, 2, 2, 2, 211, 213, 9, 6, 2, 2, 212, 211,
	3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 215, 3, 2, 2, 2, 214, 216, 9, 9,
	2, 2, 215, 214, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2,
	217, 218, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 221, 5, 27, 14, 2, 220,
	222, 9, 9, 2, 2, 221, 220, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 221,
	3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 68, 3, 2, 2, 2, 225, 226, 7, 38,
	2, 2, 226, 227, 10, 11, 2, 2, 227, 70, 3, 2, 2, 2, 228, 234, 7, 41, 2,
	2, 229, 230, 7, 41, 2, 2, 230, 233, 7, 41, 2, 2, 231, 233, 10, 12, 2, 2,
	232, 229, 3, 2, 2, 2, 232, 231, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234,
	232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 237, 3, 2, 2, 2, 236, 234,
	3, 2, 2, 2, 237, 238, 7, 41, 2, 2, 238, 72, 3, 2, 2, 2, 20, 2, 117, 119,
	129, 160, 178, 183, 187, 192, 195, 200, 204, 209, 212, 217, 223, 232, 234,
	3, 8, 2, 2,
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
	"", "'+'", "'-'", "'*'", "'/'", "'nil'", "'true'", "'false'", "'class'",
	"'self'", "'super'", "", "", "'.'", "'|'", "':'", "';'", "':='", "'['",
	"']'", "'('", "')'", "'{'", "'}'", "", "", "'2r'", "'8r'", "'16r'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "NIL", "TRUE", "FALSE", "CLASS", "SELF", "SUPER", "COMMENT",
	"WHITESPACE", "DOT", "PIPE", "COLON", "SEMICOLON", "ASSIGNMENT", "LEFT_BRK",
	"RIGHT_BRK", "LEFT_PRT", "RIGHT_PRT", "LEFT_BRC", "RIGHT_BRC", "IDENTIFIER",
	"KEYWORD", "BIN", "OCT", "HEX", "BIN_INT_NUMBER", "OCT_INT_NUMBER", "DEC_INT_NUMBER",
	"HEX_INT_NUMBER", "FLOAT_NUMBER", "CHAR", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "NIL", "TRUE", "FALSE", "CLASS", "SELF",
	"SUPER", "COMMENT", "WHITESPACE", "DOT", "PIPE", "COLON", "SEMICOLON",
	"ASSIGNMENT", "LEFT_BRK", "RIGHT_BRK", "LEFT_PRT", "RIGHT_PRT", "LEFT_BRC",
	"RIGHT_BRC", "IDENTIFIER", "KEYWORD", "BIN", "OCT", "HEX", "BIN_INT_NUMBER",
	"OCT_INT_NUMBER", "DEC_INT_NUMBER", "HEX_INT_NUMBER", "FLOAT_NUMBER", "CHAR",
	"STRING",
}

type SmallTalkLexer struct {
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

func NewSmallTalkLexer(input antlr.CharStream) *SmallTalkLexer {

	l := new(SmallTalkLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SmallTalk.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SmallTalkLexer tokens.
const (
	SmallTalkLexerT__0           = 1
	SmallTalkLexerT__1           = 2
	SmallTalkLexerT__2           = 3
	SmallTalkLexerT__3           = 4
	SmallTalkLexerNIL            = 5
	SmallTalkLexerTRUE           = 6
	SmallTalkLexerFALSE          = 7
	SmallTalkLexerCLASS          = 8
	SmallTalkLexerSELF           = 9
	SmallTalkLexerSUPER          = 10
	SmallTalkLexerCOMMENT        = 11
	SmallTalkLexerWHITESPACE     = 12
	SmallTalkLexerDOT            = 13
	SmallTalkLexerPIPE           = 14
	SmallTalkLexerCOLON          = 15
	SmallTalkLexerSEMICOLON      = 16
	SmallTalkLexerASSIGNMENT     = 17
	SmallTalkLexerLEFT_BRK       = 18
	SmallTalkLexerRIGHT_BRK      = 19
	SmallTalkLexerLEFT_PRT       = 20
	SmallTalkLexerRIGHT_PRT      = 21
	SmallTalkLexerLEFT_BRC       = 22
	SmallTalkLexerRIGHT_BRC      = 23
	SmallTalkLexerIDENTIFIER     = 24
	SmallTalkLexerKEYWORD        = 25
	SmallTalkLexerBIN            = 26
	SmallTalkLexerOCT            = 27
	SmallTalkLexerHEX            = 28
	SmallTalkLexerBIN_INT_NUMBER = 29
	SmallTalkLexerOCT_INT_NUMBER = 30
	SmallTalkLexerDEC_INT_NUMBER = 31
	SmallTalkLexerHEX_INT_NUMBER = 32
	SmallTalkLexerFLOAT_NUMBER   = 33
	SmallTalkLexerCHAR           = 34
	SmallTalkLexerSTRING         = 35
)