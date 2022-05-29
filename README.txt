Instruction to test the working of word count application. This application will provide a response to the top 10 most occuring words in a file.

http://localhost:8000/wordcounts


curl --location --request POST 'http://localhost:8000/wordcounts' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'textstring=This is with reference text to test the working of word count. Some words are repeated to increase word count. Happy count'