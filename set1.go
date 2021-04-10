package main

import (
	"bufio"
	b64 "encoding/base64"
	hex "encoding/hex"
	"fmt"
	"os"
	"strings"
)

const LOWER_LETTERS string = "abcdefghijklmnopqrstuvwxyz"
const UPPER_LETTERS string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LETTERS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NUMBERS string = "0123456789"
const PUNCTUATIONS string = ".!?,'\""
const LETTER_COUNT = 26

var AVERAGE_FREQS = map[rune]float64{
	'a': 0.08167,
	'b': 0.01492,
	'c': 0.02782,
	'd': 0.04253,
	'e': 0.12702,
	'f': 0.02228,
	'g': 0.02015,
	'h': 0.06094,
	'i': 0.06966,
	'j': 0.00153,
	'k': 0.00772,
	'l': 0.04025,
	'm': 0.02406,
	'n': 0.06749,
	'o': 0.07507,
	'p': 0.01929,
	'q': 0.00095,
	'r': 0.05987,
	's': 0.06327,
	't': 0.09056,
	'u': 0.02758,
	'v': 0.00978,
	'w': 0.02360,
	'x': 0.00150,
	'y': 0.01974,
	'z': 0.00074,
}

func main() {
	fmt.Println(AVERAGE_FREQS)
}

func getletterdistributions(str string) {
	str = strings.ToLower(str)
	// freqs := generateEmptuLetterFreqMap();
	for i := 0; i < len(str); i++ {
		// freqs[];
	}
}


func generateEmptuLetterFreqMap() map[rune]float64 {
	var freqs = map[rune]float64{}
	const start rune = 'a'
	var i rune = 0
	for i = 0; i < LETTER_COUNT; i++ {
		freqs[start+i] = 0.0
	}
	return freqs
}
func challenge1() {
	var line string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes, _ := hex.DecodeString(line)
	encoded64 := b64.RawStdEncoding.EncodeToString(bytes)
	fmt.Println(bytes)
	fmt.Println(encoded64)
}

func challenge2() {
	var line1 string = "1c0111001f010100061a024b53535009181c"
	var line2 string = "686974207468652062756c6c277320657965"
	bytes1, _ := hex.DecodeString(line1)
	bytes2, _ := hex.DecodeString(line2)
	length := len(bytes1)
	bytes3 := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes3[i] = bytes1[i] ^ bytes2[i]
	}
	fmt.Printf("%x", bytes3)
}

func scoreEnglishBytes(bytes []byte) float64 {
	checkAgainst := LETTERS + " "
	score := 0
	for i := 0; i < len(bytes); i++ {
		for j := 0; j < len(checkAgainst); j++ {
			if bytes[i] == checkAgainst[j] {
				// fmt.Printf("%c,", bytes[i])
				score++
				break
			}
		}
	}
	// fmt.Println()
	return float64(score) / float64(len(bytes))
}

func challenge3() {
	var line string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bytes, _ := hex.DecodeString(line)
	for i := 0; i < len(LETTERS); i++ {
		decBytes := make([]byte, len(bytes))
		for j := 0; j < len(decBytes); j++ {
			decBytes[j] = bytes[j] ^ byte(LETTERS[i])
		}
		fmt.Printf("%c: %s\n", LETTERS[i], decBytes)
	}
}

func challenge4() {
	f, _ := os.Open("4.txt")
	scanner := bufio.NewScanner(f)
	lineIdx := 0

	var scores [][]float64
	for scanner.Scan() {
		scores = append(scores, make([]float64, len(LETTERS)))
		line := scanner.Text()
		bytes, _ := hex.DecodeString(line)
		for keyIdx := 0; keyIdx < len(LETTERS); keyIdx++ {
			decBytes := make([]byte, len(bytes))
			for j := 0; j < len(decBytes); j++ {
				decBytes[j] = bytes[j] ^ byte(LETTERS[keyIdx])
			}
			scores[lineIdx][keyIdx] = scoreEnglishBytes(decBytes)
			fmt.Printf("Line: %d - Key: %c - score: %f : %s\n", lineIdx, LETTERS[keyIdx], scores[lineIdx][keyIdx], decBytes)
		}
		lineIdx++
	}

	maxLineIdx := 0
	maxKeyIdx := 0
	maxVal := 0.0
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			if maxVal < scores[i][j] {
				maxLineIdx = i
				maxKeyIdx = j
				maxVal = scores[i][j]
			}
		}

	}
	fmt.Printf("BEST line: %d - key: %c - score: %f\n", maxLineIdx, LETTERS[maxKeyIdx], maxVal)
}

func challenge5() {
	line := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	encBytes := make([]byte, len(line))
	key := "ICE"
	for i := 0; i < len(line); i++ {
		encBytes[i] = key[i%len(key)] ^ line[i]
	}
	fmt.Printf("%x\n", encBytes)
}
