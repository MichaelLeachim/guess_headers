# This software is Work In Progress at this moment

## Installation

```shell
go get -u <TODO>
```


## There are two purposes behind this software:

* To guess and match field names based on the data of these fields
* To fuzzy join data from two different sources (also known as [record linkage](https://en.wikipedia.org/wiki/Record_linkage))

Imagine having a hundred files with 150 fields (headers) in each one,
where headers differ, some of them slightly, some of them a lot. 
You have to figure correspondence between them and unify the data. 
This is one of the boring tasks that you might have. 

So, I decided to automate the task. 

The second purpose is result of working on the first goal. 
Record linkage is the same problem, you just have to transpose
matrix. So columns become rows, and the same algorithm applies, 
with several tweaks. 

### There are several scoring algorithms that I am working on

Right now these are implemented:

* Simple match ([Jaccard distance](https://en.wikipedia.org/wiki/Jaccard_index) between columns, compare function is simple equality)
* [TODO] Bayessian column match 
* [TODO] TFIDF match
* [TODO] Simple match + Cell2Cell(levenshtein/hamming)

##### Bayessian column match

Intuitive way to remember Bayes' theorem in context of bag of words model is:

`P(Translation|Word) = (P(Word|Translation) * P(Translation))/P(Word)`

* Probability of a word, given translation, i.e. 1/(how many words have this translation)
* Probability of a word:      (this word appears times)/(all words appears)
* Probability of translation: (this translation word appears)/(all translation words appear )


#### All input data is transformed by these rules, in order

* [TODO] Lowercased
* [TODO] Transliterated into ASCII
* [TODO] Alphanumeric only
* [TODO] Positional number tokenized (Turn 1923 into this string: "1900 900 20 3")












#### For each cell (field) full column metadata is applied

* [TODO] allnum    (means, the whole column is numeric)
* [TODO] allstring (means, the whole column is string)
* [TODO] allsame   (means, the whole column is the same)
* [TODO] maxN      (means, the column data has N different entries)

### Limiters

For the purpose of the headers match, most of the time, you don't need to perform
full comparison between whole columns. Especially if the dataset is large. 
Most of the time, random sample of ~100 cells per column is enough. 

The algorithm of matching is not efficient and, for larger datasets, it may take 
longer time to match data if no sampling is applied. 

On the other hand, when we are talking about record linkage, no sampling can take place. 
In this case, the best way to match records is to reduce input. 

But, on smaller data sets (up to ~1 000 000 records), usually, no input reducing is 
needed (but prepare to wait for several days)

Hunting for the bottlenecks is an ongoing task [TODO]

```shell
[TODO] Performance metric for two 100 000 rows files
```

## Usage examples

```shell


```









# TODO

## [2019-02-19] [19:53]


## [2019-02-18] [21:10]
* [DONE] Make **guess** function
* Implement test for **guess** function
* Implement test on small size real life data
* 
* [DONE] Create transformer for CSV  format
* Create transformer for JSON format
* [DONE] Create transformer for XLSX format
