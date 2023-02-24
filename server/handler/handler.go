package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handle() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"172.0.0.1"})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

raw_text = "Hello this is my string"

import nltk
from nltk.tokenize import sent_tokenize

//Sentence tokenization
corpus = sent_tokenize(raw_texts)

step_2_corpus = []
for sentence in corpus
step_2_corpus.append(sentence.lower())

import re
step_3_corpus = []
for sentence in step_2_corpus:
	//Remove numbers - Locate all numbers and substitute with nothing
	sentence = re.sub('[0-9]', '', sentence) 
	step_3_corpus.append(sentence)

// Step 4: De-concatination. Consider every situation
//Very important, public models don't recognize contractions
step_4_corpus = []
for sentence in step_3_corpus:
	sentence = re.sub(r"won\'t", "will not", sentence)
	sentence = re.sub(r"can\'t", "can not", sentence)
	sentence = re.sub(r"\'t", " not", sentence)
	sentence = re.sub(r"\'re", " are", sentence)
	sentence = re.sub(r"\'s", " is", sentence)
	sentence = re.sub(r"\'d", " would", sentence)
	sentence = re.sub(r"\'ll", " will", sentence)
	sentence = re.sub(r"\'t", " not", sentence)
	sentence = re.sub(r"\'ve", " have", sentence)
	sentence = re.sub(r"\'m", " am", sentence)

// Step 5: Remove punctuation, special characters, and symbols
step_5_corpus = []
for sentence in step_4_corpus:
	sentence = re.sub("[^a-z0-9<>]", ' ', sentence)
	step_5_corpus.append(sentence)

	




