package main

/*
func InitReadProbabilities(fname string, tax *Taxonomy) ([]ProbabilityDistribution, error) {
	var ps []ProbabilityDistribution

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal()
	}
	defer f.Close()

	labels := tax.labels
	nLabels := len(tax.labels)

	uniformDistributin := make([]float64, nLabels)
	for i := range uniformDistributin {
		uniformDistributin[i] = 1.0 / float64(nLabels)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ps = append(ps, ProbabilityDistribution{scanner.Text(), labels, uniformDistributin})
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ps, err
}
*/
/*
func storeTax(tax *Taxonomy) {
	db, err := leveldb.OpenFile("db/taxonomy", nil)
	if err != nil {
		log.Fatal("cannot open db file for write")
	}
	for
}
*/

func main() {
	/*
		iab := new(Taxonomy)
		iab.Level1 = make(kv)
		iab.Level2 = make(kv)
		iab.Level3 = make(kv)

		fname := os.Args[1]
		fmt.Printf("This is the command line %q\n", fname)
		iab.Init(fname)
	*/
	/*
		length, _ := InitReadProbabilities("names.txt", iab)
		fmt.Printf("let's curate %v entities", length)
		// Now get a list of subreddits.

		// key idea. ask the question by sampling a set.
		// classify that with a question.
		// write the results back to DB.

		// concept: candidates. exploit explore.
		// a list of entities to be classified.
	*/
}
